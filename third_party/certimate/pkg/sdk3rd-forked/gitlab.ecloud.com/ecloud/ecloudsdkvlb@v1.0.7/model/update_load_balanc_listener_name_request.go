// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalancListenerNameRequest struct {


	UpdateLoadBalancListenerNameBody *UpdateLoadBalancListenerNameBody `json:"updateLoadBalancListenerNameBody,omitempty"`
}

func (s UpdateLoadBalancListenerNameRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancListenerNameRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancListenerNameRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancListenerNameRequest) SetUpdateLoadBalancListenerNameBody(v *UpdateLoadBalancListenerNameBody) *UpdateLoadBalancListenerNameRequest {
	s.UpdateLoadBalancListenerNameBody = v
	return s
}


type UpdateLoadBalancListenerNameRequestBuilder struct {
	s *UpdateLoadBalancListenerNameRequest
}

func NewUpdateLoadBalancListenerNameRequestBuilder() *UpdateLoadBalancListenerNameRequestBuilder {
	s := &UpdateLoadBalancListenerNameRequest{}
	b := &UpdateLoadBalancListenerNameRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancListenerNameRequestBuilder) UpdateLoadBalancListenerNameBody(v *UpdateLoadBalancListenerNameBody) *UpdateLoadBalancListenerNameRequestBuilder {
	b.s.UpdateLoadBalancListenerNameBody = v
	return b
}

func (b *UpdateLoadBalancListenerNameRequestBuilder) Build() *UpdateLoadBalancListenerNameRequest {
	return b.s
}


