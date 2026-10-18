// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteLoadBalancePoolMemberPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s BatchDeleteLoadBalancePoolMemberPath) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteLoadBalancePoolMemberPath) GoString() string {
	return s.String()
}

func (s BatchDeleteLoadBalancePoolMemberPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteLoadBalancePoolMemberPath) SetPoolId(v string) *BatchDeleteLoadBalancePoolMemberPath {
	s.PoolId = &v
	return s
}


type BatchDeleteLoadBalancePoolMemberPathBuilder struct {
	s *BatchDeleteLoadBalancePoolMemberPath
}

func NewBatchDeleteLoadBalancePoolMemberPathBuilder() *BatchDeleteLoadBalancePoolMemberPathBuilder {
	s := &BatchDeleteLoadBalancePoolMemberPath{}
	b := &BatchDeleteLoadBalancePoolMemberPathBuilder{s: s}
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberPathBuilder) PoolId(v string) *BatchDeleteLoadBalancePoolMemberPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberPathBuilder) Build() *BatchDeleteLoadBalancePoolMemberPath {
	return b.s
}


