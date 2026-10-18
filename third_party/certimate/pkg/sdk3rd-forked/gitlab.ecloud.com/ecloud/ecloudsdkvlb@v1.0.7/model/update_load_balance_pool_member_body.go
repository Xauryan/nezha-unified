// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalancePoolMemberBody struct {
    position.Body
    // 权重
	Weight *int32 `json:"weight,omitempty"`
}

func (s UpdateLoadBalancePoolMemberBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancePoolMemberBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancePoolMemberBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancePoolMemberBody) SetWeight(v int32) *UpdateLoadBalancePoolMemberBody {
	s.Weight = &v
	return s
}


type UpdateLoadBalancePoolMemberBodyBuilder struct {
	s *UpdateLoadBalancePoolMemberBody
}

func NewUpdateLoadBalancePoolMemberBodyBuilder() *UpdateLoadBalancePoolMemberBodyBuilder {
	s := &UpdateLoadBalancePoolMemberBody{}
	b := &UpdateLoadBalancePoolMemberBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancePoolMemberBodyBuilder) Weight(v int32) *UpdateLoadBalancePoolMemberBodyBuilder {
	b.s.Weight = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberBodyBuilder) Build() *UpdateLoadBalancePoolMemberBody {
	return b.s
}


