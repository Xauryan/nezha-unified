// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceUDPListenerRequest struct {


	ListLoadBalanceUDPListenerPath *ListLoadBalanceUDPListenerPath `json:"listLoadBalanceUDPListenerPath,omitempty"`

	ListLoadBalanceUDPListenerQuery *ListLoadBalanceUDPListenerQuery `json:"listLoadBalanceUDPListenerQuery,omitempty"`
}

func (s ListLoadBalanceUDPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerRequest) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerRequest) SetListLoadBalanceUDPListenerPath(v *ListLoadBalanceUDPListenerPath) *ListLoadBalanceUDPListenerRequest {
	s.ListLoadBalanceUDPListenerPath = v
	return s
}

func (s *ListLoadBalanceUDPListenerRequest) SetListLoadBalanceUDPListenerQuery(v *ListLoadBalanceUDPListenerQuery) *ListLoadBalanceUDPListenerRequest {
	s.ListLoadBalanceUDPListenerQuery = v
	return s
}


type ListLoadBalanceUDPListenerRequestBuilder struct {
	s *ListLoadBalanceUDPListenerRequest
}

func NewListLoadBalanceUDPListenerRequestBuilder() *ListLoadBalanceUDPListenerRequestBuilder {
	s := &ListLoadBalanceUDPListenerRequest{}
	b := &ListLoadBalanceUDPListenerRequestBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerRequestBuilder) ListLoadBalanceUDPListenerPath(v *ListLoadBalanceUDPListenerPath) *ListLoadBalanceUDPListenerRequestBuilder {
	b.s.ListLoadBalanceUDPListenerPath = v
	return b
}

func (b *ListLoadBalanceUDPListenerRequestBuilder) ListLoadBalanceUDPListenerQuery(v *ListLoadBalanceUDPListenerQuery) *ListLoadBalanceUDPListenerRequestBuilder {
	b.s.ListLoadBalanceUDPListenerQuery = v
	return b
}

func (b *ListLoadBalanceUDPListenerRequestBuilder) Build() *ListLoadBalanceUDPListenerRequest {
	return b.s
}


