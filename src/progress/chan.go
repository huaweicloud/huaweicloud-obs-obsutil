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
	"assist"
	"gopkg.in/cheggaaa/pb.v2"
	"sync"
	"sync/atomic"
	"time"
)

type SingleBarChan interface {
	Start()
	SetTotalCount(count int64)
	Send(delta int)
	Send64(delta int64)
	WaitToFinished()
	SetBytes(isBytes bool)
	SetTemplate(template string)
}

type DefaultSingleBarChan struct {
	TotalCount  int64
	RefreshRate time.Duration
	Template    string
	Width       int
	IsBytes     bool
	bar         *pb.ProgressBar
	started     int32
	currentCnt  int64
	lock        sync.Mutex
}

func NewSingleBarChan() SingleBarChan {
	return &DefaultSingleBarChan{}
}

func (ch *DefaultSingleBarChan) SetTemplate(template string) {
	ch.Template = template
	if atomic.LoadInt32(&ch.started) == 1 {
		ch.bar.SetTemplateString(template)
	}
}

func (ch *DefaultSingleBarChan) SetBytes(isBytes bool) {
	ch.IsBytes = isBytes
	if atomic.LoadInt32(&ch.started) == 1 {
		ch.bar.Set(pb.Bytes, isBytes)
	}
}

func (ch *DefaultSingleBarChan) SetTotalCount(totalCount int64) {
	if totalCount >= 0 {
		atomic.StoreInt64(&ch.TotalCount, totalCount)
		if ch.bar != nil {
			ch.bar.SetTotal(totalCount)
		}
	}
}

func (ch *DefaultSingleBarChan) Start() {
	if ch.bar == nil {
		ch.lock.Lock()
		defer ch.lock.Unlock()
		if ch.bar == nil {
			bar := pb.New64(ch.TotalCount)
			if ch.RefreshRate > 0 {
				bar.SetRefreshRate(ch.RefreshRate)
			}
			if ch.Width <= 0 {
				if width, err := assist.GetTerminalWidth(); err == nil {
					if width >= 200 {
						bar.SetWidth(120)
					} else if width >= 80 {
						bar.SetWidth(79)
					} else {
						bar.SetWidth(width - 1)
					}
				}
			} else {
				bar.SetWidth(ch.Width)
			}

			if ch.Template != "" {
				bar.SetTemplateString(ch.Template)
			} else {
				bar.SetTemplateString(Standard)
			}

			if ch.IsBytes {
				bar.Set(pb.Bytes, true)
			}

			bar.Start()

			ch.bar = bar
			if cnt := atomic.LoadInt64(&ch.currentCnt); cnt > 0 {
				ch.bar.Add64(cnt)
			}
			atomic.CompareAndSwapInt32(&ch.started, 0, 1)
		}
	}
}

func (ch *DefaultSingleBarChan) Get() int64 {
	if atomic.LoadInt32(&ch.started) == 1 {
		return ch.bar.Current()
	}
	return -1
}

func (ch *DefaultSingleBarChan) Set(current int64) {
	if atomic.LoadInt32(&ch.started) == 1 {
		ch.bar.SetCurrent(current)
	}
}

func (ch *DefaultSingleBarChan) WaitToFinished() {
	if atomic.CompareAndSwapInt32(&ch.started, 1, 0) {
		ch.bar.Finish()
		ch.bar = nil
	}
}

func (ch *DefaultSingleBarChan) Send(delta int) {
	ch.Send64(int64(delta))
}

func (ch *DefaultSingleBarChan) Send64(delta int64) {
	if atomic.LoadInt32(&ch.started) == 1 {
		ch.bar.Add64(delta)
	} else {
		atomic.AddInt64(&ch.currentCnt, delta)
	}
}

type NilSingleBarChan struct{}

func NewNilSingleBarChan() SingleBarChan {
	return &NilSingleBarChan{}
}

func (ch *NilSingleBarChan) SetTemplate(template string) {}

func (ch *NilSingleBarChan) SetBytes(isBytes bool) {}

func (ch *NilSingleBarChan) SetTotalCount(totalCount int64) {}

func (ch *NilSingleBarChan) Start() {}

func (ch *NilSingleBarChan) WaitToFinished() {}

func (ch *NilSingleBarChan) Send(delta int) {}

func (ch *NilSingleBarChan) Send64(delta int64) {}
