// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdatePoolNameRequest struct {


	UpdatePoolNameBody *UpdatePoolNameBody `json:"updatePoolNameBody,omitempty"`
}

func (s UpdatePoolNameRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdatePoolNameRequest) GoString() string {
	return s.String()
}

func (s UpdatePoolNameRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePoolNameRequest) SetUpdatePoolNameBody(v *UpdatePoolNameBody) *UpdatePoolNameRequest {
	s.UpdatePoolNameBody = v
	return s
}


type UpdatePoolNameRequestBuilder struct {
	s *UpdatePoolNameRequest
}

func NewUpdatePoolNameRequestBuilder() *UpdatePoolNameRequestBuilder {
	s := &UpdatePoolNameRequest{}
	b := &UpdatePoolNameRequestBuilder{s: s}
	return b
}

func (b *UpdatePoolNameRequestBuilder) UpdatePoolNameBody(v *UpdatePoolNameBody) *UpdatePoolNameRequestBuilder {
	b.s.UpdatePoolNameBody = v
	return b
}

func (b *UpdatePoolNameRequestBuilder) Build() *UpdatePoolNameRequest {
	return b.s
}


