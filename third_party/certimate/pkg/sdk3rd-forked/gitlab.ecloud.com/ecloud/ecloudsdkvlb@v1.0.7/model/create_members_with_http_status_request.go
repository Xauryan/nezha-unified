// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersWithHttpStatusRequest struct {


	CreateMembersWithHttpStatusBody *CreateMembersWithHttpStatusBody `json:"createMembersWithHttpStatusBody,omitempty"`

	CreateMembersWithHttpStatusPath *CreateMembersWithHttpStatusPath `json:"createMembersWithHttpStatusPath,omitempty"`
}

func (s CreateMembersWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s CreateMembersWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersWithHttpStatusRequest) SetCreateMembersWithHttpStatusBody(v *CreateMembersWithHttpStatusBody) *CreateMembersWithHttpStatusRequest {
	s.CreateMembersWithHttpStatusBody = v
	return s
}

func (s *CreateMembersWithHttpStatusRequest) SetCreateMembersWithHttpStatusPath(v *CreateMembersWithHttpStatusPath) *CreateMembersWithHttpStatusRequest {
	s.CreateMembersWithHttpStatusPath = v
	return s
}


type CreateMembersWithHttpStatusRequestBuilder struct {
	s *CreateMembersWithHttpStatusRequest
}

func NewCreateMembersWithHttpStatusRequestBuilder() *CreateMembersWithHttpStatusRequestBuilder {
	s := &CreateMembersWithHttpStatusRequest{}
	b := &CreateMembersWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *CreateMembersWithHttpStatusRequestBuilder) CreateMembersWithHttpStatusBody(v *CreateMembersWithHttpStatusBody) *CreateMembersWithHttpStatusRequestBuilder {
	b.s.CreateMembersWithHttpStatusBody = v
	return b
}

func (b *CreateMembersWithHttpStatusRequestBuilder) CreateMembersWithHttpStatusPath(v *CreateMembersWithHttpStatusPath) *CreateMembersWithHttpStatusRequestBuilder {
	b.s.CreateMembersWithHttpStatusPath = v
	return b
}

func (b *CreateMembersWithHttpStatusRequestBuilder) Build() *CreateMembersWithHttpStatusRequest {
	return b.s
}


