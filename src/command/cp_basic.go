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
	"os"
	"path/filepath"
	"progress"
	"regexp"
	"strings"
)

type cpCommand struct {
	transferCommand
	multiSourceMode int
}

func (c *cpCommand) doRecover() error {
	if !c.prepareOptions() {
		return assist.ErrInvalidArgs
	}

	aclType, storageClassType, metadata, succeed := c.checkParams()
	if !succeed {
		return assist.ErrInvalidArgs
	}

	if c.crr && !c.createObsClientCrr() {
		return assist.ErrInitializing
	}

	c.outDir = strings.TrimSpace(c.outDir)
	if c.outDir == "" {
		outDir, err := getOutputDirectory()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		c.outDir = outDir
	}

	var fileUrls []string
	_fileUrls, err := assist.FindFilesV2(c.outDir, regexp.MustCompile(fmt.Sprintf("failed_report_.+?_%s.txt", c.rec)))
	if err != nil {
		printError(err)
		return assist.ErrExecuting
	}
	fileUrls = _fileUrls

	if len(fileUrls) <= 0 {
		printf("Error: Cannot find any failed-record file for task id [%s]!", c.rec)
		return assist.ErrTaskNotFound
	}

	c.printStart()

	if err := c.ensureOutputDirectory(); err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	c.printParams(true, true, true, true)
	if err := c.startLogger(true); err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	doLog(LEVEL_INFO, "Recover cp tasks by task id [%s]", c.rec)
	defer c.endLogger()
	return c.recoverTasks(fileUrls, metadata, aclType, storageClassType)
}

func (c *transferCommand) deduplicationFilePath(fileList []string) []string {
	filePathMap := make(map[string]bool, len(fileList))
	afterDeduplicationFilePath := make([]string, 0, len(fileList))
	for _, filePath := range fileList {
		trimmedFilePath := strings.TrimSpace(filePath)
		if trimmedFilePath == "" {
			continue
		}
		fileUrl := assist.NormalizeFilePath(trimmedFilePath)
		before := len(filePathMap)
		filePathMap[fileUrl] = false
		after := len(filePathMap)
		if before != after {
			afterDeduplicationFilePath = append(afterDeduplicationFilePath, fileUrl)
		}
	}
	return afterDeduplicationFilePath
}

func (c *cpCommand) uploadFileList(fileList []string, url2 string) error {
	fileList = c.deduplicationFilePath(fileList)
	if err := assist.PathListNested(fileList); err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}
	aclType, storageClassType, metadata, succeed := c.checkParams()
	if !succeed {
		return assist.ErrInvalidArgs
	}
	bucket, keyOrDir, err := c.splitCloudUrl(url2)
	dir := assist.MaybeAddTrailingSlash(keyOrDir)
	if err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	if checkEmptyFolder(bucket, dir, um) {
		printf("Error: Cannot upload to [%s], the url contains empty folder", url2)
		return assist.ErrInvalidArgs
	}

	if c.verifyMd5 {
		c.checkBucketVersion(bucket)
	}

	c.printStart()
	c.ensureParentFolder(bucket, dir)
	return c.ensureBucketsAndStartAction([]string{bucket}, func() error {
		c.printParams(true, true, true, false)
		return c.uploadMultiFilesOrFolders(bucket, dir, fileList, metadata, aclType, storageClassType)
	}, false)
}

