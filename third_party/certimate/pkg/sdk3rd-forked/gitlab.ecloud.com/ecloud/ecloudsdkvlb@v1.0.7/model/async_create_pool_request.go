// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncCreatePoolRequest struct {


	AsyncCreatePoolBody *AsyncCreatePoolBody `json:"asyncCreatePoolBody,omitempty"`
}

func (s AsyncCreatePoolRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePoolRequest) GoString() string {
	return s.String()
}

func (s AsyncCreatePoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePoolRequest) SetAsyncCreatePoolBody(v *AsyncCreatePoolBody) *AsyncCreatePoolRequest {
	s.AsyncCreatePoolBody = v
	return s
}


type AsyncCreatePoolRequestBuilder struct {
	s *AsyncCreatePoolRequest
}

func NewAsyncCreatePoolRequestBuilder() *AsyncCreatePoolRequestBuilder {
	s := &AsyncCreatePoolRequest{}
	b := &AsyncCreatePoolRequestBuilder{s: s}
	return b
}

func (b *AsyncCreatePoolRequestBuilder) AsyncCreatePoolBody(v *AsyncCreatePoolBody) *AsyncCreatePoolRequestBuilder {
	b.s.AsyncCreatePoolBody = v
	return b
}

func (b *AsyncCreatePoolRequestBuilder) Build() *AsyncCreatePoolRequest {
	return b.s
}


