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

package obs

import ()

type IFSClient interface {
	NewBucket(input *NewBucketInput) (output *BaseModel, err error)

	GetBucketFSStatus(input *GetBucketFSStatusInput) (output *GetBucketFSStatusOutput, err error)

	GetAttribute(input *GetAttributeInput) (output *GetAttributeOutput, err error)

	DropFile(input *DropFileInput) (output *DropFileOutput, err error)

	NewFolder(input *NewFolderInput) (output *NewFolderOutput, err error)

	NewFile(input *NewFileInput) (output *NewFileOutput, err error)

	RenameFile(input *RenameFileInput) (output *RenameFileOutput, err error)

	RenameFolder(input *RenameFolderInput) (output *RenameFolderOutput, err error)

	Close()
}
