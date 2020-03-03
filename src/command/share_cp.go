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
	"command/i18n"
	"concurrent"
	"fmt"
	"io"
	"net/url"
	"obs"
	"os"
	"path/filepath"
	"progress"
	"ratelimit"
	"strings"
	"sync/atomic"
)

type shareCpCommand struct {
	shareCommand
	objectKey string
}

func (c *shareCpCommand) constructGetObjectUrl(requestUrl interface{}, key string) string {
	if parsedUrl, ok := requestUrl.(*url.URL); ok {
		signedUrl := constructCommonUrl(parsedUrl, key)
		commonUrl := strings.Join(signedUrl, "")
		doLog(LEVEL_INFO, "The common url for getting object is [%s]", commonUrl)
		return commonUrl
	}
	return requestUrl.(string)
}

func (c *shareCpCommand) constructListObjectsUrlWithOutSignature(parsedUrl *url.URL, prefix string) string {
	signedUrl := make([]string, 0, 4)
	signedUrl = append(signedUrl, parsedUrl.Scheme)
	signedUrl = append(signedUrl, "://")
	signedUrl = append(signedUrl, parsedUrl.Host)
	if prefix != "" {
		signedUrl = append(signedUrl, fmt.Sprintf("/?%s=%s", "prefix", prefix))
	}
	return strings.Join(signedUrl, "")
}

func (c *shareCpCommand) constructGetObjectUrlWithOutSignature(requestUrl interface{}, key string) string {
	var parsedUrl *url.URL
	if _parsedUrl, ok := requestUrl.(*url.URL); ok {
		parsedUrl = _parsedUrl
	} else if _parsedUrl, err := url.Parse(requestUrl.(string)); err == nil {
		parsedUrl = _parsedUrl
	}

	if parsedUrl != nil {
		signedUrl := make([]string, 0, 4)
		signedUrl = append(signedUrl, parsedUrl.Scheme)
		signedUrl = append(signedUrl, "://")
		signedUrl = append(signedUrl, parsedUrl.Host)
		signedUrl = append(signedUrl, fmt.Sprintf("/%s", key))
		return strings.Join(signedUrl, "")
	}
	return requestUrl.(string)
}

func (c *shareCpCommand) getObjectMetadata(requestUrl interface{}, key string) (*MetaContext, error) {
	signedUrl := c.constructGetObjectUrl(requestUrl, key)

	requestHeaders := make(map[string][]string, 1)
	if parsedUrl, ok := requestUrl.(*url.URL); ok {
		requestHeaders["Host"] = []string{parsedUrl.Host}
	}

	output, err := obsClient.GetObjectWithSignedUrl(signedUrl, requestHeaders)
	if err == nil {
		defer output.Body.Close()
		return &MetaContext{
			Size:         output.ContentLength,
			LastModified: output.LastModified,
			ETag:         output.ETag,
			RequestId:    output.RequestId,
			Metadata:     output.Metadata,
		}, nil
	}
	return nil, err
}

func (c *shareCpCommand) checkSourceChangedForDownload(requestUrl interface{}, key string, originLastModified int64, abort *int32) error {
	if config["checkSourceChange"] == c_true {
		if metaContext, err := c.getObjectMetadata(requestUrl, key); err != nil {
			if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode == 404 {
				if abort != nil {
					atomic.CompareAndSwapInt32(abort, 0, 1)
				}
				return fmt.Errorf("Source object [%s] doesnot exist", key)
			}
		} else if originLastModified != metaContext.LastModified.Unix() {
			if abort != nil {
				atomic.CompareAndSwapInt32(abort, 0, 1)
			}
			return fmt.Errorf("Source object [%s] changed", key)
		}
	}
	return nil
}

