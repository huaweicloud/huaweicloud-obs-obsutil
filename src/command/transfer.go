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
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"obs"
	"os"
	"progress"
	"ratelimit"
	"strings"
	"sync"
	"sync/atomic"
)

type cpMode int

const (
	um cpMode = 0
	dm cpMode = 1
	cm cpMode = 2
)

const (
	fs int = 1
	fl int = 2
)

type parallelContextCommand struct {
	recursiveCommand
	parallel            int
	bigfileThresholdStr string
	partSizeStr         string
	checkpointDir       string
	tempFileDir         string
	arcDir              string
	verifyLength        bool
	verifyMd5           bool
	dryRun              bool // test mode, not real exec
	flat                bool
	update              bool

	//need to be reset in init func
	partSize         int64
	bigfileThreshold int64
	warn             atomic.Value
	warnFlag         int32
}

type transferCommand struct {
	parallelContextCommand
	acl       string
	sc        string
	metadata  string
	drange    string
	versionId string
	rec       string
	link      bool
	crr       bool
	//need to be reset in init func
	folderMap     map[string]string
	folderMapLock *sync.Mutex
}

type ObjectInfo struct {
	XMLName      xml.Name `xml:"ObjectInfo"`
	LastModified int64    `xml:"LastModified"`
	Size         int64    `xml:"Size"`
	ETag         string   `xml:"ETag"`
}

type PartEtag struct {
	PartNumber int    `xml:"PartNumber"`
	ETag       string `xml:"ETag"`
}

func (c *transferCommand) init() {
	c.recursiveCommand.init()
	c.scanContext.init()
	c.partSize = 0
	c.bigfileThreshold = 0
	c.warn = atomic.Value{}
	c.warnFlag = 0
	c.folderMap = make(map[string]string, 10)
	c.folderMapLock = new(sync.Mutex)
}

func (c *transferCommand) defineBasic() {
	c.flagSet.BoolVar(&c.dryRun, "dryRun", false, "")
	c.flagSet.BoolVar(&c.verifyLength, "vlength", false, "")
	c.flagSet.BoolVar(&c.verifyMd5, "vmd5", false, "")
	c.flagSet.BoolVar(&c.link, "link", false, "")
	c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
	c.flagSet.BoolVar(&c.crr, "crr", false, "")
	c.flagSet.BoolVar(&c.at, "at", false, "")
	c.flagSet.BoolVar(&c.disableDirObject, "disableDirObject", false, "") //disableDirObject
	c.flagSet.BoolVar(&c.matchFolder, "mf", false, "")
	c.flagSet.IntVar(&c.jobs, "j", 0, "")
	c.flagSet.IntVar(&c.parallel, "p", 0, "")
	c.flagSet.StringVar(&c.bigfileThresholdStr, "threshold", "", "")
	c.flagSet.StringVar(&c.partSizeStr, "ps", "", "")
	c.flagSet.StringVar(&c.acl, "acl", "", "")
	c.flagSet.StringVar(&c.sc, "sc", "", "")
	c.flagSet.StringVar(&c.checkpointDir, "cpd", "", "")
	c.flagSet.StringVar(&c.metadata, "meta", "", "")
	c.flagSet.StringVar(&c.outDir, "o", "", "")
	c.flagSet.StringVar(&c.include, "include", "", "")
	c.flagSet.StringVar(&c.exclude, "exclude", "", "")
	c.flagSet.StringVar(&c.timeRange, "timeRange", "", "")
	c.flagSet.StringVar(&c.arcDir, "arcDir", "", "")
	c.flagSet.StringVar(&c.tempFileDir, "tempFileDir", "", "")
	c.flagSet.StringVar(&c.payer, "payer", "", "")
}

func (c *transferCommand) needCheckNestedUrl() bool {
	return !c.crr || config["endpointCrr"] == config["endpoint"]
}

func (c *transferCommand) ensureBucketsAndStartAction(buckets []string, action func() error, recordCost bool) error {
	if buckets != nil {
		bucketMap := make(map[string]bool, len(buckets))
		for _, bucket := range buckets {
			if _, ok := bucketMap[bucket]; ok {
				continue
			}
			if err := c.ensureBucket(bucket); err != nil {
				printError(err)
				doLog(LEVEL_ERROR, err.Error())
				return assist.ErrCheckBucketStatus
			}
			bucketMap[bucket] = true
		}
	}
	return c.ensureOuputAndStartLogger(action, recordCost)
}

