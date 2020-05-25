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
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type messageBuilderCn struct {
}

func (m messageBuilderCn) buildMessage(printers map[language.Tag]*PrinterWrapper) {
	printers[language.Chinese] = newPrinterWrapper(message.NewPrinter(language.Chinese))
	m.setCommonStrings()
	m.setAbortStrings()
	m.setArchiveStrings()
	m.setChattriStrings()
	m.setClearStrings()
	m.setConfigStrings()
	m.setHelpStrings()
	m.setLsStrings()
	m.setMbStrings()
	m.setRestoreStrings()
	m.setRmStrings()
	m.setStatStrings()
	m.setCpStrings()
	m.setMvStrings()
	m.setSyncStrings()
	m.setVersionStrings()
	m.setMkdirStrings()
	m.setSignStrings()
	m.setShareLsStrings()
	m.setShareCpStrings()
	m.setShareCrtStrings()
	m.setCatStrings()
	m.setHashStrings()
	m.setDirectDownloadStrings()
	m.setBucketPolicyStrings()
}

func (messageBuilderCn) setCommonStrings() {
	message.SetString(language.Chinese, "Usage:", "用法：")
	message.SetString(language.Chinese, "You can use \"obsutil help command\" to view the specific help of each command",
		"您可以使用 \"obsutil help command\" 查看每个命令的详细帮助信息")

	message.SetString(language.Chinese, "Basic commands:", "基本命令：")
	message.SetString(language.Chinese, "Other commands:", "辅助命令：")
	message.SetString(language.Chinese, "Syntax:", "语法：")
	message.SetString(language.Chinese, "Summary:", "摘要：")
	message.SetString(language.Chinese, "Syntax:", "语法：")
	message.SetString(language.Chinese, "Syntax 1:", "语法一：")
	message.SetString(language.Chinese, "Syntax 2:", "语法二：")
	message.SetString(language.Chinese, "Syntax 3:", "语法三：")
	message.SetString(language.Chinese, "Syntax 4:", "语法四：")
	message.SetString(language.Chinese, "Syntax 5:", "语法五：")
	message.SetString(language.Chinese, "Syntax 6:", "语法六：")
	message.SetString(language.Chinese, "Syntax 7:", "语法七：")
	message.SetString(language.Chinese, "Syntax 8:", "语法八：")
	message.SetString(language.Chinese, "Syntax 9:", "语法九：")
	message.SetString(language.Chinese, "Syntax 10:", "语法十：")
	message.SetString(language.Chinese, "Options:", "选项：")
	message.SetString(language.Chinese, "force to generate the record files", "操作单个对象时生成结果清单文件")
	message.SetString(language.Chinese, "the path to the custom config file when running this command", "运行当前命令时的自定义配置文件")
}

