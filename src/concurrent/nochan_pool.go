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
	"runtime"
	"sync"
	"sync/atomic"
)

func NewNochanPool(maxWorkerCnt int) Pool {
	if maxWorkerCnt <= 0 {
		maxWorkerCnt = runtime.NumCPU()
	}

	pool := &NoChanPool{
		wg:     new(sync.WaitGroup),
		tokens: make(chan interface{}, maxWorkerCnt),
	}
	pool.isShutDown = 0
	pool.AddMaxWorkerCnt(int64(maxWorkerCnt))

	for i := 0; i < maxWorkerCnt; i++ {
		pool.tokens <- struct{}{}
	}

	return pool
}

func (pool *NoChanPool) acquire() {
	<-pool.tokens
}

func (pool *NoChanPool) release() {
	pool.tokens <- 1
}

func (pool *NoChanPool) execute(t Task) {
	pool.wg.Add(1)
	go func() {
		pool.acquire()
		defer func() {
			pool.release()
			pool.wg.Done()
		}()
		runTask(t)
	}()
}

func (pool *NoChanPool) ShutDown() {
	if !atomic.CompareAndSwapInt32(&pool.isShutDown, 0, 1) {
		return
	}
	pool.wg.Wait()
}

func (pool *NoChanPool) Execute(t Task) {
	if t != nil {
		pool.execute(t)
	}
}

func (pool *NoChanPool) ExecuteFunc(f func() interface{}) {
	fw := &funcWrapper{
		f: f,
	}
	pool.Execute(fw)
}

func (pool *NoChanPool) Submit(t Task) (Future, error) {
	if t == nil {
		return nil, ErrTaskInvalid
	}

	f := &FutureResult{}
	f.resultChan = make(chan interface{}, 1)
	tw := &taskWrapper{
		t: t,
		f: f,
	}

	pool.execute(tw)
	return f, nil
}

func (pool *NoChanPool) SubmitFunc(f func() interface{}) (Future, error) {
	fw := &funcWrapper{
		f: f,
	}
	return pool.Submit(fw)
}
