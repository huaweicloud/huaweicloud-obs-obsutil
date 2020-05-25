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
	"bytes"
	"concurrent"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"obs"
	"os"
	"path/filepath"
	"progress"
	"ratelimit"
	"repeatable"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type FileInfo struct {
	XMLName      xml.Name `xml:"FileInfo"`
	LastModified int64    `xml:"LastModified"`
	Size         int64    `xml:"Size"`
}

type UploadPart struct {
	PartEtag
	XMLName     xml.Name `xml:"UploadPart"`
	Offset      int64    `xml:"Offset"`
	PartSize    int64    `xml:"PartSize"`
	IsCompleted bool     `xml:"IsCompleted"`
}

type UploadFileCheckpoint struct {
	XMLName     xml.Name     `xml:"UploadFileCheckpoint"`
	Bucket      string       `xml:"Bucket"`
	Key         string       `xml:"Key"`
	FileUrl     string       `xml:"FileUrl"`
	UploadId    string       `xml:"UploadId,omitempty"`
	FileInfo    FileInfo     `xml:"FileInfo"`
	UploadParts []UploadPart `xml:"UploadParts>UploadPart"`
}

func (ufc *UploadFileCheckpoint) isResume() bool {
	for _, uploadPart := range ufc.UploadParts {
		if uploadPart.IsCompleted {
			return true
		}
	}
	return false
}

func (ufc *UploadFileCheckpoint) isValid(bucket, key, fileUrl string, fileStat os.FileInfo) bool {
	if ufc.Bucket != bucket || ufc.Key != key || ufc.FileUrl != fileUrl {
		return false
	}

	if ufc.FileInfo.Size != fileStat.Size() || ufc.FileInfo.LastModified != fileStat.ModTime().Unix() {
		return false
	}

	if ufc.UploadId == "" {
		return false
	}

	return true
}

type uploadPartTask struct {
	bucket     string
	key        string
	fileUrl    string
	uploadId   string
	partNumber int
	offset     int64
	partSize   int64
	verifyMd5  bool
	abort      *int32
	writer     io.Writer
	barCh      progress.SingleBarChan
	limiter    *ratelimit.RateLimiter
	fileInfo   FileInfo
	payer      string
}

