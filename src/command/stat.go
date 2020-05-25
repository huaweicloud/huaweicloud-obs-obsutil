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
	"fmt"
	"obs"
	"strings"
)

type statCommand struct {
	cloudUrlCommand
	acl   bool
	payer string
}

func (c *statCommand) getBucketStorageInfo(bucket string) (int, int64, error) {
	output, err := obsClient.GetBucketStorageInfo(bucket, obs.WithReqPaymentHeader(c.payer))
	if err == nil {
		return output.ObjectNumber, output.Size, nil
	}

	return -1, -1, err

}

func (c *statCommand) getBucketQuota(bucket string) (int64, error) {
	output, err := obsClient.GetBucketQuota(bucket, obs.WithReqPaymentHeader(c.payer))
	if err == nil {
		return output.Quota, nil
	}
	return -1, err
}

func (c *statCommand) getBucketAcl(bucket string) string {
	output, err := obsClient.GetBucketAcl(bucket, obs.WithReqPaymentHeader(c.payer))
	if err == nil {
		return c.parseAccessControlPolicy(output.AccessControlPolicy)
	}
	doLogError(err, LEVEL_ERROR, fmt.Sprintf("Get the acl of bucket [%s] error", bucket))
	return ""
}

func (c *statCommand) getObjectAcl(bucket, key string) string {
	input := &obs.GetObjectAclInput{}
	input.Bucket = bucket
	input.Key = key
	input.RequestPayer = c.payer
	output, err := obsClient.GetObjectAcl(input)
	if err == nil {
		return c.parseAccessControlPolicy(output.AccessControlPolicy)
	}
	doLogError(err, LEVEL_ERROR, fmt.Sprintf("Get the acl of object [%s] in the bucket [%s] error", key, bucket))
	return ""
}

func (c *statCommand) parseAccessControlPolicy(acp obs.AccessControlPolicy) string {
	aclXmlSlice := make([]string, 0, 19)
	aclXmlSlice = append(aclXmlSlice, "<AccessControlPolicy>")
	aclXmlSlice = append(aclXmlSlice, "<Owner>")
	aclXmlSlice = append(aclXmlSlice, "<ID>")
	aclXmlSlice = append(aclXmlSlice, acp.Owner.ID)
	aclXmlSlice = append(aclXmlSlice, "</ID>")
	aclXmlSlice = append(aclXmlSlice, "</Owner>")
	aclXmlSlice = append(aclXmlSlice, "<AccessControlList>")

	isOwnerFullContrl := false
	allUsersRead := false
	allUsersWrite := false
	hasOtherUsers := false
	allUserFullControl := false

	for _, grant := range acp.Grants {
		aclXmlSlice = append(aclXmlSlice, "<Grant>")
		aclXmlSlice = append(aclXmlSlice, "<Grantee>")
		if grant.Grantee.ID != "" {
			aclXmlSlice = append(aclXmlSlice, "<ID>")
			aclXmlSlice = append(aclXmlSlice, grant.Grantee.ID)
			aclXmlSlice = append(aclXmlSlice, "</ID>")
		} else {
			aclXmlSlice = append(aclXmlSlice, "<Canned>")
			if grant.Grantee.URI == obs.GroupAllUsers {
				aclXmlSlice = append(aclXmlSlice, aclEveryOne)
			} else {
				aclXmlSlice = append(aclXmlSlice, string(grant.Grantee.URI))
			}
			aclXmlSlice = append(aclXmlSlice, "</Canned>")
		}
		aclXmlSlice = append(aclXmlSlice, "</Grantee>")
		aclXmlSlice = append(aclXmlSlice, "<Permission>")
		aclXmlSlice = append(aclXmlSlice, string(grant.Permission))
		aclXmlSlice = append(aclXmlSlice, "</Permission>")
		aclXmlSlice = append(aclXmlSlice, "</Grant>")

		if grant.Grantee.ID == acp.Owner.ID {
			isOwnerFullContrl = grant.Permission == obs.PermissionFullControl
		} else if grant.Grantee.URI == obs.GroupAllUsers {
			if grant.Permission == obs.PermissionRead {
				allUsersRead = true
			} else if grant.Permission == obs.PermissionWrite {
				allUsersWrite = true
			} else if grant.Permission == obs.PermissionFullControl {
				allUserFullControl = true
			}
		} else {
			hasOtherUsers = true
		}
	}
	aclXmlSlice = append(aclXmlSlice, "</AccessControlList>")
	aclXmlSlice = append(aclXmlSlice, "</AccessControlPolicy>")

	if !hasOtherUsers && isOwnerFullContrl && !allUserFullControl {
		if allUsersRead {
			if allUsersWrite {
				return "public-read-write"
			}
			return "public-read"
		} else if !allUsersWrite {
			return "private"
		}
	}

	return strings.Join(aclXmlSlice, "")
}

