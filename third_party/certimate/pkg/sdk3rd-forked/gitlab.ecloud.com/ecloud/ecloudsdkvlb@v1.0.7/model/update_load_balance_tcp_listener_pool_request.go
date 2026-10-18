// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceTCPListenerPoolRequest struct {


	UpdateLoadBalanceTCPListenerPoolBody *UpdateLoadBalanceTCPListenerPoolBody `json:"updateLoadBalanceTCPListenerPoolBody,omitempty"`
}

func (s UpdateLoadBalanceTCPListenerPoolRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceTCPListenerPoolRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceTCPListenerPoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceTCPListenerPoolRequest) SetUpdateLoadBalanceTCPListenerPoolBody(v *UpdateLoadBalanceTCPListenerPoolBody) *UpdateLoadBalanceTCPListenerPoolRequest {
	s.UpdateLoadBalanceTCPListenerPoolBody = v
	return s
}


type UpdateLoadBalanceTCPListenerPoolRequestBuilder struct {
	s *UpdateLoadBalanceTCPListenerPoolRequest
}

func NewUpdateLoadBalanceTCPListenerPoolRequestBuilder() *UpdateLoadBalanceTCPListenerPoolRequestBuilder {
	s := &UpdateLoadBalanceTCPListenerPoolRequest{}
	b := &UpdateLoadBalanceTCPListenerPoolRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolRequestBuilder) UpdateLoadBalanceTCPListenerPoolBody(v *UpdateLoadBalanceTCPListenerPoolBody) *UpdateLoadBalanceTCPListenerPoolRequestBuilder {
	b.s.UpdateLoadBalanceTCPListenerPoolBody = v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolRequestBuilder) Build() *UpdateLoadBalanceTCPListenerPoolRequest {
	return b.s
}


