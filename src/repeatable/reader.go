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
	"io"
	"obs"
)

var errUnsupportReset = errors.New("UnsupportResetError")

type Reader struct {
	fd         io.Reader
	markOffset int64
	reader     *bufio.Reader
}

func (r *Reader) Read(p []byte) (n int, err error) {
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
		r.reader.Reset(r.fd)
	}

	return err
}

func NewReaderSize(fd io.Reader, bufferSize int, markOffset int64) *Reader {
	if markOffset < 0 {
		markOffset = 0
	}
	return &Reader{fd: fd, reader: bufio.NewReaderSize(fd, bufferSize), markOffset: markOffset}
}
