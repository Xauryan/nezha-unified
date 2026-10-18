// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPSListenerPoolRequest struct {


	UpdateLoadBalanceHTTPSListenerPoolBody *UpdateLoadBalanceHTTPSListenerPoolBody `json:"updateLoadBalanceHTTPSListenerPoolBody,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerPoolRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerPoolRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerPoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerPoolRequest) SetUpdateLoadBalanceHTTPSListenerPoolBody(v *UpdateLoadBalanceHTTPSListenerPoolBody) *UpdateLoadBalanceHTTPSListenerPoolRequest {
	s.UpdateLoadBalanceHTTPSListenerPoolBody = v
	return s
}


type UpdateLoadBalanceHTTPSListenerPoolRequestBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerPoolRequest
}

func NewUpdateLoadBalanceHTTPSListenerPoolRequestBuilder() *UpdateLoadBalanceHTTPSListenerPoolRequestBuilder {
	s := &UpdateLoadBalanceHTTPSListenerPoolRequest{}
	b := &UpdateLoadBalanceHTTPSListenerPoolRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolRequestBuilder) UpdateLoadBalanceHTTPSListenerPoolBody(v *UpdateLoadBalanceHTTPSListenerPoolBody) *UpdateLoadBalanceHTTPSListenerPoolRequestBuilder {
	b.s.UpdateLoadBalanceHTTPSListenerPoolBody = v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolRequestBuilder) Build() *UpdateLoadBalanceHTTPSListenerPoolRequest {
	return b.s
}


