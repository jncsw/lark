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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Attendance_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroupList(ctx, &lark.GetAttendanceGroupListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceGroup(ctx, &lark.CreateAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchAttendanceGroup(ctx, &lark.SearchAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftList(ctx, &lark.GetAttendanceShiftListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShift(ctx, &lark.GetAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftDetail(ctx, &lark.GetAttendanceShiftDetailReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserDailyShift(ctx, &lark.BatchCreateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsField(ctx, &lark.GetAttendanceUserStatsFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsView(ctx, &lark.GetAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserStatsView(ctx, &lark.UpdateAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsData(ctx, &lark.GetAttendanceUserStatsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedyAllowedRemedyList(ctx, &lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserTaskRemedy(ctx, &lark.CreateAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserSettingList(ctx, &lark.GetAttendanceUserSettingListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserSetting(ctx, &lark.UpdateAttendanceUserSettingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DownloadAttendanceFile(ctx, &lark.DownloadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceRemedyApproval(ctx, &lark.UpdateAttendanceRemedyApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceGroupList(func(ctx context.Context, request *lark.GetAttendanceGroupListReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceGroupListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceGroupList()

			_, _, err := moduleCli.GetAttendanceGroupList(ctx, &lark.GetAttendanceGroupListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceCreateAttendanceGroup(func(ctx context.Context, request *lark.CreateAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceGroup()

			_, _, err := moduleCli.CreateAttendanceGroup(ctx, &lark.CreateAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceSearchAttendanceGroup(func(ctx context.Context, request *lark.SearchAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.SearchAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceSearchAttendanceGroup()

			_, _, err := moduleCli.SearchAttendanceGroup(ctx, &lark.SearchAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceGroup(func(ctx context.Context, request *lark.GetAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceGroup()

			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceDeleteAttendanceGroup(func(ctx context.Context, request *lark.DeleteAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.DeleteAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteAttendanceGroup()

			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceShiftList(func(ctx context.Context, request *lark.GetAttendanceShiftListReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceShiftListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceShiftList()

			_, _, err := moduleCli.GetAttendanceShiftList(ctx, &lark.GetAttendanceShiftListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceShift(func(ctx context.Context, request *lark.GetAttendanceShiftReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceShift()

			_, _, err := moduleCli.GetAttendanceShift(ctx, &lark.GetAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceShiftDetail(func(ctx context.Context, request *lark.GetAttendanceShiftDetailReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceShiftDetailResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceShiftDetail()

			_, _, err := moduleCli.GetAttendanceShiftDetail(ctx, &lark.GetAttendanceShiftDetailReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceDeleteAttendanceShift(func(ctx context.Context, request *lark.DeleteAttendanceShiftReq, options ...lark.MethodOptionFunc) (*lark.DeleteAttendanceShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteAttendanceShift()

			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceCreateAttendanceShift(func(ctx context.Context, request *lark.CreateAttendanceShiftReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceShift()

			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserDailyShift(func(ctx context.Context, request *lark.GetAttendanceUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserDailyShift()

			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceBatchCreateAttendanceUserDailyShift(func(ctx context.Context, request *lark.BatchCreateAttendanceUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateAttendanceUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchCreateAttendanceUserDailyShift()

			_, _, err := moduleCli.BatchCreateAttendanceUserDailyShift(ctx, &lark.BatchCreateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserStatsField(func(ctx context.Context, request *lark.GetAttendanceUserStatsFieldReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserStatsFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserStatsField()

			_, _, err := moduleCli.GetAttendanceUserStatsField(ctx, &lark.GetAttendanceUserStatsFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserStatsView(func(ctx context.Context, request *lark.GetAttendanceUserStatsViewReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserStatsViewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserStatsView()

			_, _, err := moduleCli.GetAttendanceUserStatsView(ctx, &lark.GetAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceUpdateAttendanceUserStatsView(func(ctx context.Context, request *lark.UpdateAttendanceUserStatsViewReq, options ...lark.MethodOptionFunc) (*lark.UpdateAttendanceUserStatsViewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateAttendanceUserStatsView()

			_, _, err := moduleCli.UpdateAttendanceUserStatsView(ctx, &lark.UpdateAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserStatsData(func(ctx context.Context, request *lark.GetAttendanceUserStatsDataReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserStatsDataResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserStatsData()

			_, _, err := moduleCli.GetAttendanceUserStatsData(ctx, &lark.GetAttendanceUserStatsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceBatchGetAttendanceUserFlow(func(ctx context.Context, request *lark.BatchGetAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchGetAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchGetAttendanceUserFlow()

			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserFlow(func(ctx context.Context, request *lark.GetAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserFlow()

			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserTask(func(ctx context.Context, request *lark.GetAttendanceUserTaskReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserTask()

			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceBatchCreateAttendanceUserFlow(func(ctx context.Context, request *lark.BatchCreateAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchCreateAttendanceUserFlow()

			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList(func(ctx context.Context, request *lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserTaskRemedyAllowedRemedyListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList()

			_, _, err := moduleCli.GetAttendanceUserTaskRemedyAllowedRemedyList(ctx, &lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserTaskRemedy(func(ctx context.Context, request *lark.GetAttendanceUserTaskRemedyReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserTaskRemedyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserTaskRemedy()

			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceCreateAttendanceUserTaskRemedy(func(ctx context.Context, request *lark.CreateAttendanceUserTaskRemedyReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceUserTaskRemedyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceUserTaskRemedy()

			_, _, err := moduleCli.CreateAttendanceUserTaskRemedy(ctx, &lark.CreateAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserSettingList(func(ctx context.Context, request *lark.GetAttendanceUserSettingListReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserSettingListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserSettingList()

			_, _, err := moduleCli.GetAttendanceUserSettingList(ctx, &lark.GetAttendanceUserSettingListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceUpdateAttendanceUserSetting(func(ctx context.Context, request *lark.UpdateAttendanceUserSettingReq, options ...lark.MethodOptionFunc) (*lark.UpdateAttendanceUserSettingResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateAttendanceUserSetting()

			_, _, err := moduleCli.UpdateAttendanceUserSetting(ctx, &lark.UpdateAttendanceUserSettingReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceDownloadAttendanceFile(func(ctx context.Context, request *lark.DownloadAttendanceFileReq, options ...lark.MethodOptionFunc) (*lark.DownloadAttendanceFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDownloadAttendanceFile()

			_, _, err := moduleCli.DownloadAttendanceFile(ctx, &lark.DownloadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceUploadAttendanceFile(func(ctx context.Context, request *lark.UploadAttendanceFileReq, options ...lark.MethodOptionFunc) (*lark.UploadAttendanceFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUploadAttendanceFile()

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceGetAttendanceUserApproval(func(ctx context.Context, request *lark.GetAttendanceUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserApproval()

			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceCreateAttendanceUserApproval(func(ctx context.Context, request *lark.CreateAttendanceUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceUserApproval()

			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockAttendanceUpdateAttendanceRemedyApproval(func(ctx context.Context, request *lark.UpdateAttendanceRemedyApprovalReq, options ...lark.MethodOptionFunc) (*lark.UpdateAttendanceRemedyApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateAttendanceRemedyApproval()

			_, _, err := moduleCli.UpdateAttendanceRemedyApproval(ctx, &lark.UpdateAttendanceRemedyApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroupList(ctx, &lark.GetAttendanceGroupListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceGroup(ctx, &lark.CreateAttendanceGroupReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchAttendanceGroup(ctx, &lark.SearchAttendanceGroupReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftList(ctx, &lark.GetAttendanceShiftListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShift(ctx, &lark.GetAttendanceShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftDetail(ctx, &lark.GetAttendanceShiftDetailReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserDailyShift(ctx, &lark.BatchCreateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsField(ctx, &lark.GetAttendanceUserStatsFieldReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsView(ctx, &lark.GetAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserStatsView(ctx, &lark.UpdateAttendanceUserStatsViewReq{
				UserStatsViewID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsData(ctx, &lark.GetAttendanceUserStatsDataReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{
				UserFlowID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedyAllowedRemedyList(ctx, &lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserTaskRemedy(ctx, &lark.CreateAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserSettingList(ctx, &lark.GetAttendanceUserSettingListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserSetting(ctx, &lark.UpdateAttendanceUserSettingReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DownloadAttendanceFile(ctx, &lark.DownloadAttendanceFileReq{
				FileID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceRemedyApproval(ctx, &lark.UpdateAttendanceRemedyApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Attendance
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroupList(ctx, &lark.GetAttendanceGroupListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceGroup(ctx, &lark.CreateAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchAttendanceGroup(ctx, &lark.SearchAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftList(ctx, &lark.GetAttendanceShiftListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShift(ctx, &lark.GetAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceShiftDetail(ctx, &lark.GetAttendanceShiftDetailReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserDailyShift(ctx, &lark.BatchCreateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsField(ctx, &lark.GetAttendanceUserStatsFieldReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsView(ctx, &lark.GetAttendanceUserStatsViewReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserStatsView(ctx, &lark.UpdateAttendanceUserStatsViewReq{
				UserStatsViewID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserStatsData(ctx, &lark.GetAttendanceUserStatsDataReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{
				UserFlowID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedyAllowedRemedyList(ctx, &lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserTaskRemedy(ctx, &lark.CreateAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserSettingList(ctx, &lark.GetAttendanceUserSettingListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceUserSetting(ctx, &lark.UpdateAttendanceUserSettingReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DownloadAttendanceFile(ctx, &lark.DownloadAttendanceFileReq{
				FileID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateAttendanceRemedyApproval(ctx, &lark.UpdateAttendanceRemedyApprovalReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