func (messageBuilderCn) setAbortStrings() {
	message.SetString(language.Chinese, "abort multipart uploads", "删除分段上传任务")
	message.SetString(language.Chinese, "the ID of the multipart upload to be aborted", "待删除的单个分段上传任务的ID")
	message.SetString(language.Chinese, "batch abort multipart uploads by prefix", "按指定的对象名前缀批量删除分段上传任务")
	message.SetString(language.Chinese, "force mode, the progress will not be suspended while a multipart upload is to be aborted", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "force to generate the record files when aborting one multipart upload", "删除单个分段上传任务时生成结果清单文件")
	message.SetString(language.Chinese, "the maximum number of concurrent abort jobs, the default value can be set in the config file", "批量删除分段上传任务的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the output dir, used to record the abort results", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setArchiveStrings() {
	message.SetString(language.Chinese, "archive log files to local file system or OBS", "将日志文件归档到本地，或归档到指定的桶")
}

func (messageBuilderCn) setChattriStrings() {
	message.SetString(language.Chinese, "set bucket or object properties", "设置桶或对象的属性")
	message.SetString(language.Chinese, "batch set the properties of objects by prefix", "按指定的对象名前缀批量设置对象属性")
	message.SetString(language.Chinese, "force mode, the progress will not be suspended while an object is to be changed", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "batch set properties of versions of objects by prefix", "按指定的对象名前缀批量设置多版本对象属性")
	message.SetString(language.Chinese, "the maximum number of concurrent jobs for setting object properties, the default value can be set in the config file", "批量设置对象属性的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the storage class that can be specified for a bucket or objects, possible values are [standard|warm|cold]", "桶或对象的存储类型。支持的值：[standard|warm|cold]")
	message.SetString(language.Chinese, "the ACL that can be specified for a bucket or objects, possible values are [private|public-read|public-read-write] for a bucket or [private|public-read|public-read-write|bucket-owner-full-control] for objects",
		"桶或对象的预定义访问策略。作用于桶时支持的值：[private|public-read|public-read-write]，作用于对象时支持的值：[private|public-read|public-read-write|bucket-owner-full-control]")
	message.SetString(language.Chinese, "Bucket or object ACL, in XML format", "桶或对象的访问策略（xml格式配置）")
	message.SetString(language.Chinese, "force to generate the record files when setting the properties of a single object", "设置单个对象属性时生成结果清单文件")
	message.SetString(language.Chinese, "the version ID of the object to be made change upon", "待设置属性的单个对象的版本号")
	message.SetString(language.Chinese, "the output dir, used to record the results for setting object properties", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setClearStrings() {
	message.SetString(language.Chinese, "delete part records", "删除指定文件夹下的断点记录文件")
	message.SetString(language.Chinese, "delete the part records of all multipart upload jobs", "删除所有分段上传任务的断点记录文件")
	message.SetString(language.Chinese, "delete the part records of all multipart download jobs", "删除所有分段下载任务的断点记录文件")
	message.SetString(language.Chinese, "delete the part records of all multipart copy jobs", "删除所有分段复制任务的断点记录文件")
}

func (messageBuilderCn) setConfigStrings() {
	message.SetString(language.Chinese, "update the configuration file", "更新配置文件中的部分配置信息")
	message.SetString(language.Chinese, "update the configuration file through interactive mode", "使用交互式模式更新配置")
	message.SetString(language.Chinese, "endpoint", "配置文件中的endpoint/endpointCrr")
	message.SetString(language.Chinese, "access key ID", "配置文件中的ak/akCrr")
	message.SetString(language.Chinese, "security key ID", "配置文件中的sk/skCrr")
	message.SetString(language.Chinese, "security token", "配置文件中的token/tokenCrr")
	message.SetString(language.Chinese, "update the configuration file for crr", "更新配置文件中的akCrr/skCrr/endpointCrr/tokenCrr")
}

func (messageBuilderCn) setHelpStrings() {
	message.SetString(language.Chinese, "view command help information", "查看命令帮助")
	message.SetString(language.Chinese, "view the commands supported by this tool or the help information of a specific command", "查看工具支持的命令，或查看某个具体命令的帮助文档")
}

func (messageBuilderCn) setLsStrings() {
	message.SetString(language.Chinese, "list buckets or objects/multipart uploads in a bucket", "列举桶、列举桶内对象或列举桶内分段上传任务")
	message.SetString(language.Chinese, "show results in brief mode", "以精简格式显示查询结果")
	message.SetString(language.Chinese, "show storage class of each bucket", "查询桶列表时同时查询桶的存储类型")
	message.SetString(language.Chinese, "the maximum number of concurrent jobs for querying the storage classes of buckets, the default value can be set in the config file", "查询桶存储类型时的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "list objects and sub-folders in the current folder", "仅列举当前目录下的对象和子目录，而非递归列举所有对象和子目录")
	message.SetString(language.Chinese, "list versions of objects in a bucket", "列举桶内多版本对象，列举结果包含最新版本的对象和历史版本的对象（如果有）")
	message.SetString(language.Chinese, "list multipart uploads", "列举桶内分段上传任务")
	message.SetString(language.Chinese, "list both objects and multipart uploads", "同时列举桶内对象和桶内分段上传任务")
	message.SetString(language.Chinese, "the marker to list objects or multipart uploads", "列举桶内对象的起始位置，或列举桶内分段上传任务的起始位置")
	message.SetString(language.Chinese, "the version ID marker to list versions of objects", "列举桶内多版本对象的起始位置，必须与marker配合使用")
	message.SetString(language.Chinese, "the upload ID marker to list multipart uploads", "列举桶内分段上传任务的起始位置，必须与marker配合使用")
	message.SetString(language.Chinese, "the bytes format in results when listing objects or listing versions of objects, possible values are [human-readable|raw]", "列举桶内对象或列举桶内多版本对象的返回结果中字节数的显示格式。支持的值：[human-readable|raw]")
	message.SetString(language.Chinese, "show results by limited number", "列举结果的最大个数")
}

func (messageBuilderCn) setMbStrings() {
	message.SetString(language.Chinese, "create a bucket with the specified parameters", "按照用户指定的桶名和参数创建一个新桶")
	message.SetString(language.Chinese, "create a bucket that supports POSIX", "创建支持文件接口（POSIX）的桶")
	message.SetString(language.Chinese, "the AZ of the bucket, possible values are [multi-az]", "创桶时可指定的可用区。支持的值：[multi-az]")
	message.SetString(language.Chinese, "the ACL of the bucket, possible values are [private|public-read|public-read-write]", "创桶时可指定的预定义访问策略。支持的值：[private|public-read|public-read-write]")
	message.SetString(language.Chinese, "the default storage class of the bucket, possible values are [standard|warm|cold]", "创桶时可指定的桶的默认存储类型。支持的值：[standard|warm|cold]")
	message.SetString(language.Chinese, "the region where the bucket is located", "桶所在的区域")
}

func (messageBuilderCn) setRestoreStrings() {
	message.SetString(language.Chinese, "restore objects in a bucket to be readable", "恢复指定的归档存储对象或按指定的对象名前缀批量恢复归档存储对象")
	message.SetString(language.Chinese, "batch restore objects by prefix", "按指定的对象名前缀批量恢复归档存储对象")

	message.SetString(language.Chinese, "force mode, the progress will not be suspended while an object is to be restored", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "force to generate the record files when restoring one object", "恢复单个归档存储对象时生成结果清单文件")
	message.SetString(language.Chinese, "restore versions of objects by prefix", "按指定的对象名前缀批量恢复多版本归档存储对象")
	message.SetString(language.Chinese, "retention period of each restored object, in days. the range is [1, 30] and the default value is 1", "归档存储对象恢复后的保存时间，单位：天，取值范围是1~30。默认为1")
	message.SetString(language.Chinese, "the maximum number of concurrent restore jobs, the default value can be set in the config file", "批量恢复归档存储对象的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the version ID of the object to be restored", "待恢复的单个归档存储对象的版本号")
	message.SetString(language.Chinese, "option for restoring objects, possible values are [standard|expedited]", "恢复选项，支持的值：[standard|expedited]")
	message.SetString(language.Chinese, "the output dir, used to record the restore results", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setRmStrings() {
	message.SetString(language.Chinese, "delete a bucket or objects in a bucket", "删除桶，或删除桶内对象")
	message.SetString(language.Chinese, "batch delete objects by prefix", "按指定的对象名前缀批量删除对象")

	message.SetString(language.Chinese, "force mode, the progress will not be suspended while a bucket or an object is to be deleted", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "force to generate the record files when deleting one object", "删除单个对象时生成结果清单文件")
	message.SetString(language.Chinese, "batch delete versions of objects and the delete markers by prefix", "按指定的对象名前缀批量删除多版本对象和多版本删除标记")
	message.SetString(language.Chinese, "the version ID of the object to be deleted", "待删除的单个对象的版本号")
	message.SetString(language.Chinese, "the maximum number of concurrent delete jobs, the default value can be set in the config file", "批量删除对象时的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the output dir, used to record the deleted results", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setStatStrings() {
	message.SetString(language.Chinese, "show the properties of a bucket or an object", "查询桶属性，或查询对象属性")
	message.SetString(language.Chinese, "show the ACL of the bucket or the object", "查询桶属性或对象属性时，同时查询访问策略")
}

func (messageBuilderCn) setCpStrings() {
	message.SetString(language.Chinese, "upload, download or copy objects", "上传/下载/复制对象")
	message.SetString(language.Chinese, "conduct a dry run", "测试模式运行，不执行实际的操作")
	message.SetString(language.Chinese, "batch upload, download or copy objects by prefix (folder)", "递归上传文件夹，或按对象名前缀批量下载/复制对象")
	message.SetString(language.Chinese, "recover the failed upload, download or copy tasks by task ID", "待恢复上传/下载/复制对象任务的结果清单文件任务号")
	message.SetString(language.Chinese, "force mode, the progress will not be suspended while objects are to be uploaded, downloaded or copied", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "upload, download or copy the changed sources only", "增量上传/下载/复制对象")
	message.SetString(language.Chinese, "verify the size after the objects are uploaded or downloaded or copied through crr", "上传/下载/复制对象完成后，验证大小是否一致")
	message.SetString(language.Chinese, "verify the MD5 value after the objects are uploaded or downloaded or copied through crr", "上传/下载/复制对象时验证数据一致性")
	message.SetString(language.Chinese, "upload, download or copy the sources without the relative parent folder/prefix", "上传/下载/复制对象时不包含相对父目录")
	message.SetString(language.Chinese, "the directory where the part records reside, used to record the progress of upload, download or copy jobs", "生成断点记录文件的文件夹")
	message.SetString(language.Chinese, "the maximum number of concurrent upload, download or copy jobs, the default value can be set in the config file", "批量任务的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the maximum number of concurrent upload, download or copy tasks (a task is a sub-job), the default value can be set in the config file", "每个分段任务的最大并发数，默认为配置文件中的defaultParallels")
	message.SetString(language.Chinese, "the version ID of the object to be copied or downloaded", "下载/复制单个对象的版本号")
	message.SetString(language.Chinese, "the part size of each upload, download or copy task, the default value can be set in the config file", "每个分段任务的段大小，单位：字节，默认为配置文件中的defaultPartSize")
	message.SetString(language.Chinese, "the threshold, if it is exceeded, the upload, download or copy job will be divided into multiple tasks by the part size, the default value can be set in the config file", "开启分段任务的阈值，单位：字节，默认为配置文件中的defaultBigfileThreshold")
	message.SetString(language.Chinese, "the customized metadata of each object to be uploaded or copied", "上传/复制对象时可指定的自定义元数据")

	message.SetString(language.Chinese, "the ACL of each object to be uploaded or copied, possible values are [private|public-read|public-read-write|bucket-owner-full-control]", "上传/复制对象时可指定的预定义访问策略。支持的值：[private|public-read|public-read-write|bucket-owner-full-control]")
	message.SetString(language.Chinese, "the storage class of each object to be uploaded or copied, possible values are [standard|warm|cold]", "上传/复制对象时可指定的对象的存储类型。支持的值：[standard|warm|cold]")
	message.SetString(language.Chinese, "the objects whose names match this pattern will be included", "上传/下载/复制时对包含文件/对象的匹配模式")
	message.SetString(language.Chinese, "the objects whose names match this pattern will be excluded", "上传/下载/复制时对不包含文件/对象的匹配模式")
	message.SetString(language.Chinese, "the including pattern, the excluding pattern and the time range pattern will task effect on folders", "设置名称匹配模式和时间段匹配模式对文件夹也生效")
	message.SetString(language.Chinese, "the output dir, used to record the upload, download or copy results", "生成结果清单文件的文件夹")

	message.SetString(language.Chinese, "upload the actual path of the symbolic-link file/folder", "上传软链接文件/文件夹指向的真实路径")
	message.SetString(language.Chinese, "the archive dir, used to archive the successful uploaded file(s)", "上传文件成功后的归档路径")
	message.SetString(language.Chinese, "the files whose latest access time falls into the time range (-timeRange option) will be uploaded", "上传文件夹中文件的最后访问时间满足timeRange选项的文件列表")
	message.SetString(language.Chinese, "the folder will not be uploaded as a object", "上传时候文件夹本身不会作为单独一个对象上传")
	message.SetString(language.Chinese, "the time range for last modified time, between which the objects will be uploaded, downloaded or copied", "上传/下载/复制时文件最后修改时间的时间段匹配模式")
	message.SetString(language.Chinese, "the temp file dir, used to save temporary files during the objects are downloading", "下载时保存临时文件的文件夹")
	message.SetString(language.Chinese, "multi-source upload mode. Possible values are [1|2]. 1 indicates the source URL is a list of file names separated by commas.2 indicates the source URL is a list file", "开启多文件/文件夹上传模式，如果该值为1则代表上传的URL是一组文件列表（以逗号分隔）；如果该值为2则代表上传的URL是一个包含文件列表的文件。支持的值：[1|2]")
	message.SetString(language.Chinese, "copy the source object(s) through the client-side cross region replication mode", "复制时使用客户端跨区域复制模式")
}

func (messageBuilderCn) setMvStrings() {
	message.SetString(language.Chinese, "move objects", "移动对象")

	message.SetString(language.Chinese, "batch move objects by prefix", "按对象名前缀批量移动对象")
	message.SetString(language.Chinese, "force mode, the progress will not be suspended while objects are to be moved", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "move the changed sources only", "增量移动对象")
	message.SetString(language.Chinese, "move the sources without the relative parent prefix", "移动对象时不包含相对父目录")
	message.SetString(language.Chinese, "the directory where the part records reside, used to record the progress of movement jobs", "生成断点记录文件的文件夹")
	message.SetString(language.Chinese, "the maximum number of movement jobs, the default value can be set in the config file", "批量任务的最大并发数，默认为配置文件中的defaultJobs")

	message.SetString(language.Chinese, "the maximum number of concurrent movement tasks (a task is a sub-job), the default value can be set in the config file", "每个分段任务的最大并发数，默认为配置文件中的defaultParallels")
	message.SetString(language.Chinese, "the version ID of the object to be moved", "移动单个对象的版本号")
	message.SetString(language.Chinese, "the part size of each movement task, the default value can be set in the config file", "每个分段任务的段大小，单位：字节，默认为配置文件中的defaultPartSize")
	message.SetString(language.Chinese, "the threshold, if it is exceeded, the movement job will be divided into multiple tasks by the part size, the default value can be set in the config file", "开启分段任务的阈值，单位：字节，默认为配置文件中的defaultBigfileThreshold")
	message.SetString(language.Chinese, "the customized metadata of each object to be moved", "移动对象时可指定的自定义元数据")
	message.SetString(language.Chinese, "the ACL of each object to be moved, possible values are [private|public-read|public-read-write|bucket-owner-full-control]", "移动对象时可指定的预定义访问策略。支持的值：[private|public-read|public-read-write|bucket-owner-full-control]")
	message.SetString(language.Chinese, "the storage class of each object to be moved, possible values are [standard|warm|cold]", "移动对象时可指定的对象的存储类型。支持的值：[standard|warm|cold]")
	message.SetString(language.Chinese, "the output dir, used to record the movement results", "生成结果清单文件的文件夹")
	message.SetString(language.Chinese, "the to be moved objects whose names match this pattern will be included", "移动时对包含对象的匹配模式")
	message.SetString(language.Chinese, "the to be moved objects whose names match this pattern will be excluded", "移动时对不包含对象的匹配模式")
	message.SetString(language.Chinese, "the time range, between which the objects will be moved", "移动时的时间段匹配模式")
}

func (messageBuilderCn) setSyncStrings() {
	message.SetString(language.Chinese, "synchronize objects from the source to the destination", "增量同步上传/下载/复制对象")
	message.SetString(language.Chinese, "the directory where the part records reside, used to record the synchronization progress", "生成断点记录文件的文件夹")
	message.SetString(language.Chinese, "the maximum number of concurrent synchronization jobs, the default value can be set in the config file", "上传/下载/复制时批量任务的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the maximum number of concurrent synchronization tasks (a task is a sub-job), the default value can be set in the config file", "每个分段任务的最大并发数，默认为配置文件中的defaultParallels")
	message.SetString(language.Chinese, "the part size of each synchronization task, the default value can be set in the config file", "每个分段任务的段大小，单位：字节，默认为配置文件中的defaultPartSize")
	message.SetString(language.Chinese, "the threshold, if it is exceeded, the synchronization job will be divided into multiple tasks by the part size, the default value can be set in the config file", "开启分段任务的阈值，单位：字节，默认为配置文件中的defaultBigfileThreshold")
	message.SetString(language.Chinese, "the output dir, used to record the synchronization results", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setVersionStrings() {
	message.SetString(language.Chinese, "show version", "查看版本号")
	message.SetString(language.Chinese, "show the version of this tool", "查看工具当前的版本号")
}

func (messageBuilderCn) setMkdirStrings() {
	message.SetString(language.Chinese, "create folder(s) in a specified bucket or in the local file system", "在指定桶内或本地文件系统中创建文件夹")
}

func (messageBuilderCn) setSignStrings() {
	message.SetString(language.Chinese, "generate the download url(s) for the objects in a specified bucket", "生成指定桶内对象的下载链接")
	message.SetString(language.Chinese, "batch generate the download url(s) of objects by prefix", "按指定的对象名前缀批量生成对象的下载链接")
	message.SetString(language.Chinese, "the expiration time of the generated download url(s), in seconds, the default value is 300", "生成的对象下载链接的过期时间（单位：秒），默认为300")
	message.SetString(language.Chinese, "the output dir, used to record the generated download urls", "生成结果清单文件的文件夹")
	message.SetString(language.Chinese, "the objects whose names match this pattern will be included when generating the download urls", "生成下载链接时，包含对象的匹配模式")
	message.SetString(language.Chinese, "the objects whose names match this pattern will be excluded when generating the download urls", "生成下载链接时，不包含对象的匹配模式")
	message.SetString(language.Chinese, "the time range, between which the download url(s) of objects will be generated", "生成下载链接时的时间段匹配模式")
}

func (messageBuilderCn) setShareLsStrings() {
	message.SetString(language.Chinese, "list objects using authorization code and access code", "使用授权码和提取码列举对象")
	message.SetString(language.Chinese, "the access code", "提取码")
	message.SetString(language.Chinese, "the prefx to list objects", "列举桶内对象时的对象名前缀")
	message.SetString(language.Chinese, "the bytes format in results when listing objects, possible values are [human-readable|raw]", "列举桶内对象返回结果中字节数的显示格式。支持的值：[human-readable|raw]")
	message.SetString(language.Chinese, "the marker to list objects", "列举桶内对象的起始位置")
}

func (messageBuilderCn) setShareCpStrings() {
	message.SetString(language.Chinese, "download objects using authorization code and access code", "使用授权码和提取码下载对象")
	message.SetString(language.Chinese, "the key to download, or the prefix to batch download", "下载单个对象的对象名，或者批量下载的对象名前缀")
	message.SetString(language.Chinese, "batch downloads objects by prefix", "按对象名前缀批量下载对象")
	message.SetString(language.Chinese, "force mode, the progress will not be suspended while objects are to be downloaded", "强制操作，不进行询问提示")
	message.SetString(language.Chinese, "download the changed sources only", "增量下载对象")
	message.SetString(language.Chinese, "verify the size after the objects are downloaded", "下载对象完成后，验证大小是否一致")
	message.SetString(language.Chinese, "verify the MD5 value after the objects are downloaded", "下载对象时验证数据一致性")
	message.SetString(language.Chinese, "download the sources without the relative parent prefix", "下载对象时不包含相对父目录")
	message.SetString(language.Chinese, "the directory where the part records reside, used to record the progress of download jobs", "生成断点记录文件的文件夹")
	message.SetString(language.Chinese, "the maximum number of concurrent download jobs, the default value can be set in the config file", "批量任务的最大并发数，默认为配置文件中的defaultJobs")
	message.SetString(language.Chinese, "the maximum number of concurrent download tasks (a task is a sub-job), the default value can be set in the config file", "每个分段任务的最大并发数，默认为配置文件中的defaultParallels")
	message.SetString(language.Chinese, "the part size of each download task, the default value can be set in the config file", "每个分段任务的段大小，单位：字节，默认为配置文件中的defaultPartSize")
	message.SetString(language.Chinese, "the threshold, if it is exceeded, the download job will be divided into multiple tasks by the part size, the default value can be set in the config file", "开启分段任务的阈值，单位：字节，默认为配置文件中的defaultBigfileThreshold")
	message.SetString(language.Chinese, "the output dir, used to record the download results", "生成结果清单文件的文件夹")
}

func (messageBuilderCn) setShareCrtStrings() {
	message.SetString(language.Chinese, "create authorization code for sharing", "创建目录分享的授权码")
	message.SetString(language.Chinese, "the validity period of authorization code, the default value is 1 day", "授权码的有效期，默认为1天")
	message.SetString(language.Chinese, "the download url to which the result is generated", "生成的授权码的保存路径")
}

func (messageBuilderCn) setCatStrings() {
	message.SetString(language.Chinese, "view the content of a text object in a bucket", "查看桶内文本类型对象的内容")
}

func (messageBuilderCn) setHashStrings() {
	message.SetString(language.Chinese, "caculate the md5 or crc64 hash code of a local file", "使用crc64或md5加密算法计算本地文件的校验值")
	message.SetString(language.Chinese, "the encryption algorithm type, possible values are [md5|crc64], the default value is md5", "加密算法类型，支持的值：[md5|crc64]，默认为md5")
}

func (messageBuilderCn) setDirectDownloadStrings() {
	message.SetString(language.Chinese, "download an object directly using the specified resource url", "使用指定的资源链接支持下载对象")
}

func (messageBuilderCn) setBucketPolicyStrings() {
	message.SetString(language.Chinese, "get, put or delete bucket policy", "获取，设置或者删除桶策略")
	message.SetString(language.Chinese, "the operation you want to do,possible values are [get, put, delete]", "操作类型，支持的值：[get|put|delete]")
	message.SetString(language.Chinese, "the policy json file which you want to get or put, only support when method is get or put", "策略文件的路径，只有当操作类型为put或者get才支持")
}
