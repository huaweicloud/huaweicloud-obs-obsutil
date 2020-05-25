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
	"fmt"
	"obs"
	"os"
	"progress"
	"ratelimit"
	"regexp"
	"strings"
	"sync/atomic"
)

func (c *cpCommand) recoverTasks(fileUrls []string, metadata map[string]string, aclType obs.AclType, storageClassType obs.StorageClassType) error {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "src size", "md5 value", "src --> dst", "cost(ms)", "status code", "success message", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s]", "src size", "src --> dst", "cost(ms)", "status code", "error code", "error message", "request id")
	c.warningLogger.doRecord("[%s, %s, %s]", "src size", "src --> dst", "warn message")

	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

	barCh := newSingleBarChan()
	limiter := c.createRateLimiter()

	var mode cpMode = -1
	var totalAmounts int64
	var totalAmountsForProgress int64
	var totalTasks int64

	cloudUrlPattern := regexp.MustCompile("obs://(.+?)/(.+)")
	action := func(groups []string) {
		if len(groups) != 3 {
			return
		}
		if _mode, taskCtx := c.analyseTask(groups[1], groups[2], cloudUrlPattern); _mode != -1 {
			if mode == -1 {
				mode = _mode
				if mode == cm && !c.crr {
					if config["showBytesForCopy"] == c_true {
						barCh.SetBytes(true)
						barCh.SetTemplate(progress.TpsAndSpeed2)
					} else {
						barCh.SetTemplate(progress.TpsOnly)
					}
				} else {
					barCh.SetBytes(true)
					barCh.SetTemplate(progress.TpsAndSpeed)
				}
				if c.force {
					barCh.Start()
				}
			}
			c.recoverTask(_mode, taskCtx, metadata, aclType, storageClassType,
				pool, barCh, limiter, &totalAmounts, &totalAmountsForProgress, &totalTasks)
		}
	}

	fileUrlPattern := regexp.MustCompile("Z .*?, (.+?) --> (.+?),")
	for _, fileUrl := range fileUrls {
		if atomic.LoadInt32(&c.abort) == 1 {
			doLog(LEVEL_ERROR, "Abort to scan due to unexpected 4xx error, please check the failed manifest files to locate the root cause")
			break
		}
		assist.FindMatches(fileUrl, fileUrlPattern, action, &c.abort)
	}

	if mode == cm && !c.crr && config["showBytesForCopy"] != c_true {
		doLog(LEVEL_INFO, "Number of tasks to recover [%d]", totalTasks)
		barCh.SetTotalCount(totalAmounts)
		progress.SetTotalStream(-1)
	} else {
		doLog(LEVEL_INFO, "Number of tasks to recover [%d], total amounts to recover [%d(B)]", totalTasks, totalAmounts)
		progress.SetTotalStream(totalAmounts)
		barCh.SetTotalCount(totalAmountsForProgress)
	}
	if !c.force {
		barCh.Start()
	}
	progress.SetTotalCount(totalTasks)
	pool.ShutDown()
	barCh.WaitToFinished()
	c.recordEndWithMetricsV2(start, totalTasks, progress.GetSucceedStream(), progress.GetTotalStream())

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}
	return nil
}

func (c *cpCommand) analyseTask(url1, url2 string, cloudUrlPattern *regexp.Regexp) (mode cpMode, taskCtx map[string]string) {
	mode = -1
	if !strings.HasPrefix(url1, "obs://") {
		if groups := cloudUrlPattern.FindStringSubmatch(url2); len(groups) == 3 {
			mode = um
			taskCtx = make(map[string]string, 3+2)
			taskCtx["fileUrl"] = url1
			taskCtx["bucket"] = groups[1]
			taskCtx["key"] = groups[2]
		}
	} else if strings.HasPrefix(url2, "obs://") {
		if groups1 := cloudUrlPattern.FindStringSubmatch(url1); len(groups1) == 3 {
			if groups2 := cloudUrlPattern.FindStringSubmatch(url2); len(groups2) == 3 {
				mode = cm
				taskCtx = make(map[string]string, 5+2)
				taskCtx["dstBucket"] = groups2[1]
				taskCtx["dstKey"] = groups2[2]
				taskCtx["srcBucket"] = groups1[1]
				srcKey := groups1[2]
				if index := strings.LastIndex(srcKey, "?versionId="); index > 0 {
					taskCtx["versionId"] = srcKey[index+11:]
					srcKey = srcKey[:index]
				} else {
					taskCtx["versionId"] = ""
				}
				taskCtx["srcKey"] = srcKey
			}
		}
	} else {
		if groups := cloudUrlPattern.FindStringSubmatch(url1); len(groups) == 3 {
			mode = dm
			taskCtx = make(map[string]string, 3+2)
			taskCtx["fileUrl"] = url2
			taskCtx["bucket"] = groups[1]
			key := groups[2]
			if index := strings.LastIndex(key, "?versionId="); index > 0 {
				taskCtx["versionId"] = key[index+11:]
				key = key[:index]
			} else {
				taskCtx["versionId"] = ""
			}
			taskCtx["key"] = key
		}
	}

	if taskCtx != nil {
		taskCtx["url1"] = url1
		taskCtx["url2"] = url2
	}

	return
}

