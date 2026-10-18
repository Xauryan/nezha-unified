// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceUDPListenerRequest struct {


	UpdateLoadBalanceUDPListenerBody *UpdateLoadBalanceUDPListenerBody `json:"updateLoadBalanceUDPListenerBody,omitempty"`
}

func (s UpdateLoadBalanceUDPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerRequest) SetUpdateLoadBalanceUDPListenerBody(v *UpdateLoadBalanceUDPListenerBody) *UpdateLoadBalanceUDPListenerRequest {
	s.UpdateLoadBalanceUDPListenerBody = v
	return s
}


type UpdateLoadBalanceUDPListenerRequestBuilder struct {
	s *UpdateLoadBalanceUDPListenerRequest
}

func NewUpdateLoadBalanceUDPListenerRequestBuilder() *UpdateLoadBalanceUDPListenerRequestBuilder {
	s := &UpdateLoadBalanceUDPListenerRequest{}
	b := &UpdateLoadBalanceUDPListenerRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerRequestBuilder) UpdateLoadBalanceUDPListenerBody(v *UpdateLoadBalanceUDPListenerBody) *UpdateLoadBalanceUDPListenerRequestBuilder {
	b.s.UpdateLoadBalanceUDPListenerBody = v
	return b
}

func (b *UpdateLoadBalanceUDPListenerRequestBuilder) Build() *UpdateLoadBalanceUDPListenerRequest {
	return b.s
}


