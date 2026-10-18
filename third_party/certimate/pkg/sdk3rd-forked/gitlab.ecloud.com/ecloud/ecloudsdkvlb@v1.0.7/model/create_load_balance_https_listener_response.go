// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceHTTPSListenerResponseStateEnum string

// List of State
const (
    CreateLoadBalanceHTTPSListenerResponseStateEnumOk CreateLoadBalanceHTTPSListenerResponseStateEnum = "OK"
    CreateLoadBalanceHTTPSListenerResponseStateEnumError CreateLoadBalanceHTTPSListenerResponseStateEnum = "ERROR"
    CreateLoadBalanceHTTPSListenerResponseStateEnumException CreateLoadBalanceHTTPSListenerResponseStateEnum = "EXCEPTION"
    CreateLoadBalanceHTTPSListenerResponseStateEnumAlarm CreateLoadBalanceHTTPSListenerResponseStateEnum = "ALARM"
    CreateLoadBalanceHTTPSListenerResponseStateEnumForbidden CreateLoadBalanceHTTPSListenerResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalanceHTTPSListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadBalanceHTTPSListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalanceHTTPSListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPSListenerResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPSListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetRequestId(v string) *CreateLoadBalanceHTTPSListenerResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetErrorMessage(v string) *CreateLoadBalanceHTTPSListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetErrorCode(v string) *CreateLoadBalanceHTTPSListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetState(v CreateLoadBalanceHTTPSListenerResponseStateEnum) *CreateLoadBalanceHTTPSListenerResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetBody(v string) *CreateLoadBalanceHTTPSListenerResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerResponse) SetErrorParams(v []string) *CreateLoadBalanceHTTPSListenerResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalanceHTTPSListenerResponseBuilder struct {
	s *CreateLoadBalanceHTTPSListenerResponse
}

func NewCreateLoadBalanceHTTPSListenerResponseBuilder() *CreateLoadBalanceHTTPSListenerResponseBuilder {
	s := &CreateLoadBalanceHTTPSListenerResponse{}
	b := &CreateLoadBalanceHTTPSListenerResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) RequestId(v string) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) ErrorMessage(v string) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) ErrorCode(v string) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) State(v CreateLoadBalanceHTTPSListenerResponseStateEnum) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) Body(v string) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) ErrorParams(v []string) *CreateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerResponseBuilder) Build() *CreateLoadBalanceHTTPSListenerResponse {
	return b.s
}


