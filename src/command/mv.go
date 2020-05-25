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
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type mvCommand struct {
	transferCommand
}
type ScanObsDirCtx struct {
	input          MoveRequestInput
	relativePrefix string
	pool           concurrent.Pool
	barCh          progress.SingleBarChan
	actionName     string
	action         func(request MoveRequestInput, srcMetaError error, barCh progress.SingleBarChan, batchFlag int, fastError error) int
}
type MoveRequestInput struct {
	src       ObsObjectCtx
	dst       ObsObjectCtx
	isPosix   bool
	canRename bool
}
type ObsObjectCtx struct {
	bucket           string
	key              string
	versionId        string
	aclType          obs.AclType
	storageClassType obs.StorageClassType
	metadata         map[string]string
	objectSizeStr    string
	count            int64
	metaContext      *MetaContext
}

func (c *mvCommand) recordStartFuncForMove() time.Time { return c.recordStartFuncForCopy() }

func (c *transferCommand) scanObsDirAndDoAction(scan ScanObsDirCtx) (totalCount int64, totalBytesForProgress int64, totalObjects int64, hasListError error) {
	defer func() {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
	}()
	input := &obs.ListObjectsInput{}
	request := scan.input
	input.Bucket = request.src.bucket
	input.Prefix = request.src.key
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
			doLog(LEVEL_INFO, "List objects in the bucket [%s] to [%s] successfully, cost [%d], request id [%s]", request.src.bucket, scan.actionName, cost, output.RequestId)
		}
		for _, content := range output.Contents {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			newRequest := request
			newRequest.src.key = content.Key
			if !isObsFolder(newRequest.src.key) || c.matchFolder {
				if c.matchExclude(newRequest.src.key) {
					continue
				}

				if !c.matchInclude(newRequest.src.key) {
					continue
				}

				if !c.matchLastModifiedTime(content.LastModified) {
					continue
				}
			}

			_srcKey := newRequest.src.key
			if scan.relativePrefix != "" {
				if index := strings.Index(_srcKey, scan.relativePrefix); index >= 0 {
					_srcKey = _srcKey[len(scan.relativePrefix):]
				}
			}

			_dstDir := newRequest.dst.key

			if isObsFolder(_dstDir) {
				_dstDir = _dstDir[:len(_dstDir)-1]
			}
			_srcKey = assist.MaybeDeleteBeginningSlash(_srcKey)
			newRequest.dst.key = _dstDir + "/" + _srcKey
			newRequest.dst.key = assist.MaybeDeleteBeginningSlash(newRequest.dst.key)
			var fastFailed error
			// what to do?
			if checkEmptyFolder("", newRequest.dst.key, cm) {
				fastFailed = fmt.Errorf("Cannot [%s] to the specified key [%s] in the bucket [%s]", scan.actionName, newRequest.dst.key, newRequest.dst.bucket)
			} else if checkEmptyFolder(request.src.bucket, newRequest.src.key, cm) {
				fastFailed = fmt.Errorf("Cannot [%s] the specified key [%s] in the bucket [%s]", scan.actionName, newRequest.src.key, newRequest.src.bucket)
			}

			if newRequest.dst.key == "" {
				continue
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want [%s] key [%s] in the bucket [%s] to key [%s] in the bucket [%s] ? Please input (y/n) to confirm:", scan.actionName,
				newRequest.src.key, newRequest.src.bucket, newRequest.dst.key, newRequest.dst.bucket)) {
				continue
			}

			srcMetaContext := &MetaContext{
				Size:         content.Size,
				ETag:         content.ETag,
				LastModified: content.LastModified,
			}
			count := c.caculateCount(content.Size, false)

			atomic.AddInt64(&totalCount, count)

			if srcMetaContext.Size == 0 {
				atomic.AddInt64(&totalBytesForProgress, 1)
			} else {
				atomic.AddInt64(&totalBytesForProgress, srcMetaContext.Size)
			}
			atomic.AddInt64(&totalObjects, 1)
			newRequest.src.metaContext = srcMetaContext
			newRequest.src.count = count

			scan.pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(scan.action(newRequest, nil, scan.barCh, 2, fastFailed), 0)
			})
		}

		if !output.IsTruncated {
			doLog(LEVEL_INFO, "List objects to [%s] finished, bucket [%s], folder [%s] prefix [%s], marker [%s]", scan.actionName,
				request.src.bucket, request.src.key, input.Prefix, input.Marker)
			break
		}
		input.Marker = output.NextMarker
	}
	return
}

