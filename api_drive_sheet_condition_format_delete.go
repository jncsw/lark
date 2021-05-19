// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteSheetConditionFormat 该接口用于删除已有的条件格式，单次最多支持删除10个条件格式，每个条件格式的删除会返回成功或者失败，失败的情况包括各种参数的校验。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-delete
func (r *DriveService) DeleteSheetConditionFormat(ctx context.Context, request *DeleteSheetConditionFormatReq, options ...MethodOptionFunc) (*DeleteSheetConditionFormatResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetConditionFormat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetConditionFormat mock enable")
		return r.cli.mock.mockDriveDeleteSheetConditionFormat(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#DeleteSheetConditionFormat call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetConditionFormat request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(deleteSheetConditionFormatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#DeleteSheetConditionFormat DELETE https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#DeleteSheetConditionFormat DELETE https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "DeleteSheetConditionFormat", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetConditionFormat success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveDeleteSheetConditionFormat(f func(ctx context.Context, request *DeleteSheetConditionFormatReq, options ...MethodOptionFunc) (*DeleteSheetConditionFormatResp, *Response, error)) {
	r.mockDriveDeleteSheetConditionFormat = f
}

func (r *Mock) UnMockDriveDeleteSheetConditionFormat() {
	r.mockDriveDeleteSheetConditionFormat = nil
}

type DeleteSheetConditionFormatReq struct {
	SpreadsheetToken string                                   `path:"spreadsheetToken" json:"-"` // sheet 的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	SheetCfIDs       *DeleteSheetConditionFormatReqSheetCfIDs `json:"sheet_cf_ids,omitempty"`    // 表格条件格式id
}

type DeleteSheetConditionFormatReqSheetCfIDs struct {
	SheetID string `json:"sheet_id,omitempty"` // sheet的id
	CfID    string `json:"cf_id,omitempty"`    // 条件格式id
}

type deleteSheetConditionFormatResp struct {
	Code int64                           `json:"code,omitempty"`
	Msg  string                          `json:"msg,omitempty"`
	Data *DeleteSheetConditionFormatResp `json:"data,omitempty"`
}

type DeleteSheetConditionFormatResp struct {
	Responses *DeleteSheetConditionFormatRespResponses `json:"responses,omitempty"` // 响应
}

type DeleteSheetConditionFormatRespResponses struct {
	SheetID string `json:"sheet_id,omitempty"` // sheet的Id
	CfID    string `json:"cf_id,omitempty"`    // 条件格式id
	ResCode int64  `json:"res_code,omitempty"` // 条件格式删除状态码，0表示成功，非0表示失败
	ResMsg  string `json:"res_msg,omitempty"`  // 条件格式删除返回的状态信息，空表示成功，非空表示失败原因
}
