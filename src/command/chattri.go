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
	"concurrent"
	"encoding/xml"
	"fmt"
	"obs"
	"progress"
	"strings"
	"time"
)

type owner struct {
	XMLName xml.Name `xml:"Owner"`
	ID      string   `xml:"ID,omitempty"`
}

type grantee struct {
	XMLName xml.Name `xml:"Grantee"`
	ID      string   `xml:"ID,omitempty"`
	Canned  string   `xml:"Canned,omitempty"`
}

type grant struct {
	XMLName    xml.Name `xml:"Grant"`
	Grantee    grantee  `xml:"Grantee"`
	Permission string   `xml:"Permission"`
	Delivered  bool     `xml:"Delivered"`
}

type accessControlPolicy struct {
	XMLName xml.Name `xml:"AccessControlPolicy"`
	Owner   owner    `xml:"Owner"`
	Grants  []grant  `xml:"AccessControlList>Grant"`
}

type chattriCommand struct {
	recursiveCommand
	sc        string
	acl       string
	version   bool
	versionId string
	aclXml    string
}

func (c *chattriCommand) setBucketStorageClass(bucket string, storageClassType obs.StorageClassType) bool {
	input := &obs.SetBucketStoragePolicyInput{}
	input.Bucket = bucket
	input.StorageClass = storageClassType
	input.RequestPayer = c.payer
	output, err := obsClient.SetBucketStoragePolicy(input)
	if err == nil {
		printf("Set the default storage class of bucket [%s] to [%s] successfully, request id [%s]", bucket, c.sc, output.RequestId)
		doLog(LEVEL_INFO, "Set the default storage class of bucket [%s] to [%s] successfully, request id [%s]", bucket, c.sc, output.RequestId)
	} else {
		logError(err, LEVEL_INFO, fmt.Sprintf("Set the default storage class of bucket [%s] to [%s] failed", bucket, c.sc))
	}
	return err == nil
}