func (c *cpCommand) doUpload(url1, url2 string) error {
	if c.multiSourceMode == fl {
		fileList, err := assist.ReadContentLineByFileUrl(url1)
		if err != nil {
			printf("Error: Cannot upload to [%s], get upload file list from [%s] encountering error, [%s].", url2, url1, err.Error())
			return assist.ErrFileNotFound
		}
		return c.uploadFileList(fileList, url2)
	}

	if c.multiSourceMode == fs {
		fileList := strings.Split(url1, ",")
		return c.uploadFileList(fileList, url2)
	}

	if c.multiSourceMode != 0 {
		printf("Error: Invalid msm [%d], possible values are:[1|2]", c.multiSourceMode)
		return assist.ErrInvalidArgs
	}

	url1, err := filepath.Abs(url1)
	if err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	stat, err := os.Lstat(url1)
	if err != nil {
		printError(err)
		return assist.ErrFileNotFound
	}

	linkFolder := false
	relativeFolder := ""
	if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
		if c.link {
			_url1, _stat, err := assist.GetRealPath(url1)
			if err != nil {
				printError(err)
				return assist.ErrFileNotFound
			}
			if _stat.IsDir() {
				if !c.flat {
					relativeFolder = c.getRelativeFolder(url1)
				}
				linkFolder = true
				c.folderMap[_url1] = assist.NormalizeFilePath(url1) + "/"
			}
			url1 = _url1
			stat = _stat
		} else {
			if _stat, _err := os.Stat(url1); _err == nil && _stat.IsDir() {
				stat = _stat
			}
		}
	}

	aclType, storageClassType, metadata, succeed := c.checkParams()
	if !succeed {
		return assist.ErrInvalidArgs
	}

	bucket, keyOrDir, err := c.splitCloudUrl(url2)
	if err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	if checkEmptyFolder(bucket, keyOrDir, um) {
		printf("Error: Cannot upload to [%s], the url contains empty folder", url2)
		return assist.ErrInvalidArgs
	}

	if c.verifyMd5 {
		c.checkBucketVersion(bucket)
	}

	c.printStart()
	c.ensureParentFolder(bucket, keyOrDir)

	if c.arcDir != "" && !c.force && !confirm(fmt.Sprintf("Do you want archive source path [%s] to archive path [%s] ? Please input (y/n) to confirm:", url1, c.arcDir)) {
		return assist.ErrInterrupted
	}

	if !stat.IsDir() {
		key := keyOrDir

		if key == "" || isObsFolder(key) {
			key += stat.Name()
		}

		arcPath := ""
		if c.arcDir != "" {
			arcPath = c.arcDir + "/" + stat.Name()
		}

		if c.forceRecord {
			return c.ensureBucketsAndStartAction([]string{bucket}, func() error {
				c.printParams(true, true, true, false)
				c.recordStartFuncForUpload()
				ret := c.uploadFile(bucket, key, arcPath, url1, stat, metadata, aclType, storageClassType, nil, nil, 1, nil)
				if ret >= 1 {
					progress.AddSucceedCount(1)
					return nil
				}
				progress.AddFailedCount(1)
				return assist.ErrExecuting
			}, true)
		}
		c.printParams(false, true, true, false)
		ret := c.uploadFile(bucket, key, arcPath, url1, stat, metadata, aclType, storageClassType, nil, nil, 0, nil)
		if warn, ok := c.warn.Load().(error); ok {
			printWarn(warn)
		}
		if ret == 0 {
			return assist.ErrExecuting
		}
		return nil
	}
	dir := keyOrDir
	if c.recursive {
		if c.arcDir == assist.NormalizeFilePath(url1) {
			printf("Error: The folder to be uploaded and the archive dir are same!")
			return assist.ErrInvalidArgs
		}

		return c.ensureBucketsAndStartAction([]string{bucket}, func() error {
			c.printParams(true, true, true, false)
			doLog(LEVEL_INFO, "Upload objects from local folder [%s] to cloud folder [%s] in the bucket [%s]", url1, dir, bucket)
			arcDir := c.arcDir
			if linkFolder {
				if dir != "" && !isObsFolder(dir) {
					dir = dir + "/"
				}
				dir += relativeFolder

				if arcDir != "" {
					if !isObsFolder(arcDir) {
						arcDir = arcDir + "/"
					}
					arcDir += relativeFolder
				}
			}
			return c.uploadDir(bucket, dir, arcDir, url1, linkFolder, stat, metadata, aclType, storageClassType)
		}, false)
	}
	printf("Error: Must pass -r to upload folder!")
	return assist.ErrInvalidArgs
}

