// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPListenerRequest struct {


	ListLoadBalanceHTTPListenerQuery *ListLoadBalanceHTTPListenerQuery `json:"listLoadBalanceHTTPListenerQuery,omitempty"`

	ListLoadBalanceHTTPListenerPath *ListLoadBalanceHTTPListenerPath `json:"listLoadBalanceHTTPListenerPath,omitempty"`
}

func (s ListLoadBalanceHTTPListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerRequest) SetListLoadBalanceHTTPListenerQuery(v *ListLoadBalanceHTTPListenerQuery) *ListLoadBalanceHTTPListenerRequest {
	s.ListLoadBalanceHTTPListenerQuery = v
	return s
}

func (s *ListLoadBalanceHTTPListenerRequest) SetListLoadBalanceHTTPListenerPath(v *ListLoadBalanceHTTPListenerPath) *ListLoadBalanceHTTPListenerRequest {
	s.ListLoadBalanceHTTPListenerPath = v
	return s
}


type ListLoadBalanceHTTPListenerRequestBuilder struct {
	s *ListLoadBalanceHTTPListenerRequest
}

func NewListLoadBalanceHTTPListenerRequestBuilder() *ListLoadBalanceHTTPListenerRequestBuilder {
	s := &ListLoadBalanceHTTPListenerRequest{}
	b := &ListLoadBalanceHTTPListenerRequestBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerRequestBuilder) ListLoadBalanceHTTPListenerQuery(v *ListLoadBalanceHTTPListenerQuery) *ListLoadBalanceHTTPListenerRequestBuilder {
	b.s.ListLoadBalanceHTTPListenerQuery = v
	return b
}

func (b *ListLoadBalanceHTTPListenerRequestBuilder) ListLoadBalanceHTTPListenerPath(v *ListLoadBalanceHTTPListenerPath) *ListLoadBalanceHTTPListenerRequestBuilder {
	b.s.ListLoadBalanceHTTPListenerPath = v
	return b
}

func (b *ListLoadBalanceHTTPListenerRequestBuilder) Build() *ListLoadBalanceHTTPListenerRequest {
	return b.s
}


