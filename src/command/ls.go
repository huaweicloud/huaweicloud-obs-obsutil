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
	"sync"
	"sync/atomic"
)

type lsCommand struct {
	cloudUrlCommand
	short           bool
	storageClass    bool
	dir             bool
	multipart       bool
	all             bool
	version         bool
	jobs            int
	limit           int64
	bytesFormat     string
	marker          string
	versionIdMarker string
	uploadIdMarker  string
}

type storageClassResult struct {
	storageClass string
	bucket       string
	err          error
}

func (c *lsCommand) getBucketStorageClass(bucket string) *storageClassResult {
	input := &obs.GetBucketMetadataInput{}
	input.Bucket = bucket
	input.RequestPayer = c.payer
	output, err := obsClient.GetBucketMetadata(input)
	if err == nil {
		return &storageClassResult{bucket: bucket, storageClass: transStorageClassType(output.StorageClass)}
	}
	return &storageClassResult{err: err}
}

func (c *lsCommand) listBuckets() error {
	if c.limit == 0 {
		printf("Bucket Number is: %d", 0)
		return nil
	}

	input := &obs.ListBucketsInput{}
	input.QueryLocation = true
	output, err := obsClient.ListBuckets(input)

	if err == nil {
		count := int64(len(output.Buckets))
		if c.limit > 0 && c.limit < count {
			output.Buckets = output.Buckets[:c.limit]
		}

		if c.short {
			for _, val := range output.Buckets {
				printf("obs://%s", val.Name)
			}
		} else {
			if c.storageClass {
				storageClasses := make(map[string]string, len(output.Buckets))
				if c.jobs <= 0 {
					c.jobs = assist.MaxInt(assist.StringToInt(config["defaultJobs"], defaultJobs), 1)
				}

				poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
				futureChan := make(chan concurrent.Future, poolCacheCount)
				pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)
				wg := new(sync.WaitGroup)
				wg.Add(1)
				go func() {
					for {
						future, ok := <-futureChan
						if !ok {
							break
						}
						result := future.Get().(*storageClassResult)
						if result.err == nil {
							storageClasses[result.bucket] = result.storageClass
						}
					}
					wg.Done()
				}()

				h := &assist.HintV2{}
				h.Message = "Querying the storage classes of buckets"
				h.Start()

				for _, val := range output.Buckets {
					bucket := val.Name
					future, _err := pool.SubmitFunc(func() interface{} {
						return c.getBucketStorageClass(bucket)
					})
					if _err != nil {
						doLog(LEVEL_ERROR, "Submit task to pool failed bucket [%s], %s", bucket, _err.Error())
					} else if future != nil {
						futureChan <- future
					}
				}

				close(futureChan)
				pool.ShutDown()
				wg.Wait()
				h.End()

				format := "%-10s%-15s%-25s%-15s%-15s"
				printf(format, "Bucket", "StorageClass", "CreationDate", "Location", "BucketType")
				for _, val := range output.Buckets {
					bucketName := "obs://" + val.Name
					printf("%s", bucketName)
					bucketType := c_object
					if val.BucketType != "" {
						bucketType = val.BucketType
					}
					printf(format, "", storageClasses[val.Name], val.CreationDate.Format(ISO8601_DATE_FORMAT), val.Location, bucketType)
					printf("")
				}

			} else {
				format := "%-25s%-25s%-15s%-15s"
				printf(format, "Bucket", "CreationDate", "Location", "BucketType")
				for _, val := range output.Buckets {
					bucketName := "obs://" + val.Name
					bucketType := c_object
					if val.BucketType != "" {
						bucketType = val.BucketType
					}
					if len(bucketName) >= 25 {
						printf("%s", bucketName)
						printf(format, "", val.CreationDate.Format(ISO8601_DATE_FORMAT), val.Location, bucketType)
					} else {
						printf(format, bucketName, val.CreationDate.Format(ISO8601_DATE_FORMAT), val.Location, bucketType)
					}
					printf("")
				}
			}
		}
		printf("Bucket number is: %d", len(output.Buckets))
		return nil
	}
	logError(err, LEVEL_INFO, "List buckets failed")
	return assist.ErrExecuting
}

