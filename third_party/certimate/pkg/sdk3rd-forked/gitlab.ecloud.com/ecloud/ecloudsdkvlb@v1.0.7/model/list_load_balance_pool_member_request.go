// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalancePoolMemberRequest struct {


	ListLoadBalancePoolMemberPath *ListLoadBalancePoolMemberPath `json:"listLoadBalancePoolMemberPath,omitempty"`

	ListLoadBalancePoolMemberQuery *ListLoadBalancePoolMemberQuery `json:"listLoadBalancePoolMemberQuery,omitempty"`
}

func (s ListLoadBalancePoolMemberRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberRequest) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberRequest) SetListLoadBalancePoolMemberPath(v *ListLoadBalancePoolMemberPath) *ListLoadBalancePoolMemberRequest {
	s.ListLoadBalancePoolMemberPath = v
	return s
}

func (s *ListLoadBalancePoolMemberRequest) SetListLoadBalancePoolMemberQuery(v *ListLoadBalancePoolMemberQuery) *ListLoadBalancePoolMemberRequest {
	s.ListLoadBalancePoolMemberQuery = v
	return s
}


type ListLoadBalancePoolMemberRequestBuilder struct {
	s *ListLoadBalancePoolMemberRequest
}

func NewListLoadBalancePoolMemberRequestBuilder() *ListLoadBalancePoolMemberRequestBuilder {
	s := &ListLoadBalancePoolMemberRequest{}
	b := &ListLoadBalancePoolMemberRequestBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberRequestBuilder) ListLoadBalancePoolMemberPath(v *ListLoadBalancePoolMemberPath) *ListLoadBalancePoolMemberRequestBuilder {
	b.s.ListLoadBalancePoolMemberPath = v
	return b
}

func (b *ListLoadBalancePoolMemberRequestBuilder) ListLoadBalancePoolMemberQuery(v *ListLoadBalancePoolMemberQuery) *ListLoadBalancePoolMemberRequestBuilder {
	b.s.ListLoadBalancePoolMemberQuery = v
	return b
}

func (b *ListLoadBalancePoolMemberRequestBuilder) Build() *ListLoadBalancePoolMemberRequest {
	return b.s
}


