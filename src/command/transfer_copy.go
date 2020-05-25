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
	"concurrent"
	"encoding/xml"
	"fmt"
	"obs"
	"os"
	"progress"
	"ratelimit"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type CopyPart struct {
	PartEtag
	XMLName     xml.Name `xml:"CopyPart"`
	RangeStart  int64    `xml:"RangeStart"`
	RangeEnd    int64    `xml:"RangeEnd"`
	IsCompleted bool     `xml:"IsCompleted"`
}

type CopyObjectCheckpoint struct {
	XMLName           xml.Name   `xml:"CopyObjectCheckpoint"`
	SourceBucket      string     `xml:"SourceBucket"`
	SourceKey         string     `xml:"SourceKey"`
	SourceVersionId   string     `xml:"SourceVersionId"`
	DestinationBucket string     `xml:"DestinationBucket"`
	DestinationKey    string     `xml:"DestinationKey"`
	UploadId          string     `xml:"UploadId,omitempty"`
	ObjectInfo        ObjectInfo `xml:"ObjectInfo"`
	CopyParts         []CopyPart `xml:"CopyParts>CopyPart"`
}

func (cfc *CopyObjectCheckpoint) isValid(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext) bool {
	if cfc.SourceBucket != srcBucket || cfc.SourceKey != srcKey || cfc.SourceVersionId != versionId || cfc.DestinationBucket != dstBucket || cfc.DestinationKey != dstKey {
		return false
	}

	if cfc.ObjectInfo.Size != srcMetaContext.Size || cfc.ObjectInfo.ETag != srcMetaContext.ETag ||
		cfc.ObjectInfo.LastModified != srcMetaContext.LastModified.Unix() {
		return false
	}

	if cfc.UploadId == "" {
		return false
	}

	return true
}

type copyPartTask struct {
	dstBucket    string
	dstKey       string
	srcBucket    string
	srcKey       string
	srcVersionId string
	uploadId     string
	partNumber   int
	rangeStart   int64
	rangeEnd     int64
	abort        *int32
	barCh        progress.SingleBarChan
	limiter      *ratelimit.RateLimiter
	verifyMd5    bool
	crr          bool
	objectInfo   ObjectInfo
	payer        string
}

type copyPartResult struct {
	PartEtag
	metadata map[string]string
}

func (t *copyPartTask) Run() interface{} {
	if atomic.LoadInt32(t.abort) == 1 {
		return errAbort
	}

	if !t.crr {
		input := &obs.CopyPartInput{}
		input.Bucket = t.dstBucket
		input.Key = t.dstKey
		input.CopySourceBucket = t.srcBucket
		input.CopySourceKey = t.srcKey
		input.CopySourceVersionId = t.srcVersionId
		input.UploadId = t.uploadId
		input.CopySourceRangeStart = t.rangeStart
		input.CopySourceRangeEnd = t.rangeEnd
		input.PartNumber = t.partNumber
		input.RequestPayer = t.payer
		output, err := obsClient.CopyPart(input)
		if err == nil {
			if changedErr := checkSourceChangedForCopy(t.srcBucket, t.srcKey, t.srcVersionId, t.objectInfo.LastModified, t.abort, t.payer); changedErr != nil {
				return &errorWrapper{
					err:       changedErr,
					requestId: output.RequestId,
				}
			}

			return PartEtag{
				PartNumber: t.partNumber,
				ETag:       output.ETag,
			}
		}

		if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			atomic.CompareAndSwapInt32(t.abort, 0, 1)
		}

		doLogError(err, LEVEL_ERROR, fmt.Sprintf("SrcBucket [%s], SrcKey [%s], DstBucket [%s], DstKey [%s], UploadId [%s], PartNumber [%d]", t.srcBucket, t.srcKey, t.dstBucket, t.dstKey, t.uploadId, t.partNumber))

		return err
	}

	var result interface{}
	retryCount := 0
	maxRetryCount := assist.StringToInt(config["maxRetryCount"], defaultMaxRetryCount)
	for {
		ret, noRepeatable, readed, copyPartError := t.copyPartCrr()
		if copyPartError == nil {
			result = ret
			break
		}
		if noRepeatable || retryCount >= maxRetryCount {
			result = copyPartError
			doLogError(copyPartError, LEVEL_ERROR, fmt.Sprintf("SrcBucket [%s], SrcKey [%s], DstBucket [%s], DstKey [%s], UploadId [%s], PartNumber [%d]", t.srcBucket, t.srcKey, t.dstBucket, t.dstKey, t.uploadId, t.partNumber))
			break
		}

		if readed > 0 {
			t.barCh.Send64(-readed)
			progress.AddEffectiveStream(-readed)
			progress.AddFinishedStream(-readed)
		}

		doLog(LEVEL_WARN, "Failed to copy part crr and will try again, err:%s", copyPartError.Error())
		assist.SleepByCount(retryCount)
		retryCount++
	}

	return result
}