func (c *shareCpCommand) downloadBigFile(requestUrl interface{}, downloadCloudUrl, key, fileUrl string, fileStat os.FileInfo, metaContext *MetaContext,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int) (requestId string, status int, md5Value string, downloadFileError error) {

	if metaContext.Size == 0 {
		return c.downloadSmallFileWithRetry(requestUrl, downloadCloudUrl, key, fileUrl, fileStat, metaContext, barCh, limiter)
	}

	_downloadCloudUrl := downloadCloudUrl
	if strings.HasPrefix(downloadCloudUrl, "http://") {
		_downloadCloudUrl = downloadCloudUrl[len("http://"):]
	} else if strings.HasPrefix(downloadCloudUrl, "https://") {
		_downloadCloudUrl = downloadCloudUrl[len("https://"):]
	}

	checkpointFile := c.getCheckpointFile(_downloadCloudUrl, key, "", dm)
	dfc := &DownloadFileCheckpoint{}
	stat, err := os.Stat(checkpointFile)
	needPrepare := true
	if err == nil {
		if stat.IsDir() {
			downloadFileError = fmt.Errorf("Checkpoint file for downloading [%s] is a folder", key)
			return
		}
		err = c.loadCheckpoint(checkpointFile, dfc)
		if err != nil {
			if err = os.Remove(checkpointFile); err != nil {
				downloadFileError = err
				return
			}
		} else if !dfc.isValid("", key, "", fileUrl, metaContext) {
			if dfc.TempFileInfo.TempFileUrl != "" {
				os.Remove(dfc.TempFileInfo.TempFileUrl)
			}

			if err = os.Remove(checkpointFile); err != nil {
				downloadFileError = err
				return
			}
		} else {
			needPrepare = false
		}
	}
	if needPrepare {
		err = c.prepareDownloadFileCheckpoint("", key, "", fileUrl, metaContext, dfc, barCh == nil)
		if err != nil {
			os.Remove(dfc.TempFileInfo.TempFileUrl)
			downloadFileError = err
			return
		}
		err = c.recordCheckpointFile(checkpointFile, dfc)
		if err != nil {
			os.Remove(dfc.TempFileInfo.TempFileUrl)
			downloadFileError = err
			return
		}
	}

	barChFlag := false
	if barCh == nil {
		barCh = newSingleBarChan()
		barCh.SetBytes(true)
		barCh.SetTemplate(progress.SpeedOnly)
		barCh.SetTotalCount(dfc.ObjectInfo.Size)
		progress.SetTotalStream(dfc.ObjectInfo.Size)
		barChFlag = true
	}

	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	var abort int32
	var metadata map[string]string
	abort, status, requestId, metadata, downloadFileError = c.downloadBigFileConcurrent(dfc, checkpointFile, barChFlag, barCh, limiter, requestUrl)

	if barChFlag {
		barCh.WaitToFinished()
	}
	if abort == 1 {
		os.Remove(dfc.TempFileInfo.TempFileUrl)
		if err = os.Remove(checkpointFile); err != nil {
			downloadFileError = err
			return
		}
	}
	if downloadFileError != nil {
		return
	}
	start := assist.GetUtcNow()
	if batchFlag <= 1 {
		printf("Waiting to rename temporary file...")
	}
	_writeBufferIoSize, transErr := assist.TranslateToInt64(config["writeBufferIoSize"])
	if transErr != nil {
		_writeBufferIoSize = defaultWriteBufferIoSize
	}
	if err = assist.RenameFile(dfc.TempFileInfo.TempFileUrl, dfc.FileUrl, config["forceOverwriteForDownload"] == c_true, int(_writeBufferIoSize), config["fsyncForDownload"] == c_true); err != nil {
		downloadFileError = err
	} else {
		cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
		doLog(LEVEL_DEBUG, "Rename temporary file [%s] to [%s] successfully, cost [%d]",
			dfc.TempFileInfo.TempFileUrl, dfc.FileUrl, cost)
		if err = os.Remove(checkpointFile); err != nil {
			doLog(LEVEL_WARN, "Download big file [%s] from key [%s] successfully, but remove checkpoint file [%s] failed",
				dfc.FileUrl, dfc.Key, checkpointFile)
		}
	}

	if c.verifyMd5 {
		if _md5Value, ok := metadata[checkSumKey]; ok && _md5Value != "" {
			md5Value = _md5Value
			if barChFlag {
				h := &assist.HintV2{}
				h.Message = c_waiting_caculate_md5
				h.Start()
				if etag, err := md5File(fileUrl); err != nil {
					downloadFileError = &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after downloading file [%s], %s", fileUrl, err.Error())}
				} else {
					localMd5 := assist.Hex(etag)
					if localMd5 != md5Value {
						doLog(LEVEL_ERROR, "Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s], will try to delete downloaded file", fileUrl, localMd5, md5Value)
						downloadFileError = &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s]", fileUrl, localMd5, md5Value)}
					}
				}
				h.End()
			} else {
				if etag, err := md5File(fileUrl); err != nil {
					downloadFileError = &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after downloading file [%s], %s", fileUrl, err.Error())}
				} else {
					localMd5 := assist.Hex(etag)
					if localMd5 != md5Value {
						doLog(LEVEL_ERROR, "Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s], will try to delete downloaded file", fileUrl, localMd5, md5Value)
						downloadFileError = &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s]", fileUrl, localMd5, md5Value)}
					}
				}
			}
		} else {
			objectSizeStr := c_na
			if metaContext != nil {
				objectSizeStr = normalizeBytes(metaContext.Size)
			}

			warnMessage := fmt.Sprintf("Cannot get the valid md5 value of key [%s] to check", key)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> %s, warn message [%s]",
				objectSizeStr, downloadCloudUrl, fileUrl, warnMessage)
			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	}
	return
}

