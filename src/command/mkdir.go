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
	"obs"
	"os"
	"strings"
)

type mkdirCommand struct {
	cloudUrlCommand
}

func initMkdir() command {

	c := &mkdirCommand{}
	c.key = "mkdir"
	c.usage = "cloud_url|folder_url"
	c.description = "create folder(s) in a specified bucket or in the local file system"
	c.define = func() {
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}
	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err != nil {
			if cloudUrl == "" {
				printError(err)
				return assist.ErrInvalidArgs
			}

			folderUrl := assist.NormalizeFilePath(cloudUrl)
			if err := assist.MkdirAll(folderUrl, os.ModePerm); err != nil {
				printError(err)
				return assist.ErrExecuting
			}
			printf("Create folder(s) [%s] in the local file system successfully", folderUrl)
			return nil
		}
		bucket, key, err := c.splitCloudUrl(cloudUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		_, succeed := getRequestPayerType(c.payer)
		if !succeed {
			return assist.ErrInvalidArgs
		}

		if key == "" {
			printf("Error: No folder(s) specified to create in the bucket [%s]", bucket)
			return assist.ErrInvalidArgs
		}

		fsStatus, err := c.checkBucketFSStatus(bucket)
		if err != nil {
			printError(err)
			return assist.ErrCheckBucketStatus
		}

		if fsStatus == c_enabled {
			printf("The bucket [%s] supports POSIX, create folder(s) directly", bucket)
			input := &obs.NewFolderInput{}
			input.Bucket = bucket
			input.Key = key
			input.RequestPayer = c.payer
			output, err := obsClient.NewFolder(input)
			if err == nil {
				printf("Create folder [obs://%s/%s] successfully, request id [%s]", bucket, key, output.RequestId)
				doLog(LEVEL_INFO, "Create folder [obs://%s/%s] successfully, request id [%s]", bucket, key, output.RequestId)
				return nil
			}
			logError(err, LEVEL_INFO, fmt.Sprintf("Create folder [obs://%s/%s] failed", bucket, key))
			return assist.ErrExecuting
		}

		if fsStatus == c_disabled {
			printf("The bucket [%s] does not support POSIX, create folder(s) step by step", bucket)
		} else if fsStatus == c_unknown {
			printf("Can not identify whether the bucket [%s] supports POSIX, create folder(s) step by step", bucket)
		}

		if isObsFolder(key) {
			key = key[:len(key)-1]
		}
		folders := strings.Split(key, "/")
		current := ""
		input := &obs.PutObjectInput{}
		input.Bucket = bucket
		input.ContentLength = 0
		input.RequestPayer = c.payer
		allSucceed := true
		for _, folder := range folders {
			current += folder + "/"
			input.Key = current
			output, err := obsClient.PutObject(input)
			if err == nil {
				printf("Create folder [obs://%s/%s] successfully, request id [%s]", bucket, current, output.RequestId)
				doLog(LEVEL_INFO, "Create folder [obs://%s/%s] successfully, request id [%s]", bucket, current, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Create folder [obs://%s/%s] failed", bucket, current))
				allSucceed = false
			}
		}

		if !allSucceed {
			return assist.ErrUncompeleted
		}
		return nil

	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("create folder(s) in a specified bucket or in the local file system"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil mkdir obs://bucket/folder1/folder2/folder3/ [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil mkdir folder_url [-config=xxx]")
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
