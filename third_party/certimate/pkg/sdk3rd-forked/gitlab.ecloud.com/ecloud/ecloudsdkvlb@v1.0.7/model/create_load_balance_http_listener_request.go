// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalanceHTTPListenerRequest struct {


	CreateLoadBalanceHTTPListenerBody *CreateLoadBalanceHTTPListenerBody `json:"createLoadBalanceHTTPListenerBody,omitempty"`
}

func (s CreateLoadBalanceHTTPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPListenerRequest) SetCreateLoadBalanceHTTPListenerBody(v *CreateLoadBalanceHTTPListenerBody) *CreateLoadBalanceHTTPListenerRequest {
	s.CreateLoadBalanceHTTPListenerBody = v
	return s
}


type CreateLoadBalanceHTTPListenerRequestBuilder struct {
	s *CreateLoadBalanceHTTPListenerRequest
}

func NewCreateLoadBalanceHTTPListenerRequestBuilder() *CreateLoadBalanceHTTPListenerRequestBuilder {
	s := &CreateLoadBalanceHTTPListenerRequest{}
	b := &CreateLoadBalanceHTTPListenerRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPListenerRequestBuilder) CreateLoadBalanceHTTPListenerBody(v *CreateLoadBalanceHTTPListenerBody) *CreateLoadBalanceHTTPListenerRequestBuilder {
	b.s.CreateLoadBalanceHTTPListenerBody = v
	return b
}

func (b *CreateLoadBalanceHTTPListenerRequestBuilder) Build() *CreateLoadBalanceHTTPListenerRequest {
	return b.s
}


