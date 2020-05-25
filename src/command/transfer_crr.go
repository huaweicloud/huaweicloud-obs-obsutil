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
	"bufio"
	"fmt"
	"io"
	"obs"
	"os"
	"progress"
	"ratelimit"
	"strings"
	"sync/atomic"
)

func (t *copyPartTask) copyPartCrr() (ret copyPartResult, noRepeatable bool, readed int64, copyPartError error) {
	input := &obs.GetObjectInput{}
	input.Bucket = t.srcBucket
	input.Key = t.srcKey
	input.VersionId = t.srcVersionId
	input.RangeStart = t.rangeStart
	input.RangeEnd = t.rangeEnd
	input.RequestPayer = t.payer

	output, err := obsClientCrr.GetObject(input)
	if err != nil {
		copyPartError = err
		if obsError, ok := copyPartError.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			atomic.CompareAndSwapInt32(t.abort, 0, 1)
			noRepeatable = true
		}
		return
	}

	defer output.Body.Close()
	uploadPartInput := &obs.UploadPartInput{}
	uploadPartInput.Bucket = t.dstBucket
	uploadPartInput.Key = t.dstKey
	uploadPartInput.UploadId = t.uploadId

	partSize := input.RangeEnd - input.RangeStart + 1

	_readBufferIoSize, transErr := assist.TranslateToInt64(config["readBufferIoSize"])
	if transErr != nil {
		_readBufferIoSize = defaultReadBufferIoSize
	}

	if _readBufferIoSize < minReadBufferIoSize {
		_readBufferIoSize = minReadBufferIoSize
	}
	if partSize < _readBufferIoSize {
		_readBufferIoSize = partSize
	}
	_body := progress.NewSingleProgressReader(bufio.NewReaderSize(output.Body, int(_readBufferIoSize)), -1, t.verifyMd5, t.barCh)
	var body io.Reader = _body
	if t.limiter != nil {
		body = ratelimit.NewRateLimitReaderWithLimiter(body, t.limiter)
	}
	uploadPartInput.Body = body
	uploadPartInput.PartSize = partSize
	uploadPartInput.PartNumber = t.partNumber
	uploadPartInput.RequestPayer = t.payer
	var uploadPartOutput *obs.UploadPartOutput
	uploadPartOutput, err = obsClient.UploadPart(uploadPartInput)

	readed += _body.ReadedCount

	if err != nil {
		copyPartError = err
		if obsError, ok := copyPartError.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			atomic.CompareAndSwapInt32(t.abort, 0, 1)
			noRepeatable = true
		}
		return
	}

	if changedErr := checkSourceChangedForCopyCrr(t.srcBucket, t.srcKey, t.srcVersionId, t.objectInfo.LastModified, t.abort, t.payer); changedErr != nil {
		copyPartError = &errorWrapper{
			err:       changedErr,
			requestId: uploadPartOutput.RequestId,
		}
		noRepeatable = true
		return
	}

	if t.verifyMd5 {
		sourceETag := _body.HexMd5()
		distinationETag := uploadPartOutput.ETag
		if !compareETag(sourceETag, distinationETag) {
			copyPartError = &errorWrapper{
				err:       &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after copying part [%d] of key [%s] in the bucket [%s], source md5 [%s] destination md5 [%s]", t.partNumber, t.srcKey, t.srcBucket, sourceETag, distinationETag)},
				requestId: uploadPartOutput.RequestId,
			}
			return
		}
	}
	ret = copyPartResult{
		PartEtag: PartEtag{
			PartNumber: t.partNumber,
			ETag:       uploadPartOutput.ETag,
		},
		metadata: output.Metadata,
	}

	return
}

func (c *transferCommand) createObsClientCrr() bool {
	if config["akCrr"] == defaultAccessKey || config["skCrr"] == defaultSecurityKey || config["endpointCrr"] == defaultEndpoint || config["endpointCrr"] == "" {
		printf("Warn: Please set akCrr, skCrr and endpointCrr in the configuration file!")
		return false
	}

	if err := refreshObsClientCrr(config["autoChooseSecurityProvider"] == c_true); err != nil {
		printError(err)
		return false
	}

	return true
}

