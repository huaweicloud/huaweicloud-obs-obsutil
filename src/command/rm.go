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
	"sync"
	"sync/atomic"
	"time"
)

type rmCommand struct {
	recursiveCommand
	version   bool
	versionId string
}

func (c *rmCommand) deleteBucket(bucket string) bool {
	if !c.force && !confirm(fmt.Sprintf("Do you want delete bucket [%s] ? Please input (y/n) to confirm:", bucket)) {
		return false
	}
	output, err := obsClient.DeleteBucket(bucket, obs.WithReqPaymentHeader(c.payer))
	if err == nil {
		printf("Delete bucket [%s] successfully!", bucket)
		doLog(LEVEL_INFO, "Delete bucket [%s] successfully, request id [%s]", bucket, output.RequestId)
	} else {
		logError(err, LEVEL_INFO, fmt.Sprintf("Delete bucket [%s] failed", bucket))
	}
	return err == nil
}

func (c *rmCommand) deleteObject(bucket, key, versionId string, batchFlag int) bool {
	abortHandler := func() {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		c.failedLogger.doRecord("Bucket [%s], key [%s], %s, n/a, n/a, error code [%s], error message [%s], n/a", bucket, key, versionIdStr,
			"AbortError", "Task is aborted")
	}

	actionFunc := func() (output *obs.BaseModel, err error) {
		input := &obs.DropFileInput{}
		input.Bucket = bucket
		input.Key = key
		input.VersionId = versionId
		input.RequestPayer = c.payer
		if dropFileOutput, dropErr := obsClient.DropFile(input); dropErr != nil {
			err = dropErr
		} else {
			output = &dropFileOutput.BaseModel
		}
		return
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
				printf("Delete object [%s] with version id [%s] in the bucket [%s] successfully, cost [%d], request id [%s]", key, versionId, bucket, cost, output.RequestId)
				doLog(LEVEL_INFO, "Delete object [%s] with version id [%s] in the bucket [%s] successfully, cost [%d], request id [%s]", key, versionId, bucket, cost, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Delete object [%s] with version id [%s] in the bucket [%s] failed, cost [%d]", key, versionId, bucket, cost))
			}
		} else {
			if err == nil {
				printf("Delete object [%s] in the bucket [%s] successfully, cost [%d], request id [%s]", key, bucket, cost, output.RequestId)
				doLog(LEVEL_INFO, "Delete object [%s] in the bucket [%s] successfully, cost [%d], request id [%s]", key, bucket, cost, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Delete object [%s] in the bucket [%s] failed, cost [%d]", key, bucket, cost))
			}
		}
	}

	return c.simpleAction(batchFlag, abortHandler, actionFunc, recordHandler, printHandler)
}

func (c *rmCommand) deleteMultiObjects(bucket string, keys []obs.ObjectToDelete, ch progress.SingleBarChan) {
	if atomic.LoadInt32(&c.abort) == 1 {
		for _, key := range keys {
			versionIdStr := c_na
			if key.VersionId != "" {
				versionIdStr = fmt.Sprintf("version id [%s]", key.VersionId)
			}
			c.failedLogger.doRecord("Bucket [%s], Key [%s], %s, n/a, n/a, error code [%s], error message [%s], n/a", bucket, key.Key,
				versionIdStr, "AbortError", "Task is aborted")
		}
		return
	}

	deleteInput := &obs.DeleteObjectsInput{}
	deleteInput.Bucket = bucket
	deleteInput.Quiet = false

	deleteInput.Objects = keys
	deleteInput.RequestPayer = c.payer
	start := assist.GetUtcNow()

	deleteOutput, err := obsClient.DeleteObjects(deleteInput)

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

	progress.AddFinishedCount(int64(len(keys)))

	if err == nil {
		succeedCntOnce := len(deleteOutput.Deleteds)
		failedCntOnce := len(deleteOutput.Errors)

		progress.AddTransaction(int64(succeedCntOnce + failedCntOnce))

		progress.AddSucceedCount(int64(succeedCntOnce))
		progress.AddFailedCount(int64(failedCntOnce))
		ch.Send(succeedCntOnce)
		if succeedCntOnce > 0 {
			for _, deleted := range deleteOutput.Deleteds {
				versionIdStr := c_na
				if deleted.VersionId != "" {
					versionIdStr = fmt.Sprintf("version id [%s]", deleted.VersionId)
				}
				c.succeedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], request id [%s]", bucket, deleted.Key, versionIdStr, cost, deleteOutput.StatusCode, deleteOutput.RequestId)
			}
		}

		if failedCntOnce > 0 {
			for _, _error := range deleteOutput.Errors {
				versionIdStr := c_na
				if _error.VersionId != "" {
					versionIdStr = fmt.Sprintf("version id [%s]", _error.VersionId)
				}
				c.failedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", bucket, _error.Key,
					versionIdStr, cost, deleteOutput.StatusCode, _error.Code, _error.Message, deleteOutput.RequestId)
			}
		}
	} else {
		progress.AddTransaction(1)
		progress.AddFailedCount(int64(len(keys)))
		status, code, message, requestId := c.checkAbort(err, 401, 405, 403)
		for _, key := range keys {
			versionIdStr := c_na
			if key.VersionId != "" {
				versionIdStr = fmt.Sprintf("version id [%s]", key.VersionId)
			}
			c.failedLogger.doRecord("Bucket [%s], key [%s], %s, n/a, status [%d], error code [%s], error message [%s], request id [%s]", bucket, key.Key,
				versionIdStr, status, code, message, requestId)
		}

	}
}

