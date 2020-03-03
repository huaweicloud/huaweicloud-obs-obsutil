// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.
package obs

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ACCESS_KEY_ENV     = "OBS_ACCESS_KEY_ID"
	SECURITY_KEY_ENV   = "OBS_SECRET_ACCESS_KEY"
	SECURITY_TOKEN_ENV = "OBS_SECURITY_TOKEN"
	ECS_REQUEST_URL    = "http://169.254.169.254/openstack/latest/securitykey"
)

type SecurityHolder struct {
	AK            string
	SK            string
	SecurityToken string
}

var emptySecurityHolder = SecurityHolder{}

type SecurityProvider interface {
	GetSecurity() SecurityHolder
}

type BasicSecurityProvider struct {
	val atomic.Value
}

func (bsp *BasicSecurityProvider) GetSecurity() SecurityHolder {
	if sh, ok := bsp.val.Load().(SecurityHolder); ok {
		return sh
	}
	return emptySecurityHolder
}

func (bsp *BasicSecurityProvider) Refresh(ak, sk, securityToken string) {
	bsp.val.Store(SecurityHolder{AK: strings.TrimSpace(ak), SK: strings.TrimSpace(sk), SecurityToken: strings.TrimSpace(securityToken)})
}

func NewBasicSecurityProvider(ak, sk, securityToken string) *BasicSecurityProvider {
	bsp := &BasicSecurityProvider{}
	bsp.Refresh(ak, sk, securityToken)
	return bsp
}

type EnvSecurityProvider struct {
	sh     SecurityHolder
	suffix string
	once   sync.Once
}

func (esp *EnvSecurityProvider) GetSecurity() SecurityHolder {
	//ensure run only once
	esp.once.Do(func() {
		esp.sh = SecurityHolder{
			AK:            strings.TrimSpace(os.Getenv(ACCESS_KEY_ENV + esp.suffix)),
			SK:            strings.TrimSpace(os.Getenv(SECURITY_KEY_ENV + esp.suffix)),
			SecurityToken: strings.TrimSpace(os.Getenv(SECURITY_TOKEN_ENV + esp.suffix)),
		}
	})

	return esp.sh
}

func NewEnvSecurityProvider(suffix string) *EnvSecurityProvider {
	suffix = strings.TrimSpace(strings.ToUpper(suffix))
	if suffix != "" {
		suffix = "_" + suffix
	}
	esp := &EnvSecurityProvider{
		suffix: suffix,
	}
	return esp
}

type TemporarySecurityHolder struct {
	SecurityHolder
	ExpireDate time.Time
}

var emptyTemporarySecurityHolder = TemporarySecurityHolder{}

type EcsSecurityProvider struct {
	val        atomic.Value
	lock       sync.Mutex
	httpClient *http.Client
	prefetch   int32
	retryCount int
}

func (ecsSp *EcsSecurityProvider) loadTemporarySecurityHolder() (TemporarySecurityHolder, bool) {
	if sh := ecsSp.val.Load(); sh == nil {
		return emptyTemporarySecurityHolder, false
	} else if _sh, ok := sh.(TemporarySecurityHolder); !ok {
		return emptyTemporarySecurityHolder, false
	} else {
		return _sh, true
	}
}

