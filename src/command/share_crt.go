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
	"os"
	"path/filepath"
	"strings"
	"time"
)

var shareConsoleUrl = ""

func SetShareConsoleUrl(s string) {
	if s != "" {
		if !strings.HasPrefix(s, "https://") {
			shareConsoleUrl = "https://" + s
		} else {
			shareConsoleUrl = s
		}
	}
}

type shareCrtCommand struct {
	cloudUrlCommand
	accessCode     string
	validityPeriod string
	downloadUrl    string
}

func (c *shareCrtCommand) getBucketVersion(bucket string) string {
	input := &obs.GetBucketMetadataInput{}
	input.Bucket = bucket
	input.RequestPayer = c.payer
	output, err := obsClient.GetBucketMetadata(input)
	obsVersion := OBS_VERSION_UNKNOWN
	if err == nil {
		obsVersion = output.ObsVersion
	}
	return obsVersion
}

func (c *shareCrtCommand) parseValidityPeriod() (int64, error) {
	vp := strings.ToLower(c.validityPeriod)
	var unit int64 = 1
	if strings.HasSuffix(vp, "m") {
		unit = 30 * 24 * 60 * 60
		vp = vp[:len(vp)-1]
	} else if strings.HasSuffix(vp, "w") {
		unit = 7 * 24 * 60 * 60
		vp = vp[:len(vp)-1]
	} else if strings.HasSuffix(vp, "d") {
		unit = 24 * 60 * 60
		vp = vp[:len(vp)-1]
	} else if strings.HasSuffix(vp, "h") {
		unit = 60 * 60
		vp = vp[:len(vp)-1]
	} else if strings.HasSuffix(vp, "min") {
		unit = 60
		vp = vp[:len(vp)-3]
	} else if strings.HasSuffix(vp, "s") {
		vp = vp[:len(vp)-1]
	}

	ret, err := assist.StringToInt64V2(vp)
	if err != nil {
		return -1, fmt.Errorf("Invalid validity period [%s], %s", c.validityPeriod, err.Error())
	}

	if ret <= 0 {
		return -1, fmt.Errorf("Invalid validity period [%s], the value must greater than 0", c.validityPeriod)
	}

	return ret * unit, nil
}

