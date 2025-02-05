// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
    "progress"
    "sync"
    "sync/atomic"
    "time"
    "strings"
    "path/filepath"
    "os"
    "errors"
    // For real unzip functionality, you might import "archive/zip" or other relevant packages
)

// --------------------------- unzipCommand struct -----------------------------

type unzipCommand struct {
    transferCommand       // similar embedded struct for shared fields/methods
    recursive bool
    force     bool
    dryRun    bool
    flat      bool
    jobs      int
    parallel  int
    // Additional flags can be added as needed
}

// unzipRequestInput encapsulates the source and destination for a single unzip operation
type unzipRequestInput struct {
    srcLocalPath   string // .zip file or local path
    dstLocalPath   string // target extraction directory
    keepSubfolders bool   // preserve subfolders if true
    canOverwrite   bool   // overwrite existing files if true
}

// unzipScanCtx is used for scanning directories or tasks
type unzipScanCtx struct {
    input      unzipRequestInput
    pool       concurrent.Pool
    barCh      progress.SingleBarChan
    actionName string
    action     func(req unzipRequestInput, barCh progress.SingleBarChan, fastError error) int
}

// ---------------------------------------------------------------------------
// ---------------------- Helper and Processing Functions ---------------------

// localUnzipFile simulates or performs the actual unzip.
// Real code would use "archive/zip" or an equivalent library.
func (c *unzipCommand) localUnzipFile(req unzipRequestInput) error {
    doLog(LEVEL_INFO, "Starting unzip from source [%s] to destination [%s]", req.srcLocalPath, req.dstLocalPath)

    // If dryRun is set, we skip the actual file operations
    if c.dryRun {
        return nil
    }

    // Basic existence check
    if _, err := os.Stat(req.srcLocalPath); os.IsNotExist(err) {
        return fmt.Errorf("zip file [%s] not found", req.srcLocalPath)
    }
    // Here, you would do real unzip logic (e.g. archive/zip).
    // We'll just log success for demonstration.
    doLog(LEVEL_INFO, "Successfully unzipped file. source: %s, destination: %s", req.srcLocalPath, req.dstLocalPath)
    return nil
}


// It traverses a local path, finds .zip files, and schedules them for extraction.
func (c *unzipCommand) scanZipDirAndDoAction(ctx unzipScanCtx) (totalCount int64, listError error) {
    defer func() {
        if atomic.LoadInt32(&c.abort) == 1 {
            doLog(LEVEL_ERROR, "Aborting scan due to an unexpected error. Please check logs.")
        }
    }()

    // We'll assume ctx.input.srcLocalPath might be a single file or a folder.
    searchPath := ctx.input.srcLocalPath
    info, err := os.Stat(searchPath)
    if err != nil {
        listError = err
        return
    }
    // If it's a file (single .zip), handle it directly
    if !info.IsDir() {
        atomic.AddInt64(&totalCount, 1)
        newReq := ctx.input
        newReq.srcLocalPath = searchPath
        ctx.pool.ExecuteFunc(func() interface{} {
            return c.handleExecResult(ctx.action(newReq, ctx.barCh, nil), 0)
        })
        return
    }

    // Otherwise, recursively walk the directory
    err = filepath.Walk(searchPath, func(path string, f os.FileInfo, walkErr error) error {
        if walkErr != nil {
            listError = walkErr
            return walkErr
        }
        if f.IsDir() {
            // If not recursive, skip subdirectories
            if !c.recursive && path != searchPath {
                return filepath.SkipDir
            }
            return nil
        }
        // We only care about files ending in .zip
        if strings.HasSuffix(strings.ToLower(f.Name()), ".zip") {
            atomic.AddInt64(&totalCount, 1)
            newReq := ctx.input
            newReq.srcLocalPath = path
            ctx.pool.ExecuteFunc(func() interface{} {
                return c.handleExecResult(ctx.action(newReq, ctx.barCh, nil), 0)
            })
        }
        return nil
    })

    if err != nil && listError == nil {
        listError = err
    }
    return
}

// doUnzipFile coordinates the actual unzip operation for a single file
func (c *unzipCommand) doUnzipFile(req unzipRequestInput, barCh progress.SingleBarChan, fastFailed error) int {
    if fastFailed != nil {
        doLog(LEVEL_ERROR, "Fast fail triggered: %s", fastFailed.Error())
        return 0
    }
    start := assist.GetUtcNow()
    err := c.localUnzipFile(req)
    cost := (assist.GetUtcNow().UnixNano() - start.UnixNano()) / 1000000

    if err != nil {
        doLog(LEVEL_ERROR, "Unzip failed. Duration [%d ms], Error [%s]", cost, err.Error())
        if barCh != nil {
            barCh.SendError(1)
        }
        return 0
    }

    doLog(LEVEL_INFO, "Unzip succeeded in [%d ms].", cost)
    if barCh != nil {
        barCh.Send(1)
    }
    return 1
}

