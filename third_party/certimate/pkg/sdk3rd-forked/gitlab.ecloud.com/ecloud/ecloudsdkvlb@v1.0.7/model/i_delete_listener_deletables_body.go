// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type IDeleteListenerDeletablesBody struct {
    position.Body
    // 监听器ID列表
	ListenerIds []string 
}

func (s IDeleteListenerDeletablesBody) String() string {
	return utils.Beautify(s)
}

func (s IDeleteListenerDeletablesBody) GoString() string {
	return s.String()
}

func (s IDeleteListenerDeletablesBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IDeleteListenerDeletablesBody) SetListenerIds(v []string) *IDeleteListenerDeletablesBody {
	s.ListenerIds = v
	return s
}


type IDeleteListenerDeletablesBodyBuilder struct {
	s *IDeleteListenerDeletablesBody
}

func NewIDeleteListenerDeletablesBodyBuilder() *IDeleteListenerDeletablesBodyBuilder {
	s := &IDeleteListenerDeletablesBody{}
	b := &IDeleteListenerDeletablesBodyBuilder{s: s}
	return b
}

func (b *IDeleteListenerDeletablesBodyBuilder) ListenerIds(v []string) *IDeleteListenerDeletablesBodyBuilder {
	b.s.ListenerIds = v
	return b
}

func (b *IDeleteListenerDeletablesBodyBuilder) Build() *IDeleteListenerDeletablesBody {
	return b.s
}