func (t *uploadPartTask) Run() interface{} {
	if atomic.LoadInt32(t.abort) == 1 {
		return errAbort
	}

	fd, err := os.Open(t.fileUrl)
	if err != nil {
		return err
	}
	defer fd.Close()

	input := &obs.UploadPartInput{}
	input.Bucket = t.bucket
	input.Key = t.key
	input.UploadId = t.uploadId

	if _, err = fd.Seek(t.offset, 0); err != nil {
		return err
	}

	_readBufferIoSize, transErr := assist.TranslateToInt64(config["readBufferIoSize"])
	if transErr != nil {
		_readBufferIoSize = defaultReadBufferIoSize
	}
	if _readBufferIoSize < minReadBufferIoSize {
		_readBufferIoSize = minReadBufferIoSize
	}
	if t.partSize < _readBufferIoSize {
		_readBufferIoSize = t.partSize
	}
	var _body *progress.SingleProgressReaderV2
	if t.writer != nil {
		_body = progress.NewSingleProgressReaderV2(assist.Wrap(bufio.NewReaderSize(fd, int(_readBufferIoSize)), t.writer), t.partSize, t.verifyMd5, t.barCh)
	} else {
		_body = progress.NewSingleProgressReaderV2(repeatable.NewReaderSize(fd, int(_readBufferIoSize), t.offset, config["faultTolerantMode"] == "true"), t.partSize, t.verifyMd5, t.barCh)
	}
	var body io.Reader = _body
	if t.limiter != nil {
		body = ratelimit.NewRateLimitReaderWithLimiterV2(body, t.limiter)
	}
	input.Body = body
	input.PartSize = t.partSize
	input.PartNumber = t.partNumber
	input.RequestPayer = t.payer
	output, err := obsClient.UploadPart(input)

	if err == nil {
		if changedErr := checkSourceChangedForUpload(t.fileUrl, t.fileInfo.Size, t.fileInfo.LastModified, t.abort); changedErr != nil {
			return &errorWrapper{
				err:       changedErr,
				requestId: output.RequestId,
			}
		}

		if t.verifyMd5 {
			localETag := _body.HexMd5()
			remoteETag := output.ETag
			if !compareETag(localETag, remoteETag) {
				return &errorWrapper{
					err:       &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after uploading part [%d] of file [%s], local md5 [%s] remote md5 [%s]", t.partNumber, t.fileUrl, localETag, remoteETag)},
					requestId: output.RequestId,
				}
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

	doLogError(err, LEVEL_ERROR, fmt.Sprintf("Bucket [%s], Key [%s], UploadId [%s], PartNumber [%d]", t.bucket, t.key, t.uploadId, t.partNumber))

	return err
}

func (c *transferCommand) uploadFolder(bucket, key, arcDir, folderUrl string, folderStat os.FileInfo, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, batchFlag int, fastFailed error) int {

	fileSizeStr := normalizeBytes(0)

	if fastFailed != nil {
		c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", fileSizeStr, folderUrl, bucket, key, fastFailed.Error())
		return 0
	}

	if batchFlag == 2 && atomic.LoadInt32(&c.abort) == 1 {
		c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, error code [%s], error message [%s], n/a", fileSizeStr, folderUrl, bucket, key,
			"AbortError", "Task is aborted")
		return 0
	}

	if !isObsFolder(key) {
		key = key + "/"
	}

	if c.update {
		if changed, err := c.ensureKeyForFolder(bucket, key); !changed {
			if err == nil {
				if barCh != nil {
					barCh.Send64(1)
				}
				if batchFlag >= 1 {
					c.succeedLogger.doRecord("%s, n/a, %s --> obs://%s/%s, n/a, n/a, success message [skip since the source is not changed], n/a", fileSizeStr, folderUrl, bucket, key)
				}
				if batchFlag != 2 {
					printf("%s --> obs://%s/%s, skip since the source is not changed", folderUrl, bucket, key)
				}
				return 2
			}
			if batchFlag >= 1 {
				c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, n/a, error message [skip since the status of source is unknown], n/a", fileSizeStr, folderUrl, bucket, key)
			}
			if batchFlag != 2 {
				printf("%s --> obs://%s/%s, skip since the status of source is unknown", folderUrl, bucket, key)
			}
			return 0
		}
	}

	if c.dryRun {
		if barCh != nil {
			barCh.Send64(1)
		}
		if batchFlag >= 1 {
			c.succeedLogger.doRecord("%s, n/a, %s --> obs://%s/%s, n/a, n/a, success message [dry run done], n/a", fileSizeStr, folderUrl, bucket, key)
		}
		if batchFlag != 2 {
			printf("%s --> obs://%s/%s, dry run done", folderUrl, bucket, key)
		}
		return 1
	}

	input := &obs.PutObjectInput{}
	input.Bucket = bucket
	input.Key = key
	input.ACL = aclType
	input.ContentLength = 0
	input.StorageClass = storageClassType
	input.RequestPayer = c.payer
	start := assist.GetUtcNow()
	output, err := obsClient.PutObject(input)

	if err == nil && arcDir != "" {
		// archive foler
		doLog(LEVEL_INFO, "Start archive folder [%s]", folderUrl)
		c.archiveFolderCommit(bucket, key, arcDir, folderUrl, fileSizeStr)
	}

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
	if batchFlag >= 1 {
		if err == nil {
			c.succeedLogger.doRecord("%s, n/a, %s --> obs://%s/%s, cost [%d], status [%d], success message [succeed], request id [%s]", fileSizeStr, folderUrl, bucket, key, cost, output.StatusCode, output.RequestId)
		} else {
			status, code, message, requestId := c.checkAbort(err, 401, 405, 403)
			c.failedLogger.doRecord("%s, %s --> obs://%s/%s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", fileSizeStr, folderUrl, bucket, key,
				cost, status, code, message, requestId)
		}
	}

	if batchFlag == 2 {
		c.ensureMaxCostAndMinCost(cost)
		atomic.AddInt64(&c.totalCost, cost)
	} else {
		if err == nil {
			printf("Upload successfully, %s --> obs://%s/%s, cost [%d]", folderUrl, bucket, key, cost)
			doLog(LEVEL_DEBUG, "Upload successfully, %s --> obs://%s/%s, cost [%d], request id [%s]", folderUrl, bucket, key, cost, output.RequestId)
		} else {
			logError(err, LEVEL_INFO, fmt.Sprintf("Upload failed, %s --> obs://%s/%s, cost [%d]", folderUrl, bucket, key, cost))
		}
	}

	if err == nil {
		if barCh != nil {
			barCh.Send(1)
		}
		return 1
	}
	return 0
}

func checkSourceChangedForUpload(fileUrl string, originSize, originLastModified int64, abort *int32) error {
	if config["checkSourceChange"] == c_true {
		if stat, err := os.Lstat(fileUrl); err != nil {
			if abort != nil {
				atomic.CompareAndSwapInt32(abort, 0, 1)
			}
			return fmt.Errorf("Source file [%s] doesnot exist", fileUrl)
		} else if originSize != stat.Size() || originLastModified != stat.ModTime().Unix() {
			if abort != nil {
				atomic.CompareAndSwapInt32(abort, 0, 1)
			}
			return fmt.Errorf("Source file [%s] changed", fileUrl)
		}
	}
	return nil
}

func (c *transferCommand) prepareUploadFileCheckpoint(bucket, key, fileUrl string,
	fileStat os.FileInfo, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType, ufc *UploadFileCheckpoint) error {

	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = bucket
	input.Key = key
	input.Metadata = metadata
	input.ACL = aclType
	input.StorageClass = storageClassType
	input.RequestPayer = c.payer
	output, err := obsClient.InitiateMultipartUpload(input)
	if err != nil {
		return err
	}
	doLog(LEVEL_DEBUG, "Initiate multipart upload [%s] in the bucket [%s] successfully, request id [%s]", key, bucket, output.RequestId)

	ufc.Bucket = bucket
	ufc.Key = key
	ufc.FileUrl = fileUrl
	ufc.UploadId = output.UploadId
	ufc.FileInfo = FileInfo{
		Size:         fileStat.Size(),
		LastModified: fileStat.ModTime().Unix(),
	}
	partSize := c.autoSelectPartSize(ufc.FileInfo.Size, um)
	count := ufc.FileInfo.Size / partSize
	if count >= 10000 {
		partSize = ufc.FileInfo.Size / 10000
		if ufc.FileInfo.Size%10000 != 0 {
			partSize++
		}
		count = ufc.FileInfo.Size / partSize
	}

	if ufc.FileInfo.Size%partSize != 0 {
		count++
	}

	if partSize > serverBigFileThreshold {
		return fmt.Errorf("The source file [%s] is too large", fileUrl)
	}

	if count == 0 {
		uploadPart := UploadPart{}
		uploadPart.PartNumber = 1
		ufc.UploadParts = []UploadPart{uploadPart}
	} else {
		uploadParts := make([]UploadPart, 0, count)
		var i int64
		for i = 0; i < count; i++ {
			uploadPart := UploadPart{
				Offset:   i * partSize,
				PartSize: partSize,
			}
			uploadPart.PartNumber = int(i) + 1
			uploadParts = append(uploadParts, uploadPart)
		}
		if lastPartSize := ufc.FileInfo.Size % partSize; lastPartSize != 0 {
			uploadParts[count-1].PartSize = lastPartSize
		}
		ufc.UploadParts = uploadParts
	}
	return nil
}

func (c *transferCommand) handleUploadPartResult(ufc *UploadFileCheckpoint, checkpointFile string, result interface{}, lock *sync.Mutex) (uploadPartError error) {
	if partETag, ok := result.(PartEtag); ok {
		lock.Lock()
		defer lock.Unlock()
		ufc.UploadParts[partETag.PartNumber-1].IsCompleted = true
		ufc.UploadParts[partETag.PartNumber-1].ETag = partETag.ETag
		uploadPartError = c.recordCheckpointFile(checkpointFile, ufc)
	} else if result != errAbort {
		if resultError, ok := result.(error); ok {
			uploadPartError = resultError
		}
	}
	return
}

func (c *transferCommand) uploadBigFileConcurrent(ufc *UploadFileCheckpoint, checkpointFile string, barChFlag bool,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (int32, error) {
	pool := concurrent.NewRoutinePool(c.parallel, defaultParallelsCacheCount)
	var uploadPartError atomic.Value
	var uploadPartErrorFlag int32

	var abort int32
	lock := new(sync.Mutex)

	for _, uploadPart := range ufc.UploadParts {
		if atomic.LoadInt32(&abort) == 1 {
			break
		}
		if !uploadPart.IsCompleted {
			task := &uploadPartTask{
				bucket:     ufc.Bucket,
				key:        ufc.Key,
				uploadId:   ufc.UploadId,
				fileUrl:    ufc.FileUrl,
				partNumber: uploadPart.PartNumber,
				offset:     uploadPart.Offset,
				partSize:   uploadPart.PartSize,
				abort:      &abort,
				barCh:      barCh,
				limiter:    limiter,
				verifyMd5:  c.verifyMd5,
				fileInfo:   ufc.FileInfo,
				payer:      c.payer,
			}

			pool.ExecuteFunc(func() interface{} {
				ret := task.Run()
				if _uploadPartError := c.handleUploadPartResult(ufc, checkpointFile, ret, lock); _uploadPartError != nil {
					if atomic.CompareAndSwapInt32(&uploadPartErrorFlag, 0, 1) {
						uploadPartError.Store(_uploadPartError)
					}
				}
				return nil
			})

		} else {
			barCh.Send64(uploadPart.PartSize)
			progress.AddFinishedStream(uploadPart.PartSize)
		}
	}
	if barChFlag {
		barCh.Start()
	}
	pool.ShutDown()

	if e, ok := uploadPartError.Load().(error); ok {
		return abort, e
	}

	return abort, nil
}

func (c *transferCommand) uploadBigFileSerial(ufc *UploadFileCheckpoint, checkpointFile string, barChFlag bool,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, writer io.Writer) (int32, error) {
	var uploadPartError error
	var abort int32

	if barChFlag {
		barCh.Start()
	}
	for _, uploadPart := range ufc.UploadParts {
		if atomic.LoadInt32(&abort) == 1 {
			break
		}
		_writer := writer
		if uploadPartError != nil {
			_writer = nil
		}
		task := &uploadPartTask{
			bucket:     ufc.Bucket,
			key:        ufc.Key,
			uploadId:   ufc.UploadId,
			fileUrl:    ufc.FileUrl,
			partNumber: uploadPart.PartNumber,
			offset:     uploadPart.Offset,
			partSize:   uploadPart.PartSize,
			abort:      &abort,
			barCh:      barCh,
			limiter:    limiter,
			verifyMd5:  c.verifyMd5,
			writer:     _writer,
			payer:      c.payer,
		}
		result := task.Run()
		if partETag, ok := result.(PartEtag); ok {
			ufc.UploadParts[partETag.PartNumber-1].IsCompleted = true
			ufc.UploadParts[partETag.PartNumber-1].ETag = partETag.ETag
			if uploadPartError == nil {
				uploadPartError = c.recordCheckpointFile(checkpointFile, ufc)
			}
		} else if uploadPartError == nil && result != errAbort {
			if resultError, ok := result.(error); ok {
				uploadPartError = resultError
			}
		}
	}

	return abort, uploadPartError
}

func (c *transferCommand) completeMultipartUploadForUploadFile(ufc *UploadFileCheckpoint) (string, int, error) {
	input := &obs.CompleteMultipartUploadInput{}
	input.Bucket = ufc.Bucket
	input.Key = ufc.Key
	input.UploadId = ufc.UploadId
	parts := make([]obs.Part, 0, len(ufc.UploadParts))
	for _, uploadPart := range ufc.UploadParts {
		part := obs.Part{
			ETag:       uploadPart.ETag,
			PartNumber: uploadPart.PartNumber,
		}
		parts = append(parts, part)
	}
	input.Parts = parts
	input.RequestPayer = c.payer
	output, err := obsClient.CompleteMultipartUpload(input)
	if err == nil {
		doLog(LEVEL_DEBUG, "Complete multipart upload [%s] in the bucket [%s] successfully, request id [%s]", ufc.Key, ufc.Bucket, output.RequestId)
		return output.RequestId, output.StatusCode, nil
	}
	return "", 0, err
}

func (c *transferCommand) uploadBigFile(bucket, key, fileUrl string, fileStat os.FileInfo, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int) (requestId string, status int, md5Value string, uploadFileError error) {

	if fileStat.Size() == 0 {
		return c.uploadSmallFile(bucket, key, fileUrl, fileStat, metadata, aclType, storageClassType, barCh, limiter)
	}

	fileSizeStr := normalizeBytes(fileStat.Size())
	checkpointFile := c.getCheckpointFile(bucket, key, "", um)
	ufc := &UploadFileCheckpoint{}
	stat, err := os.Stat(checkpointFile)
	needPrepare := true
	if err == nil {
		if stat.IsDir() {
			uploadFileError = fmt.Errorf("Checkpoint file for uploading [%s]-[%s] is a folder", bucket, key)
			return
		}
		err = c.loadCheckpoint(checkpointFile, ufc)
		if err != nil {
			if err = os.Remove(checkpointFile); err != nil {
				uploadFileError = err
				return
			}
		} else if !ufc.isValid(bucket, key, fileUrl, fileStat) {
			if ufc.Bucket != "" && ufc.Key != "" && ufc.UploadId != "" {
				if isContinue, _err := c.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId); !isContinue {
					uploadFileError = _err
					return
				}
			}
			if err = os.Remove(checkpointFile); err != nil {
				uploadFileError = err
				return
			}
		} else {
			needPrepare = false
		}
	}
	if needPrepare {
		err = c.prepareUploadFileCheckpoint(bucket, key, fileUrl, fileStat, metadata, aclType, storageClassType, ufc)
		if err != nil {
			uploadFileError = err
			return
		}
		err = c.recordCheckpointFile(checkpointFile, ufc)
		if err != nil {
			if isContinue, _err := c.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId); !isContinue {
				uploadFileError = _err
				return
			}
			uploadFileError = err
			return
		}
	}

	defer func() {
		if r := recover(); r != nil {
			c.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId)
			panic(r)
		}
	}()

	barChFlag := false
	if barCh == nil {
		barCh = newSingleBarChan()
		barCh.SetTemplate(progress.SpeedOnly)
		barCh.SetBytes(true)
		barCh.SetTotalCount(ufc.FileInfo.Size)
		progress.SetTotalStream(ufc.FileInfo.Size)
		barChFlag = true
	}

	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	var abort int32
	var md5Writer io.Writer
	if c.verifyMd5 {
		if batchFlag <= 1 || ufc.FileInfo.Size < serialVerifyMd5Threshold || ufc.isResume() {
			abort, uploadFileError = c.uploadBigFileConcurrent(ufc, checkpointFile, barChFlag, barCh, limiter)
		} else {
			md5Writer = assist.GetMd5Writer()
			abort, uploadFileError = c.uploadBigFileSerial(ufc, checkpointFile, barChFlag, barCh, limiter, md5Writer)
		}
	} else {
		abort, uploadFileError = c.uploadBigFileConcurrent(ufc, checkpointFile, barChFlag, barCh, limiter)
	}

	if barChFlag {
		barCh.WaitToFinished()
	}

	if abort == 1 {
		if isContinue, _err := c.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId); !isContinue {
			uploadFileError = _err
			return
		}
		if err = os.Remove(checkpointFile); err != nil {
			uploadFileError = err
			return
		}
	}
	if uploadFileError != nil {
		return
	}

	if c.verifyMd5 {
		if barChFlag {
			h := &assist.HintV2{}
			h.Message = c_waiting_caculate_md5
			h.Start()
			if md5Writer != nil {
				md5Value = assist.GetHexMd5(md5Writer)
			} else {
				if etag, err := md5File(fileUrl); err == nil {
					md5Value = assist.Hex(etag)
				} else {
					warnMessage := fmt.Sprintf("Get local md5 failed after uploading file [%s] - %s", fileUrl, err.Error())
					warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, fileUrl, bucket, key, warnMessage)
					c.recordWarnMessage(warnMessage, warnLoggerMessage)
				}
			}
			h.End()
		} else {
			if md5Writer != nil {
				md5Value = assist.GetHexMd5(md5Writer)
			} else {
				if etag, err := md5File(fileUrl); err == nil {
					md5Value = assist.Hex(etag)
				} else {
					warnMessage := fmt.Sprintf("Get local md5 failed after uploading file [%s] - %s", fileUrl, err.Error())
					warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, fileUrl, bucket, key, warnMessage)
					c.recordWarnMessage(warnMessage, warnLoggerMessage)
				}
			}
		}
	}

	if barChFlag {
		h := &assist.HintV2{}
		h.Message = "Waiting for the uploaded key to be completed on server side"
		h.Start()
		defer h.End()
	}

	if _requestId, _status, completeErr := c.completeMultipartUploadForUploadFile(ufc); completeErr != nil {
		if obsError, ok := completeErr.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			if isContinue, _err := c.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId); !isContinue {
				uploadFileError = _err
				return
			}
			if removeErr := os.Remove(checkpointFile); removeErr != nil {
				uploadFileError = removeErr
				return
			}
		}
		uploadFileError = completeErr
	} else {
		if removeErr := os.Remove(checkpointFile); removeErr != nil {
			doLog(LEVEL_WARN, "Upload big file [%s] as key [%s] in the bucket [%s] successfully, but remove checkpoint file [%s] failed",
				ufc.FileUrl, ufc.Key, ufc.Bucket, checkpointFile)
		}
		requestId = _requestId
		status = _status
	}

	return
}