func (c *mvCommand) submitMoveTask(request MoveRequestInput, relativePrefix string,
	barCh progress.SingleBarChan, pool concurrent.Pool) (totalCount int64, totalBytesForProgress int64, totalObjects int64, hasListError error) {
	return c.scanObsDirAndDoAction(ScanObsDirCtx{
		input:          request,
		relativePrefix: relativePrefix,
		barCh:          barCh,
		pool:           pool,
		actionName:     "move",
		action:         c.moveObjectWithMetaContext,
	})
}

func (c *mvCommand) doScanEmptyDir(bucket, dirKey string, totalCnt *int64, wg *sync.WaitGroup) {
	c.scanPool.ExecuteFunc(func() (r interface{}) {
		subWg := new(sync.WaitGroup)
		input := &obs.ListObjectsInput{}
		input.Bucket = bucket
		input.Prefix = dirKey
		input.Delimiter = "/"
		input.MaxKeys = defaultListMaxKeys
		input.RequestPayer = c.payer
		for {
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

			for _, subFolder := range output.CommonPrefixes {
				subWg.Add(1)
				c.doScanEmptyDir(bucket, subFolder, totalCnt, subWg)
			}

			if !output.IsTruncated {
				doLog(LEVEL_INFO, "List objects to delete finished, bucket [%s], folder [%s], marker [%s]", bucket, dirKey, input.Marker)
				break
			}
			input.Marker = output.NextMarker
		}
		go func() {
			subWg.Wait()
			defer func() {
				wg.Done()
			}()

			if isObsFolder(dirKey) {
				atomic.AddInt64(totalCnt, 1)
				requestId, delSrcError := c.deleteObject(bucket, dirKey, "")
				if delSrcError != nil {
					delSrcErrStatus, delSrcErrCode, delSrcErrMessage, delSrcErrRequestId := getErrorInfo(delSrcError)
					doLog(LEVEL_INFO, "Delete source key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]",
						dirKey, bucket, delSrcErrStatus, delSrcErrCode, delSrcErrMessage, delSrcErrRequestId)
				} else {
					doLog(LEVEL_INFO, "Delete source key [%s] in the bucket [%s] successfully - request id [%s]", dirKey, bucket, requestId)
				}
			}
		}()

		return
	})
}

func (c *mvCommand) cleanEmptyDir(request MoveRequestInput, totalCnt int64) error {
	c.scanPool = concurrent.NewNochanPool(-1)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	c.doScanEmptyDir(request.src.bucket, request.src.key, &totalCnt, wg)
	wg.Wait()
	c.scanPool.ShutDown()
	if _err, ok := c.scanError.Load().(error); ok {
		return _err
	}
	return nil
}

func (c *transferCommand) shouldCleanSource(request MoveRequestInput) bool {
	if !isObsFolder(request.src.key) {
		return true
	}

	return !request.isPosix
}