func checkSourceChangedForCopy(srcBucket, srcKey, srcVersionId string, originLastModified int64, abort *int32, payer string) error {
	return checkSourceChangedForCopyByClient(srcBucket, srcKey, srcVersionId, originLastModified, abort, obsClient, payer)
}

func checkSourceChangedForCopyByClient(srcBucket, srcKey, srcVersionId string, originLastModified int64, abort *int32, client *obs.ObsClient, payer string) error {
	if config["checkSourceChange"] == c_true {
		if metaContext, err := getObjectMetadataByClient(srcBucket, srcKey, srcVersionId, client, payer); err != nil {
			if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == 404 {
				if abort != nil {
					atomic.CompareAndSwapInt32(abort, 0, 1)
				}
				return fmt.Errorf("Source object [%s] in the bucket [%s] doesnot exist", srcKey, srcBucket)
			}
		} else if originLastModified != metaContext.LastModified.Unix() {
			if abort != nil {
				atomic.CompareAndSwapInt32(abort, 0, 1)
			}
			return fmt.Errorf("Source object [%s] in the bucket [%s] changed", srcKey, srcBucket)
		}
	}
	return nil
}

func (c *transferCommand) ensureKeyForCopy(srcMetaContext *MetaContext, srcMetaErr error, dstBucket string, dstKey string) (bool, error) {
	var changed bool
	if srcMetaErr == nil {
		dstMetaContext, dstMetaErr := getObjectMetadata(dstBucket, dstKey, "", c.payer)
		if dstMetaErr != nil {
			changed = true
		} else {
			changed = srcMetaContext.Size != dstMetaContext.Size || srcMetaContext.LastModified.After(dstMetaContext.LastModified)
		}
	} else if obsError, ok := srcMetaErr.(obs.ObsError); ok && obsError.StatusCode >= 300 && obsError.StatusCode < 500 && obsError.StatusCode != 408 {
		changed = true
	} else {
		changed = false
	}
	return changed, srcMetaErr
}