func (c *transferCommand) getRelativeFolder(folder string) string {
	relativeFolder := strings.Replace(folder, "\\", "/", -1)

	if isObsFolder(relativeFolder) {
		relativeFolder = relativeFolder[:len(relativeFolder)-1]
	}
	if index := strings.LastIndex(relativeFolder, "/"); index >= 0 {
		relativeFolder = relativeFolder[index+1:] + "/"
	} else {
		// handler root folder on windows
		relativeFolder = ""
	}
	return relativeFolder
}

func (c *transferCommand) compareLocation(srcBucket, dstBucket string) bool {
	if output, err := obsClient.GetBucketLocation(srcBucket, obs.WithReqPaymentHeader(c.payer)); err != nil {
		doLog(LEVEL_WARN, "Cannot get the location of bucket [%s], skip to compare location, %s", srcBucket, err.Error())
		return true
	} else if output2, err := obsClient.GetBucketLocation(dstBucket, obs.WithReqPaymentHeader(c.payer)); err != nil {
		doLog(LEVEL_WARN, "Cannot get the location of bucket [%s], skip to compare location, %s", dstBucket, err.Error())
		return true
	} else {
		ret := output.Location == output2.Location
		if !ret {
			printf("Check the locations of two buckets failed, the location of source bucket [%s] is [%s], the location of dest bucket [%s] is [%s]",
				srcBucket, output.Location, dstBucket, output2.Location)
		}
		return ret
	}
}

func (c *transferCommand) ensureKeyForFile(bucket, key string, fileStat os.FileInfo) (bool, error) {
	var changed bool
	output, err := getObjectMetadata(bucket, key, "", c.payer)
	if err == nil {
		changed = fileStat.Size() != output.Size || fileStat.ModTime().After(output.LastModified)
	} else if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode >= 300 && obsError.StatusCode < 500 && obsError.StatusCode != 408 {
		changed = true
	} else {
		changed = false
	}
	return changed, err
}

func (c *transferCommand) ensureKeyForFolder(bucket, key string) (bool, error) {
	var changed bool
	_, err := getObjectMetadata(bucket, key, "", c.payer)
	if err == nil {
		changed = false
	} else if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode >= 300 && obsError.StatusCode < 500 && obsError.StatusCode != 408 {
		changed = true
	} else {
		changed = false
	}
	return changed, err
}

func (c *transferCommand) abortMultipartUpload(bucket, key, uploadId string) (bool, error) {
	input := &obs.AbortMultipartUploadInput{}
	input.Bucket = bucket
	input.Key = key
	input.UploadId = uploadId
	input.RequestPayer = c.payer
	output, err := obsClient.AbortMultipartUpload(input)
	if err == nil {
		doLog(LEVEL_INFO, "Abort multipart upload [%s] in the bucket [%s] with upload id [%s] successfully, request id [%s]", key, bucket, uploadId, output.RequestId)
		return true, nil
	}

	if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode >= 400 && obsError.StatusCode < 500 &&
		strings.Index(config["abortHttpStatusForResumableTasks"], assist.IntToString(obsError.StatusCode)) >= 0 {
		doLogError(obsError, LEVEL_WARN, fmt.Sprintf("Abort multipart upload [%s] in the bucket [%s] with upload id [%s] failed", key, bucket, uploadId))
		return true, obsError
	}
	doLogError(err, LEVEL_WARN, fmt.Sprintf("Abort multipart upload [%s] in the bucket [%s] with upload id [%s] failed", key, bucket, uploadId))
	return false, err
}

func (c *transferCommand) deleteObject(bucket, key, versionId string) (string, error) {
	input := &obs.DeleteObjectInput{}
	input.Bucket = bucket
	input.Key = key
	input.VersionId = versionId
	input.RequestPayer = c.payer
	output, err := obsClient.DeleteObject(input)
	if err == nil {
		return output.RequestId, nil
	}

	return "", err
}

