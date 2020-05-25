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
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"net/url"
	"os"
	"path/filepath"
	"progress"
	"strings"
	"sync/atomic"
)

type directDownloadCommand struct {
	shareCpCommand
}

func (c *directDownloadCommand) checkArgs() (resourceUrl, key, originFileUrl, fileUrl string, err error) {
	args := c.flagSet.Args()
	if len(args) < 2 {
		c.showHelp()
		err = errors.New("Error: Invalid args, please refer to help doc")
		return
	}

	resourceUrl = args[0]
	parsedUrl, _err := url.Parse(resourceUrl)
	if _err != nil {
		err = _err
		return
	}

	key = parsedUrl.Path
	if strings.HasPrefix(key, "/") {
		key = key[1:]
	}

	originFileUrl = args[1]
	fileUrl, err = filepath.Abs(originFileUrl)
	if err != nil {
		return
	}

	if err = c.flagSet.Parse(args[2:]); err != nil {
		c.showHelp()
		return
	}

	if len(c.flagSet.Args()) >= 1 {
		c.showHelp()
		err = fmt.Errorf("Invalid args \"%v\", please refer to help doc", c.flagSet.Args())
		return
	}

	return
}

func initDirectDownload() command {
	c := &directDownloadCommand{}
	c.key = "download"
	c.usage = c_direct_download_usage
	c.description = "download an object directly using the specified resource url"
	c.skipCheckAkSk = true

	c.define = func() {
		c.recursiveCommand.init()
		c.partSize = 0
		c.bigfileThreshold = 0
		c.warn = atomic.Value{}
		c.warnFlag = 0

		c.flagSet.BoolVar(&c.dryRun, "dryRun", false, "")
		c.flagSet.BoolVar(&c.verifyLength, "vlength", false, "")
		c.flagSet.BoolVar(&c.verifyMd5, "vmd5", false, "")
		c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
		c.flagSet.IntVar(&c.parallel, "p", 0, "")
		c.flagSet.StringVar(&c.bigfileThresholdStr, "threshold", "", "")
		c.flagSet.StringVar(&c.partSizeStr, "ps", "", "")
		c.flagSet.StringVar(&c.checkpointDir, "cpd", "", "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.tempFileDir, "tempFileDir", "", "")
		c.flagSet.BoolVar(&c.update, "u", false, "")
	}

	c.action = func() error {
		resourceUrl, key, originFileUrl, fileUrl, err := c.checkArgs()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		if !c.prepareOptions() {
			return assist.ErrInvalidArgs
		}

		stat, err := os.Lstat(fileUrl)

		if (err != nil && (isObsFolder(originFileUrl) || strings.HasSuffix(originFileUrl, "\\"))) || (err == nil && stat.IsDir()) {
			fileName := key
			if fileName == "" {
				_fileName, err := uuid.NewV4()
				if err != nil {
					printf("Error: Cannot generate the random file name, err:%s", err.Error())
					return assist.ErrExecuting
				}
				fileName = _fileName.String()
				printf("The resolved file name is empty, will use the random file name [%s] instead", fileName)
			} else if index := strings.LastIndex(fileName, "/"); index >= 0 {
				fileName = fileName[index+1:]
			}
			fileUrl = assist.NormalizeFilePath(fileUrl + "/" + fileName)
			stat = nil
		}

		c.printStart()

		if c.forceRecord {
			return c.ensureOuputAndStartLogger(func() error {
				c.printParams(true, true, false, true)
				c.recordStartFuncForDownload()
				ret := c.downloadFile(resourceUrl, key, fileUrl, stat, nil, nil, 1, nil)
				if ret >= 1 {
					progress.AddSucceedCount(1)
					return nil
				}
				progress.AddFailedCount(1)
				return assist.ErrExecuting
			}, true)
		}

		c.printParams(false, true, false, true)
		ret := c.downloadFile(resourceUrl, key, fileUrl, stat, nil, nil, 0, nil)
		if warn, ok := c.warn.Load().(error); ok {
			printWarn(warn)
		}
		if ret == 0 {
			return assist.ErrExecuting
		}
		return nil
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("download an object directly using the specified resource url"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil download resource_url file_url|folder_url [-dryRun] [-tempFileDir=xxx] [-u] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]")
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-dryRun")
		printf("%4s%s", "", p.Sprintf("conduct a dry run"))
		printf("")
		printf("%2s%s", "", "-tempFileDir=xxx")
		printf("%4s%s", "", p.Sprintf("the temp file dir, used to save temporary files during the objects are downloading"))
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
		printf("%2s%s", "", "-cpd=xxx")
		printf("%4s%s", "", p.Sprintf("the directory where the part records reside, used to record the progress of download jobs"))
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
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the download results"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
	}

	return c
}
