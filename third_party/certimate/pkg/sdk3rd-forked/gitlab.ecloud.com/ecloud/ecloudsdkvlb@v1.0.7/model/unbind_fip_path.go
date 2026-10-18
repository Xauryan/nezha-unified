// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnbindFipPath struct {
    position.Path
    // 绑定的负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s UnbindFipPath) String() string {
	return utils.Beautify(s)
}

func (s UnbindFipPath) GoString() string {
	return s.String()
}

func (s UnbindFipPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindFipPath) SetLoadBalanceId(v string) *UnbindFipPath {
	s.LoadBalanceId = &v
	return s
}


type UnbindFipPathBuilder struct {
	s *UnbindFipPath
}

func NewUnbindFipPathBuilder() *UnbindFipPathBuilder {
	s := &UnbindFipPath{}
	b := &UnbindFipPathBuilder{s: s}
	return b
}

func (b *UnbindFipPathBuilder) LoadBalanceId(v string) *UnbindFipPathBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *UnbindFipPathBuilder) Build() *UnbindFipPath {
	return b.s
}


