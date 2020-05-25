package assist

import (
	"os"
	"syscall"
	"time"
)

func GetFileAccessTime(fileInfo os.FileInfo) time.Time {
	defer func() {
		if err := recover(); err != nil {
			// ignore not get the file access time
		}
	}()
	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	nanoseconds := fileSys.LastAccessTime.Nanoseconds()
	return time.Unix(nanoseconds/1e9, 0)
}