func (c *rmCommand) submitDeleteTask(bucket, prefix string, pool concurrent.Pool, ch progress.SingleBarChan) (totalCnt int64, hasListError error) {
	defer func() {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
	}()
	if c.version {
		input := &obs.ListVersionsInput{}
		input.Bucket = bucket
		input.Prefix = prefix
		input.MaxKeys = defaultListMaxKeys
		input.RequestPayer = c.payer
		for {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			start := assist.GetUtcNow()
			output, err := obsClient.ListVersions(input)
			if err != nil {
				hasListError = err
				break
			} else {
				cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
				doLog(LEVEL_INFO, "List objects in the bucket [%s] to delete successfully, cost [%d], request id [%s]", bucket, cost, output.RequestId)
			}

			length := len(output.Versions)
			length += len(output.DeleteMarkers)

			if length > 0 {
				keys := make([]obs.ObjectToDelete, 0, defaultListMaxKeys)
				for _, version := range output.Versions {
					key := version.Key
					versionId := version.VersionId
					if !c.force && !confirm(fmt.Sprintf("Do you want delete object [%s] with version id [%s] ? Please input (y/n) to confirm:", key, versionId)) {
						continue
					}
					atomic.AddInt64(&totalCnt, 1)
					keys = append(keys, obs.ObjectToDelete{Key: key, VersionId: versionId})
				}

				for _, deleteMarker := range output.DeleteMarkers {
					key := deleteMarker.Key
					versionId := deleteMarker.VersionId
					if !c.force && !confirm(fmt.Sprintf("Do you want delete marker [%s] with version id [%s] ? Please input (y/n) to confirm:", key, versionId)) {
						continue
					}
					atomic.AddInt64(&totalCnt, 1)
					keys = append(keys, obs.ObjectToDelete{Key: key, VersionId: versionId})
				}

				if len(keys) > 0 {
					pool.ExecuteFunc(func() interface{} {
						c.deleteMultiObjects(bucket, keys, ch)
						return nil
					})

				}
			}

			if !output.IsTruncated {
				doLog(LEVEL_INFO, "List versioning objects to delete finished, bucket [%s], prefix [%s], marker [%s], versionIdMarker [%s]", bucket, input.Prefix, input.KeyMarker, input.VersionIdMarker)
				break
			}
			input.KeyMarker = output.NextKeyMarker
			input.VersionIdMarker = output.NextVersionIdMarker
		}
	} else {
		input := &obs.ListObjectsInput{}
		input.Bucket = bucket
		input.Prefix = prefix
		input.MaxKeys = defaultListMaxKeys
		input.RequestPayer = c.payer
		for {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			start := assist.GetUtcNow()
			output, err := obsClient.ListObjects(input)
			if err != nil {
				hasListError = err
				break
			} else {
				cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
				doLog(LEVEL_INFO, "List objects in the bucket [%s] to delete successfully, cost [%d], request id [%s]", bucket, cost, output.RequestId)
			}
			length := len(output.Contents)

			if length > 0 {
				keys := make([]obs.ObjectToDelete, 0, defaultListMaxKeys)
				for _, content := range output.Contents {
					key := content.Key
					if !c.force && !confirm(fmt.Sprintf("Do you want delete object [%s] ? Please input (y/n) to confirm:", key)) {
						continue
					}
					atomic.AddInt64(&totalCnt, 1)
					keys = append(keys, obs.ObjectToDelete{Key: key})
				}

				if len(keys) > 0 {
					pool.ExecuteFunc(func() interface{} {
						c.deleteMultiObjects(bucket, keys, ch)
						return nil
					})

				}
			}

			if !output.IsTruncated {
				doLog(LEVEL_INFO, "List objects to restore finished, bucket [%s], prefix [%s], marker [%s]", bucket, input.Prefix, input.Marker)
				break
			}
			input.Marker = output.NextMarker
		}
	}
	return
}

