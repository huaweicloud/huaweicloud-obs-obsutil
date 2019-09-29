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
	"io"
)

type MultiWritersReader struct {
	reader  io.Reader
	writers []io.Writer
}

func (mwr *MultiWritersReader) Read(p []byte) (n int, err error) {
	n, err = mwr.reader.Read(p)
	if mwr.writers != nil {
		if n > 0 {
			for _, writer := range mwr.writers {
				if writer != nil {
					writer.Write(p[:n])
				}
			}
		}
	}
	return
}

func Wrap(reader io.Reader, writers ...io.Writer) io.Reader {
	length := len(writers)
	if length == 0 {
		return reader
	}

	_writers := make([]io.Writer, 0, length)
	for _, writer := range writers {
		_writers = append(_writers, writer)
	}

	return &MultiWritersReader{reader: reader, writers: _writers}
}
