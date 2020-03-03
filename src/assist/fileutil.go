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
package assist

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"hash/crc64"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync/atomic"
	"syscall"
	"unicode/utf8"
)

const (
	defaultChunkSize     = 8192
	readBufferIoSize     = 8192
	windowsMaxFileLength = 259
)

var uid = os.Getuid()
var gid = os.Getgid()
var clearSlashRegex = regexp.MustCompile("/+")

func findFirstValidPath(value []byte) (string, bool) {
	vals := strings.Split(BytesToString(value), "\n")
	for _, val := range vals {
		if _val := strings.TrimSpace(val); _val == "" {
			continue
		} else {
			return _val, true
		}
	}
	return "", false
}

func GetOsPath(path string) (ret string) {
	var c *exec.Cmd
	if IsWindows() {
		c = exec.Command("where", path)
		if value, err := c.CombinedOutput(); err == nil {
			if _value, flag := findFirstValidPath(value); flag {
				ret = _value
			}
		}
	} else {
		c = exec.Command("which", path)
		if value, err := c.CombinedOutput(); err == nil {
			if _value, flag := findFirstValidPath(value); flag {
				ret = _value
			}
		}
	}

	if ret == "" {
		ret = path
		_, err := os.Lstat(path)
		if err == nil {
			ret, _ = filepath.Abs(path)
		}
	}

	return
}

func checkLength(path string) error {
	if IsWindows() {
		path = NormalizeFilePath(path)
		if length := utf8.RuneCountInString(path); length >= windowsMaxFileLength {
			return fmt.Errorf("the length:%d of path:%s exceed the max length %d", length, path, windowsMaxFileLength)
		}
	}
	return nil
}

func ReadLine(rd *bufio.Reader) ([]byte, error) {
	line := make([]byte, 0, 4096)
	for {
		lineByte, isPrefix, err := rd.ReadLine()
		if err != nil {
			return nil, err
		}
		line = append(line, lineByte...)
		if !isPrefix {
			break
		}
	}
	return line, nil
}

func MkdirAll(path string, perm os.FileMode) error {
	if err := checkLength(path); err != nil {
		return err
	}
	return os.MkdirAll(path, perm)
}

func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	if err := checkLength(path); err != nil {
		return nil, err
	}

	return os.OpenFile(path, flag, perm)
}

func Chown(path string) error {
	return os.Chown(path, uid, gid)
}

func FindFiles(folder string, pattern *regexp.Regexp, action func(fileUrl string)) error {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && pattern.MatchString(path) {
			action(path)
		}
		return nil
	}
	return filepath.Walk(folder, walkFn)
}

func FindFilesV2(folder string, pattern *regexp.Regexp) ([]string, error) {
	files := make([]string, 0)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && pattern.MatchString(path) {
			files = append(files, path)
		}
		return nil
	}
	if err := filepath.Walk(folder, walkFn); err != nil {
		return nil, err
	}
	return files, nil
}

func FindMatches(fileUrl string, pattern *regexp.Regexp, action func(groups []string), abort *int32) error {
	fd, err := os.Open(fileUrl)
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	for {
		if atomic.LoadInt32(abort) == 1 {
			break
		}

		lineByte, err := ReadLine(rd)
		if err != nil {
			break
		}
		line := strings.TrimSpace(BytesToString(lineByte))
		if line == "" {
			continue
		}
		if pattern.MatchString(line) {
			action(pattern.FindStringSubmatch(line))
		}
	}
	return nil
}

