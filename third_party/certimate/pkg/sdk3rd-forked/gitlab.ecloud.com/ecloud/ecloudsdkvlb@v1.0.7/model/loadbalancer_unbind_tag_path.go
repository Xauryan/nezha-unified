// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoadbalancerUnbindTagPath struct {
    position.Path
    // 负载均衡器ID
	LoadbalanceId *string `json:"loadbalanceId,omitempty"`
}

func (s LoadbalancerUnbindTagPath) String() string {
	return utils.Beautify(s)
}

func (s LoadbalancerUnbindTagPath) GoString() string {
	return s.String()
}

func (s LoadbalancerUnbindTagPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoadbalancerUnbindTagPath) SetLoadbalanceId(v string) *LoadbalancerUnbindTagPath {
	s.LoadbalanceId = &v
	return s
}


type LoadbalancerUnbindTagPathBuilder struct {
	s *LoadbalancerUnbindTagPath
}

func NewLoadbalancerUnbindTagPathBuilder() *LoadbalancerUnbindTagPathBuilder {
	s := &LoadbalancerUnbindTagPath{}
	b := &LoadbalancerUnbindTagPathBuilder{s: s}
	return b
}

func (b *LoadbalancerUnbindTagPathBuilder) LoadbalanceId(v string) *LoadbalancerUnbindTagPathBuilder {
	b.s.LoadbalanceId = &v
	return b
}

func (b *LoadbalancerUnbindTagPathBuilder) Build() *LoadbalancerUnbindTagPath {
	return b.s
}


