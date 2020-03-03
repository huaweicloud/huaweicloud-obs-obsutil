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
package progress

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"obs"
)

var errUnsupportReset = errors.New("UnsupportResetError")

type ReaderWrapper struct {
	Reader      io.Reader
	TotalCount  int64
	ReadedCount int64
}

func (rw *ReaderWrapper) Read(p []byte) (n int, err error) {
	if rw.TotalCount == 0 {
		return 0, io.EOF
	}
	n, err = rw.Reader.Read(p)
	if rw.TotalCount > 0 {
		readedOnce := int64(n)
		remainCount := rw.TotalCount - rw.ReadedCount
		if remainCount > readedOnce {
			rw.ReadedCount += readedOnce
			return
		}
		rw.ReadedCount += remainCount
		return int(remainCount), io.EOF
	}
	return
}

type CheckSumReader struct {
	ReaderWrapper
	checksum []byte
	md5Hash  hash.Hash
}

func (csr *CheckSumReader) Base64Md5() string {
	if csr.md5Hash != nil {
		if csr.checksum == nil {
			csr.checksum = csr.md5Hash.Sum(nil)
		}
		return base64.StdEncoding.EncodeToString(csr.checksum)
	}
	return ""
}

func (csr *CheckSumReader) HexMd5() string {
	if csr.md5Hash != nil {
		if csr.checksum == nil {
			csr.checksum = csr.md5Hash.Sum(nil)
		}
		return hex.EncodeToString(csr.checksum)
	}
	return ""
}

func (csr *CheckSumReader) Read(p []byte) (n int, err error) {
	n, err = csr.ReaderWrapper.Read(p)
	if csr.md5Hash == nil {
		return
	}
	if n > 0 {
		csr.md5Hash.Write(p[:n])
	}
	return
}

type SingleProgressReader struct {
	CheckSumReader
	BarCh SingleBarChan
}

func (spr *SingleProgressReader) Read(p []byte) (n int, err error) {
	n, err = spr.CheckSumReader.Read(p)
	if n > 0 {
		ctx.AddEffectiveStream(int64(n))
		ctx.AddFinishedStream(int64(n))
		if spr.BarCh != nil {
			spr.BarCh.Send(n)
		}
	}
	return
}

func NewSingleProgressReader(reader io.Reader, totalCount int64, enableChecksum bool, barCh SingleBarChan) *SingleProgressReader {
	spr := &SingleProgressReader{}
	spr.Reader = reader
	spr.TotalCount = totalCount
	if enableChecksum {
		spr.md5Hash = md5.New()
	}
	spr.BarCh = barCh
	return spr
}

type SingleProgressReaderV2 struct {
	SingleProgressReader
}

func (spr *SingleProgressReaderV2) Reset() error {
	if r, ok := spr.Reader.(obs.IRepeatable); ok {
		if err := r.Reset(); err != nil {
			return err
		}

		if spr.md5Hash != nil {
			spr.md5Hash.Reset()
		}

		if spr.ReadedCount > 0 {
			if spr.BarCh != nil {
				spr.BarCh.Send64(-spr.ReadedCount)
			}
			ctx.AddEffectiveStream(-spr.ReadedCount)
			ctx.AddFinishedStream(-spr.ReadedCount)
		}
		spr.ReadedCount = 0
		return nil
	}

	return errUnsupportReset
}

func NewSingleProgressReaderV2(reader io.Reader, totalCount int64, enableChecksum bool, barCh SingleBarChan) *SingleProgressReaderV2 {
	spr := &SingleProgressReaderV2{}
	spr.Reader = reader
	spr.TotalCount = totalCount
	if enableChecksum {
		spr.md5Hash = md5.New()
	}
	spr.BarCh = barCh
	return spr
}
