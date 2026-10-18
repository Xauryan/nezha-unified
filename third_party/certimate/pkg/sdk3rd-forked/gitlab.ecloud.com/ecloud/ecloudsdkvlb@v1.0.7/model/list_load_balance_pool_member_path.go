// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalancePoolMemberPath struct {
    position.Path
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s ListLoadBalancePoolMemberPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberPath) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberPath) SetPoolId(v string) *ListLoadBalancePoolMemberPath {
	s.PoolId = &v
	return s
}


type ListLoadBalancePoolMemberPathBuilder struct {
	s *ListLoadBalancePoolMemberPath
}

func NewListLoadBalancePoolMemberPathBuilder() *ListLoadBalancePoolMemberPathBuilder {
	s := &ListLoadBalancePoolMemberPath{}
	b := &ListLoadBalancePoolMemberPathBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberPathBuilder) PoolId(v string) *ListLoadBalancePoolMemberPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalancePoolMemberPathBuilder) Build() *ListLoadBalancePoolMemberPath {
	return b.s
}


