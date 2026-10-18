// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAliAclRequest struct {


	CreateAliAclHeader *CreateAliAclHeader `json:"createAliAclHeader,omitempty"`

	CreateAliAclBody *CreateAliAclBody `json:"createAliAclBody,omitempty"`
}

func (s CreateAliAclRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateAliAclRequest) GoString() string {
	return s.String()
}

func (s CreateAliAclRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAliAclRequest) SetCreateAliAclHeader(v *CreateAliAclHeader) *CreateAliAclRequest {
	s.CreateAliAclHeader = v
	return s
}

func (s *CreateAliAclRequest) SetCreateAliAclBody(v *CreateAliAclBody) *CreateAliAclRequest {
	s.CreateAliAclBody = v
	return s
}


type CreateAliAclRequestBuilder struct {
	s *CreateAliAclRequest
}

func NewCreateAliAclRequestBuilder() *CreateAliAclRequestBuilder {
	s := &CreateAliAclRequest{}
	b := &CreateAliAclRequestBuilder{s: s}
	return b
}

func (b *CreateAliAclRequestBuilder) CreateAliAclHeader(v *CreateAliAclHeader) *CreateAliAclRequestBuilder {
	b.s.CreateAliAclHeader = v
	return b
}

func (b *CreateAliAclRequestBuilder) CreateAliAclBody(v *CreateAliAclBody) *CreateAliAclRequestBuilder {
	b.s.CreateAliAclBody = v
	return b
}

func (b *CreateAliAclRequestBuilder) Build() *CreateAliAclRequest {
	return b.s
}


