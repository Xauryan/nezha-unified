// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersWithHttpStatusBody struct {
    position.Body
    // 负载均衡后端服务器创建请求体
	PoolMemberCreateReqs *[]CreateMembersWithHttpStatusRequestPoolMemberCreateReqs 
}

func (s CreateMembersWithHttpStatusBody) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersWithHttpStatusBody) GoString() string {
	return s.String()
}

func (s CreateMembersWithHttpStatusBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersWithHttpStatusBody) SetPoolMemberCreateReqs(v []CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) *CreateMembersWithHttpStatusBody {
	s.PoolMemberCreateReqs = &v
	return s
}


type CreateMembersWithHttpStatusBodyBuilder struct {
	s *CreateMembersWithHttpStatusBody
}

func NewCreateMembersWithHttpStatusBodyBuilder() *CreateMembersWithHttpStatusBodyBuilder {
	s := &CreateMembersWithHttpStatusBody{}
	b := &CreateMembersWithHttpStatusBodyBuilder{s: s}
	return b
}

func (b *CreateMembersWithHttpStatusBodyBuilder) PoolMemberCreateReqs(v []CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) *CreateMembersWithHttpStatusBodyBuilder {
	b.s.PoolMemberCreateReqs = &v
	return b
}

func (b *CreateMembersWithHttpStatusBodyBuilder) Build() *CreateMembersWithHttpStatusBody {
	return b.s
}


