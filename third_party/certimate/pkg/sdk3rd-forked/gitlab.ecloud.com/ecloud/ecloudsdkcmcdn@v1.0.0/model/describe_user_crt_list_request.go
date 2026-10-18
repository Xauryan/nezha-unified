// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserCrtListRequest struct {


	DescribeUserCrtListQuery *DescribeUserCrtListQuery `json:"describeUserCrtListQuery,omitempty"`
}

func (s DescribeUserCrtListRequest) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserCrtListRequest) GoString() string {
	return s.String()
}

func (s DescribeUserCrtListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserCrtListRequest) SetDescribeUserCrtListQuery(v *DescribeUserCrtListQuery) *DescribeUserCrtListRequest {
	s.DescribeUserCrtListQuery = v
	return s
}


type DescribeUserCrtListRequestBuilder struct {
	s *DescribeUserCrtListRequest
}

func NewDescribeUserCrtListRequestBuilder() *DescribeUserCrtListRequestBuilder {
	s := &DescribeUserCrtListRequest{}
	b := &DescribeUserCrtListRequestBuilder{s: s}
	return b
}

func (b *DescribeUserCrtListRequestBuilder) DescribeUserCrtListQuery(v *DescribeUserCrtListQuery) *DescribeUserCrtListRequestBuilder {
	b.s.DescribeUserCrtListQuery = v
	return b
}

func (b *DescribeUserCrtListRequestBuilder) Build() *DescribeUserCrtListRequest {
	return b.s
}