func (c *shareCpCommand) downloadSmallFile(requestUrl interface{}, downloadCloudUrl, key, fileUrl string, fileStat os.FileInfo, metaContext *MetaContext,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (requestId string, status int, md5Value string,
	noRepeatable bool, readed int64, downloadFileError error) {

	requestHeaders := make(map[string][]string, 1)
	var signedUrl string
	if parsedUrl, ok := requestUrl.(*url.URL); ok {
		requestHeaders["Host"] = []string{parsedUrl.Host}
		signedUrl = c.constructGetObjectUrl(parsedUrl, key)
	} else {
		signedUrl = requestUrl.(string)
	}

	output, err := obsClient.GetObjectWithSignedUrl(signedUrl, requestHeaders)
	if err != nil {
		downloadFileError = err
		if obsError, ok := downloadFileError.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
			strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
			noRepeatable = true
		}
		return
	}

	_readBufferIoSize, transErr := assist.TranslateToInt64(config["readBufferIoSize"])
	if transErr != nil {
		_readBufferIoSize = defaultReadBufferIoSize
	}
	if _readBufferIoSize < minReadBufferIoSize {
		_readBufferIoSize = minReadBufferIoSize
	}
	if output.ContentLength < _readBufferIoSize {
		_readBufferIoSize = output.ContentLength
	}

	defer output.Body.Close()
	var reader io.Reader = progress.NewSingleProgressReader(bufio.NewReaderSize(output.Body, int(_readBufferIoSize)), -1, false, barCh)

	if limiter == nil {
		limiter = c.createRateLimiter()
	}

	if limiter != nil {
		reader = ratelimit.NewRateLimitReaderWithLimiter(reader, limiter)
	}

	fd, err := c.createFile(key, fileUrl, fileStat)
	if err != nil {
		downloadFileError = &errorWrapper{
			err:       err,
			requestId: output.RequestId,
		}
		noRepeatable = true
		return
	}

	var md5Writer io.Writer = nilWriter
	if c.verifyMd5 {
		if _md5Value, ok := output.Metadata[checkSumKey]; ok && _md5Value != "" {
			md5Value = _md5Value
			md5Writer = assist.GetMd5Writer()
		} else {
			warnMessage := fmt.Sprintf("Cannot get the valid md5 value of key [%s] to check", key)
			warnLoggerMessage := fmt.Sprintf("%s, %s --> %s, warn message [%s]",
				normalizeBytes(metaContext.Size), downloadCloudUrl, fileUrl, warnMessage)

			c.recordWarnMessage(warnMessage, warnLoggerMessage)
		}
	}

	if fd != nil && err == nil {
		defer fd.Close()

		_writeBufferIoSize, transErr := assist.TranslateToInt64(config["writeBufferIoSize"])
		if transErr != nil {
			_writeBufferIoSize = defaultWriteBufferIoSize
		}

		bufWriter := bufio.NewWriterSize(fd, int(_writeBufferIoSize))
		p := assist.GetByteArrayFromPool()
		for {
			n, err := reader.Read(p)
			if n > 0 {
				readed += int64(n)
				slice := p[0:n]
				wcnt, werr := bufWriter.Write(slice)
				if _, writeErr := md5Writer.Write(slice); writeErr != nil {
					doLog(LEVEL_WARN, "Write md5 value failed, %s", writeErr.Error())
				}
				if werr != nil {
					downloadFileError = &errorWrapper{
						err:       werr,
						requestId: output.RequestId,
					}
					assist.PutByteArrayToPool(p)
					noRepeatable = true
					return
				}

				if wcnt != n {
					downloadFileError = &errorWrapper{
						err:       fmt.Errorf("Write to file [%s] failed, expect [%d], actual [%d]", fileUrl, n, wcnt),
						requestId: output.RequestId,
					}
					assist.PutByteArrayToPool(p)
					noRepeatable = true
					return
				}
			}

			if err != nil {
				if err != io.EOF {
					downloadFileError = &errorWrapper{
						err:       err,
						requestId: output.RequestId,
					}
					assist.PutByteArrayToPool(p)
					return
				}
				break
			}
		}
		assist.PutByteArrayToPool(p)

		if err := bufWriter.Flush(); err != nil {
			downloadFileError = &errorWrapper{
				err:       err,
				requestId: output.RequestId,
			}
			noRepeatable = true
			return
		}

		if config["fsyncForDownload"] == c_true {
			if err := fd.Sync(); err != nil {
				downloadFileError = &errorWrapper{
					err:       err,
					requestId: output.RequestId,
				}
				noRepeatable = true
				return
			}
		}

	}

	if changedErr := c.checkSourceChangedForDownload(requestUrl, key, metaContext.LastModified.Unix(), nil); changedErr != nil {
		downloadFileError = &errorWrapper{
			err:       changedErr,
			requestId: output.RequestId,
		}
		noRepeatable = true
		return
	}

	if md5Value != "" {
		localMd5 := assist.GetHexMd5(md5Writer)
		if localMd5 != md5Value {
			doLog(LEVEL_ERROR, "Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s], will try to delete downloaded file", fileUrl, localMd5, md5Value)
			downloadFileError = &errorWrapper{
				err:       &verifyMd5Error{msg: fmt.Sprintf("Verify md5 failed after downloading file [%s], local md5 [%s] remote md5 [%s]", fileUrl, localMd5, md5Value)},
				requestId: output.RequestId,
			}
			return
		}
	}

	requestId = output.RequestId

	if requestId == "" {
		if ret, ok := output.ResponseHeaders["x-obs-request-id"]; ok {
			requestId = ret[0]
		}
	}

	status = output.StatusCode
	return
}

