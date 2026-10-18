// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateL7PolicyRequest struct {


	AsyncUpdateL7PolicyBody *AsyncUpdateL7PolicyBody `json:"asyncUpdateL7PolicyBody,omitempty"`
}

func (s AsyncUpdateL7PolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateL7PolicyRequest) GoString() string {
	return s.String()
}

func (s AsyncUpdateL7PolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateL7PolicyRequest) SetAsyncUpdateL7PolicyBody(v *AsyncUpdateL7PolicyBody) *AsyncUpdateL7PolicyRequest {
	s.AsyncUpdateL7PolicyBody = v
	return s
}


type AsyncUpdateL7PolicyRequestBuilder struct {
	s *AsyncUpdateL7PolicyRequest
}

func NewAsyncUpdateL7PolicyRequestBuilder() *AsyncUpdateL7PolicyRequestBuilder {
	s := &AsyncUpdateL7PolicyRequest{}
	b := &AsyncUpdateL7PolicyRequestBuilder{s: s}
	return b
}

func (b *AsyncUpdateL7PolicyRequestBuilder) AsyncUpdateL7PolicyBody(v *AsyncUpdateL7PolicyBody) *AsyncUpdateL7PolicyRequestBuilder {
	b.s.AsyncUpdateL7PolicyBody = v
	return b
}

func (b *AsyncUpdateL7PolicyRequestBuilder) Build() *AsyncUpdateL7PolicyRequest {
	return b.s
}


