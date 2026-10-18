// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncCreateWithHttpStatusRequest struct {


	AsyncCreateWithHttpStatusBody *AsyncCreateWithHttpStatusBody `json:"asyncCreateWithHttpStatusBody,omitempty"`
}

func (s AsyncCreateWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreateWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s AsyncCreateWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreateWithHttpStatusRequest) SetAsyncCreateWithHttpStatusBody(v *AsyncCreateWithHttpStatusBody) *AsyncCreateWithHttpStatusRequest {
	s.AsyncCreateWithHttpStatusBody = v
	return s
}


type AsyncCreateWithHttpStatusRequestBuilder struct {
	s *AsyncCreateWithHttpStatusRequest
}

func NewAsyncCreateWithHttpStatusRequestBuilder() *AsyncCreateWithHttpStatusRequestBuilder {
	s := &AsyncCreateWithHttpStatusRequest{}
	b := &AsyncCreateWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *AsyncCreateWithHttpStatusRequestBuilder) AsyncCreateWithHttpStatusBody(v *AsyncCreateWithHttpStatusBody) *AsyncCreateWithHttpStatusRequestBuilder {
	b.s.AsyncCreateWithHttpStatusBody = v
	return b
}

func (b *AsyncCreateWithHttpStatusRequestBuilder) Build() *AsyncCreateWithHttpStatusRequest {
	return b.s
}