func (c *transferCommand) checkCloudUrlIfChange(request MoveRequestInput, srcMetaErr error, barCh progress.SingleBarChan, batchFlag int, _versionId string) int {
	changed, err := c.ensureKeyForCopy(request.src.metaContext, srcMetaErr, request.dst.bucket, request.dst.key)
	if !changed {
		if err == nil {
			if c.shouldCleanSource(request) {
				_, _, cleanError := c.cleanSource(request, _versionId)
				if cleanError != nil {
					if batchFlag >= 1 {
						c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a",
							request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cleanError.Error())
					}
					if batchFlag != 2 {
						printf("obs://%s/%s%s --> obs://%s/%s, %s",
							request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cleanError.Error())
					}
					return 0
				}
			}

			if barCh != nil {
				barCh.Send64(request.src.count)
			}
			if batchFlag >= 1 {
				c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, success message [skip since the source is not changed], n/a",
					request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
			}
			if batchFlag != 2 {
				printf("%s, obs://%s/%s%s --> obs://%s/%s, skip since the source is not changed",
					request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
			}
			return 2
		}
		if batchFlag >= 1 {
			c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [skip since the status of source is unknown], n/a",
				request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
		}
		if batchFlag != 2 {
			printf("obs://%s/%s%s --> obs://%s/%s, skip since the status of source is unknown",
				request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
		}
		return 0
	}
	return -1
}

func (c *transferCommand) doMove(request MoveRequestInput, barCh progress.SingleBarChan, _versionId string) (status int, requestId string, moveObjectError error) {
	if moveObjectError == nil {
		if request.src.metaContext.Size >= c.bigfileThreshold || request.src.metaContext.Size >= serverBigFileThreshold {
			status, requestId, moveObjectError = c.copyBigObject(request.src.bucket, request.src.key, request.src.versionId, request.dst.bucket, request.dst.key,
				request.src.metaContext, request.dst.metadata, request.dst.aclType, request.dst.storageClassType, barCh)
		} else {
			status, requestId, moveObjectError = c.copySmallObject(request.src.bucket, request.src.key, request.src.versionId, request.dst.bucket, request.dst.key,
				request.src.metaContext, request.dst.metadata, request.dst.aclType, request.dst.storageClassType, barCh)
		}
	}

	// delete source object
	if moveObjectError == nil && c.shouldCleanSource(request) {
		if _status, _requestId, err := c.cleanSource(request, _versionId); err != nil {
			status = _status
			requestId = _requestId
			moveObjectError = err
		}
	}
	return
}

func (c *transferCommand) cleanSource(request MoveRequestInput, _versionId string) (status int, requestId string, cleanObjectError error) {
	if deleteSrcSuccRequestId, err := c.deleteObject(request.src.bucket, request.src.key, request.src.versionId); err != nil {
		delSrcErrStatus, delSrcErrCode, delSrcErrMessage, delSrcErrRequestId := getErrorInfo(err)
		cleanObjectError = err
		requestId = delSrcErrRequestId
		status = delSrcErrStatus
		deleteErrorMessage := fmt.Sprintf("Delete source key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]",
			request.src.key, request.src.bucket, status, delSrcErrCode, delSrcErrMessage, delSrcErrRequestId)
		warnLoggerMessage := fmt.Sprintf("%s, obs:%s/%s%s --> obs://%s/%s, warn message [%s]", request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, deleteErrorMessage)
		c.recordWarnMessage(deleteErrorMessage, warnLoggerMessage)
		if delSrcErrStatus >= 300 { // ensure delete the src failed
			// delete dst object
			if deleteDstSuccRequestId, err := c.deleteObject(request.dst.bucket, request.dst.key, request.dst.versionId); err != nil {
				delDstErrStatus, delDstErrCode, delDstErrMessage, delDstErrRequestId := getErrorInfo(err)
				// delete warning, record warning
				deleteErrorMessage := fmt.Sprintf("Delete destination key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]",
					request.src.key, request.src.bucket, delDstErrStatus, delDstErrCode, delDstErrMessage, delDstErrRequestId)
				warnLoggerMessage := fmt.Sprintf("%s, obs:%s/%s%s --> obs://%s/%s, warn message [%s]", request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, deleteErrorMessage)
				c.recordWarnMessage(deleteErrorMessage, warnLoggerMessage)
			} else {
				doLog(LEVEL_INFO, "Delete destination key [%s] in the bucket [%s] successfully - request id [%s]", request.dst.key, request.dst.bucket, deleteDstSuccRequestId)
			}
		}
	} else {
		doLog(LEVEL_INFO, "Delete source key [%s] in the bucket [%s] successfully - request id [%s]", request.src.key, request.src.bucket, deleteSrcSuccRequestId)
	}
	return
}

func (c *mvCommand) moveDryRun(request MoveRequestInput, _versionId string, barCh progress.SingleBarChan, batchFlag int, moveObjectError error) int {
	if moveObjectError == nil {
		if barCh != nil {
			barCh.Send64(request.src.count)
		}
		if batchFlag >= 1 {
			c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, success message [dry run done], n/a",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
		}
		if batchFlag != 2 {
			printf("\nMove dry run successfully, %s, obs://%s/%s%s --> obs://%s/%s",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
		}

		return 1
	}

	if batchFlag >= 1 {
		c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [dry run done], n/a",
			request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key)
	}
	if batchFlag != 2 {
		logError(moveObjectError, LEVEL_INFO, fmt.Sprintf("\nMove dry run failed, obs://%s/%s%s --> obs://%s/%s",
			request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key))
	}
	return 0
}

func (c *mvCommand) moveObjectWithMetaContext(request MoveRequestInput, srcMetaErr error, barCh progress.SingleBarChan, batchFlag int, fastFailed error) int {
	var _versionId string
	if request.src.versionId != "" {
		_versionId = "?versionId=" + request.src.versionId
	}
	srcObjectSizeStr := c_na
	if request.src.metaContext != nil {
		srcObjectSizeStr = normalizeBytes(request.src.metaContext.Size)
	}
	request.src.objectSizeStr = srcObjectSizeStr
	if fastFailed != nil {
		c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", request.src.objectSizeStr, request.src.bucket,
			request.src.key, _versionId, request.dst.bucket, request.dst.key, fastFailed.Error())
		return 0
	}

	if batchFlag == 2 && atomic.LoadInt32(&c.abort) == 1 {
		c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, error code [%s], error message [%s], n/a",
			request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, "AbortError", "Task is aborted")
		return 0
	}

	moveObjectError := srcMetaErr
	if c.update {
		if changed := c.checkCloudUrlIfChange(request, srcMetaErr, barCh, batchFlag, _versionId); changed >= 0 {
			return changed
		}
	}

	if c.dryRun {
		return c.moveDryRun(request, _versionId, barCh, batchFlag, moveObjectError)
	}
	start := assist.GetUtcNow()
	// start move
	var requestId string
	var status int
	if moveObjectError == nil {
		if request.canRename && !isObsFolder(request.src.key) {
			status, requestId, moveObjectError = c.doRename(request, barCh)
		} else {
			status, requestId, moveObjectError = c.doMove(request, barCh, _versionId)
		}
	}
	// end move
	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
	// record and return
	return c.recordMoveRate(request, _versionId, moveObjectError, cost, status, requestId, batchFlag)
}

