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
	"errors"
	"os"
)

var ErrFileNotFound = errors.New("FileNotFoundError")
var ErrTaskNotFound = errors.New("TaskNotFoundError")
var ErrInvalidArgs = errors.New("InvalidArgsError")
var ErrInitializing = errors.New("InitializingError")
var ErrInterrupted = errors.New("InterruptedError")
var ErrCheckBucketStatus = errors.New("CheckBucketStatusError")
var ErrExecuting = errors.New("ExecutedError")
var ErrUncompeleted = errors.New("UncompeletedError")
var ErrUnsupported = errors.New("UnsupportedError")

func CheckErrorAndExit(err error) {
	if err == nil {
		os.Exit(0)
		return
	}

	if err == ErrFileNotFound {
		os.Exit(1)
		return
	}

	if err == ErrTaskNotFound {
		os.Exit(2)
		return
	}

	if err == ErrInvalidArgs {
		os.Exit(3)
		return
	}

	if err == ErrCheckBucketStatus {
		os.Exit(4)
		return
	}

	if err == ErrInitializing {
		os.Exit(5)
		return
	}

	if err == ErrExecuting {
		os.Exit(6)
		return
	}

	if err == ErrUnsupported {
		os.Exit(7)
		return
	}

	if err == ErrUncompeleted {
		os.Exit(8)
		return
	}

	if err == ErrInterrupted {
		os.Exit(9)
		return
	}

	if err != nil {
		os.Exit(-1)
		return
	}
}
