// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type IDeleteListenerDeletablesResponseStateEnum string

// List of State
const (
    IDeleteListenerDeletablesResponseStateEnumOk IDeleteListenerDeletablesResponseStateEnum = "OK"
    IDeleteListenerDeletablesResponseStateEnumError IDeleteListenerDeletablesResponseStateEnum = "ERROR"
    IDeleteListenerDeletablesResponseStateEnumException IDeleteListenerDeletablesResponseStateEnum = "EXCEPTION"
    IDeleteListenerDeletablesResponseStateEnumAlarm IDeleteListenerDeletablesResponseStateEnum = "ALARM"
    IDeleteListenerDeletablesResponseStateEnumForbidden IDeleteListenerDeletablesResponseStateEnum = "FORBIDDEN"
)

type IDeleteListenerDeletablesResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *IDeleteListenerDeletablesResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]IDeleteListenerDeletablesResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s IDeleteListenerDeletablesResponse) String() string {
	return utils.Beautify(s)
}

func (s IDeleteListenerDeletablesResponse) GoString() string {
	return s.String()
}

func (s IDeleteListenerDeletablesResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IDeleteListenerDeletablesResponse) SetAsyncKey(v string) *IDeleteListenerDeletablesResponse {
	s.AsyncKey = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetRequestId(v string) *IDeleteListenerDeletablesResponse {
	s.RequestId = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetErrorMessage(v string) *IDeleteListenerDeletablesResponse {
	s.ErrorMessage = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetErrorCode(v string) *IDeleteListenerDeletablesResponse {
	s.ErrorCode = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetState(v IDeleteListenerDeletablesResponseStateEnum) *IDeleteListenerDeletablesResponse {
	s.State = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetBody(v []IDeleteListenerDeletablesResponseBody) *IDeleteListenerDeletablesResponse {
	s.Body = &v
	return s
}

func (s *IDeleteListenerDeletablesResponse) SetErrorParams(v []string) *IDeleteListenerDeletablesResponse {
	s.ErrorParams = v
	return s
}


type IDeleteListenerDeletablesResponseBuilder struct {
	s *IDeleteListenerDeletablesResponse
}

func NewIDeleteListenerDeletablesResponseBuilder() *IDeleteListenerDeletablesResponseBuilder {
	s := &IDeleteListenerDeletablesResponse{}
	b := &IDeleteListenerDeletablesResponseBuilder{s: s}
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) AsyncKey(v string) *IDeleteListenerDeletablesResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) RequestId(v string) *IDeleteListenerDeletablesResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) ErrorMessage(v string) *IDeleteListenerDeletablesResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) ErrorCode(v string) *IDeleteListenerDeletablesResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) State(v IDeleteListenerDeletablesResponseStateEnum) *IDeleteListenerDeletablesResponseBuilder {
	b.s.State = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) Body(v []IDeleteListenerDeletablesResponseBody) *IDeleteListenerDeletablesResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) ErrorParams(v []string) *IDeleteListenerDeletablesResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *IDeleteListenerDeletablesResponseBuilder) Build() *IDeleteListenerDeletablesResponse {
	return b.s
}


