// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnbindFipRequest struct {


	UnbindFipPath *UnbindFipPath `json:"unbindFipPath,omitempty"`

	UnbindFipQuery *UnbindFipQuery `json:"unbindFipQuery,omitempty"`
}

func (s UnbindFipRequest) String() string {
	return utils.Beautify(s)
}

func (s UnbindFipRequest) GoString() string {
	return s.String()
}

func (s UnbindFipRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindFipRequest) SetUnbindFipPath(v *UnbindFipPath) *UnbindFipRequest {
	s.UnbindFipPath = v
	return s
}

func (s *UnbindFipRequest) SetUnbindFipQuery(v *UnbindFipQuery) *UnbindFipRequest {
	s.UnbindFipQuery = v
	return s
}


type UnbindFipRequestBuilder struct {
	s *UnbindFipRequest
}

func NewUnbindFipRequestBuilder() *UnbindFipRequestBuilder {
	s := &UnbindFipRequest{}
	b := &UnbindFipRequestBuilder{s: s}
	return b
}

func (b *UnbindFipRequestBuilder) UnbindFipPath(v *UnbindFipPath) *UnbindFipRequestBuilder {
	b.s.UnbindFipPath = v
	return b
}

func (b *UnbindFipRequestBuilder) UnbindFipQuery(v *UnbindFipQuery) *UnbindFipRequestBuilder {
	b.s.UnbindFipQuery = v
	return b
}

func (b *UnbindFipRequestBuilder) Build() *UnbindFipRequest {
	return b.s
}