func (c *transferCommand) getObjectMetadataCrr(bucket, key, versionId string, payer string) (*MetaContext, error) {
	return getObjectMetadataByClient(bucket, key, versionId, obsClientCrr, payer)
}

func (c *transferCommand) ensureBucketCrr(bucket string) error {
	if isAnonymousUserCrr() {
		return nil
	}

	return c.ensureBucketByClient(bucket, obsClientCrr)
}

func (c *transferCommand) ensureBucketsAndStartActionCrr(srcBucket string, dstBucket string, action func() error, recordCost bool) error {
	if err := c.ensureBucketCrr(srcBucket); err != nil {
		printError(err)
		doLog(LEVEL_ERROR, err.Error())
		return assist.ErrCheckBucketStatus
	}

	if err := c.ensureBucket(dstBucket); err != nil {
		printError(err)
		doLog(LEVEL_ERROR, err.Error())
		return assist.ErrCheckBucketStatus
	}

	return c.ensureOuputAndStartLogger(action, recordCost)
}

func checkSourceChangedForCopyCrr(srcBucket, srcKey, srcVersionId string, originLastModified int64, abort *int32, payer string) error {
	return checkSourceChangedForCopyByClient(srcBucket, srcKey, srcVersionId, originLastModified, abort, obsClientCrr, payer)
}

func (c *transferCommand) ensureObjectAttributesCrr(bucket, key, versionId string, srcMetaContext *MetaContext, metadata map[string]string) (map[string]string, string, obs.StorageClassType, string) {
	return c.ensureObjectAttributesByClient(bucket, key, versionId, srcMetaContext, metadata, obsClientCrr)
}

func (c *transferCommand) copyObjectCrr(srcBucket, srcKey, versionId, dstBucket, dstKey string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int, fastFailed error) int {
	srcMetaContext, srcMetaErr := c.getObjectMetadataCrr(srcBucket, srcKey, versionId, c.payer)
	return c.copyObjectCrrWithMetaContext(srcBucket, srcKey, versionId, srcMetaContext, srcMetaErr, dstBucket, dstKey, metadata, aclType, storageClassType, barCh, limiter, batchFlag, fastFailed)
}

