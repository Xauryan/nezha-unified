// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceUDPListenerPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListLoadBalanceUDPListenerPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerPath) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerPath) SetLoadBalanceId(v string) *ListLoadBalanceUDPListenerPath {
	s.LoadBalanceId = &v
	return s
}


type ListLoadBalanceUDPListenerPathBuilder struct {
	s *ListLoadBalanceUDPListenerPath
}

func NewListLoadBalanceUDPListenerPathBuilder() *ListLoadBalanceUDPListenerPathBuilder {
	s := &ListLoadBalanceUDPListenerPath{}
	b := &ListLoadBalanceUDPListenerPathBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerPathBuilder) LoadBalanceId(v string) *ListLoadBalanceUDPListenerPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerPathBuilder) Build() *ListLoadBalanceUDPListenerPath {
	return b.s
}


