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
	"sync/atomic"
)

type context struct {
	totalCount      int64
	succeedCount    int64
	failedCount     int64
	warningCount    int64
	resumeCount     int64
	finishedCount   int64
	totalStream     int64
	finishedStream  int64
	transaction     int64
	effectiveStream int64
	succeedStream   int64
}

func (ctx *context) GetTotalCount() int64 {
	return atomic.LoadInt64(&ctx.totalCount)
}

func (ctx *context) SetTotalCount(value int64) {
	atomic.StoreInt64(&ctx.totalCount, value)
}

func (ctx *context) GetSucceedCount() int64 {
	return atomic.LoadInt64(&ctx.succeedCount)
}

func (ctx *context) SetSucceedCount(value int64) {
	atomic.StoreInt64(&ctx.succeedCount, value)
}

func (ctx *context) AddSucceedCount(value int64) int64 {
	return atomic.AddInt64(&ctx.succeedCount, value)
}

func (ctx *context) GetResumeCount() int64 {
	return atomic.LoadInt64(&ctx.resumeCount)
}

func (ctx *context) SetResumeCount(value int64) {
	atomic.StoreInt64(&ctx.resumeCount, value)
}

func (ctx *context) AddResumeCount(value int64) int64 {
	return atomic.AddInt64(&ctx.resumeCount, value)
}

func (ctx *context) GetFailedCount() int64 {
	return atomic.LoadInt64(&ctx.failedCount)
}

func (ctx *context) SetFailedCount(value int64) {
	atomic.StoreInt64(&ctx.failedCount, value)
}

func (ctx *context) AddFailedCount(value int64) int64 {
	return atomic.AddInt64(&ctx.failedCount, value)
}

func (ctx *context) GetWarningCount() int64 {
	return atomic.LoadInt64(&ctx.warningCount)
}

func (ctx *context) SetWarningCount(value int64) {
	atomic.StoreInt64(&ctx.warningCount, value)
}

func (ctx *context) AddWarningCount(value int64) int64 {
	return atomic.AddInt64(&ctx.warningCount, value)
}

func (ctx *context) CompareAndSetWarningCount(oldValue, newValue int64) bool {
	return atomic.CompareAndSwapInt64(&ctx.warningCount, oldValue, newValue)
}

func (ctx *context) GetFinishedCount() int64 {
	return atomic.LoadInt64(&ctx.finishedCount)
}

func (ctx *context) SetFinishedCount(value int64) {
	atomic.StoreInt64(&ctx.finishedCount, value)
}

func (ctx *context) AddFinishedCount(value int64) int64 {
	return atomic.AddInt64(&ctx.finishedCount, value)
}

func (ctx *context) GetTotalStream() int64 {
	return atomic.LoadInt64(&ctx.totalStream)
}

func (ctx *context) SetTotalStream(value int64) {
	atomic.StoreInt64(&ctx.totalStream, value)
}

func (ctx *context) GetFinishedStream() int64 {
	return atomic.LoadInt64(&ctx.finishedStream)
}

func (ctx *context) SetFinishedStream(value int64) {
	atomic.StoreInt64(&ctx.finishedStream, value)
}

func (ctx *context) AddFinishedStream(value int64) int64 {
	return atomic.AddInt64(&ctx.finishedStream, value)
}

func (ctx *context) GetTransaction() int64 {
	return atomic.LoadInt64(&ctx.transaction)
}

func (ctx *context) AddTransaction(value int64) int64 {
	return atomic.AddInt64(&ctx.transaction, value)
}

func (ctx *context) GetEffectiveStream() int64 {
	return atomic.LoadInt64(&ctx.effectiveStream)
}

func (ctx *context) AddEffectiveStream(value int64) int64 {
	return atomic.AddInt64(&ctx.effectiveStream, value)
}

func (ctx *context) GetSucceedStream() int64 {
	return atomic.LoadInt64(&ctx.succeedStream)
}

func (ctx *context) AddSucceedStream(value int64) int64 {
	return atomic.AddInt64(&ctx.succeedStream, value)
}

func (ctx *context) Reset() {
	atomic.StoreInt64(&ctx.totalCount, -1)
	atomic.StoreInt64(&ctx.totalStream, -1)
	atomic.StoreInt64(&ctx.succeedCount, 0)
	atomic.StoreInt64(&ctx.failedCount, 0)
	atomic.StoreInt64(&ctx.warningCount, 0)
	atomic.StoreInt64(&ctx.resumeCount, 0)
	atomic.StoreInt64(&ctx.finishedCount, 0)
	atomic.StoreInt64(&ctx.finishedStream, 0)
	atomic.StoreInt64(&ctx.transaction, 0)
	atomic.StoreInt64(&ctx.effectiveStream, 0)
	atomic.StoreInt64(&ctx.succeedStream, 0)
}

var ctx = &context{
	totalCount:  -1,
	totalStream: -1,
}

func ResetContext() {
	ctx.Reset()
}

func AddTransaction(value int64) int64 {
	return ctx.AddTransaction(value)
}

func GetTransaction() int64 {
	return ctx.GetTransaction()
}

func AddEffectiveStream(value int64) int64 {
	return ctx.AddEffectiveStream(value)
}

func GetEffectiveStream() int64 {
	return ctx.GetEffectiveStream()
}

func SetTotalCount(value int64) {
	ctx.SetTotalCount(value)
}

func AddFinishedCount(value int64) int64 {
	return ctx.AddFinishedCount(value)
}

func AddSucceedCount(value int64) int64 {
	return ctx.AddSucceedCount(value)
}

func GetSucceedCount() int64 {
	return ctx.GetSucceedCount()
}

func AddFailedCount(value int64) int64 {
	return ctx.AddFailedCount(value)
}

func GetFailedCount() int64 {
	return ctx.GetFailedCount()
}

func AddWarningCount(value int64) int64 {
	return ctx.AddWarningCount(value)
}

func GetWarningCount() int64 {
	return ctx.GetWarningCount()
}

func CompareAndSetWarningCount(oldValue, newValue int64) bool {
	return ctx.CompareAndSetWarningCount(oldValue, newValue)
}

func GetResumeCount() int64 {
	return ctx.GetResumeCount()
}

func AddResumeCount(value int64) int64 {
	return ctx.AddResumeCount(value)
}

func SetTotalStream(value int64) {
	ctx.SetTotalStream(value)
}

func GetTotalStream() int64 {
	return ctx.GetTotalStream()
}

func GetFinishedStream() int64 {
	return ctx.GetFinishedStream()
}

func AddFinishedStream(value int64) int64 {
	return ctx.AddFinishedStream(value)
}

func GetSucceedStream() int64 {
	return ctx.GetSucceedStream()
}

func AddSucceedStream(value int64) int64 {
	return ctx.AddSucceedStream(value)
}
