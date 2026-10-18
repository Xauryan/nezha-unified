// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateRequest struct {


	AsyncUpdateBody *AsyncUpdateBody `json:"asyncUpdateBody,omitempty"`
}

func (s AsyncUpdateRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateRequest) GoString() string {
	return s.String()
}

func (s AsyncUpdateRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateRequest) SetAsyncUpdateBody(v *AsyncUpdateBody) *AsyncUpdateRequest {
	s.AsyncUpdateBody = v
	return s
}


type AsyncUpdateRequestBuilder struct {
	s *AsyncUpdateRequest
}

func NewAsyncUpdateRequestBuilder() *AsyncUpdateRequestBuilder {
	s := &AsyncUpdateRequest{}
	b := &AsyncUpdateRequestBuilder{s: s}
	return b
}

func (b *AsyncUpdateRequestBuilder) AsyncUpdateBody(v *AsyncUpdateBody) *AsyncUpdateRequestBuilder {
	b.s.AsyncUpdateBody = v
	return b
}

func (b *AsyncUpdateRequestBuilder) Build() *AsyncUpdateRequest {
	return b.s
}


