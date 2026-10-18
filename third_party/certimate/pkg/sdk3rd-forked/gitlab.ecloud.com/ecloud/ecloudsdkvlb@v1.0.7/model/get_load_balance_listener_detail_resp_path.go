// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceListenerDetailRespPath struct {
    position.Path
    // 监听器的ID
	ListenerId *string `json:"listenerId,omitempty"`
}

func (s GetLoadBalanceListenerDetailRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceListenerDetailRespPath) GoString() string {
	return s.String()
}

func (s GetLoadBalanceListenerDetailRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceListenerDetailRespPath) SetListenerId(v string) *GetLoadBalanceListenerDetailRespPath {
	s.ListenerId = &v
	return s
}


type GetLoadBalanceListenerDetailRespPathBuilder struct {
	s *GetLoadBalanceListenerDetailRespPath
}

func NewGetLoadBalanceListenerDetailRespPathBuilder() *GetLoadBalanceListenerDetailRespPathBuilder {
	s := &GetLoadBalanceListenerDetailRespPath{}
	b := &GetLoadBalanceListenerDetailRespPathBuilder{s: s}
	return b
}

func (b *GetLoadBalanceListenerDetailRespPathBuilder) ListenerId(v string) *GetLoadBalanceListenerDetailRespPathBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespPathBuilder) Build() *GetLoadBalanceListenerDetailRespPath {
	return b.s
}