func (c *chattriCommand) setObjectStorageClass(bucket, key, versionId string, storageClassType obs.StorageClassType, batchFlag int) bool {
	abortHandler := func() {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		c.failedLogger.doRecord("Bucket [%s], key [%s], %s, n/a, n/a, error code [%s], error message [%s], n/a", bucket, key, versionIdStr,
			"AbortError", "Task is aborted")
	}
	actionFunc := func() (output *obs.BaseModel, err error) {
		input := &obs.SetObjectMetadataInput{}
		input.Bucket = bucket
		input.Key = key
		input.StorageClass = storageClassType
		input.MetadataDirective = obs.ReplaceMetadataNew
		input.VersionId = versionId
		input.RequestPayer = c.payer
		if setObjectMetadataOutput, setErr := obsClient.SetObjectMetadata(input); setErr != nil {
			err = setErr
		} else {
			output = &setObjectMetadataOutput.BaseModel
		}
		return
	}
	recordHandler := func(cost int64, output *obs.BaseModel, err error) {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		if err == nil {
			c.succeedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], request id [%s]", bucket, key, versionIdStr, cost, output.StatusCode, output.RequestId)
		} else {
			status, code, message, requestId := c.checkAbort(err, 401, 405)
			c.failedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", bucket, key,
				versionIdStr, cost, status, code, message, requestId)
		}
	}
	printHandler := func(cost int64, output *obs.BaseModel, err error) {
		if versionId != "" {
			if err == nil {
				printf("Set the storage class of object [%s] with version id [%s] in the bucket [%s] to [%s] successfully, request id [%s]", key, versionId, bucket, c.sc, output.RequestId)
				doLog(LEVEL_INFO, "Set the storage class of object [%s] in the bucket [%s] to [%s] successfully, request id [%s]", key, bucket, c.sc, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Set the storage class of object [%s] with version id [%s] in the bucket [%s] to [%s] failed", key, versionId, bucket, c.sc))
			}

		} else {
			if err == nil {
				printf("Set the storage class of object [%s] in the bucket [%s] to [%s] successfully, request id [%s]", key, bucket, c.sc, output.RequestId)
				doLog(LEVEL_INFO, "Set the storage class of object [%s] in the bucket [%s] to [%s] successfully, request id [%s]", key, bucket, c.sc, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Set the storage class of object [%s] in the bucket [%s] to [%s] failed", key, bucket, c.sc))
			}
		}
	}

	return c.simpleAction(batchFlag, abortHandler, actionFunc, recordHandler, printHandler)
}

func (c *chattriCommand) transGrants(acp *accessControlPolicy) []obs.Grant {
	grants := make([]obs.Grant, 0, len(acp.Grants))
	for _, grant := range acp.Grants {
		g := obs.Grant{}
		if userId := strings.TrimSpace(grant.Grantee.ID); userId != "" {
			g.Grantee.ID = userId
			g.Grantee.Type = obs.GranteeUser
		} else if canned := strings.TrimSpace(grant.Grantee.Canned); canned == aclEveryOne {
			g.Grantee.URI = obs.GroupAllUsers
			g.Grantee.Type = obs.GranteeGroup
		}
		g.Permission = obs.PermissionType(strings.ToUpper(grant.Permission))
		grants = append(grants, g)
	}
	return grants
}

func (c *chattriCommand) transBucketAcl(input *obs.SetBucketAclInput, acp *accessControlPolicy) {
	ownerId := strings.TrimSpace(acp.Owner.ID)
	if ownerId == "" {
		if output, err := obsClient.GetBucketAcl(input.Bucket, obs.WithReqPaymentHeader(c.payer)); err == nil {
			ownerId = output.Owner.ID
		}
	}
	input.Owner.ID = ownerId
	input.Grants = c.transGrants(acp)
}

func (c *chattriCommand) transObjectAcl(input *obs.SetObjectAclInput, acp *accessControlPolicy) {
	ownerId := strings.TrimSpace(acp.Owner.ID)
	if ownerId == "" {
		getObjectAclInput := &obs.GetObjectAclInput{}
		getObjectAclInput.Bucket = input.Bucket
		getObjectAclInput.Key = input.Key
		getObjectAclInput.VersionId = input.VersionId
		getObjectAclInput.RequestPayer = c.payer
		if output, err := obsClient.GetObjectAcl(getObjectAclInput); err == nil {
			ownerId = output.Owner.ID
		}
	}
	input.Owner.ID = ownerId
	input.Grants = c.transGrants(acp)
}

func (c *chattriCommand) setBucketAcl(bucket string, aclType obs.AclType, acp *accessControlPolicy) bool {
	input := &obs.SetBucketAclInput{}
	input.Bucket = bucket
	if acp != nil {
		c.transBucketAcl(input, acp)
	} else {
		input.ACL = aclType
	}
	input.RequestPayer = c.payer

	output, err := obsClient.SetBucketAcl(input)
	if err == nil {
		if aclType != "" {
			printf("Set the acl of bucket [%s] to [%s] successfully, request id [%s]", bucket, aclType, output.RequestId)
			doLog(LEVEL_INFO, "Set the acl of bucket [%s] to [%s] successfully, request id [%s]", bucket, aclType, output.RequestId)
		} else {
			printf("Set the acl of bucket [%s] successfully, request id [%s]", bucket, output.RequestId)
			doLog(LEVEL_INFO, "Set the acl of bucket [%s] successfully, request id [%s]", bucket, output.RequestId)
		}
	} else {
		if aclType != "" {
			logError(err, LEVEL_INFO, fmt.Sprintf("Set the acl of bucket [%s] to [%s] failed", bucket, aclType))
		} else {
			logError(err, LEVEL_INFO, fmt.Sprintf("Set the acl of bucket [%s] failed", bucket))
		}
	}
	return err == nil
}

func (c *chattriCommand) setObjectAcl(bucket, key, versionId string, aclType obs.AclType, acp *accessControlPolicy, batchFlag int) bool {
	abortHandler := func() {

		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}

		c.failedLogger.doRecord("Bucket [%s], key [%s], %s, n/a, n/a, error code [%s], error message [%s], n/a", bucket, key, versionIdStr,
			"AbortError", "Task is aborted")
	}
	actionFunc := func() (output *obs.BaseModel, err error) {
		input := &obs.SetObjectAclInput{}
		input.Bucket = bucket
		input.Key = key
		input.VersionId = versionId
		if acp != nil {
			c.transObjectAcl(input, acp)
		} else {
			input.ACL = aclType
		}
		input.RequestPayer = c.payer
		return obsClient.SetObjectAcl(input)
	}
	recordHandler := func(cost int64, output *obs.BaseModel, err error) {
		versionIdStr := c_na
		if versionId != "" {
			versionIdStr = fmt.Sprintf("version id [%s]", versionId)
		}
		if err == nil {
			c.succeedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], request id [%s]", bucket, key, versionIdStr, cost, output.StatusCode, output.RequestId)
		} else {
			status, code, message, requestId := c.checkAbort(err, 401, 405)
			c.failedLogger.doRecord("Bucket [%s], key [%s], %s, cost [%d], status [%d], error code [%s], error message [%s], request id [%s]", bucket, key,
				versionIdStr, cost, status, code, message, requestId)
		}
	}
	printHandler := func(cost int64, output *obs.BaseModel, err error) {
		aclMessage := ""
		if aclType != "" {
			aclMessage = fmt.Sprintf(" to [%s]", aclType)
		}

		if versionId != "" {
			if err == nil {
				printf("Set the acl of object [%s] with version id [%s] in the bucket [%s]%s successfully, request id [%s]", key, versionId, bucket, aclMessage, output.RequestId)
				doLog(LEVEL_INFO, "Set the acl of object [%s] with version id [%s] in the bucket [%s]%s successfully, request id [%s]", key, versionId, bucket, aclMessage, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Set the acl of object [%s] in the bucket [%s]%s failed", key, bucket, c.acl))
			}
		} else {
			if err == nil {
				printf("Set the acl of object [%s] in the bucket [%s]%s successfully, request id [%s]", key, bucket, aclMessage, output.RequestId)
				doLog(LEVEL_INFO, "Set the acl of object [%s] in the bucket [%s]%s successfully, request id [%s]", key, bucket, aclMessage, output.RequestId)
			} else {
				logError(err, LEVEL_INFO, fmt.Sprintf("Set the acl of object [%s] in the bucket [%s]%s failed", key, bucket, aclMessage))
			}
		}
	}

	return c.simpleAction(batchFlag, abortHandler, actionFunc, recordHandler, printHandler)
}

