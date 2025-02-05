// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// CreateDocx 创建新版文档, 文档标题和目录可选。
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnby5Y0yoACL3PdfZqrJEm6f#doxcnyoyCgwS8ywWwMtQr9yjZ2f), 了解相关规则及约束。
// 频率限制: 单个应用调用频率上限为每秒 3 次。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/create
func (r *DriveService) CreateDocx(ctx context.Context, request *CreateDocxReq, options ...MethodOptionFunc) (*CreateDocxResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDocx != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDocx mock enable")
		return r.cli.mock.mockDriveCreateDocx(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDocx",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDocxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDocx mock DriveCreateDocx method
func (r *Mock) MockDriveCreateDocx(f func(ctx context.Context, request *CreateDocxReq, options ...MethodOptionFunc) (*CreateDocxResp, *Response, error)) {
	r.mockDriveCreateDocx = f
}

// UnMockDriveCreateDocx un-mock DriveCreateDocx method
func (r *Mock) UnMockDriveCreateDocx() {
	r.mockDriveCreateDocx = nil
}

// CreateDocxReq ...
type CreateDocxReq struct {
	FolderToken *string `json:"folder_token,omitempty"` // 文件夹 token, 获取方式见云文档接口快速入门；空表示根目录, tenant_access_token应用权限仅允许操作应用创建的目录, 示例值: "fldcnqquW1svRIYVT2Np6IuLCKd"
	Title       *string `json:"title,omitempty"`        // 文档标题, 只支持纯文本, 示例值: "undefined", 长度范围: `1` ～ `800` 字符
}

// CreateDocxResp ...
type CreateDocxResp struct {
	Document *CreateDocxRespDocument `json:"document,omitempty"` // 新建文档的文档信息
}

// CreateDocxRespDocument ...
type CreateDocxRespDocument struct {
	DocumentID string `json:"document_id,omitempty"` // 文档唯一标识
	RevisionID int64  `json:"revision_id,omitempty"` // 文档版本 ID
	Title      string `json:"title,omitempty"`       // 文档标题
}

// createDocxResp ...
type createDocxResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateDocxResp `json:"data,omitempty"`
}
