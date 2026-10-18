// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateLoadBalancePoolMembersBody struct {
    position.Body
    // 多个负载均衡后端服务器
	PoolMemberCreateReqs *[]CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs `json:"poolMemberCreateReqs,omitempty"`
}

func (s CreateLoadBalancePoolMembersBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalancePoolMembersBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalancePoolMembersBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalancePoolMembersBody) SetPoolMemberCreateReqs(v []CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) *CreateLoadBalancePoolMembersBody {
	s.PoolMemberCreateReqs = &v
	return s
}


type CreateLoadBalancePoolMembersBodyBuilder struct {
	s *CreateLoadBalancePoolMembersBody
}

func NewCreateLoadBalancePoolMembersBodyBuilder() *CreateLoadBalancePoolMembersBodyBuilder {
	s := &CreateLoadBalancePoolMembersBody{}
	b := &CreateLoadBalancePoolMembersBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalancePoolMembersBodyBuilder) PoolMemberCreateReqs(v []CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) *CreateLoadBalancePoolMembersBodyBuilder {
	b.s.PoolMemberCreateReqs = &v
	return b
}

func (b *CreateLoadBalancePoolMembersBodyBuilder) Build() *CreateLoadBalancePoolMembersBody {
	return b.s
}