func (c *mvCommand) recordMoveRate(request MoveRequestInput, _versionId string, moveObjectError error, cost int64, status int, requestId string, batchFlag int) int {
	if batchFlag >= 1 {
		if moveObjectError == nil {
			c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], success message [move succeed], request id [%s]",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cost, status, requestId)
		} else {
			_status, _code, _message, _requestId := c.checkAbort(moveObjectError, 401, 405)
			c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cost, _status, _code, _message, _requestId)
		}
	}

	if batchFlag == 2 {
		c.ensureMaxCostAndMinCost(cost)
		atomic.AddInt64(&c.totalCost, cost)
	} else {
		if moveObjectError == nil {
			printf("\nMove successfully, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cost, status, requestId)
			doLog(LEVEL_DEBUG, "Move successfully, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
				request.src.objectSizeStr, request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cost, status, requestId)
		} else {
			logError(moveObjectError, LEVEL_INFO, fmt.Sprintf("\nMove failed, obs://%s/%s%s --> obs://%s/%s, cost [%d]",
				request.src.bucket, request.src.key, _versionId, request.dst.bucket, request.dst.key, cost))
		}
	}
	if moveObjectError == nil {
		return 1
	}
	return 0
}

func (c *mvCommand) moveDir(request MoveRequestInput) error {
	start := c.recordStartFuncForMove()
	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	barCh := newSingleBarChan()
	barCh.SetTemplate(progress.TpsOnly)
	if c.force {
		barCh.Start()
	}

	var relativePrefix string
	if c.flat {
		if request.src.key != "" && !isObsFolder(request.src.key) {
			request.src.key += "/"
		}
		relativePrefix = request.src.key
	} else {
		relativePrefix = request.src.key

		if isObsFolder(relativePrefix) {
			relativePrefix = relativePrefix[:len(relativePrefix)-1]
		}
		if index := strings.LastIndex(relativePrefix, "/"); index >= 0 {
			relativePrefix = relativePrefix[:index+1]
		} else {
			relativePrefix = ""
		}
	}

	totalCount, _, totalObjects, hasListError := c.submitMoveTask(request, relativePrefix, barCh, pool)

	doLog(LEVEL_INFO, "Number of objects to move [%d]", totalObjects)
	progress.SetTotalCount(totalObjects)
	barCh.SetTotalCount(totalCount)
	progress.SetTotalStream(-1)

	if !c.force {
		barCh.Start()
	}

	pool.ShutDown()
	barCh.WaitToFinished()

	if request.isPosix && atomic.LoadInt32(&c.abort) == 0 {
		// hint
		h := &assist.HintV2{}
		var totalDeletedCount int64
		h.MessageFunc = func() string {
			count := ""
			if tc := atomic.LoadInt64(&totalDeletedCount); tc > 0 {
				count = "[" + assist.Int64ToString(tc) + "]"
			}
			return fmt.Sprintf("Waitting for clean up move surplus %s", count)
		}
		h.Start()

		cleanError := c.cleanEmptyDir(request, totalDeletedCount)
		h.End()

		if cleanError != nil {
			logError(cleanError, LEVEL_ERROR, fmt.Sprintf("\nList source dir from bucket [%s] to clean failed", request.src.bucket))
			return assist.ErrUncompeleted
		}
	}

	c.recordEndWithMetricsV2(start, totalObjects, progress.GetSucceedStream(), progress.GetTotalStream())
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects from bucket [%s] to move failed", request.src.bucket))
		return assist.ErrUncompeleted
	}

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}
	return nil
}

