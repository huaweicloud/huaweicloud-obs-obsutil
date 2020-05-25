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
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"progress"
	"regexp"
	"runtime/debug"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

var commands = make(map[string]command)
var basic []command
var other []command
var basicCommandCnt int
var otherCommandCnt int
var terminalWidth int
var runningRound int

type NilWriter struct {
}

func (NilWriter) Write(p []byte) (n int, err error) {
	if p != nil {
		return len(p), nil
	}
	return 0, nil
}

var nilWriter = &NilWriter{}
var splitRegex = regexp.MustCompile("\\s+")

func initFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("obsutil", flag.ContinueOnError)
	flagSet.SetOutput(nilWriter)
	return flagSet
}

func register(c command) {
	if c != nil {
		commands[c.getKey()] = c
		if c.getAdditional() {
			otherCommandCnt++
		} else {
			basicCommandCnt++
		}
	}
}

func initCommands() {
	commandFuncs := []func() command{
		initVersion,
		initHelp,
		initConfig,
		initLs,
		initMb,
		initRm,
		initStat,
		initCp,
		initSync,
		initRestore,
		initAbort,
		initClear,
		initChattri,
		initArchive,
		initMkdir,
		initSign,
		initMv,
		initCat,
		initHash,
		initDirectDownload,
		initBucketPolicyCommand,
	}

	if assist.IsHec() {
		commandFuncs = append(commandFuncs, initShareCrt)
		commandFuncs = append(commandFuncs, initShareLs)
		commandFuncs = append(commandFuncs, initShareCp)
	} else {
		//dt
		delete(defaultConfigMap, "autoChooseSecurityProvider")
		newDefaultConfigSlice := defaultConfigSlice[:0]
		for _, key := range defaultConfigSlice {
			if key != "autoChooseSecurityProvider" {
				newDefaultConfigSlice = append(newDefaultConfigSlice, key)
			}
		}
		defaultConfigSlice = newDefaultConfigSlice
	}

	for _, commandFunc := range commandFuncs {
		register(commandFunc())
	}

	basicKeys := make([]string, 0, basicCommandCnt)
	otherKeys := make([]string, 0, otherCommandCnt)

	for k, c := range commands {
		if c.getAdditional() {
			otherKeys = append(otherKeys, k)
		} else {
			basicKeys = append(basicKeys, k)
		}
	}

	sort.Strings(basicKeys)
	sort.Strings(otherKeys)

	basic = make([]command, 0, len(basicKeys))
	other = make([]command, 0, len(otherKeys))

	for _, k := range basicKeys {
		basic = append(basic, commands[k])
	}

	for _, k := range otherKeys {
		other = append(other, commands[k])
	}
}

func printCommands(commands []command, p *i18n.PrinterWrapper) {
	for _, c := range commands {
		if usages, ok := c.getUsage().([]string); ok && len(usages) > 0 {
			printf("%2s%-14s%-30s", "", c.getKey(), usages[0])
			for _, usage := range usages[1:] {
				printf("%16s%-30s", "", usage)
			}
		} else if usage, ok := c.getUsage().(string); ok {
			printf("%2s%-14s%-30s", "", c.getKey(), usage)
		} else {
			printf("%2s%-14s%-30s", "", c.getKey(), "")
		}
		printf("%-16s%-30s", "", c.getDescription(p))
		printf("")
	}
}

func usage() {
	p := i18n.GetCurrentPrinter()
	printf(p.Sprintf("Usage:") + " obsutil [command] [args...] [options...]")
	p.Printf("You can use \"obsutil help command\" to view the specific help of each command")

	printf("")
	p.Printf("Basic commands:")
	printCommands(basic, p)
	printf("")
	p.Printf("Other commands:")
	printCommands(other, p)
}

func runCommand(args []string, additionalAction func(c command)) error {
	cmd := args[0]
	if c, ok := commands[cmd]; ok {
		if additionalAction != nil {
			additionalAction(c)
		}
		return c.parse(args[1:])
	}
	printf("Error: No such command: \"%s\", please try \"help\" for more information!", cmd)
	return assist.ErrInvalidArgs
}