func (c *lsCommand) listVersions(bucket, prefix string) error {
	input := &obs.ListVersionsInput{}
	input.Bucket = bucket
	input.Prefix = prefix
	input.KeyMarker = c.marker
	input.VersionIdMarker = c.versionIdMarker
	input.RequestPayer = c.payer
	if c.dir {
		input.Delimiter = "/"
		if input.Prefix != "" && !isObsFolder(input.Prefix) {
			input.Prefix += "/"
		}
	}
	limit := c.limit
	count := limit
	if count <= 0 {
		count = 1000
	}
	truncated := false
	totalFolders := make([]string, 0, count)
	totalVersions := make([]obs.Version, 0, count)
	totalDeleteMarkers := make([]obs.DeleteMarker, 0, count)
	nextKeyMarker := ""
	nextVersionIdMarker := ""
	var totalCount int64
	h := &assist.HintV2{}
	h.MessageFunc = func() string {
		count := ""
		if tc := atomic.LoadInt64(&totalCount); tc > 0 {
			count = "[" + assist.Int64ToString(tc) + "]"
		}
		return fmt.Sprintf("Listing versioning objects %s", count)
	}
	h.Start()
	var hasListError error
	for {
		if limit > 0 {
			if limit <= 1000 {
				input.MaxKeys = int(limit)
				truncated = true
			} else {
				input.MaxKeys = 1000
				limit -= 1000
			}
		}

		output, err := obsClient.ListVersions(input)
		if err != nil {
			hasListError = err
			break
		}

		atomic.AddInt64(&totalCount, int64(len(output.Versions)))
		atomic.AddInt64(&totalCount, int64(len(output.DeleteMarkers)))
		atomic.AddInt64(&totalCount, int64(len(output.CommonPrefixes)))

		folders, versions, deleteMarkers := getVersionsResult(output)
		totalFolders = append(totalFolders, folders...)
		totalVersions = append(totalVersions, versions...)
		totalDeleteMarkers = append(totalDeleteMarkers, deleteMarkers...)
		if !output.IsTruncated {
			break
		}

		if truncated {
			nextKeyMarker = output.NextKeyMarker
			nextVersionIdMarker = output.NextVersionIdMarker
			break
		}

		input.KeyMarker = output.NextKeyMarker
		input.VersionIdMarker = output.NextVersionIdMarker
	}

	h.End()

	if hasListError != nil {
		printError(hasListError)
		return assist.ErrUncompeleted
	}

	totalFoldersNumber := len(totalFolders)
	totalVersionsNumber := len(totalVersions)
	totalDeleteMarkersNumber := len(totalDeleteMarkers)
	var totalSize int64
	if c.short {
		if totalFoldersNumber > 0 {
			printf("Folder list:")
			for _, prefix := range totalFolders {
				printf("obs://%s/%s", bucket, prefix)
			}
			printf("")
		}
		if totalVersionsNumber > 0 {
			printf("Versioning Object list:")
			for _, val := range totalVersions {
				printf("obs://%s/%s", bucket, val.Key)
				totalSize += val.Size
			}
			printf("")
		}
		if totalDeleteMarkersNumber > 0 {
			printf("DeleteMarker list:")
			for _, val := range totalDeleteMarkers {
				printf("obs://%s/%s", bucket, val.Key)
			}
			printf("")
		}
	} else {
		if totalFoldersNumber > 0 {
			printf("Folder list:")
			for _, prefix := range totalFolders {
				printf("obs://%s/%s", bucket, prefix)
			}
			printf("")
		}
		if totalVersionsNumber > 0 {
			printf("Versioning Object list:")
			printf("%-30s%-40s%-30s%-10s%-20s%-20s", "Key", "VersionId", "LastModified", "Size", "StorageClass", "ETag")
			for _, val := range totalVersions {
				objectKey := "obs://" + bucket + "/" + val.Key
				printf("%s", objectKey)
				objectSizeStr := normalizeBytesByBytesFormat(c.bytesFormat, val.Size)
				if len(objectSizeStr) < 10 {
					printf("%-30s%-40s%-30s%-10s%-20s%-20s", "", val.VersionId, val.LastModified.Format(ISO8601_DATE_FORMAT),
						objectSizeStr, transStorageClassType(val.StorageClass), val.ETag)
				} else {
					printf("%-100s%s", "", objectSizeStr)
					printf("%-30s%-40s%-30s%-10s%-20s%-20s", "", val.VersionId, val.LastModified.Format(ISO8601_DATE_FORMAT),
						"", transStorageClassType(val.StorageClass), val.ETag)
				}
				printf("")
				totalSize += val.Size
			}
		}
		if totalDeleteMarkersNumber > 0 {
			printf("DeleteMarker list:")
			printf("%-30s%-40s%-30s%-20s", "Key", "VersionId", "LastModified", "StorageClass")
			for _, val := range totalDeleteMarkers {
				objectKey := "obs://" + bucket + "/" + val.Key
				printf("%s", objectKey)
				printf("%-30s%-40s%-30s%-20s", "", val.VersionId, val.LastModified.Format(ISO8601_DATE_FORMAT),
					transStorageClassType(val.StorageClass))
				printf("")
			}
		}
	}

	if nextKeyMarker != "" || nextVersionIdMarker != "" {
		printf("Next key marker is: %s", nextKeyMarker)
		printf("Next version id marker is: %s", nextVersionIdMarker)
	} else if !c.dir {
		if prefix == "" {
			printf("Total size of bucket is: %s", normalizeBytesByBytesFormat(c.bytesFormat, totalSize))
		} else {
			printf("Total size of prefix [%s] is: %s", prefix, normalizeBytesByBytesFormat(c.bytesFormat, totalSize))
		}
	}

	printf("Folder number is: %d", totalFoldersNumber)
	printf("Versioning file number is: %d", totalVersionsNumber)
	printf("DeleteMarker number is: %d", totalDeleteMarkersNumber)
	return nil
}

