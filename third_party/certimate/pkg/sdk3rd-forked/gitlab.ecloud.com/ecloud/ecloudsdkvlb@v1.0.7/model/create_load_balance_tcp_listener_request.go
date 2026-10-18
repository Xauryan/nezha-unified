// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalanceTCPListenerRequest struct {


	CreateLoadBalanceTCPListenerBody *CreateLoadBalanceTCPListenerBody `json:"createLoadBalanceTCPListenerBody,omitempty"`
}

func (s CreateLoadBalanceTCPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceTCPListenerRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceTCPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceTCPListenerRequest) SetCreateLoadBalanceTCPListenerBody(v *CreateLoadBalanceTCPListenerBody) *CreateLoadBalanceTCPListenerRequest {
	s.CreateLoadBalanceTCPListenerBody = v
	return s
}


type CreateLoadBalanceTCPListenerRequestBuilder struct {
	s *CreateLoadBalanceTCPListenerRequest
}

func NewCreateLoadBalanceTCPListenerRequestBuilder() *CreateLoadBalanceTCPListenerRequestBuilder {
	s := &CreateLoadBalanceTCPListenerRequest{}
	b := &CreateLoadBalanceTCPListenerRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceTCPListenerRequestBuilder) CreateLoadBalanceTCPListenerBody(v *CreateLoadBalanceTCPListenerBody) *CreateLoadBalanceTCPListenerRequestBuilder {
	b.s.CreateLoadBalanceTCPListenerBody = v
	return b
}

func (b *CreateLoadBalanceTCPListenerRequestBuilder) Build() *CreateLoadBalanceTCPListenerRequest {
	return b.s
}