func (c *transferCommand) copySmallObject(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, barCh progress.SingleBarChan) (int, string, error) {
	input := &obs.CopyObjectInput{}
	input.Bucket = dstBucket
	input.Key = dstKey
	input.CopySourceBucket = srcBucket
	input.CopySourceKey = srcKey
	input.ACL = aclType
	input.CopySourceVersionId = versionId
	input.RequestPayer = c.payer

	if metadata == nil || len(metadata) == 0 {
		input.StorageClass = storageClassType
		input.MetadataDirective = obs.CopyMetadata
	} else {
		input.MetadataDirective = obs.ReplaceMetadata
		_metadata, contentType, storageClass, webredirectLocation := c.ensureObjectAttributes(srcBucket, srcKey, versionId, srcMetaContext, metadata)
		input.ContentType = contentType
		input.WebsiteRedirectLocation = webredirectLocation
		input.Metadata = _metadata
		if storageClassType == "" {
			input.StorageClass = storageClass
		} else {
			input.StorageClass = storageClassType
		}
	}
	output, err := obsClient.CopyObject(input)
	if err == nil {
		if changedErr := checkSourceChangedForCopy(srcBucket, srcKey, versionId, srcMetaContext.LastModified.Unix(), nil, c.payer); changedErr != nil {
			return 0, "", &errorWrapper{
				err:       changedErr,
				requestId: output.RequestId,
			}
		}
		if barCh != nil {
			if config["showBytesForCopy"] == c_true {
				if srcMetaContext.Size <= 0 {
					barCh.Send64(1)
				} else {
					barCh.Send64(srcMetaContext.Size)
				}
				progress.AddEffectiveStream(srcMetaContext.Size)
				progress.AddFinishedStream(srcMetaContext.Size)
			} else {
				barCh.Send(1)
			}
		}
	}
	if err != nil {
		return 0, "", err
	}
	return output.StatusCode, output.RequestId, nil
}

func (c *transferCommand) prepareCopyObjectCheckpoint(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, cfc *CopyObjectCheckpoint) error {
	return c.prepareCopyObjectCheckpointByClient(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, cfc, obsClient)
}

func (c *transferCommand) prepareCopyObjectCheckpointByClient(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, cfc *CopyObjectCheckpoint, client *obs.ObsClient) error {

	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = dstBucket
	input.Key = dstKey
	_metadata, contentType, storageClass, webredirectLocation := c.ensureObjectAttributesByClient(srcBucket, srcKey, versionId, srcMetaContext, metadata, client)
	input.ContentType = contentType
	input.WebsiteRedirectLocation = webredirectLocation
	input.Metadata = _metadata
	input.ACL = aclType
	input.RequestPayer = c.payer
	if storageClassType == "" {
		input.StorageClass = storageClass
	} else {
		input.StorageClass = storageClassType
	}
	output, err := obsClient.InitiateMultipartUpload(input)
	if err != nil {
		return err
	}
	doLog(LEVEL_DEBUG, "Initiate multipart upload [%s] in the bucket [%s] successfully, request id [%s]", dstKey, dstBucket, output.RequestId)

	cfc.DestinationBucket = dstBucket
	cfc.DestinationKey = dstKey
	cfc.SourceBucket = srcBucket
	cfc.SourceKey = srcKey
	cfc.SourceVersionId = versionId
	cfc.UploadId = output.UploadId
	cfc.ObjectInfo = ObjectInfo{
		Size:         srcMetaContext.Size,
		LastModified: srcMetaContext.LastModified.Unix(),
		ETag:         srcMetaContext.ETag,
	}
	partSize := c.autoSelectPartSize(cfc.ObjectInfo.Size, cm)
	count := cfc.ObjectInfo.Size / partSize
	if count >= 10000 {
		partSize = cfc.ObjectInfo.Size / 10000
		if cfc.ObjectInfo.Size%10000 != 0 {
			partSize++
		}
		count = cfc.ObjectInfo.Size / partSize
	}

	if cfc.ObjectInfo.Size%partSize != 0 {
		count++
	}

	if partSize > serverBigFileThreshold {
		return fmt.Errorf("The source key [%s] in bucket [%s] is too large", srcKey, srcBucket)
	}

	copyParts := make([]CopyPart, 0, count)
	var i int64
	for i = 0; i < count; i++ {
		copyPart := CopyPart{
			RangeStart: i * partSize,
			RangeEnd:   (i+1)*partSize - 1,
		}
		copyPart.PartNumber = int(i) + 1
		copyParts = append(copyParts, copyPart)
	}
	if lastPartSize := cfc.ObjectInfo.Size % partSize; lastPartSize != 0 {
		copyParts[count-1].RangeEnd = cfc.ObjectInfo.Size - 1
	}
	cfc.CopyParts = copyParts

	return nil
}