// batchFlag: {0 single object, 1 single object and must record, 2 multi object}
func (c *mvCommand) moveObject(request MoveRequestInput, barCh progress.SingleBarChan, batchFlag int, fastFailed error) int {
	srcMetaContext, srcMetaErr := getObjectMetadata(request.src.bucket, request.src.key, request.src.versionId, c.payer)
	request.src.metaContext = srcMetaContext
	var count int64 = 1
	if srcMetaErr == nil {
		count = c.caculateCount(request.src.metaContext.Size, false)
	}
	request.src.count = count

	return c.moveObjectWithMetaContext(request, srcMetaErr, barCh, batchFlag, fastFailed)
}

func (c *transferCommand) doRename(request MoveRequestInput, barCh progress.SingleBarChan) (status int, requestId string, renameError error) {
	input := &obs.RenameFileInput{}
	input.Bucket = request.src.bucket
	input.Key = request.src.key
	input.NewObjectKey = request.dst.key
	input.RequestPayer = c.payer
	if output, err := obsClient.RenameFile(input); err != nil {
		if index := strings.Index(input.NewObjectKey, "/"); index >= 0 {
			// the destination parant folder may not exist
			if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == 404 {
				newFolderInput := &obs.NewFolderInput{}
				newFolderInput.Bucket = request.dst.bucket
				newFolderInput.Key = input.NewObjectKey[:index+1]
				newFolderInput.RequestPayer = c.payer
				_, newFolderErr := obsClient.NewFolder(newFolderInput)
				if newFolderErr == nil {
					if _output, _err := obsClient.RenameFile(input); _err != nil {
						renameError = err
					} else {
						if barCh != nil {
							barCh.Send(1)
						}
						status = _output.StatusCode
						requestId = _output.RequestId
					}
					return
				}
				doLogError(newFolderErr, LEVEL_DEBUG, fmt.Sprintf("Try to create new folder [%s] failed", newFolderInput.Key))
			}
		}
		renameError = err
	} else {
		if barCh != nil {
			barCh.Send(1)
		}
		status = output.StatusCode
		requestId = output.RequestId
	}
	return
}