func (c *transferCommand) uploadSmallFile(bucket, key, fileUrl string, fileStat os.FileInfo, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (requestId string, status int, md5Value string, uploadFileError error) {

	fileSize := fileStat.Size()

	fileSizeStr := normalizeBytes(fileSize)
	var fd io.ReadSeeker
	if fileSize > 0 {
		_fd, err := os.Open(fileUrl)
		if err != nil {
			uploadFileError = err
			return
		}
		fd = _fd
		defer _fd.Close()
	} else {
		fd = bytes.NewReader([]byte{})
	}

	input := &obs.PutObjectInput{}
	input.Bucket = bucket
	input.Key = key
	input.Metadata = metadata
	input.ACL = aclType
	input.ContentLength = fileSize
	input.StorageClass = storageClassType

	barChFlag := false
	if barCh == nil && fileSize > 0 {
		barCh = newSingleBarChan()
		barCh.SetTemplate(progress.SpeedOnly)
		barCh.SetBytes(true)
		barCh.SetTotalCount(fileSize)
		progress.SetTotalStream(fileSize)
		barCh.Start()
		barChFlag = true
	}

	_readBufferIoSize, transErr := assist.TranslateToInt64(config["readBufferIoSize"])
	if transErr != nil {
		_readBufferIoSize = defaultReadBufferIoSize
	}
	if _readBufferIoSize < minReadBufferIoSize {
		_readBufferIoSize = minReadBufferIoSize
	}
	if fileSize < _readBufferIoSize {
		_readBufferIoSize = fileSize
	}

	_body := progress.NewSingleProgressReaderV2(repeatable.NewReaderSize(fd, int(_readBufferIoSize), 0, config["faultTolerantMode"] == "true"), input.ContentLength, c.verifyMd5, barCh)
	var body io.Reader = _body
	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	if limiter != nil {
		body = ratelimit.NewRateLimitReaderWithLimiterV2(body, limiter)
	}

	input.Body = body
	input.RequestPayer = c.payer
	output, err := obsClient.PutObject(input)

	if barChFlag {
		barCh.WaitToFinished()
	}
	if err != nil {
		uploadFileError = err
		return
	}

	if changedErr := checkSourceChangedForUpload(fileUrl, fileSize, fileStat.ModTime().Unix(), nil); changedErr != nil {
		uploadFileError = &errorWrapper{
			err:       changedErr,
			requestId: output.RequestId,
		}
		return
	}

	md5Value = _body.HexMd5()
	if c.verifyMd5 && !compareETag(md5Value, output.ETag) {
		doLog(LEVEL_ERROR, "Verify md5 failed after uploading file [%s], local md5 [%s] remote md5 [%s], will try to delete uploaded key", fileUrl, md5Value, output.ETag)
		if deleteRequestId, err := c.deleteObject(bucket, key, ""); err == nil {
			doLog(LEVEL_INFO, "Delete key [%s] in the bucket [%s] successfully, request id [%s]", key, bucket, deleteRequestId)
		} else {
			status, code, message, deleteRequestId := getErrorInfo(err)
			warnMessage := fmt.Sprintf("Delete key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]", key, bucket, status, code, message, deleteRequestId)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]",
				fileSizeStr, fileUrl, bucket, key, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
		uploadFileError = &errorWrapper{
			err:       &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after uploading file [%s], local md5 [%s] remote md5 [%s]", fileUrl, md5Value, output.ETag)},
			requestId: output.RequestId,
		}
		return
	}

	requestId = output.RequestId
	status = output.StatusCode
	return
}

func (c *transferCommand) beforeUploadFile(bucket, key, realArcPath, fileUrl string, fileStat os.FileInfo, startUploadModifiedTime, fileSize int64,
	fileSizeStr string, barCh progress.SingleBarChan, batchFlag int, fastFailed error) int {

	if fastFailed != nil {
		c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", fileSizeStr, fileUrl, bucket, key, fastFailed.Error())
		return 0
	}

	if batchFlag == 2 && atomic.LoadInt32(&c.abort) == 1 {
		c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, error code [%s], error message [%s], n/a", fileSizeStr, fileUrl, bucket, key,
			"AbortError", "Task is aborted")
		return 0
	}

	if c.update {
		if changed, err := c.ensureKeyForFile(bucket, key, fileStat); !changed {
			if err == nil {
				if realArcPath != "" {
					// check data is not changed by modified time. and archive file
					doLog(LEVEL_INFO, "Source is not change, start archive file [%s]", fileUrl)
					if retErr := c.archiveFileCommit(bucket, key, realArcPath, fileUrl, fileSizeStr, startUploadModifiedTime, fileSize); retErr != nil {
						if batchFlag >= 1 {
							c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", fileSizeStr, fileUrl, bucket, key, retErr.Error())
						}

						if batchFlag != 2 {
							printf("%s, %s --> obs://%s/%s, %s", fileSizeStr, fileUrl, bucket, key, retErr.Error())
						}
						return 0
					}
				}

				if barCh != nil {
					if fileSize <= 0 {
						barCh.Send64(1)
					} else {
						barCh.Send64(fileSize)
					}
				}
				progress.AddFinishedStream(fileSize)

				if batchFlag >= 1 {
					c.succeedLogger.doRecord("%s, n/a, %s --> obs://%s/%s, n/a, n/a, success message [skip since the source is not changed], n/a", fileSizeStr, fileUrl, bucket, key)
				}
				if batchFlag != 2 {
					printf("%s, %s --> obs://%s/%s, skip since the source is not changed", fileSizeStr, fileUrl, bucket, key)
				}

				return 2
			}
			if batchFlag >= 1 {
				c.failedLogger.doRecord("%s, %s --> obs://%s/%s, n/a, n/a, n/a, error message [skip since the status of source is unknown], n/a", fileSizeStr, fileUrl, bucket, key)
			}
			if batchFlag != 2 {
				printf("%s, %s --> obs://%s/%s, skip since the status of source is unknown", fileSizeStr, fileUrl, bucket, key)
			}
			return 0
		}
	}

	return -1
}

