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
	"net/url"
	"obs"
	"strings"
	"sync/atomic"
)

type shareLsCommand struct {
	shareCommand
	short       bool
	dir         bool
	marker      string
	bytesFormat string
	prefix      string
	limit       int64
}

func (c *shareLsCommand) constructListObjectsUrl(parsedUrl *url.URL) string {
	signedUrl := constructCommonUrl(parsedUrl, "")
	prefix := c.prefix
	if c.dir {
		signedUrl = append(signedUrl, fmt.Sprintf("%s=%s&", "delimiter", "/"))
		if prefix != "" && !isObsFolder(prefix) {
			prefix += "/"
		}
	}

	if prefix != "" {
		signedUrl = append(signedUrl, fmt.Sprintf("%s=%s&", "prefix", url.QueryEscape(prefix)))
	}

	commonUrl := strings.Join(signedUrl, "")
	doLog(LEVEL_INFO, "The common url for listing objects is [%s]", commonUrl)
	return commonUrl
}

func (c *shareLsCommand) listObjects(parsedUrl *url.URL) error {
	commonUrl := c.constructListObjectsUrl(parsedUrl)

	limit := c.limit
	count := limit
	if count <= 0 {
		count = 1000
	}

	truncated := false
	totalFolders := make([]string, 0, count)
	totalObjects := make([]obs.Content, 0, count)
	nextMarker := ""
	bucket := ""
	var totalCount int64

	h := &assist.HintV2{}
	h.MessageFunc = func() string {
		count := ""
		if tc := atomic.LoadInt64(&totalCount); tc > 0 {
			count = "[" + assist.Int64ToString(tc) + "]"
		}
		return fmt.Sprintf("Listing objects %s", count)
	}
	h.Start()

	requestHeaders := map[string][]string{"Host": []string{parsedUrl.Host}}
	marker := c.marker
	var hasListError error
	var signedUrl []string
	for {
		maxKeys := 1000
		if limit > 0 {
			if limit <= 1000 {
				maxKeys = int(limit)
				truncated = true
			} else {
				limit -= 1000
			}
		}

		if signedUrl == nil {
			signedUrl = make([]string, 0, 3)
		} else {
			signedUrl = signedUrl[:0]
		}

		signedUrl = append(signedUrl, commonUrl)

		if marker != "" {
			signedUrl = append(signedUrl, fmt.Sprintf("%s=%s&", "marker", url.QueryEscape(marker)))
		}

		signedUrl = append(signedUrl, fmt.Sprintf("%s=%d", "max-keys", maxKeys))
		requestUrl := strings.Join(signedUrl, "")
		output, err := obsClient.ListObjectsWithSignedUrl(requestUrl, requestHeaders)
		if err != nil {
			hasListError = err
			break
		}

		if bucket == "" {
			bucket = output.Name
		}

		atomic.AddInt64(&totalCount, int64(len(output.Contents)))
		atomic.AddInt64(&totalCount, int64(len(output.CommonPrefixes)))

		folders, objects := getObjectsResult(output)
		totalFolders = append(totalFolders, folders...)
		totalObjects = append(totalObjects, objects...)
		if !output.IsTruncated {
			break
		}

		if truncated {
			nextMarker = output.NextMarker
			break
		}

		marker = output.NextMarker
	}

	h.End()

	if hasListError != nil {
		printError(hasListError)
		return assist.ErrUncompeleted
	}

	printListObjectsResult(totalFolders, totalObjects, c.short, c.dir, bucket, c.prefix, nextMarker, c.bytesFormat)
	return nil
}

func initShareLs() command {
	c := &shareLsCommand{}
	c.key = "share-ls"
	c.usage = c_share_usage
	c.description = "list objects using authorization code and access code"
	c.skipCheckAkSk = true

	c.define = func() {
		c.init()
		c.flagSet.BoolVar(&c.short, "s", false, "")
		c.flagSet.BoolVar(&c.dir, "d", false, "")
		c.flagSet.StringVar(&c.marker, "marker", "", "")
		c.flagSet.StringVar(&c.bytesFormat, "bf", "", "")
		c.flagSet.StringVar(&c.prefix, "prefix", "", "")
		c.flagSet.Int64Var(&c.limit, "limit", 1000, "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) < 1 {
			c.showHelp()
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		parsedUrl, allowedPrefix, err := c.prepareAccessUrl(args[0], args[1:])
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		defer func() {
			c.printAuthorizedPrefix(allowedPrefix)
		}()

		if c.prefix == "" {
			c.prefix = allowedPrefix
		} else if allowedPrefix != "" && !strings.HasPrefix(c.prefix, allowedPrefix) {
			printf("Error: Invalid prefix [%s], must start with [%s]", c.prefix, allowedPrefix)
			return assist.ErrInvalidArgs
		}

		if c.bytesFormat != "" && c.bytesFormat != "human-readable" && c.bytesFormat != c_raw {
			printf("Error: Invalid bf [%s], possible values are:[human-readable|raw]", c.bytesFormat)
			return assist.ErrInvalidArgs
		}

		return c.listObjects(parsedUrl)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("list objects using authorization code and access code"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil share-ls authorization_code [-ac=xxx] [-prefix=xxx] [-s] [-d] [-marker=xxx] [-bf=xxx] [-limit=1] [-config=xxx]"+commandCommonSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil share-ls file://authorization_code_file_url [-ac=xxx] [-prefix=xxx] [-s] [-d] [-marker=xxx] [-bf=xxx] [-limit=1] [-config=xxx]"+commandCommonSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-ac=xxx")
		printf("%4s%s", "", p.Sprintf("the access code"))
		printf("")
		printf("%2s%s", "", "-prefix=xxx")
		printf("%4s%s", "", p.Sprintf("the prefx to list objects"))
		printf("")
		printf("%2s%s", "", "-s")
		printf("%4s%s", "", p.Sprintf("show results in brief mode"))
		printf("")
		printf("%2s%s", "", "-d")
		printf("%4s%s", "", p.Sprintf("list objects and sub-folders in the current folder"))
		printf("")
		printf("%2s%s", "", "-marker=xxx")
		printf("%4s%s", "", p.Sprintf("the marker to list objects"))
		printf("")
		printf("%2s%s", "", "-bf=xxx")
		printf("%4s%s", "", p.Sprintf("the bytes format in results when listing objects, possible values are [human-readable|raw]"))
		printf("")
		printf("%2s%s", "", "-limit=1000")
		printf("%4s%s", "", p.Sprintf("show results by limited number"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
	}

	return c
}
