// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteLoadBalancePoolMemberBody struct {
    position.Body
    // 后端服务器ID列表
	MemberIds []string `json:"memberIds,omitempty"`
}

func (s BatchDeleteLoadBalancePoolMemberBody) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteLoadBalancePoolMemberBody) GoString() string {
	return s.String()
}

func (s BatchDeleteLoadBalancePoolMemberBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteLoadBalancePoolMemberBody) SetMemberIds(v []string) *BatchDeleteLoadBalancePoolMemberBody {
	s.MemberIds = v
	return s
}


type BatchDeleteLoadBalancePoolMemberBodyBuilder struct {
	s *BatchDeleteLoadBalancePoolMemberBody
}

func NewBatchDeleteLoadBalancePoolMemberBodyBuilder() *BatchDeleteLoadBalancePoolMemberBodyBuilder {
	s := &BatchDeleteLoadBalancePoolMemberBody{}
	b := &BatchDeleteLoadBalancePoolMemberBodyBuilder{s: s}
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberBodyBuilder) MemberIds(v []string) *BatchDeleteLoadBalancePoolMemberBodyBuilder {
	b.s.MemberIds = v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberBodyBuilder) Build() *BatchDeleteLoadBalancePoolMemberBody {
	return b.s
}