func (c *transferCommand) dryRunUploadFile(bucket, key, fileUrl string, fileSize int64, fileSizeStr string, barCh progress.SingleBarChan, batchFlag int) int {
	if c.dryRun {
		if barCh != nil {
			if fileSize <= 0 {
				barCh.Send64(1)
			} else {
				barCh.Send64(fileSize)
			}
		}
		progress.AddFinishedStream(fileSize)
		var caculateMd5Err error
		md5Value := c_na
		if c.verifyMd5 {
			if etag, err := md5File(fileUrl); err == nil {
				md5Value = assist.Hex(etag)
			} else {
				caculateMd5Err = err
			}
		}
		if batchFlag >= 1 {
			if caculateMd5Err == nil {
				c.succeedLogger.doRecord("%s, %s, %s --> obs://%s/%s, n/a, n/a, success message [dry run done], n/a", fileSizeStr, md5Value, fileUrl, bucket, key)
			} else {
				c.failedLogger.doRecord("%s, n/a, %s --> obs://%s/%s, n/a, n/a, error message [dry run done with error - %s], n/a", fileSizeStr, fileUrl, bucket, key, caculateMd5Err.Error())
			}
		}
		if batchFlag != 2 {
			if caculateMd5Err == nil {
				printf("\nUpload dry run successfully, %s, %s, %s --> obs://%s/%s", fileSizeStr, md5Value, fileUrl, bucket, key)
			} else {
				printf("\nUpload dry run failed, %s, %s --> obs://%s/%s, error [%s]", fileSizeStr, fileUrl, bucket, key, caculateMd5Err.Error())
			}
		}
		return 1
	}

	return -1
}

