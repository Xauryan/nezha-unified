// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindFipRequest struct {


	BindFipBody *BindFipBody `json:"bindFipBody,omitempty"`
}

func (s BindFipRequest) String() string {
	return utils.Beautify(s)
}

func (s BindFipRequest) GoString() string {
	return s.String()
}

func (s BindFipRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindFipRequest) SetBindFipBody(v *BindFipBody) *BindFipRequest {
	s.BindFipBody = v
	return s
}


type BindFipRequestBuilder struct {
	s *BindFipRequest
}

func NewBindFipRequestBuilder() *BindFipRequestBuilder {
	s := &BindFipRequest{}
	b := &BindFipRequestBuilder{s: s}
	return b
}

func (b *BindFipRequestBuilder) BindFipBody(v *BindFipBody) *BindFipRequestBuilder {
	b.s.BindFipBody = v
	return b
}

func (b *BindFipRequestBuilder) Build() *BindFipRequest {
	return b.s
}


