// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableLoadBalancerQuery struct {
    position.Query
    // 是否启用弹性负载均衡
	Enable *bool `json:"enable,omitempty"`
}

func (s EnableLoadBalancerQuery) String() string {
	return utils.Beautify(s)
}

func (s EnableLoadBalancerQuery) GoString() string {
	return s.String()
}

func (s EnableLoadBalancerQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableLoadBalancerQuery) SetEnable(v bool) *EnableLoadBalancerQuery {
	s.Enable = &v
	return s
}


type EnableLoadBalancerQueryBuilder struct {
	s *EnableLoadBalancerQuery
}

func NewEnableLoadBalancerQueryBuilder() *EnableLoadBalancerQueryBuilder {
	s := &EnableLoadBalancerQuery{}
	b := &EnableLoadBalancerQueryBuilder{s: s}
	return b
}

func (b *EnableLoadBalancerQueryBuilder) Enable(v bool) *EnableLoadBalancerQueryBuilder {
	b.s.Enable = &v
	return b
}

func (b *EnableLoadBalancerQueryBuilder) Build() *EnableLoadBalancerQuery {
	return b.s
}