func (c *transferCommand) afterUploadFile(bucket, key, fileUrl string, metadata map[string]string,
	fileSize int64, fileSizeStr string, requestId, md5Value string, barCh progress.SingleBarChan) (uploadFileError error) {
	if barCh != nil && fileSize <= 0 {
		barCh.Send64(1)
	}

	if c.verifyMd5 && md5Value != "" {
		if obsVersion, ok := c.bucketsVersionMap[bucket]; ok && obsVersion == OBS_VERSION_UNKNOWN {
			warnMessage := fmt.Sprintf("Bucket [%s] cannot support setObjectMetadata interface, because of obs version check failed - so skip set object md5", bucket)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, fileUrl, bucket, key, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		} else if ok && obsVersion >= "3.0" {
			if _, _err := c.setObjectMd5(bucket, key, "", md5Value, metadata); _err != nil {
				_status, _code, _message, _requestId := getErrorInfo(_err)
				warnMessage := fmt.Sprintf("Upload file [%s] as key [%s] in the bucket [%s] successfully - but set object md5 failed - status [%d] - error code [%s] - error message [%s] - request id [%s]",
					fileUrl, key, bucket, _status, _code, _message, _requestId)
				warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]",
					fileSizeStr, fileUrl, bucket, key, warnMessage)
				c.recordWarnMessage(warnMessage, warnLoggerMessage)
			}
		} else {
			warnMessage := fmt.Sprintf("Bucket [%s] cannot support setObjectMetadata interface - so skip set object md5", bucket)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, fileUrl, bucket, key, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	} else if c.verifyLength {
		if metaContext, err := getObjectMetadata(bucket, key, "", c.payer); err == nil {
			if metaContext.Size != fileSize {
				doLog(LEVEL_ERROR, "Verify length failed after uploading file [%s], local length [%d] remote length [%d], will try to delete uploaded key", fileUrl, fileSize, metaContext.Size)
				if _requestId, _err := c.deleteObject(bucket, key, ""); _err == nil {
					doLog(LEVEL_INFO, "Delete key [%s] in the bucket [%s] successfully, request id [%s]", key, bucket, _requestId)
				} else {
					_status, _code, _message, _requestId := getErrorInfo(_err)
					warnMessage := fmt.Sprintf("Delete key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]", key, bucket, _status, _code, _message, _requestId)
					warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]",
						fileSizeStr, fileUrl, bucket, key, warnMessage)
					c.recordWarnMessage(warnMessage, warnLoggerMessage)
				}
				uploadFileError = &errorWrapper{
					err:       &verifyLengthError{msg: fmt.Sprintf("Verify length failed after uploading file [%s], local length [%d] remote length [%d]", fileUrl, fileSize, metaContext.Size)},
					requestId: requestId,
				}
			}
		} else {
			warnMessage := fmt.Sprintf("Upload file [%s] as key [%s] in the bucket [%s] successfully - but can not verify length - %s",
				fileUrl, key, bucket, err.Error())
			warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]",
				fileSizeStr, fileUrl, bucket, key, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	}

	return
}

func (c *transferCommand) uploadFile(bucket, key, realArcPath, fileUrl string, fileStat os.FileInfo, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int, fastFailed error) int {

	startUploadModifiedTime := fileStat.ModTime().Unix()
	fileSize := fileStat.Size()
	fileSizeStr := normalizeBytes(fileSize)

	if ret := c.beforeUploadFile(bucket, key, realArcPath, fileUrl, fileStat, startUploadModifiedTime, fileSize, fileSizeStr, barCh, batchFlag, fastFailed); ret >= 0 {
		return ret
	}

	if ret := c.dryRunUploadFile(bucket, key, fileUrl, fileSize, fileSizeStr, barCh, batchFlag); ret >= 0 {
		return ret
	}

	var uploadFileError error
	var requestId string
	var md5Value string
	var status int
	start := assist.GetUtcNow()

	if fileSize >= c.bigfileThreshold || fileSize >= serverBigFileThreshold {
		requestId, status, md5Value, uploadFileError = c.uploadBigFile(bucket, key, fileUrl, fileStat, metadata, aclType, storageClassType, barCh, limiter, batchFlag)
	} else {
		requestId, status, md5Value, uploadFileError = c.uploadSmallFile(bucket, key, fileUrl, fileStat, metadata, aclType, storageClassType, barCh, limiter)
	}

	if uploadFileError == nil {
		uploadFileError = c.afterUploadFile(bucket, key, fileUrl, metadata, fileSize, fileSizeStr, requestId, md5Value, barCh)
	}

	if md5Value == "" {
		md5Value = c_na
	}
	if uploadFileError == nil && realArcPath != "" {
		// check data is not changed by modified time. and archive file
		doLog(LEVEL_INFO, "Start archive file [%s]", fileUrl)
		uploadFileError = c.archiveFileCommit(bucket, key, realArcPath, fileUrl, fileSizeStr, startUploadModifiedTime, fileSize)
	}

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

	// batchFlag = {0: single file upload, 1: single file upload and record log, 2: multi file upload}
	if batchFlag >= 1 {
		if uploadFileError == nil {
			c.succeedLogger.doRecord("%s, %s, %s --> obs://%s/%s, cost [%d], status [%d], success message [upload succeed], request id [%s]", fileSizeStr, md5Value, fileUrl, bucket, key, cost, status, requestId)
		} else {
			_status, _code, _message, _requestId := c.checkAbort(uploadFileError, 401, 405, 403)
			c.failedLogger.doRecord("%s, %s --> obs://%s/%s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", fileSizeStr, fileUrl, bucket, key, cost, _status, _code, _message, _requestId)
		}
	}

	if batchFlag == 2 {
		c.ensureMaxCostAndMinCost(cost)
		atomic.AddInt64(&c.totalCost, cost)
	} else {
		if uploadFileError == nil {
			printf("\nUpload successfully, %s, %s, %s --> obs://%s/%s, cost [%d], status [%d], request id [%s]", fileSizeStr, md5Value, fileUrl, bucket, key, cost, status, requestId)
			doLog(LEVEL_DEBUG, "Upload successfully, %s, %s, %s --> obs://%s/%s, cost [%d], status [%d], request id [%s]", fileSizeStr, md5Value, fileUrl, bucket, key, cost, status, requestId)
		} else {
			logError(uploadFileError, LEVEL_INFO, fmt.Sprintf("\nUpload failed, %s --> obs://%s/%s, cost [%d]", fileUrl, bucket, key, cost))
		}
	}

	if uploadFileError == nil {
		return 1
	}
	return 0
}

func (c *transferCommand) isConsistency(info os.FileInfo, startUploadMonifiedTime, startUploadSize int64) bool {
	return info.ModTime().Unix() == startUploadMonifiedTime && info.Size() == startUploadSize
}

func (c *transferCommand) archiveFolderCommit(bucket, key, realArcPath, folderUrl, fileSizeStr string) {
	var archiveErr error
	var info os.FileInfo
	info, archiveErr = os.Lstat(folderUrl)
	if archiveErr == nil {
		if info.IsDir() {
			dstPath := assist.NormalizeFilePath(realArcPath)
			if err := assist.MkdirAll(dstPath, os.ModePerm); err == nil {
				doLog(LEVEL_INFO, "Archive folder %s --> %s succeed", folderUrl, realArcPath)
			} else {
				archiveErr = err
			}
		} else {
			archiveErr = errors.New("The source is not a folder")
		}
	}
	if archiveErr != nil {
		warnMessage := fmt.Sprintf("Archive folder %s --> %s failed due to %s", folderUrl, realArcPath, archiveErr.Error())
		warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, folderUrl, bucket, key, warnMessage)
		c.recordWarnMessage(warnMessage, warnLoggerMessage)
	}
	return
}

func (c *transferCommand) getStat(url string) (os.FileInfo, error) {
	if c.link {
		return os.Stat(url)
	}
	return os.Lstat(url)
}

