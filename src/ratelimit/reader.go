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
package ratelimit

import (
	"errors"
	"io"
	"obs"
	"time"
)

var errUnsupportReset = errors.New("UnsupportResetError")

type RateLimitReader struct {
	Reader  io.Reader
	Limiter *RateLimiter
}

func (rlr *RateLimitReader) Read(p []byte) (n int, err error) {
	count := len(p)
	var ret int64
	for {
		ret = rlr.Limiter.Acquire(int64(count))
		if ret == 0 {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(ret))
	}
	return rlr.Reader.Read(p)
}

func NewRateLimitReader(reader io.Reader, capacity, rate int64) *RateLimitReader {
	rlr := &RateLimitReader{}
	rlr.Reader = reader
	rlr.Limiter = NewRateLimiter(capacity, rate)
	return rlr
}

func NewRateLimitReaderWithLimiter(reader io.Reader, limiter *RateLimiter) *RateLimitReader {
	rlr := &RateLimitReader{}
	rlr.Reader = reader
	rlr.Limiter = limiter
	return rlr
}

type RateLimitReaderV2 struct {
	Reader  io.Reader
	Limiter *RateLimiter
}

func (rlr *RateLimitReaderV2) Read(p []byte) (n int, err error) {
	count := len(p)
	var ret int64
	for {
		ret = rlr.Limiter.Acquire(int64(count))
		if ret == 0 {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(ret))
	}
	return rlr.Reader.Read(p)
}

func (rlr *RateLimitReaderV2) Reset() error {
	if r, ok := rlr.Reader.(obs.IRepeatable); ok {
		return r.Reset()
	}

	return errUnsupportReset
}

func NewRateLimitReaderV2(reader io.Reader, capacity, rate int64) *RateLimitReaderV2 {
	rlr := &RateLimitReaderV2{}
	rlr.Reader = reader
	rlr.Limiter = NewRateLimiter(capacity, rate)
	return rlr
}

func NewRateLimitReaderWithLimiterV2(reader io.Reader, limiter *RateLimiter) *RateLimitReaderV2 {
	rlr := &RateLimitReaderV2{}
	rlr.Reader = reader
	rlr.Limiter = limiter
	return rlr
}
