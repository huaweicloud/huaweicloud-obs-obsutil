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
	"progress"
	"strings"
	"sync/atomic"
)

type signCommand struct {
	recursiveCommand
	expires int
}

func (c *signCommand) init() {
	c.recursiveCommand.init()
}

func (*signCommand) getSignedUrl(output *obs.CreateSignedUrlOutput) string {
	return strings.Replace(output.SignedUrl, "?AWSAccessKeyId=", "?AccessKeyId=", 1)
}

func (c *signCommand) generateDownloadUrl(bucket, key string, batchFlag bool) bool {
	input := &obs.CreateSignedUrlInput{}
	input.Bucket = bucket
	input.Key = key
	input.Method = obs.HTTP_GET
	input.Expires = c.expires
	input.RequestPayer = c.payer
	output, err := obsClient.CreateSignedUrl(input)

	if batchFlag {
		if err != nil {
			progress.AddFailedCount(1)
			c.failedLogger.doRecord("Generate the download url of [obs://%s/%s] failed", bucket, key)
			return false
		}
		progress.AddSucceedCount(1)
		c.succeedLogger.doRecord("obs://%s/%s: %s", bucket, key, c.getSignedUrl(output))
	} else {
		if err != nil {
			printf("Error: Generate the download url of [obs://%s/%s] failed", bucket, key)
			printf("%2sError: %s", "", err.Error())
			return false
		}

		printf("Download url of [obs://%s/%s] is:", bucket, key)
		printf("%2s%s", "", c.getSignedUrl(output))
	}

	return true
}

func (c *signCommand) generateDownloadUrls(bucket, prefix string) error {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefix
	input.MaxKeys = defaultListMaxKeys
	input.RequestPayer = c.payer
	action := "generate download url(s)"
	var hasListError error

	var totalCount int64
	h := &assist.HintV2{}
	h.MessageFunc = func() string {
		count := ""
		if tc := atomic.LoadInt64(&totalCount); tc > 0 {
			count = "[" + assist.Int64ToString(tc) + "]"
		}
		return fmt.Sprintf("Generate download urls for objects %s", count)
	}
	h.Start()

	for {
		start := assist.GetUtcNow()
		output, err := obsClient.ListObjects(input)
		if err != nil {
			hasListError = err
			break
		} else {
			cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
			doLog(LEVEL_DEBUG, "List objects in the bucket [%s] to %s successfully, cost [%d], request id [%s]", bucket, action, cost, output.RequestId)
		}
		for _, content := range output.Contents {
			if isObsFolder(content.Key) {
				continue
			}

			if c.matchExclude(content.Key) {
				continue
			}

			if !c.matchInclude(content.Key) {
				continue
			}

			if !c.matchLastModifiedTime(content.LastModified) {
				continue
			}

			c.generateDownloadUrl(bucket, content.Key, true)
			atomic.AddInt64(&totalCount, 1)
		}

		if !output.IsTruncated {
			doLog(LEVEL_INFO, "List objects to %s finished, bucket [%s], prefix [%s], marker [%s]", action, bucket, input.Prefix, input.Marker)
			break
		}
		input.Marker = output.NextMarker
	}

	h.End()
	var ret error
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to %s failed", bucket, action))
		ret = assist.ErrUncompeleted
	} else if progress.GetFailedCount() > 0 {
		ret = assist.ErrUncompeleted
	}
	printf("\nGenerate the download url(s) for the objects in the bucket [%s] finished", bucket)
	return ret
}

func initSign() command {

	c := &signCommand{}
	c.key = "sign"
	c.usage = c_cloud_url_usage
	c.description = "generate the download url(s) for the objects in a specified bucket"

	c.define = func() {
		c.init()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.IntVar(&c.expires, "e", 300, "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.include, "include", "", "")
		c.flagSet.StringVar(&c.exclude, "exclude", "", "")
		c.flagSet.StringVar(&c.timeRange, "timeRange", "", "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		if c.expires < 60 {
			printf("Error: The value of expiration time must greater than 60")
			return assist.ErrInvalidArgs
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

		if !c.recursive {
			if key == "" {
				printf("Error: No key specified to generate the download url")
				return assist.ErrInvalidArgs
			}
			if !c.generateDownloadUrl(bucket, key, false) {
				return assist.ErrExecuting
			}
			return nil
		}

		if succeed := c.checkInclude(); !succeed {
			return assist.ErrInvalidArgs
		}

		if succeed := c.checkExclude(); !succeed {
			return assist.ErrInvalidArgs
		}

		if succeed := c.checkTimeRange(); !succeed {
			return assist.ErrInvalidArgs
		}

		if err := c.ensureOutputDirectory(); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		if err := c.startLogger(true); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		defer c.endLogger()
		return c.generateDownloadUrls(bucket, key)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("generate the download url(s) for the objects in a specified bucket"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil sign obs://bucket/key [-e=300] [-config=xxx]"+signCommandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil sign obs://bucket[/prefix] -r [-e=300] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-o=xxx] [-config=xxx]"+signCommandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch generate the download url(s) of objects by prefix"))
		printf("")
		printf("%2s%s", "", "-e=300")
		printf("%4s%s", "", p.Sprintf("the expiration time of the generated download url(s), in seconds, the default value is 300"))
		printf("")
		printf("%2s%s", "", "-include=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be included when generating the download urls"))
		printf("")
		printf("%2s%s", "", "-exclude=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be excluded when generating the download urls"))
		printf("")
		printf("%2s%s", "", "-timeRange=time1-time2")
		printf("%4s%s", "", p.Sprintf("the time range, between which the download url(s) of objects will be generated"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the generated download urls"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		if assist.IsHec() {
			printf("%2s%s", "", "-endpoint=xxx")
			printf("%4s%s", "", p.Sprintf("endpoint"))
			printf("")
			printf("%2s%s", "", "-i=xxx")
			printf("%4s%s", "", p.Sprintf("access key ID"))
			printf("")
			printf("%2s%s", "", "-k=xxx")
			printf("%4s%s", "", p.Sprintf("security key ID"))
			printf("")
			printf("%2s%s", "", "-t=xxx")
			printf("%4s%s", "", p.Sprintf("security token"))
			printf("")
		}
		commandRequestPayerHelp(p)
	}

	return c
}