func (c *transferCommand) archiveFileCommit(bucket, key, realArcPath, fileUrl,
	fileSizeStr string, startUploadMonifiedTime, startUploadSize int64) (retErr error) {

	var doDeleteObject = func() {
		if deleteRequestId, err := c.deleteObject(bucket, key, ""); err == nil {
			doLog(LEVEL_INFO, "Delete key [%s] in the bucket [%s] successfully, request id [%s]", key, bucket, deleteRequestId)
		} else {
			status, code, message, deleteRequestId := getErrorInfo(err)
			warnMessage := fmt.Sprintf("Delete key [%s] in the bucket [%s] failed - status [%d] - error code [%s] - error message [%s] - request id [%s]", key, bucket, status, code, message, deleteRequestId)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]",
				fileSizeStr, fileUrl, bucket, key, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	}

	info, archiveErr := c.getStat(fileUrl)
	if archiveErr == nil {
		if info.IsDir() {
			archiveErr = errors.New("The source is not a file")
		} else if c.isConsistency(info, startUploadMonifiedTime, startUploadSize) {
			_writeBufferIoSize, transErr := assist.TranslateToInt64(config["writeBufferIoSize"])
			if transErr != nil {
				_writeBufferIoSize = defaultWriteBufferIoSize
			}
			if cpErr := assist.CopyFile(fileUrl, realArcPath, true, int(_writeBufferIoSize), false); cpErr == nil {
				if _info, _archiveErr := c.getStat(fileUrl); _archiveErr == nil {
					if _info.IsDir() {
						archiveErr = errors.New("The source is not a file")
					} else if c.isConsistency(_info, startUploadMonifiedTime, startUploadSize) {
						if rmErr := os.Remove(fileUrl); rmErr == nil {
							doLog(LEVEL_INFO, "Archive file %s --> %s succeed", fileUrl, realArcPath)
						} else {
							archiveErr = rmErr
						}
					} else {
						if rmErr := os.Remove(realArcPath); rmErr != nil {
							doLog(LEVEL_WARN, "Delete file [%s] failed due to %s", realArcPath, rmErr.Error())
						}
						doDeleteObject()
						retErr = fmt.Errorf("Try to archive file but source file [%s] changed", fileUrl)
					}
				} else {
					archiveErr = _archiveErr
				}
			} else {
				archiveErr = cpErr
			}
		} else {
			doDeleteObject()
			retErr = fmt.Errorf("Try to archive file but source file [%s] changed", fileUrl)
		}
	}

	if archiveErr != nil {
		warnMessage := fmt.Sprintf("Archive file %s --> %s failed due to %s", fileUrl, realArcPath, archiveErr.Error())
		warnLoggerMessage := fmt.Sprintf("%s, %s --> obs://%s/%s, warn message [%s]", fileSizeStr, fileUrl, bucket, key, warnMessage)
		c.recordWarnMessage(warnMessage, warnLoggerMessage)
	}

	return
}

func (c *transferCommand) getWalkFunc(bucket, dir, arcDir, folder string, linkFolder bool, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType, pool concurrent.Pool,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter,
	totalBytes, totalBytesForProgress, totalFiles *int64) func(path string, info os.FileInfo, err error) error {
	if dir != "" && !isObsFolder(dir) {
		dir = dir + "/"
	}

	if arcDir != "" && !isObsFolder(arcDir) {
		arcDir = arcDir + "/"
	}

	var relativeFolder string
	var isRootFolder = false

	_relativeFolder := c.getRelativeFolder(folder)

	if _relativeFolder == "" && !linkFolder {
		isRootFolder = true
	}

	if !linkFolder && !c.flat {
		relativeFolder = _relativeFolder
	}

	folderLength := len(folder)
	return func(path string, info os.FileInfo, err error) error {
		arcPath := ""
		commonDirSuffix := relativeFolder
		if len(path) == folderLength || isRootFolder {
			commonDirSuffix += strings.Replace(path[folderLength:], "\\", "/", -1)
		} else if len(path) > folderLength {
			commonDirSuffix += strings.Replace(path[folderLength+1:], "\\", "/", -1)
		}
		key := dir + commonDirSuffix

		if arcDir != "" {
			arcPath = arcDir + commonDirSuffix
		}

		if err != nil {
			c.recordPrepareFailed(bucket, key, path, totalBytesForProgress, totalFiles, err.Error())
			return err
		}

		if key == "/" || key == "" {
			return nil
		}

		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if c.link {
				_path, _info, _err := assist.GetRealPath(path)
				if _err != nil {
					c.recordPrepareFailed(bucket, key, path, totalBytesForProgress, totalFiles, _err.Error())
					return nil
				}
				if _info.IsDir() {
					c.folderMapLock.Lock()
					pathToCheck := folder + "-" + assist.NormalizeFilePath(path) + "/"
					if previousFolder, ok := c.folderMap[_path]; ok && strings.Index(previousFolder, pathToCheck) >= 0 {
						panicErr := fmt.Errorf("the symbolic-link folder [%s] --> [%s] result in a circle with folder [%s], path to check is [%s]", path, _path, previousFolder, pathToCheck)
						c.recordPrepareFailed(bucket, key, path, totalBytesForProgress, totalFiles, panicErr.Error())
						c.folderMapLock.Unlock()
						if config["panicForSymbolicLinkCircle"] == c_true {
							c.recordCntV2()
							c.endLogger()
							panic(panicErr)
						}
						if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
							c.scanError.Store(panicErr)
						}
						return nil
					}

					c.folderMap[_path] = pathToCheck
					c.folderMapLock.Unlock()

					return filepath.Walk(_path, c.getWalkFunc(bucket, key, arcPath, _path, true, metadata,
						aclType, storageClassType, pool, barCh, limiter, totalBytes, totalBytesForProgress, totalFiles))
				}
				path = _path
				info = _info
			} else {
				if stat, err := os.Stat(path); err == nil && stat.IsDir() {
					info = stat
				}
			}
		}

		if info.IsDir() {

			if c.matchFolder {
				if c.matchExclude(path) {
					return nil
				}

				if !c.matchInclude(path) {
					return nil
				}
			}

			if !c.matchUploadTimeRange(info) {
				return nil
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want upload folder [%s] to bucket [%s] ? Please input (y/n) to confirm:", path, bucket)) {
				return nil
			}
			// modify by w00468571 wanghongbao, if the disableDirObject is true the dir will not upload as a object
			if c.disableDirObject {
				return nil
			}
			atomic.AddInt64(totalBytesForProgress, 1)
			atomic.AddInt64(totalFiles, 1)

			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.uploadFolder(bucket, key, arcPath, path, info, aclType, storageClassType, barCh, 2, nil), 0)
			})

		} else {
			if c.matchExclude(path) {
				return nil
			}

			if !c.matchInclude(path) {
				return nil
			}

			if !c.matchUploadTimeRange(info) {
				return nil
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want upload file [%s] to bucket [%s] ? Please input (y/n) to confirm:", path, bucket)) {
				return nil
			}
			atomic.AddInt64(totalBytes, info.Size())
			if info.Size() == 0 {
				atomic.AddInt64(totalBytesForProgress, 1)
			} else {
				atomic.AddInt64(totalBytesForProgress, info.Size())
			}
			atomic.AddInt64(totalFiles, 1)

			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.uploadFile(bucket, key, arcPath, path, info, metadata, aclType, storageClassType, barCh, limiter, 2, nil), info.Size())
			})
		}

		return nil
	}
}

func (c *transferCommand) readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return names, nil
}

func (c *transferCommand) scanFolder(path string, info os.FileInfo, doAction func(path string, info os.FileInfo, err error) error) {
	if names, err := c.readDirNames(path); err != nil {
		doAction(path, info, err)
		if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
			c.scanError.Store(err)
		}
	} else {
		doAction(path, info, nil)
		for _, name := range names {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			filename := filepath.Join(path, name)
			subFileInfo, err := os.Lstat(filename)
			if err != nil {
				doAction(filename, nil, err)
				if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
					c.scanError.Store(err)
				}
				continue
			}
			if !subFileInfo.IsDir() {
				doAction(filename, subFileInfo, nil)
				continue
			}
			c.doScan(filename, doAction)
		}
	}
}

