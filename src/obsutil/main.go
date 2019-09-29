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
package main

import (
	"assist"
	"command"
	"runtime"
)

var (
	AesKey            = ""
	AesIv             = ""
	CloudType         = ""
	AesShareKeyPrefix = ""
	AesShareIv        = ""
	ShareConsole      = ""
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	assist.SetCloudType(CloudType)
	command.SetAesKeyAndIv(AesKey, AesIv)
	command.SetAesShareIv(AesShareKeyPrefix, AesShareIv)
	command.SetShareConsoleUrl(ShareConsole)
	command.Run()
}