func (c *transferCommand) doRenameDir(request MoveRequestInput) (status int, requestId string, renameError error) {
	input := &obs.RenameFolderInput{}
	input.Bucket = request.src.bucket
	input.Key = request.src.key
	input.NewObjectKey = request.dst.key
	input.RequestPayer = c.payer
	if output, err := obsClient.RenameFolder(input); err != nil {
		renameError = err
	} else {
		status = output.StatusCode
		requestId = output.RequestId
	}
	return
}

func (c *mvCommand) movePosixDir(moveRequestInput MoveRequestInput) error {

	if !c.force && !confirm(fmt.Sprintf("Do you want move key [%s] in the bucket [%s] to key [%s] in the bucket [%s] ? Please input (y/n) to confirm:",
		moveRequestInput.src.key, moveRequestInput.src.bucket, moveRequestInput.dst.key, moveRequestInput.dst.bucket)) {
		return nil
	}

	if c.dryRun {
		printf("\nMove dry run done, obs://%s/%s --> obs://%s/%s", moveRequestInput.src.bucket, moveRequestInput.src.key, moveRequestInput.dst.bucket, moveRequestInput.dst.key)
		doLog(LEVEL_DEBUG, "Move dry run done, %s, obs://%s/%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
			moveRequestInput.src.bucket, moveRequestInput.src.key, moveRequestInput.dst.bucket, moveRequestInput.dst.key)
		return nil
	}

	start := assist.GetUtcNow()
	status, requestId, renameFolderError := c.doRenameDir(moveRequestInput)
	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
	if renameFolderError == nil {
		printf("\nMove successfully, obs://%s/%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
			moveRequestInput.src.bucket, moveRequestInput.src.key, moveRequestInput.dst.bucket, moveRequestInput.dst.key, cost, status, requestId)
		doLog(LEVEL_DEBUG, "Move successfully, %s, obs://%s/%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
			moveRequestInput.src.bucket, moveRequestInput.src.key, moveRequestInput.dst.bucket, moveRequestInput.dst.key, cost, status, requestId)
		return nil
	}
	logError(renameFolderError, LEVEL_ERROR, fmt.Sprintf("\nMove failed, obs://%s/%s --> obs://%s/%s, cost [%d]",
		moveRequestInput.src.bucket, moveRequestInput.src.key, moveRequestInput.dst.bucket, moveRequestInput.dst.key, cost))
	return assist.ErrExecuting
}

