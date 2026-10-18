// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceUDPListenerResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceUDPListenerResponseStateEnumOk UpdateLoadBalanceUDPListenerResponseStateEnum = "OK"
    UpdateLoadBalanceUDPListenerResponseStateEnumError UpdateLoadBalanceUDPListenerResponseStateEnum = "ERROR"
    UpdateLoadBalanceUDPListenerResponseStateEnumException UpdateLoadBalanceUDPListenerResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceUDPListenerResponseStateEnumAlarm UpdateLoadBalanceUDPListenerResponseStateEnum = "ALARM"
    UpdateLoadBalanceUDPListenerResponseStateEnumForbidden UpdateLoadBalanceUDPListenerResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceUDPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceUDPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceUDPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetRequestId(v string) *UpdateLoadBalanceUDPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetErrorMessage(v string) *UpdateLoadBalanceUDPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetErrorCode(v string) *UpdateLoadBalanceUDPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetState(v UpdateLoadBalanceUDPListenerResponseStateEnum) *UpdateLoadBalanceUDPListenerResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetBody(v string) *UpdateLoadBalanceUDPListenerResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerResponse) SetErrorParams(v []string) *UpdateLoadBalanceUDPListenerResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceUDPListenerResponseBuilder struct {
	s *UpdateLoadBalanceUDPListenerResponse
}

func NewUpdateLoadBalanceUDPListenerResponseBuilder() *UpdateLoadBalanceUDPListenerResponseBuilder {
	s := &UpdateLoadBalanceUDPListenerResponse{}
	b := &UpdateLoadBalanceUDPListenerResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) RequestId(v string) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) State(v UpdateLoadBalanceUDPListenerResponseStateEnum) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) Body(v string) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceUDPListenerResponseBuilder) Build() *UpdateLoadBalanceUDPListenerResponse {
	return b.s
}