func (c *lsCommand) listObjects(bucket, prefix string) error {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefix
	input.Marker = c.marker
	input.RequestPayer = c.payer
	if c.dir {
		input.Delimiter = "/"
		if input.Prefix != "" && !isObsFolder(input.Prefix) {
			input.Prefix += "/"
		}
	}

	limit := c.limit
	count := limit
	if count <= 0 {
		count = 1000
	}
	truncated := false

	totalFolders := make([]string, 0, count)
	totalObjects := make([]obs.Content, 0, count)
	nextMarker := ""
	var totalCount int64
	h := &assist.HintV2{}
	h.MessageFunc = func() string {
		count := ""
		if tc := atomic.LoadInt64(&totalCount); tc > 0 {
			count = "[" + assist.Int64ToString(tc) + "]"
		}
		return fmt.Sprintf("Listing objects %s", count)
	}
	h.Start()

	var hasListError error
	for {
		if limit > 0 {
			if limit <= 1000 {
				input.MaxKeys = int(limit)
				truncated = true
			} else {
				input.MaxKeys = 1000
				limit -= 1000
			}
		}

		output, err := obsClient.ListObjects(input)
		if err != nil {
			hasListError = err
			break
		}

		atomic.AddInt64(&totalCount, int64(len(output.Contents)))
		atomic.AddInt64(&totalCount, int64(len(output.CommonPrefixes)))

		folders, objects := getObjectsResult(output)
		totalFolders = append(totalFolders, folders...)
		totalObjects = append(totalObjects, objects...)
		if !output.IsTruncated {
			break
		}

		if truncated {
			nextMarker = output.NextMarker
			break
		}

		input.Marker = output.NextMarker
	}

	h.End()

	if hasListError != nil {
		printError(hasListError)
		return assist.ErrUncompeleted
	}

	printListObjectsResult(totalFolders, totalObjects, c.short, c.dir, bucket, prefix, nextMarker, c.bytesFormat)
	return nil
}

func (c *lsCommand) getUploadsResult(output *obs.ListMultipartUploadsOutput) ([]string, []obs.Upload) {
	folders := make([]string, 0, len(output.CommonPrefixes))
	uploads := make([]obs.Upload, 0, len(output.Uploads))

	folders = append(folders, output.CommonPrefixes...)
	for _, upload := range output.Uploads {
		if isObsFolder(upload.Key) {
			folders = append(folders, upload.Key)
		} else {
			uploads = append(uploads, upload)
		}
	}
	return folders, uploads

}