func (c *transferCommand) setObjectMd5(bucket, key, versionId, md5Value string, metadata map[string]string) (string, error) {
	input := &obs.SetObjectMetadataInput{}
	input.Bucket = bucket
	input.Key = key
	input.VersionId = versionId
	input.MetadataDirective = obs.ReplaceMetadataNew

	var inputMetadata map[string]string
	if metadata == nil {
		inputMetadata = make(map[string]string, 1)
	} else {
		inputMetadata = make(map[string]string, len(metadata))
		for k, v := range metadata {
			inputMetadata[k] = v
		}
	}

	inputMetadata[checkSumKey] = md5Value
	input.Metadata = inputMetadata
	input.RequestPayer = c.payer
	output, err := obsClient.SetObjectMetadata(input)
	if err == nil {
		return output.RequestId, nil
	}

	return "", err
}

func (c *transferCommand) caculateCount(size int64, isDownload bool) int64 {
	if size >= c.bigfileThreshold || (!isDownload && size >= serverBigFileThreshold) {
		partSize := c.partSize
		count := size / partSize
		if count >= 10000 {
			partSize = size / 10000
			if size%10000 != 0 {
				partSize++
			}
			count = size / partSize
		}

		if size%partSize != 0 {
			count++
		}
		return count
	}
	return 1
}

func (c *transferCommand) prepareUrls(args []string) (url1 string, url2 string, mode cpMode, err error) {
	mode = -1
	_url1 := args[0]
	_url2 := args[1]

	if !assist.IsObsFilePath(_url1) {
		if !assist.IsObsFilePath(_url2) {
			err = fmt.Errorf("Missing a valid cloud_url")
			return
		}

		if len(_url2[assist.OBS_PREFIX_LEN:]) == 0 {
			err = fmt.Errorf("cloud_url [%s] is not well-formed", _url2)
			return
		}

		mode = um
	} else {
		if len(_url1[assist.OBS_PREFIX_LEN:]) == 0 {
			err = fmt.Errorf("cloud_url [%s] is not well-formed", _url1)
			return
		}

		if assist.IsObsFilePath(_url2) {
			if len(_url2[assist.OBS_PREFIX_LEN:]) == 0 {
				err = fmt.Errorf("cloud_url [%s] is not well-formed", _url2)
				return
			}
			mode = cm
		} else {
			mode = dm
		}
	}

	if _err := c.checkArgs(args[2:]); _err != nil {
		mode = -1
		err = _err
		return
	}

	url1 = _url1
	url2 = _url2
	return
}

func (c *parallelContextCommand) recordCheckpointFile(checkpointFile string, value interface{}) error {
	ret, err := assist.TransToXml(value)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(checkpointFile, ret, 0666)
}

func (c *parallelContextCommand) loadCheckpoint(checkpointFile string, result interface{}) error {
	ret, err := ioutil.ReadFile(checkpointFile)
	if err != nil {
		return err
	}
	return assist.ParseXml(ret, result)
}

func (c *parallelContextCommand) getCheckpointFile(bucket, key, versionId string, mode cpMode) string {
	var prefix string
	if mode == dm {
		prefix = "download"
	} else if mode == cm {
		prefix = "copy"
	} else {
		prefix = "upload"
	}
	if versionId == "" {
		return fmt.Sprintf("%s/%s/%s.txt", c.checkpointDir, prefix, assist.HexMd5(assist.StringToBytes(fmt.Sprintf("%s_%s", bucket, key))))
	}
	return fmt.Sprintf("%s/%s/%s.txt", c.checkpointDir, prefix, assist.HexMd5(assist.StringToBytes(fmt.Sprintf("%s_%s_%s", bucket, key, versionId))))
}

func (c *parallelContextCommand) handleExecResult(ret interface{}, succeedStream int64) interface{} {
	return c.handleExecResultTransAction(ret, succeedStream, 1)
}

