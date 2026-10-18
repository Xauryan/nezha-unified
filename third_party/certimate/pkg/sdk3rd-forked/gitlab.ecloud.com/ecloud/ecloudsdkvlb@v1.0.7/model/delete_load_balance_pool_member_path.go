// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteLoadBalancePoolMemberPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 负载均衡后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s DeleteLoadBalancePoolMemberPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteLoadBalancePoolMemberPath) GoString() string {
	return s.String()
}

func (s DeleteLoadBalancePoolMemberPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteLoadBalancePoolMemberPath) SetPoolId(v string) *DeleteLoadBalancePoolMemberPath {
	s.PoolId = &v
	return s
}

func (s *DeleteLoadBalancePoolMemberPath) SetMemberId(v string) *DeleteLoadBalancePoolMemberPath {
	s.MemberId = &v
	return s
}


type DeleteLoadBalancePoolMemberPathBuilder struct {
	s *DeleteLoadBalancePoolMemberPath
}

func NewDeleteLoadBalancePoolMemberPathBuilder() *DeleteLoadBalancePoolMemberPathBuilder {
	s := &DeleteLoadBalancePoolMemberPath{}
	b := &DeleteLoadBalancePoolMemberPathBuilder{s: s}
	return b
}

func (b *DeleteLoadBalancePoolMemberPathBuilder) PoolId(v string) *DeleteLoadBalancePoolMemberPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *DeleteLoadBalancePoolMemberPathBuilder) MemberId(v string) *DeleteLoadBalancePoolMemberPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *DeleteLoadBalancePoolMemberPathBuilder) Build() *DeleteLoadBalancePoolMemberPath {
	return b.s
}


