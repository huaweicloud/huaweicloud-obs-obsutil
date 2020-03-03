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
)

type mbCommand struct {
	cloudUrlCommand
	fs       bool
	acl      string
	sc       string
	location string
	az       string
}

func initMb() command {
	c := &mbCommand{}
	c.key = "mb"
	c.usage = c_cloud_url_usage
	c.description = "create a bucket with the specified parameters"

	c.define = func() {
		if assist.IsHec() {
			c.flagSet.BoolVar(&c.fs, "fs", false, "")
			c.flagSet.StringVar(&c.az, "az", "", "")
		}
		c.flagSet.StringVar(&c.acl, "acl", "", "")
		c.flagSet.StringVar(&c.sc, "sc", "", "")
		c.flagSet.StringVar(&c.location, "location", "", "")
	}

	c.additionalValidate = func(cloudUrl string) bool {
		return cloudUrlRegex.MatchString(cloudUrl)
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		c.printStart()

		bucket := cloudUrl[6:]
		input := &obs.CreateBucketInput{}
		input.Bucket = bucket
		if c.location != "" {
			input.Location = c.location
		}

		aclType, succeed := getBucketAclType(c.acl)
		if !succeed {
			return assist.ErrInvalidArgs
		}
		input.ACL = aclType

		storageClassType, succeed := getStorageClassType(c.sc)
		if !succeed {
			return assist.ErrInvalidArgs
		}
		input.StorageClass = storageClassType

		availableZone, succeed := getAvailableZoneType(c.az)
		if !succeed {
			return assist.ErrInvalidArgs
		}
		input.AvailableZone = availableZone

		var output *obs.BaseModel
		if c.fs {
			newBucketInput := &obs.NewBucketInput{}
			newBucketInput.CreateBucketInput = *input
			output, err = obsClient.NewBucket(newBucketInput)
		} else {
			output, err = obsClient.CreateBucket(input)
		}
		if err == nil {
			printf("Create bucket [%s] successfully, request id [%s]", bucket, output.RequestId)
			doLog(LEVEL_INFO, "Create bucket [%s] successfully, request id [%s]", bucket, output.RequestId)
			printf("Notice: If the configured endpoint is a global domain name, " +
				"you may need to wait serveral minutes before performing uploading operations on the created bucket. " +
				"Therefore, configure the endpoint to a regional domain name if you want instant uploading operations on the bucket.")
			return nil
		}
		logError(err, LEVEL_INFO, fmt.Sprintf("Create bucket [%s] failed", bucket))
		return assist.ErrExecuting
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("create a bucket with the specified parameters"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil mb obs://bucket [-fs] [-az=xxx] [-acl=xxx] [-sc=xxx] [-location=xxx] [-config=xxx]"+commandCommonSyntax())
		printf("")

		p.Printf("Options:")
		if assist.IsHec() {
			printf("%2s%s", "", "-fs")
			printf("%4s%s", "", p.Sprintf("create a bucket that supports POSIX"))
			printf("")
			printf("%2s%s", "", "-az=xxx")
			printf("%4s%s", "", p.Sprintf("the AZ of the bucket, possible values are [multi-az]"))
			printf("")
		}
		printf("%2s%s", "", "-acl=xxx")
		printf("%4s%s", "", p.Sprintf("the ACL of the bucket, possible values are [private|public-read|public-read-write]"))
		printf("")
		printf("%2s%s", "", "-sc=xxx")
		printf("%4s%s", "", p.Sprintf("the default storage class of the bucket, possible values are [standard|warm|cold]"))
		printf("")
		printf("%2s%s", "", "-location=xxx")
		printf("%4s%s", "", p.Sprintf("the region where the bucket is located"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
	}

	return c
}
