// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceListenerAsyncResponseStateEnum string

// List of State
const (
    CreateLoadBalanceListenerAsyncResponseStateEnumOk CreateLoadBalanceListenerAsyncResponseStateEnum = "OK"
    CreateLoadBalanceListenerAsyncResponseStateEnumError CreateLoadBalanceListenerAsyncResponseStateEnum = "ERROR"
    CreateLoadBalanceListenerAsyncResponseStateEnumException CreateLoadBalanceListenerAsyncResponseStateEnum = "EXCEPTION"
    CreateLoadBalanceListenerAsyncResponseStateEnumAlarm CreateLoadBalanceListenerAsyncResponseStateEnum = "ALARM"
    CreateLoadBalanceListenerAsyncResponseStateEnumForbidden CreateLoadBalanceListenerAsyncResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalanceListenerAsyncResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadBalanceListenerAsyncResponseStateEnum `json:"state,omitempty"`
    // 请求成功后返回的监听器ID
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalanceListenerAsyncResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceListenerAsyncResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceListenerAsyncResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetAsyncKey(v string) *CreateLoadBalanceListenerAsyncResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetRequestId(v string) *CreateLoadBalanceListenerAsyncResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetErrorMessage(v string) *CreateLoadBalanceListenerAsyncResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetErrorCode(v string) *CreateLoadBalanceListenerAsyncResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetState(v CreateLoadBalanceListenerAsyncResponseStateEnum) *CreateLoadBalanceListenerAsyncResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetBody(v string) *CreateLoadBalanceListenerAsyncResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncResponse) SetErrorParams(v []string) *CreateLoadBalanceListenerAsyncResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalanceListenerAsyncResponseBuilder struct {
	s *CreateLoadBalanceListenerAsyncResponse
}

func NewCreateLoadBalanceListenerAsyncResponseBuilder() *CreateLoadBalanceListenerAsyncResponseBuilder {
	s := &CreateLoadBalanceListenerAsyncResponse{}
	b := &CreateLoadBalanceListenerAsyncResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) AsyncKey(v string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) RequestId(v string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) ErrorMessage(v string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) ErrorCode(v string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) State(v CreateLoadBalanceListenerAsyncResponseStateEnum) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) Body(v string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) ErrorParams(v []string) *CreateLoadBalanceListenerAsyncResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalanceListenerAsyncResponseBuilder) Build() *CreateLoadBalanceListenerAsyncResponse {
	return b.s
}