func (c *transferCommand) copyObjectCrrWithMetaContext(srcBucket, srcKey, versionId string, srcMetaContext *MetaContext, srcMetaErr error,
	dstBucket, dstKey string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int, fastFailed error) int {
	var _versionId string
	if versionId != "" {
		_versionId = "?versionId=" + versionId
	}

	srcObjectSizeStr := c_na
	if srcMetaContext != nil {
		srcObjectSizeStr = normalizeBytes(srcMetaContext.Size)
	}

	if fastFailed != nil {
		c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, fastFailed.Error())
		return 0
	}

	if batchFlag == 2 && atomic.LoadInt32(&c.abort) == 1 {
		c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, error code [%s], error message [%s], n/a",
			srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, "AbortError", "Task is aborted")
		return 0
	}

	if c.update {
		changed, err := c.ensureKeyForCopy(srcMetaContext, srcMetaErr, dstBucket, dstKey)
		if !changed {
			if err == nil {
				if barCh != nil {
					if srcMetaContext.Size <= 0 {
						barCh.Send64(1)
					} else {
						barCh.Send64(srcMetaContext.Size)
					}
				}
				progress.AddFinishedStream(srcMetaContext.Size)
				if batchFlag >= 1 {
					c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, success message [skip since the source is not changed], n/a",
						srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey)
				}
				if batchFlag != 2 {
					printf("%s, obs://%s/%s%s --> obs://%s/%s, skip since the source is not changed",
						srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey)
				}
				return 2
			}
			if batchFlag >= 1 {
				c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [skip since the status of source is unknown], n/a",
					srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey)
			}
			if batchFlag != 2 {
				printf("obs://%s/%s%s --> obs://%s/%s, skip since the status of source is unknown",
					srcBucket, srcKey, _versionId, dstBucket, dstKey)
			}
			return 0
		}
	}

	copyObjectError := srcMetaErr

	if c.dryRun {
		if copyObjectError == nil {
			md5Value := c_na
			if c.verifyMd5 && srcMetaContext.Metadata != nil {
				if _md5Value, ok := srcMetaContext.Metadata[checkSumKey]; ok && _md5Value != "" {
					md5Value = _md5Value
				}
			}
			if barCh != nil {
				if srcMetaContext.Size <= 0 {
					barCh.Send64(1)
				} else {
					barCh.Send64(srcMetaContext.Size)
				}
			}
			progress.AddFinishedStream(srcMetaContext.Size)
			if batchFlag >= 1 {
				c.succeedLogger.doRecord("%s, %s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, success message [dry run done], n/a", srcObjectSizeStr, md5Value, srcBucket, srcKey,
					_versionId, dstBucket, dstKey)
			}
			if batchFlag != 2 {
				printf("\nCopy dry run successfully, %s, %s, obs://%s/%s%s --> obs://%s/%s", srcObjectSizeStr, md5Value, srcBucket, srcKey, _versionId, dstBucket, dstKey)
			}

			return 1
		}

		if batchFlag >= 1 {
			c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, n/a, error message [dry run done], n/a", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey)
		}
		if batchFlag != 2 {
			logError(copyObjectError, LEVEL_INFO, fmt.Sprintf("\nCopy dry run failed, obs://%s/%s%s --> obs://%s/%s", srcBucket, srcKey, _versionId, dstBucket, dstKey))
		}
		return 0
	}

	var requestId string
	var md5Value string
	var status int
	start := assist.GetUtcNow()
	if copyObjectError == nil {
		if srcMetaContext.Size >= c.bigfileThreshold || srcMetaContext.Size >= serverBigFileThreshold {
			status, requestId, md5Value, copyObjectError = c.copyBigObjectCrr(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh, limiter)
		} else {
			status, requestId, md5Value, copyObjectError = c.copySmallObjectCrrWithRetry(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh, limiter)
		}
	}

	if copyObjectError == nil {
		if barCh != nil && srcMetaContext.Size <= 0 {
			barCh.Send64(1)
		}

		if c.verifyMd5 && md5Value != "" {
			if obsVersion, ok := c.bucketsVersionMap[dstBucket]; ok && obsVersion == OBS_VERSION_UNKNOWN {
				warnMessage := fmt.Sprintf("Bucket [%s] cannot support setObjectMetadata interface, because of obs version check failed - so skip set object md5", dstBucket)
				warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
				c.recordWarnMessage(warnMessage, warnLoggerMessage)
			} else if ok && obsVersion >= "3.0" {
				if _, err := c.setObjectMd5(dstBucket, dstKey, "", md5Value, metadata); err != nil {
					_status, _code, _message, _requestId := getErrorInfo(err)
					warnMessage := fmt.Sprintf("Copy key [%s] in the bucket [%s] as key [%s] in the bucket [%s] successfully - but set object md5 failed status [%d] - error code [%s] - error message [%s] - request id [%s]",
						srcKey, srcBucket, dstKey, dstBucket, _status, _code, _message, _requestId)
					warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]",
						srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
					c.recordWarnMessage(warnMessage, warnLoggerMessage)
				}
			} else {
				warnMessage := fmt.Sprintf("Bucket [%s] cannot support setObjectMetadata interface - so skip set object md5", dstBucket)
				warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
				c.recordWarnMessage(warnMessage, warnLoggerMessage)
			}
		} else if c.verifyLength {
			if metaContext, err := getObjectMetadata(dstBucket, dstKey, "", c.payer); err == nil {
				if metaContext.Size != srcMetaContext.Size {
					doLog(LEVEL_ERROR, "Verify length failed after copying key [%s] in the bucket [%s], source length [%d] destination length [%d], will try to delete copied key", srcKey, srcBucket, srcMetaContext.Size, metaContext.Size)
					if _requestId, _err := c.deleteObject(dstBucket, dstKey, ""); _err == nil {
						doLog(LEVEL_INFO, "Delete key [%s] in the bucket [%s] successfully, request id [%s]", dstKey, dstBucket, _requestId)
					} else {
						_status, _code, _message, _requestId := getErrorInfo(_err)
						warnMessage := fmt.Sprintf("Delete key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]", dstKey, dstBucket, _status, _code, _message, _requestId)
						warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
						c.recordWarnMessage(warnMessage, warnLoggerMessage)
					}

					copyObjectError = &errorWrapper{
						err:       &verifyLengthError{msg: fmt.Sprintf("Verify length failed after copying key [%s] in the bucket [%s], source length [%d] destination length [%d]", srcKey, srcBucket, srcMetaContext.Size, metaContext.Size)},
						requestId: requestId,
					}
				}
			} else {
				warnMessage := fmt.Sprintf("Copy key [%s] in the bucket [%s] as key [%s] in the bucket [%s] successfully - but can not verify length - %s",
					srcKey, srcBucket, dstKey, dstBucket, err.Error())
				warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]",
					srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
				c.recordWarnMessage(warnMessage, warnLoggerMessage)
			}
		}
	}

	if md5Value == "" {
		md5Value = c_na
	}

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

	if batchFlag >= 1 {
		if copyObjectError == nil {
			c.succeedLogger.doRecord("%s, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], success message [copy done], request id [%s]", srcObjectSizeStr, md5Value, srcBucket, srcKey,
				_versionId, dstBucket, dstKey, cost, status, requestId)
		} else {
			_status, _code, _message, _requestId := c.checkAbort(copyObjectError, 401, 405)
			c.failedLogger.doRecord("%s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]",
				srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, cost, _status, _code, _message, _requestId)
		}
	}

	if batchFlag == 2 {
		c.ensureMaxCostAndMinCost(cost)
		atomic.AddInt64(&c.totalCost, cost)
	} else {
		if copyObjectError == nil {
			printf("\nCopy successfully, %s, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]", srcObjectSizeStr, md5Value, srcBucket, srcKey, _versionId, dstBucket, dstKey, cost, status, requestId)
			doLog(LEVEL_DEBUG, "Copy successfully, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
				srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, cost, status, requestId)
		} else {
			logError(copyObjectError, LEVEL_INFO, fmt.Sprintf("\nCopy failed, obs://%s/%s%s --> obs://%s/%s, cost [%d]", srcBucket, srcKey, _versionId, dstBucket, dstKey, cost))
		}
	}
	if copyObjectError == nil {
		return 1
	}
	return 0
}

