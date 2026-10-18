// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalancePoolMembersPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s CreateLoadBalancePoolMembersPath) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalancePoolMembersPath) GoString() string {
	return s.String()
}

func (s CreateLoadBalancePoolMembersPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalancePoolMembersPath) SetPoolId(v string) *CreateLoadBalancePoolMembersPath {
	s.PoolId = &v
	return s
}


type CreateLoadBalancePoolMembersPathBuilder struct {
	s *CreateLoadBalancePoolMembersPath
}

func NewCreateLoadBalancePoolMembersPathBuilder() *CreateLoadBalancePoolMembersPathBuilder {
	s := &CreateLoadBalancePoolMembersPath{}
	b := &CreateLoadBalancePoolMembersPathBuilder{s: s}
	return b
}

func (b *CreateLoadBalancePoolMembersPathBuilder) PoolId(v string) *CreateLoadBalancePoolMembersPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalancePoolMembersPathBuilder) Build() *CreateLoadBalancePoolMembersPath {
	return b.s
}


