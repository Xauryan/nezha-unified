// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateWithHttpStatusRequest struct {


	AsyncUpdateWithHttpStatusBody *AsyncUpdateWithHttpStatusBody `json:"asyncUpdateWithHttpStatusBody,omitempty"`
}

func (s AsyncUpdateWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s AsyncUpdateWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateWithHttpStatusRequest) SetAsyncUpdateWithHttpStatusBody(v *AsyncUpdateWithHttpStatusBody) *AsyncUpdateWithHttpStatusRequest {
	s.AsyncUpdateWithHttpStatusBody = v
	return s
}


type AsyncUpdateWithHttpStatusRequestBuilder struct {
	s *AsyncUpdateWithHttpStatusRequest
}

func NewAsyncUpdateWithHttpStatusRequestBuilder() *AsyncUpdateWithHttpStatusRequestBuilder {
	s := &AsyncUpdateWithHttpStatusRequest{}
	b := &AsyncUpdateWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *AsyncUpdateWithHttpStatusRequestBuilder) AsyncUpdateWithHttpStatusBody(v *AsyncUpdateWithHttpStatusBody) *AsyncUpdateWithHttpStatusRequestBuilder {
	b.s.AsyncUpdateWithHttpStatusBody = v
	return b
}

func (b *AsyncUpdateWithHttpStatusRequestBuilder) Build() *AsyncUpdateWithHttpStatusRequest {
	return b.s
}


