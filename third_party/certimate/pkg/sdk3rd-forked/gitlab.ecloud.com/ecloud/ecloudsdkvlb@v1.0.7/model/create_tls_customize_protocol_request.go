// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateTLSCustomizeProtocolRequest struct {


	CreateTLSCustomizeProtocolBody *CreateTLSCustomizeProtocolBody `json:"createTLSCustomizeProtocolBody,omitempty"`
}

func (s CreateTLSCustomizeProtocolRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateTLSCustomizeProtocolRequest) GoString() string {
	return s.String()
}

func (s CreateTLSCustomizeProtocolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTLSCustomizeProtocolRequest) SetCreateTLSCustomizeProtocolBody(v *CreateTLSCustomizeProtocolBody) *CreateTLSCustomizeProtocolRequest {
	s.CreateTLSCustomizeProtocolBody = v
	return s
}


type CreateTLSCustomizeProtocolRequestBuilder struct {
	s *CreateTLSCustomizeProtocolRequest
}

func NewCreateTLSCustomizeProtocolRequestBuilder() *CreateTLSCustomizeProtocolRequestBuilder {
	s := &CreateTLSCustomizeProtocolRequest{}
	b := &CreateTLSCustomizeProtocolRequestBuilder{s: s}
	return b
}

func (b *CreateTLSCustomizeProtocolRequestBuilder) CreateTLSCustomizeProtocolBody(v *CreateTLSCustomizeProtocolBody) *CreateTLSCustomizeProtocolRequestBuilder {
	b.s.CreateTLSCustomizeProtocolBody = v
	return b
}

func (b *CreateTLSCustomizeProtocolRequestBuilder) Build() *CreateTLSCustomizeProtocolRequest {
	return b.s
}


