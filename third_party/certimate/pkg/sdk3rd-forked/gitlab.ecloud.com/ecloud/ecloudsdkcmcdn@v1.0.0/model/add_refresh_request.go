// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddRefreshRequest struct {


	AddRefreshBody *AddRefreshBody `json:"addRefreshBody,omitempty"`
}

func (s AddRefreshRequest) String() string {
	return utils.Beautify(s)
}

func (s AddRefreshRequest) GoString() string {
	return s.String()
}

func (s AddRefreshRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddRefreshRequest) SetAddRefreshBody(v *AddRefreshBody) *AddRefreshRequest {
	s.AddRefreshBody = v
	return s
}


type AddRefreshRequestBuilder struct {
	s *AddRefreshRequest
}

func NewAddRefreshRequestBuilder() *AddRefreshRequestBuilder {
	s := &AddRefreshRequest{}
	b := &AddRefreshRequestBuilder{s: s}
	return b
}

func (b *AddRefreshRequestBuilder) AddRefreshBody(v *AddRefreshBody) *AddRefreshRequestBuilder {
	b.s.AddRefreshBody = v
	return b
}

func (b *AddRefreshRequestBuilder) Build() *AddRefreshRequest {
	return b.s
}


