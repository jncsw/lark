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

func Test_Bitable_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Bitable

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableViewList(ctx, &lark.GetBitableViewListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableView(ctx, &lark.CreateBitableViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecord(ctx, &lark.GetBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableFieldList(ctx, &lark.GetBitableFieldListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableField(ctx, &lark.CreateBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableField(ctx, &lark.UpdateBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableTableList(ctx, &lark.GetBitableTableListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableTable(ctx, &lark.CreateBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableMeta(ctx, &lark.GetBitableMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Bitable

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableViewList(func(ctx context.Context, request *lark.GetBitableViewListReq, options ...lark.MethodOptionFunc) (*lark.GetBitableViewListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableViewList()

			_, _, err := moduleCli.GetBitableViewList(ctx, &lark.GetBitableViewListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableCreateBitableView(func(ctx context.Context, request *lark.CreateBitableViewReq, options ...lark.MethodOptionFunc) (*lark.CreateBitableViewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableCreateBitableView()

			_, _, err := moduleCli.CreateBitableView(ctx, &lark.CreateBitableViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableDeleteBitableView(func(ctx context.Context, request *lark.DeleteBitableViewReq, options ...lark.MethodOptionFunc) (*lark.DeleteBitableViewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableDeleteBitableView()

			_, _, err := moduleCli.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableRecordList(func(ctx context.Context, request *lark.GetBitableRecordListReq, options ...lark.MethodOptionFunc) (*lark.GetBitableRecordListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableRecordList()

			_, _, err := moduleCli.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableRecord(func(ctx context.Context, request *lark.GetBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.GetBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableRecord()

			_, _, err := moduleCli.GetBitableRecord(ctx, &lark.GetBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableCreateBitableRecord(func(ctx context.Context, request *lark.CreateBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.CreateBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableCreateBitableRecord()

			_, _, err := moduleCli.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableBatchCreateBitableRecord(func(ctx context.Context, request *lark.BatchCreateBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableBatchCreateBitableRecord()

			_, _, err := moduleCli.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableUpdateBitableRecord(func(ctx context.Context, request *lark.UpdateBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.UpdateBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableUpdateBitableRecord()

			_, _, err := moduleCli.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableBatchUpdateBitableRecord(func(ctx context.Context, request *lark.BatchUpdateBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.BatchUpdateBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableBatchUpdateBitableRecord()

			_, _, err := moduleCli.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableDeleteBitableRecord(func(ctx context.Context, request *lark.DeleteBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.DeleteBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableDeleteBitableRecord()

			_, _, err := moduleCli.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableBatchDeleteBitableRecord(func(ctx context.Context, request *lark.BatchDeleteBitableRecordReq, options ...lark.MethodOptionFunc) (*lark.BatchDeleteBitableRecordResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableBatchDeleteBitableRecord()

			_, _, err := moduleCli.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableFieldList(func(ctx context.Context, request *lark.GetBitableFieldListReq, options ...lark.MethodOptionFunc) (*lark.GetBitableFieldListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableFieldList()

			_, _, err := moduleCli.GetBitableFieldList(ctx, &lark.GetBitableFieldListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableCreateBitableField(func(ctx context.Context, request *lark.CreateBitableFieldReq, options ...lark.MethodOptionFunc) (*lark.CreateBitableFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableCreateBitableField()

			_, _, err := moduleCli.CreateBitableField(ctx, &lark.CreateBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableUpdateBitableField(func(ctx context.Context, request *lark.UpdateBitableFieldReq, options ...lark.MethodOptionFunc) (*lark.UpdateBitableFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableUpdateBitableField()

			_, _, err := moduleCli.UpdateBitableField(ctx, &lark.UpdateBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableDeleteBitableField(func(ctx context.Context, request *lark.DeleteBitableFieldReq, options ...lark.MethodOptionFunc) (*lark.DeleteBitableFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableDeleteBitableField()

			_, _, err := moduleCli.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableTableList(func(ctx context.Context, request *lark.GetBitableTableListReq, options ...lark.MethodOptionFunc) (*lark.GetBitableTableListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableTableList()

			_, _, err := moduleCli.GetBitableTableList(ctx, &lark.GetBitableTableListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableCreateBitableTable(func(ctx context.Context, request *lark.CreateBitableTableReq, options ...lark.MethodOptionFunc) (*lark.CreateBitableTableResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableCreateBitableTable()

			_, _, err := moduleCli.CreateBitableTable(ctx, &lark.CreateBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableBatchCreateBitableTable(func(ctx context.Context, request *lark.BatchCreateBitableTableReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateBitableTableResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableBatchCreateBitableTable()

			_, _, err := moduleCli.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableDeleteBitableTable(func(ctx context.Context, request *lark.DeleteBitableTableReq, options ...lark.MethodOptionFunc) (*lark.DeleteBitableTableResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableDeleteBitableTable()

			_, _, err := moduleCli.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableBatchDeleteBitableTable(func(ctx context.Context, request *lark.BatchDeleteBitableTableReq, options ...lark.MethodOptionFunc) (*lark.BatchDeleteBitableTableResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableBatchDeleteBitableTable()

			_, _, err := moduleCli.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableUpdateBitableMeta(func(ctx context.Context, request *lark.UpdateBitableMetaReq, options ...lark.MethodOptionFunc) (*lark.UpdateBitableMetaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableUpdateBitableMeta()

			_, _, err := moduleCli.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockBitableGetBitableMeta(func(ctx context.Context, request *lark.GetBitableMetaReq, options ...lark.MethodOptionFunc) (*lark.GetBitableMetaResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBitableGetBitableMeta()

			_, _, err := moduleCli.GetBitableMeta(ctx, &lark.GetBitableMetaReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Bitable

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableViewList(ctx, &lark.GetBitableViewListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableView(ctx, &lark.CreateBitableViewReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{
				AppToken: "x",
				TableID:  "x",
				ViewID:   "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecord(ctx, &lark.GetBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableFieldList(ctx, &lark.GetBitableFieldListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableField(ctx, &lark.CreateBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableField(ctx, &lark.UpdateBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
				FieldID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
				FieldID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableTableList(ctx, &lark.GetBitableTableListReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableTable(ctx, &lark.CreateBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Bitable
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableViewList(ctx, &lark.GetBitableViewListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableView(ctx, &lark.CreateBitableViewReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{
				AppToken: "x",
				TableID:  "x",
				ViewID:   "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableRecord(ctx, &lark.GetBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
				RecordID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableFieldList(ctx, &lark.GetBitableFieldListReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableField(ctx, &lark.CreateBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableField(ctx, &lark.UpdateBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
				FieldID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{
				AppToken: "x",
				TableID:  "x",
				FieldID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableTableList(ctx, &lark.GetBitableTableListReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateBitableTable(ctx, &lark.CreateBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{
				AppToken: "x",
				TableID:  "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
				AppToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
