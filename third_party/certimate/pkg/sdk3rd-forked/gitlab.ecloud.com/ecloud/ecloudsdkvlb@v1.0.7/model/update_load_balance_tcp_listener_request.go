// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceTCPListenerRequest struct {


	UpdateLoadBalanceTCPListenerBody *UpdateLoadBalanceTCPListenerBody `json:"updateLoadBalanceTCPListenerBody,omitempty"`
}

func (s UpdateLoadBalanceTCPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceTCPListenerRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceTCPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceTCPListenerRequest) SetUpdateLoadBalanceTCPListenerBody(v *UpdateLoadBalanceTCPListenerBody) *UpdateLoadBalanceTCPListenerRequest {
	s.UpdateLoadBalanceTCPListenerBody = v
	return s
}


type UpdateLoadBalanceTCPListenerRequestBuilder struct {
	s *UpdateLoadBalanceTCPListenerRequest
}

func NewUpdateLoadBalanceTCPListenerRequestBuilder() *UpdateLoadBalanceTCPListenerRequestBuilder {
	s := &UpdateLoadBalanceTCPListenerRequest{}
	b := &UpdateLoadBalanceTCPListenerRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceTCPListenerRequestBuilder) UpdateLoadBalanceTCPListenerBody(v *UpdateLoadBalanceTCPListenerBody) *UpdateLoadBalanceTCPListenerRequestBuilder {
	b.s.UpdateLoadBalanceTCPListenerBody = v
	return b
}

func (b *UpdateLoadBalanceTCPListenerRequestBuilder) Build() *UpdateLoadBalanceTCPListenerRequest {
	return b.s
}


