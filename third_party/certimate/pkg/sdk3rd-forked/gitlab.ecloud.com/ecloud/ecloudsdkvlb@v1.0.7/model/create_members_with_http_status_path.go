// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersWithHttpStatusPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s CreateMembersWithHttpStatusPath) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersWithHttpStatusPath) GoString() string {
	return s.String()
}

func (s CreateMembersWithHttpStatusPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersWithHttpStatusPath) SetPoolId(v string) *CreateMembersWithHttpStatusPath {
	s.PoolId = &v
	return s
}


type CreateMembersWithHttpStatusPathBuilder struct {
	s *CreateMembersWithHttpStatusPath
}

func NewCreateMembersWithHttpStatusPathBuilder() *CreateMembersWithHttpStatusPathBuilder {
	s := &CreateMembersWithHttpStatusPath{}
	b := &CreateMembersWithHttpStatusPathBuilder{s: s}
	return b
}

func (b *CreateMembersWithHttpStatusPathBuilder) PoolId(v string) *CreateMembersWithHttpStatusPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateMembersWithHttpStatusPathBuilder) Build() *CreateMembersWithHttpStatusPath {
	return b.s
}


