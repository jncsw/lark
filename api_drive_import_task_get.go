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

// GetDriveImportTask 根据创建导入任务返回的 ticket 查询导入结果。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/get
func (r *DriveService) GetDriveImportTask(ctx context.Context, request *GetDriveImportTaskReq, options ...MethodOptionFunc) (*GetDriveImportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveImportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveImportTask mock enable")
		return r.cli.mock.mockDriveGetDriveImportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveImportTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/import_tasks/:ticket",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveImportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveImportTask mock DriveGetDriveImportTask method
func (r *Mock) MockDriveGetDriveImportTask(f func(ctx context.Context, request *GetDriveImportTaskReq, options ...MethodOptionFunc) (*GetDriveImportTaskResp, *Response, error)) {
	r.mockDriveGetDriveImportTask = f
}

// UnMockDriveGetDriveImportTask un-mock DriveGetDriveImportTask method
func (r *Mock) UnMockDriveGetDriveImportTask() {
	r.mockDriveGetDriveImportTask = nil
}

// GetDriveImportTaskReq ...
type GetDriveImportTaskReq struct {
	Ticket string `path:"ticket" json:"-"` // 导入任务ID, 示例值: "6990281865xxxxxxxx7843"
}

// GetDriveImportTaskResp ...
type GetDriveImportTaskResp struct {
	Result *GetDriveImportTaskRespResult `json:"result,omitempty"` // 导入结果
}

// GetDriveImportTaskRespResult ...
type GetDriveImportTaskRespResult struct {
	Ticket      string   `json:"ticket,omitempty"`        // 任务ID
	Type        string   `json:"type,omitempty"`          // 导入目标云文档格式
	JobStatus   int64    `json:"job_status,omitempty"`    // 任务状态, 可选值有: 0: 成功, 1: 初始化, 2: 处理中, 3: 内部错误, 100: 导入文档已加密, 101: 内部错误, 102: 内部错误, 103: 内部错误, 104: 租户容量不足, 105: 文件夹节点太多, 106: 内部错误, 108: 处理超时, 109: 内部错误, 110: 无权限, 112: 格式不支持, 113: office格式不支持, 114: 内部错误, 115: 导入文件过大, 116: 目录无权限, 117: 目录已删除, 118: 导入文件和任务指定后缀不匹配, 119: 目录不存在, 120: 导入文件和任务指定文件类型不匹配, 121: 导入文件已过期, 122: 创建副本中禁止导出, 5000: 内部错误, 7000: docx block 数量超过系统上限, 7001: docx block 层级超过系统上线, 7002: docx block 大小超过系统上限
	JobErrorMsg string   `json:"job_error_msg,omitempty"` // 任务失败原因
	Token       string   `json:"token,omitempty"`         // 导入云文档Token
	URL         string   `json:"url,omitempty"`           // 导入云文档URL
	Extra       []string `json:"extra,omitempty"`         // 任务成功后的提示信息
}

// getDriveImportTaskResp ...
type getDriveImportTaskResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveImportTaskResp `json:"data,omitempty"`
}
