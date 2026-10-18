// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type IsPoolDeletablesBody struct {
    position.Body
    // 后端服务器组ID列表
	PoolIds []string 
}

func (s IsPoolDeletablesBody) String() string {
	return utils.Beautify(s)
}

func (s IsPoolDeletablesBody) GoString() string {
	return s.String()
}

func (s IsPoolDeletablesBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IsPoolDeletablesBody) SetPoolIds(v []string) *IsPoolDeletablesBody {
	s.PoolIds = v
	return s
}


type IsPoolDeletablesBodyBuilder struct {
	s *IsPoolDeletablesBody
}

func NewIsPoolDeletablesBodyBuilder() *IsPoolDeletablesBodyBuilder {
	s := &IsPoolDeletablesBody{}
	b := &IsPoolDeletablesBodyBuilder{s: s}
	return b
}

func (b *IsPoolDeletablesBodyBuilder) PoolIds(v []string) *IsPoolDeletablesBodyBuilder {
	b.s.PoolIds = v
	return b
}

func (b *IsPoolDeletablesBodyBuilder) Build() *IsPoolDeletablesBody {
	return b.s
}


