// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateTLSNameRequest struct {


	UpdateTLSNameBody *UpdateTLSNameBody `json:"updateTLSNameBody,omitempty"`
}

func (s UpdateTLSNameRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSNameRequest) GoString() string {
	return s.String()
}

func (s UpdateTLSNameRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSNameRequest) SetUpdateTLSNameBody(v *UpdateTLSNameBody) *UpdateTLSNameRequest {
	s.UpdateTLSNameBody = v
	return s
}


type UpdateTLSNameRequestBuilder struct {
	s *UpdateTLSNameRequest
}

func NewUpdateTLSNameRequestBuilder() *UpdateTLSNameRequestBuilder {
	s := &UpdateTLSNameRequest{}
	b := &UpdateTLSNameRequestBuilder{s: s}
	return b
}

func (b *UpdateTLSNameRequestBuilder) UpdateTLSNameBody(v *UpdateTLSNameBody) *UpdateTLSNameRequestBuilder {
	b.s.UpdateTLSNameBody = v
	return b
}

func (b *UpdateTLSNameRequestBuilder) Build() *UpdateTLSNameRequest {
	return b.s
}


