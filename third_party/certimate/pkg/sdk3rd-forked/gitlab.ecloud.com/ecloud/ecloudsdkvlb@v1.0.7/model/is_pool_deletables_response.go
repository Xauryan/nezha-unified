// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type IsPoolDeletablesResponseStateEnum string

// List of State
const (
    IsPoolDeletablesResponseStateEnumOk IsPoolDeletablesResponseStateEnum = "OK"
    IsPoolDeletablesResponseStateEnumError IsPoolDeletablesResponseStateEnum = "ERROR"
    IsPoolDeletablesResponseStateEnumException IsPoolDeletablesResponseStateEnum = "EXCEPTION"
    IsPoolDeletablesResponseStateEnumAlarm IsPoolDeletablesResponseStateEnum = "ALARM"
    IsPoolDeletablesResponseStateEnumForbidden IsPoolDeletablesResponseStateEnum = "FORBIDDEN"
)

type IsPoolDeletablesResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *IsPoolDeletablesResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]IsPoolDeletablesResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s IsPoolDeletablesResponse) String() string {
	return utils.Beautify(s)
}

func (s IsPoolDeletablesResponse) GoString() string {
	return s.String()
}

func (s IsPoolDeletablesResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IsPoolDeletablesResponse) SetAsyncKey(v string) *IsPoolDeletablesResponse {
	s.AsyncKey = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetRequestId(v string) *IsPoolDeletablesResponse {
	s.RequestId = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetErrorMessage(v string) *IsPoolDeletablesResponse {
	s.ErrorMessage = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetErrorCode(v string) *IsPoolDeletablesResponse {
	s.ErrorCode = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetState(v IsPoolDeletablesResponseStateEnum) *IsPoolDeletablesResponse {
	s.State = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetBody(v []IsPoolDeletablesResponseBody) *IsPoolDeletablesResponse {
	s.Body = &v
	return s
}

func (s *IsPoolDeletablesResponse) SetErrorParams(v []string) *IsPoolDeletablesResponse {
	s.ErrorParams = v
	return s
}


type IsPoolDeletablesResponseBuilder struct {
	s *IsPoolDeletablesResponse
}

func NewIsPoolDeletablesResponseBuilder() *IsPoolDeletablesResponseBuilder {
	s := &IsPoolDeletablesResponse{}
	b := &IsPoolDeletablesResponseBuilder{s: s}
	return b
}

func (b *IsPoolDeletablesResponseBuilder) AsyncKey(v string) *IsPoolDeletablesResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) RequestId(v string) *IsPoolDeletablesResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) ErrorMessage(v string) *IsPoolDeletablesResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) ErrorCode(v string) *IsPoolDeletablesResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) State(v IsPoolDeletablesResponseStateEnum) *IsPoolDeletablesResponseBuilder {
	b.s.State = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) Body(v []IsPoolDeletablesResponseBody) *IsPoolDeletablesResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) ErrorParams(v []string) *IsPoolDeletablesResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *IsPoolDeletablesResponseBuilder) Build() *IsPoolDeletablesResponse {
	return b.s
}


