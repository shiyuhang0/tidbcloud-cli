// Code generated by mockery v2.43.0. DO NOT EDIT.

package mock

import (
	br "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/br"
	branch "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/branch"

	cluster "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/cluster"

	context "context"

	export "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/export"

	iam "tidbcloud-cli/pkg/tidbcloud/v1beta1/iam"

	imp "tidbcloud-cli/pkg/tidbcloud/v1beta1/serverless/import"

	mock "github.com/stretchr/testify/mock"

	pingchat "tidbcloud-cli/pkg/tidbcloud/pingchat"
)

// TiDBCloudClient is an autogenerated mock type for the TiDBCloudClient type
type TiDBCloudClient struct {
	mock.Mock
}

// CancelExport provides a mock function with given fields: ctx, clusterId, exportId
func (_m *TiDBCloudClient) CancelExport(ctx context.Context, clusterId string, exportId string) (*export.Export, error) {
	ret := _m.Called(ctx, clusterId, exportId)

	if len(ret) == 0 {
		panic("no return value specified for CancelExport")
	}

	var r0 *export.Export
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*export.Export, error)); ok {
		return rf(ctx, clusterId, exportId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *export.Export); ok {
		r0 = rf(ctx, clusterId, exportId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.Export)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, exportId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelImport provides a mock function with given fields: ctx, clusterId, id
func (_m *TiDBCloudClient) CancelImport(ctx context.Context, clusterId string, id string) error {
	ret := _m.Called(ctx, clusterId, id)

	if len(ret) == 0 {
		panic("no return value specified for CancelImport")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, clusterId, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelUpload provides a mock function with given fields: ctx, clusterId, uploadId
func (_m *TiDBCloudClient) CancelUpload(ctx context.Context, clusterId string, uploadId *string) error {
	ret := _m.Called(ctx, clusterId, uploadId)

	if len(ret) == 0 {
		panic("no return value specified for CancelUpload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string) error); ok {
		r0 = rf(ctx, clusterId, uploadId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chat provides a mock function with given fields: ctx, chatInfo
func (_m *TiDBCloudClient) Chat(ctx context.Context, chatInfo *pingchat.PingchatChatInfo) (*pingchat.PingchatChatResponse, error) {
	ret := _m.Called(ctx, chatInfo)

	if len(ret) == 0 {
		panic("no return value specified for Chat")
	}

	var r0 *pingchat.PingchatChatResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pingchat.PingchatChatInfo) (*pingchat.PingchatChatResponse, error)); ok {
		return rf(ctx, chatInfo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pingchat.PingchatChatInfo) *pingchat.PingchatChatResponse); ok {
		r0 = rf(ctx, chatInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pingchat.PingchatChatResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pingchat.PingchatChatInfo) error); ok {
		r1 = rf(ctx, chatInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteUpload provides a mock function with given fields: ctx, clusterId, uploadId, parts
func (_m *TiDBCloudClient) CompleteUpload(ctx context.Context, clusterId string, uploadId *string, parts *[]imp.CompletePart) error {
	ret := _m.Called(ctx, clusterId, uploadId, parts)

	if len(ret) == 0 {
		panic("no return value specified for CompleteUpload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *[]imp.CompletePart) error); ok {
		r0 = rf(ctx, clusterId, uploadId, parts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBranch provides a mock function with given fields: ctx, clusterId, body
func (_m *TiDBCloudClient) CreateBranch(ctx context.Context, clusterId string, body *branch.Branch) (*branch.Branch, error) {
	ret := _m.Called(ctx, clusterId, body)

	if len(ret) == 0 {
		panic("no return value specified for CreateBranch")
	}

	var r0 *branch.Branch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *branch.Branch) (*branch.Branch, error)); ok {
		return rf(ctx, clusterId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *branch.Branch) *branch.Branch); ok {
		r0 = rf(ctx, clusterId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*branch.Branch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *branch.Branch) error); ok {
		r1 = rf(ctx, clusterId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCluster provides a mock function with given fields: ctx, body
func (_m *TiDBCloudClient) CreateCluster(ctx context.Context, body *cluster.TidbCloudOpenApiserverlessv1beta1Cluster) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ret := _m.Called(ctx, body)

	if len(ret) == 0 {
		panic("no return value specified for CreateCluster")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1Cluster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.TidbCloudOpenApiserverlessv1beta1Cluster) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.TidbCloudOpenApiserverlessv1beta1Cluster) *cluster.TidbCloudOpenApiserverlessv1beta1Cluster); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1Cluster)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cluster.TidbCloudOpenApiserverlessv1beta1Cluster) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateExport provides a mock function with given fields: ctx, clusterId, body
func (_m *TiDBCloudClient) CreateExport(ctx context.Context, clusterId string, body *export.ExportServiceCreateExportBody) (*export.Export, error) {
	ret := _m.Called(ctx, clusterId, body)

	if len(ret) == 0 {
		panic("no return value specified for CreateExport")
	}

	var r0 *export.Export
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *export.ExportServiceCreateExportBody) (*export.Export, error)); ok {
		return rf(ctx, clusterId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *export.ExportServiceCreateExportBody) *export.Export); ok {
		r0 = rf(ctx, clusterId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.Export)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *export.ExportServiceCreateExportBody) error); ok {
		r1 = rf(ctx, clusterId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateImport provides a mock function with given fields: ctx, clusterId, body
func (_m *TiDBCloudClient) CreateImport(ctx context.Context, clusterId string, body *imp.ImportServiceCreateImportBody) (*imp.Import, error) {
	ret := _m.Called(ctx, clusterId, body)

	if len(ret) == 0 {
		panic("no return value specified for CreateImport")
	}

	var r0 *imp.Import
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *imp.ImportServiceCreateImportBody) (*imp.Import, error)); ok {
		return rf(ctx, clusterId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *imp.ImportServiceCreateImportBody) *imp.Import); ok {
		r0 = rf(ctx, clusterId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*imp.Import)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *imp.ImportServiceCreateImportBody) error); ok {
		r1 = rf(ctx, clusterId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSQLUser provides a mock function with given fields: ctx, clusterID, body
func (_m *TiDBCloudClient) CreateSQLUser(ctx context.Context, clusterID string, body *iam.ApiCreateSqlUserReq) (*iam.ApiSqlUser, error) {
	ret := _m.Called(ctx, clusterID, body)

	if len(ret) == 0 {
		panic("no return value specified for CreateSQLUser")
	}

	var r0 *iam.ApiSqlUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *iam.ApiCreateSqlUserReq) (*iam.ApiSqlUser, error)); ok {
		return rf(ctx, clusterID, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *iam.ApiCreateSqlUserReq) *iam.ApiSqlUser); ok {
		r0 = rf(ctx, clusterID, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiSqlUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *iam.ApiCreateSqlUserReq) error); ok {
		r1 = rf(ctx, clusterID, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBackup provides a mock function with given fields: ctx, backupId
func (_m *TiDBCloudClient) DeleteBackup(ctx context.Context, backupId string) (*br.V1beta1Backup, error) {
	ret := _m.Called(ctx, backupId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBackup")
	}

	var r0 *br.V1beta1Backup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*br.V1beta1Backup, error)); ok {
		return rf(ctx, backupId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *br.V1beta1Backup); ok {
		r0 = rf(ctx, backupId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*br.V1beta1Backup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, backupId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBranch provides a mock function with given fields: ctx, clusterId, branchId
func (_m *TiDBCloudClient) DeleteBranch(ctx context.Context, clusterId string, branchId string) (*branch.Branch, error) {
	ret := _m.Called(ctx, clusterId, branchId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBranch")
	}

	var r0 *branch.Branch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*branch.Branch, error)); ok {
		return rf(ctx, clusterId, branchId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *branch.Branch); ok {
		r0 = rf(ctx, clusterId, branchId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*branch.Branch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, branchId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCluster provides a mock function with given fields: ctx, clusterId
func (_m *TiDBCloudClient) DeleteCluster(ctx context.Context, clusterId string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ret := _m.Called(ctx, clusterId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCluster")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1Cluster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error)); ok {
		return rf(ctx, clusterId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *cluster.TidbCloudOpenApiserverlessv1beta1Cluster); ok {
		r0 = rf(ctx, clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1Cluster)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExport provides a mock function with given fields: ctx, clusterId, exportId
func (_m *TiDBCloudClient) DeleteExport(ctx context.Context, clusterId string, exportId string) (*export.Export, error) {
	ret := _m.Called(ctx, clusterId, exportId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExport")
	}

	var r0 *export.Export
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*export.Export, error)); ok {
		return rf(ctx, clusterId, exportId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *export.Export); ok {
		r0 = rf(ctx, clusterId, exportId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.Export)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, exportId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSQLUser provides a mock function with given fields: ctx, clusterID, userName
func (_m *TiDBCloudClient) DeleteSQLUser(ctx context.Context, clusterID string, userName string) (*iam.ApiBasicResp, error) {
	ret := _m.Called(ctx, clusterID, userName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSQLUser")
	}

	var r0 *iam.ApiBasicResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*iam.ApiBasicResp, error)); ok {
		return rf(ctx, clusterID, userName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *iam.ApiBasicResp); ok {
		r0 = rf(ctx, clusterID, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiBasicResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterID, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadExport provides a mock function with given fields: ctx, clusterId, exportId
func (_m *TiDBCloudClient) DownloadExport(ctx context.Context, clusterId string, exportId string) (*export.DownloadExportsResponse, error) {
	ret := _m.Called(ctx, clusterId, exportId)

	if len(ret) == 0 {
		panic("no return value specified for DownloadExport")
	}

	var r0 *export.DownloadExportsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*export.DownloadExportsResponse, error)); ok {
		return rf(ctx, clusterId, exportId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *export.DownloadExportsResponse); ok {
		r0 = rf(ctx, clusterId, exportId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.DownloadExportsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, exportId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackup provides a mock function with given fields: ctx, backupId
func (_m *TiDBCloudClient) GetBackup(ctx context.Context, backupId string) (*br.V1beta1Backup, error) {
	ret := _m.Called(ctx, backupId)

	if len(ret) == 0 {
		panic("no return value specified for GetBackup")
	}

	var r0 *br.V1beta1Backup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*br.V1beta1Backup, error)); ok {
		return rf(ctx, backupId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *br.V1beta1Backup); ok {
		r0 = rf(ctx, backupId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*br.V1beta1Backup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, backupId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBranch provides a mock function with given fields: ctx, clusterId, branchId
func (_m *TiDBCloudClient) GetBranch(ctx context.Context, clusterId string, branchId string) (*branch.Branch, error) {
	ret := _m.Called(ctx, clusterId, branchId)

	if len(ret) == 0 {
		panic("no return value specified for GetBranch")
	}

	var r0 *branch.Branch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*branch.Branch, error)); ok {
		return rf(ctx, clusterId, branchId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *branch.Branch); ok {
		r0 = rf(ctx, clusterId, branchId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*branch.Branch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, branchId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: ctx, clusterId
func (_m *TiDBCloudClient) GetCluster(ctx context.Context, clusterId string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ret := _m.Called(ctx, clusterId)

	if len(ret) == 0 {
		panic("no return value specified for GetCluster")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1Cluster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error)); ok {
		return rf(ctx, clusterId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *cluster.TidbCloudOpenApiserverlessv1beta1Cluster); ok {
		r0 = rf(ctx, clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1Cluster)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExport provides a mock function with given fields: ctx, clusterId, exportId
func (_m *TiDBCloudClient) GetExport(ctx context.Context, clusterId string, exportId string) (*export.Export, error) {
	ret := _m.Called(ctx, clusterId, exportId)

	if len(ret) == 0 {
		panic("no return value specified for GetExport")
	}

	var r0 *export.Export
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*export.Export, error)); ok {
		return rf(ctx, clusterId, exportId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *export.Export); ok {
		r0 = rf(ctx, clusterId, exportId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.Export)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, exportId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImport provides a mock function with given fields: ctx, clusterId, id
func (_m *TiDBCloudClient) GetImport(ctx context.Context, clusterId string, id string) (*imp.Import, error) {
	ret := _m.Called(ctx, clusterId, id)

	if len(ret) == 0 {
		panic("no return value specified for GetImport")
	}

	var r0 *imp.Import
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*imp.Import, error)); ok {
		return rf(ctx, clusterId, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *imp.Import); ok {
		r0 = rf(ctx, clusterId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*imp.Import)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterId, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSQLUser provides a mock function with given fields: ctx, clusterID, userName
func (_m *TiDBCloudClient) GetSQLUser(ctx context.Context, clusterID string, userName string) (*iam.ApiSqlUser, error) {
	ret := _m.Called(ctx, clusterID, userName)

	if len(ret) == 0 {
		panic("no return value specified for GetSQLUser")
	}

	var r0 *iam.ApiSqlUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*iam.ApiSqlUser, error)); ok {
		return rf(ctx, clusterID, userName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *iam.ApiSqlUser); ok {
		r0 = rf(ctx, clusterID, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiSqlUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterID, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBackups provides a mock function with given fields: ctx, clusterId, pageSize, pageToken
func (_m *TiDBCloudClient) ListBackups(ctx context.Context, clusterId *string, pageSize *int32, pageToken *string) (*br.V1beta1ListBackupsResponse, error) {
	ret := _m.Called(ctx, clusterId, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListBackups")
	}

	var r0 *br.V1beta1ListBackupsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, *int32, *string) (*br.V1beta1ListBackupsResponse, error)); ok {
		return rf(ctx, clusterId, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string, *int32, *string) *br.V1beta1ListBackupsResponse); ok {
		r0 = rf(ctx, clusterId, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*br.V1beta1ListBackupsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string, *int32, *string) error); ok {
		r1 = rf(ctx, clusterId, pageSize, pageToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBranches provides a mock function with given fields: ctx, clusterId, pageSize, pageToken
func (_m *TiDBCloudClient) ListBranches(ctx context.Context, clusterId string, pageSize *int32, pageToken *string) (*branch.ListBranchesResponse, error) {
	ret := _m.Called(ctx, clusterId, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListBranches")
	}

	var r0 *branch.ListBranchesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string) (*branch.ListBranchesResponse, error)); ok {
		return rf(ctx, clusterId, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string) *branch.ListBranchesResponse); ok {
		r0 = rf(ctx, clusterId, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*branch.ListBranchesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *int32, *string) error); ok {
		r1 = rf(ctx, clusterId, pageSize, pageToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: ctx, filter, pageSize, pageToken, orderBy, skip
func (_m *TiDBCloudClient) ListClusters(ctx context.Context, filter *string, pageSize *int32, pageToken *string, orderBy *string, skip *int32) (*cluster.TidbCloudOpenApiserverlessv1beta1ListClustersResponse, error) {
	ret := _m.Called(ctx, filter, pageSize, pageToken, orderBy, skip)

	if len(ret) == 0 {
		panic("no return value specified for ListClusters")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1ListClustersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, *int32, *string, *string, *int32) (*cluster.TidbCloudOpenApiserverlessv1beta1ListClustersResponse, error)); ok {
		return rf(ctx, filter, pageSize, pageToken, orderBy, skip)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string, *int32, *string, *string, *int32) *cluster.TidbCloudOpenApiserverlessv1beta1ListClustersResponse); ok {
		r0 = rf(ctx, filter, pageSize, pageToken, orderBy, skip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1ListClustersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string, *int32, *string, *string, *int32) error); ok {
		r1 = rf(ctx, filter, pageSize, pageToken, orderBy, skip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExportFiles provides a mock function with given fields: ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl
func (_m *TiDBCloudClient) ListExportFiles(ctx context.Context, clusterId string, exportId string, pageSize *int32, pageToken *string, isGenerateUrl bool) (*export.ListExportFilesResponse, error) {
	ret := _m.Called(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)

	if len(ret) == 0 {
		panic("no return value specified for ListExportFiles")
	}

	var r0 *export.ListExportFilesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, *string, bool) (*export.ListExportFilesResponse, error)); ok {
		return rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, *string, bool) *export.ListExportFilesResponse); ok {
		r0 = rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.ListExportFilesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *int32, *string, bool) error); ok {
		r1 = rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExportFilesWithRetry provides a mock function with given fields: ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl
func (_m *TiDBCloudClient) ListExportFilesWithRetry(ctx context.Context, clusterId string, exportId string, pageSize *int32, pageToken *string, isGenerateUrl bool) (*export.ListExportFilesResponse, error) {
	ret := _m.Called(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)

	if len(ret) == 0 {
		panic("no return value specified for ListExportFilesWithRetry")
	}

	var r0 *export.ListExportFilesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, *string, bool) (*export.ListExportFilesResponse, error)); ok {
		return rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, *string, bool) *export.ListExportFilesResponse); ok {
		r0 = rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.ListExportFilesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *int32, *string, bool) error); ok {
		r1 = rf(ctx, clusterId, exportId, pageSize, pageToken, isGenerateUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExports provides a mock function with given fields: ctx, clusterId, pageSize, pageToken, orderBy
func (_m *TiDBCloudClient) ListExports(ctx context.Context, clusterId string, pageSize *int32, pageToken *string, orderBy *string) (*export.ListExportsResponse, error) {
	ret := _m.Called(ctx, clusterId, pageSize, pageToken, orderBy)

	if len(ret) == 0 {
		panic("no return value specified for ListExports")
	}

	var r0 *export.ListExportsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string, *string) (*export.ListExportsResponse, error)); ok {
		return rf(ctx, clusterId, pageSize, pageToken, orderBy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string, *string) *export.ListExportsResponse); ok {
		r0 = rf(ctx, clusterId, pageSize, pageToken, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*export.ListExportsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *int32, *string, *string) error); ok {
		r1 = rf(ctx, clusterId, pageSize, pageToken, orderBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImports provides a mock function with given fields: ctx, clusterId, pageSize, pageToken, orderBy
func (_m *TiDBCloudClient) ListImports(ctx context.Context, clusterId string, pageSize *int32, pageToken *string, orderBy *string) (*imp.ListImportsResp, error) {
	ret := _m.Called(ctx, clusterId, pageSize, pageToken, orderBy)

	if len(ret) == 0 {
		panic("no return value specified for ListImports")
	}

	var r0 *imp.ListImportsResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string, *string) (*imp.ListImportsResp, error)); ok {
		return rf(ctx, clusterId, pageSize, pageToken, orderBy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string, *string) *imp.ListImportsResp); ok {
		r0 = rf(ctx, clusterId, pageSize, pageToken, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*imp.ListImportsResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *int32, *string, *string) error); ok {
		r1 = rf(ctx, clusterId, pageSize, pageToken, orderBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *TiDBCloudClient) ListProjects(ctx context.Context, pageSize *int32, pageToken *string) (*iam.ApiListProjectsRsp, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListProjects")
	}

	var r0 *iam.ApiListProjectsRsp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int32, *string) (*iam.ApiListProjectsRsp, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int32, *string) *iam.ApiListProjectsRsp); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiListProjectsRsp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int32, *string) error); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviderRegions provides a mock function with given fields: ctx
func (_m *TiDBCloudClient) ListProviderRegions(ctx context.Context) (*cluster.TidbCloudOpenApiserverlessv1beta1ListRegionsResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListProviderRegions")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1ListRegionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*cluster.TidbCloudOpenApiserverlessv1beta1ListRegionsResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *cluster.TidbCloudOpenApiserverlessv1beta1ListRegionsResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1ListRegionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSQLUsers provides a mock function with given fields: ctx, clusterID, pageSize, pageToken
func (_m *TiDBCloudClient) ListSQLUsers(ctx context.Context, clusterID string, pageSize *int32, pageToken *string) (*iam.ApiListSqlUsersRsp, error) {
	ret := _m.Called(ctx, clusterID, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListSQLUsers")
	}

	var r0 *iam.ApiListSqlUsersRsp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string) (*iam.ApiListSqlUsersRsp, error)); ok {
		return rf(ctx, clusterID, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32, *string) *iam.ApiListSqlUsersRsp); ok {
		r0 = rf(ctx, clusterID, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiListSqlUsersRsp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *int32, *string) error); ok {
		r1 = rf(ctx, clusterID, pageSize, pageToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PartialUpdateCluster provides a mock function with given fields: ctx, clusterId, body
func (_m *TiDBCloudClient) PartialUpdateCluster(ctx context.Context, clusterId string, body *cluster.V1beta1ServerlessServicePartialUpdateClusterBody) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error) {
	ret := _m.Called(ctx, clusterId, body)

	if len(ret) == 0 {
		panic("no return value specified for PartialUpdateCluster")
	}

	var r0 *cluster.TidbCloudOpenApiserverlessv1beta1Cluster
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.V1beta1ServerlessServicePartialUpdateClusterBody) (*cluster.TidbCloudOpenApiserverlessv1beta1Cluster, error)); ok {
		return rf(ctx, clusterId, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *cluster.V1beta1ServerlessServicePartialUpdateClusterBody) *cluster.TidbCloudOpenApiserverlessv1beta1Cluster); ok {
		r0 = rf(ctx, clusterId, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.TidbCloudOpenApiserverlessv1beta1Cluster)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *cluster.V1beta1ServerlessServicePartialUpdateClusterBody) error); ok {
		r1 = rf(ctx, clusterId, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Restore provides a mock function with given fields: ctx, body
func (_m *TiDBCloudClient) Restore(ctx context.Context, body *br.V1beta1RestoreRequest) (*br.V1beta1RestoreResponse, error) {
	ret := _m.Called(ctx, body)

	if len(ret) == 0 {
		panic("no return value specified for Restore")
	}

	var r0 *br.V1beta1RestoreResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *br.V1beta1RestoreRequest) (*br.V1beta1RestoreResponse, error)); ok {
		return rf(ctx, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *br.V1beta1RestoreRequest) *br.V1beta1RestoreResponse); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*br.V1beta1RestoreResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *br.V1beta1RestoreRequest) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartUpload provides a mock function with given fields: ctx, clusterId, fileName, targetDatabase, targetTable, partNumber
func (_m *TiDBCloudClient) StartUpload(ctx context.Context, clusterId string, fileName *string, targetDatabase *string, targetTable *string, partNumber *int32) (*imp.StartUploadResponse, error) {
	ret := _m.Called(ctx, clusterId, fileName, targetDatabase, targetTable, partNumber)

	if len(ret) == 0 {
		panic("no return value specified for StartUpload")
	}

	var r0 *imp.StartUploadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *string, *string, *int32) (*imp.StartUploadResponse, error)); ok {
		return rf(ctx, clusterId, fileName, targetDatabase, targetTable, partNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *string, *string, *int32) *imp.StartUploadResponse); ok {
		r0 = rf(ctx, clusterId, fileName, targetDatabase, targetTable, partNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*imp.StartUploadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *string, *string, *string, *int32) error); ok {
		r1 = rf(ctx, clusterId, fileName, targetDatabase, targetTable, partNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSQLUser provides a mock function with given fields: ctx, clusterID, userName, body
func (_m *TiDBCloudClient) UpdateSQLUser(ctx context.Context, clusterID string, userName string, body *iam.ApiUpdateSqlUserReq) (*iam.ApiSqlUser, error) {
	ret := _m.Called(ctx, clusterID, userName, body)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSQLUser")
	}

	var r0 *iam.ApiSqlUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *iam.ApiUpdateSqlUserReq) (*iam.ApiSqlUser, error)); ok {
		return rf(ctx, clusterID, userName, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *iam.ApiUpdateSqlUserReq) *iam.ApiSqlUser); ok {
		r0 = rf(ctx, clusterID, userName, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ApiSqlUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *iam.ApiUpdateSqlUserReq) error); ok {
		r1 = rf(ctx, clusterID, userName, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTiDBCloudClient creates a new instance of TiDBCloudClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTiDBCloudClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *TiDBCloudClient {
	mock := &TiDBCloudClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
