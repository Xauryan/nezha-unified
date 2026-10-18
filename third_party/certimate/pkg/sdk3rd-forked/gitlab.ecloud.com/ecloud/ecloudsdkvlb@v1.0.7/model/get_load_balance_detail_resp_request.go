// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceDetailRespRequest struct {


	GetLoadBalanceDetailRespQuery *GetLoadBalanceDetailRespQuery `json:"getLoadBalanceDetailRespQuery,omitempty"`

	GetLoadBalanceDetailRespPath *GetLoadBalanceDetailRespPath `json:"getLoadBalanceDetailRespPath,omitempty"`
}

func (s GetLoadBalanceDetailRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespRequest) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespRequest) SetGetLoadBalanceDetailRespQuery(v *GetLoadBalanceDetailRespQuery) *GetLoadBalanceDetailRespRequest {
	s.GetLoadBalanceDetailRespQuery = v
	return s
}

func (s *GetLoadBalanceDetailRespRequest) SetGetLoadBalanceDetailRespPath(v *GetLoadBalanceDetailRespPath) *GetLoadBalanceDetailRespRequest {
	s.GetLoadBalanceDetailRespPath = v
	return s
}


type GetLoadBalanceDetailRespRequestBuilder struct {
	s *GetLoadBalanceDetailRespRequest
}

func NewGetLoadBalanceDetailRespRequestBuilder() *GetLoadBalanceDetailRespRequestBuilder {
	s := &GetLoadBalanceDetailRespRequest{}
	b := &GetLoadBalanceDetailRespRequestBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespRequestBuilder) GetLoadBalanceDetailRespQuery(v *GetLoadBalanceDetailRespQuery) *GetLoadBalanceDetailRespRequestBuilder {
	b.s.GetLoadBalanceDetailRespQuery = v
	return b
}

func (b *GetLoadBalanceDetailRespRequestBuilder) GetLoadBalanceDetailRespPath(v *GetLoadBalanceDetailRespPath) *GetLoadBalanceDetailRespRequestBuilder {
	b.s.GetLoadBalanceDetailRespPath = v
	return b
}

func (b *GetLoadBalanceDetailRespRequestBuilder) Build() *GetLoadBalanceDetailRespRequest {
	return b.s
}


