// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPSListenerRequest struct {


	UpdateLoadBalanceHTTPSListenerBody *UpdateLoadBalanceHTTPSListenerBody `json:"updateLoadBalanceHTTPSListenerBody,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerRequest) SetUpdateLoadBalanceHTTPSListenerBody(v *UpdateLoadBalanceHTTPSListenerBody) *UpdateLoadBalanceHTTPSListenerRequest {
	s.UpdateLoadBalanceHTTPSListenerBody = v
	return s
}


type UpdateLoadBalanceHTTPSListenerRequestBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerRequest
}

func NewUpdateLoadBalanceHTTPSListenerRequestBuilder() *UpdateLoadBalanceHTTPSListenerRequestBuilder {
	s := &UpdateLoadBalanceHTTPSListenerRequest{}
	b := &UpdateLoadBalanceHTTPSListenerRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerRequestBuilder) UpdateLoadBalanceHTTPSListenerBody(v *UpdateLoadBalanceHTTPSListenerBody) *UpdateLoadBalanceHTTPSListenerRequestBuilder {
	b.s.UpdateLoadBalanceHTTPSListenerBody = v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerRequestBuilder) Build() *UpdateLoadBalanceHTTPSListenerRequest {
	return b.s
}