func (c *transferCommand) handleCopyPartResult(cfc *CopyObjectCheckpoint, checkpointFile string, result interface{}, lock *sync.Mutex) (metadata map[string]string, copyPartError error) {
	if partETag, ok := result.(PartEtag); ok {
		lock.Lock()
		defer lock.Unlock()
		cfc.CopyParts[partETag.PartNumber-1].IsCompleted = true
		cfc.CopyParts[partETag.PartNumber-1].ETag = partETag.ETag
		copyPartError = c.recordCheckpointFile(checkpointFile, cfc)
	} else if ret, ok := result.(copyPartResult); ok {
		lock.Lock()
		defer lock.Unlock()
		cfc.CopyParts[ret.PartNumber-1].IsCompleted = true
		cfc.CopyParts[ret.PartNumber-1].ETag = ret.ETag
		metadata = ret.metadata
		copyPartError = c.recordCheckpointFile(checkpointFile, cfc)
	} else if result != errAbort {
		if resultError, ok := result.(error); ok {
			copyPartError = resultError
		}
	}
	return
}

func (c *transferCommand) completeMultipartUploadForCopyObject(cfc *CopyObjectCheckpoint) (int, string, error) {
	input := &obs.CompleteMultipartUploadInput{}
	input.Bucket = cfc.DestinationBucket
	input.Key = cfc.DestinationKey
	input.UploadId = cfc.UploadId
	input.RequestPayer = c.payer
	parts := make([]obs.Part, 0, len(cfc.CopyParts))
	for _, copyPart := range cfc.CopyParts {
		part := obs.Part{
			ETag:       copyPart.ETag,
			PartNumber: copyPart.PartNumber,
		}
		parts = append(parts, part)
	}
	input.Parts = parts
	output, err := obsClient.CompleteMultipartUpload(input)
	if err == nil {
		doLog(LEVEL_DEBUG, "Complete multipart upload [%s] in the bucket [%s] successfully, request id [%s]", cfc.DestinationKey, cfc.DestinationBucket, output.RequestId)
		return output.StatusCode, output.RequestId, nil
	}
	return 0, "", err
}