func (c *parallelContextCommand) handleExecResultTransAction(ret interface{}, succeedStream int64, transAction int64) interface{} {
	progress.AddFinishedCount(1)
	if transAction > 0 {
		if ret == 2 {
			progress.AddSucceedCount(1)
			progress.AddResumeCount(1)
		} else if ret == 1 {
			progress.AddSucceedCount(1)
			if succeedStream > 0 {
				progress.AddSucceedStream(succeedStream)
			}
			progress.AddTransaction(transAction)
		} else {
			progress.AddFailedCount(1)
			progress.AddTransaction(transAction)
		}
	} else {
		if ret == 2 {
			progress.AddSucceedCount(1)
			progress.AddResumeCount(1)
		} else if ret == 1 {
			progress.AddSucceedCount(1)
		} else {
			progress.AddFailedCount(1)
		}
	}
	return ret
}

func (c *parallelContextCommand) createRateLimiter() *ratelimit.RateLimiter {
	if rateLimitThreshold, err := assist.TranslateToInt64(config["rateLimitThreshold"]); err == nil && rateLimitThreshold > 0 && rateLimitThreshold >= 10*kb {
		return ratelimit.NewRateLimiter(rateLimitThreshold*defaultBrustRate, rateLimitThreshold)
	}
	return nil
}

func (c *parallelContextCommand) recordWarnMessage(warnMessage string, warnLoggerMessage string) {
	if warnMessage == "" && warnLoggerMessage == "" {
		return
	}
	doLog(LEVEL_WARN, warnMessage)
	progress.AddWarningCount(1)
	if atomic.CompareAndSwapInt32(&c.warnFlag, 0, 1) {
		c.warn.Store(errors.New(warnMessage))
	}
	c.warningLogger.doRecord(warnLoggerMessage)
}

func (c *parallelContextCommand) ensureTempFileDirectory() error {
	c.tempFileDir = strings.TrimSpace(c.tempFileDir)
	if c.tempFileDir == "" {
		tempFileDir, err := getTempFileDirectory()
		if err != nil {
			return err
		}
		c.tempFileDir = tempFileDir
	}

	if c.tempFileDir != "" {
		stat, err := os.Stat(c.tempFileDir)
		if err == nil && !stat.IsDir() {
			return fmt.Errorf("The specified temp file folder [%s] is a file, not a folder", c.tempFileDir)
		}

		if err = assist.MkdirAll(c.tempFileDir, 0750); err != nil {
			return err
		}
	}

	return nil
}

func (c *parallelContextCommand) ensureCheckpointDirectory() error {
	c.checkpointDir = strings.TrimSpace(c.checkpointDir)
	if c.checkpointDir == "" {
		checkpointDirectory, err := getCheckpointDirectory()
		if err != nil {
			return err
		}
		c.checkpointDir = checkpointDirectory
	}

	stat, err := os.Stat(c.checkpointDir)
	if err == nil && !stat.IsDir() {
		return fmt.Errorf("The specified checkpoint folder [%s] is a file, not a folder", c.checkpointDir)
	}

	if err = assist.MkdirAll(c.checkpointDir+"/upload", 0750); err != nil {
		return err
	}

	if err = assist.MkdirAll(c.checkpointDir+"/download", 0750); err != nil {
		return err
	}
	if err = assist.MkdirAll(c.checkpointDir+"/copy", 0750); err != nil {
		return err
	}

	return nil
}

