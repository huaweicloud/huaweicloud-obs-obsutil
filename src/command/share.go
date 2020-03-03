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
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var aesShareKeyPrefix = "mockaeskey0123456789012345"
var aesShareIv = "mockaesiv0123456"
var authorizationCodeRegex = regexp.MustCompile("http.+?token=(.+)")

func SetAesShareIv(_aesShareKeyPrefix, _aesShareIv string) {
	if _aesShareKeyPrefix != "" {
		aesShareKeyPrefix = _aesShareKeyPrefix
	}

	if _aesShareIv != "" {
		aesShareIv = _aesShareIv
	}
}

type shareCommand struct {
	parallelContextCommand
	accessCode string
}

func (c *shareCommand) init() {
	c.recursiveCommand.init()
	c.flagSet.StringVar(&c.accessCode, "ac", "", "")
}

func (c *shareCommand) prepareAccessUrl(authorizationCode string, args []string) (parsedUrl *url.URL, allowedPrefix string, err error) {
	if err = c.flagSet.Parse(args); err != nil {
		c.showHelp()
		return
	}

	if len(c.flagSet.Args()) >= 1 {
		c.showHelp()
		err = fmt.Errorf("Invalid args \"%v\", please refer to help doc", c.flagSet.Args())
		return
	}

	if c.accessCode == "" {
		c.accessCode, err = getUserInput("Please input your access code:")
		if err != nil {
			return
		}
	}

	accessUrl, _err := c.checkAuthorizationAndAccess(authorizationCode, c.accessCode)
	if _err != nil {
		err = _err
		return
	}

	parsedUrl, err = url.Parse(accessUrl)

	if err != nil {
		return
	}

	allowedPrefix = parsedUrl.Query().Get("prefix")
	c.printAuthorizedPrefix(allowedPrefix)
	return
}

func (c *shareCommand) printAuthorizedPrefix(prefix string) {
	if prefix == "" {
		printf("The authorized prefix is empty, all the content of bucket can be accessed\n")
	} else {
		printf("The authorized prefix is [%s]\n", prefix)
	}
}

func (c *shareCommand) checkAuthorizationAndAccess(authorizationCode, accessCode string) (accessUrl string, err error) {
	if l := len(accessCode); l != 6 {
		err = fmt.Errorf("Invalid access code, the length [%d] does not equal to 6", l)
		return
	}

	if strings.HasPrefix(authorizationCode, "file://") {
		fd, rerr := os.Open(authorizationCode[len("file://"):])
		if rerr != nil {
			err = rerr
			return
		}
		defer fd.Close()
		ret, rerr := ioutil.ReadAll(fd)
		if rerr != nil {
			err = rerr
			return
		}

		authorizationCode = strings.TrimSpace(assist.BytesToString(ret))
	}

	if authorizationCodeRegex.MatchString(authorizationCode) {
		ret := authorizationCodeRegex.FindStringSubmatch(authorizationCode)
		if len(ret) != 2 {
			err = errors.New("Invalid authorization code, can not parse the authorization code")
			return
		}
		authorizationCode = ret[1]
	} else if strings.HasPrefix(authorizationCode, "token=") {
		authorizationCode = authorizationCode[len("token="):]
	} else if strings.HasPrefix(authorizationCode, "?token=") {
		authorizationCode = authorizationCode[len("?token="):]
	}

	if retByte, retErr := AesDecrypt(authorizationCode, assist.StringToBytes(aesShareKeyPrefix+accessCode), assist.StringToBytes(aesShareIv)); retErr != nil {
		err = fmt.Errorf("Invalid authorization code or access code, can not decrypt the authorization code, [%s]", retErr.Error())
		return
	} else if _accessUrl, _err := assist.Base64Decode(assist.BytesToString(retByte)); _err != nil {
		err = fmt.Errorf("Invalid authorization code or access code, can not decode the authorization code, [%s]", _err.Error())
		return
	} else {
		accessUrl = assist.BytesToString(_accessUrl)
	}

	if accessUrl == "" {
		err = errors.New("Invalid authorization code or access code")
		return
	}

	doLog(LEVEL_INFO, "The access url from authorization code and access code is [%s]", accessUrl)
	return
}