func (c *cpCommand) doDownload(url1, url2 string) error {
	originUrl2 := url2
	url2, err := filepath.Abs(url2)
	if err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}
	bucket, keyOrDir, err := c.splitCloudUrl(url1)
	if err != nil {
		printError(err)
		return assist.ErrInvalidArgs
	}

	if checkEmptyFolder(bucket, keyOrDir, dm) {
		printf("Error: Cannot download from [%s], the url contains empty folder", url1)
		return assist.ErrInvalidArgs
	}

	c.printStart()

	stat, err := os.Lstat(url2)
	if !c.recursive {
		key := keyOrDir
		if key == "" {
			printf("Error: The object key is empty!")
			return assist.ErrInvalidArgs
		}

		if isObsFolder(key) {
			if !c.force && !confirm(fmt.Sprintf("Do you forget pass \"-r\" to recursively download? \nThis command will only download a empty folder as [%s]. Please input (y/n) to confirm:", url2)) {
				return nil
			}
		}

		if (err != nil && (isObsFolder(originUrl2) || strings.HasSuffix(originUrl2, "\\"))) || (err == nil && stat.IsDir()) {
			fileName := key
			if index := strings.LastIndex(key, "/"); index >= 0 {
				fileName = key[index+1:]
			}
			url2 = assist.NormalizeFilePath(url2 + "/" + fileName)
			stat = nil
		}

		if c.forceRecord {
			return c.ensureBucketsAndStartAction([]string{bucket}, func() error {
				c.printParams(true, true, false, true)
				c.recordStartFuncForDownload()
				ret := c.downloadFile(bucket, key, c.versionId, url2, stat, nil, nil, 1, nil)
				if ret >= 1 {
					progress.AddSucceedCount(1)
					return nil
				}
				progress.AddFailedCount(1)
				return assist.ErrExecuting
			}, true)
		}
		c.printParams(false, true, false, true)
		ret := c.downloadFile(bucket, key, c.versionId, url2, stat, nil, nil, 0, nil)
		if warn, ok := c.warn.Load().(error); ok {
			printWarn(warn)
		}
		if ret == 0 {
			return assist.ErrExecuting
		}
		return nil
	}
	if _err := c.ensureBucket(bucket); _err != nil {
		printError(_err)
		doLog(LEVEL_ERROR, _err.Error())
		return assist.ErrCheckBucketStatus
	}

	if err != nil {
		if err = assist.MkdirAll(url2, os.ModePerm); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}
		stat, err = os.Lstat(url2)
		if err != nil {
			printError(err)
			return assist.ErrFileNotFound
		}

	} else if !stat.IsDir() {
		printf("Error: Cannot download to the folder [%s] due to a file with the same name exits", url2)
		return assist.ErrInvalidArgs
	}

	if _err := c.ensureOutputDirectory(); _err != nil {
		printError(_err)
		return assist.ErrInvalidArgs
	}
	if _err := c.startLogger(true); _err != nil {
		printError(_err)
		return assist.ErrInvalidArgs
	}
	defer c.endLogger()
	dir := keyOrDir
	c.printParams(true, true, false, true)
	doLog(LEVEL_INFO, "Download objects from cloud folder [%s] in the bucket [%s] to local folder [%s] ", dir, bucket, url2)
	return c.downloadDir(bucket, dir, url2, stat)
}

