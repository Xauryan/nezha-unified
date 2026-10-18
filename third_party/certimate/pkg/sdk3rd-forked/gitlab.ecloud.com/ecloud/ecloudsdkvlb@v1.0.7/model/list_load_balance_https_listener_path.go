// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPSListenerPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListLoadBalanceHTTPSListenerPath) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPSListenerPath) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPSListenerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPSListenerPath) SetLoadBalanceId(v string) *ListLoadBalanceHTTPSListenerPath {
	s.LoadBalanceId = &v
	return s
}


type ListLoadBalanceHTTPSListenerPathBuilder struct {
	s *ListLoadBalanceHTTPSListenerPath
}

func NewListLoadBalanceHTTPSListenerPathBuilder() *ListLoadBalanceHTTPSListenerPathBuilder {
	s := &ListLoadBalanceHTTPSListenerPath{}
	b := &ListLoadBalanceHTTPSListenerPathBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPSListenerPathBuilder) LoadBalanceId(v string) *ListLoadBalanceHTTPSListenerPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerPathBuilder) Build() *ListLoadBalanceHTTPSListenerPath {
	return b.s
}


