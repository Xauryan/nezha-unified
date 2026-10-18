// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPListenerPoolRequest struct {


	UpdateLoadBalanceHTTPListenerPoolBody *UpdateLoadBalanceHTTPListenerPoolBody `json:"updateLoadBalanceHTTPListenerPoolBody,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerPoolRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerPoolRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerPoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerPoolRequest) SetUpdateLoadBalanceHTTPListenerPoolBody(v *UpdateLoadBalanceHTTPListenerPoolBody) *UpdateLoadBalanceHTTPListenerPoolRequest {
	s.UpdateLoadBalanceHTTPListenerPoolBody = v
	return s
}


type UpdateLoadBalanceHTTPListenerPoolRequestBuilder struct {
	s *UpdateLoadBalanceHTTPListenerPoolRequest
}

func NewUpdateLoadBalanceHTTPListenerPoolRequestBuilder() *UpdateLoadBalanceHTTPListenerPoolRequestBuilder {
	s := &UpdateLoadBalanceHTTPListenerPoolRequest{}
	b := &UpdateLoadBalanceHTTPListenerPoolRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolRequestBuilder) UpdateLoadBalanceHTTPListenerPoolBody(v *UpdateLoadBalanceHTTPListenerPoolBody) *UpdateLoadBalanceHTTPListenerPoolRequestBuilder {
	b.s.UpdateLoadBalanceHTTPListenerPoolBody = v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolRequestBuilder) Build() *UpdateLoadBalanceHTTPListenerPoolRequest {
	return b.s
}


