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
	"io/ioutil"
	"obs"
	"os"
	"path/filepath"
	"progress"
	"strings"
	"sync"
)

type clearCommand struct {
	defaultCommand
	uploadMode   bool
	downloadMode bool
	copyMode     bool
}

type clearResult struct {
	succeed bool
	path    string
}

type clearTask struct {
	path string
}

func (*clearTask) abortMultipartUpload(bucket, key, uploadId string) (bool, error) {
	if bucket == "" || key == "" || uploadId == "" {
		return true, nil
	}
	input := &obs.AbortMultipartUploadInput{}
	input.Bucket = bucket
	input.Key = key
	input.UploadId = uploadId
	output, err := obsClient.AbortMultipartUpload(input)
	if err == nil {
		doLog(LEVEL_DEBUG, "Abort multipart upload [%s] in the bucket [%s] successfully, request id [%s]", key, bucket, output.RequestId)
		return true, nil
	}

	if obsError, ok := err.(obs.ObsError); ok && obsError.StatusCode >= 300 && obsError.StatusCode < 500 && obsError.StatusCode != 408 {
		return true, obsError
	}
	doLog(LEVEL_DEBUG, "Abort multipart upload [%s] in the bucket [%s] failed, %s", key, bucket, err.Error())
	return false, err
}

func (*clearTask) loadCheckpoint(checkpointFile string, result interface{}) error {
	ret, err := ioutil.ReadFile(checkpointFile)
	if err != nil {
		return err
	}
	return assist.ParseXml(ret, result)
}

type clearUploadTask struct {
	clearTask
}

func (t *clearUploadTask) Run() interface{} {
	ufc := &UploadFileCheckpoint{}
	if err := t.loadCheckpoint(t.path, ufc); err == nil {
		aborted, _ := t.abortMultipartUpload(ufc.Bucket, ufc.Key, ufc.UploadId)
		if !aborted {
			return &clearResult{
				succeed: false,
				path:    t.path,
			}
		}
	}
	if err := os.Remove(t.path); err != nil {
		return &clearResult{
			succeed: false,
			path:    t.path,
		}
	}
	return &clearResult{
		succeed: true,
		path:    t.path,
	}
}

type clearCopyTask struct {
	clearTask
}

func (t *clearCopyTask) Run() interface{} {
	cfc := &CopyObjectCheckpoint{}
	if err := t.loadCheckpoint(t.path, cfc); err == nil {
		aborted, _ := t.abortMultipartUpload(cfc.DestinationBucket, cfc.DestinationKey, cfc.UploadId)
		if !aborted {
			return &clearResult{
				succeed: false,
				path:    t.path,
			}
		}
	}
	if err := os.Remove(t.path); err != nil {
		return &clearResult{
			succeed: false,
			path:    t.path,
		}
	}
	return &clearResult{
		succeed: true,
		path:    t.path,
	}
}

type clearDownTask struct {
	clearTask
}

func (t *clearDownTask) Run() interface{} {
	dfc := &DownloadFileCheckpoint{}
	if err := t.loadCheckpoint(t.path, dfc); err == nil {
		if dfc.TempFileInfo.TempFileUrl != "" {
			if stat, err := os.Stat(dfc.TempFileInfo.TempFileUrl); err == nil && !stat.IsDir() {
				if err := os.Remove(dfc.TempFileInfo.TempFileUrl); err != nil {
					return &clearResult{
						succeed: false,
						path:    t.path,
					}
				}
			}
		}
	}
	if err := os.Remove(t.path); err != nil {
		return &clearResult{
			succeed: false,
			path:    t.path,
		}
	}
	return &clearResult{
		succeed: true,
		path:    t.path,
	}
}

func (c *clearCommand) clearCheckpointDir(checkpointDir string, mode int) error {
	c.printStart()

	poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
	futureChan := make(chan concurrent.Future, poolCacheCount)
	pool := concurrent.NewRoutinePool(assist.StringToInt(config["defaultJobs"], defaultJobs), poolCacheCount)

	var totalFiles int64
	wg := new(sync.WaitGroup)
	wg.Add(1)
	barCh := newSingleBarChan()
	barCh.SetTemplate(progress.Simple)
	failedPaths := make([]string, 0, 10)
	go func() {
		for {
			future, ok := <-futureChan
			if !ok {
				break
			}
			result := future.Get().(*clearResult)
			if result.succeed {
				barCh.Send64(1)
				progress.AddSucceedCount(1)
			} else {
				failedPaths = append(failedPaths, result.path)
				progress.AddFailedCount(1)
			}
			progress.AddFinishedCount(1)
		}
		wg.Done()
	}()

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			totalFiles++
			var task concurrent.Task
			if mode == 1 {
				t := &clearDownTask{}
				t.path = path
				task = t
			} else if mode == 2 {
				t := &clearCopyTask{}
				t.path = path
				task = t
			} else {
				t := &clearUploadTask{}
				t.path = path
				task = t
			}
			future, _ := pool.Submit(task)
			if future != nil {
				futureChan <- future
			}
		}
		return err
	}
	hasListError := filepath.Walk(checkpointDir, walkFunc)
	barCh.SetTotalCount(totalFiles)
	progress.SetTotalCount(totalFiles)

	barCh.Start()
	close(futureChan)
	pool.ShutDown()
	wg.Wait()
	barCh.WaitToFinished()
	printf("%-20s%-10d%-20s%-10d", "Succeed files is: ", progress.GetSucceedCount(), "Failed files is: ", progress.GetFailedCount())
	for index, failedPath := range failedPaths {
		printf("The [%d] failed path: %s", index, failedPath)
	}
	if hasListError != nil {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("List local files in local folder [%s] failed", checkpointDir))
		return assist.ErrUncompeleted
	}

	if progress.GetFailedCount() > 0 {
		return assist.ErrUncompeleted
	}
	return nil
}

