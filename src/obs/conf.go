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
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type urlHolder struct {
	scheme string
	host   string
	port   int
}

type config struct {
	securityProviders  []SecurityProvider
	urlHolder          *urlHolder
	endpoint           string
	signature          SignatureType
	pathStyle          bool
	region             string
	connectTimeout     int
	socketTimeout      int
	headerTimeout      int
	idleConnTimeout    int
	finalTimeout       int
	maxRetryCount      int
	proxyUrl           string
	maxConnsPerHost    int
	sslVerify          bool
	pemCerts           []byte
	userAgent          string
	disableCompression bool
}

func (conf config) String() string {
	return fmt.Sprintf("[endpoint:%s, signature:%s, pathStyle:%v, region:%s"+
		"\nconnectTimeout:%d, socketTimeout:%dheaderTimeout:%d, idleConnTimeout:%d"+
		"\nmaxRetryCount:%d, maxConnsPerHost:%d, sslVerify:%v, proxyUrl:%s]",
		conf.endpoint, conf.signature, conf.pathStyle, conf.region,
		conf.connectTimeout, conf.socketTimeout, conf.headerTimeout, conf.idleConnTimeout,
		conf.maxRetryCount, conf.maxConnsPerHost, conf.sslVerify, conf.proxyUrl,
	)
}

type configurer func(conf *config)

func WithSecurityProviders(sps ...SecurityProvider) configurer {
	return func(conf *config) {
		for _, sp := range sps {
			if sp != nil {
				conf.securityProviders = append(conf.securityProviders, sp)
			}
		}
	}
}

func WithSslVerify(sslVerify bool) configurer {
	return WithSslVerifyAndPemCerts(sslVerify, nil)
}

func WithSslVerifyAndPemCerts(sslVerify bool, pemCerts []byte) configurer {
	return func(conf *config) {
		conf.sslVerify = sslVerify
		conf.pemCerts = pemCerts
	}
}

func WithHeaderTimeout(headerTimeout int) configurer {
	return func(conf *config) {
		conf.headerTimeout = headerTimeout
	}
}

func WithProxyUrl(proxyUrl string) configurer {
	return func(conf *config) {
		conf.proxyUrl = proxyUrl
	}
}

func WithMaxConnections(maxConnsPerHost int) configurer {
	return func(conf *config) {
		conf.maxConnsPerHost = maxConnsPerHost
	}
}

func WithPathStyle(pathStyle bool) configurer {
	return func(conf *config) {
		conf.pathStyle = pathStyle
	}
}

func WithSignature(signature SignatureType) configurer {
	return func(conf *config) {
		conf.signature = signature
	}
}

func WithRegion(region string) configurer {
	return func(conf *config) {
		conf.region = region
	}
}

func WithConnectTimeout(connectTimeout int) configurer {
	return func(conf *config) {
		conf.connectTimeout = connectTimeout
	}
}

func WithSocketTimeout(socketTimeout int) configurer {
	return func(conf *config) {
		conf.socketTimeout = socketTimeout
	}
}

func WithIdleConnTimeout(idleConnTimeout int) configurer {
	return func(conf *config) {
		conf.idleConnTimeout = idleConnTimeout
	}
}

func WithMaxRetryCount(maxRetryCount int) configurer {
	return func(conf *config) {
		conf.maxRetryCount = maxRetryCount
	}
}

func WithSecurityToken(securityToken string) configurer {
	return func(conf *config) {
		if len(conf.securityProviders) == 1 {
			if bsp, ok := conf.securityProviders[0].(*BasicSecurityProvider); ok {
				sh := bsp.GetSecurity()
				bsp.Refresh(sh.AK, sh.SK, securityToken)
			}
		}
	}
}

func WithUserAgent(userAgent string) configurer {
	return func(conf *config) {
		conf.userAgent = userAgent
	}
}

