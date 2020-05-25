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
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat_t.Atim.Sec), 0)
}
