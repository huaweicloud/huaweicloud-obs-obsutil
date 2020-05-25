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
	"concurrent"
	"fmt"
	"obs"
	"progress"
	"time"
)

type restoreCommand struct {
	recursiveCommand
	days      int
	tier      string
	version   bool
	versionId string
}

func (c *restoreCommand) restoreObject(bucket, key, versionId string, restoreTier obs.RestoreTierType, batchFlag int) bool {
	abortHandler := func() {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		c.failedLogger.doRecord("Bucket [%s], key [%s], %s, n/a, n/a, error code [%s], error message [%s], n/a", bucket, key, versionIdStr,
			"AbortError", "Task is aborted")
	}
	actionFunc := func() (output *obs.BaseModel, err error) {
		input := &obs.RestoreObjectInput{}
		input.Bucket = bucket
		input.Key = key
		input.Days = c.days
		input.Tier = restoreTier
		input.VersionId = versionId
		input.RequestPayer = c.payer
		return obsClient.RestoreObject(input)
	}
	recordHandler := func(cost int64, output *obs.BaseModel, err error) {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		if err == nil {
			c.succeedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], request id [%s]", bucket, key, versionIdStr, cost, output.StatusCode, output.RequestId)
		} else {
			status, code, message, requestId := c.checkAbort(err, 401, 405)
			c.failedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", bucket, key,
				versionIdStr, cost, status, code, message, requestId)
		}
	}
	printHandler := func(cost int64, output *obs.BaseModel, err error) {
		if versionId != "" {
			if err == nil {
				printf("Start to restore object [%s] with version id [%s] in the bucket [%s] successfully, cost [%d] ms, request id [%s]", key, versionId, bucket, cost, output.RequestId)
				doLog(LEVEL_DEBUG, "Start to restore object [%s] with version id [%s] in the bucket [%s] successfully, cost [%d] ms, request id [%s]", key, versionId, bucket, cost, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Start to restore object [%s] with version id [%s] in the bucket [%s] failed, cost [%d] ms", key, versionId, bucket, cost))
			}
		} else {
			if err == nil {
				printf("Start to restore object [%s] in the bucket [%s] successfully, cost [%d] ms, request id [%s] ", key, bucket, cost, output.RequestId)
				doLog(LEVEL_DEBUG, "Start to restore object [%s] in the bucket [%s] successfully, cost [%d] ms, request id [%s]", key, bucket, cost, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Start to restore object [%s] in the bucket [%s] failed, cost [%d] ms", key, bucket, cost))
			}
		}
	}

	return c.simpleAction(batchFlag, abortHandler, actionFunc, recordHandler, printHandler)
}

func (c *restoreCommand) submitRestoreTask(bucket, prefix string, restoreTier obs.RestoreTierType,
	pool concurrent.Pool, ch progress.SingleBarChan) (totalCnt int64, hasListError error) {

	if c.version {
		actionFunc := func(bucket, key, versionId string) bool {
			return c.restoreObject(bucket, key, versionId, restoreTier, 2)
		}
		isSkipFunc := func(version obs.Version) bool {
			return version.StorageClass != obs.StorageClassCold
		}
		return c.submitListVersionsTask(bucket, prefix, "restore", pool, ch, actionFunc, isSkipFunc, nil, false)
	}

	actionFunc := func(bucket, key string) bool {
		return c.restoreObject(bucket, key, "", restoreTier, 2)
	}
	isSkipFunc := func(content obs.Content) bool {
		return content.StorageClass != obs.StorageClassCold
	}
	return c.submitListObjectsTask(bucket, prefix, "restore", pool, ch, actionFunc, isSkipFunc)
}

func (c *restoreCommand) recordStartFunc() time.Time {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "error code", "error message", "request id")
	return start
}

func (c *restoreCommand) restoreObjects(bucket, prefix string, restoreTier obs.RestoreTierType) error {
	submitFunc := func(pool concurrent.Pool, ch progress.SingleBarChan) (int64, error) {
		return c.submitRestoreTask(bucket, prefix, restoreTier, pool, ch)
	}

	errorHandleFunc := func(hasListError error) {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to restore failed", bucket))
	}

	return c.recursiveAction(bucket, prefix, submitFunc, errorHandleFunc, c.recordStartFunc, true)
}

func initRestore() command {
	c := &restoreCommand{}
	c.key = "restore"
	c.usage = c_cloud_url_usage
	c.description = "restore objects in a bucket to be readable"
	c.define = func() {
		c.init()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
		c.flagSet.BoolVar(&c.version, "v", false, "")
		c.flagSet.IntVar(&c.jobs, "j", 0, "")
		c.flagSet.IntVar(&c.days, "d", 1, "")
		c.flagSet.StringVar(&c.tier, "t", "expedited", "")
		c.flagSet.StringVar(&c.versionId, "versionId", "", "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		var restoreTier obs.RestoreTierType
		checkParamFunc := func(prefix string) bool {
			if c.days < 1 || c.days > 30 {
				printf("Error: Invalid d [%d], the range is [1, 30]", c.days)
				return false
			}

			if c.tier != "" {
				if _restoreTier, ok := restoreTierType[c.tier]; ok {
					restoreTier = _restoreTier
				} else {
					printf("Error: Invalid t [%s], possible values are [standard|expedited]", c.tier)
					return false
				}
			}
			_, ok := getRequestPayerType(c.payer)
			if !ok {
				return false
			}
			return true
		}

		prefixFunc := func(bucket, prefix string, batchFlag int) error {
			if c.restoreObject(bucket, prefix, c.versionId, restoreTier, batchFlag) {
				return nil
			}

			return assist.ErrExecuting
		}
		recursivePrefixFun := func(bucket, prefix string) error {
			return c.restoreObjects(bucket, prefix, restoreTier)
		}

		return c.chooseAction(checkParamFunc, nil, nil, prefixFunc, recursivePrefixFun, c.recordStartFunc)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s\n", "", p.Sprintf("restore objects in a bucket to be readable"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil restore obs://bucket/key [-d=1] [-t=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]"+restoreCommandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil restore obs://bucket/[prefix] -r [-f] [-v] [-d=1] [-t=xxx] [-o=xxx] [-j=1] [-config=xxx]"+restoreCommandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch restore objects by prefix"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while an object is to be restored"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files when restoring one object"))
		printf("")
		printf("%2s%s", "", "-v")
		printf("%4s%s", "", p.Sprintf("restore versions of objects by prefix"))
		printf("")
		printf("%2s%s", "", "-d=1")
		printf("%4s%s", "", p.Sprintf("retention period of each restored object, in days. the range is [1, 30] and the default value is 1"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent restore jobs, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-versionId=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID of the object to be restored"))
		printf("")
		printf("%2s%s", "", "-t=xxx")
		printf("%4s%s", "", p.Sprintf("option for restoring objects, possible values are [standard|expedited]"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the restore results"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		if assist.IsHec() {
			printf("%2s%s", "", "-e=xxx")
			printf("%4s%s", "", p.Sprintf("endpoint"))
			printf("")
			printf("%2s%s", "", "-i=xxx")
			printf("%4s%s", "", p.Sprintf("access key ID"))
			printf("")
			printf("%2s%s", "", "-k=xxx")
			printf("%4s%s", "", p.Sprintf("security key ID"))
			printf("")
			printf("%2s%s", "", "-token=xxx")
			printf("%4s%s", "", p.Sprintf("security token"))
			printf("")
		}
		commandRequestPayerHelp(p)
	}

	return c
}
