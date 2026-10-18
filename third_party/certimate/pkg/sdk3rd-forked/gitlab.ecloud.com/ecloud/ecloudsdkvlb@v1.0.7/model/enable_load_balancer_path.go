// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableLoadBalancerPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s EnableLoadBalancerPath) String() string {
	return utils.Beautify(s)
}

func (s EnableLoadBalancerPath) GoString() string {
	return s.String()
}

func (s EnableLoadBalancerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableLoadBalancerPath) SetLoadBalanceId(v string) *EnableLoadBalancerPath {
	s.LoadBalanceId = &v
	return s
}


type EnableLoadBalancerPathBuilder struct {
	s *EnableLoadBalancerPath
}

func NewEnableLoadBalancerPathBuilder() *EnableLoadBalancerPathBuilder {
	s := &EnableLoadBalancerPath{}
	b := &EnableLoadBalancerPathBuilder{s: s}
	return b
}

func (b *EnableLoadBalancerPathBuilder) LoadBalanceId(v string) *EnableLoadBalancerPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *EnableLoadBalancerPathBuilder) Build() *EnableLoadBalancerPath {
	return b.s
}


