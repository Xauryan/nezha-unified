// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateTLSCustomizeProtocolResponseStateEnum string

// List of State
const (
    CreateTLSCustomizeProtocolResponseStateEnumOk CreateTLSCustomizeProtocolResponseStateEnum = "OK"
    CreateTLSCustomizeProtocolResponseStateEnumError CreateTLSCustomizeProtocolResponseStateEnum = "ERROR"
    CreateTLSCustomizeProtocolResponseStateEnumException CreateTLSCustomizeProtocolResponseStateEnum = "EXCEPTION"
    CreateTLSCustomizeProtocolResponseStateEnumAlarm CreateTLSCustomizeProtocolResponseStateEnum = "ALARM"
    CreateTLSCustomizeProtocolResponseStateEnumForbidden CreateTLSCustomizeProtocolResponseStateEnum = "FORBIDDEN"
)

type CreateTLSCustomizeProtocolResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *CreateTLSCustomizeProtocolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *int64 `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateTLSCustomizeProtocolResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateTLSCustomizeProtocolResponse) GoString() string {
	return s.String()
}

func (s CreateTLSCustomizeProtocolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTLSCustomizeProtocolResponse) SetAsyncKey(v string) *CreateTLSCustomizeProtocolResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetRequestId(v string) *CreateTLSCustomizeProtocolResponse {
	s.RequestId = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetErrorMessage(v string) *CreateTLSCustomizeProtocolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetErrorCode(v string) *CreateTLSCustomizeProtocolResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetState(v CreateTLSCustomizeProtocolResponseStateEnum) *CreateTLSCustomizeProtocolResponse {
	s.State = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetBody(v int64) *CreateTLSCustomizeProtocolResponse {
	s.Body = &v
	return s
}

func (s *CreateTLSCustomizeProtocolResponse) SetErrorParams(v []string) *CreateTLSCustomizeProtocolResponse {
	s.ErrorParams = v
	return s
}


type CreateTLSCustomizeProtocolResponseBuilder struct {
	s *CreateTLSCustomizeProtocolResponse
}

func NewCreateTLSCustomizeProtocolResponseBuilder() *CreateTLSCustomizeProtocolResponseBuilder {
	s := &CreateTLSCustomizeProtocolResponse{}
	b := &CreateTLSCustomizeProtocolResponseBuilder{s: s}
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) AsyncKey(v string) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) RequestId(v string) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) ErrorMessage(v string) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) ErrorCode(v string) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) State(v CreateTLSCustomizeProtocolResponseStateEnum) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) Body(v int64) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) ErrorParams(v []string) *CreateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateTLSCustomizeProtocolResponseBuilder) Build() *CreateTLSCustomizeProtocolResponse {
	return b.s
}


