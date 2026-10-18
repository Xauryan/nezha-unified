// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeOperationLogsRequest struct {


	DescribeOperationLogsQuery *DescribeOperationLogsQuery `json:"describeOperationLogsQuery,omitempty"`
}

func (s DescribeOperationLogsRequest) String() string {
	return utils.Beautify(s)
}

func (s DescribeOperationLogsRequest) GoString() string {
	return s.String()
}

func (s DescribeOperationLogsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeOperationLogsRequest) SetDescribeOperationLogsQuery(v *DescribeOperationLogsQuery) *DescribeOperationLogsRequest {
	s.DescribeOperationLogsQuery = v
	return s
}


type DescribeOperationLogsRequestBuilder struct {
	s *DescribeOperationLogsRequest
}

func NewDescribeOperationLogsRequestBuilder() *DescribeOperationLogsRequestBuilder {
	s := &DescribeOperationLogsRequest{}
	b := &DescribeOperationLogsRequestBuilder{s: s}
	return b
}

func (b *DescribeOperationLogsRequestBuilder) DescribeOperationLogsQuery(v *DescribeOperationLogsQuery) *DescribeOperationLogsRequestBuilder {
	b.s.DescribeOperationLogsQuery = v
	return b
}

func (b *DescribeOperationLogsRequestBuilder) Build() *DescribeOperationLogsRequest {
	return b.s
}


