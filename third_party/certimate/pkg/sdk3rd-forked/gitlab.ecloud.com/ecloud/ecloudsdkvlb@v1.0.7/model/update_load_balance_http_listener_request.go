// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPListenerRequest struct {


	UpdateLoadBalanceHTTPListenerBody *UpdateLoadBalanceHTTPListenerBody `json:"updateLoadBalanceHTTPListenerBody,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerRequest) SetUpdateLoadBalanceHTTPListenerBody(v *UpdateLoadBalanceHTTPListenerBody) *UpdateLoadBalanceHTTPListenerRequest {
	s.UpdateLoadBalanceHTTPListenerBody = v
	return s
}


type UpdateLoadBalanceHTTPListenerRequestBuilder struct {
	s *UpdateLoadBalanceHTTPListenerRequest
}

func NewUpdateLoadBalanceHTTPListenerRequestBuilder() *UpdateLoadBalanceHTTPListenerRequestBuilder {
	s := &UpdateLoadBalanceHTTPListenerRequest{}
	b := &UpdateLoadBalanceHTTPListenerRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerRequestBuilder) UpdateLoadBalanceHTTPListenerBody(v *UpdateLoadBalanceHTTPListenerBody) *UpdateLoadBalanceHTTPListenerRequestBuilder {
	b.s.UpdateLoadBalanceHTTPListenerBody = v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerRequestBuilder) Build() *UpdateLoadBalanceHTTPListenerRequest {
	return b.s
}


