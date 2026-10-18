// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersRequest struct {


	CreateMembersPath *CreateMembersPath `json:"createMembersPath,omitempty"`

	CreateMembersBody *CreateMembersBody `json:"createMembersBody,omitempty"`
}

func (s CreateMembersRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersRequest) GoString() string {
	return s.String()
}

func (s CreateMembersRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersRequest) SetCreateMembersPath(v *CreateMembersPath) *CreateMembersRequest {
	s.CreateMembersPath = v
	return s
}

func (s *CreateMembersRequest) SetCreateMembersBody(v *CreateMembersBody) *CreateMembersRequest {
	s.CreateMembersBody = v
	return s
}


type CreateMembersRequestBuilder struct {
	s *CreateMembersRequest
}

func NewCreateMembersRequestBuilder() *CreateMembersRequestBuilder {
	s := &CreateMembersRequest{}
	b := &CreateMembersRequestBuilder{s: s}
	return b
}

func (b *CreateMembersRequestBuilder) CreateMembersPath(v *CreateMembersPath) *CreateMembersRequestBuilder {
	b.s.CreateMembersPath = v
	return b
}

func (b *CreateMembersRequestBuilder) CreateMembersBody(v *CreateMembersBody) *CreateMembersRequestBuilder {
	b.s.CreateMembersBody = v
	return b
}

func (b *CreateMembersRequestBuilder) Build() *CreateMembersRequest {
	return b.s
}