func (c *transferCommand) doScan(folder string, doAction func(path string, info os.FileInfo, err error) error) {
	c.scanPool.ExecuteFunc(func() (ret interface{}) {
		if atomic.LoadInt32(&c.abort) == 1 {
			return
		}

		info, err := os.Lstat(folder)
		if err != nil {
			doAction(folder, nil, err)
			if atomic.CompareAndSwapInt32(&c.scanErrorFlag, 0, 1) {
				c.scanError.Store(err)
			}
		} else if !info.IsDir() {
			doAction(folder, info, nil)
		} else {
			// start scan folder
			c.scanFolder(folder, info, doAction)
			// end scan folder
		}
		return
	})
}

func (c *transferCommand) walkAndCheckFileAccessTimes(folder string, folderMapLock *sync.Mutex,
	folderMap map[string]string) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if c.link {
				_path, _info, _err := assist.GetRealPath(path)
				if _err != nil {
					realPathErr := fmt.Sprintf("get symbolic-link file [%s] real path failed", path)
					doLogError(_err, LEVEL_ERROR, realPathErr)
					return _err
				}
				if _info.IsDir() {
					folderMapLock.Lock()
					pathToCheck := folder + "-" + assist.NormalizeFilePath(path) + "/"
					if previousFolder, ok := folderMap[_path]; ok && strings.Index(previousFolder, pathToCheck) >= 0 {
						panicErr := fmt.Errorf("the symbolic-link folder [%s] --> [%s] result in a circle with folder [%s], path to check is [%s]", path, _path, previousFolder, pathToCheck)
						doLogError(panicErr, LEVEL_ERROR, "")
						folderMapLock.Unlock()
						return panicErr
					}
					folderMap[_path] = pathToCheck
					folderMapLock.Unlock()
					return filepath.Walk(_path, c.walkAndCheckFileAccessTimes(_path, folderMapLock, folderMap))
				}
				path = _path
				info = _info
			} else {
				if stat, err := os.Stat(path); err == nil && stat.IsDir() {
					info = stat
				}
			}
		}

		if !info.IsDir() {
			fileAccessTime := assist.GetFileAccessTime(info)
			if !c.matchLastAccessTime(fileAccessTime) {
				notMatchErr := errors.New("file access time not match time range")
				notMatchMsg := fmt.Sprintf("file [%s] access time [%v] not in range [%s]",
					path, fileAccessTime.UTC(), c.timeRange)
				doLogError(notMatchErr, LEVEL_ERROR, notMatchMsg)
				return notMatchErr
			}
		}
		return nil
	}
}

func (c *transferCommand) submitUploadTask(bucket, dir, arcDir, folder string, linkFolder bool, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType, pool concurrent.Pool, barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter,
	totalBytes, totalBytesForProgress, totalFiles *int64) error {

	defer func() {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
	}()

	// when -at accessTimeAccordingFolders and memoryEconomicalScanForUpload set true
	if c.at && config["accessTimeAccordingFolders"] == c_true {
		if config["memoryEconomicalScanForUpload"] == c_true {
			folderMap := make(map[string]string, 10)
			folderMapLock := new(sync.Mutex)
			if matchErr := filepath.Walk(folder, c.walkAndCheckFileAccessTimes(folder, folderMapLock, folderMap)); matchErr != nil {
				printf("the folder's sub-file access time is not in range [%s]", c.timeRange)
				return nil
			}
			c.timeRange = ""
		} else {
			printf("use the -at option, must set memoryEconomicalScanForUpload value true in config file")
			return nil
		}
	}

	if config["memoryEconomicalScanForUpload"] == c_true || c.link {
		if err := filepath.Walk(folder, c.getWalkFunc(bucket, dir, arcDir, folder, linkFolder, metadata,
			aclType, storageClassType, pool, barCh, limiter, totalBytes, totalBytesForProgress, totalFiles)); err != nil {
			return err
		}
	} else {
		c.scanPool = concurrent.NewNochanPool(-1)
		c.doScan(folder, c.getWalkFunc(bucket, dir, arcDir, folder, linkFolder, metadata,
			aclType, storageClassType, pool, barCh, limiter, totalBytes, totalBytesForProgress, totalFiles))
		c.scanPool.ShutDown()
	}

	if err, ok := c.scanError.Load().(error); ok {
		return err
	}
	return nil
}

func (c *transferCommand) recordStartFuncForUpload() time.Time {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "file size", "md5 value", "src --> dst", "cost(ms)", "status code", "success message", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "file size", "src --> dst", "cost(ms)", "status code", "error code", "error message", "request id")
	c.warningLogger.doRecord("[%s, %s, %s]", "file size", "src --> dst", "warn message")
	return start
}

func (c *transferCommand) recordPrepareFailed(bucket, key, filePath string, totalBytesForProgress, totalFiles *int64, failedError string) {
	atomic.AddInt64(totalFiles, 1)
	atomic.AddInt64(totalBytesForProgress, 1)
	doLog(LEVEL_ERROR, "Error happened when dealing with file path [%s], error message [%s]", filePath, failedError)
	c.failedLogger.doRecord("n/a, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", filePath, bucket, key, failedError)
	c.handleExecResultTransAction(0, 0, 0)
}

// error has been logged, return error just for print for user.
func (c *transferCommand) uploadFileOrFolder(bucket, dir, filePath string, limiter *ratelimit.RateLimiter, pool concurrent.Pool, barCh progress.SingleBarChan,
	totalBytes, totalBytesForProgress, totalFiles *int64, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType, cleanFolders *[]string) error {
	fileUrl, fileAbsErr := filepath.Abs(filePath)
	if fileAbsErr != nil {
		fileAbsErrorMsg := fmt.Sprintf("Get file path [%s] failed, [%s]", filePath, fileAbsErr.Error())
		c.recordPrepareFailed(bucket, dir, filePath, totalBytesForProgress, totalFiles, fileAbsErrorMsg)
		return fileAbsErr
	}
	stat, fileStatErr := os.Lstat(fileUrl)
	if fileStatErr != nil {
		fileStatErrorMsg := fmt.Sprintf("Gte file [%s] status faied, [%s]", fileUrl, fileStatErr.Error())
		c.recordPrepareFailed(bucket, dir, filePath, totalBytesForProgress, totalFiles, fileStatErrorMsg)
		return fileStatErr
	}

	linkFolder := false
	relativeFolder := ""
	if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
		if c.link {
			_url1, _stat, _err := assist.GetRealPath(fileUrl)
			if _err != nil {
				linkErrorMsg := fmt.Sprintf("Get real path for path [%s] failed, [%s]", fileUrl, _err.Error())
				c.recordPrepareFailed(bucket, dir, filePath, totalBytesForProgress, totalFiles, linkErrorMsg)
				return _err
			}
			if _stat.IsDir() {
				if !c.flat {
					relativeFolder = c.getRelativeFolder(fileUrl)
				}
				linkFolder = true
			}
			fileUrl = _url1
			stat = _stat
		} else {
			if _stat, err := os.Stat(fileUrl); err == nil && _stat.IsDir() {
				stat = _stat
			}
		}
	}
	arcDir := c.arcDir
	if !stat.IsDir() {
		key := dir
		key += stat.Name()
		if arcDir != "" {
			arcDir = arcDir + "/" + stat.Name()
		}
		atomic.AddInt64(totalBytes, stat.Size())
		if stat.Size() == 0 {
			atomic.AddInt64(totalBytesForProgress, 1)
		} else {
			atomic.AddInt64(totalBytesForProgress, stat.Size())
		}
		atomic.AddInt64(totalFiles, 1)
		pool.ExecuteFunc(func() interface{} {
			return c.handleExecResult(c.uploadFile(bucket, key, arcDir, fileUrl, stat, metadata, aclType, storageClassType, barCh, limiter, 2, nil), stat.Size())
		})

		return nil
	}
	if c.recursive {
		if c.arcDir == assist.NormalizeFilePath(fileUrl) {
			// print is not current
			archiveErrorMsg := fmt.Sprintf("The folder [%s] to be uploaded and the archive dir are same!", fileUrl)
			c.recordPrepareFailed(bucket, dir, filePath, totalBytesForProgress, totalFiles, archiveErrorMsg)
			return fmt.Errorf(archiveErrorMsg)
		}
		doLog(LEVEL_INFO, "Upload objects from local folder [%s] to cloud folder [%s] in the bucket [%s]", fileUrl, dir, bucket)
		if linkFolder {
			dir += relativeFolder
			if arcDir != "" {
				if !isObsFolder(arcDir) {
					arcDir = arcDir + "/"
				}
				arcDir += relativeFolder
			}
		}
		uploadDirError := c.submitUploadTask(bucket, dir, arcDir, fileUrl, linkFolder, metadata, aclType,
			storageClassType, pool, barCh, limiter, totalBytes, totalBytesForProgress, totalFiles)

		if uploadDirError != nil {
			doLogError(uploadDirError, LEVEL_ERROR, fmt.Sprintf("List local files in folder [%s] failed", fileUrl))
			return uploadDirError
		}
		if arcDir != "" {
			*cleanFolders = append(*cleanFolders, fileUrl)
		}
		return nil
	}

	notSupportErrorMsg := fmt.Sprintf("Must pass -r to upload folder [%s]!", fileUrl)
	c.recordPrepareFailed(bucket, dir, filePath, totalBytesForProgress, totalFiles, notSupportErrorMsg)
	return fmt.Errorf(notSupportErrorMsg)
}

