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
package concurrent

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Future interface {
	Get() interface{}
}

type FutureResult struct {
	result     interface{}
	resultChan chan interface{}
	lock       sync.Mutex
}

type panicResult struct {
	presult interface{}
}

func (f *FutureResult) checkPanic() interface{} {
	if r, ok := f.result.(panicResult); ok {
		panic(r.presult)
	}
	return f.result
}

func (f *FutureResult) Get() interface{} {
	if f.resultChan == nil {
		return f.checkPanic()
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.resultChan == nil {
		return f.checkPanic()
	}

	f.result = <-f.resultChan
	close(f.resultChan)
	f.resultChan = nil
	return f.checkPanic()
}

type Task interface {
	Run() interface{}
}

type funcWrapper struct {
	f func() interface{}
}

func (fw *funcWrapper) Run() interface{} {
	if fw.f != nil {
		return fw.f()
	}
	return nil
}

type taskWrapper struct {
	t Task
	f *FutureResult
}

func (tw *taskWrapper) Run() interface{} {
	if tw.t != nil {
		return tw.t.Run()
	}
	return nil
}

type signalTask struct {
	id string
}

func (signalTask) Run() interface{} {
	return nil
}

type worker struct {
	name      string
	taskQueue chan Task
	wg        *sync.WaitGroup
	pool      *RoutinePool
}

func runTask(t Task) {
	if tw, ok := t.(*taskWrapper); ok {
		defer func() {
			if r := recover(); r != nil {
				tw.f.resultChan <- panicResult{
					presult: r,
				}
			}
		}()
		ret := t.Run()
		tw.f.resultChan <- ret
	} else {
		t.Run()
	}
}

func (*worker) runTask(t Task) {
	runTask(t)
}

func (w *worker) start() {
	go func() {
		defer func() {
			if w.wg != nil {
				w.wg.Done()
			}
		}()
		for {
			task, ok := <-w.taskQueue
			if !ok {
				break
			}
			w.pool.AddCurrentWorkingCnt(1)
			w.runTask(task)
			w.pool.AddCurrentWorkingCnt(-1)
			if w.pool.autoTuneWorker(w) {
				break
			}
		}
	}()
}

func (w *worker) release() {
	w.taskQueue = nil
	w.wg = nil
	w.pool = nil
}

type Pool interface {
	ShutDown()
	Submit(t Task) (Future, error)
	SubmitFunc(f func() interface{}) (Future, error)
	Execute(t Task)
	ExecuteFunc(f func() interface{})
	GetMaxWorkerCnt() int64
	AddMaxWorkerCnt(value int64) int64
	GetCurrentWorkingCnt() int64
	AddCurrentWorkingCnt(value int64) int64
	GetWorkerCnt() int64
	AddWorkerCnt(value int64) int64
	EnableAutoTune()
}

type basicPool struct {
	maxWorkerCnt      int64
	workerCnt         int64
	currentWorkingCnt int64
	isShutDown        int32
}

var ErrTaskInvalid = errors.New("Task is nil")

func (pool *basicPool) GetCurrentWorkingCnt() int64 {
	return atomic.LoadInt64(&pool.currentWorkingCnt)
}

func (pool *basicPool) AddCurrentWorkingCnt(value int64) int64 {
	return atomic.AddInt64(&pool.currentWorkingCnt, value)
}

func (pool *basicPool) GetWorkerCnt() int64 {
	return atomic.LoadInt64(&pool.workerCnt)
}

func (pool *basicPool) AddWorkerCnt(value int64) int64 {
	return atomic.AddInt64(&pool.workerCnt, value)
}

func (pool *basicPool) GetMaxWorkerCnt() int64 {
	return atomic.LoadInt64(&pool.maxWorkerCnt)
}

func (pool *basicPool) AddMaxWorkerCnt(value int64) int64 {
	return atomic.AddInt64(&pool.maxWorkerCnt, value)
}

func (pool *basicPool) CompareAndSwapCurrentWorkingCnt(oldValue, newValue int64) bool {
	return atomic.CompareAndSwapInt64(&pool.currentWorkingCnt, oldValue, newValue)
}

func (pool *basicPool) EnableAutoTune() {

}

type RoutinePool struct {
	basicPool
	taskQueue     chan Task
	dispatchQueue chan Task
	workers       map[string]*worker
	cacheCnt      int
	wg            *sync.WaitGroup
	lock          *sync.Mutex
	shutDownWg    *sync.WaitGroup
	autoTune      int32
}

type NoChanPool struct {
	basicPool
	wg     *sync.WaitGroup
	tokens chan interface{}
}
