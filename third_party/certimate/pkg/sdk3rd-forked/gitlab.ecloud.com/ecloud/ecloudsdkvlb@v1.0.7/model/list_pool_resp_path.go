// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListPoolRespPath struct {
    position.Path
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ListPoolRespPath) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespPath) GoString() string {
	return s.String()
}

func (s ListPoolRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespPath) SetLoadBalanceId(v string) *ListPoolRespPath {
	s.LoadBalanceId = &v
	return s
}


type ListPoolRespPathBuilder struct {
	s *ListPoolRespPath
}

func NewListPoolRespPathBuilder() *ListPoolRespPathBuilder {
	s := &ListPoolRespPath{}
	b := &ListPoolRespPathBuilder{s: s}
	return b
}

func (b *ListPoolRespPathBuilder) LoadBalanceId(v string) *ListPoolRespPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListPoolRespPathBuilder) Build() *ListPoolRespPath {
	return b.s
}


