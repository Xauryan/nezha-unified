// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalancePoolMemberRequest struct {


	UpdateLoadBalancePoolMemberBody *UpdateLoadBalancePoolMemberBody `json:"updateLoadBalancePoolMemberBody,omitempty"`

	UpdateLoadBalancePoolMemberPath *UpdateLoadBalancePoolMemberPath `json:"updateLoadBalancePoolMemberPath,omitempty"`
}

func (s UpdateLoadBalancePoolMemberRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancePoolMemberRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancePoolMemberRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancePoolMemberRequest) SetUpdateLoadBalancePoolMemberBody(v *UpdateLoadBalancePoolMemberBody) *UpdateLoadBalancePoolMemberRequest {
	s.UpdateLoadBalancePoolMemberBody = v
	return s
}

func (s *UpdateLoadBalancePoolMemberRequest) SetUpdateLoadBalancePoolMemberPath(v *UpdateLoadBalancePoolMemberPath) *UpdateLoadBalancePoolMemberRequest {
	s.UpdateLoadBalancePoolMemberPath = v
	return s
}


type UpdateLoadBalancePoolMemberRequestBuilder struct {
	s *UpdateLoadBalancePoolMemberRequest
}

func NewUpdateLoadBalancePoolMemberRequestBuilder() *UpdateLoadBalancePoolMemberRequestBuilder {
	s := &UpdateLoadBalancePoolMemberRequest{}
	b := &UpdateLoadBalancePoolMemberRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancePoolMemberRequestBuilder) UpdateLoadBalancePoolMemberBody(v *UpdateLoadBalancePoolMemberBody) *UpdateLoadBalancePoolMemberRequestBuilder {
	b.s.UpdateLoadBalancePoolMemberBody = v
	return b
}

func (b *UpdateLoadBalancePoolMemberRequestBuilder) UpdateLoadBalancePoolMemberPath(v *UpdateLoadBalancePoolMemberPath) *UpdateLoadBalancePoolMemberRequestBuilder {
	b.s.UpdateLoadBalancePoolMemberPath = v
	return b
}

func (b *UpdateLoadBalancePoolMemberRequestBuilder) Build() *UpdateLoadBalancePoolMemberRequest {
	return b.s
}


