// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceListenerRespRequest struct {


	ListLoadBalanceListenerRespPath *ListLoadBalanceListenerRespPath `json:"listLoadBalanceListenerRespPath,omitempty"`

	ListLoadBalanceListenerRespQuery *ListLoadBalanceListenerRespQuery `json:"listLoadBalanceListenerRespQuery,omitempty"`
}

func (s ListLoadBalanceListenerRespRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespRequest) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespRequest) SetListLoadBalanceListenerRespPath(v *ListLoadBalanceListenerRespPath) *ListLoadBalanceListenerRespRequest {
	s.ListLoadBalanceListenerRespPath = v
	return s
}

func (s *ListLoadBalanceListenerRespRequest) SetListLoadBalanceListenerRespQuery(v *ListLoadBalanceListenerRespQuery) *ListLoadBalanceListenerRespRequest {
	s.ListLoadBalanceListenerRespQuery = v
	return s
}


type ListLoadBalanceListenerRespRequestBuilder struct {
	s *ListLoadBalanceListenerRespRequest
}

func NewListLoadBalanceListenerRespRequestBuilder() *ListLoadBalanceListenerRespRequestBuilder {
	s := &ListLoadBalanceListenerRespRequest{}
	b := &ListLoadBalanceListenerRespRequestBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespRequestBuilder) ListLoadBalanceListenerRespPath(v *ListLoadBalanceListenerRespPath) *ListLoadBalanceListenerRespRequestBuilder {
	b.s.ListLoadBalanceListenerRespPath = v
	return b
}

func (b *ListLoadBalanceListenerRespRequestBuilder) ListLoadBalanceListenerRespQuery(v *ListLoadBalanceListenerRespQuery) *ListLoadBalanceListenerRespRequestBuilder {
	b.s.ListLoadBalanceListenerRespQuery = v
	return b
}

func (b *ListLoadBalanceListenerRespRequestBuilder) Build() *ListLoadBalanceListenerRespRequest {
	return b.s
}


