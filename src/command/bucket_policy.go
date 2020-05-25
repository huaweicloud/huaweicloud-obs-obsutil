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
	"bytes"
	"command/i18n"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"obs"
	"os"
	"strings"
)

const (
	c_put    = "put"
	c_get    = "get"
	c_delete = "delete"
)

type bucketPolicyCommand struct {
	cloudUrlCommand
	method    string
	localfile string
}

var bucketPolicyMethod = map[string]string{
	c_put:    "put",
	c_get:    "get",
	c_delete: "delete",
}

func getBucketPolicyMethod(method string) (string, bool) {
	if method != "" {
		policyMethod, ok := bucketPolicyMethod[method]
		if !ok {
			printf("Error: Invalid method [%s], possible values are:[%s][%s][%s]", method, c_put, c_get, c_delete)
			return "", false
		}
		return policyMethod, true
	}
	return "", false
}

func (c *bucketPolicyCommand) confirm(str string) bool {
	var val string

	printf("output file already exist, are you sure overwrite \"%s\"(y or N)? ", str)
	if _, err := fmt.Scanln(&val); err != nil || (strings.ToLower(val) != "yes" && strings.ToLower(val) != "y") {
		return false
	}
	return true
}

func (c *bucketPolicyCommand) getBucketPolicy(bucket string) error {
	output, err := obsClient.GetBucketPolicy(bucket, obs.WithReqPaymentHeader(c.payer))
	if err != nil {
		return err
	}

	var outFile *os.File
	fileName := c.localfile
	outFile = os.Stdout
	if fileName != "" {
		_, err = os.Stat(fileName)
		if err == nil {
			isContinue := c.confirm(fileName)
			if !isContinue {
				return nil
			}
		}
		outFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
		if err != nil {
			doLog(LEVEL_ERROR, "open file [%s] failed, %s", fileName, err.Error())
			return err
		}
		defer func() {
			if closeErr := outFile.Close(); closeErr != nil {
				doLog(LEVEL_ERROR, "open file [%s] failed, %s", fileName, err.Error())
			}
		}()
	}
	var jsonText bytes.Buffer
	err = json.Indent(&jsonText, []byte(output.Policy), "", "    ")
	if err != nil {
		doLog(LEVEL_ERROR, "open file [%s] failed, %s", fileName, err.Error())
	}
	outFile.Write(jsonText.Bytes())
	if c.localfile != "" {
		printf("Export bucketPolicy to [%s] succeed, requestId is [%s]", fileName, output.RequestId)
		return nil
	}
	printf("\n Get bucketPolicy succeed, requestId is [%s]", output.RequestId)
	return nil
}

func (c *bucketPolicyCommand) deleteBucketPolicy(bucket string) error {
	output, err := obsClient.DeleteBucketPolicy(bucket, obs.WithReqPaymentHeader(c.payer))
	if err != nil {
		return err
	}
	printf("Delete bucketPolicy succeed, requestId is [%s]", output.RequestId)
	return nil
}

func (c *bucketPolicyCommand) putBucketPolicy(bucket string) error {
	localFile := c.localfile
	if localFile == "" {
		return fmt.Errorf("localFile should be set")
	}
	inputFile, err := os.Stat(localFile)
	if err != nil {
		doLogError(err, LEVEL_ERROR, "stat file failed")
		return fmt.Errorf("stat file [%s] failed", localFile)
	}

	if inputFile.IsDir() {
		return fmt.Errorf("%s is a directory, localFile must be a exist file", localFile)
	}

	if inputFile.Size() == 0 {
		return fmt.Errorf("%s is a empty file", localFile)
	}

	fileFd, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := fileFd.Close(); closeErr != nil {
			doLog(LEVEL_ERROR, "open file [%s] failed, %s", localFile, err.Error())
		}
	}()

	policy, err := ioutil.ReadAll(fileFd)
	if err != nil {
		return err
	}

	input := &obs.SetBucketPolicyInput{}
	input.Bucket = bucket
	input.Policy = string(policy)
	input.RequestPayer = c.payer
	output, err := obsClient.SetBucketPolicy(input)
	if err != nil {
		return err
	}

	printf("Put bucketPolicy succeed, requestId is [%s]", output.RequestId)
	return nil
}

func initBucketPolicyCommand() command {
	c := &bucketPolicyCommand{}
	c.key = "bucketpolicy"
	c.usage = c_cloud_url_usage
	c.description = ""

	c.define = func() {
		c.flagSet.StringVar(&c.method, "method", "", "")
		c.flagSet.StringVar(&c.localfile, "localfile", "", "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
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

		method, succeed := getBucketPolicyMethod(c.method)
		if !succeed {
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}
		_, ok := getRequestPayerType(c.payer)
		if !ok {
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}
		var ret error
		switch method {
		case c_get:
			ret = c.getBucketPolicy(bucket)
		case c_put:
			ret = c.putBucketPolicy(bucket)
		case c_delete:
			ret = c.deleteBucketPolicy(bucket)
		default:
			ret = assist.ErrInvalidArgs
		}
		if ret != nil {
			printError(ret)
		}
		return ret
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("get, put or delete bucket policy"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil bucketpolicy obs://bucket -method=xxx [-localfile=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-method=xxx")
		printf("%4s%s", "", p.Sprintf("the operation you want to do,possible values are [get, put, delete]"))
		printf("")
		printf("%2s%s", "", "-localFile=xxx")
		printf("%4s%s", "", p.Sprintf("the policy json file which you want to get or put, only support when method is get or put"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}
	return c
}
