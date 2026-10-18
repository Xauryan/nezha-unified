// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalanceHTTPSListenerRequest struct {


	CreateLoadBalanceHTTPSListenerBody *CreateLoadBalanceHTTPSListenerBody `json:"createLoadBalanceHTTPSListenerBody,omitempty"`
}

func (s CreateLoadBalanceHTTPSListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPSListenerRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPSListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPSListenerRequest) SetCreateLoadBalanceHTTPSListenerBody(v *CreateLoadBalanceHTTPSListenerBody) *CreateLoadBalanceHTTPSListenerRequest {
	s.CreateLoadBalanceHTTPSListenerBody = v
	return s
}


type CreateLoadBalanceHTTPSListenerRequestBuilder struct {
	s *CreateLoadBalanceHTTPSListenerRequest
}

func NewCreateLoadBalanceHTTPSListenerRequestBuilder() *CreateLoadBalanceHTTPSListenerRequestBuilder {
	s := &CreateLoadBalanceHTTPSListenerRequest{}
	b := &CreateLoadBalanceHTTPSListenerRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPSListenerRequestBuilder) CreateLoadBalanceHTTPSListenerBody(v *CreateLoadBalanceHTTPSListenerBody) *CreateLoadBalanceHTTPSListenerRequestBuilder {
	b.s.CreateLoadBalanceHTTPSListenerBody = v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerRequestBuilder) Build() *CreateLoadBalanceHTTPSListenerRequest {
	return b.s
}