func (c *chattriCommand) submitSetStorageClassTask(bucket, prefix string, storageClassType obs.StorageClassType,
	pool concurrent.Pool, ch progress.SingleBarChan) (totalCnt int64, hasListError error) {

	if c.version {
		actionFunc := func(bucket, key, versionId string) bool {
			return c.setObjectStorageClass(bucket, key, versionId, storageClassType, 2)
		}
		isSkipFunc := func(version obs.Version) bool {
			return version.StorageClass == storageClassType
		}
		return c.submitListVersionsTask(bucket, prefix, "change the storage class", pool, ch, actionFunc, isSkipFunc, nil, false)
	}

	actionFunc := func(bucket, key string) bool {
		return c.setObjectStorageClass(bucket, key, "", storageClassType, 2)
	}
	isSkipFunc := func(content obs.Content) bool {
		return content.StorageClass == storageClassType
	}
	return c.submitListObjectsTask(bucket, prefix, "change the storage class", pool, ch, actionFunc, isSkipFunc)
}

func (c *chattriCommand) recordStartFunc() time.Time {
	start := c.recordStart()
	c.succeedLogger.doRecord("[%s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "request id")
	c.failedLogger.doRecord("[%s, %s, %s, %s, %s, %s, %s, %s]", "bucket name", "object key", "version id", "cost(ms)", "status code", "error code", "error message", "request id")
	return start
}

func (c *chattriCommand) setObjectsStorageClass(bucket, prefix string, storageClassType obs.StorageClassType) error {
	submitFunc := func(pool concurrent.Pool, ch progress.SingleBarChan) (int64, error) {
		return c.submitSetStorageClassTask(bucket, prefix, storageClassType, pool, ch)
	}

	errorHandleFunc := func(hasListError error) {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to change the storage class failed", bucket))
	}

	return c.recursiveAction(bucket, prefix, submitFunc, errorHandleFunc, c.recordStartFunc, true)
}

func (c *chattriCommand) submitSetAclTask(bucket, prefix string, aclType obs.AclType, acp *accessControlPolicy,
	pool concurrent.Pool, ch progress.SingleBarChan) (totalCnt int64, hasListError error) {

	if c.version {
		actionFunc := func(bucket, key, versionId string) bool {
			return c.setObjectAcl(bucket, key, versionId, aclType, acp, 2)
		}
		return c.submitListVersionsTask(bucket, prefix, "change the acl", pool, ch, actionFunc, nil, nil, false)
	}

	actionFunc := func(bucket, key string) bool {
		return c.setObjectAcl(bucket, key, "", aclType, acp, 2)
	}
	return c.submitListObjectsTask(bucket, prefix, "change the acl", pool, ch, actionFunc, nil)
}

func (c *chattriCommand) setObjectsAcl(bucket, prefix string, aclType obs.AclType, acp *accessControlPolicy) error {
	submitFunc := func(pool concurrent.Pool, ch progress.SingleBarChan) (int64, error) {
		return c.submitSetAclTask(bucket, prefix, aclType, acp, pool, ch)
	}

	errorHandleFunc := func(hasListError error) {
		logError(hasListError, LEVEL_ERROR, fmt.Sprintf("\nList objects in the bucket [%s] to change the acl failed", bucket))
	}

	return c.recursiveAction(bucket, prefix, submitFunc, errorHandleFunc, c.recordStartFunc, true)
}

func (c *chattriCommand) isNewVersionBucket(bucket string) error {
	obsVersion := c.checkBucketVersion(bucket)

	if obsVersion < "3.0" {
		return fmt.Errorf("Error: Bucket [%s] cannot support setObjectMetadata interface, skip the storage class action", bucket)
	}

	if obsVersion == OBS_VERSION_UNKNOWN {
		return fmt.Errorf("Error: Bucket [%s] cannot support setObjectMetadata interface, due to check obs version failed, skip the storage class action", bucket)
	}
	return nil
}

func initChattri() command {
	c := &chattriCommand{}
	c.key = "chattri"
	c.usage = c_cloud_url_usage
	c.description = "set bucket or object properties"
	c.define = func() {
		c.init()
		c.flagSet.BoolVar(&c.recursive, "r", false, "")
		c.flagSet.BoolVar(&c.force, "f", false, "")
		c.flagSet.BoolVar(&c.version, "v", false, "")
		c.flagSet.IntVar(&c.jobs, "j", 0, "")
		c.flagSet.StringVar(&c.sc, "sc", "", "")
		c.flagSet.StringVar(&c.acl, "acl", "", "")
		c.flagSet.StringVar(&c.aclXml, "aclXml", "", "")
		c.flagSet.BoolVar(&c.forceRecord, "fr", false, "")
		c.flagSet.StringVar(&c.versionId, "versionId", "", "")
		c.flagSet.StringVar(&c.outDir, "o", "", "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		var storageClassType obs.StorageClassType
		var aclType obs.AclType
		var acp *accessControlPolicy
		checkParamFunc := func(prefix string) bool {
			values := []string{c.sc, c.acl, c.aclXml}
			valuesJoin := strings.Join(values, "")

			invalidInput := false
			if valuesJoin != "" {
				for _, value := range values {
					if value == valuesJoin {
						invalidInput = true
						break
					}
				}
			}

			if !invalidInput {
				printf("Error: Invalid options, must provide only one of [sc|acl|aclXml]")
				return false
			}

			_, ok := getRequestPayerType(c.payer)
			if !ok {
				return false
			}

			if c.sc != "" {
				if _storageClassType, succeed := getStorageClassType(c.sc); succeed {
					storageClassType = _storageClassType
				} else {
					return false
				}
			}

			if c.acl != "" {
				if prefix == "" && !c.recursive {
					if _aclType, succeed := getBucketAclType(c.acl); succeed {
						aclType = _aclType
					} else {
						return false
					}
				} else {
					if _aclType, succeed := getObjectAclType(c.acl); succeed {
						aclType = _aclType
					} else {
						return false
					}
				}
			}

			if c.aclXml != "" {
				acp = &accessControlPolicy{}
				if err := xml.Unmarshal(assist.StringToBytes(c.aclXml), acp); err != nil {
					printf("Error: aclXml is not in well-format, %s", err.Error())
					return false
				}
			}

			return true
		}

		emptyPrefixFunc := func(bucket string) error {
			ret := true
			if storageClassType != "" {
				ret = c.setBucketStorageClass(bucket, storageClassType)
			} else if aclType != "" || acp != nil {
				ret = c.setBucketAcl(bucket, aclType, acp)
			}

			if !ret {
				return assist.ErrExecuting
			}
			return nil
		}

		prefixFunc := func(bucket, prefix string, batchFlag int) error {
			var ret bool
			if storageClassType != "" {
				if err := c.isNewVersionBucket(bucket); err != nil {
					printError(err)
					return assist.ErrUnsupported
				}
				ret = c.setObjectStorageClass(bucket, prefix, c.versionId, storageClassType, batchFlag)
			} else if aclType != "" || acp != nil {
				ret = c.setObjectAcl(bucket, prefix, c.versionId, aclType, acp, batchFlag)
			}

			if ret {
				return nil
			}

			return assist.ErrExecuting
		}
		recursivePrefixFun := func(bucket, prefix string) error {
			if storageClassType != "" {
				if err := c.isNewVersionBucket(bucket); err != nil {
					printError(err)
					return assist.ErrUnsupported
				}
				return c.setObjectsStorageClass(bucket, prefix, storageClassType)
			} else if aclType != "" || acp != nil {
				return c.setObjectsAcl(bucket, prefix, aclType, acp)
			}

			return nil
		}

		return c.chooseAction(checkParamFunc, emptyPrefixFunc, nil, prefixFunc, recursivePrefixFun, c.recordStartFunc)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s\n", "", p.Sprintf("set bucket or object properties"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil chattri obs://bucket [-sc=xxx] [-acl=xxx] [-aclJson=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil chattri obs://bucket/key [-sc=xxx] [-acl=xxx] [-aclJson=xxx] [-versionId=xxx] [-fr] [-o=xxx] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 3:")
		printf("%2s%s", "", "obsutil chattri obs://bucket/[prefix] -r [-f] [-v] [-sc=xxx] [-acl=xxx] [-aclJson=xxx] [-o=xxx] [-j=1] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-r")
		printf("%4s%s", "", p.Sprintf("batch set the properties of objects by prefix"))
		printf("")
		printf("%2s%s", "", "-f")
		printf("%4s%s", "", p.Sprintf("force mode, the progress will not be suspended while an object is to be changed"))
		printf("")
		printf("%2s%s", "", "-v")
		printf("%4s%s", "", p.Sprintf("batch set properties of versions of objects by prefix"))
		printf("")
		printf("%2s%s", "", "-j=1")
		printf("%4s%s", "", p.Sprintf("the maximum number of concurrent jobs for setting object properties, the default value can be set in the config file"))
		printf("")
		printf("%2s%s", "", "-sc=xxx")
		printf("%4s%s", "", p.Sprintf("the storage class that can be specified for a bucket or objects, possible values are [standard|warm|cold]"))
		printf("")
		printf("%2s%s", "", "-acl=xxx")
		printf("%4s%s", "", p.Sprintf("the ACL that can be specified for a bucket or objects, possible values are [private|public-read|public-read-write] for a bucket or [private|public-read|public-read-write|bucket-owner-full-control] for objects"))
		printf("")
		printf("%2s%s", "", "-aclXml=xxx")
		printf("%4s%s", "", p.Sprintf("Bucket or object ACL, in XML format"))
		printf("")
		printf("%2s%s", "", "-fr")
		printf("%4s%s", "", p.Sprintf("force to generate the record files when setting the properties of a single object"))
		printf("")
		printf("%2s%s", "", "-versionId=xxx")
		printf("%4s%s", "", p.Sprintf("the version ID of the object to be made change upon"))
		printf("")
		printf("%2s%s", "", "-o=xxx")
		printf("%4s%s", "", p.Sprintf("the output dir, used to record the results for setting object properties"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
