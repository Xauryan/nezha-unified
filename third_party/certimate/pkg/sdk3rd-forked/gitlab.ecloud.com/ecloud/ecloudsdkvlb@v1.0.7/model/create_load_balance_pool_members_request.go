// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalancePoolMembersRequest struct {


	CreateLoadBalancePoolMembersBody *CreateLoadBalancePoolMembersBody `json:"createLoadBalancePoolMembersBody,omitempty"`

	CreateLoadBalancePoolMembersPath *CreateLoadBalancePoolMembersPath `json:"createLoadBalancePoolMembersPath,omitempty"`
}

func (s CreateLoadBalancePoolMembersRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalancePoolMembersRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalancePoolMembersRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalancePoolMembersRequest) SetCreateLoadBalancePoolMembersBody(v *CreateLoadBalancePoolMembersBody) *CreateLoadBalancePoolMembersRequest {
	s.CreateLoadBalancePoolMembersBody = v
	return s
}

func (s *CreateLoadBalancePoolMembersRequest) SetCreateLoadBalancePoolMembersPath(v *CreateLoadBalancePoolMembersPath) *CreateLoadBalancePoolMembersRequest {
	s.CreateLoadBalancePoolMembersPath = v
	return s
}


type CreateLoadBalancePoolMembersRequestBuilder struct {
	s *CreateLoadBalancePoolMembersRequest
}

func NewCreateLoadBalancePoolMembersRequestBuilder() *CreateLoadBalancePoolMembersRequestBuilder {
	s := &CreateLoadBalancePoolMembersRequest{}
	b := &CreateLoadBalancePoolMembersRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalancePoolMembersRequestBuilder) CreateLoadBalancePoolMembersBody(v *CreateLoadBalancePoolMembersBody) *CreateLoadBalancePoolMembersRequestBuilder {
	b.s.CreateLoadBalancePoolMembersBody = v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestBuilder) CreateLoadBalancePoolMembersPath(v *CreateLoadBalancePoolMembersPath) *CreateLoadBalancePoolMembersRequestBuilder {
	b.s.CreateLoadBalancePoolMembersPath = v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestBuilder) Build() *CreateLoadBalancePoolMembersRequest {
	return b.s
}