func WithDisableCompression(disableCompression bool) configurer {
	return func(conf *config) {
		conf.disableCompression = disableCompression
	}
}

const c_https = "https"
const c_http = "http"

func (conf *config) initConfigWithDefault() error {
	conf.endpoint = strings.TrimSpace(conf.endpoint)
	if conf.endpoint == "" {
		return errors.New("endpoint is not set")
	}

	if index := strings.Index(conf.endpoint, "?"); index > 0 {
		conf.endpoint = conf.endpoint[:index]
	}

	for strings.LastIndex(conf.endpoint, "/") == len(conf.endpoint)-1 {
		conf.endpoint = conf.endpoint[:len(conf.endpoint)-1]
	}

	if conf.signature == "" {
		conf.signature = DEFAULT_SIGNATURE
	}

	urlHolder := &urlHolder{}
	var address string
	if strings.HasPrefix(conf.endpoint, "https://") {
		urlHolder.scheme = c_https
		address = conf.endpoint[len("https://"):]
	} else if strings.HasPrefix(conf.endpoint, "http://") {
		urlHolder.scheme = c_http
		address = conf.endpoint[len("http://"):]
	} else {
		urlHolder.scheme = c_https
		address = conf.endpoint
	}

	addr := strings.Split(address, ":")
	if len(addr) == 2 {
		if port, err := strconv.Atoi(addr[1]); err == nil {
			urlHolder.port = port
		}
	}
	urlHolder.host = addr[0]
	if urlHolder.port == 0 {
		if urlHolder.scheme == c_https {
			urlHolder.port = 443
		} else {
			urlHolder.port = 80
		}
	}

	if IsIP(urlHolder.host) {
		conf.pathStyle = true
	}

	conf.urlHolder = urlHolder

	conf.region = strings.TrimSpace(conf.region)
	if conf.region == "" {
		conf.region = DEFAULT_REGION
	}

	if conf.connectTimeout <= 0 {
		conf.connectTimeout = DEFAULT_CONNECT_TIMEOUT
	}

	if conf.socketTimeout <= 0 {
		conf.socketTimeout = DEFAULT_SOCKET_TIMEOUT
	}

	conf.finalTimeout = conf.socketTimeout * 10

	if conf.headerTimeout <= 0 {
		conf.headerTimeout = DEFAULT_HEADER_TIMEOUT
	}

	if conf.idleConnTimeout < 0 {
		conf.idleConnTimeout = DEFAULT_IDLE_CONN_TIMEOUT
	}

	if conf.maxRetryCount < 0 {
		conf.maxRetryCount = DEFAULT_MAX_RETRY_COUNT
	}

	if conf.maxConnsPerHost <= 0 {
		conf.maxConnsPerHost = DEFAULT_MAX_CONN_PER_HOST
	}

	conf.proxyUrl = strings.TrimSpace(conf.proxyUrl)
	return nil
}