func (c *transferCommand) uploadMultiFilesOrFolders(bucket, dir string, fileList []string, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType) error {
	start := c.recordStartFuncForUpload()
	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	barCh := newSingleBarChan()
	barCh.SetBytes(true)
	barCh.SetTemplate(progress.TpsAndSpeed)
	if c.force {
		barCh.Start()
	}

	limiter := c.createRateLimiter()

	var totalBytes int64
	var totalFiles int64
	var totalBytesForProgress int64
	var cleanFolders []string
	const maxErrorNums = 2
	uploadErrors := [maxErrorNums]error{}
	errInx := 0

	for _, filePath := range fileList {
		uploadError := c.uploadFileOrFolder(bucket, dir, strings.TrimSpace(filePath), limiter, pool, barCh,
			&totalBytes, &totalBytesForProgress, &totalFiles, metadata, aclType, storageClassType, &cleanFolders)
		if uploadError != nil {
			if errInx < maxErrorNums {
				uploadErrors[errInx] = uploadError
			}
			errInx = errInx + 1
		}
	}

	doLog(LEVEL_INFO, "Number of files to upload [%d], total size to upload [%d(B)]", totalFiles, totalBytes)
	progress.SetTotalCount(totalFiles)
	progress.SetTotalStream(totalBytes)
	barCh.SetTotalCount(totalBytesForProgress)
	if !c.force {
		barCh.Start()
	}

	pool.ShutDown()
	barCh.WaitToFinished()

	if c.arcDir != "" && cleanFolders != nil && atomic.LoadInt32(&c.abort) == 0 {
		// hint
		h := &assist.HintV2{}
		var totalDeletedCount int64
		h.MessageFunc = func() string {
			count := ""
			if tc := atomic.LoadInt64(&totalDeletedCount); tc > 0 {
				count = "[" + assist.Int64ToString(tc) + "]"
			}
			return fmt.Sprintf("Waitting for clean up archive surplus %s", count)
		}
		h.Start()
		needCleanFolders := []string{}
		for _, folder := range cleanFolders {
			if dirs, err := assist.FindDirsToDelete(folder); err == nil {
				needCleanFolders = append(needCleanFolders, dirs...)
			}
		}
		sort.Strings(needCleanFolders)
		for inx := len(needCleanFolders) - 1; inx >= 0; inx-- {
			if de := os.Remove(needCleanFolders[inx]); de == nil {
				atomic.AddInt64(&totalDeletedCount, 1)
				doLog(LEVEL_INFO, "Delete folder [%s] succeed", needCleanFolders[inx])
			} else {
				doLog(LEVEL_INFO, "Delete folder [%s] failed due to %s", needCleanFolders[inx], de.Error())
			}
		}
		h.End()
	}

	c.recordEndWithMetricsV2(start, totalFiles, progress.GetSucceedStream(), progress.GetTotalStream())

	if errInx > 0 {
		for _, uploadError := range uploadErrors {
			if uploadError != nil {
				printError(uploadError)
			}
		}
		if errInx >= maxErrorNums {
			printf("Upload file list failed large than max print size [%d]. Please checkout the failed records file to lookup all failed message.\n", maxErrorNums)
		}
		return assist.ErrUncompeleted
	}

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}

	return nil
}

func (c *transferCommand) uploadDir(bucket, dir, arcDir, folder string, linkFolder bool, folderStat os.FileInfo, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType) error {

	start := c.recordStartFuncForUpload()
	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	barCh := newSingleBarChan()
	barCh.SetBytes(true)
	barCh.SetTemplate(progress.TpsAndSpeed)
	if c.force {
		barCh.Start()
	}

	limiter := c.createRateLimiter()

	var totalBytes int64
	var totalFiles int64
	var totalBytesForProgress int64
	hasListError := c.submitUploadTask(bucket, dir, arcDir, folder, linkFolder, metadata, aclType,
		storageClassType, pool, barCh, limiter, &totalBytes, &totalBytesForProgress, &totalFiles)

	doLog(LEVEL_INFO, "Number of files to upload [%d], total size to upload [%d(B)]", totalFiles, totalBytes)
	progress.SetTotalCount(totalFiles)
	progress.SetTotalStream(totalBytes)
	barCh.SetTotalCount(totalBytesForProgress)
	if !c.force {
		barCh.Start()
	}

	pool.ShutDown()
	barCh.WaitToFinished()

	if arcDir != "" && atomic.LoadInt32(&c.abort) == 0 {
		// hint
		h := &assist.HintV2{}
		var totalDeletedCount int64
		h.MessageFunc = func() string {
			count := ""
			if tc := atomic.LoadInt64(&totalDeletedCount); tc > 0 {
				count = "[" + assist.Int64ToString(tc) + "]"
			}
			return fmt.Sprintf("Waitting for clean up archive surplus %s", count)
		}
		h.Start()

		if dirs, err := assist.FindDirsToDelete(folder); err == nil {
			sort.Strings(dirs)
			for inx := len(dirs) - 1; inx >= 0; inx-- {
				if de := os.Remove(dirs[inx]); de == nil {
					atomic.AddInt64(&totalDeletedCount, 1)
					doLog(LEVEL_INFO, "Delete folder [%s] succeed", dirs[inx])
				} else {
					doLog(LEVEL_INFO, "Delete folder [%s] failed due to %s", dirs[inx], de.Error())
				}
			}

		}
		h.End()
	}

	c.recordEndWithMetricsV2(start, totalFiles, progress.GetSucceedStream(), progress.GetTotalStream())
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("List local files in local folder [%s] failed", folder))
		return assist.ErrUncompeleted
	}

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}

	return nil
}