func (c *transferCommand) prepareCopyObjectCheckpointCrr(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, cfc *CopyObjectCheckpoint) error {
	return c.prepareCopyObjectCheckpointByClient(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, cfc, obsClientCrr)
}

func (c *transferCommand) copyBigObjectCrr(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (status int, requestId string, md5Value string, copyObjectError error) {

	if srcMetaContext.Size == 0 {
		return c.copySmallObjectCrrWithRetry(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh, limiter)
	}

	checkpointFile := c.getCheckpointFile(dstBucket, dstKey, versionId, cm)
	cfc := &CopyObjectCheckpoint{}
	stat, err := os.Stat(checkpointFile)
	needPrepare := true
	if err == nil {
		if stat.IsDir() {
			copyObjectError = fmt.Errorf("Checkpoint file for copying [%s]-[%s] is a folder", dstBucket, dstKey)
			return
		}
		err = c.loadCheckpoint(checkpointFile, cfc)
		if err != nil {
			if err = os.Remove(checkpointFile); err != nil {
				copyObjectError = err
				return
			}
		} else if !cfc.isValid(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext) {
			if cfc.DestinationBucket != "" && cfc.DestinationKey != "" && cfc.UploadId != "" {
				if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
					copyObjectError = _err
					return
				}
			}
			if err = os.Remove(checkpointFile); err != nil {
				copyObjectError = err
				return
			}
		} else {
			needPrepare = false
		}
	}

	if needPrepare {
		err = c.prepareCopyObjectCheckpointCrr(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, cfc)
		if err != nil {
			copyObjectError = err
			return
		}
		err = c.recordCheckpointFile(checkpointFile, cfc)
		if err != nil {
			if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
				copyObjectError = _err
				return
			}
			copyObjectError = err
			return
		}
	}

	defer func() {
		if r := recover(); r != nil {
			c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId)
			panic(r)
		}
	}()

	barChFlag := false
	if barCh == nil {
		barCh = newSingleBarChan()
		barCh.SetTemplate(progress.SpeedOnly)
		barCh.SetBytes(true)
		barCh.SetTotalCount(cfc.ObjectInfo.Size)
		progress.SetTotalStream(cfc.ObjectInfo.Size)
		barChFlag = true
	}

	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	var abort int32
	var srcMetadata map[string]string
	abort, srcMetadata, copyObjectError = c.copyBigObjectConcurrent(cfc, checkpointFile, barChFlag, barCh, limiter)

	if barChFlag {
		barCh.Start()
	}

	if barChFlag {
		barCh.WaitToFinished()
	}
	if abort == 1 {
		if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
			copyObjectError = _err
			return
		}
		if err = os.Remove(checkpointFile); err != nil {
			copyObjectError = err
			return
		}
	}
	if copyObjectError != nil {
		return
	}

	if c.verifyMd5 {
		if _md5Value, ok := srcMetadata[checkSumKey]; ok && _md5Value != "" {
			md5Value = _md5Value
		} else {
			warnMessage := fmt.Sprintf("Cannot get the valid md5 value of key [%s] in bucket [%s] to check", cfc.SourceKey, cfc.SourceBucket)
			var _versionId string
			if versionId != "" {
				_versionId = "?versionId=" + versionId
			}
			srcObjectSizeStr := c_na
			if srcMetaContext != nil {
				srcObjectSizeStr = normalizeBytes(srcMetaContext.Size)
			}
			warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]",
				srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	}

	if barChFlag {
		h := &assist.HintV2{}
		h.Message = "Waiting for the copied object to be completed on server side"
		h.Start()
		defer h.End()
	}

	_status, _requestId, completeErr := c.completeMultipartUploadForCopyObject(cfc)
	if completeErr != nil {
		if obsError, ok := completeErr.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
				copyObjectError = _err
				return
			}
			if _err := os.Remove(checkpointFile); _err != nil {
				copyObjectError = _err
				return
			}
		}
		copyObjectError = completeErr
		return
	}
	if _err := os.Remove(checkpointFile); _err != nil {
		doLog(LEVEL_WARN, "Copy key [%s] in the bucket [%s] to key [%s] in the bucket [%s] successfully, but remove checkpoint file [%s] failed",
			cfc.SourceKey, cfc.SourceBucket, cfc.DestinationKey, cfc.DestinationBucket, checkpointFile)
	}
	requestId = _requestId
	status = _status
	return

}

