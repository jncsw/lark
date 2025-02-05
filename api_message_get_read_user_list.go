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

// GetMessageReadUserList 查询消息的已读信息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 只能查询机器人自己发送, 且发送时间不超过7天的消息
// - 查询消息已读信息时机器人仍需要在会话内
// - 本接口不支持查询批量消息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users
func (r *MessageService) GetMessageReadUserList(ctx context.Context, request *GetMessageReadUserListReq, options ...MethodOptionFunc) (*GetMessageReadUserListResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessageReadUserList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetMessageReadUserList mock enable")
		return r.cli.mock.mockMessageGetMessageReadUserList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessageReadUserList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/read_users",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMessageReadUserListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessageReadUserList mock MessageGetMessageReadUserList method
func (r *Mock) MockMessageGetMessageReadUserList(f func(ctx context.Context, request *GetMessageReadUserListReq, options ...MethodOptionFunc) (*GetMessageReadUserListResp, *Response, error)) {
	r.mockMessageGetMessageReadUserList = f
}

// UnMockMessageGetMessageReadUserList un-mock MessageGetMessageReadUserList method
func (r *Mock) UnMockMessageGetMessageReadUserList() {
	r.mockMessageGetMessageReadUserList = nil
}

// GetMessageReadUserListReq ...
type GetMessageReadUserListReq struct {
	MessageID  string  `path:"message_id" json:"-"`    // 待查询的消息的ID, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 注意: 不支持查询批量消息, 示例值: "om_dc13264520392913993dd051dba21dcf"
	UserIDType IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageSize   *int64  `query:"page_size" json:"-"`    // 此次调用中使用的分页的大小, 示例值: 20, 取值范围: `1` ～ `100`
	PageToken  *string `query:"page_token" json:"-"`   // 下一页分页的token, 示例值: "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ["
}

// GetMessageReadUserListResp ...
type GetMessageReadUserListResp struct {
	Items     []*GetMessageReadUserListRespItem `json:"items,omitempty"`
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有下一页
	PageToken string                            `json:"page_token,omitempty"` // 下一页分页的token
}

// GetMessageReadUserListRespItem ...
type GetMessageReadUserListRespItem struct {
	UserIDType IDType `json:"user_id_type,omitempty"` // 用户id类型
	UserID     string `json:"user_id,omitempty"`      // 用户id
	Timestamp  string `json:"timestamp,omitempty"`    // 阅读时间
	TenantKey  string `json:"tenant_key,omitempty"`   // tenant key
}

// getMessageReadUserListResp ...
type getMessageReadUserListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageReadUserListResp `json:"data,omitempty"`
}
