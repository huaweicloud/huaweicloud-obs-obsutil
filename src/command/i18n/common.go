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
package i18n

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sync"
)

var printers map[language.Tag]*PrinterWrapper
var currentTag = language.BritishEnglish
var tagLock sync.Mutex

func SetI18nStrings() {
	printers = make(map[language.Tag]*PrinterWrapper)
	(&messageBuilderEn{}).buildMessage(printers)
	(&messageBuilderCn{}).buildMessage(printers)
}

type PrinterWrapper struct {
	p *message.Printer
}

func newPrinterWrapper(p *message.Printer) *PrinterWrapper {
	pw := &PrinterWrapper{
		p: p,
	}
	return pw
}

func (pw PrinterWrapper) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = pw.p.Printf(format, a...)
	fmt.Println()
	return
}

func (pw PrinterWrapper) Sprintf(format string, a ...interface{}) string {
	return pw.p.Sprintf(format, a...)
}

func getPrinter(tag language.Tag) *PrinterWrapper {
	if printer, ok := printers[tag]; ok {
		return printer
	}
	return printers[language.BritishEnglish]
}

func SetCurrentTag(tag language.Tag) {
	tagLock.Lock()
	currentTag = tag
	tagLock.Unlock()
}

func SetCurrentLanguage(lan string) {
	if lan == "chinese" {
		SetCurrentTag(language.Chinese)
	} else {
		SetCurrentTag(language.BritishEnglish)
	}
}

func GetCurrentPrinter() *PrinterWrapper {
	tagLock.Lock()
	defer tagLock.Unlock()
	return getPrinter(currentTag)
}
