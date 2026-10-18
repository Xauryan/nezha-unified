// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserDomainsRequest struct {


	DescribeUserDomainsQuery *DescribeUserDomainsQuery `json:"describeUserDomainsQuery,omitempty"`
}

func (s DescribeUserDomainsRequest) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserDomainsRequest) GoString() string {
	return s.String()
}

func (s DescribeUserDomainsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserDomainsRequest) SetDescribeUserDomainsQuery(v *DescribeUserDomainsQuery) *DescribeUserDomainsRequest {
	s.DescribeUserDomainsQuery = v
	return s
}


type DescribeUserDomainsRequestBuilder struct {
	s *DescribeUserDomainsRequest
}

func NewDescribeUserDomainsRequestBuilder() *DescribeUserDomainsRequestBuilder {
	s := &DescribeUserDomainsRequest{}
	b := &DescribeUserDomainsRequestBuilder{s: s}
	return b
}

func (b *DescribeUserDomainsRequestBuilder) DescribeUserDomainsQuery(v *DescribeUserDomainsQuery) *DescribeUserDomainsRequestBuilder {
	b.s.DescribeUserDomainsQuery = v
	return b
}

func (b *DescribeUserDomainsRequestBuilder) Build() *DescribeUserDomainsRequest {
	return b.s
}