func (c *transferCommand) copyBigObjectConcurrent(cfc *CopyObjectCheckpoint, checkpointFile string, barChFlag bool,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (int32, map[string]string, error) {
	pool := concurrent.NewRoutinePool(c.parallel, defaultParallelsCacheCount)

	var copyPartError atomic.Value
	var copyPartErrorFlag int32
	var abort int32
	var metadata atomic.Value
	var metadataFlag int32
	lock := new(sync.Mutex)
	for _, copyPart := range cfc.CopyParts {
		if atomic.LoadInt32(&abort) == 1 {
			break
		}
		if !copyPart.IsCompleted {
			task := &copyPartTask{
				srcBucket:    cfc.SourceBucket,
				srcKey:       cfc.SourceKey,
				srcVersionId: cfc.SourceVersionId,
				dstBucket:    cfc.DestinationBucket,
				dstKey:       cfc.DestinationKey,
				uploadId:     cfc.UploadId,
				partNumber:   copyPart.PartNumber,
				rangeStart:   copyPart.RangeStart,
				rangeEnd:     copyPart.RangeEnd,
				abort:        &abort,
				barCh:        barCh,
				limiter:      limiter,
				verifyMd5:    c.verifyMd5,
				crr:          c.crr,
				objectInfo:   cfc.ObjectInfo,
				payer:        c.payer,
			}

			_copyPart := copyPart
			pool.ExecuteFunc(func() interface{} {
				ret := task.Run()
				if _metadata, _copyPartError := c.handleCopyPartResult(cfc, checkpointFile, ret, lock); _copyPartError != nil {
					if atomic.CompareAndSwapInt32(&copyPartErrorFlag, 0, 1) {
						copyPartError.Store(_copyPartError)
					}
				} else if _metadata != nil {
					if atomic.CompareAndSwapInt32(&metadataFlag, 0, 1) {
						metadata.Store(_metadata)
					}
				}

				if !c.crr {
					if config["showBytesForCopy"] == c_true {
						completed := _copyPart.RangeEnd - _copyPart.RangeStart + 1
						barCh.Send64(completed)
						progress.AddEffectiveStream(completed)
						progress.AddFinishedStream(completed)
					} else {
						barCh.Send64(1)
					}

					if barChFlag {
						progress.AddFinishedCount(1)
					}
				}

				return nil
			})
		} else if c.crr || config["showBytesForCopy"] == c_true {
			completed := copyPart.RangeEnd - copyPart.RangeStart + 1
			barCh.Send64(completed)
			progress.AddFinishedStream(completed)
		} else {
			barCh.Send64(1)
			if barChFlag {
				progress.AddFinishedCount(1)
			}
		}
	}
	if barChFlag {
		barCh.Start()
	}
	pool.ShutDown()

	var e error
	if _e, ok := copyPartError.Load().(error); ok {
		e = _e
	}

	var m map[string]string
	if _m, ok := metadata.Load().(map[string]string); ok {
		m = _m
	}

	return abort, m, e
}

func (c *transferCommand) copyBigObject(srcBucket, srcKey, versionId, dstBucket, dstKey string, srcMetaContext *MetaContext,
	metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, barCh progress.SingleBarChan) (int, string, error) {

	if srcMetaContext.Size == 0 {
		return c.copySmallObject(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh)
	}

	checkpointFile := c.getCheckpointFile(dstBucket, dstKey, versionId, cm)
	cfc := &CopyObjectCheckpoint{}
	stat, err := os.Stat(checkpointFile)
	needPrepare := true
	if err == nil {
		if stat.IsDir() {
			return 0, "", fmt.Errorf("Checkpoint file for copying [%s]-[%s] is a folder", dstBucket, dstKey)
		}
		err = c.loadCheckpoint(checkpointFile, cfc)
		if err != nil {
			if err = os.Remove(checkpointFile); err != nil {
				return 0, "", err
			}
		} else if !cfc.isValid(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext) {
			if cfc.DestinationBucket != "" && cfc.DestinationKey != "" && cfc.UploadId != "" {
				if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
					return 0, "", _err
				}
			}
			if err = os.Remove(checkpointFile); err != nil {
				return 0, "", err
			}
		} else {
			needPrepare = false
		}
	}

	if needPrepare {
		err = c.prepareCopyObjectCheckpoint(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, cfc)
		if err != nil {
			return 0, "", err
		}
		err = c.recordCheckpointFile(checkpointFile, cfc)
		if err != nil {
			if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
				return 0, "", _err
			}
			return 0, "", err
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
		if config["showBytesForCopy"] == c_true {
			barCh.SetBytes(true)
			barCh.SetTemplate(progress.SpeedOnly)
			barCh.SetTotalCount(cfc.ObjectInfo.Size)
			progress.SetTotalStream(cfc.ObjectInfo.Size)
		} else {
			totalCount := int64(len(cfc.CopyParts))
			barCh.SetTotalCount(totalCount)
			barCh.SetTemplate(progress.Simple)
			progress.SetTotalCount(totalCount)
		}
		barChFlag = true
	}

	abort, _, copyObjectError := c.copyBigObjectConcurrent(cfc, checkpointFile, barChFlag, barCh, nil)
	if barChFlag {
		barCh.WaitToFinished()
	}
	if abort == 1 {
		if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
			return 0, "", _err
		}
		if err = os.Remove(checkpointFile); err != nil {
			return 0, "", err
		}
	}
	if copyObjectError != nil {
		return 0, "", copyObjectError
	}

	if barChFlag {
		h := &assist.HintV2{}
		h.Message = "Waiting for the copied key to be completed on server side"
		h.Start()
		defer h.End()
	}

	status, requestId, completeErr := c.completeMultipartUploadForCopyObject(cfc)
	if completeErr != nil {
		if obsError, ok := completeErr.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			if isContinue, _err := c.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId); !isContinue {
				return 0, "", _err
			}
			if _err := os.Remove(checkpointFile); _err != nil {
				return 0, "", _err
			}
		}
		return 0, "", completeErr
	}

	if _err := os.Remove(checkpointFile); _err != nil {
		doLog(LEVEL_WARN, "Copy key [%s] in the bucket [%s] to key [%s] in the bucket [%s] successfully, but remove checkpoint file [%s] failed",
			cfc.SourceKey, cfc.SourceBucket, cfc.DestinationKey, cfc.DestinationBucket, checkpointFile)
	}
	return status, requestId, nil

}

