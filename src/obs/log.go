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

package obs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Level int

const (
	LEVEL_OFF   Level = 500
	LEVEL_ERROR Level = 400
	LEVEL_WARN  Level = 300
	LEVEL_INFO  Level = 200
	LEVEL_DEBUG Level = 100
)

var logLevelMap = map[Level]string{
	LEVEL_OFF:   "[OFF]: ",
	LEVEL_ERROR: "[ERROR]: ",
	LEVEL_WARN:  "[WARN]: ",
	LEVEL_INFO:  "[INFO]: ",
	LEVEL_DEBUG: "[DEBUG]: ",
}

type logConfType struct {
	level        Level
	logToConsole bool
	logFullPath  string
	maxLogSize   int64
	backups      int
}

func getDefaultLogConf() logConfType {
	return logConfType{
		level:        LEVEL_WARN,
		logToConsole: false,
		logFullPath:  "",
		maxLogSize:   1024 * 1024 * 30, //30MB
		backups:      10,
	}
}

var logConf logConfType

type loggerWrapper struct {
	fullPath   string
	fd         *os.File
	ch         chan string
	wg         sync.WaitGroup
	queue      []string
	logger     *log.Logger
	index      int
	cacheCount int
	closed     int32
}

func (lw *loggerWrapper) doInit() {
	lw.queue = make([]string, 0, lw.cacheCount)
	lw.logger = log.New(lw.fd, "", 0)
	lw.ch = make(chan string, lw.cacheCount)
	lw.wg.Add(1)
	go lw.doWrite()
}

func (lw *loggerWrapper) rotate() {
	stat, err := lw.fd.Stat()
	if err != nil {
		lw.fd.Close()
		panic(err)
	}
	if stat.Size() >= logConf.maxLogSize {
		lw.fd.Close()
		if lw.index > logConf.backups {
			lw.index = 1
		}
		renamePath := lw.fullPath + "." + IntToString(lw.index)
		if _stat, _ := os.Stat(renamePath); _stat != nil {
			if _err := os.Remove(renamePath); _err != nil {
				panic(_err)
			}
		}

		if _err := os.Rename(lw.fullPath, renamePath); _err != nil {
			panic(_err)
		}
		if chmodErr := os.Chmod(renamePath, 0400); chmodErr != nil {
			// ignore change mode error
		}
		lw.index++

		fd, err := os.OpenFile(lw.fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			panic(err)
		}
		lw.fd = fd
		lw.logger.SetOutput(lw.fd)
	}
}

func (lw *loggerWrapper) doFlush() {
	lw.rotate()
	if len(lw.queue) > 0 {
		lw.logger.Println(strings.Join(lw.queue, "\n"))
	}
}

func (lw *loggerWrapper) doClose() {
	if atomic.CompareAndSwapInt32(&lw.closed, 0, 1) {
		close(lw.ch)
		lw.wg.Wait()
	}
}

func (lw *loggerWrapper) doWrite() {
	defer lw.wg.Done()
	for {
		select {
		case msg, ok := <-lw.ch:
			if !ok {
				lw.doFlush()
				lw.fd.Close()
				return
			}
			if len(lw.queue) >= lw.cacheCount {
				lw.doFlush()
				lw.queue = lw.queue[:0]
			}
			lw.queue = append(lw.queue, msg)
		case <-time.After(time.Second * 5):
			if len(lw.queue) > 0 {
				lw.doFlush()
				lw.queue = lw.queue[:0]
			}
		}
	}

}

func (lw *loggerWrapper) Printf(format string, v ...interface{}) {
	if atomic.LoadInt32(&lw.closed) == 0 {
		msg := fmt.Sprintf(format, v...)
		lw.ch <- msg
	}
}

var consoleLogger *log.Logger
var fileLogger *loggerWrapper
var globalClosed int32 = 1
var lock = new(sync.RWMutex)

func isDebugLogEnabled() bool {
	return logConf.level <= LEVEL_DEBUG
}