func (c *cpCommand) checkBucketVersionByTaskCtx(mode cpMode, taskCtx map[string]string) string {
	var bucket string
	if mode == cm {
		bucket = taskCtx["dstBucket"]
	} else {
		bucket = taskCtx["bucket"]
	}
	return c.checkBucketVersion(bucket)
}

func (c *cpCommand) recoverTask(mode cpMode, taskCtx map[string]string, metadata map[string]string,
	aclType obs.AclType, storageClassType obs.StorageClassType, pool concurrent.Pool,
	barCh progress.SingleBarChan, limiter *ratelimit.RateLimiter,
	totalAmounts, totalAmountsForProgress, totalTasks *int64) {

	if c.verifyMd5 {
		c.checkBucketVersionByTaskCtx(mode, taskCtx)
	}

	if mode == um {
		path := taskCtx["fileUrl"]
		bucket := taskCtx["bucket"]
		key := taskCtx["key"]
		info, err := os.Lstat(path)
		if err != nil {
			atomic.AddInt64(totalTasks, 1)
			atomic.AddInt64(totalAmountsForProgress, 1)
			progress.AddFailedCount(1)
			c.failedLogger.doRecord("n/a, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", path, bucket, key, err.Error())
			return
		}

		var fastFailed error
		if checkEmptyFolder(bucket, key, um) {
			fastFailed = fmt.Errorf("Cannot upload to the specified key [%s] in the bucket [%s]", key, bucket)
		}

		if c.link && info.Mode()&os.ModeSymlink == os.ModeSymlink {
			_path, _info, _err := assist.GetRealPath(path)
			if _err != nil {
				atomic.AddInt64(totalTasks, 1)
				atomic.AddInt64(totalAmountsForProgress, 1)
				progress.AddFailedCount(1)
				doLog(LEVEL_ERROR, "Get real path for path [%s] failed, %s", path, err.Error())
				c.failedLogger.doRecord("n/a, %s --> obs://%s/%s, n/a, n/a, n/a, error message [%s], n/a", path, bucket, key, err.Error())
				return
			}
			path = _path
			info = _info
		}
		arcPathPrefix := c.arcDir + "/"

		if info.IsDir() {

			if c.matchFolder {
				if c.matchExclude(path) {
					return
				}

				if !c.matchInclude(path) {
					return
				}
			}

			if !c.matchUploadTimeRange(info) {
				return
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want upload folder [%s] to bucket [%s] ? Please input (y/n) to confirm:", path, bucket)) {
				return
			}
			// modify by w00468571 wanghongbao, if the disableDirObject is true the dir will not upload as a object
			if c.disableDirObject {
				return
			}
			atomic.AddInt64(totalAmountsForProgress, 1)
			atomic.AddInt64(totalTasks, 1)
			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.uploadFolder(bucket, key, arcPathPrefix, path, info, aclType, storageClassType, barCh, 2, fastFailed), 0)
			})
		} else {
			if c.matchExclude(path) {
				return
			}

			if !c.matchInclude(path) {
				return
			}

			if !c.matchUploadTimeRange(info) {
				return
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want upload file [%s] to bucket [%s] ? Please input (y/n) to confirm:", path, bucket)) {
				return
			}
			atomic.AddInt64(totalAmounts, info.Size())
			if info.Size() == 0 {
				atomic.AddInt64(totalAmountsForProgress, 1)
			} else {
				atomic.AddInt64(totalAmountsForProgress, info.Size())
			}
			atomic.AddInt64(totalTasks, 1)

			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.uploadFile(bucket, key, arcPathPrefix, path, info, metadata, aclType, storageClassType, barCh, limiter, 2, fastFailed), info.Size())
			})
		}

	} else if mode == dm {
		bucket := taskCtx["bucket"]
		key := taskCtx["key"]
		versionId := taskCtx["versionId"]
		var _versionId string
		if versionId != "" {
			_versionId = "?versionId=" + versionId
		}
		fileUrl := taskCtx["fileUrl"]

		metaContext, err := getObjectMetadata(bucket, key, versionId, c.payer)
		if err != nil {
			status, code, message, requestId := getErrorInfo(err)
			atomic.AddInt64(totalTasks, 1)
			atomic.AddInt64(totalAmountsForProgress, 1)
			progress.AddFailedCount(1)
			c.failedLogger.doRecord("n/a, obs://%s/%s%s --> %s, n/a, status [%d], error code [%s], error message [%s], request id [%s]", bucket, key, _versionId, fileUrl,
				status, code, message, requestId)
			return
		}

		var fastFailed error
		if checkEmptyFolder(bucket, key, dm) {
			fastFailed = fmt.Errorf("Cannot download the specified key [%s] in the bucket [%s]", key, bucket)
		}

		fileStat, statErr := os.Stat(fileUrl)
		if statErr != nil {
			doLog(LEVEL_WARN, "Stat file failed, %s", statErr.Error())
		}
		if isObsFolder(key) {

			if c.matchFolder {
				if c.matchExclude(key) {
					return
				}

				if !c.matchInclude(key) {
					return
				}

				if !c.matchLastModifiedTime(metaContext.LastModified) {
					return
				}
			}

			atomic.AddInt64(totalAmountsForProgress, 1)
			atomic.AddInt64(totalTasks, 1)
			metaContext.Size = 0

			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResultTransAction(c.downloadFileWithMetaContext(bucket, key, versionId, metaContext, nil, fileUrl, fileStat, barCh, limiter, 2, fastFailed),
					0, 0)
			})
		} else {
			if c.matchExclude(key) {
				return
			}

			if !c.matchInclude(key) {
				return
			}

			if !c.matchLastModifiedTime(metaContext.LastModified) {
				return
			}

			if !c.force && !confirm(fmt.Sprintf("Do you want download key [%s] in the bucket [%s] to [%s] ? Please input (y/n) to confirm:", key, bucket, fileUrl)) {
				return
			}

			atomic.AddInt64(totalAmounts, metaContext.Size)
			atomic.AddInt64(totalTasks, 1)
			if metaContext.Size == 0 {
				atomic.AddInt64(totalAmountsForProgress, 1)
			} else {
				atomic.AddInt64(totalAmountsForProgress, metaContext.Size)
			}

			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.downloadFileWithMetaContext(bucket, key, versionId, metaContext, nil, fileUrl, fileStat, barCh, limiter, 2, fastFailed), metaContext.Size)
			})
		}

	} else if mode == cm {
		srcBucket := taskCtx["srcBucket"]
		srcKey := taskCtx["srcKey"]
		versionId := taskCtx["versionId"]
		dstBucket := taskCtx["dstBucket"]
		dstKey := taskCtx["dstKey"]
		var _versionId string
		if versionId != "" {
			_versionId = "?versionId=" + versionId
		}

		var fastFailed error
		if checkEmptyFolder(dstBucket, dstKey, cm) {
			fastFailed = fmt.Errorf("Cannot copy to the specified key [%s] in the bucket [%s]", dstKey, dstBucket)
		} else if checkEmptyFolder(srcBucket, srcKey, cm) {
			fastFailed = fmt.Errorf("Cannot copy the specified key [%s] in the bucket [%s]", srcKey, srcBucket)
		}

		var srcMetaContext *MetaContext
		var err error
		if c.crr {
			srcMetaContext, err = c.getObjectMetadataCrr(srcBucket, srcKey, versionId, c.payer)
		} else {
			srcMetaContext, err = getObjectMetadata(srcBucket, srcKey, versionId, c.payer)
		}
		if err != nil {
			atomic.AddInt64(totalTasks, 1)
			atomic.AddInt64(totalAmountsForProgress, 1)
			progress.AddFailedCount(1)
			status, code, message, requestId := getErrorInfo(err)
			c.failedLogger.doRecord("n/a, obs://%s/%s%s --> obs://%s/%s, n/a, status [%d], error code [%s], error message [%s], request id [%s]",
				srcBucket, srcKey, _versionId, dstBucket, dstKey, status, code, message, requestId)
			return
		}

		if !isObsFolder(srcKey) || c.matchFolder {
			if c.matchExclude(srcKey) {
				return
			}

			if !c.matchInclude(srcKey) {
				return
			}

			if !c.matchLastModifiedTime(srcMetaContext.LastModified) {
				return
			}

		}

		if !c.force && !confirm(fmt.Sprintf("Do you want copy object [%s] in the bucket [%s] to object [%s] in the bucket [%s] ? Please input (y/n) to confirm:",
			srcKey, srcBucket, dstKey, dstBucket)) {
			return
		}

		count := c.caculateCount(srcMetaContext.Size, false)
		if c.crr || config["showBytesForCopy"] == c_true {
			atomic.AddInt64(totalAmounts, srcMetaContext.Size)
		} else {
			atomic.AddInt64(totalAmounts, count)
		}
		if srcMetaContext.Size == 0 {
			atomic.AddInt64(totalAmountsForProgress, 1)
		} else {
			atomic.AddInt64(totalAmountsForProgress, srcMetaContext.Size)
		}

		atomic.AddInt64(totalTasks, 1)

		if c.crr {
			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.copyObjectCrrWithMetaContext(srcBucket, srcKey, versionId, srcMetaContext, nil, dstBucket, dstKey,
					metadata, aclType, storageClassType, barCh, limiter, 2, fastFailed), srcMetaContext.Size)
			})
		} else {
			pool.ExecuteFunc(func() interface{} {
				return c.handleExecResult(c.copyObjectWithMetaContext(srcBucket, srcKey, versionId, srcMetaContext, nil, dstBucket, dstKey,
					metadata, aclType, storageClassType, barCh, 2, count, fastFailed), 0)
			})
		}
	}
}
