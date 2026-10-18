// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeOperationLogsResponseBody struct {


	PageSize *int32 `json:"pageSize,omitempty"`

	Page *int32 `json:"page,omitempty"`

	List *[]DescribeOperationLogsResponseList `json:"list,omitempty"`
}

func (s DescribeOperationLogsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DescribeOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s DescribeOperationLogsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeOperationLogsResponseBody) SetPageSize(v int32) *DescribeOperationLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOperationLogsResponseBody) SetPage(v int32) *DescribeOperationLogsResponseBody {
	s.Page = &v
	return s
}

func (s *DescribeOperationLogsResponseBody) SetList(v []DescribeOperationLogsResponseList) *DescribeOperationLogsResponseBody {
	s.List = &v
	return s
}


type DescribeOperationLogsResponseBodyBuilder struct {
	s *DescribeOperationLogsResponseBody
}

func NewDescribeOperationLogsResponseBodyBuilder() *DescribeOperationLogsResponseBodyBuilder {
	s := &DescribeOperationLogsResponseBody{}
	b := &DescribeOperationLogsResponseBodyBuilder{s: s}
	return b
}

func (b *DescribeOperationLogsResponseBodyBuilder) PageSize(v int32) *DescribeOperationLogsResponseBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *DescribeOperationLogsResponseBodyBuilder) Page(v int32) *DescribeOperationLogsResponseBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *DescribeOperationLogsResponseBodyBuilder) List(v []DescribeOperationLogsResponseList) *DescribeOperationLogsResponseBodyBuilder {
	b.s.List = &v
	return b
}

func (b *DescribeOperationLogsResponseBodyBuilder) Build() *DescribeOperationLogsResponseBody {
	return b.s
}


