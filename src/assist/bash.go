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
	"bufio"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func EnterBashMode(additionalTips func(), callback func(value string)) {
	fmt.Println("Enter \"exit\" or \"quit\" to logout")
	if additionalTips != nil {
		additionalTips()
	}

	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input your command:")
		fmt.Printf("%s", "-->")
		input, err := ReadLine(rd)
		if err == nil {
			command := strings.TrimSpace(BytesToString(input))
			if command != "" {
				if strings.ToLower(command) == "exit" || strings.ToLower(command) == "quit" {
					return
				}
				callback(command)
				fmt.Println()
				continue
			}
		}
		if err == nil || err != io.EOF {
			fmt.Println("Error: Invalid input, please try again")
		}
	}
}

func PreprocessInput(input string) (string, string, bool, error) {
	_placeHolder, err := uuid.NewV4()
	if err != nil {
		return "", "", false, err
	}
	placeHolder := _placeHolder.String()
	length := utf8.RuneCountInString(input)
	temp := make([]string, 0, length)
	enter := false
	flag := false
	for _, s := range input {
		_s := string(s)
		if _s == "\"" {
			enter = !enter
		} else if _s == " " && enter {
			_s = placeHolder
			flag = true
		}

		temp = append(temp, _s)
	}

	return strings.Join(temp, ""), placeHolder, flag, nil
}
