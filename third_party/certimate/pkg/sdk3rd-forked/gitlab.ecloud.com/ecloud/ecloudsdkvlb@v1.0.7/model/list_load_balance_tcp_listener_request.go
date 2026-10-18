// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceTCPListenerRequest struct {


	ListLoadBalanceTCPListenerPath *ListLoadBalanceTCPListenerPath `json:"listLoadBalanceTCPListenerPath,omitempty"`

	ListLoadBalanceTCPListenerQuery *ListLoadBalanceTCPListenerQuery `json:"listLoadBalanceTCPListenerQuery,omitempty"`
}

func (s ListLoadBalanceTCPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerRequest) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerRequest) SetListLoadBalanceTCPListenerPath(v *ListLoadBalanceTCPListenerPath) *ListLoadBalanceTCPListenerRequest {
	s.ListLoadBalanceTCPListenerPath = v
	return s
}

func (s *ListLoadBalanceTCPListenerRequest) SetListLoadBalanceTCPListenerQuery(v *ListLoadBalanceTCPListenerQuery) *ListLoadBalanceTCPListenerRequest {
	s.ListLoadBalanceTCPListenerQuery = v
	return s
}


type ListLoadBalanceTCPListenerRequestBuilder struct {
	s *ListLoadBalanceTCPListenerRequest
}

func NewListLoadBalanceTCPListenerRequestBuilder() *ListLoadBalanceTCPListenerRequestBuilder {
	s := &ListLoadBalanceTCPListenerRequest{}
	b := &ListLoadBalanceTCPListenerRequestBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerRequestBuilder) ListLoadBalanceTCPListenerPath(v *ListLoadBalanceTCPListenerPath) *ListLoadBalanceTCPListenerRequestBuilder {
	b.s.ListLoadBalanceTCPListenerPath = v
	return b
}

func (b *ListLoadBalanceTCPListenerRequestBuilder) ListLoadBalanceTCPListenerQuery(v *ListLoadBalanceTCPListenerQuery) *ListLoadBalanceTCPListenerRequestBuilder {
	b.s.ListLoadBalanceTCPListenerQuery = v
	return b
}

func (b *ListLoadBalanceTCPListenerRequestBuilder) Build() *ListLoadBalanceTCPListenerRequest {
	return b.s
}


