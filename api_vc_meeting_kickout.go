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

// KickoutVCMeeting 将参会人从会议中移除
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/kickout
func (r *VCService) KickoutVCMeeting(ctx context.Context, request *KickoutVCMeetingReq, options ...MethodOptionFunc) (*KickoutVCMeetingResp, *Response, error) {
	if r.cli.mock.mockVCKickoutVCMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#KickoutVCMeeting mock enable")
		return r.cli.mock.mockVCKickoutVCMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "KickoutVCMeeting",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/meetings/:meeting_id/kickout",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(kickoutVCMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCKickoutVCMeeting mock VCKickoutVCMeeting method
func (r *Mock) MockVCKickoutVCMeeting(f func(ctx context.Context, request *KickoutVCMeetingReq, options ...MethodOptionFunc) (*KickoutVCMeetingResp, *Response, error)) {
	r.mockVCKickoutVCMeeting = f
}

// UnMockVCKickoutVCMeeting un-mock VCKickoutVCMeeting method
func (r *Mock) UnMockVCKickoutVCMeeting() {
	r.mockVCKickoutVCMeeting = nil
}

// KickoutVCMeetingReq ...
type KickoutVCMeetingReq struct {
	MeetingID    string                            `path:"meeting_id" json:"-"`     // 会议ID, 示例值: "6911188411932033028"
	UserIDType   *IDType                           `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	KickoutUsers []*KickoutVCMeetingReqKickoutUser `json:"kickout_users,omitempty"` // 需移除的用户列表
}

// KickoutVCMeetingReqKickoutUser ...
type KickoutVCMeetingReqKickoutUser struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int64  `json:"user_type,omitempty"` // 用户类型, 示例值: 1, 可选值有: 1: lark用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
}

// KickoutVCMeetingResp ...
type KickoutVCMeetingResp struct {
	KickoutResults []*KickoutVCMeetingRespKickoutResult `json:"kickout_results,omitempty"` // 移除结果
}

// KickoutVCMeetingRespKickoutResult ...
type KickoutVCMeetingRespKickoutResult struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: 1: lark用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
	Result   int64  `json:"result,omitempty"`    // 移除结果, 可选值有: 1: 移除成功, 2: 移除失败
}

// kickoutVCMeetingResp ...
type kickoutVCMeetingResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *KickoutVCMeetingResp `json:"data,omitempty"`
}
