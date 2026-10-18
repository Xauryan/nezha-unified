// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteMemberWithHttpStatusPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 负载均衡后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s DeleteMemberWithHttpStatusPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteMemberWithHttpStatusPath) GoString() string {
	return s.String()
}

func (s DeleteMemberWithHttpStatusPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteMemberWithHttpStatusPath) SetPoolId(v string) *DeleteMemberWithHttpStatusPath {
	s.PoolId = &v
	return s
}

func (s *DeleteMemberWithHttpStatusPath) SetMemberId(v string) *DeleteMemberWithHttpStatusPath {
	s.MemberId = &v
	return s
}


type DeleteMemberWithHttpStatusPathBuilder struct {
	s *DeleteMemberWithHttpStatusPath
}

func NewDeleteMemberWithHttpStatusPathBuilder() *DeleteMemberWithHttpStatusPathBuilder {
	s := &DeleteMemberWithHttpStatusPath{}
	b := &DeleteMemberWithHttpStatusPathBuilder{s: s}
	return b
}

func (b *DeleteMemberWithHttpStatusPathBuilder) PoolId(v string) *DeleteMemberWithHttpStatusPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *DeleteMemberWithHttpStatusPathBuilder) MemberId(v string) *DeleteMemberWithHttpStatusPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *DeleteMemberWithHttpStatusPathBuilder) Build() *DeleteMemberWithHttpStatusPath {
	return b.s
}


