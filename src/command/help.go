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
)

func initHelp() command {

	c := &defaultCommand{
		key:         "help",
		usage:       "[command]",
		description: "view command help information",
		additional:  true,
	}

	c.action = func() error {
		args := c.flagSet.Args()
		length := len(args)
		if length <= 0 {
			usage()
			return nil
		}
		if len(args) > 1 {
			c.showHelp()
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}
		if _command, ok := commands[args[0]]; ok && _command.getHelp() != nil {
			_command.getHelp()()
			return nil
		}
		printf("Error: No such command: \"%s\", please try \"help\" for more information!", args[0])
		return assist.ErrInvalidArgs
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("view the commands supported by this tool or the help information of a specific command"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil help [command]")
		printf("")
	}

	return c
}
