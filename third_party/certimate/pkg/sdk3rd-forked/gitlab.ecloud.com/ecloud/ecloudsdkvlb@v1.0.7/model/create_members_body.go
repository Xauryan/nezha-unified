// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateMembersBody struct {
    position.Body
    // 后端服务器创建请求体
	PoolMemberCreateReqs *[]CreateMembersRequestPoolMemberCreateReqs 
}

func (s CreateMembersBody) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersBody) GoString() string {
	return s.String()
}

func (s CreateMembersBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersBody) SetPoolMemberCreateReqs(v []CreateMembersRequestPoolMemberCreateReqs) *CreateMembersBody {
	s.PoolMemberCreateReqs = &v
	return s
}


type CreateMembersBodyBuilder struct {
	s *CreateMembersBody
}

func NewCreateMembersBodyBuilder() *CreateMembersBodyBuilder {
	s := &CreateMembersBody{}
	b := &CreateMembersBodyBuilder{s: s}
	return b
}

func (b *CreateMembersBodyBuilder) PoolMemberCreateReqs(v []CreateMembersRequestPoolMemberCreateReqs) *CreateMembersBodyBuilder {
	b.s.PoolMemberCreateReqs = &v
	return b
}

func (b *CreateMembersBodyBuilder) Build() *CreateMembersBody {
	return b.s
}