func (ecsSp *EcsSecurityProvider) getAndSetSecurityWithOutLock() SecurityHolder {
	_sh := TemporarySecurityHolder{}
	_sh.ExpireDate = time.Now().Add(time.Minute * 5)
	retryCount := 0
	for {
		if req, err := http.NewRequest("GET", ECS_REQUEST_URL, nil); err == nil {
			start := GetCurrentTimestamp()
			res, err := ecsSp.httpClient.Do(req)
			if err == nil {
				if data, _err := ioutil.ReadAll(res.Body); _err == nil {
					temp := &struct {
						Credential struct {
							AK            string    `json:"access,omitempty"`
							SK            string    `json:"secret,omitempty"`
							SecurityToken string    `json:"securitytoken,omitempty"`
							ExpireDate    time.Time `json:"expires_at,omitempty"`
						} `json:"credential"`
					}{}

					doLog(LEVEL_DEBUG, "Get the json data from ecs succeed, data %s", data)

					if jsonErr := json.Unmarshal(data, temp); jsonErr == nil {
						_sh.AK = temp.Credential.AK
						_sh.SK = temp.Credential.SK
						_sh.SecurityToken = temp.Credential.SecurityToken
						_sh.ExpireDate = temp.Credential.ExpireDate.Add(time.Minute * -1)

						doLog(LEVEL_INFO, "Get security from ecs succeed, AK:xxxx, SK:xxxx, SecurityToken:xxxx, ExprireDate %s", _sh.ExpireDate)

						doLog(LEVEL_INFO, "Get security from ecs succeed, cost %d ms", (GetCurrentTimestamp() - start))
						break
					} else {
						err = jsonErr
					}
				} else {
					err = _err
				}
			}

			doLog(LEVEL_WARN, "Try to get security from ecs failed, cost %d ms, err %s", (GetCurrentTimestamp() - start), err.Error())
		}

		if retryCount > ecsSp.retryCount {
			doLog(LEVEL_WARN, "Try to get security from ecs failed and exceed the max retry count")
			break
		}
		sleepTime := float64(retryCount+2) * rand.Float64()
		if sleepTime > 10 {
			sleepTime = 10
		}
		time.Sleep(time.Duration(sleepTime * float64(time.Second)))
		retryCount++
	}

	ecsSp.val.Store(_sh)
	return _sh.SecurityHolder
}

func (ecsSp *EcsSecurityProvider) getAndSetSecurity() SecurityHolder {
	ecsSp.lock.Lock()
	defer ecsSp.lock.Unlock()
	tsh, succeed := ecsSp.loadTemporarySecurityHolder()
	if !succeed {
		return ecsSp.getAndSetSecurityWithOutLock()
	}
	return tsh.SecurityHolder
}

func (ecsSp *EcsSecurityProvider) GetSecurity() SecurityHolder {
	if tsh, succeed := ecsSp.loadTemporarySecurityHolder(); succeed {
		if time.Now().Before(tsh.ExpireDate) {
			//not expire
			if time.Now().Add(time.Minute*5).After(tsh.ExpireDate) && atomic.CompareAndSwapInt32(&ecsSp.prefetch, 0, 1) {
				//do prefetch
				sh := ecsSp.getAndSetSecurityWithOutLock()
				atomic.CompareAndSwapInt32(&ecsSp.prefetch, 1, 0)
				return sh
			}
			return tsh.SecurityHolder
		}
		return ecsSp.getAndSetSecurity()
	}

	return ecsSp.getAndSetSecurity()
}

func getInternalTransport() *http.Transport {

	timeout := 10
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			start := GetCurrentTimestamp()
			conn, err := (&net.Dialer{
				Timeout:  time.Second * time.Duration(timeout),
				Resolver: net.DefaultResolver,
			}).Dial(network, addr)

			if isInfoLogEnabled() {
				doLog(LEVEL_INFO, "Do http dial cost %d ms", (GetCurrentTimestamp() - start))
			}
			if err != nil {
				return nil, err
			}
			return getConnDelegate(conn, timeout, timeout*10), nil
		},
		MaxIdleConns:          10,
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		IdleConnTimeout:       time.Second * time.Duration(DEFAULT_IDLE_CONN_TIMEOUT),
		DisableCompression:    true,
	}

	return transport
}

func NewEcsSecurityProvider(retryCount int) *EcsSecurityProvider {
	ecsSp := &EcsSecurityProvider{
		retryCount: retryCount,
	}
	ecsSp.httpClient = &http.Client{Transport: getInternalTransport(), CheckRedirect: checkRedirectFunc}
	return ecsSp
}