func (c *rmCommand) recordStartFunc() time.Time {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "error code", "error message", "request id")
	return start
}

func (c *rmCommand) deleteObjects(bucket, prefix string) error {

	submitFunc := func(pool concurrent.Pool, ch progress.SingleBarChan) (int64, error) {
		return c.submitDeleteTask(bucket, prefix, pool, ch)
	}

	errorHandleFunc := func(hasListError error) {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to delete failed", bucket))
	}

	return c.recursiveAction(bucket, prefix, submitFunc, errorHandleFunc, c.recordStartFunc, false)
}

func (c *rmCommand) doScan(bucket, folder string, pool concurrent.Pool, ch progress.SingleBarChan, totalCnt *int64, wg *sync.WaitGroup) {
	c.scanPool.ExecuteFunc(func() (r interface{}) {
		subWg := new(sync.WaitGroup)
		if c.version {
			input := &obs.ListVersionsInput{}
			input.Bucket = bucket
			input.Prefix = folder
			input.Delimiter = "/"
			input.MaxKeys = defaultListMaxKeys
			input.RequestPayer = c.payer
			for {
				if atomic.LoadInt32(&c.abort) == 1 {
					break
				}

				start := assist.GetUtcNow()
				output, err := obsClient.ListVersions(input)
				if err != nil {
					if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
						c.scanError.Store(err)
					}
					break
				} else {
					cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
					doLog(LEVEL_INFO, "List versioning objects in the bucket [%s] to delete successfully, cost [%d], request id [%s]", bucket, cost, output.RequestId)
				}

				for _, version := range output.Versions {
					if atomic.LoadInt32(&c.abort) == 1 {
						break
					}
					key := version.Key
					if isObsFolder(key) {
						continue
					}
					versionId := version.VersionId
					if !c.force && !confirm(fmt.Sprintf("Do you want delete object [%s] with version id [%s] ? Please input (y/n) to confirm:", key, versionId)) {
						continue
					}
					subWg.Add(1)
					atomic.AddInt64(totalCnt, 1)
					pool.ExecuteFunc(func() interface{} {
						ret := handleResult(c.deleteObject(bucket, key, versionId, 2), ch)
						subWg.Done()
						return ret
					})
				}

				for _, deleteMarker := range output.DeleteMarkers {
					if atomic.LoadInt32(&c.abort) == 1 {
						break
					}
					key := deleteMarker.Key
					if isObsFolder(key) {
						continue
					}
					versionId := deleteMarker.VersionId
					if !c.force && !confirm(fmt.Sprintf("Do you want delete marker [%s] with version id [%s] ? Please input (y/n) to confirm:", key, versionId)) {
						continue
					}
					subWg.Add(1)
					atomic.AddInt64(totalCnt, 1)
					pool.ExecuteFunc(func() interface{} {
						ret := handleResult(c.deleteObject(bucket, key, versionId, 2), ch)
						subWg.Done()
						return ret
					})
				}

				for _, subFolder := range output.CommonPrefixes {
					if atomic.LoadInt32(&c.abort) == 1 {
						break
					}
					subWg.Add(1)
					c.doScan(bucket, subFolder, pool, ch, totalCnt, subWg)
				}

				if !output.IsTruncated {
					doLog(LEVEL_INFO, "List versioning objects to delete finished, bucket [%s], folder [%s], marker [%s], versionIdMarker [%s]", bucket, folder, input.KeyMarker, input.VersionIdMarker)
					break
				}
				input.KeyMarker = output.NextKeyMarker
				input.VersionIdMarker = output.NextVersionIdMarker
			}

		} else {
			input := &obs.ListObjectsInput{}
			input.Bucket = bucket
			input.Prefix = folder
			input.Delimiter = "/"
			input.MaxKeys = defaultListMaxKeys
			input.RequestPayer = c.payer
			for {
				if atomic.LoadInt32(&c.abort) == 1 {
					break
				}
				start := assist.GetUtcNow()
				output, err := obsClient.ListObjects(input)
				if err != nil {
					if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
						c.scanError.Store(err)
					}
					break
				}
				cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
				doLog(LEVEL_INFO, "List objects in the bucket [%s] to delete successfully, cost [%d], request id [%s]", bucket, cost, output.RequestId)

				for _, content := range output.Contents {
					if atomic.LoadInt32(&c.abort) == 1 {
						break
					}
					key := content.Key
					if isObsFolder(key) {
						continue
					}
					if !c.force && !confirm(fmt.Sprintf("Do you want delete object [%s] ? Please input (y/n) to confirm:", key)) {
						continue
					}
					subWg.Add(1)
					atomic.AddInt64(totalCnt, 1)
					pool.ExecuteFunc(func() interface{} {
						ret := handleResult(c.deleteObject(bucket, key, "", 2), ch)
						subWg.Done()
						return ret
					})
				}

				for _, subFolder := range output.CommonPrefixes {
					if atomic.LoadInt32(&c.abort) == 1 {
						break
					}
					subWg.Add(1)
					c.doScan(bucket, subFolder, pool, ch, totalCnt, subWg)
				}

				if !output.IsTruncated {
					doLog(LEVEL_INFO, "List objects to delete finished, bucket [%s], folder [%s], marker [%s]", bucket, folder, input.Marker)
					break
				}
				input.Marker = output.NextMarker
			}
		}

		go func() {
			subWg.Wait()
			defer func() {
				wg.Done()
			}()

			if isObsFolder(folder) {
				atomic.AddInt64(totalCnt, 1)
				handleResult(c.deleteObject(bucket, folder, "", 2), ch)
			}
		}()

		return
	})
}

