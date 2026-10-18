// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPListenerPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListLoadBalanceHTTPListenerPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerPath) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerPath) SetLoadBalanceId(v string) *ListLoadBalanceHTTPListenerPath {
	s.LoadBalanceId = &v
	return s
}


type ListLoadBalanceHTTPListenerPathBuilder struct {
	s *ListLoadBalanceHTTPListenerPath
}

func NewListLoadBalanceHTTPListenerPathBuilder() *ListLoadBalanceHTTPListenerPathBuilder {
	s := &ListLoadBalanceHTTPListenerPath{}
	b := &ListLoadBalanceHTTPListenerPathBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerPathBuilder) LoadBalanceId(v string) *ListLoadBalanceHTTPListenerPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerPathBuilder) Build() *ListLoadBalanceHTTPListenerPath {
	return b.s
}


