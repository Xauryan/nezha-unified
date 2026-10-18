// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalanceUDPListenerRequest struct {


	CreateLoadBalanceUDPListenerBody *CreateLoadBalanceUDPListenerBody `json:"createLoadBalanceUDPListenerBody,omitempty"`
}

func (s CreateLoadBalanceUDPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceUDPListenerRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceUDPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceUDPListenerRequest) SetCreateLoadBalanceUDPListenerBody(v *CreateLoadBalanceUDPListenerBody) *CreateLoadBalanceUDPListenerRequest {
	s.CreateLoadBalanceUDPListenerBody = v
	return s
}


type CreateLoadBalanceUDPListenerRequestBuilder struct {
	s *CreateLoadBalanceUDPListenerRequest
}

func NewCreateLoadBalanceUDPListenerRequestBuilder() *CreateLoadBalanceUDPListenerRequestBuilder {
	s := &CreateLoadBalanceUDPListenerRequest{}
	b := &CreateLoadBalanceUDPListenerRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceUDPListenerRequestBuilder) CreateLoadBalanceUDPListenerBody(v *CreateLoadBalanceUDPListenerBody) *CreateLoadBalanceUDPListenerRequestBuilder {
	b.s.CreateLoadBalanceUDPListenerBody = v
	return b
}

func (b *CreateLoadBalanceUDPListenerRequestBuilder) Build() *CreateLoadBalanceUDPListenerRequest {
	return b.s
}