func (c *shareCpCommand) downloadSmallFileWithRetry(requestUrl interface{}, downloadCloudUrl, key, fileUrl string, fileStat os.FileInfo, metaContext *MetaContext,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter) (requestId string, status int, md5Value string, downloadFileError error) {

	objectSize := metaContext.Size
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
		requestId, status, md5Value, noRepeatable, readed, downloadFileError = c.downloadSmallFile(requestUrl, downloadCloudUrl, key, fileUrl, fileStat, metaContext, barCh, limiter)
		if downloadFileError == nil || noRepeatable || retryCount >= maxRetryCount {
			break
		}

		if readed > 0 {
			barCh.Send64(-readed)
			progress.AddEffectiveStream(-readed)
			progress.AddFinishedStream(-readed)
		}

		doLog(LEVEL_WARN, "Failed to download object and will try again, err:%s", downloadFileError.Error())
		assist.SleepByCount(retryCount)
		retryCount++
	}

	if barChFlag {
		barCh.WaitToFinished()
	}

	return
}

func (c *shareCpCommand) downloadFile(requestUrl interface{}, key, fileUrl string, fileStat os.FileInfo,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int, fastFailed error) int {
	metaContext, metaErr := c.getObjectMetadata(requestUrl, key)
	return c.downloadFileWithMetaContext(requestUrl, key, metaContext, metaErr, fileUrl, fileStat, barCh, limiter, batchFlag, fastFailed)
}

