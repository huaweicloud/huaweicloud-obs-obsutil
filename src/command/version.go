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
	"command/i18n"
	"obs"
)

const obsUtilVersion = "5.1.13"

func initVersion() command {

	c := &defaultCommand{
		key:         "version",
		description: "show version",
		additional:  true,
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) > 0 {
			c.showHelp()
			printf("Error: Invalid args: %v", args)
			return assist.ErrInvalidArgs
		}
		printf("obsutil version:%s, obssdk version:%s", obsUtilVersion, obs.OBS_SDK_VERSION)
		printf("operating system:%s, arch:%s", assist.GetOS(), assist.GetArch())
		return nil
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("show the version of this tool"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil version")
		printf("")
	}

	return c
}