func (c *statCommand) getBucketAttribute(bucket string) error {
	input := &obs.GetBucketFSStatusInput{}
	input.Bucket = bucket
	input.RequestPayer = c.payer
	output, err := obsClient.GetBucketFSStatus(input)
	if err == nil {
		objectNumber, size, storageInfoErr := c.getBucketStorageInfo(bucket)
		quota, quotaErr := c.getBucketQuota(bucket)

		aclResult := ""
		if c.acl {
			aclResult = c.getBucketAcl(bucket)
		}

		printf("Bucket:")
		printf("%2s%s", "", "obs://"+bucket)
		printf("StorageClass:")
		printf("%2s%s", "", transStorageClassType(output.StorageClass))
		printf("Location:")
		printf("%2s%s", "", output.Location)
		printf("ObsVersion:")
		printf("%2s%s", "", output.ObsVersion)
		if output.AvailableZone != "" {
			if az := transAvailableZoneType(output.AvailableZone); az != "" {
				printf("AvailableZone:")
				printf("%2s%s", "", az)
			}
		}
		printf("BucketType:")
		bucketType := c_object
		if transFSStatusType(output.FSStatus) == "enabled" {
			bucketType = "POSIX"
		}
		printf("%2s%s", "", bucketType)
		if storageInfoErr == nil {
			printf("ObjectNumber:")
			printf("%2s%d", "", objectNumber)
			printf("Size:")
			printf("%2s%d", "", size)
		} else {
			doLogError(storageInfoErr, LEVEL_INFO, fmt.Sprintf("Get the storage info of bucket [%s] failed", bucket))
		}
		if quotaErr == nil {
			printf("Quota:")
			printf("%2s%d", "", quota)
		} else {
			doLogError(quotaErr, LEVEL_INFO, fmt.Sprintf("Get the quota of bucket [%s] failed", bucket))
		}

		if aclResult != "" {
			printf("Acl:")
			printf("%2s%s", "", aclResult)
		}
		return nil
	}
	printError(err)
	return assist.ErrExecuting
}

func (c *statCommand) getObjectAttribute(bucket, key string) error {
	input := &obs.GetAttributeInput{}
	input.Bucket = bucket
	input.Key = key
	input.RequestPayer = c.payer
	output, err := obsClient.GetAttribute(input)
	if err == nil {
		aclResult := ""
		if c.acl {
			aclResult = c.getObjectAcl(bucket, key)
		}

		printf("Key:")
		printf("%2s%s", "", "obs://"+bucket+"/"+key)
		printf("LastModified:")
		printf("%2s%s", "", output.LastModified.Format(ISO8601_DATE_FORMAT))
		printf("Size:")
		printf("%2s%d", "", output.ContentLength)
		printf("StorageClass:")
		printf("%2s%s", "", transStorageClassType(output.StorageClass))
		if md5, ok := output.Metadata[checkSumKey]; ok {
			printf("MD5:")
			printf("%2s%s", "", md5)
		}
		printf("ETag:")
		etag := output.ETag
		if strings.HasPrefix(etag, "\"") {
			etag = etag[1:]
		}
		if strings.HasSuffix(etag, "\"") {
			etag = etag[:len(etag)-1]
		}
		printf("%2s%s", "", etag)
		printf("ContentType:")
		printf("%2s%s", "", output.ContentType)

		printf("Type:")
		t := "file"
		if output.Mode != -1 {
			if output.Mode&0040000 != 0 {
				t = "folder"
			}
		} else if isObsFolder(key) {
			t = "folder"
		}
		printf("%2s%s", "", t)

		if len(output.Metadata) > 0 {
			meta := make(map[string]string, len(output.Metadata))
			for k, v := range output.Metadata {
				if k != checkSumKey {
					meta[k] = v
				}
			}
			if len(meta) > 0 {
				printf("Metadata:")
				for k, v := range meta {
					printf("%2s%s=%s", "", k, v)
				}
			}
		}

		if aclResult != "" {
			printf("Acl:")
			printf("%2s%s", "", aclResult)
		}
		return nil
	}
	printError(err)
	return assist.ErrExecuting
}

func initStat() command {
	c := &statCommand{}
	c.key = "stat"
	c.usage = "cloud_url"
	c.description = "show the properties of a bucket or an object"

	c.define = func() {
		c.flagSet.BoolVar(&c.acl, "acl", false, "")
		c.flagSet.StringVar(&c.payer, "payer", "", "")
	}

	c.action = func() error {
		cloudUrl, err := c.prepareCloudUrl()
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		bucket, key, err := c.splitCloudUrl(cloudUrl)
		if err != nil {
			printError(err)
			return assist.ErrInvalidArgs
		}

		_, ok := getRequestPayerType(c.payer)
		if !ok {
			printf("Error: Invalid args, please refer to help doc")
			return assist.ErrInvalidArgs
		}

		c.printStart()

		if key == "" {
			return c.getBucketAttribute(bucket)
		}
		return c.getObjectAttribute(bucket, key)
	}

	c.help = func() {
		p := i18n.GetCurrentPrinter()
		p.Printf("Summary:")
		printf("%2s%s", "", p.Sprintf("show the properties of a bucket or an object"))
		printf("")
		p.Printf("Syntax 1:")
		printf("%2s%s", "", "obsutil stat obs://bucket [-acl] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")
		p.Printf("Syntax 2:")
		printf("%2s%s", "", "obsutil stat obs://bucket/key [-acl] [-config=xxx]"+commandCommonSyntax()+commandRequestPayerSyntax())
		printf("")

		p.Printf("Options:")
		printf("%2s%s", "", "-acl")
		printf("%4s%s", "", p.Sprintf("show the ACL of the bucket or the object"))
		printf("")
		printf("%2s%s", "", "-config=xxx")
		printf("%4s%s", "", p.Sprintf("the path to the custom config file when running this command"))
		printf("")
		commandCommonHelp(p)
		commandRequestPayerHelp(p)
	}

	return c
}