func (c *shareCpCommand) downloadFileWithMetaContext(requestUrl interface{}, key string, metaContext *MetaContext, metaErr error, fileUrl string, fileStat os.FileInfo,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter, batchFlag int, fastFailed error) int {
	objectSizeStr := c_na
	if metaContext != nil {
		objectSizeStr = normalizeBytes(metaContext.Size)
	}

	downloadCloudUrl := c.constructGetObjectUrlWithOutSignature(requestUrl, key)

	if fastFailed != nil {
		c.failedLogger.doRecord("%s, %s --> %s, n/a, n/a, n/a, error message [%s], n/a", objectSizeStr, downloadCloudUrl, fileUrl, fastFailed.Error())
		return 0
	}

	if batchFlag == 2 && atomic.LoadInt32(&c.abort) == 1 {
		c.failedLogger.doRecord("%s, %s --> %s, n/a, n/a, error code [%s], error message [%s], n/a", objectSizeStr, downloadCloudUrl, fileUrl,
			"AbortError", "Task is aborted")
		return 0
	}

	if c.update {
		changed, err := c.ensureKeyForDownload(metaContext, metaErr, fileStat, batchFlag, key)
		if !changed {
			if err == nil {
				if barCh != nil {
					if metaContext.Size <= 0 {
						barCh.Send64(1)
					} else {
						barCh.Send64(metaContext.Size)
					}
				}

				progress.AddFinishedStream(metaContext.Size)
				if batchFlag >= 1 {
					c.succeedLogger.doRecord("%s, n/a, %s --> %s, n/a, n/a, success message [skip since the source is not changed], n/a", objectSizeStr, downloadCloudUrl, fileUrl)
				}

				if batchFlag != 2 {
					printf("%s, %s --> %s, skip since the source is not changed", objectSizeStr, downloadCloudUrl, fileUrl)
				}

				return 2
			}

			if batchFlag >= 1 {
				c.failedLogger.doRecord("%s, %s --> %s, n/a, n/a, n/a, error message [skip since the status of source is unknown], n/a", objectSizeStr, downloadCloudUrl, fileUrl)
			}
			if batchFlag != 2 {
				printf("%s --> %s, skip since the status of source is unknown", downloadCloudUrl, fileUrl)
			}

			return 0
		}
	}

	downloadFileError := metaErr
	if c.dryRun {
		if downloadFileError == nil {
			if barCh != nil {
				if metaContext.Size <= 0 {
					barCh.Send64(1)
				} else {
					barCh.Send64(metaContext.Size)
				}
			}
			progress.AddFinishedStream(metaContext.Size)
			if batchFlag >= 1 {
				c.succeedLogger.doRecord("%s, n/a, %s --> %s, n/a, n/a, success message [dry run done], n/a", objectSizeStr, downloadCloudUrl, fileUrl)
			}
			if batchFlag != 2 {
				printf("\nDownload dry run successfully, %s, %s --> %s", objectSizeStr, downloadCloudUrl, fileUrl)
			}
			return 1
		}
		if batchFlag >= 1 {
			c.failedLogger.doRecord("%s, %s --> %s, n/a, n/a, n/a, error message [dry run done with error - %s], n/a", objectSizeStr, downloadCloudUrl, fileUrl, downloadFileError.Error())
		}
		if batchFlag != 2 {
			logError(downloadFileError, LEVEL_INFO, fmt.Sprintf("\nDownload dry run failed, %s --> %s", downloadCloudUrl, fileUrl))
		}
		return 0
	}

	var requestId string
	var status int
	var md5Value string
	start := assist.GetUtcNow()
	addCostFlag := false
	if downloadFileError == nil {
		if isObsFolder(key) {
			if fileStat == nil {
				downloadFileError = assist.MkdirAll(fileUrl, os.ModePerm)
			} else if !fileStat.IsDir() {
				downloadFileError = fmt.Errorf("Cannot create the folder [%s] due to a same file exits", fileUrl)
			}
			if downloadFileError == nil && barCh != nil {
				barCh.Send64(1)
			}
		} else {
			objectSize := metaContext.Size
			addCostFlag = true
			if objectSize >= c.bigfileThreshold {
				requestId, status, md5Value, downloadFileError = c.downloadBigFile(requestUrl, downloadCloudUrl, key, fileUrl, fileStat, metaContext, barCh, limiter, batchFlag)
			} else {
				requestId, status, md5Value, downloadFileError = c.downloadSmallFileWithRetry(requestUrl, downloadCloudUrl, key, fileUrl, fileStat, metaContext, barCh, limiter)
			}

			if _, ok := downloadFileError.(*verifyMd5Error); ok {
				if err := os.Remove(fileUrl); err == nil {
					doLog(LEVEL_INFO, "Delete file [%s] successfully", fileUrl)
				} else {
					warnMessage := fmt.Sprintf("Delete file [%s] failed - %s", fileUrl, err.Error())
					warnLoggerMessage := fmt.Sprintf("%s, %s --> %s, warn message [%s]",
						objectSizeStr, downloadCloudUrl, fileUrl, warnMessage)
					c.recordWarnMessage(warnMessage, warnLoggerMessage)
				}
			}

			if downloadFileError == nil {
				if barCh != nil && objectSize <= 0 {
					barCh.Send64(1)
				}

				if (!c.verifyMd5 || md5Value == "") && c.verifyLength {
					if stat, err := os.Stat(fileUrl); err == nil {
						if !stat.IsDir() && stat.Size() != metaContext.Size {
							doLog(LEVEL_ERROR, "Verify length failed after downloading file [%s], local length [%d] remote length [%d], will try to delete downloaded file", fileUrl, stat.Size(), metaContext.Size)
							if _err := os.Remove(fileUrl); _err == nil {
								doLog(LEVEL_INFO, "Delete local file [%s] successfully", fileUrl)
							} else {
								warnMessage := fmt.Sprintf("Delete local file [%s] failed - %s", fileUrl, _err.Error())
								warnLoggerMessage := fmt.Sprintf("%s, %s --> %s, warn message [%s]",
									objectSizeStr, downloadCloudUrl, fileUrl, warnMessage)
								c.recordWarnMessage(warnMessage, warnLoggerMessage)
							}
							downloadFileError = &errorWrapper{
								err:       &verifyLengthError{msg: fmt.Sprintf("Verify length failed after downloading file [%s], local length [%d] remote length [%d]", fileUrl, stat.Size(), metaContext.Size)},
								requestId: requestId,
							}
						}
					} else {
						warnMessage := fmt.Sprintf("Download file [%s] from key [%s] of [%s] successfully - but can not verify length - %s",
							fileUrl, key, downloadCloudUrl, err.Error())
						warnLoggerMessage := fmt.Sprintf("%s, %s --> %s, warn message [%s]",
							objectSizeStr, downloadCloudUrl, fileUrl, warnMessage)
						c.recordWarnMessage(warnMessage, warnLoggerMessage)
					}
				}
			}
		}
	}

	if md5Value == "" {
		md5Value = c_na
	}

	cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

	if batchFlag >= 1 {
		if downloadFileError == nil {
			c.succeedLogger.doRecord("%s, %s, %s --> %s, cost [%d], status [%d], success message [download succeed], request id [%s]", objectSizeStr, md5Value, downloadCloudUrl, fileUrl, cost, status, requestId)
		} else {
			_status, _code, _message, _requestId := c.checkAbort(downloadFileError, 401, 405)
			c.failedLogger.doRecord("%s, %s --> %s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", objectSizeStr, downloadCloudUrl, fileUrl, cost,
				_status, _code, _message, _requestId)
		}
	}

	if batchFlag == 2 {
		if addCostFlag {
			c.ensureMaxCostAndMinCost(cost)
		}
		atomic.AddInt64(&c.totalCost, cost)
	} else {
		if downloadFileError == nil {
			printf("\nDownload successfully, %s, %s, %s --> %s, cost [%d], status [%d], request id [%s]", objectSizeStr, md5Value, downloadCloudUrl, fileUrl, cost, status, requestId)
			doLog(LEVEL_DEBUG, "Download successfully, %s, %s, %s --> %s, cost [%d], status [%d], request id [%s]", objectSizeStr, md5Value, downloadCloudUrl, fileUrl, cost, status, requestId)
		} else {
			logError(downloadFileError, LEVEL_INFO, fmt.Sprintf("\nDownload failed, %s --> %s, cost [%d]", downloadCloudUrl, fileUrl, cost))
		}
	}
	if downloadFileError == nil {
		return 1
	}

	return 0
}