func (c *transferCommand) copyObjectWithMetaContext(srcBucket, srcKey, versionId string, srcMetaContext *MetaContext, srcMetaErr error,
	dstBucket, dstKey string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, batchFlag int, count int64, fastFailed error) int {
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
				if config["showBytesForCopy"] == c_true {
					if barCh != nil {
						if srcMetaContext.Size <= 0 {
							barCh.Send64(1)
						} else {
							barCh.Send64(srcMetaContext.Size)
						}
					}
					progress.AddFinishedStream(srcMetaContext.Size)
				} else if barCh != nil {
					barCh.Send64(count)
				}

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
					srcBucket, srcKey, _versionId, dstBucket, dstKey)
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
			if config["showBytesForCopy"] == c_true {
				if barCh != nil {
					if srcMetaContext.Size <= 0 {
						barCh.Send64(1)
					} else {
						barCh.Send64(srcMetaContext.Size)
					}
				}
				progress.AddFinishedStream(srcMetaContext.Size)
			} else if barCh != nil {
				barCh.Send64(count)
			}

			if batchFlag >= 1 {
				c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, n/a, n/a, success message [dry run done], n/a", srcObjectSizeStr, srcBucket, srcKey,
					_versionId, dstBucket, dstKey)
			}
			if batchFlag != 2 {
				printf("\nCopy dry run successfully, %s, obs://%s/%s%s --> obs://%s/%s", srcObjectSizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey)
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
	var status int
	start := assist.GetUtcNow()
	if copyObjectError == nil {
		if srcMetaContext.Size >= c.bigfileThreshold || srcMetaContext.Size >= serverBigFileThreshold {
			status, requestId, copyObjectError = c.copyBigObject(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh)
		} else {
			status, requestId, copyObjectError = c.copySmallObject(srcBucket, srcKey, versionId, dstBucket, dstKey, srcMetaContext, metadata, aclType, storageClassType, barCh)
		}
	}

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

	if batchFlag >= 1 {
		if copyObjectError == nil {
			c.succeedLogger.doRecord("%s, n/a, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], success message [copy succeed], request id [%s]", srcObjectSizeStr, srcBucket, srcKey,
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
			sizeStr := normalizeBytes(srcMetaContext.Size)
			printf("\nCopy successfully, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]", sizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, cost, status, requestId)
			doLog(LEVEL_DEBUG, "Copy successfully, %s, obs://%s/%s%s --> obs://%s/%s, cost [%d], status [%d], request id [%s]",
				sizeStr, srcBucket, srcKey, _versionId, dstBucket, dstKey, cost, status, requestId)
		} else {
			logError(copyObjectError, LEVEL_INFO, fmt.Sprintf("\nCopy failed, obs://%s/%s%s --> obs://%s/%s, cost [%d]", srcBucket, srcKey, _versionId, dstBucket, dstKey, cost))
		}
	}
	if copyObjectError == nil {
		return 1
	}
	return 0
}

func (c *transferCommand) submitCopyTask(srcBucket, srcDir, dstBucket, dstDir, relativePrefix string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, pool concurrent.Pool) (totalCount int64, totalBytesForProgress int64, totalObjects int64, hasListError error) {
	defer func() {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
	}()
	input := &obs.ListObjectsInput{}
	input.Bucket = srcBucket
	input.Prefix = srcDir
	input.MaxKeys = defaultListMaxKeys
	input.RequestPayer = c.payer
	var client *obs.ObsClient
	if c.crr {
		client = obsClientCrr
	} else {
		client = obsClient
	}
	for {
		if atomic.LoadInt32(&c.abort) == 1 {
			return
		}
		start := assist.GetUtcNow()
		output, err := client.ListObjects(input)
		if err != nil {
			hasListError = err
			break
		} else {
			cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
			doLog(LEVEL_INFO, "List objects in the bucket [%s] to copy successfully, cost [%d], request id [%s]", srcBucket, cost, output.RequestId)
		}
		for _, content := range output.Contents {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			srcKey := content.Key
			if !isObsFolder(srcKey) || c.matchFolder {
				if c.matchExclude(srcKey) {
					continue
				}

				if !c.matchInclude(srcKey) {
					continue
				}

				if !c.matchLastModifiedTime(content.LastModified) {
					continue
				}
			}

			_srcKey := srcKey
			if relativePrefix != "" {
				if index := strings.Index(_srcKey, relativePrefix); index >= 0 {
					_srcKey = _srcKey[len(relativePrefix):]
				}
			}

			_dstDir := dstDir

			if isObsFolder(_dstDir) {
				_dstDir = _dstDir[:len(_dstDir)-1]
			}

			if strings.HasPrefix(_srcKey, "/") {
				_srcKey = _srcKey[1:]
			}

			dstKey := _dstDir + "/" + _srcKey

			if strings.HasPrefix(dstKey, "/") {
				dstKey = dstKey[1:]
			}

			var fastFailed error
			if checkEmptyFolder("", dstKey, cm) {
				fastFailed = fmt.Errorf("Cannot copy to the specified key [%s] in the bucket [%s]", dstKey, dstBucket)
			} else if checkEmptyFolder(srcBucket, srcKey, cm) {
				fastFailed = fmt.Errorf("Cannot copy the specified key [%s] in the bucket [%s]", srcKey, srcBucket)
			}

			if dstKey == "" {
				continue
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want copy key [%s] in the bucket [%s] to key [%s] in the bucket [%s] ? Please input (y/n) to confirm:",
				srcKey, srcBucket, dstKey, dstBucket)) {
				continue
			}

			srcMetaContext := &MetaContext{
				Size:         content.Size,
				ETag:         content.ETag,
				LastModified: content.LastModified,
			}
			count := c.caculateCount(content.Size, false)
			if c.crr || config["showBytesForCopy"] == c_true {
				atomic.AddInt64(&totalCount, srcMetaContext.Size)
			} else {
				atomic.AddInt64(&totalCount, count)
			}
			if srcMetaContext.Size == 0 {
				atomic.AddInt64(&totalBytesForProgress, 1)
			} else {
				atomic.AddInt64(&totalBytesForProgress, srcMetaContext.Size)
			}
			atomic.AddInt64(&totalObjects, 1)

			if c.crr {
				pool.ExecuteFunc(func() interface{} {
					return c.handleExecResult(c.copyObjectCrrWithMetaContext(srcBucket, srcKey, "", srcMetaContext, nil, dstBucket, dstKey,
						metadata, aclType, storageClassType, barCh, limiter, 2, fastFailed), srcMetaContext.Size)
				})
			} else if config["showBytesForCopy"] == c_true {
				pool.ExecuteFunc(func() interface{} {
					return c.handleExecResult(c.copyObjectWithMetaContext(srcBucket, srcKey, "", srcMetaContext, nil, dstBucket, dstKey,
						metadata, aclType, storageClassType, barCh, 2, count, fastFailed), srcMetaContext.Size)
				})
			} else {
				pool.ExecuteFunc(func() interface{} {
					return c.handleExecResult(c.copyObjectWithMetaContext(srcBucket, srcKey, "", srcMetaContext, nil, dstBucket, dstKey,
						metadata, aclType, storageClassType, barCh, 2, count, fastFailed), 0)
				})
			}
		}

		if !output.IsTruncated {
			doLog(LEVEL_INFO, "List objects to copy finished, bucket [%s], prefix [%s], marker [%s]", srcBucket, input.Prefix, input.Marker)
			break
		}
		input.Marker = output.NextMarker
	}
	return
}

