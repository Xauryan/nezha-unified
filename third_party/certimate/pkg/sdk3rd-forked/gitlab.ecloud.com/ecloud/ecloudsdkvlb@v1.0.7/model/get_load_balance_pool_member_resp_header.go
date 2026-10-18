// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalancePoolMemberRespHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s GetLoadBalancePoolMemberRespHeader) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalancePoolMemberRespHeader) GoString() string {
	return s.String()
}

func (s GetLoadBalancePoolMemberRespHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalancePoolMemberRespHeader) SetRegion(v string) *GetLoadBalancePoolMemberRespHeader {
	s.Region = &v
	return s
}


type GetLoadBalancePoolMemberRespHeaderBuilder struct {
	s *GetLoadBalancePoolMemberRespHeader
}

func NewGetLoadBalancePoolMemberRespHeaderBuilder() *GetLoadBalancePoolMemberRespHeaderBuilder {
	s := &GetLoadBalancePoolMemberRespHeader{}
	b := &GetLoadBalancePoolMemberRespHeaderBuilder{s: s}
	return b
}

func (b *GetLoadBalancePoolMemberRespHeaderBuilder) Region(v string) *GetLoadBalancePoolMemberRespHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespHeaderBuilder) Build() *GetLoadBalancePoolMemberRespHeader {
	return b.s
}