func (c *transferCommand) copySmallObjectCrrWithRetry(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (status int, requestId string, md5Value string, copyObjectError error) {

	input := &obs.GetObjectMetadataInput{}
	input.Bucket = srcBucket
	input.Key = srcKey
	input.VersionId = versionId
	input.RequestPayer = c.payer

	output, err := obsClientCrr.GetObjectMetadata(input)
	if err != nil {
		copyObjectError = err
		return
	}

	objectSize := output.ContentLength

	if srcMetaContext.Size != objectSize {
		copyObjectError = fmt.Errorf("Object size changed, expect [%d], actual [%d], get object request id [%s]", srcMetaContext.Size, objectSize, output.RequestId)
		return
	}

	barChFlag := false
	if barCh == nil && objectSize > 0 {
		barCh = newSingleBarChan()
		barCh.SetBytes(true)
		barCh.SetTemplate(progress.SpeedOnly)
		barCh.SetTotalCount(objectSize)
		progress.SetTotalStream(objectSize)
		barCh.Start()
		barChFlag = true
	}

	var noRepeatable bool
	retryCount := 0
	var readed int64
	maxRetryCount := assist.StringToInt(config["maxRetryCount"], defaultMaxRetryCount)
	for {
		status, requestId, md5Value, noRepeatable, readed, copyObjectError = c.copySmallObjectCrr(srcBucket, srcKey, versionId, dstBucket, dstKey,
			srcMetaContext, metadata, aclType, storageClassType, barCh, limiter)
		if copyObjectError == nil || noRepeatable || retryCount >= maxRetryCount {
			break
		}

		if readed > 0 {
			barCh.Send64(-readed)
			progress.AddEffectiveStream(-readed)
			progress.AddFinishedStream(-readed)
		}

		doLog(LEVEL_WARN, "Failed to copy object and will try again, err:%s", copyObjectError.Error())
		assist.SleepByCount(retryCount)
		retryCount++
	}

	if barChFlag {
		barCh.WaitToFinished()
	}
	return
}

func (c *transferCommand) copySmallObjectCrr(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (status int, requestId string, md5Value string,
	noRepeatable bool, readed int64, copyObjectError error) {

	input := &obs.GetObjectInput{}
	input.Bucket = srcBucket
	input.Key = srcKey
	input.VersionId = versionId
	input.RequestPayer = c.payer
	output, err := obsClientCrr.GetObject(input)
	if err != nil {
		copyObjectError = err
		if obsError, ok := copyObjectError.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			noRepeatable = true
		}
		return
	}

	objectSize := output.ContentLength

	if srcMetaContext.Size != objectSize {
		copyObjectError = fmt.Errorf("Object size changed, expect [%d], actual [%d], get object request id [%s]", srcMetaContext.Size, objectSize, output.RequestId)
		noRepeatable = true
		return
	}

	_readBufferIoSize, transErr := assist.TranslateToInt64(config["readBufferIoSize"])
	if transErr != nil {
		_readBufferIoSize = defaultReadBufferIoSize
	}
	if _readBufferIoSize < minReadBufferIoSize {
		_readBufferIoSize = minReadBufferIoSize
	}
	if objectSize < _readBufferIoSize {
		_readBufferIoSize = objectSize
	}
	defer output.Body.Close()
	_body := progress.NewSingleProgressReader(bufio.NewReaderSize(output.Body, int(_readBufferIoSize)), -1, c.verifyMd5, barCh)
	var body io.Reader = _body
	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	if limiter != nil {
		body = ratelimit.NewRateLimitReaderWithLimiter(body, limiter)
	}

	putInput := &obs.PutObjectInput{}
	putInput.Bucket = dstBucket
	putInput.Key = dstKey
	putInput.ACL = aclType
	putInput.ContentLength = objectSize
	putInput.Body = body
	_metadata, contentType, storageClass, webredirectLocation := c.ensureObjectAttributesCrr(srcBucket, srcKey, versionId, srcMetaContext, metadata)
	putInput.ContentType = contentType
	putInput.WebsiteRedirectLocation = webredirectLocation
	putInput.Metadata = _metadata
	putInput.RequestPayer = c.payer
	if storageClassType == "" {
		putInput.StorageClass = storageClass
	} else {
		putInput.StorageClass = storageClassType
	}

	putOutput, err := obsClient.PutObject(putInput)

	if err != nil {
		copyObjectError = err
		if obsError, ok := copyObjectError.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			noRepeatable = true
		}
		return
	}

	if changedErr := checkSourceChangedForCopyCrr(srcBucket, srcKey, versionId, srcMetaContext.LastModified.Unix(), nil, c.payer); changedErr != nil {
		copyObjectError = &errorWrapper{
			err:       changedErr,
			requestId: putOutput.RequestId,
		}
		noRepeatable = true
		return
	}

	md5Value = _body.HexMd5()
	if c.verifyMd5 && !compareETag(md5Value, putOutput.ETag) {
		doLog(LEVEL_ERROR, "Verify md5 failed after copying key [%s] in the bucket [%s], source md5 [%s] destination md5 [%s], will try to delete copied key", srcKey, srcBucket, md5Value, putOutput.ETag)
		if deleteRequestId, err := c.deleteObject(dstBucket, dstKey, ""); err == nil {
			doLog(LEVEL_INFO, "Delete key [%s] in the bucket [%s] successfully, request id [%s]", dstKey, dstBucket, deleteRequestId)
		} else {
			status, code, message, deleteRequestId := getErrorInfo(err)

			var _versionId string
			if versionId != "" {
				_versionId = "?versionId=" + versionId
			}
			srcObjectSizeStr := c_na
			if srcMetaContext != nil {
				srcObjectSizeStr = normalizeBytes(srcMetaContext.Size)
			}
			warnMessage := fmt.Sprintf("Delete key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]", dstKey, dstBucket, status, code, message, deleteRequestId)
			warnLoggerMessage := fmt.Sprintf("%s, obs://%s/%s%s --> obs://%s/%s, warn message [%s]",
				srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, warnMessage)

			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
		copyObjectError = &errorWrapper{
			err:       &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after copying key [%s] in the bucket [%s], source md5 [%s] destination md5 [%s]", srcKey, srcBucket, md5Value, putOutput.ETag)},
			requestId: putOutput.RequestId,
		}
		return
	}

	requestId = putOutput.RequestId
	status = putOutput.StatusCode
	return
}