func (c *parallelContextCommand) prepareOptions() bool {
	if c.parallel <= 0 {
		c.parallel = assist.MaxInt(assist.StringToInt(config["defaultParallels"], defaultParallels), 1)
	}

	if c.partSizeStr == c_auto {
		c.partSize = -1
	} else {
		if c.partSizeStr != "" {
			if partSize, err := assist.TranslateToInt64(c.partSizeStr); err == nil {
				c.partSize = partSize
			} else {
				printf("Error: The part size [%s] is not well-formed", c.partSizeStr)
				return false
			}
		}

		if c.partSize <= 0 {
			if config["defaultPartSize"] == c_auto {
				c.partSize = -1
			} else if partSize, err := assist.TranslateToInt64(config["defaultPartSize"]); err == nil {
				if partSize <= 0 {
					printf("Error: The default part size [%s] is in config file is less than 0", config["defaultPartSize"])
					return false
				}
				c.partSize = partSize
			} else {
				doLog(LEVEL_INFO, "The default part size [%s] in config file is not well-formed, will use [%d] instead", config["defaultPartSize"], defaultPartSize)
				c.partSize = defaultPartSize
			}
		}
	}

	if c.partSize > serverBigFileThreshold {
		c.partSize = serverBigFileThreshold
	}

	if c.bigfileThresholdStr != "" {
		if bigfileThreshold, err := assist.TranslateToInt64(c.bigfileThresholdStr); err == nil {
			c.bigfileThreshold = bigfileThreshold
		} else {
			printf("Error: The threshold [%s] is not well-formed", c.bigfileThresholdStr)
			return false
		}
	}

	if c.bigfileThreshold <= 0 {
		if bigfileThreshold, err := assist.TranslateToInt64(config["defaultBigfileThreshold"]); err == nil {
			c.bigfileThreshold = bigfileThreshold
		} else {
			doLog(LEVEL_INFO, "The default threshold [%s] in config file is not well-formed,  will use [%d] instead", config["defaultBigfileThreshold"], defaultBigfileThreshold)
			c.bigfileThreshold = defaultBigfileThreshold
		}
	}

	if c.jobs <= 0 {
		c.jobs = assist.MaxInt(assist.StringToInt(config["defaultJobs"], defaultJobs), 1)
	}

	if succeed := c.checkInclude(); !succeed {
		return false
	}

	if succeed := c.checkExclude(); !succeed {
		return false
	}

	if succeed := c.checkTimeRange(); !succeed {
		return false
	}

	if c.arcDir != "" {
		c.arcDir = assist.NormalizeFilePath(strings.TrimSpace(c.arcDir))
		if err := assist.EnsureDirectory(c.arcDir); err != nil {
			printf("Error: The archive dir is not valid, %s", err.Error())
			return false
		}
	}

	if err := c.ensureCheckpointDirectory(); err != nil {
		printError(err)
		return false
	}

	if err := c.ensureTempFileDirectory(); err != nil {
		printError(err)
		return false
	}

	return true
}

func (c *parallelContextCommand) printParams(printOutDir bool, printVerifyMd5 bool, printArcDir bool, printTempFileDir bool) {
	printf("")
	printf("%-15s%-20d%-15s%-20d", "Parallel:", c.parallel, "Jobs:", c.jobs)
	partSizeStr := c_auto
	if c.partSize > 0 {
		partSizeStr = assist.NormalizeBytes(c.partSize)
	}
	printf("%-15s%-20s%-15s%-20s", "Threshold:", assist.NormalizeBytes(c.bigfileThreshold), "PartSize:", partSizeStr)
	if printVerifyMd5 {
		printf("%-15s%-20t%-15s%-20t", "VerifyLength:", c.verifyLength, "VerifyMd5:", c.verifyMd5)
	}

	if c.exclude != "" {
		printf("%-15s%-30s", "Exclude:", c.exclude)
	}

	if c.exclude != "" {
		printf("%-15s%-30s", "Include:", c.include)
	}

	if c.timeRange != "" {
		printf("%-15s%-30s", "TimeRange:", c.timeRange)
	}

	printf("%s%-30s", "CheckpointDir: ", assist.NormalizeFilePath(c.checkpointDir))
	if printOutDir {
		printf("%s%-30s", "OutputDir: ", assist.NormalizeFilePath(c.outDir))
	}

	if printTempFileDir {
		printf("%s%-30s", "TempFileDir: ", assist.NormalizeFilePath(c.tempFileDir))
	}

	if printArcDir && c.arcDir != "" {
		printf("%s%-30s", "ArcDir: ", assist.NormalizeFilePath(c.arcDir))
	}
	printf("")
}

func (c *parallelContextCommand) ensureOuputAndStartLogger(action func() error, recordCost bool) error {
	if err := c.ensureOutputDirectory(); err != nil {
		printError(err)
		return assist.ErrInitializing
	}
	if err := c.startLogger(true); err != nil {
		printError(err)
		return assist.ErrInitializing
	}
	defer c.endLogger()
	var ret error
	if action != nil {
		if recordCost {
			start := assist.GetUtcNow()
			ret = action()
			c.recordEnd(start)
		} else {
			ret = action()
		}
	}

	return ret
}

