// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceUDPListenerPoolRequest struct {


	UpdateLoadBalanceUDPListenerPoolBody *UpdateLoadBalanceUDPListenerPoolBody `json:"updateLoadBalanceUDPListenerPoolBody,omitempty"`
}

func (s UpdateLoadBalanceUDPListenerPoolRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerPoolRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerPoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerPoolRequest) SetUpdateLoadBalanceUDPListenerPoolBody(v *UpdateLoadBalanceUDPListenerPoolBody) *UpdateLoadBalanceUDPListenerPoolRequest {
	s.UpdateLoadBalanceUDPListenerPoolBody = v
	return s
}


type UpdateLoadBalanceUDPListenerPoolRequestBuilder struct {
	s *UpdateLoadBalanceUDPListenerPoolRequest
}

func NewUpdateLoadBalanceUDPListenerPoolRequestBuilder() *UpdateLoadBalanceUDPListenerPoolRequestBuilder {
	s := &UpdateLoadBalanceUDPListenerPoolRequest{}
	b := &UpdateLoadBalanceUDPListenerPoolRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolRequestBuilder) UpdateLoadBalanceUDPListenerPoolBody(v *UpdateLoadBalanceUDPListenerPoolBody) *UpdateLoadBalanceUDPListenerPoolRequestBuilder {
	b.s.UpdateLoadBalanceUDPListenerPoolBody = v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolRequestBuilder) Build() *UpdateLoadBalanceUDPListenerPoolRequest {
	return b.s
}


