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
package command

import (
	"assist"
	"command/i18n"
	"os"
)

type hashCommand struct {
	defaultCommand
	algorithmType string
}

func initHash() command {
	c := &hashCommand{}
	c.key = "hash"
	c.usage = "file_url [options...]"
	c.description = "caculate the md5 or crc64 hash code of a local file"
	c.additional = true

	c.define = func() {
		c.flagSet.StringVar(&c.algorithmType, "type", "md5", "")
	}

	c.action = func() error {
		args := c.flagSet.Args()
		if len(args) <= 0 {
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		fileUrl := assist.NormalizeFilePath(args[0])
		stat, err := os.Lstat(fileUrl)
		if err != nil {
			printError(err)
			return assist.ErrFileNotFound
		}

		originFileUrl := fileUrl
		if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
			fileUrl, stat, err = assist.GetRealPath(fileUrl)
			if err != nil {
				printError(err)
				return assist.ErrFileNotFound
			}
		}

		if stat.IsDir() {
			printf("Error: The file url [%s] is a folder, not a valid file", originFileUrl)
			return assist.ErrInvalidArgs
		}

		if _err := c.checkArgs(args[1:]); _err != nil {
			printError(_err)
			return assist.ErrInvalidArgs
		}

		if c.algorithmType == c_md5 {
			ret, err := assist.Md5File(fileUrl)
			if err != nil {
				printError(err)
				return assist.ErrExecuting
			}

			printf("%-20s%s", "hex_md5:", assist.Hex(ret))
			printf("%-20s%s", "base64_md5:", assist.Base64Encode(ret))
			return nil
		}

		if c.algorithmType == c_crc64 {
			ret, err := assist.Crc64File(fileUrl)
			if err != nil {
				printError(err)
				return assist.ErrExecuting
			}
			printf("%-20s%d", "crc64_ecma:", ret)
			return nil
		}

		printf("Error: Invalid type [%s], possible values are:[%s|%s]", c.algorithmType, c_md5, c_crc64)
		return assist.ErrInvalidArgs
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("caculate the md5 or crc64 hash code of a local file"))
		printf("")
		p.Printf("Syntax:")
		printf("%2s%s", "", "obsutil hash file_url [-type=xxx]")
		printf("")

		p.Printf("Options:")
		printf("")
		printf("%2s%s", "", "-type=xxx")
		printf("%4s%s", "", p.Sprintf("the encryption algorithm type, possible values are [md5|crc64], the default value is md5"))
		printf("")
	}

	return c
}
