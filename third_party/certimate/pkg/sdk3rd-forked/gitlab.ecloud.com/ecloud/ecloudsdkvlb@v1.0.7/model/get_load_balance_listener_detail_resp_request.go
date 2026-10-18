// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceListenerDetailRespRequest struct {


	GetLoadBalanceListenerDetailRespPath *GetLoadBalanceListenerDetailRespPath `json:"getLoadBalanceListenerDetailRespPath,omitempty"`
}

func (s GetLoadBalanceListenerDetailRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceListenerDetailRespRequest) GoString() string {
	return s.String()
}

func (s GetLoadBalanceListenerDetailRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceListenerDetailRespRequest) SetGetLoadBalanceListenerDetailRespPath(v *GetLoadBalanceListenerDetailRespPath) *GetLoadBalanceListenerDetailRespRequest {
	s.GetLoadBalanceListenerDetailRespPath = v
	return s
}


type GetLoadBalanceListenerDetailRespRequestBuilder struct {
	s *GetLoadBalanceListenerDetailRespRequest
}

func NewGetLoadBalanceListenerDetailRespRequestBuilder() *GetLoadBalanceListenerDetailRespRequestBuilder {
	s := &GetLoadBalanceListenerDetailRespRequest{}
	b := &GetLoadBalanceListenerDetailRespRequestBuilder{s: s}
	return b
}

func (b *GetLoadBalanceListenerDetailRespRequestBuilder) GetLoadBalanceListenerDetailRespPath(v *GetLoadBalanceListenerDetailRespPath) *GetLoadBalanceListenerDetailRespRequestBuilder {
	b.s.GetLoadBalanceListenerDetailRespPath = v
	return b
}

func (b *GetLoadBalanceListenerDetailRespRequestBuilder) Build() *GetLoadBalanceListenerDetailRespRequest {
	return b.s
}