func (c *shareCpCommand) constructListObjectsUrl(parsedUrl *url.URL, prefix string) string {
	signedUrl := constructCommonUrl(parsedUrl, "")

	if prefix != "" {
		signedUrl = append(signedUrl, fmt.Sprintf("%s=%s&", "prefix", url.QueryEscape(prefix)))
	}

	signedUrl = append(signedUrl, fmt.Sprintf("%s=%d&", "max-keys", defaultListMaxKeys))
	commonUrl := strings.Join(signedUrl, "")
	doLog(LEVEL_INFO, "The common url for listing objects is [%s]", commonUrl)
	return commonUrl
}

func (c *shareCpCommand) submitDownloadTask(parsedUrl *url.URL, listCloudUrl, dir, folder, relativePrefix string, barCh progress.SingleBarChan,
	limiter *ratelimit.RateLimiter, pool concurrent.Pool) (totalBytes int64, totalBytesForProgress int64, totalObjects int64, hasListError error) {

	defer func() {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
		}
	}()

	commonUrl := c.constructListObjectsUrl(parsedUrl, dir)

	requestHeaders := map[string][]string{"Host": []string{parsedUrl.Host}}
	var signedUrl []string
	var marker string
	for {
		if atomic.LoadInt32(&c.abort) == 1 {
			return
		}
		start := assist.GetUtcNow()

		if signedUrl == nil {
			signedUrl = make([]string, 0, 2)
		} else {
			signedUrl = signedUrl[:0]
		}

		signedUrl = append(signedUrl, commonUrl)

		if marker != "" {
			signedUrl = append(signedUrl, fmt.Sprintf("%s=%s", "marker", url.QueryEscape(marker)))
		}

		requestUrl := strings.Join(signedUrl, "")

		output, err := obsClient.ListObjectsWithSignedUrl(requestUrl, requestHeaders)
		if err != nil {
			hasListError = err
			break
		} else {
			cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000
			doLog(LEVEL_INFO, "List objects through [%s] to download successfully, cost [%d], request id [%s]", listCloudUrl, cost, output.RequestId)
		}

		for _, content := range output.Contents {
			if atomic.LoadInt32(&c.abort) == 1 {
				return
			}
			key := content.Key
			fileName := key
			if relativePrefix != "" {
				if index := strings.Index(fileName, relativePrefix); index >= 0 {
					fileName = fileName[len(relativePrefix):]
				}
			}

			var fastFailed error
			if checkEmptyFolder("", key, dm) {
				fastFailed = fmt.Errorf("Cannot download the specified key [%s]", key)
			}

			fileUrl := assist.NormalizeFilePath(folder + "/" + fileName)
			fileStat, statErr := os.Stat(fileUrl)
			if statErr != nil {
				doLog(LEVEL_WARN, "Stat file failed, %s", statErr.Error())
			}
			if isObsFolder(key) {

				if c.matchFolder {
					if c.matchExclude(key) {
						continue
					}

					if !c.matchInclude(key) {
						continue
					}

					if !c.matchLastModifiedTime(content.LastModified) {
						continue
					}
				}

				atomic.AddInt64(&totalBytesForProgress, 1)
				atomic.AddInt64(&totalObjects, 1)

				metaContext := &MetaContext{
					ETag:         content.ETag,
					LastModified: content.LastModified,
					Size:         0,
				}
				pool.ExecuteFunc(func() interface{} {
					return c.handleExecResultTransAction(c.downloadFileWithMetaContext(parsedUrl, key,
						metaContext, nil, fileUrl, fileStat, barCh, limiter, 2, fastFailed), 0, 0)
				})
			} else {
				if c.matchExclude(key) {
					continue
				}

				if !c.matchInclude(key) {
					continue
				}

				if !c.matchLastModifiedTime(content.LastModified) {
					continue
				}

				if !c.force && !confirm(fmt.Sprintf("Do you want download key [%s] to [%s] ? Please input (y/n) to confirm:", key, fileUrl)) {
					continue
				}

				metaContext := &MetaContext{
					ETag:         content.ETag,
					LastModified: content.LastModified,
					Size:         content.Size,
				}
				atomic.AddInt64(&totalBytes, metaContext.Size)
				atomic.AddInt64(&totalObjects, 1)
				if metaContext.Size == 0 {
					atomic.AddInt64(&totalBytesForProgress, 1)
				} else {
					atomic.AddInt64(&totalBytesForProgress, metaContext.Size)
				}
				pool.ExecuteFunc(func() interface{} {
					return c.handleExecResult(c.downloadFileWithMetaContext(parsedUrl, key, metaContext, nil, fileUrl, fileStat, barCh, limiter, 2, fastFailed), metaContext.Size)
				})
			}

		}

		if !output.IsTruncated {
			doLog(LEVEL_INFO, "List objects through [%s] to download finished, marker [%s]", listCloudUrl, marker)
			break
		}
		marker = output.NextMarker
	}

	return
}