func (c *lsCommand) listMultipartUploads(bucket, prefix string) error {
	input := &obs.ListMultipartUploadsInput{}
	input.Bucket = bucket
	input.Prefix = prefix
	input.KeyMarker = c.marker
	input.UploadIdMarker = c.uploadIdMarker
	input.RequestPayer = c.payer
	if c.dir {
		input.Delimiter = "/"
		if input.Prefix != "" && !isObsFolder(input.Prefix) {
			input.Prefix += "/"
		}
	}
	limit := c.limit
	count := limit
	if count <= 0 {
		count = 1000
	}

	truncated := false
	totalFolders := make([]string, 0, count)
	totalUploads := make([]obs.Upload, 0, count)
	nextKeyMarker := ""
	nextUploadIdMarker := ""
	var totalCount int64
	h := &assist.HintV2{}
	h.MessageFunc = func() string {
		return fmt.Sprintf("Listing multipart uploads [%d]", atomic.LoadInt64(&totalCount))
	}
	h.Start()
	var hasListError error
	for {
		if limit > 0 {
			if limit <= 1000 {
				input.MaxUploads = int(limit)
				truncated = true
			} else {
				input.MaxUploads = 1000
				limit -= 1000
			}
		}
		output, err := obsClient.ListMultipartUploads(input)
		if err != nil {
			hasListError = err
			break
		}

		atomic.AddInt64(&totalCount, int64(len(output.Uploads)))
		atomic.AddInt64(&totalCount, int64(len(output.CommonPrefixes)))

		folders, uploads := c.getUploadsResult(output)
		totalFolders = append(totalFolders, folders...)
		totalUploads = append(totalUploads, uploads...)
		if !output.IsTruncated {
			break
		}
		if truncated {
			nextKeyMarker = output.NextKeyMarker
			nextUploadIdMarker = output.NextUploadIdMarker
			break
		}

		input.KeyMarker = output.NextKeyMarker
		input.UploadIdMarker = output.NextUploadIdMarker
	}

	h.End()

	if hasListError != nil {
		printError(hasListError)
		return assist.ErrUncompeleted
	}

	totalFoldersNumber := len(totalFolders)
	totalUploadsNumber := len(totalUploads)
	if c.short {
		if totalFoldersNumber > 0 {
			printf("Folder list:")
			for _, prefix := range totalFolders {
				printf("obs://%s/%s", bucket, prefix)
			}
			printf("")
		}
		if totalUploadsNumber > 0 {
			printf("Upload list:")
			for _, val := range totalUploads {
				printf("obs://%s/%s %s", bucket, val.Key, val.UploadId)
			}
			printf("")
		}
	} else {
		if totalFoldersNumber > 0 {
			printf("Folder list:")
			for _, prefix := range totalFolders {
				printf("obs://%s/%s", bucket, prefix)
			}
			printf("")
		}
		if totalUploadsNumber > 0 {
			printf("Upload list:")
			printf("%-50s%-30s%-20s%-20s", "Key", "Initiated", "StorageClass", "UploadId")
			for _, val := range totalUploads {
				objectKey := "obs://" + bucket + "/" + val.Key
				if len(objectKey) >= 50 || assist.HasUnicode(objectKey) {
					printf("%s", objectKey)
					printf("%-50s%-30s%-20s%-20s", "", val.Initiated.Format(ISO8601_DATE_FORMAT),
						transStorageClassType(val.StorageClass), val.UploadId)
					printf("")
				} else {
					printf("%-50s%-30s%-20s%-20s", objectKey, val.Initiated.Format(ISO8601_DATE_FORMAT),
						transStorageClassType(val.StorageClass), val.UploadId)
					printf("")
				}
			}
		}
	}

	if nextKeyMarker != "" || nextUploadIdMarker != "" {
		printf("Next keyMarker is: %s", nextKeyMarker)
		printf("Next uploadIdMarker is: %s", nextUploadIdMarker)
	}

	printf("Folder number is: %d", totalFoldersNumber)
	printf("Upload number is: %d", totalUploadsNumber)
	return nil
}

