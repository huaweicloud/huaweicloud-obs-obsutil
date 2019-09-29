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
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	capacity       int64
	rate           int64
	tokens         int64
	lastCheckpoint int64
}

func NewRateLimiter(capacity, rate int64) *RateLimiter {
	rl := &RateLimiter{}
	rl.capacity = capacity
	rl.rate = rate
	rl.lastCheckpoint = time.Now().Unix()
	return rl
}

func (rl *RateLimiter) SetCapacityAndRate(newCapacity, newRate int64) {
	rl.SetCapacity(newCapacity)
	rl.SetRate(newRate)
}

func (rl *RateLimiter) SetCapacity(newCapacity int64) {
	atomic.StoreInt64(&rl.capacity, newCapacity)
}

func (rl *RateLimiter) SetRate(newRate int64) {
	atomic.StoreInt64(&rl.rate, newRate)
}

func (rl *RateLimiter) Acquire(permits int64) int64 {
	capacity := atomic.LoadInt64(&rl.capacity)
	rate := atomic.LoadInt64(&rl.rate)

	var now int64
	var lastCheckpoint int64
	var tokens int64
	var originTokens int64
	var delta int64

	for {
		now = time.Now().Unix()
		lastCheckpoint = atomic.LoadInt64(&rl.lastCheckpoint)
		if atomic.CompareAndSwapInt64(&rl.lastCheckpoint, lastCheckpoint, now) {
			delta = (now - lastCheckpoint) * rate
			for {
				originTokens = atomic.LoadInt64(&rl.tokens)
				tokens = min(capacity, originTokens+delta)
				if tokens >= permits {
					if atomic.CompareAndSwapInt64(&rl.tokens, originTokens, tokens-permits) {
						return 0
					}
				} else {
					atomic.AddInt64(&rl.tokens, min(capacity-originTokens, delta))
					goto outer
				}
			}
		}
	}

outer:
	ret := (permits-tokens)/rate*1000 + 1
	if ret < 10 {
		ret = 10
	}
	return ret
}

func min(va, vb int64) int64 {
	if va <= vb {
		return va
	}
	return vb
}