func (c *shareCpCommand) downloadDir(parsedUrl *url.URL, dir, folder string, folderStat os.FileInfo) error {
	start := c.recordStartFuncForDownload()
	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	barCh := newSingleBarChan()
	barCh.SetBytes(true)
	barCh.SetTemplate(progress.TpsAndSpeed)
	if c.force {
		barCh.Start()
	}

	limiter := c.createRateLimiter()

	var relativePrefix string
	if c.flat {
		if dir != "" && !isObsFolder(dir) {
			dir += "/"
		}
		relativePrefix = dir
	} else {
		relativePrefix = dir
		if isObsFolder(relativePrefix) {
			relativePrefix = relativePrefix[:len(relativePrefix)-1]
		}
		if index := strings.LastIndex(relativePrefix, "/"); index >= 0 {
			relativePrefix = relativePrefix[:index+1]
		} else {
			relativePrefix = ""
		}
	}

	listCloudUrl := c.constructListObjectsUrlWithOutSignature(parsedUrl, dir)
	totalBytes, totalBytesForProgress, totalObjects, hasListError := c.submitDownloadTask(parsedUrl, listCloudUrl, dir, folder, relativePrefix, barCh, limiter, pool)

	doLog(LEVEL_INFO, "Number of objects to download [%d], total size to download [%d(B)]", totalObjects, totalBytes)
	progress.SetTotalCount(totalObjects)
	progress.SetTotalStream(totalBytes)
	barCh.SetTotalCount(totalBytesForProgress)
	if !c.force {
		barCh.Start()
	}

	pool.ShutDown()
	barCh.WaitToFinished()
	c.recordEndWithMetricsV2(start, totalObjects, progress.GetSucceedStream(), progress.GetTotalStream())
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("List objects through [%s] to download failed", listCloudUrl))
		return assist.ErrUncompeleted
	}
	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}
	return nil
}