func (conf *config) getTransport() (*http.Transport, error) {
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			start := GetCurrentTimestamp()
			conn, err := (&net.Dialer{
				Timeout:  time.Second * time.Duration(conf.connectTimeout),
				Resolver: net.DefaultResolver,
			}).Dial(network, addr)

			if isInfoLogEnabled() {
				doLog(LEVEL_INFO, "Do http dial cost %d ms", (GetCurrentTimestamp() - start))
			}
			if err != nil {
				return nil, err
			}
			return getConnDelegate(conn, conf.socketTimeout, conf.finalTimeout), nil
		},
		MaxIdleConns:          conf.maxConnsPerHost,
		MaxIdleConnsPerHost:   conf.maxConnsPerHost,
		ResponseHeaderTimeout: time.Second * time.Duration(conf.headerTimeout),
		IdleConnTimeout:       time.Second * time.Duration(conf.idleConnTimeout),
		DisableCompression:    conf.disableCompression,
	}

	if conf.proxyUrl != "" {
		if !strings.HasPrefix(conf.proxyUrl, "http://") && !strings.HasPrefix(conf.proxyUrl, "https://") {
			conf.proxyUrl = fmt.Sprintf("%s%s", "http://", conf.proxyUrl)
		}

		var proxyUrl *url.URL
		var parseErr error
		userInfoGroups := regexp.MustCompile(".*?//(.+?)@.*?").FindStringSubmatch(conf.proxyUrl)
		if len(userInfoGroups) <= 1 {
			proxyUrl, parseErr = url.Parse(conf.proxyUrl)
		} else {
			proxyUrl, parseErr = url.Parse(strings.Replace(conf.proxyUrl, userInfoGroups[1]+"@", "", 1))
			if index := strings.Index(userInfoGroups[1], ":"); index >= 0 {
				proxyUrl.User = url.UserPassword(userInfoGroups[1][0:index], userInfoGroups[1][index+1:])
			} else {
				proxyUrl.User = url.User(userInfoGroups[1])
			}

		}
		if parseErr != nil {
			return nil, parseErr
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: !conf.sslVerify}
	if conf.sslVerify && conf.pemCerts != nil {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(conf.pemCerts)
		tlsConfig.RootCAs = pool
	}
	transport.TLSClientConfig = tlsConfig

	return transport, nil
}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func DummyQueryEscape(s string) string {
	return s
}

func (conf *config) formatUrls(bucketName, objectKey string, params map[string]string, escape bool) (requestUrl string, canonicalizedUrl string) {

	urlHolder := conf.urlHolder

	if bucketName == "" {
		requestUrl = fmt.Sprintf("%s://%s:%d", urlHolder.scheme, urlHolder.host, urlHolder.port)
		canonicalizedUrl = "/"
	} else {
		if conf.pathStyle {
			requestUrl = fmt.Sprintf("%s://%s:%d/%s", urlHolder.scheme, urlHolder.host, urlHolder.port, bucketName)
			canonicalizedUrl = "/" + bucketName
		} else {
			requestUrl = fmt.Sprintf("%s://%s.%s:%d", urlHolder.scheme, bucketName, urlHolder.host, urlHolder.port)
			if conf.signature == "v2" {
				canonicalizedUrl = "/" + bucketName + "/"
			} else {
				canonicalizedUrl = "/"
			}
		}
	}
	var escapeFunc func(s string) string
	if escape {
		escapeFunc = url.QueryEscape
	} else {
		escapeFunc = DummyQueryEscape
	}

	if objectKey != "" {
		encodeObjectKey := escapeFunc(objectKey)
		requestUrl += "/" + encodeObjectKey
		if !strings.HasSuffix(canonicalizedUrl, "/") {
			canonicalizedUrl += "/"
		}
		canonicalizedUrl += encodeObjectKey
	}

	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, strings.TrimSpace(key))
	}
	sort.Strings(keys)
	i := 0

	for index, key := range keys {
		if index == 0 {
			requestUrl += "?"
		} else {
			requestUrl += "&"
		}
		_key := url.QueryEscape(key)
		requestUrl += _key

		_value := params[key]

		if conf.signature == "v4" {
			requestUrl += "=" + url.QueryEscape(_value)
		} else {
			if _value != "" {
				requestUrl += "=" + url.QueryEscape(_value)
				_value = "=" + _value
			} else {
				_value = ""
			}
			lowerKey := strings.ToLower(key)
			_, ok := allowedResourceParameterNames[lowerKey]
			ok = ok || strings.HasPrefix(lowerKey, HEADER_PREFIX)
			if ok {
				if i == 0 {
					canonicalizedUrl += "?"
				} else {
					canonicalizedUrl += "&"
				}
				canonicalizedUrl += getQueryUrl(_key, _value)
				i++
			}
		}
	}
	return
}

func getQueryUrl(key, value string) string {
	queryUrl := ""
	queryUrl += key
	queryUrl += value
	return queryUrl
}
