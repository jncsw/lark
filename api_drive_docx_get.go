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

// GetDocxDocument 获取文档最新版本号、标题等
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnby5Y0yoACL3PdfZqrJEm6f#doxcnWKAE4aSaIU4GcdLInSaVde), 了解相关规则及约束。
// 频率限制: 单个应用调用频率上限为每秒 5 次。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/get
func (r *DriveService) GetDocxDocument(ctx context.Context, request *GetDocxDocumentReq, options ...MethodOptionFunc) (*GetDocxDocumentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDocxDocument != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDocxDocument mock enable")
		return r.cli.mock.mockDriveGetDocxDocument(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDocxDocument",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDocxDocumentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDocxDocument mock DriveGetDocxDocument method
func (r *Mock) MockDriveGetDocxDocument(f func(ctx context.Context, request *GetDocxDocumentReq, options ...MethodOptionFunc) (*GetDocxDocumentResp, *Response, error)) {
	r.mockDriveGetDocxDocument = f
}

// UnMockDriveGetDocxDocument un-mock DriveGetDocxDocument method
func (r *Mock) UnMockDriveGetDocxDocument() {
	r.mockDriveGetDocxDocument = nil
}

// GetDocxDocumentReq ...
type GetDocxDocumentReq struct {
	DocumentID string `path:"document_id" json:"-"` // 文档的唯一标识, 示例值: "doxcnePuYufKa49ISjhD8Ih0ikh", 长度范围: `27` ～ `27` 字符
}

// GetDocxDocumentResp ...
type GetDocxDocumentResp struct {
	Document *GetDocxDocumentRespDocument `json:"document,omitempty"` // 文档信息
}

// GetDocxDocumentRespDocument ...
type GetDocxDocumentRespDocument struct {
	DocumentID string `json:"document_id,omitempty"` // 文档唯一标识
	RevisionID int64  `json:"revision_id,omitempty"` // 文档版本 ID
	Title      string `json:"title,omitempty"`       // 文档标题
}

// getDocxDocumentResp ...
type getDocxDocumentResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetDocxDocumentResp `json:"data,omitempty"`
}
