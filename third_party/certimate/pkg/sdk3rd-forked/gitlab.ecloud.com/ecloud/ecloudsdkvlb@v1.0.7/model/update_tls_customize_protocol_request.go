// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateTLSCustomizeProtocolRequest struct {


	UpdateTLSCustomizeProtocolBody *UpdateTLSCustomizeProtocolBody `json:"updateTLSCustomizeProtocolBody,omitempty"`
}

func (s UpdateTLSCustomizeProtocolRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSCustomizeProtocolRequest) GoString() string {
	return s.String()
}

func (s UpdateTLSCustomizeProtocolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSCustomizeProtocolRequest) SetUpdateTLSCustomizeProtocolBody(v *UpdateTLSCustomizeProtocolBody) *UpdateTLSCustomizeProtocolRequest {
	s.UpdateTLSCustomizeProtocolBody = v
	return s
}


type UpdateTLSCustomizeProtocolRequestBuilder struct {
	s *UpdateTLSCustomizeProtocolRequest
}

func NewUpdateTLSCustomizeProtocolRequestBuilder() *UpdateTLSCustomizeProtocolRequestBuilder {
	s := &UpdateTLSCustomizeProtocolRequest{}
	b := &UpdateTLSCustomizeProtocolRequestBuilder{s: s}
	return b
}

func (b *UpdateTLSCustomizeProtocolRequestBuilder) UpdateTLSCustomizeProtocolBody(v *UpdateTLSCustomizeProtocolBody) *UpdateTLSCustomizeProtocolRequestBuilder {
	b.s.UpdateTLSCustomizeProtocolBody = v
	return b
}

func (b *UpdateTLSCustomizeProtocolRequestBuilder) Build() *UpdateTLSCustomizeProtocolRequest {
	return b.s
}