func (c *rmCommand) dropFolder(bucket, folder string) error {
	submitFunc := func(pool concurrent.Pool, ch progress.SingleBarChan) (totalCnt int64, err error) {
		c.scanPool = concurrent.NewNochanPool(-1)
		wg := new(sync.WaitGroup)
		wg.Add(1)
		c.doScan(bucket, folder, pool, ch, &totalCnt, wg)
		wg.Wait()
		c.scanPool.ShutDown()
		if _err, ok := c.scanError.Load().(error); ok {
			err = _err
		}

		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
		return
	}

	errorHandleFunc := func(hasListError error) {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to delete failed", bucket))
	}

	return c.recursiveAction(bucket, folder, submitFunc, errorHandleFunc, c.recordStartFunc, true)
}

func initRm() command {
	c := &rmCommand{}
	c.key = "rm"
	c.usage = c_cloud_url_usage
	c.description = "delete a bucket or objects in a bucket"
	c.define = func() {
		c.init()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
		c.flagSet.BoolVar(&c.version, "v", false, "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.versionId, "versionId", "", "")
		c.flagSet.IntVar(&c.jobs, "j", 0, "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		checkParamsFunc := func(prefix string) bool {
			_, ok := getRequestPayerType(c.payer)
			if !ok {
				return false
			}
			return true
		}

		emptyPrefixFunc := func(bucket string) error {
			if !c.deleteBucket(bucket) {
				return assist.ErrExecuting
			}
			return nil
		}

		confirmFunc := func(bucket, prefix string) bool {
			return confirm(fmt.Sprintf("Do you want delete object [%s] in the bucket [%s]? Please input (y/n) to confirm:", prefix, bucket))
		}
		prefixFunc := func(bucket, prefix string, batchFlag int) error {
			if c.deleteObject(bucket, prefix, c.versionId, batchFlag) {
				return nil
			}
			return assist.ErrExecuting
		}
		recursivePrefixFun := func(bucket, prefix string) error {
			fsStatus, err := c.checkBucketFSStatus(bucket)
			if err != nil {
				printError(err)
				return assist.ErrCheckBucketStatus
			}
			if fsStatus == c_disabled {
				return c.deleteObjects(bucket, prefix)
			}
			return c.dropFolder(bucket, prefix)
		}

		return c.chooseAction(checkParamsFunc, emptyPrefixFunc, confirmFunc, prefixFunc, recursivePrefixFun, c.recordStartFunc)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s\n", "", p.Sprintf("delete a bucket or objects in a bucket"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil rm obs://bucket [-f] [-config=xxx]")
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil rm obs://bucket/key [-f] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 3:")
		printf("%2s%s", "", "obsutil rm obs://bucket/[prefix] -r [-j=1] [-f] [-v] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch delete objects by prefix"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while a bucket or an object is to be deleted"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files when deleting one object"))
		printf("")
		printf("%2s%s", "", "-v")
		printf("%4s%s", "", p.Sprintf("batch delete versions of objects and the delete markers by prefix"))
		printf("")
		printf("%2s%s", "", "-versionId=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID of the object to be deleted"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent delete jobs, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the deleted results"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
