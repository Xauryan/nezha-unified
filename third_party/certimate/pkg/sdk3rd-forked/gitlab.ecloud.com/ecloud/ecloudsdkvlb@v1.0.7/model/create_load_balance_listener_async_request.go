// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalanceListenerAsyncRequest struct {


	CreateLoadBalanceListenerAsyncBody *CreateLoadBalanceListenerAsyncBody `json:"createLoadBalanceListenerAsyncBody,omitempty"`
}

func (s CreateLoadBalanceListenerAsyncRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceListenerAsyncRequest) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceListenerAsyncRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceListenerAsyncRequest) SetCreateLoadBalanceListenerAsyncBody(v *CreateLoadBalanceListenerAsyncBody) *CreateLoadBalanceListenerAsyncRequest {
	s.CreateLoadBalanceListenerAsyncBody = v
	return s
}


type CreateLoadBalanceListenerAsyncRequestBuilder struct {
	s *CreateLoadBalanceListenerAsyncRequest
}

func NewCreateLoadBalanceListenerAsyncRequestBuilder() *CreateLoadBalanceListenerAsyncRequestBuilder {
	s := &CreateLoadBalanceListenerAsyncRequest{}
	b := &CreateLoadBalanceListenerAsyncRequestBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceListenerAsyncRequestBuilder) CreateLoadBalanceListenerAsyncBody(v *CreateLoadBalanceListenerAsyncBody) *CreateLoadBalanceListenerAsyncRequestBuilder {
	b.s.CreateLoadBalanceListenerAsyncBody = v
	return b
}

func (b *CreateLoadBalanceListenerAsyncRequestBuilder) Build() *CreateLoadBalanceListenerAsyncRequest {
	return b.s
}


