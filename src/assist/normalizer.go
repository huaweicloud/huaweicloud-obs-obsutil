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
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024

	K   = 1000
	MLN = K * K
	BLN = K * K * K
)

var unitsInt64 = map[string]int64{
	"KB": KB,
	"MB": MB,
	"GB": GB,
	"TB": TB,
}

var unitsFloat64 = map[string]float64{
	"KB": KB,
	"MB": MB,
	"GB": GB,
	"TB": TB,
}

func TranslateToInt64(value string) (int64, error) {
	value = strings.TrimSpace(strings.ToUpper(value))

	for suffix, weight := range unitsInt64 {
		if strings.HasSuffix(value, suffix) {
			ret, err := StringToInt64V2(value[:len(value)-2])
			if err != nil {
				return -1, err
			}
			return ret * weight, nil
		}
	}

	if strings.HasSuffix(value, "B") {
		value = value[:len(value)-1]
	}

	return StringToInt64V2(value)
}

func TranslateToFloat64(value string) (float64, error) {
	value = strings.TrimSpace(strings.ToUpper(value))

	for suffix, weight := range unitsFloat64 {
		if strings.HasSuffix(value, suffix) {
			ret, err := StringToFloat64V2(value[:len(value)-2])
			if err != nil {
				return -1, err
			}
			return ret * weight, nil
		}
	}

	if strings.HasSuffix(value, "B") {
		value = value[:len(value)-1]
	}

	return StringToFloat64V2(value)
}

func NormalizeFilePath(fileUrl string) string {
	if absPath, err := filepath.Abs(fileUrl); err == nil {
		return absPath
	}
	return fileUrl
}

func NormalizeFileSize(fileUrl string) string {
	stat, err := os.Stat(fileUrl)
	if err == nil {
		return normalizeBytes(stat.Size(), "B")
	}
	return ""
}

func NormalizeBytes(size int64) string {
	return normalizeBytes(size, "B")
}

func normalizeBytes(size int64, unit string) string {
	if size <= 0 {
		return "0" + unit
	}

	if size < KB {
		return Int64ToString(size) + unit
	}

	_size := float64(size)

	if size < MB {
		return fmt.Sprintf("%.2fK%s", _size/KB, unit)
	}

	if size < GB {
		return fmt.Sprintf("%.2fM%s", _size/MB, unit)
	}

	if size < TB {
		return fmt.Sprintf("%.2fG%s", _size/GB, unit)
	}

	return fmt.Sprintf("%.2fT%s", _size/TB, unit)
}

func NormalizeCount(count int64, unit string) string {
	if count <= 0 {
		return "0" + unit
	}

	if count < K {
		return Int64ToString(count) + unit
	}

	_count := float64(count)

	if count < MLN {
		return fmt.Sprintf("%.2fK%s", _count/K, unit)
	}

	if count < BLN {
		return fmt.Sprintf("%.2fMln%s", _count/MLN, unit)
	}

	return fmt.Sprintf("%.2fBln%s", _count/BLN, unit)
}
