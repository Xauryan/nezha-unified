// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateTCAclRequest struct {


	CreateTCAclBody *CreateTCAclBody `json:"createTCAclBody,omitempty"`

	CreateTCAclHeader *CreateTCAclHeader `json:"createTCAclHeader,omitempty"`
}

func (s CreateTCAclRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateTCAclRequest) GoString() string {
	return s.String()
}

func (s CreateTCAclRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTCAclRequest) SetCreateTCAclBody(v *CreateTCAclBody) *CreateTCAclRequest {
	s.CreateTCAclBody = v
	return s
}

func (s *CreateTCAclRequest) SetCreateTCAclHeader(v *CreateTCAclHeader) *CreateTCAclRequest {
	s.CreateTCAclHeader = v
	return s
}


type CreateTCAclRequestBuilder struct {
	s *CreateTCAclRequest
}

func NewCreateTCAclRequestBuilder() *CreateTCAclRequestBuilder {
	s := &CreateTCAclRequest{}
	b := &CreateTCAclRequestBuilder{s: s}
	return b
}

func (b *CreateTCAclRequestBuilder) CreateTCAclBody(v *CreateTCAclBody) *CreateTCAclRequestBuilder {
	b.s.CreateTCAclBody = v
	return b
}

func (b *CreateTCAclRequestBuilder) CreateTCAclHeader(v *CreateTCAclHeader) *CreateTCAclRequestBuilder {
	b.s.CreateTCAclHeader = v
	return b
}

func (b *CreateTCAclRequestBuilder) Build() *CreateTCAclRequest {
	return b.s
}