func (c *cpCommand) doCopy(url1, url2 string) error {
	if c.needCheckNestedUrl() && url1 == url2 {
		printf("The source url and the destination url are same!")
		return assist.ErrInvalidArgs
	}

	aclType, storageClassType, metadata, succeed := c.checkParams()
	if !succeed {
		return assist.ErrInvalidArgs
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

	if checkEmptyFolder(srcBucket, srcKeyOrDir, cm) {
		printf("Error: Cannot copy from [%s], the url contains empty folder", url1)
		return assist.ErrInvalidArgs
	}

	if checkEmptyFolder(dstBucket, dstKeyOrDir, cm) {
		printf("Error: Cannot copy to [%s], the url contains empty folder", url2)
		return assist.ErrInvalidArgs
	}

	if c.crr {
		if !c.createObsClientCrr() {
			return assist.ErrInitializing
		}
		if c.verifyMd5 {
			c.checkBucketVersion(dstBucket)
		}
	}

	c.printStart()
	c.ensureParentFolder(dstBucket, dstKeyOrDir)

	if !c.recursive {
		srcKey := srcKeyOrDir
		if srcKey == "" {
			printf("Error: The source object key is empty!")
			return assist.ErrInvalidArgs
		}
		dstKey := dstKeyOrDir
		if dstKey == "" || isObsFolder(dstKey) {
			if index := strings.LastIndex(srcKey, "/"); index >= 0 {
				dstKey += srcKey[index+1:]
			} else {
				dstKey += srcKey
			}
		}

		if c.forceRecord {
			if c.crr {
				return c.ensureBucketsAndStartActionCrr(srcBucket, dstBucket, func() error {
					c.printParams(true, true, false, false)
					c.recordStartFuncForCopy()
					ret := c.copyObjectCrr(srcBucket, srcKey, c.versionId, dstBucket, dstKey, metadata, aclType, storageClassType, nil, nil, 1, nil)
					if ret >= 1 {
						progress.AddSucceedCount(1)
						return nil
					}
					progress.AddFailedCount(1)
					return assist.ErrExecuting
				}, true)
			}
			if succeed := c.compareLocation(srcBucket, dstBucket); !succeed {
				return assist.ErrInvalidArgs
			}

			return c.ensureBucketsAndStartAction([]string{srcBucket, dstBucket}, func() error {
				c.printParams(true, false, false, false)
				c.recordStartFuncForCopy()
				ret := c.copyObject(srcBucket, srcKey, c.versionId, dstBucket, dstKey, metadata, aclType, storageClassType, nil, 1, nil)
				if ret >= 1 {
					progress.AddSucceedCount(1)
					return nil
				}
				progress.AddFailedCount(1)
				return assist.ErrExecuting
			}, true)
		}
		var ret int
		if c.crr {
			c.printParams(false, true, false, false)
			ret = c.copyObjectCrr(srcBucket, srcKey, c.versionId, dstBucket, dstKey, metadata, aclType, storageClassType, nil, nil, 0, nil)
		} else {
			c.printParams(false, false, false, false)
			ret = c.copyObject(srcBucket, srcKey, c.versionId, dstBucket, dstKey, metadata, aclType, storageClassType, nil, 0, nil)
		}
		if warn, ok := c.warn.Load().(error); ok {
			printWarn(warn)
		}
		if ret == 0 {
			return assist.ErrExecuting
		}
		return nil
	}
	if dstBucket == srcBucket {
		_url1 := url1
		_url2 := url2

		if !isObsFolder(_url1) {
			_url1 += "/"
		}

		if !isObsFolder(_url2) {
			_url2 += "/"
		}

		if c.needCheckNestedUrl() {
			if index := strings.Index(_url2, _url1); index >= 0 {
				printf("The source cloud_url and the destination cloud_url are nested")
				return assist.ErrInvalidArgs
			}

			if index := strings.Index(_url1, _url2); index >= 0 {
				printf("The source cloud_url and the destination cloud_url are nested")
				return assist.ErrInvalidArgs
			}
		}
	}

	if c.crr {
		return c.ensureBucketsAndStartActionCrr(srcBucket, dstBucket, func() error {
			srcDir := srcKeyOrDir
			dstDir := dstKeyOrDir

			c.printParams(true, true, false, false)
			doLog(LEVEL_INFO, "Copy objects from cloud folder [%s] in the bucket [%s] to cloud folder [%s] in the bucket [%s] ",
				srcDir, srcBucket, dstDir, dstBucket)
			return c.copyDir(srcBucket, srcDir, dstBucket, dstDir, metadata, aclType, storageClassType)
		}, false)
	}
	if succeed := c.compareLocation(srcBucket, dstBucket); !succeed {
		return assist.ErrInvalidArgs
	}
	return c.ensureBucketsAndStartAction([]string{srcBucket, dstBucket}, func() error {
		srcDir := srcKeyOrDir
		dstDir := dstKeyOrDir

		c.printParams(true, false, false, false)
		doLog(LEVEL_INFO, "Copy objects from cloud folder [%s] in the bucket [%s] to cloud folder [%s] in the bucket [%s] ",
			srcDir, srcBucket, dstDir, dstBucket)
		return c.copyDir(srcBucket, srcDir, dstBucket, dstDir, metadata, aclType, storageClassType)
	}, false)
}

func initCp() command {
	c := &cpCommand{}
	c.key = "cp"
	c.usage = []string{
		"file_url cloud_url [options...]",
		"cloud_url file_url [options...]",
		"cloud_url cloud_url [options...]",
	}
	c.description = "upload, download or copy objects"
	c.define = func() {
		c.init()
		c.defineBasic()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.update, "u", false, "")
		c.flagSet.IntVar(&c.multiSourceMode, "msm", 0, "")
		c.flagSet.BoolVar(&c.flat, "flat", false, "")
		c.flagSet.StringVar(&c.versionId, "versionId", "", "")
		c.flagSet.StringVar(&c.rec, "recover", "", "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) <= 1 {
			if c.rec != "" {
				return c.doRecover()
			}
			c.showHelp()
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		url1, url2, mode, err := c.prepareUrls(args)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		if !c.prepareOptions() {
			return assist.ErrInvalidArgs
		}
		_, succeed := getRequestPayerType(c.payer)
		if !succeed {
			return assist.ErrInvalidArgs
		}

		var ret error
		if mode == um {
			ret = c.doUpload(url1, url2)
		} else if mode == dm {
			ret = c.doDownload(url1, url2)
		} else if mode == cm {
			//TODO
			if !c.crr {
				if c.jobs > 10 {
					printf("Error: The max jobs for copy is 10")
					return assist.ErrInvalidArgs
				}
				if c.parallel > 10 {
					printf("Error: The max parallel for copy is 10")
					return assist.ErrInvalidArgs
				}
			}
			ret = c.doCopy(url1, url2)
		}
		return ret
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("upload, download or copy objects"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil cp file_url|folder_url obs://bucket[/key] [-dryRun] [-link] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-arcDir=xxx] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil cp folder_url obs://bucket[/prefix] -r [-dryRun] [-link] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-arcDir=xxx] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 3:")
		printf("%2s%s", "", "obsutil cp file1_url,folder1_url|filelist_url obs://bucket[/prefix] -msm=1 [-r] [-dryRun] [-link] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-arcDir=xxx] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 4:")
		printf("%2s%s", "", "obsutil cp obs://bucket/key file_url|folder_url [-dryRun] [-tempFileDir=xxx] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 5:")
		printf("%2s%s", "", "obsutil cp obs://bucket[/prefix] folder_url -r [-dryRun] [-tempFileDir=xxx] [-f] [-u] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-at] [-disableDirObject] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 6:")
		printf("%2s%s", "", "obsutil cp obs://srcbucket/key obs://dstbucket/[dest] [-dryRun] [-u] [-crr] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 7:")
		printf("%2s%s", "", "obsutil cp obs://srcbucket[/src_prefix] obs://dstbucket[/dest_prefix] -r [-dryRun] [-f] [-u] [-crr] [-vlength] [-vmd5] [-flat] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 8:")
		printf("%2s%s", "", "obsutil cp -recover=xxx [-dryRun] [-crr] [-f] [-u] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-arcDir=xxx] [-tempFileDir=xxx] [-o=xxx] [-cpd=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-dryRun")
		printf("%4s%s", "", p.Sprintf("conduct a dry run"))
		printf("")
		printf("%2s%s", "", "-tempFileDir=xxx")
		printf("%4s%s", "", p.Sprintf("the temp file dir, used to save temporary files during the objects are downloading"))
		printf("")
		printf("%2s%s", "", "-link")
		printf("%4s%s", "", p.Sprintf("upload the actual path of the symbolic-link file/folder"))
		printf("")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch upload, download or copy objects by prefix (folder)"))
		printf("")
		printf("%2s%s", "", "-recover")
		printf("%4s%s", "", p.Sprintf("recover the failed upload, download or copy tasks by task ID"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while objects are to be uploaded, downloaded or copied"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files"))
		printf("")
		printf("%2s%s", "", "-u")
		printf("%4s%s", "", p.Sprintf("upload, download or copy the changed sources only"))
		printf("")
		printf("%2s%s", "", "-crr")
		printf("%4s%s", "", p.Sprintf("copy the source object(s) through the client-side cross region replication mode"))
		printf("")
		printf("%2s%s", "", "-vlength")
		printf("%4s%s", "", p.Sprintf("verify the size after the objects are uploaded or downloaded or copied through crr"))
		printf("")
		printf("%2s%s", "", "-vmd5")
		printf("%4s%s", "", p.Sprintf("verify the MD5 value after the objects are uploaded or downloaded or copied through crr"))
		printf("")
		printf("%2s%s", "", "-flat")
		printf("%4s%s", "", p.Sprintf("upload, download or copy the sources without the relative parent folder/prefix"))
		printf("")
		printf("%2s%s", "", "-cpd=xxx")
		printf("%4s%s", "", p.Sprintf("the directory where the part records reside, used to record the progress of upload, download or copy jobs"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent upload, download or copy jobs, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-p=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent upload, download or copy tasks (a task is a sub-job), the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-versionId=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID of the object to be copied or downloaded"))
		printf("")
		printf("%2s%s", "", "-ps=auto")
		printf("%4s%s", "", p.Sprintf("the part size of each upload, download or copy task, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-threshold=52428800")
		printf("%4s%s", "", p.Sprintf("the threshold, if it is exceeded, the upload, download or copy job will be divided into multiple tasks by the part size, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-meta=aaa:bbb#ccc:ddd")
		printf("%4s%s", "", p.Sprintf("the customized metadata of each object to be uploaded or copied"))
		printf("")
		printf("%2s%s", "", "-acl=xxx")
		printf("%4s%s", "", p.Sprintf("the ACL of each object to be uploaded or copied, possible values are [private|public-read|public-read-write|bucket-owner-full-control]"))
		printf("")
		printf("%2s%s", "", "-sc=xxx")
		printf("%4s%s", "", p.Sprintf("the storage class of each object to be uploaded or copied, possible values are [standard|warm|cold]"))
		printf("")
		printf("%2s%s", "", "-include=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be included"))
		printf("")
		printf("%2s%s", "", "-exclude=*.xxx")
		printf("%4s%s", "", p.Sprintf("the objects whose names match this pattern will be excluded"))
		printf("")
		printf("%2s%s", "", "-at")
		printf("%4s%s", "", p.Sprintf("the files whose latest access time falls into the time range (-timeRange option) will be uploaded"))
		printf("")
		printf("%2s%s", "", "-disableDirObject")
		printf("%4s%s", "", p.Sprintf("the folder will not be uploaded as a object"))
		printf("")
		printf("%2s%s", "", "-timeRange=time1-time2")
		printf("%4s%s", "", p.Sprintf("the time range for last modified time, between which the objects will be uploaded, downloaded or copied"))
		printf("")
		printf("%2s%s", "", "-mf")
		printf("%4s%s", "", p.Sprintf("the including pattern, the excluding pattern and the time range pattern will task effect on folders"))
		printf("")
		printf("%2s%s", "", "-arcDir=xxx")
		printf("%4s%s", "", p.Sprintf("the archive dir, used to archive the successful uploaded file(s)"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the upload, download or copy results"))
		printf("")
		printf("%2s%s", "", "-msm=1")
		printf("%4s%s", "", p.Sprintf("multi-source upload mode. Possible values are [1|2]. 1 indicates the source URL is a list of file names separated by commas.2 indicates the source URL is a list file"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}
	return c
}