func ReadContentLineByFileUrl(fileUrl string) ([]string, error) {
	lines := []string{}
	fd, err := os.Open(fileUrl)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	for {
		lineByte, err := ReadLine(rd)
		if err != nil {
			break
		}
		line := strings.TrimSpace(BytesToString(lineByte))
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func Md5File(fileUrl string) ([]byte, error) {
	fd, err := os.Open(fileUrl)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	m := md5.New()

	p := GetByteArrayFromPool()
	reader := bufio.NewReaderSize(fd, readBufferIoSize)
	for {
		n, err := reader.Read(p)
		if n > 0 {
			_, werr := m.Write(p[0:n])
			if werr != nil {
				PutByteArrayToPool(p)
				return nil, werr
			}
		}

		if err != nil {
			if err != io.EOF {
				PutByteArrayToPool(p)
				return nil, err
			}
			break
		}
	}
	PutByteArrayToPool(p)
	return m.Sum(nil), nil
}

func Crc64File(fileUrl string) (uint64, error) {
	fd, err := os.Open(fileUrl)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	m := crc64.New(crc64.MakeTable(crc64.ECMA))

	p := GetByteArrayFromPool()
	reader := bufio.NewReaderSize(fd, readBufferIoSize)
	for {
		n, err := reader.Read(p)
		if n > 0 {
			_, werr := m.Write(p[0:n])
			if werr != nil {
				PutByteArrayToPool(p)
				return 0, werr
			}
		}

		if err != nil {
			if err != io.EOF {
				PutByteArrayToPool(p)
				return 0, err
			}
			break
		}
	}
	PutByteArrayToPool(p)
	return m.Sum64(), nil
}

func GetRealPath(path string) (realPath string, realStat os.FileInfo, err error) {
	realPath = path
	realStat, err = os.Lstat(realPath)
	if err != nil {
		return
	}
	count := 0
	for realStat.Mode()&os.ModeSymlink == os.ModeSymlink {
		originRealPath := realPath
		realPath, err = os.Readlink(realPath)
		if err != nil {
			return
		}

		realPath = strings.TrimSpace(realPath)

		if !IsWindows() {
			if !strings.HasPrefix(realPath, "/") {
				realPath = filepath.Dir(originRealPath) + "/" + realPath
			}
		} else {
			if !strings.Contains(realPath, ":") {
				realPath = filepath.Dir(originRealPath) + "\\" + realPath
			}
		}

		realPath = NormalizeFilePath(realPath)
		if realPath == originRealPath && count >= 10 {
			err = errors.New("No such file or directory")
			return
		}

		realStat, err = os.Lstat(realPath)
		if err != nil {
			return
		}
		count++
	}

	//fixbug
	if !IsWindows() {
		realPath = clearSlashRegex.ReplaceAllString(realPath, "/")
		if strings.HasSuffix(realPath, "/") {
			realPath = realPath[:len(realPath)-1]
		}
	} else if strings.HasSuffix(realPath, "\\") {
		realPath = realPath[:len(realPath)-1]
	}

	return
}

func copyFile(oldpath, newpath string, writeBufferIoSize int, fsync bool) error {
	fd, fdErr := os.Open(oldpath)
	if fdErr != nil {
		return fdErr
	}
	defer fd.Close()
	wfd, wfdErr := OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if wfdErr != nil {
		return wfdErr
	}

	defer wfd.Close()
	p := GetByteArrayFromPool()
	bufWriter := bufio.NewWriterSize(wfd, writeBufferIoSize)
	for {
		n, err := fd.Read(p)
		if n > 0 {
			_, werr := bufWriter.Write(p[0:n])
			if werr != nil {
				PutByteArrayToPool(p)
				return werr
			}
		}

		if err != nil {
			if err != io.EOF {
				PutByteArrayToPool(p)
				return err
			}
			break
		}
	}

	PutByteArrayToPool(p)

	if ferr := bufWriter.Flush(); ferr != nil {
		return ferr
	}

	if fsync {
		if err := wfd.Sync(); err != nil {
			return err
		}
	}

	return nil
}

func CopyFile(oldpath, newpath string, checkNewpathLength bool, writeBufferIoSize int, fsync bool) error {
	if checkNewpathLength {
		if err := checkLength(newpath); err != nil {
			return err
		}
	}
	oldStat, oldStatErr := os.Stat(oldpath)
	if oldStatErr != nil {
		return oldStatErr
	}
	if oldStat.IsDir() {
		return fmt.Errorf("oldpath:%s is a dir", oldpath)
	}

	newfullpath, newfullpathErr := filepath.Abs(newpath)
	if newfullpathErr != nil {
		return newfullpathErr
	}

	newStat, newStatErr := os.Stat(newfullpath)
	if newStatErr == nil && newStat.IsDir() {
		return fmt.Errorf("newpath:%s is a dir", newpath)
	}

	if err := MkdirAll(filepath.Dir(newfullpath), os.ModePerm); err != nil {
		return err
	}

	if err := copyFile(oldpath, newfullpath, writeBufferIoSize, fsync); err != nil {
		return err
	}

	return nil
}

func RenameFile(oldpath, newpath string, forceOverwrite bool, writeBufferIoSize int, fsync bool) error {
	if err := checkLength(newpath); err != nil {
		return err
	}

	if forceOverwrite {
		if stat, _ := os.Stat(newpath); stat != nil {
			if err := os.Remove(newpath); err != nil {
				return err
			}
		}
	}

	if err := os.Rename(oldpath, newpath); err == nil {
		return nil
	}

	if err := CopyFile(oldpath, newpath, false, writeBufferIoSize, fsync); err != nil {
		return err
	}

	if err := os.Remove(oldpath); err != nil {
		return err
	}
	return nil
}

func PathListNested(fileList []string) error {
	afterRemoveNestedFilePath := make([]string, 0, len(fileList))
	sep := PathSeparator()
	for _, path := range fileList {
		for _, afterPath := range afterRemoveNestedFilePath {
			_path := path + sep
			_afterPath := afterPath + sep
			if strings.HasPrefix(_path, _afterPath) {
				return fmt.Errorf("The path [%s] is nested with path [%s]", path, afterPath)
			}
			if strings.HasPrefix(_afterPath, _path) {
				return fmt.Errorf("The path [%s] is nested with path [%s]", afterPath, path)
			}
		}
		afterRemoveNestedFilePath = append(afterRemoveNestedFilePath, path)
	}
	return nil
}

func PathSeparator() string {
	if IsWindows() {
		return "\\"
	}
	return "/"
}

func QuickCreateFile(fileUrl string, size int64) error {
	fd, err := syscall.Open(fileUrl, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer syscall.Close(fd)

	err = syscall.Ftruncate(fd, size)
	if err != nil {
		return err
	}
	return nil
}