func logUserInput(inputs []string, c command) {
	userInput := strings.Join(inputs, " ")
	userInput += " "
	userInput = cleanUpAkRegex1.ReplaceAllString(userInput, "-i=xxx ")
	userInput = cleanUpAkRegex2.ReplaceAllString(userInput, "-i xxx ")
	userInput = cleanUpSkRegex1.ReplaceAllString(userInput, "-k=xxx ")
	userInput = cleanUpSkRegex2.ReplaceAllString(userInput, "-k xxx ")
	userInput = cleanUpTokenRegex1.ReplaceAllString(userInput, "-t=xxx ")
	userInput = cleanUpTokenRegex2.ReplaceAllString(userInput, "-t xxx ")
	userInput = cleanUpTokenRegex3.ReplaceAllString(userInput, "-token=xxx ")
	userInput = cleanUpTokenRegex4.ReplaceAllString(userInput, "-token xxx ")
	doLog(LEVEL_INFO, "User input command \"%s\"", userInput)
}

func setCurrentLanguage() {
	i18n.SetCurrentLanguage(strings.ToLower(config["helpLanguage"]))
}

func Run() {
	if err := InitConfigFile(GetDefaultConfig(), true); err == nil {
		rand.Seed(time.Now().Unix())

		var exitErr error
		i18n.SetI18nStrings()
		setCurrentLanguage()

		var exitFlag int32

		//handle unexpected error
		defer func() {
			r := recover()
			if r != nil {
				printf("Unexpect error, please collect the logs and contact our support, %v", r)
				printf("%s", debug.Stack())
				doLog(LEVEL_ERROR, "%v", r)
				doLog(LEVEL_ERROR, "%s", debug.Stack())
			}

			if atomic.CompareAndSwapInt32(&exitFlag, 0, 1) {
				doClean()
			}
			assist.CheckErrorAndExit(exitErr)
		}()

		progress.InitCustomizeElements(config["colorfulProgress"] == c_true)
		initCommands()

		if _terminalWidth, err := assist.GetTerminalWidth(); err == nil && _terminalWidth > 80 {
			terminalWidth = _terminalWidth
		} else {
			terminalWidth = 80
		}

		args := os.Args[1:]
		if len(args) <= 0 && !assist.IsWindows() {
			usage()
			exitErr = assist.ErrInvalidArgs
			return
		}

		if len(args) <= 0 {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt, os.Kill)
			var currentCommand command
			go func() {
				<-ch
				if currentCommand != nil {
					if _c, ok := currentCommand.(taskRecorder); ok && _c.getTaskId() != "" {
						printf("\nTask id is: %s", _c.getTaskId())
					}
				}

				if atomic.CompareAndSwapInt32(&exitFlag, 0, 1) {
					doClean()
				}
				assist.CheckErrorAndExit(assist.ErrInterrupted)
			}()

			additionalTips := func() {
				printf("Enter \"help\" or \"help command\" to show help docs")
			}

			callback := func(input string) {
				_input, placeHolder, flag, err := assist.PreprocessInput(input)
				if err != nil {
					printf("Error: Invalid input, please try again, %s", err.Error())
					return
				}

				inputs := splitRegex.Split(_input, -1)

				if flag {
					for index, item := range inputs {
						if strings.HasPrefix(item, "\"") && strings.HasSuffix(item, "\"") {
							item = strings.Replace(item, placeHolder, " ", -1)
							item = item[1 : len(item)-1]
							inputs[index] = item
						}
					}
				}

				if len(inputs) == 1 && inputs[0] == "obsutil" {
					printf("Error: Invalid input, please try again")
					return
				}

				originInputs := inputs
				if inputs[0] == "obsutil" {
					inputs = inputs[1:]
				}

				progress.ResetContext()
				runningRound++

				runCommand(inputs, func(c command) {
					logUserInput(originInputs, c)
					currentCommand = c
				})
			}
			assist.EnterBashMode(additionalTips, callback)
			return
		}

		exitErr = runCommand(args, func(c command) {
			logUserInput(os.Args, c)

			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt, os.Kill)
			go func() {
				<-ch
				if _c, ok := c.(taskRecorder); ok {
					_c.printTaskId()
				}

				if atomic.CompareAndSwapInt32(&exitFlag, 0, 1) {
					doClean()
				}

				assist.CheckErrorAndExit(assist.ErrInterrupted)
			}()
		})
	} else {
		printError(err)
		assist.CheckErrorAndExit(assist.ErrInitializing)
	}

}
