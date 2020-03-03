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
package repeatable

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"obs"
)

var errUnsupportReset = errors.New("UnsupportResetError")

type Reader struct {
	fd         io.Reader
	reader     *bufio.Reader
	markOffset int64
	bufferSize int
}

type ReaderV2 struct {
	Reader
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r *ReaderV2) Read(p []byte) (n int, err error) {
	defer func() {
		e := recover()
		if _err, ok := e.(error); ok {
			err = _err
		} else if e != nil {
			err = fmt.Errorf("Unexpect error, please collect the logs and contact our support, %v", e)
		}
	}()

	return r.reader.Read(p)
}

func (r *Reader) Reset() error {
	err := errUnsupportReset
	if _fd, ok := r.fd.(io.ReadSeeker); ok {
		_, err = _fd.Seek(r.markOffset, 0)
	} else if _fd, ok := r.fd.(obs.IRepeatable); ok {
		err = _fd.Reset()
	}

	if err == nil {
		//fix bug, need rebuild bufio rather than reset it
		r.reader = bufio.NewReaderSize(r.fd, r.bufferSize)
	}

	return err
}

func NewReaderSize(fd io.Reader, bufferSize int, markOffset int64, ignoreReadPanic bool) io.Reader {
	if markOffset < 0 {
		markOffset = 0
	}

	if ignoreReadPanic {
		r := &ReaderV2{}
		r.fd = fd
		r.reader = bufio.NewReaderSize(fd, bufferSize)
		r.markOffset = markOffset
		r.bufferSize = bufferSize
		return r
	}

	return &Reader{fd: fd, reader: bufio.NewReaderSize(fd, bufferSize), markOffset: markOffset, bufferSize: bufferSize}
}
