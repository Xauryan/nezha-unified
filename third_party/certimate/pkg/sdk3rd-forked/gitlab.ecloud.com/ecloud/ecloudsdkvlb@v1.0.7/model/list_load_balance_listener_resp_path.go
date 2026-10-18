// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceListenerRespPath struct {
    position.Path
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListLoadBalanceListenerRespPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespPath) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespPath) SetLoadBalanceId(v string) *ListLoadBalanceListenerRespPath {
	s.LoadBalanceId = &v
	return s
}


type ListLoadBalanceListenerRespPathBuilder struct {
	s *ListLoadBalanceListenerRespPath
}

func NewListLoadBalanceListenerRespPathBuilder() *ListLoadBalanceListenerRespPathBuilder {
	s := &ListLoadBalanceListenerRespPath{}
	b := &ListLoadBalanceListenerRespPathBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespPathBuilder) LoadBalanceId(v string) *ListLoadBalanceListenerRespPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceListenerRespPathBuilder) Build() *ListLoadBalanceListenerRespPath {
	return b.s
}


