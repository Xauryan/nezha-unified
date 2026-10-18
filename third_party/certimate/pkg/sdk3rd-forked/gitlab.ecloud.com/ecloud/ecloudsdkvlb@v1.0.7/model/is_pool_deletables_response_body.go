// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type IsPoolDeletablesResponseBodyOperateDisplayEnumEnum string

// List of OperateDisplayEnum
const (
    IsPoolDeletablesResponseBodyOperateDisplayEnumEnumAllow IsPoolDeletablesResponseBodyOperateDisplayEnumEnum = "ALLOW"
    IsPoolDeletablesResponseBodyOperateDisplayEnumEnumTempDeny IsPoolDeletablesResponseBodyOperateDisplayEnumEnum = "TEMP_DENY"
    IsPoolDeletablesResponseBodyOperateDisplayEnumEnumForeverDeny IsPoolDeletablesResponseBodyOperateDisplayEnumEnum = "FOREVER_DENY"
)

type IsPoolDeletablesResponseBody struct {

    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 后端服务器组能否删除校验结果，ALLOW：可操作，TEMP_DENY：置灰，FOREVER_DENY：不可见
	OperateDisplayEnum *IsPoolDeletablesResponseBodyOperateDisplayEnumEnum `json:"operateDisplayEnum,omitempty"`
    // 后端服务器组能否删除信息
	Message *string `json:"message,omitempty"`
}

func (s IsPoolDeletablesResponseBody) String() string {
	return utils.Beautify(s)
}

func (s IsPoolDeletablesResponseBody) GoString() string {
	return s.String()
}

func (s IsPoolDeletablesResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IsPoolDeletablesResponseBody) SetPoolId(v string) *IsPoolDeletablesResponseBody {
	s.PoolId = &v
	return s
}

func (s *IsPoolDeletablesResponseBody) SetOperateDisplayEnum(v IsPoolDeletablesResponseBodyOperateDisplayEnumEnum) *IsPoolDeletablesResponseBody {
	s.OperateDisplayEnum = &v
	return s
}

func (s *IsPoolDeletablesResponseBody) SetMessage(v string) *IsPoolDeletablesResponseBody {
	s.Message = &v
	return s
}


type IsPoolDeletablesResponseBodyBuilder struct {
	s *IsPoolDeletablesResponseBody
}

func NewIsPoolDeletablesResponseBodyBuilder() *IsPoolDeletablesResponseBodyBuilder {
	s := &IsPoolDeletablesResponseBody{}
	b := &IsPoolDeletablesResponseBodyBuilder{s: s}
	return b
}

func (b *IsPoolDeletablesResponseBodyBuilder) PoolId(v string) *IsPoolDeletablesResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *IsPoolDeletablesResponseBodyBuilder) OperateDisplayEnum(v IsPoolDeletablesResponseBodyOperateDisplayEnumEnum) *IsPoolDeletablesResponseBodyBuilder {
	b.s.OperateDisplayEnum = &v
	return b
}

func (b *IsPoolDeletablesResponseBodyBuilder) Message(v string) *IsPoolDeletablesResponseBodyBuilder {
	b.s.Message = &v
	return b
}

func (b *IsPoolDeletablesResponseBodyBuilder) Build() *IsPoolDeletablesResponseBody {
	return b.s
}