func initLs() command {

	c := &lsCommand{}
	c.key = "ls"
	c.usage = "[cloud_url] [options...]"
	c.description = "list buckets or objects/multipart uploads in a bucket"

	c.define = func() {
		c.flagSet.BoolVar(&c.short, "s", false, "")
		c.flagSet.BoolVar(&c.storageClass, "sc", false, "")
		c.flagSet.BoolVar(&c.dir, "d", false, "")
		c.flagSet.BoolVar(&c.multipart, "m", false, "")
		c.flagSet.BoolVar(&c.all, "a", false, "")
		c.flagSet.BoolVar(&c.version, "v", false, "")
		c.flagSet.StringVar(&c.marker, "marker", "", "")
		c.flagSet.StringVar(&c.versionIdMarker, "versionIdMarker", "", "")
		c.flagSet.StringVar(&c.uploadIdMarker, "uploadIdMarker", "", "")
		c.flagSet.StringVar(&c.bytesFormat, "bf", "", "")
		c.flagSet.Int64Var(&c.limit, "limit", 1000, "")
		c.flagSet.IntVar(&c.jobs, "j", 0, "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.emptyArgsAction = func() error {
		c.printStart()
		return c.listBuckets()
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err == errEmptyArgs {
			return c.emptyArgsAction()
		}

		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		bucket, prefix, err := c.splitCloudUrl(cloudUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		_, succeed := getRequestPayerType(c.payer)
		if !succeed {
			return assist.ErrInvalidArgs
		}

		if c.bytesFormat != "" && c.bytesFormat != "human-readable" && c.bytesFormat != c_raw {
			printf("Error: Invalid bf [%s], possible values are:[human-readable|raw]", c.bytesFormat)
			return assist.ErrInvalidArgs
		}

		c.printStart()

		if c.all {
			var ret error
			if c.version {
				ret = c.listVersions(bucket, prefix)
			} else {
				ret = c.listObjects(bucket, prefix)
			}
			if ret != nil {
				return ret
			}
			printf("")
			return c.listMultipartUploads(bucket, prefix)
		}
		if c.multipart {
			return c.listMultipartUploads(bucket, prefix)
		}
		if c.version {
			return c.listVersions(bucket, prefix)
		}
		return c.listObjects(bucket, prefix)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("list buckets or objects/multipart uploads in a bucket"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil ls [-s] [-sc] [-j=1] [-limit=1] [-config=xxx]"+commandCommonSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil ls obs://bucket[/prefix] [-s] [-d] [-v] [-marker=xxx] [-versionIdMarker=xxx] [-bf=xxx] [-limit=1] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 3:")
		printf("%2s%s", "", "obsutil ls obs://bucket[/prefix] [-s] [-d] [-v] -m [-a] [-uploadIdMarker=xxx] [-marker=xxx] [-versionIdMarker=xxx] [-limit=1] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-s")
		printf("%4s%s", "", p.Sprintf("show results in brief mode"))
		printf("")
		printf("%2s%s", "", "-sc")
		printf("%4s%s", "", p.Sprintf("show storage class of each bucket"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent jobs for querying the storage classes of buckets, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-d")
		printf("%4s%s", "", p.Sprintf("list objects and sub-folders in the current folder"))
		printf("")
		printf("%2s%s", "", "-v")
		printf("%4s%s", "", p.Sprintf("list versions of objects in a bucket"))
		printf("")
		printf("%2s%s", "", "-m")
		printf("%4s%s", "", p.Sprintf("list multipart uploads"))
		printf("")
		printf("%2s%s", "", "-a")
		printf("%4s%s", "", p.Sprintf("list both objects and multipart uploads"))
		printf("")
		printf("%2s%s", "", "-marker=xxx")
		printf("%4s%s", "", p.Sprintf("the marker to list objects or multipart uploads"))
		printf("")
		printf("%2s%s", "", "-versionIdMarker=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID marker to list versions of objects"))
		printf("")
		printf("%2s%s", "", "-uploadIdMarker=xxx")
		printf("%4s%s", "", p.Sprintf("the upload ID marker to list multipart uploads"))
		printf("")
		printf("%2s%s", "", "-bf=xxx")
		printf("%4s%s", "", p.Sprintf("the bytes format in results when listing objects or listing versions of objects, possible values are [human-readable|raw]"))
		printf("")
		printf("%2s%s", "", "-limit=1000")
		printf("%4s%s", "", p.Sprintf("show results by limited number"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
	}
	return c
}
