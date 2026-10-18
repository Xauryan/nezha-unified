// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalancePoolMemberRespPath struct {
    position.Path
    // 后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s GetLoadBalancePoolMemberRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalancePoolMemberRespPath) GoString() string {
	return s.String()
}

func (s GetLoadBalancePoolMemberRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalancePoolMemberRespPath) SetMemberId(v string) *GetLoadBalancePoolMemberRespPath {
	s.MemberId = &v
	return s
}


type GetLoadBalancePoolMemberRespPathBuilder struct {
	s *GetLoadBalancePoolMemberRespPath
}

func NewGetLoadBalancePoolMemberRespPathBuilder() *GetLoadBalancePoolMemberRespPathBuilder {
	s := &GetLoadBalancePoolMemberRespPath{}
	b := &GetLoadBalancePoolMemberRespPathBuilder{s: s}
	return b
}

func (b *GetLoadBalancePoolMemberRespPathBuilder) MemberId(v string) *GetLoadBalancePoolMemberRespPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespPathBuilder) Build() *GetLoadBalancePoolMemberRespPath {
	return b.s
}


