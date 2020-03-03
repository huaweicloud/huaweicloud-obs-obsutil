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
package assist

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var homePath string

func Home() (string, error) {
	if homePath != "" {
		return homePath, nil
	}

	h, err := home()
	if err == nil {
		homePath = h
		return homePath, nil
	}
	return "", err
}

func home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	if "windows" == runtime.GOOS {
		return getWindowsHome()
	}

	return getUnixHome()
}

func getUnixHome() (ret string, err error) {
	ret = strings.TrimSpace(os.Getenv("HOME"))
	if ret != "" {
		return
	}

	evalCommand := exec.Command("sh", "-c", "eval echo ~$USER")
	evalCommand.Stdout = new(bytes.Buffer)
	err = evalCommand.Run()
	if err != nil {
		return
	}

	output, ok := evalCommand.Stdout.(*bytes.Buffer)
	if ok {
		ret = strings.TrimSpace(output.String())
	}

	if ret == "" {
		err = errors.New("Get blank output after exectuing command to read home directory")
		return
	}

	return
}

func getWindowsHome() (ret string, err error) {
	homePath := strings.TrimSpace(os.Getenv("HOMEPATH"))
	homeDrive := strings.TrimSpace(os.Getenv("HOMEDRIVE"))

	if homePath == "" || homeDrive == "" {
		ret = strings.TrimSpace(os.Getenv("USERPROFILE"))
	} else {
		ret = homeDrive + homePath
	}

	if ret == "" {
		err = errors.New("Either HOMEDRIVE, HOMEPATH, or USERPROFILE is blank")
		return
	}

	return
}