func (c *parallelContextCommand) autoSelectPartSize(totalSize int64, mode cpMode) int64 {
	if c.partSize > 0 {
		return c.partSize
	}

	if mode == dm {
		if totalSize < 1*gb {
			return 50 * mb
		}

		return 100 * mb
	}

	if totalSize < 100*mb {
		return 9 * mb
	}

	if totalSize < 500*mb {
		return 20 * mb
	}

	if totalSize < 1*gb {
		return 30 * mb
	}

	if totalSize < 5*gb {
		return 40 * mb
	}

	return 50 * mb
}

func (c *transferCommand) checkParams() (aclType obs.AclType, storageType obs.StorageClassType, metadata map[string]string, flag bool) {
	if c.acl != "" {
		_aclType, ok := objectAclType[c.acl]
		if !ok {
			printf("Error: Invalid acl [%s], possible values are:[private|public-read|public-read-write|bucket-owner-full-control]", c.acl)
			return
		}
		aclType = _aclType
	}

	if c.sc != "" {
		_storageClassType, ok := storageClassType[c.sc]
		if !ok {
			printf("Error: Invalid sc [%s], possible values are:[standard|warm|cold]", c.sc)
			return
		}
		storageType = _storageClassType
	}

	if c.metadata != "" {
		metadataList := strings.Split(c.metadata, "#")
		metadata = make(map[string]string, len(metadataList))
		for _, meta := range metadataList {
			metaPair := strings.Split(meta, ":")
			length := len(metaPair)
			if length == 0 || length > 2 {
				printf("Error: Invalid metadata [%s]", c.metadata)
				return
			}
			key := metaPair[0]
			if key == "" {
				printf("Error: Invalid metadata [%s]", c.metadata)
				return
			}

			val := ""
			if length == 2 {
				val = metaPair[1]
			}
			metadata[key] = val
		}
	}

	flag = true
	return
}

func (c *transferCommand) ensureParentFolder(bucket, key string) {
	ignoreLast := true
	if isObsFolder(key) {
		key = key[:len(key)-1]
		ignoreLast = false
	}

	folders := strings.Split(key, "/")
	length := len(folders)
	var current string
	for index, folder := range folders {
		if index == length-1 && ignoreLast {
			break
		}
		current += folder + "/"
		input := &obs.PutObjectInput{}
		input.Bucket = bucket
		input.Key = current
		input.ContentLength = 0
		input.RequestPayer = c.payer
		if _, err := obsClient.PutObject(input); err != nil {
			doLogError(err, LEVEL_WARN, "Cannot create folder "+current)
		}
	}
}

func (c *transferCommand) ensureObjectAttributes(bucket, key, versionId string, srcMetaContext *MetaContext, metadata map[string]string) (map[string]string, string, obs.StorageClassType, string) {
	return c.ensureObjectAttributesByClient(bucket, key, versionId, srcMetaContext, metadata, obsClient)
}

func (c *transferCommand) ensureObjectAttributesByClient(bucket, key, versionId string, srcMetaContext *MetaContext, metadata map[string]string, client *obs.ObsClient) (map[string]string, string, obs.StorageClassType, string) {
	var _metadata map[string]string
	if metadata == nil {
		_metadata = make(map[string]string)
	} else {
		_metadata = make(map[string]string, len(metadata))
		for k, v := range metadata {
			_metadata[k] = v
		}
	}

	if srcMetaContext == nil || srcMetaContext.Metadata == nil {
		srcMetaContext, _ = getObjectMetadataByClient(bucket, key, versionId, client, c.payer)
		if versionId == "" {
			doLog(LEVEL_WARN, "Cannot get object metadata of object [%s] in the bucket [%s]", key, bucket)
		} else {
			doLog(LEVEL_WARN, "Cannot get object metadata of object [%s] with version id [%s] in the bucket [%s]", key, versionId, bucket)
		}
	}

	if srcMetaContext != nil {
		if srcMetaContext.Metadata != nil {
			for k, v := range srcMetaContext.Metadata {
				if _, ok := _metadata[k]; !ok || k == checkSumKey {
					_metadata[k] = v
				}
			}
		}
		return _metadata, srcMetaContext.ContentType, srcMetaContext.StorageClass, srcMetaContext.WebsiteRedirectLocation
	}

	return _metadata, "", "", ""
}
