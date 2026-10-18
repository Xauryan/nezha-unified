// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceUDPListenerResponseStateEnum string

// List of State
const (
    CreateLoadBalanceUDPListenerResponseStateEnumOk CreateLoadBalanceUDPListenerResponseStateEnum = "OK"
    CreateLoadBalanceUDPListenerResponseStateEnumError CreateLoadBalanceUDPListenerResponseStateEnum = "ERROR"
    CreateLoadBalanceUDPListenerResponseStateEnumException CreateLoadBalanceUDPListenerResponseStateEnum = "EXCEPTION"
    CreateLoadBalanceUDPListenerResponseStateEnumAlarm CreateLoadBalanceUDPListenerResponseStateEnum = "ALARM"
    CreateLoadBalanceUDPListenerResponseStateEnumForbidden CreateLoadBalanceUDPListenerResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalanceUDPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadBalanceUDPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalanceUDPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceUDPListenerResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceUDPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceUDPListenerResponse) SetRequestId(v string) *CreateLoadBalanceUDPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerResponse) SetErrorMessage(v string) *CreateLoadBalanceUDPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerResponse) SetErrorCode(v string) *CreateLoadBalanceUDPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerResponse) SetState(v CreateLoadBalanceUDPListenerResponseStateEnum) *CreateLoadBalanceUDPListenerResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerResponse) SetBody(v string) *CreateLoadBalanceUDPListenerResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerResponse) SetErrorParams(v []string) *CreateLoadBalanceUDPListenerResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalanceUDPListenerResponseBuilder struct {
	s *CreateLoadBalanceUDPListenerResponse
}

func NewCreateLoadBalanceUDPListenerResponseBuilder() *CreateLoadBalanceUDPListenerResponseBuilder {
	s := &CreateLoadBalanceUDPListenerResponse{}
	b := &CreateLoadBalanceUDPListenerResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) RequestId(v string) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) ErrorMessage(v string) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) ErrorCode(v string) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) State(v CreateLoadBalanceUDPListenerResponseStateEnum) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) Body(v string) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) ErrorParams(v []string) *CreateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalanceUDPListenerResponseBuilder) Build() *CreateLoadBalanceUDPListenerResponse {
	return b.s
}


