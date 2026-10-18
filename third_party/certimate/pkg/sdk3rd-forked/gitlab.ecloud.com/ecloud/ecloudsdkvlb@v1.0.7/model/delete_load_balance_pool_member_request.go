// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteLoadBalancePoolMemberRequest struct {


	DeleteLoadBalancePoolMemberPath *DeleteLoadBalancePoolMemberPath `json:"deleteLoadBalancePoolMemberPath,omitempty"`

	DeleteLoadBalancePoolMemberHeader *DeleteLoadBalancePoolMemberHeader `json:"deleteLoadBalancePoolMemberHeader,omitempty"`
}

func (s DeleteLoadBalancePoolMemberRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteLoadBalancePoolMemberRequest) GoString() string {
	return s.String()
}

func (s DeleteLoadBalancePoolMemberRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteLoadBalancePoolMemberRequest) SetDeleteLoadBalancePoolMemberPath(v *DeleteLoadBalancePoolMemberPath) *DeleteLoadBalancePoolMemberRequest {
	s.DeleteLoadBalancePoolMemberPath = v
	return s
}

func (s *DeleteLoadBalancePoolMemberRequest) SetDeleteLoadBalancePoolMemberHeader(v *DeleteLoadBalancePoolMemberHeader) *DeleteLoadBalancePoolMemberRequest {
	s.DeleteLoadBalancePoolMemberHeader = v
	return s
}


type DeleteLoadBalancePoolMemberRequestBuilder struct {
	s *DeleteLoadBalancePoolMemberRequest
}

func NewDeleteLoadBalancePoolMemberRequestBuilder() *DeleteLoadBalancePoolMemberRequestBuilder {
	s := &DeleteLoadBalancePoolMemberRequest{}
	b := &DeleteLoadBalancePoolMemberRequestBuilder{s: s}
	return b
}

func (b *DeleteLoadBalancePoolMemberRequestBuilder) DeleteLoadBalancePoolMemberPath(v *DeleteLoadBalancePoolMemberPath) *DeleteLoadBalancePoolMemberRequestBuilder {
	b.s.DeleteLoadBalancePoolMemberPath = v
	return b
}

func (b *DeleteLoadBalancePoolMemberRequestBuilder) DeleteLoadBalancePoolMemberHeader(v *DeleteLoadBalancePoolMemberHeader) *DeleteLoadBalancePoolMemberRequestBuilder {
	b.s.DeleteLoadBalancePoolMemberHeader = v
	return b
}

func (b *DeleteLoadBalancePoolMemberRequestBuilder) Build() *DeleteLoadBalancePoolMemberRequest {
	return b.s
}