func initShareCrt() command {
	c := &shareCrtCommand{}
	c.key = "create-share"
	c.usage = c_cloud_url_usage
	c.description = "create authorization code for sharing"

	c.define = func() {
		c.flagSet.StringVar(&c.accessCode, "ac", "", "")
		c.flagSet.StringVar(&c.validityPeriod, "vp", "1d", "")
		c.flagSet.StringVar(&c.downloadUrl, "dst", "", "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		bucket, prefix, err := c.splitCloudUrl(cloudUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		_, succeed := getRequestPayerType(c.payer)
		if !succeed {
			return assist.ErrInvalidArgs
		}

		if c.accessCode == "" {
			c.accessCode, err = getUserInput("Please input your access code:")
			if err != nil {
				printError(err)
				return assist.ErrInvalidArgs
			}
		}

		if l := len(c.accessCode); l != 6 {
			printf("Error: Invalid access code, the length [%d] does not equal to 6", l)
			return assist.ErrInvalidArgs
		}

		validityPeriodSeconds, err := c.parseValidityPeriod()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		obsVersion := c.getBucketVersion(bucket)

		if obsVersion < "3.0" {
			printf("Error: Bucket [%s] cannot support to create authorization code for sharing", bucket)
			return assist.ErrUnsupported
		}

		if obsVersion == OBS_VERSION_UNKNOWN {
			printf("Error: Bucket [%s] cannot support to create authorization code for sharing, due to check obs version failed", bucket)
			return assist.ErrUnsupported
		}

		expiration := time.Now().UTC().Add(time.Second * time.Duration(validityPeriodSeconds))
		originPolicySlice := make([]string, 0, 5)
		originPolicySlice = append(originPolicySlice, fmt.Sprintf("{\"expiration\":\"%s\",", expiration.Format(ISO8601_DATE_FORMAT)))
		originPolicySlice = append(originPolicySlice, "\"conditions\":[")
		originPolicySlice = append(originPolicySlice, fmt.Sprintf("{\"%s\":\"%s\"}", "bucket", bucket))
		if prefix != "" {
			originPolicySlice = append(originPolicySlice, fmt.Sprintf(",[\"starts-with\", \"$key\", \"%s\"]", prefix))
		}
		originPolicySlice = append(originPolicySlice, "]}")
		originPolicy := strings.Join(originPolicySlice, "")

		input := &obs.CreateSignedUrlInput{}
		input.Bucket = bucket
		input.Key = prefix
		input.Policy = originPolicy
		input.Method = obs.HTTP_GET
		if prefix != "" {
			input.QueryParams = map[string]string{"prefix": prefix}
		}

		output, err := obsClient.CreateSignedUrl(input)
		if err != nil {
			printError(err)
			return assist.ErrExecuting
		}

		doLog(LEVEL_DEBUG, "The access url for sharing is [%s]", output.SignedUrl)

		token, err := AesEncrypt(assist.Base64Encode(assist.StringToBytes(output.SignedUrl)), assist.StringToBytes(aesShareKeyPrefix+c.accessCode), assist.StringToBytes(aesShareIv))
		if err != nil {
			printf("Error: Try to encrypt access url failed, %s", err.Error())
			return assist.ErrExecuting
		}

		result := make([]string, 0, 6)
		result = append(result, "Authorization Code:")
		result = append(result, fmt.Sprintf("%s?token=%s", shareConsoleUrl, token))
		result = append(result, "\nAccess Code:")
		result = append(result, c.accessCode)
		result = append(result, "\nValid Until:")
		result = append(result, fmt.Sprintf("%s", expiration.Local().Format(RFC1123_FORMAT)))

		if c.downloadUrl == "" {
			printf(strings.Join(result, "\n"))
			return nil
		}

		downloadUrl := assist.NormalizeFilePath(c.downloadUrl)
		stat, err := os.Lstat(downloadUrl)
		if err == nil {
			if stat.IsDir() {
				downloadUrl = assist.NormalizeFilePath(downloadUrl + "/result.txt")
			}
		} else if strings.HasSuffix(c.downloadUrl, "/") {
			if merr := assist.MkdirAll(downloadUrl, os.ModePerm); merr != nil {
				printf("Error: Try to create folder [%s] failed, %s", downloadUrl, merr.Error())
				return assist.ErrExecuting
			}
			downloadUrl = assist.NormalizeFilePath(downloadUrl + "/result.txt")
		} else {
			parent := filepath.Dir(downloadUrl)
			if merr := assist.MkdirAll(parent, os.ModePerm); merr != nil {
				printf("Error: Try to create folder [%s] failed, %s", parent, merr.Error())
				return assist.ErrExecuting
			}
		}

		if werr := ioutil.WriteFile(downloadUrl, assist.StringToBytes(strings.Join(result, "\n")), 0666); werr != nil {
			printf("Error: Try to write result to [%s] failed, %s", downloadUrl, werr.Error())
			return assist.ErrExecuting
		}

		printf("The result is generated to [%s]", downloadUrl)

		return nil
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("create authorization code for sharing"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil create-share obs://bucket[/prefix] [-ac=xxx] [-vp=xxx] [-dst=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-ac=xxx")
		printf("%4s%s", "", p.Sprintf("the access code"))
		printf("")
		printf("%2s%s", "", "-vp=xxx")
		printf("%4s%s", "", p.Sprintf("the validity period of authorization code, the default value is 1 day"))
		printf("")
		printf("%2s%s", "", "-dst=xxx")
		printf("%4s%s", "", p.Sprintf("the download url to which the result is generated"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