func (c *clearCommand) ensureCheckpointDirectory(checkpointDir string) error {
	checkpointDir = strings.TrimSpace(checkpointDir)
	if checkpointDir == "" {
		dir, err := getCheckpointDirectory()
		if err != nil {
			return err
		}
		checkpointDir = dir
	}

	stat, err := os.Stat(checkpointDir)
	if err == nil && !stat.IsDir() {
		return fmt.Errorf("The specified checkpoint folder [%s] is a file", checkpointDir)
	}

	return nil
}

func initClear() command {
	c := &clearCommand{}
	c.key = "clear"
	c.usage = "[checkpoint_dir] [options...]"
	c.description = "delete part records"
	c.additional = true

	c.define = func() {
		c.flagSet.BoolVar(&c.uploadMode, "u", false, "")
		c.flagSet.BoolVar(&c.downloadMode, "d", false, "")
		c.flagSet.BoolVar(&c.copyMode, "c", false, "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		var checkpointDir string
		if len(args) <= 0 {
			dir, err := getCheckpointDirectory()
			if err != nil {
				printError(err)
				return assist.ErrExecuting
			}
			checkpointDir = dir
		} else if len(args) >= 1 {
			checkpointDir = args[0]
			if err := c.flagSet.Parse(args[1:]); err != nil {
				c.showHelp()
				printError(err)
				return assist.ErrInvalidArgs
			}
			if len(c.flagSet.Args()) >= 1 {
				c.showHelp()
				printf("Error: Invalid args, please refer to help doc")
				return assist.ErrInvalidArgs
			}
		}
		if err := c.ensureCheckpointDirectory(checkpointDir); err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		if !c.uploadMode && !c.downloadMode && !c.copyMode {
			printf("Error: No mode specified!")
			return assist.ErrInvalidArgs
		}

		if c.uploadMode {
			uploadCheckpointDir := checkpointDir + "/upload"
			printf("Clear checkpoint files for uploading in folder [%s]", uploadCheckpointDir)
			if stat, err := os.Stat(uploadCheckpointDir); err == nil && stat.IsDir() {
				return c.clearCheckpointDir(uploadCheckpointDir, 0)
			}
			printf("Error: Clear checkpoint files for uploading failed, [%s] is not a valid folder", uploadCheckpointDir)
			return assist.ErrInvalidArgs
		}

		if c.downloadMode {
			downloadCheckpointDir := checkpointDir + "/download"
			printf("Clear checkpoint files for downloading in folder [%s]", downloadCheckpointDir)
			if stat, err := os.Stat(downloadCheckpointDir); err == nil && stat.IsDir() {
				return c.clearCheckpointDir(downloadCheckpointDir, 1)
			}
			printf("Error: Clear checkpoint files for downloading failed, [%s] is not a valid folder", downloadCheckpointDir)
			return assist.ErrInvalidArgs
		}

		if c.copyMode {
			copyCheckpointDir := checkpointDir + "/copy"
			printf("Clear checkpoint files for copying in folder [%s]", copyCheckpointDir)
			if stat, err := os.Stat(copyCheckpointDir); err == nil && stat.IsDir() {
				return c.clearCheckpointDir(copyCheckpointDir, 2)
			}
			printf("Error: Clear checkpoint files for copying failed, [%s] is not a valid folder", copyCheckpointDir)
			return assist.ErrInvalidArgs
		}

		return nil
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("delete part records"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil clear [checkpoint_dir] [-u] [-d] [-c] [-config=xxx]"+commandCommonSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-u")
		printf("%4s%s", "", p.Sprintf("delete the part records of all multipart upload jobs"))
		printf("")
		printf("%2s%s", "", "-d")
		printf("%4s%s", "", p.Sprintf("delete the part records of all multipart download jobs"))
		printf("")
		printf("%2s%s", "", "-c")
		printf("%4s%s", "", p.Sprintf("delete the part records of all multipart copy jobs"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
	}

	return c
}
