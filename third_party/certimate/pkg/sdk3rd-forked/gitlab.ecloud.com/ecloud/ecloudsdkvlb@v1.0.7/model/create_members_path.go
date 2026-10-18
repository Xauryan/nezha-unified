// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersPath struct {
    position.Path
    // 负载均衡后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s CreateMembersPath) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersPath) GoString() string {
	return s.String()
}

func (s CreateMembersPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersPath) SetMemberId(v string) *CreateMembersPath {
	s.MemberId = &v
	return s
}


type CreateMembersPathBuilder struct {
	s *CreateMembersPath
}

func NewCreateMembersPathBuilder() *CreateMembersPathBuilder {
	s := &CreateMembersPath{}
	b := &CreateMembersPathBuilder{s: s}
	return b
}

func (b *CreateMembersPathBuilder) MemberId(v string) *CreateMembersPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *CreateMembersPathBuilder) Build() *CreateMembersPath {
	return b.s
}


