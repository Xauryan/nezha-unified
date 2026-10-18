// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceHTTPListenerResponseStateEnum string

// List of State
const (
    CreateLoadBalanceHTTPListenerResponseStateEnumOk CreateLoadBalanceHTTPListenerResponseStateEnum = "OK"
    CreateLoadBalanceHTTPListenerResponseStateEnumError CreateLoadBalanceHTTPListenerResponseStateEnum = "ERROR"
    CreateLoadBalanceHTTPListenerResponseStateEnumException CreateLoadBalanceHTTPListenerResponseStateEnum = "EXCEPTION"
    CreateLoadBalanceHTTPListenerResponseStateEnumAlarm CreateLoadBalanceHTTPListenerResponseStateEnum = "ALARM"
    CreateLoadBalanceHTTPListenerResponseStateEnumForbidden CreateLoadBalanceHTTPListenerResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalanceHTTPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadBalanceHTTPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalanceHTTPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPListenerResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetRequestId(v string) *CreateLoadBalanceHTTPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetErrorMessage(v string) *CreateLoadBalanceHTTPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetErrorCode(v string) *CreateLoadBalanceHTTPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetState(v CreateLoadBalanceHTTPListenerResponseStateEnum) *CreateLoadBalanceHTTPListenerResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetBody(v string) *CreateLoadBalanceHTTPListenerResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerResponse) SetErrorParams(v []string) *CreateLoadBalanceHTTPListenerResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalanceHTTPListenerResponseBuilder struct {
	s *CreateLoadBalanceHTTPListenerResponse
}

func NewCreateLoadBalanceHTTPListenerResponseBuilder() *CreateLoadBalanceHTTPListenerResponseBuilder {
	s := &CreateLoadBalanceHTTPListenerResponse{}
	b := &CreateLoadBalanceHTTPListenerResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) RequestId(v string) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) ErrorMessage(v string) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) ErrorCode(v string) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) State(v CreateLoadBalanceHTTPListenerResponseStateEnum) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) Body(v string) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) ErrorParams(v []string) *CreateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalanceHTTPListenerResponseBuilder) Build() *CreateLoadBalanceHTTPListenerResponse {
	return b.s
}


