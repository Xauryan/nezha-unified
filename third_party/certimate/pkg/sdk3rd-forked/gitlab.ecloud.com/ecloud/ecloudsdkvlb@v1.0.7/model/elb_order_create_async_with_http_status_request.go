// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncWithHttpStatusRequest struct {


	ElbOrderCreateAsyncWithHttpStatusBody *ElbOrderCreateAsyncWithHttpStatusBody `json:"elbOrderCreateAsyncWithHttpStatusBody,omitempty"`
}

func (s ElbOrderCreateAsyncWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncWithHttpStatusRequest) SetElbOrderCreateAsyncWithHttpStatusBody(v *ElbOrderCreateAsyncWithHttpStatusBody) *ElbOrderCreateAsyncWithHttpStatusRequest {
	s.ElbOrderCreateAsyncWithHttpStatusBody = v
	return s
}


type ElbOrderCreateAsyncWithHttpStatusRequestBuilder struct {
	s *ElbOrderCreateAsyncWithHttpStatusRequest
}

func NewElbOrderCreateAsyncWithHttpStatusRequestBuilder() *ElbOrderCreateAsyncWithHttpStatusRequestBuilder {
	s := &ElbOrderCreateAsyncWithHttpStatusRequest{}
	b := &ElbOrderCreateAsyncWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusRequestBuilder) ElbOrderCreateAsyncWithHttpStatusBody(v *ElbOrderCreateAsyncWithHttpStatusBody) *ElbOrderCreateAsyncWithHttpStatusRequestBuilder {
	b.s.ElbOrderCreateAsyncWithHttpStatusBody = v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusRequestBuilder) Build() *ElbOrderCreateAsyncWithHttpStatusRequest {
	return b.s
}


