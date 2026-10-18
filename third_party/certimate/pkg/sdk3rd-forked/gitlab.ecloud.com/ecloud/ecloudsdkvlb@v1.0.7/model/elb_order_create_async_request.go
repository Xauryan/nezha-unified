// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncRequest struct {


	ElbOrderCreateAsyncBody *ElbOrderCreateAsyncBody `json:"elbOrderCreateAsyncBody,omitempty"`
}

func (s ElbOrderCreateAsyncRequest) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncRequest) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncRequest) SetElbOrderCreateAsyncBody(v *ElbOrderCreateAsyncBody) *ElbOrderCreateAsyncRequest {
	s.ElbOrderCreateAsyncBody = v
	return s
}


type ElbOrderCreateAsyncRequestBuilder struct {
	s *ElbOrderCreateAsyncRequest
}

func NewElbOrderCreateAsyncRequestBuilder() *ElbOrderCreateAsyncRequestBuilder {
	s := &ElbOrderCreateAsyncRequest{}
	b := &ElbOrderCreateAsyncRequestBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncRequestBuilder) ElbOrderCreateAsyncBody(v *ElbOrderCreateAsyncBody) *ElbOrderCreateAsyncRequestBuilder {
	b.s.ElbOrderCreateAsyncBody = v
	return b
}

func (b *ElbOrderCreateAsyncRequestBuilder) Build() *ElbOrderCreateAsyncRequest {
	return b.s
}


