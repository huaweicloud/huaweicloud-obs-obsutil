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
package command

import (
	"assist"
	"obs"
	"os"
	"strings"
)

var obsClient *obs.ObsClient
var obsClientCrr *obs.ObsClient

func refreshObsClient(autoChooseSecurityProvider bool) (err error) {
	if obsClient != nil {
		obsClient.Close()
	}
	obsClient, err = createObsClient("", autoChooseSecurityProvider)
	return
}

func refreshObsClientCrr(autoChooseSecurityProvider bool) (err error) {
	if obsClientCrr != nil {
		obsClientCrr.Close()
	}
	obsClientCrr, err = createObsClient("Crr", autoChooseSecurityProvider)
	return
}

func createObsClient(suffix string, autoChooseSecurityProvider bool) (*obs.ObsClient, error) {
	connectTimeout := assist.StringToInt(config["connectTimeout"], defaultConnectTimeout)
	socketTimeout := assist.StringToInt(config["socketTimeout"], defaultSocketTimeout)
	maxRetryCount := assist.StringToInt(config["maxRetryCount"], defaultMaxRetryCount)
	maxConnections := assist.StringToInt(config["maxConnections"], defaultMaxConnections)

	endpoint := strings.TrimSpace(config["endpoint"+suffix])
	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	proxyUrl := strings.TrimSpace(config["proxyUrl"])
	if proxyUrl == "" {
		proxyUrl = strings.TrimSpace(os.Getenv("HTTPS_PROXY"))
		if proxyUrl == "" {
			proxyUrl = strings.TrimSpace(os.Getenv("HTTP_PROXY"))
		}
	}

	ak := config["ak"+suffix]
	sk := config["sk"+suffix]
	if ak == defaultAccessKey {
		ak = ""
	}

	if sk == defaultSecurityKey {
		sk = ""
	}

	if autoChooseSecurityProvider {
		return obs.New(ak, sk, endpoint,
			obs.WithSecurityToken(config["token"+suffix]),
			obs.WithConnectTimeout(connectTimeout),
			obs.WithSocketTimeout(socketTimeout),
			obs.WithHeaderTimeout(socketTimeout),
			obs.WithMaxRetryCount(maxRetryCount),
			obs.WithMaxConnections(maxConnections),
			obs.WithProxyUrl(proxyUrl),
			obs.WithUserAgent("obsutil/"+obsUtilVersion),
			obs.WithDisableCompression(config["enableCompression"] != c_true),
			obs.WithSecurityProviders(obs.NewEnvSecurityProvider(suffix), obs.NewEcsSecurityProvider(1)),
		)
	}

	return obs.New(ak, sk, endpoint,
		obs.WithSecurityToken(config["token"+suffix]),
		obs.WithConnectTimeout(connectTimeout),
		obs.WithSocketTimeout(socketTimeout),
		obs.WithHeaderTimeout(socketTimeout),
		obs.WithMaxRetryCount(maxRetryCount),
		obs.WithMaxConnections(maxConnections),
		obs.WithProxyUrl(proxyUrl),
		obs.WithUserAgent("obsutil/"+obsUtilVersion),
		obs.WithDisableCompression(config["enableCompression"] != c_true),
	)
}

func initClientAndLog(autoChooseSecurityProvider bool) bool {
	var err error

	err = refreshObsClient(autoChooseSecurityProvider)
	if err != nil {
		printError(err)
		return false
	}

	helper := assist.MapHelper(config)

	if sdkLogPath := helper.Get("sdkLogPath"); sdkLogPath != "" {
		sdkLogBackups := assist.StringToInt(config["sdkLogBackups"], defaultLogBackups)
		var sdkMaxLogSize int64
		if _sdkMaxLogSize, err := assist.TranslateToInt64(config["sdkMaxLogSize"]); err == nil && _sdkMaxLogSize > 0 {
			sdkMaxLogSize = _sdkMaxLogSize
		} else {
			sdkMaxLogSize = defaultMaxLogSize
		}
		if err := obs.InitLog(sdkLogPath, sdkMaxLogSize, sdkLogBackups, obs.Level(transLogLevel(config["sdkLogLevel"])), false); err != nil {
			printError(err)
		}
	}

	if utilLogPath := helper.Get("utilLogPath"); utilLogPath != "" {
		utilLogBackups := assist.StringToInt(config["utilLogBackups"], defaultLogBackups)
		var utilMaxLogSize int64
		if _utilMaxLogSize, err := assist.TranslateToInt64(config["utilMaxLogSize"]); err == nil && _utilMaxLogSize > 0 {
			utilMaxLogSize = _utilMaxLogSize
		} else {
			utilMaxLogSize = defaultMaxLogSize
		}
		if err := initLog(utilLogPath, utilMaxLogSize, utilLogBackups, Level(transLogLevel(config["utilLogLevel"])), defaultLogCacheCnt); err != nil {
			printError(err)
		}
	}

	return true
}

func doClean() {
	if obsClient != nil {
		obsClient.Close()
		obsClient = nil
	}

	if obsClientCrr != nil {
		obsClientCrr.Close()
		obsClientCrr = nil
	}

	obs.CloseLog()
	closeLog()
}
