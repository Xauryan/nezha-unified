// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DescribeOperationLogsQueryOperateTypeEnum string

// List of OperateType
const (
    DescribeOperationLogsQueryOperateTypeEnumRefresh DescribeOperationLogsQueryOperateTypeEnum = "REFRESH"
    DescribeOperationLogsQueryOperateTypeEnumReload DescribeOperationLogsQueryOperateTypeEnum = "RELOAD"
)
type DescribeOperationLogsQueryRefreshTypeEnum string

// List of RefreshType
const (
    DescribeOperationLogsQueryRefreshTypeEnumFile DescribeOperationLogsQueryRefreshTypeEnum = "FILE"
    DescribeOperationLogsQueryRefreshTypeEnumCatalog DescribeOperationLogsQueryRefreshTypeEnum = "CATALOG"
)

type DescribeOperationLogsQuery struct {
    position.Query
    // 操作类型，预热或刷新
	OperateType *DescribeOperationLogsQueryOperateTypeEnum `json:"operateType,omitempty"`
    // 一页中个数，默认为10
	PageSize *string `json:"pageSize,omitempty"`
    // 刷新类型，URL或目录
	RefreshType *DescribeOperationLogsQueryRefreshTypeEnum `json:"refreshType,omitempty"`
    // 开始时间，格式yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间，格式yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"endTime,omitempty"`
    // 页数，默认为1
	Page *string `json:"page,omitempty"`
}

func (s DescribeOperationLogsQuery) String() string {
	return utils.Beautify(s)
}

func (s DescribeOperationLogsQuery) GoString() string {
	return s.String()
}

func (s DescribeOperationLogsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeOperationLogsQuery) SetOperateType(v DescribeOperationLogsQueryOperateTypeEnum) *DescribeOperationLogsQuery {
	s.OperateType = &v
	return s
}

func (s *DescribeOperationLogsQuery) SetPageSize(v string) *DescribeOperationLogsQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeOperationLogsQuery) SetRefreshType(v DescribeOperationLogsQueryRefreshTypeEnum) *DescribeOperationLogsQuery {
	s.RefreshType = &v
	return s
}

func (s *DescribeOperationLogsQuery) SetStartTime(v string) *DescribeOperationLogsQuery {
	s.StartTime = &v
	return s
}

func (s *DescribeOperationLogsQuery) SetEndTime(v string) *DescribeOperationLogsQuery {
	s.EndTime = &v
	return s
}

func (s *DescribeOperationLogsQuery) SetPage(v string) *DescribeOperationLogsQuery {
	s.Page = &v
	return s
}


type DescribeOperationLogsQueryBuilder struct {
	s *DescribeOperationLogsQuery
}

func NewDescribeOperationLogsQueryBuilder() *DescribeOperationLogsQueryBuilder {
	s := &DescribeOperationLogsQuery{}
	b := &DescribeOperationLogsQueryBuilder{s: s}
	return b
}

func (b *DescribeOperationLogsQueryBuilder) OperateType(v DescribeOperationLogsQueryOperateTypeEnum) *DescribeOperationLogsQueryBuilder {
	b.s.OperateType = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) PageSize(v string) *DescribeOperationLogsQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) RefreshType(v DescribeOperationLogsQueryRefreshTypeEnum) *DescribeOperationLogsQueryBuilder {
	b.s.RefreshType = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) StartTime(v string) *DescribeOperationLogsQueryBuilder {
	b.s.StartTime = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) EndTime(v string) *DescribeOperationLogsQueryBuilder {
	b.s.EndTime = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) Page(v string) *DescribeOperationLogsQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *DescribeOperationLogsQueryBuilder) Build() *DescribeOperationLogsQuery {
	return b.s
}


