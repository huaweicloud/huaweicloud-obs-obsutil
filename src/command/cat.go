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
	"fmt"
	"io/ioutil"
	"obs"
)

type catCommand struct {
	cloudUrlCommand
}

func initCat() command {
	c := &catCommand{}
	c.key = "cat"
	c.usage = "cloud_url"
	c.description = "view the content of a text object in a bucket"
	c.define = func() {
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()

		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		bucket, key, err := c.splitCloudUrl(cloudUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		if key == "" {
			printf("Error: The key is not specified")
			return assist.ErrInvalidArgs
		}

		_, ok := getRequestPayerType(c.payer)
		if !ok {
			return assist.ErrInvalidArgs
		}

		input := &obs.GetObjectInput{}
		input.Bucket = bucket
		input.Key = key
		input.RequestPayer = c.payer
		output, err := obsClient.GetObject(input)
		if err != nil {
			printError(err)
			return assist.ErrExecuting
		}

		if output.ContentLength > 10*mb {
			printf("Error: The file size [%d] is too large, the supported maximum size is 10MB", output.ContentLength)
			return assist.ErrExecuting
		}

		defer output.Body.Close()

		ret, err := ioutil.ReadAll(output.Body)
		if err != nil {
			printf("Error: Read content error, err:%s", err.Error())
			return assist.ErrExecuting
		}

		fmt.Println(assist.BytesToString(ret))

		return nil
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("view the content of a text object in a bucket"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil cat obs://bucket/key [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
