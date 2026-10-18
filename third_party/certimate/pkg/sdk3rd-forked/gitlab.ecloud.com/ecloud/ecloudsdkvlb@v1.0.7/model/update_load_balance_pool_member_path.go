// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalancePoolMemberPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 后端服务器组下需更新的后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s UpdateLoadBalancePoolMemberPath) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancePoolMemberPath) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancePoolMemberPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancePoolMemberPath) SetPoolId(v string) *UpdateLoadBalancePoolMemberPath {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberPath) SetMemberId(v string) *UpdateLoadBalancePoolMemberPath {
	s.MemberId = &v
	return s
}


type UpdateLoadBalancePoolMemberPathBuilder struct {
	s *UpdateLoadBalancePoolMemberPath
}

func NewUpdateLoadBalancePoolMemberPathBuilder() *UpdateLoadBalancePoolMemberPathBuilder {
	s := &UpdateLoadBalancePoolMemberPath{}
	b := &UpdateLoadBalancePoolMemberPathBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancePoolMemberPathBuilder) PoolId(v string) *UpdateLoadBalancePoolMemberPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberPathBuilder) MemberId(v string) *UpdateLoadBalancePoolMemberPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberPathBuilder) Build() *UpdateLoadBalancePoolMemberPath {
	return b.s
}


