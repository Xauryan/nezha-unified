// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalancePoolMemberRespRequest struct {


	GetLoadBalancePoolMemberRespHeader *GetLoadBalancePoolMemberRespHeader `json:"getLoadBalancePoolMemberRespHeader,omitempty"`

	GetLoadBalancePoolMemberRespPath *GetLoadBalancePoolMemberRespPath `json:"getLoadBalancePoolMemberRespPath,omitempty"`
}

func (s GetLoadBalancePoolMemberRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalancePoolMemberRespRequest) GoString() string {
	return s.String()
}

func (s GetLoadBalancePoolMemberRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalancePoolMemberRespRequest) SetGetLoadBalancePoolMemberRespHeader(v *GetLoadBalancePoolMemberRespHeader) *GetLoadBalancePoolMemberRespRequest {
	s.GetLoadBalancePoolMemberRespHeader = v
	return s
}

func (s *GetLoadBalancePoolMemberRespRequest) SetGetLoadBalancePoolMemberRespPath(v *GetLoadBalancePoolMemberRespPath) *GetLoadBalancePoolMemberRespRequest {
	s.GetLoadBalancePoolMemberRespPath = v
	return s
}


type GetLoadBalancePoolMemberRespRequestBuilder struct {
	s *GetLoadBalancePoolMemberRespRequest
}

func NewGetLoadBalancePoolMemberRespRequestBuilder() *GetLoadBalancePoolMemberRespRequestBuilder {
	s := &GetLoadBalancePoolMemberRespRequest{}
	b := &GetLoadBalancePoolMemberRespRequestBuilder{s: s}
	return b
}

func (b *GetLoadBalancePoolMemberRespRequestBuilder) GetLoadBalancePoolMemberRespHeader(v *GetLoadBalancePoolMemberRespHeader) *GetLoadBalancePoolMemberRespRequestBuilder {
	b.s.GetLoadBalancePoolMemberRespHeader = v
	return b
}

func (b *GetLoadBalancePoolMemberRespRequestBuilder) GetLoadBalancePoolMemberRespPath(v *GetLoadBalancePoolMemberRespPath) *GetLoadBalancePoolMemberRespRequestBuilder {
	b.s.GetLoadBalancePoolMemberRespPath = v
	return b
}

func (b *GetLoadBalancePoolMemberRespRequestBuilder) Build() *GetLoadBalancePoolMemberRespRequest {
	return b.s
}


