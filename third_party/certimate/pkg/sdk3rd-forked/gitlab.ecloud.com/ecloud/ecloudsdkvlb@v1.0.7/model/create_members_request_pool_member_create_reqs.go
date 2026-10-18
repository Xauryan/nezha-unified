// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersRequestPoolMemberCreateReqs struct {

    // 业务主机用来被监听的端口
	Port *int32 `json:"port,omitempty"`
    // 业务主机负载均衡的权重，取值范围：1-100。取值越大，在所有业务主机中占的权重越高，承担的负载均衡的比例越大
	Weight *int32 `json:"weight,omitempty"`
}

func (s CreateMembersRequestPoolMemberCreateReqs) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersRequestPoolMemberCreateReqs) GoString() string {
	return s.String()
}

func (s CreateMembersRequestPoolMemberCreateReqs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersRequestPoolMemberCreateReqs) SetPort(v int32) *CreateMembersRequestPoolMemberCreateReqs {
	s.Port = &v
	return s
}

func (s *CreateMembersRequestPoolMemberCreateReqs) SetWeight(v int32) *CreateMembersRequestPoolMemberCreateReqs {
	s.Weight = &v
	return s
}


type CreateMembersRequestPoolMemberCreateReqsBuilder struct {
	s *CreateMembersRequestPoolMemberCreateReqs
}

func NewCreateMembersRequestPoolMemberCreateReqsBuilder() *CreateMembersRequestPoolMemberCreateReqsBuilder {
	s := &CreateMembersRequestPoolMemberCreateReqs{}
	b := &CreateMembersRequestPoolMemberCreateReqsBuilder{s: s}
	return b
}

func (b *CreateMembersRequestPoolMemberCreateReqsBuilder) Port(v int32) *CreateMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Port = &v
	return b
}

func (b *CreateMembersRequestPoolMemberCreateReqsBuilder) Weight(v int32) *CreateMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Weight = &v
	return b
}

func (b *CreateMembersRequestPoolMemberCreateReqsBuilder) Build() *CreateMembersRequestPoolMemberCreateReqs {
	return b.s
}


