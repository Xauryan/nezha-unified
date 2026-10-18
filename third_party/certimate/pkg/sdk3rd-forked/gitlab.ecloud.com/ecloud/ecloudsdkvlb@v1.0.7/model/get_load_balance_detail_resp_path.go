// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceDetailRespPath struct {
    position.Path
    // 弹性负载均衡的ID或调用API创建弹性负载均衡接口返回的orderId，例如MOP-O-XXX
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s GetLoadBalanceDetailRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespPath) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespPath) SetLoadBalanceId(v string) *GetLoadBalanceDetailRespPath {
	s.LoadBalanceId = &v
	return s
}


type GetLoadBalanceDetailRespPathBuilder struct {
	s *GetLoadBalanceDetailRespPath
}

func NewGetLoadBalanceDetailRespPathBuilder() *GetLoadBalanceDetailRespPathBuilder {
	s := &GetLoadBalanceDetailRespPath{}
	b := &GetLoadBalanceDetailRespPathBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespPathBuilder) LoadBalanceId(v string) *GetLoadBalanceDetailRespPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *GetLoadBalanceDetailRespPathBuilder) Build() *GetLoadBalanceDetailRespPath {
	return b.s
}


