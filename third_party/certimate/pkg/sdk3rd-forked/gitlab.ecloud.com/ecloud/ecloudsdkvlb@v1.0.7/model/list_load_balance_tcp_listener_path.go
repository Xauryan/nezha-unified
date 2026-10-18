// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceTCPListenerPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListLoadBalanceTCPListenerPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerPath) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerPath) SetLoadBalanceId(v string) *ListLoadBalanceTCPListenerPath {
	s.LoadBalanceId = &v
	return s
}


type ListLoadBalanceTCPListenerPathBuilder struct {
	s *ListLoadBalanceTCPListenerPath
}

func NewListLoadBalanceTCPListenerPathBuilder() *ListLoadBalanceTCPListenerPathBuilder {
	s := &ListLoadBalanceTCPListenerPath{}
	b := &ListLoadBalanceTCPListenerPathBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerPathBuilder) LoadBalanceId(v string) *ListLoadBalanceTCPListenerPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerPathBuilder) Build() *ListLoadBalanceTCPListenerPath {
	return b.s
}