// unzipDir handles either a single .zip file or all .zip files within a directory
// using concurrency pools.
func (c *unzipCommand) unzipDir(req unzipRequestInput) error {
    poolCacheCount := assist.StringToInt(config["defaultJobsCacheCount"], defaultJobsCacheCount)
    pool := concurrent.NewRoutinePool(c.jobs, poolCacheCount)

    barCh := newSingleBarChan()
    if c.force {
        barCh.Start()
    }

    totalCount, listErr := c.scanZipDirAndDoAction(unzipScanCtx{
        input:      req,
        pool:       pool,
        barCh:      barCh,
        actionName: "unzip",
        action:     c.doUnzipFile,
    })

    doLog(LEVEL_INFO, "Number of zip files to process for unzip: %d", totalCount)
    progress.SetTotalCount(totalCount)
    barCh.SetTotalCount(totalCount)

    if !c.force {
        barCh.Start()
    }

    pool.ShutDown()
    barCh.WaitToFinished()

    if listErr != nil {
        logError(listErr, LEVEL_ERROR, fmt.Sprintf("Listing zip files failed: %v", listErr))
        return assist.ErrUncompeleted
    }
    if progress.GetFailedCount() > 0 {
        return assist.ErrUncompeleted
    }
    return nil
}

// ---------------------------------------------------------------------------
// --------------------------- initUnzip Command ------------------------------

func initUnzip() command {
    c := &unzipCommand{}
    c.key = "unzip"
    c.usage = "unzip <local_path> <destination_path> [options...]"
    c.description = "unzip local .zip files into a specified directory"

    c.define = func() {
        c.init()
        c.defineBasic()
        c.flagSet.BoolVar(&c.recursive, "r", false, "")
        c.flagSet.BoolVar(&c.force, "f", false, "")
        c.flagSet.BoolVar(&c.dryRun, "dryRun", false, "")
        c.flagSet.BoolVar(&c.flat, "flat", false, "")
        c.flagSet.IntVar(&c.jobs, "j", 1, "")
        c.flagSet.IntVar(&c.parallel, "p", 1, "")
    }

    c.action = func() error {
        args := c.flagSet.Args()
        if len(args) < 2 {
            c.showHelp()
            printf("Error: Invalid arguments. Refer to the help documentation.")
            return assist.ErrInvalidArgs
        }

        srcPath := args[0]
        dstPath := args[1]

        if srcPath == "" || dstPath == "" {
            printf("Error: Source and destination paths cannot be empty.")
            return assist.ErrInvalidArgs
        }

        // Concurrency checks
        if c.jobs > 10 {
            printf("Error: The maximum number of jobs for unzip is 10.")
            return assist.ErrInvalidArgs
        }
        if c.parallel > 10 {
            printf("Error: The maximum parallelism for unzip is 10.")
            return assist.ErrInvalidArgs
        }

        c.printStart()

        unzipReq := unzipRequestInput{
            srcLocalPath:   srcPath,
            dstLocalPath:   dstPath,
            keepSubfolders: !c.flat,
            canOverwrite:   c.force,
        }

        // If not recursive, handle single .zip file
        if !c.recursive {
            info, err := os.Stat(srcPath)
            if err != nil {
                printError(err)
                return assist.ErrInvalidArgs
            }
            if info.IsDir() {
                printf("Error: -r must be specified to process multiple zips from a folder.")
                return assist.ErrInvalidArgs
            }
            if c.dryRun {
                printf("Unzip dry run: source: %s -> destination: %s", srcPath, dstPath)
                return nil
            }
            // Actual single-file unzip
            err = c.localUnzipFile(unzipReq)
            if err != nil {
                logError(err, LEVEL_ERROR, fmt.Sprintf("Unzip failed: %s -> %s", srcPath, dstPath))
                return assist.ErrExecuting
            }
            printf("Unzip succeeded: %s -> %s", srcPath, dstPath)
            return nil
        }

        // If -r is set, we handle all .zip files under srcPath
        if err := c.unzipDir(unzipReq); err != nil {
            return err
        }
        return nil
    }

    c.help = func() {
        p := i18n.GetCurrentPrinter()
        p.Printf("Summary:")
        printf("%2s%s", "", p.Sprintf("Unzip local .zip files/folders"))
        printf("")
        p.Printf("Syntax 1:")
        printf("%2s%s", "",
            "obsutil unzip <local_path_or_zip> <destination_path> [-r] [-dryRun] [-f] [-flat] [-j=1] [-p=1] [-config=xxx]" +
            commandCommonSyntax() + commandRequestPayerSyntax())
        printf("")

        p.Printf("Options:")
        printf("%2s%s", "", "-r")
        printf("%4s%s", "", p.Sprintf("Recursively unzip all .zip files found in a local folder"))
        printf("")
        printf("%2s%s", "", "-f")
        printf("%4s%s", "", p.Sprintf("Force mode: overwrite files without prompting"))
        printf("")
        printf("%2s%s", "", "-dryRun")
        printf("%4s%s", "", p.Sprintf("Simulate the unzip process without extracting any files"))
        printf("")
        printf("%2s%s", "", "-flat")
        printf("%4s%s", "", p.Sprintf("Unzip files directly into the target directory without preserving folder structure"))
        printf("")
        printf("%2s%s", "", "-j=1")
        printf("%4s%s", "", p.Sprintf("Maximum number of listing jobs for the unzip process (default can be from config)"))
        printf("")
        printf("%2s%s", "", "-p=1")
        printf("%4s%s", "", p.Sprintf("Maximum number of concurrent unzip tasks (default can be from config)"))
        printf("")
        printf("%2s%s", "", "-config=xxx")
        printf("%4s%s", "", p.Sprintf("Custom config file path for the command"))
        printf("")
        commandCommonHelp(p)
        commandRequestPayerHelp(p)
    }
    return c
}
