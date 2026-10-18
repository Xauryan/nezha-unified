// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum string

// List of OperateDisplayEnum
const (
    IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnumAllow IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum = "ALLOW"
    IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnumTempDeny IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum = "TEMP_DENY"
    IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnumForeverDeny IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum = "FOREVER_DENY"
)

type IDeleteListenerDeletablesResponseBody struct {

    // 监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 校验监听器能否删除校验结果
	OperateDisplayEnum *IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum `json:"operateDisplayEnum,omitempty"`
    // 校验监听器能否删除信息
	Message *string `json:"message,omitempty"`
}

func (s IDeleteListenerDeletablesResponseBody) String() string {
	return utils.Beautify(s)
}

func (s IDeleteListenerDeletablesResponseBody) GoString() string {
	return s.String()
}

func (s IDeleteListenerDeletablesResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IDeleteListenerDeletablesResponseBody) SetListenerId(v string) *IDeleteListenerDeletablesResponseBody {
	s.ListenerId = &v
	return s
}

func (s *IDeleteListenerDeletablesResponseBody) SetOperateDisplayEnum(v IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum) *IDeleteListenerDeletablesResponseBody {
	s.OperateDisplayEnum = &v
	return s
}

func (s *IDeleteListenerDeletablesResponseBody) SetMessage(v string) *IDeleteListenerDeletablesResponseBody {
	s.Message = &v
	return s
}


type IDeleteListenerDeletablesResponseBodyBuilder struct {
	s *IDeleteListenerDeletablesResponseBody
}

func NewIDeleteListenerDeletablesResponseBodyBuilder() *IDeleteListenerDeletablesResponseBodyBuilder {
	s := &IDeleteListenerDeletablesResponseBody{}
	b := &IDeleteListenerDeletablesResponseBodyBuilder{s: s}
	return b
}

func (b *IDeleteListenerDeletablesResponseBodyBuilder) ListenerId(v string) *IDeleteListenerDeletablesResponseBodyBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBodyBuilder) OperateDisplayEnum(v IDeleteListenerDeletablesResponseBodyOperateDisplayEnumEnum) *IDeleteListenerDeletablesResponseBodyBuilder {
	b.s.OperateDisplayEnum = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBodyBuilder) Message(v string) *IDeleteListenerDeletablesResponseBodyBuilder {
	b.s.Message = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBodyBuilder) Build() *IDeleteListenerDeletablesResponseBody {
	return b.s
}