func initShareCp() command {
	c := &shareCpCommand{}
	c.key = "share-cp"
	c.usage = c_share_cp_usage
	c.description = "download objects using authorization code and access code"
	c.skipCheckAkSk = true

	c.define = func() {
		c.init()
		c.partSize = 0
		c.bigfileThreshold = 0
		c.warn = atomic.Value{}
		c.warnFlag = 0

		c.flagSet.BoolVar(&c.dryRun, "dryRun", false, "")
		c.flagSet.BoolVar(&c.verifyLength, "vlength", false, "")
		c.flagSet.BoolVar(&c.verifyMd5, "vmd5", false, "")
		c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
		c.flagSet.BoolVar(&c.matchFolder, "mf", false, "")
		c.flagSet.IntVar(&c.jobs, "j", 0, "")
		c.flagSet.IntVar(&c.parallel, "p", 0, "")
		c.flagSet.StringVar(&c.bigfileThresholdStr, "threshold", "", "")
		c.flagSet.StringVar(&c.partSizeStr, "ps", "", "")
		c.flagSet.StringVar(&c.checkpointDir, "cpd", "", "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.include, "include", "", "")
		c.flagSet.StringVar(&c.exclude, "exclude", "", "")
		c.flagSet.StringVar(&c.timeRange, "timeRange", "", "")
		c.flagSet.StringVar(&c.tempFileDir, "tempFileDir", "", "")
		c.flagSet.StringVar(&c.objectKey, "key", "", "")

		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.update, "u", false, "")
		c.flagSet.BoolVar(&c.flat, "flat", false, "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) < 2 {
			c.showHelp()
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		originFileUrl := args[1]
		fileUrl, err := filepath.Abs(originFileUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		parsedUrl, allowedPrefix, err := c.prepareAccessUrl(args[0], args[2:])
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		defer func() {
			c.printAuthorizedPrefix(allowedPrefix)
		}()

		if c.objectKey == "" && c.recursive {
			c.objectKey = allowedPrefix
		} else if allowedPrefix != "" && !strings.HasPrefix(c.objectKey, allowedPrefix) {
			printf("Error: Invalid key [%s], must start with [%s]", c.objectKey, allowedPrefix)
			return assist.ErrInvalidArgs
		}

		if !c.prepareOptions() {
			return assist.ErrInvalidArgs
		}

		c.printStart()

		stat, err := os.Lstat(fileUrl)
		if !c.recursive {
			if c.objectKey == "" {
				printf("Error: No key specified for download, please specify the key parameter")
				return assist.ErrInvalidArgs
			}

			if isObsFolder(c.objectKey) {
				if !c.force && !confirm(fmt.Sprintf("Do you forget pass \"-r\" to recursively download? \nThis command will only download a empty folder as [%s]. Please input (y/n) to confirm:", fileUrl)) {
					return nil
				}
			}

			if (err != nil && (isObsFolder(originFileUrl) || strings.HasSuffix(originFileUrl, "\\"))) || (err == nil && stat.IsDir()) {
				fileName := c.objectKey
				if index := strings.LastIndex(fileName, "/"); index >= 0 {
					fileName = fileName[index+1:]
				}
				fileUrl = assist.NormalizeFilePath(fileUrl + "/" + fileName)
				stat = nil
			}

			if c.forceRecord {
				return c.ensureOuputAndStartLogger(func() error {
					c.printParams(true, true, false, true)
					c.recordStartFuncForDownload()
					ret := c.downloadFile(parsedUrl, c.objectKey, fileUrl, stat, nil, nil, 1, nil)
					if ret >= 1 {
						progress.AddSucceedCount(1)
						return nil
					}
					progress.AddFailedCount(1)
					return assist.ErrExecuting
				}, true)
			}

			c.printParams(false, true, false, true)
			ret := c.downloadFile(parsedUrl, c.objectKey, fileUrl, stat, nil, nil, 0, nil)
			if warn, ok := c.warn.Load().(error); ok {
				printWarn(warn)
			}
			if ret == 0 {
				return assist.ErrExecuting
			}
			return nil
		}

		if err != nil {
			if err = assist.MkdirAll(fileUrl, os.ModePerm); err != nil {
				printError(err)
				return assist.ErrInvalidArgs
			}

			stat, err = os.Lstat(fileUrl)
			if err != nil {
				printError(err)
				return assist.ErrFileNotFound
			}
		} else if !stat.IsDir() {
			printf("Error: Cannot download to the folder [%s] due to a file with the same name exits", fileUrl)
			return assist.ErrInvalidArgs
		}

		if err = c.ensureOutputDirectory(); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		if err = c.startLogger(true); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		defer c.endLogger()

		key := c.objectKey
		if key == "" {
			key = allowedPrefix
		}
		c.printParams(true, true, false, true)
		doLog(LEVEL_INFO, "Download objects from cloud folder [%s] of [%s] to local folder [%s]", key, parsedUrl.Host, fileUrl)
		return c.downloadDir(parsedUrl, key, fileUrl, stat)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("download objects using authorization code and access code"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil share-cp authorization_code file_url|folder_url -key=xxx [-ac=xxx] [-dryRun] [-tempFileDir=xxx] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil share-cp file://authorization_code_file_url file_url|folder_url -key=xxx [-ac=xxx] [-dryRun] [-tempFileDir=xxx] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax())
		printf("")
		p.Printf("Syntax 3:")
		printf("%2s%s", "", "obsutil share-cp authorization_code folder_url -r [-key=xxx] [-ac=xxx] [-dryRun] [-tempFileDir=xxx] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax())
		printf("")
		p.Printf("Syntax 4:")
		printf("%2s%s", "", "obsutil share-cp file://authorization_code_file_url folder_url -r [-key=xxx] [-ac=xxx] [-dryRun] [-tempFileDir=xxx] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-key=xxx")
		printf("%4s%s", "", p.Sprintf("the key to download, or the prefix to batch download"))
		printf("")
		printf("%2s%s", "", "-ac=xxx")
		printf("%4s%s", "", p.Sprintf("the access code"))
		printf("")
		printf("%2s%s", "", "-dryRun")
		printf("%4s%s", "", p.Sprintf("conduct a dry run"))
		printf("")
		printf("%2s%s", "", "-tempFileDir=xxx")
		printf("%4s%s", "", p.Sprintf("the temp file dir, used to save temporary files during the objects are downloading"))
		printf("")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch downloads objects by prefix"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while objects are to be downloaded"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files"))
		printf("")
		printf("%2s%s", "", "-u")
		printf("%4s%s", "", p.Sprintf("download the changed sources only"))
		printf("")
		printf("%2s%s", "", "-vlength")
		printf("%4s%s", "", p.Sprintf("verify the size after the objects are downloaded"))
		printf("")
		printf("%2s%s", "", "-vmd5")
		printf("%4s%s", "", p.Sprintf("verify the MD5 value after the objects are downloaded"))
		printf("")
		printf("%2s%s", "", "-flat")
		printf("%4s%s", "", p.Sprintf("download the sources without the relative parent prefix"))
		printf("")
		printf("%2s%s", "", "-cpd=xxx")
		printf("%4s%s", "", p.Sprintf("the directory where the part records reside, used to record the progress of download jobs"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent download jobs, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-p=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent download tasks (a task is a sub-job), the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-ps=auto")
		printf("%4s%s", "", p.Sprintf("the part size of each download task, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-threshold=52428800")
		printf("%4s%s", "", p.Sprintf("the threshold, if it is exceeded, the download job will be divided into multiple tasks by the part size, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-include=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be included"))
		printf("")
		printf("%2s%s", "", "-exclude=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be excluded"))
		printf("")
		printf("%2s%s", "", "-timeRange=time1-time2")
		printf("%4s%s", "", p.Sprintf("the time range, between which the objects will be uploaded, downloaded or copied"))
		printf("")
		printf("%2s%s", "", "-mf")
		printf("%4s%s", "", p.Sprintf("the including pattern, the excluding pattern and the time range pattern will task effect on folders"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the download results"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
	}

	return c
}
