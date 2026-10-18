// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateControlGroupRequest struct {


	CreateControlGroupBody *CreateControlGroupBody `json:"createControlGroupBody,omitempty"`
}

func (s CreateControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupRequest) GoString() string {
	return s.String()
}

func (s CreateControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupRequest) SetCreateControlGroupBody(v *CreateControlGroupBody) *CreateControlGroupRequest {
	s.CreateControlGroupBody = v
	return s
}


type CreateControlGroupRequestBuilder struct {
	s *CreateControlGroupRequest
}

func NewCreateControlGroupRequestBuilder() *CreateControlGroupRequestBuilder {
	s := &CreateControlGroupRequest{}
	b := &CreateControlGroupRequestBuilder{s: s}
	return b
}

func (b *CreateControlGroupRequestBuilder) CreateControlGroupBody(v *CreateControlGroupBody) *CreateControlGroupRequestBuilder {
	b.s.CreateControlGroupBody = v
	return b
}

func (b *CreateControlGroupRequestBuilder) Build() *CreateControlGroupRequest {
	return b.s
}


