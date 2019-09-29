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
	"fmt"
	"gopkg.in/VividCortex/ewma.v1"
	"gopkg.in/cheggaaa/pb.v2"
	"gopkg.in/fatih/color.v1"
	"math"
	"time"
)

var speedAddLimit = time.Second / 2

type speed struct {
	ewma        ewma.MovingAverage
	lastStateId uint64
	prevValue   int64
	startValue  int64
	prevTime    time.Time
	startTime   time.Time
}

func (s *speed) value(state *pb.State) float64 {
	if s.ewma == nil {
		s.ewma = ewma.NewMovingAverage()
	}
	if state.IsFirst() || state.Id() < s.lastStateId {
		s.reset(state)
		return 0
	}
	if state.Id() == s.lastStateId {
		return s.ewma.Value()
	}
	if state.IsFinished() {
		return s.absValue(state)
	}
	dur := state.Time().Sub(s.prevTime)
	if dur < speedAddLimit {
		return s.ewma.Value()
	}
	current := ctx.GetEffectiveStream()
	diff := math.Abs(float64(current - s.prevValue))
	lastSpeed := diff / dur.Seconds()
	s.prevTime = state.Time()
	s.prevValue = current
	s.lastStateId = state.Id()
	s.ewma.Add(lastSpeed)
	return s.ewma.Value()
}

func (s *speed) reset(state *pb.State) {
	s.lastStateId = state.Id()
	s.startTime = state.Time()
	s.prevTime = state.Time()
	s.startValue = ctx.GetEffectiveStream()
	s.prevValue = s.startValue
	s.ewma = ewma.NewMovingAverage()
}

func (s *speed) absValue(state *pb.State) float64 {
	if dur := state.Time().Truncate(time.Millisecond).Sub(s.startTime).Truncate(time.Millisecond); dur > 0 {
		return float64(ctx.GetEffectiveStream()) / dur.Seconds()
	}
	return 0
}

func (s *speed) eval(state *pb.State, args ...string) string {
	sp := s.value(state)
	if sp == 0 {
		return "?/s"
	}
	return fmt.Sprintf("%s/s", assist.NormalizeBytes(int64(assist.Round(sp))))
}

func speedBarColorFunc(colorful bool) func(a ...interface{}) string {
	if !colorful {
		return color.New(color.FgWhite).SprintFunc()
	}
	return func(s ...interface{}) string {
		rate := float64(ctx.GetFinishedStream()) / float64(ctx.GetTotalStream())
		var attr color.Attribute
		if rate <= 0.3 {
			attr = color.FgRed
		} else if rate <= 0.6 {
			attr = color.FgYellow
		} else {
			attr = color.FgGreen
		}
		return color.New(attr).Sprint(s...)
	}
}

func speedColorFunc(colorful bool, sp *speed) func(a ...interface{}) string {
	if !colorful {
		return color.New(color.FgWhite).SprintFunc()
	}
	return func(s ...interface{}) string {
		if sp == nil || sp.ewma == nil {
			return color.New(color.Attribute(color.FgWhite)).Sprint(s...)
		}
		rate := sp.ewma.Value()
		var attr color.Attribute
		if rate <= assist.KB {
			attr = color.FgRed
		} else if rate <= assist.MB {
			attr = color.FgYellow
		} else if rate <= assist.GB {
			attr = color.FgCyan
		} else {
			attr = color.FgGreen
		}
		return color.New(attr).Sprint(s...)
	}
}
