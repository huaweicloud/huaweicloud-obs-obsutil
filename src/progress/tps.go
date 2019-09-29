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
	"fmt"
	"gopkg.in/cheggaaa/pb.v2"
	"gopkg.in/fatih/color.v1"
	"time"
)

type tps struct {
	startTime time.Time
}

func (t *tps) absValue(state *pb.State) float64 {
	if dur := state.Time().Sub(t.startTime); dur > 0 {
		return float64(ctx.GetTransaction()) / dur.Seconds()
	}
	return 0
}

func (t *tps) eval(state *pb.State, args ...string) string {
	if state.IsFirst() {
		t.startTime = state.Time()
	}
	return fmt.Sprintf("tps:%.2f", t.absValue(state))
}

func tpsBarColorFunc(colorful bool) func(a ...interface{}) string {
	if !colorful {
		return color.New(color.FgWhite).SprintFunc()
	}
	return func(s ...interface{}) string {
		rate := float64(ctx.GetFinishedCount()) / float64(ctx.GetTotalCount())
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