func isErrorLogEnabled() bool {
	return logConf.level <= LEVEL_ERROR
}

func isWarnLogEnabled() bool {
	return logConf.level <= LEVEL_WARN
}

func isInfoLogEnabled() bool {
	return logConf.level <= LEVEL_INFO
}

func reset() {
	if fileLogger != nil {
		fileLogger.doClose()
		fileLogger = nil
	}
	consoleLogger = nil
	logConf = getDefaultLogConf()
}

func InitLog(logFullPath string, maxLogSize int64, backups int, level Level, logToConsole bool) error {
	return InitLogWithCacheCnt(logFullPath, maxLogSize, backups, level, logToConsole, 50)
}

func InitLogWithCacheCnt(logFullPath string, maxLogSize int64, backups int, level Level, logToConsole bool, cacheCnt int) error {
	lock.Lock()
	defer lock.Unlock()
	if cacheCnt <= 0 {
		cacheCnt = 50
	}
	reset()
	if fullPath := strings.TrimSpace(logFullPath); fullPath != "" {
		_fullPath, err := filepath.Abs(fullPath)
		if err != nil {
			return err
		}

		if !strings.HasSuffix(_fullPath, ".log") {
			_fullPath += ".log"
		}

		stat, err := os.Stat(_fullPath)
		if err == nil && stat.IsDir() {
			return fmt.Errorf("logFullPath:[%s] is a directory", _fullPath)
		} else if err = os.MkdirAll(filepath.Dir(_fullPath), 0750); err != nil {
			return err
		}

		fd, err := os.OpenFile(_fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
		if err != nil {
			return err
		}

		if stat == nil {
			stat, err = os.Stat(_fullPath)
			if err != nil {
				fd.Close()
				return err
			}
		}

		prefix := stat.Name() + "."
		index := 1
		walkFunc := func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if name := info.Name(); strings.HasPrefix(name, prefix) {
					if i := StringToInt(name[len(prefix):], 0); i >= index {
						index = i + 1
					}
				}
			}
			return err
		}

		if err = filepath.Walk(filepath.Dir(_fullPath), walkFunc); err != nil {
			fd.Close()
			return err
		}

		fileLogger = &loggerWrapper{fullPath: _fullPath, fd: fd, index: index, cacheCount: cacheCnt, closed: 0}
		fileLogger.doInit()
	}
	if maxLogSize > 0 {
		logConf.maxLogSize = maxLogSize
	}
	if backups > 0 {
		logConf.backups = backups
	}
	logConf.level = level
	if logToConsole {
		consoleLogger = log.New(os.Stdout, "", log.LstdFlags)
	}
	atomic.StoreInt32(&globalClosed, 0)
	return nil
}

func CloseLog() {
	if atomic.CompareAndSwapInt32(&globalClosed, 0, 1) {
		lock.Lock()
		defer lock.Unlock()
		reset()
	}
}

func SyncLog() {
}

func logEnabled() bool {
	return atomic.LoadInt32(&globalClosed) == 0
}

func DoLog(level Level, format string, v ...interface{}) {
	doLog(level, format, v)
}

func doLog(level Level, format string, v ...interface{}) {
	if logEnabled() && logConf.level <= level {
		msg := fmt.Sprintf(format, v...)
		if _, file, line, ok := runtime.Caller(1); ok {
			index := strings.LastIndex(file, "/")
			if index >= 0 {
				file = file[index+1:]
			}
			msg = fmt.Sprintf("%s:%d|%s", file, line, msg)
		}
		prefix := logLevelMap[level]
		defer func() {
			recover()
			// ignore ch closed error
		}()
		if consoleLogger != nil {
			consoleLogger.Printf("%s%s", prefix, msg)
		}
		if fileLogger != nil {
			nowDate := FormatUtcNow("2006-01-02T15:04:05Z")
			fileLogger.Printf("%s %s%s", nowDate, prefix, msg)
		}
	}
}
