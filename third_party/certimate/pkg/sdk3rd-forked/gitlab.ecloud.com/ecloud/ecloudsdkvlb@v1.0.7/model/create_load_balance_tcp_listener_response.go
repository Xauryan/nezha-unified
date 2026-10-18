// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceTCPListenerResponseStateEnum string

// List of State
const (
    CreateLoadBalanceTCPListenerResponseStateEnumOk CreateLoadBalanceTCPListenerResponseStateEnum = "OK"
    CreateLoadBalanceTCPListenerResponseStateEnumError CreateLoadBalanceTCPListenerResponseStateEnum = "ERROR"
    CreateLoadBalanceTCPListenerResponseStateEnumException CreateLoadBalanceTCPListenerResponseStateEnum = "EXCEPTION"
    CreateLoadBalanceTCPListenerResponseStateEnumAlarm CreateLoadBalanceTCPListenerResponseStateEnum = "ALARM"
    CreateLoadBalanceTCPListenerResponseStateEnumForbidden CreateLoadBalanceTCPListenerResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalanceTCPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadBalanceTCPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalanceTCPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceTCPListenerResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceTCPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceTCPListenerResponse) SetRequestId(v string) *CreateLoadBalanceTCPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerResponse) SetErrorMessage(v string) *CreateLoadBalanceTCPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerResponse) SetErrorCode(v string) *CreateLoadBalanceTCPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerResponse) SetState(v CreateLoadBalanceTCPListenerResponseStateEnum) *CreateLoadBalanceTCPListenerResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerResponse) SetBody(v string) *CreateLoadBalanceTCPListenerResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerResponse) SetErrorParams(v []string) *CreateLoadBalanceTCPListenerResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalanceTCPListenerResponseBuilder struct {
	s *CreateLoadBalanceTCPListenerResponse
}

func NewCreateLoadBalanceTCPListenerResponseBuilder() *CreateLoadBalanceTCPListenerResponseBuilder {
	s := &CreateLoadBalanceTCPListenerResponse{}
	b := &CreateLoadBalanceTCPListenerResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) RequestId(v string) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) ErrorMessage(v string) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) ErrorCode(v string) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) State(v CreateLoadBalanceTCPListenerResponseStateEnum) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) Body(v string) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) ErrorParams(v []string) *CreateLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalanceTCPListenerResponseBuilder) Build() *CreateLoadBalanceTCPListenerResponse {
	return b.s
}