func initMv() command {
	c := &mvCommand{}
	c.key = "mv"
	c.usage = "cloud_url cloud_url [options...]"
	c.description = "move objects"
	c.define = func() {
		c.init()
		c.defineBasic()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.update, "u", false, "")
		c.flagSet.BoolVar(&c.flat, "flat", false, "")
		c.flagSet.StringVar(&c.versionId, "versionId", "", "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) <= 1 {
			c.showHelp()
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		url1, url2, mode, err := c.prepareUrls(args)

		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		_, succeed := getRequestPayerType(c.payer)
		if !succeed {
			return assist.ErrInvalidArgs
		}

		if mode != cm {
			printf("Error: The source url [%s] or the destination url [%s] is not the valid cloud_url!", url1, url2)
			return assist.ErrInvalidArgs
		}

		if url1 == url2 {
			printf("Error: The source url [%s] and the destination url [%s] are same!", url1, url2)
			return assist.ErrInvalidArgs
		}

		if !c.prepareOptions() {
			return assist.ErrInvalidArgs
		}

		aclType, storageClassType, metadata, succeed := c.checkParams()
		if !succeed {
			return assist.ErrExecuting
		}

		srcBucket, srcKeyOrDir, err := c.splitCloudUrl(url1)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		dstBucket, dstKeyOrDir, err := c.splitCloudUrl(url2)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		c.printStart()
		// start prepare move params
		src := ObsObjectCtx{
			bucket:    srcBucket,
			key:       srcKeyOrDir,
			versionId: c.versionId,
		}
		dst := ObsObjectCtx{
			bucket:           dstBucket,
			key:              dstKeyOrDir,
			aclType:          aclType,
			storageClassType: storageClassType,
			metadata:         metadata,
		}

		fsStatus, err := c.checkBucketFSStatus(src.bucket)
		if err != nil {
			printError(err)
			return assist.ErrCheckBucketStatus
		}

		moveRequestInput := MoveRequestInput{
			src:       src,
			dst:       dst,
			isPosix:   fsStatus == c_enabled,
			canRename: src.bucket == dst.bucket && fsStatus == c_enabled,
		}

		//TODO
		if c.jobs > 10 {
			printf("Error: The max jobs for move is 10")
			return assist.ErrInvalidArgs
		}
		if c.parallel > 10 {
			printf("Error: The max parallel for move is 10")
			return assist.ErrInvalidArgs
		}

		// create parent folder if need
		c.ensureParentFolder(dst.bucket, dst.key)
		// end prepare move params
		// start move
		if !c.recursive {
			// start move object
			if src.key == "" {
				printf("Error: The source object key is empty!")
				return assist.ErrInvalidArgs
			}

			if moveRequestInput.canRename {
				input := &obs.GetAttributeInput{}
				input.Bucket = src.bucket
				input.Key = src.key
				input.RequestPayer = c.payer
				if output, err := obsClient.GetAttribute(input); err != nil {
					printError(err)
					return assist.ErrExecuting
				} else if output.Mode != -1 && output.Mode&0040000 != 0 {
					printf("Error: Must pass -r to move folder!")
					return assist.ErrInvalidArgs
				}
			}

			if dst.key == "" || isObsFolder(dst.key) {
				if index := strings.LastIndex(src.key, "/"); index >= 0 {
					moveRequestInput.dst.key += src.key[index+1:]
				} else {
					moveRequestInput.dst.key += src.key
				}
			}

			if c.forceRecord {
				if succeed := c.compareLocation(src.bucket, dst.bucket); !succeed {
					return assist.ErrInvalidArgs
				}
				return c.ensureBucketsAndStartAction([]string{src.bucket, dst.bucket}, func() error {
					c.printParams(true, false, false, false)
					c.recordStartFuncForMove()
					ret := c.moveObject(moveRequestInput, nil, 1, nil)
					if ret >= 1 {
						progress.AddSucceedCount(1)
						return nil
					}
					progress.AddFailedCount(1)
					return assist.ErrExecuting
				}, true)
			}
			c.printParams(false, false, false, false)
			ret := c.moveObject(moveRequestInput, nil, 0, nil)
			if warn, ok := c.warn.Load().(error); ok {
				printWarn(warn)
			}
			if ret == 0 {
				return assist.ErrExecuting
			}
			// end move object
			return nil
		}

		if dst.bucket == src.bucket {
			_url1 := url1
			_url2 := url2

			if !isObsFolder(_url1) {
				_url1 += "/"
			}

			if !isObsFolder(_url2) {
				_url2 += "/"
			}

			if index := strings.Index(_url2, _url1); index >= 0 {
				printf("The source cloud_url and the destination cloud_url are nested")
				return assist.ErrInvalidArgs
			}

			if index := strings.Index(_url1, _url2); index >= 0 {
				printf("The source cloud_url and the destination cloud_url are nested")
				return assist.ErrInvalidArgs
			}
		}

		if succeed := c.compareLocation(src.bucket, dst.bucket); !succeed {
			return assist.ErrInvalidArgs
		}

		if moveRequestInput.canRename && c.exclude == "" && c.include == "" && c.timeRange == "" {
			if c.flat {
				if dst.key == "" {
					// can not rename, need to copy and delete
					return c.ensureBucketsAndStartAction([]string{src.bucket, dst.bucket}, func() error {
						c.printParams(true, false, false, false)
						doLog(LEVEL_INFO, "Move objects from cloud folder [%s] in the bucket [%s] to cloud folder [%s] in the bucket [%s] ",
							src.key, src.bucket, dst.key, dst.bucket)
						return c.moveDir(moveRequestInput)
					}, false)
				}
			} else {
				srcFolder := ""
				if index := strings.LastIndex(src.key, "/"); index >= 0 {
					srcFolder = src.key[index+1:]
				} else {
					srcFolder = src.key
				}
				if dst.key == "" {
					moveRequestInput.dst.key = srcFolder
				} else {
					if !isObsFolder(dst.key) {
						moveRequestInput.dst.key += "/"
					}
					moveRequestInput.dst.key += srcFolder
				}
			}

			return c.movePosixDir(moveRequestInput)
		}
		// can not rename, need to copy and delete
		return c.ensureBucketsAndStartAction([]string{src.bucket, dst.bucket}, func() error {
			c.printParams(true, false, false, false)
			doLog(LEVEL_INFO, "Move objects from cloud folder [%s] in the bucket [%s] to cloud folder [%s] in the bucket [%s] ",
				src.key, src.bucket, dst.key, dst.bucket)
			return c.moveDir(moveRequestInput)
		}, false)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("move objects"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil mv obs://srcbucket/key obs://dstbucket/[dest] [-dryRun] [-u] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil mv obs://srcbucket[/src_prefix] obs://dstbucket[/dest_prefix] -r [-dryRun] [-f] [-u] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-cpd=xxx] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-dryRun")
		printf("%4s%s", "", p.Sprintf("conduct a dry run"))
		printf("")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch move objects by prefix"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while objects are to be moved"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files"))
		printf("")
		printf("%2s%s", "", "-u")
		printf("%4s%s", "", p.Sprintf("move the changed sources only"))
		printf("")
		printf("%2s%s", "", "-flat")
		printf("%4s%s", "", p.Sprintf("move the sources without the relative parent prefix"))
		printf("")
		printf("%2s%s", "", "-cpd=xxx")
		printf("%4s%s", "", p.Sprintf("the directory where the part records reside, used to record the progress of movement jobs"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of movement jobs, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-p=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent movement tasks (a task is a sub-job), the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-versionId=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID of the object to be moved"))
		printf("")
		printf("%2s%s", "", "-ps=auto")
		printf("%4s%s", "", p.Sprintf("the part size of each movement task, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-threshold=52428800")
		printf("%4s%s", "", p.Sprintf("the threshold, if it is exceeded, the movement job will be divided into multiple tasks by the part size, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-meta=aaa:bbb#ccc:ddd")
		printf("%4s%s", "", p.Sprintf("the customized metadata of each object to be moved"))
		printf("")
		printf("%2s%s", "", "-acl=xxx")
		printf("%4s%s", "", p.Sprintf("the ACL of each object to be moved, possible values are [private|public-read|public-read-write|bucket-owner-full-control]"))
		printf("")
		printf("%2s%s", "", "-sc=xxx")
		printf("%4s%s", "", p.Sprintf("the storage class of each object to be moved, possible values are [standard|warm|cold]"))
		printf("")
		printf("%2s%s", "", "-include=*.xxx")
		printf("%4s%s", "", p.Sprintf("the to be moved objects whose names match this pattern will be included"))
		printf("")
		printf("%2s%s", "", "-exclude=*.xxx")
		printf("%4s%s", "", p.Sprintf("the to be moved objects whose names match this pattern will be excluded"))
		printf("")
		printf("%2s%s", "", "-timeRange=time1-time2")
		printf("%4s%s", "", p.Sprintf("the time range, between which the objects will be moved"))
		printf("")
		printf("%2s%s", "", "-mf")
		printf("%4s%s", "", p.Sprintf("the including pattern, the excluding pattern and the time range pattern will task effect on folders"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the movement results"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}
	return c
}