func (c *transferCommand) recordStartFuncForCopy() time.Time {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "object size", "md5 value", "src --> dst", "cost(ms)", "status code", "success message", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "object size", "src --> dst", "cost(ms)", "status code", "error code", "error message", "request id")
	c.warningLogger.doRecord("[%s, %s, %s]", "object size", "src --> dst", "warn message")
	return start
}

func (c *transferCommand) copyDir(srcBucket, srcDir, dstBucket, dstDir string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType) error {
	start := c.recordStartFuncForCopy()
	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	var limiter *ratelimit.RateLimiter
	barCh := newSingleBarChan()
	if c.crr {
		barCh.SetBytes(true)
		barCh.SetTemplate(progress.TpsAndSpeed)
		limiter = c.createRateLimiter()
	} else if config["showBytesForCopy"] == c_true {
		barCh.SetBytes(true)
		barCh.SetTemplate(progress.TpsAndSpeed2)
	} else {
		barCh.SetTemplate(progress.TpsOnly)
	}
	if c.force {
		barCh.Start()
	}

	var relativePrefix string
	if c.flat {
		if srcDir != "" && !isObsFolder(srcDir) {
			srcDir += "/"
		}
		relativePrefix = srcDir
	} else {
		relativePrefix = srcDir

		if isObsFolder(relativePrefix) {
			relativePrefix = relativePrefix[:len(relativePrefix)-1]
		}
		if index := strings.LastIndex(relativePrefix, "/"); index >= 0 {
			relativePrefix = relativePrefix[:index+1]
		} else {
			relativePrefix = ""
		}
	}

	totalCount, totalBytesForProgress, totalObjects, hasListError := c.submitCopyTask(srcBucket, srcDir, dstBucket, dstDir, relativePrefix,
		metadata, aclType, storageClassType, barCh, limiter, pool)

	doLog(LEVEL_INFO, "Number of objects to copy [%d]", totalObjects)
	progress.SetTotalCount(totalObjects)
	if c.crr || config["showBytesForCopy"] == c_true {
		progress.SetTotalStream(totalCount)
		barCh.SetTotalCount(totalBytesForProgress)
	} else {
		barCh.SetTotalCount(totalCount)
		progress.SetTotalStream(-1)
	}

	if !c.force {
		barCh.Start()
	}

	pool.ShutDown()
	barCh.WaitToFinished()
	c.recordEndWithMetricsV2(start, totalObjects, progress.GetSucceedStream(), progress.GetTotalStream())
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects from bucket [%s] to copy failed", srcBucket))
		return assist.ErrUncompeleted
	}

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}
	return nil
}

func (c *transferCommand) copyObject(srcBucket, srcKey, versionId, dstBucket, dstKey string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, batchFlag int, fastFailed error) int {
	srcMetaContext, srcMetaErr := getObjectMetadata(srcBucket, srcKey, versionId, c.payer)
	var count int64 = 1
	if srcMetaErr == nil {
		count = c.caculateCount(srcMetaContext.Size, false)
	}
	return c.copyObjectWithMetaContext(srcBucket, srcKey, versionId, srcMetaContext, srcMetaErr, dstBucket, dstKey, metadata, aclType, storageClassType, barCh, batchFlag, count, fastFailed)
}
