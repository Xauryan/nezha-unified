// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceUDPListenerPoolResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceUDPListenerPoolResponseStateEnumOk UpdateLoadBalanceUDPListenerPoolResponseStateEnum = "OK"
    UpdateLoadBalanceUDPListenerPoolResponseStateEnumError UpdateLoadBalanceUDPListenerPoolResponseStateEnum = "ERROR"
    UpdateLoadBalanceUDPListenerPoolResponseStateEnumException UpdateLoadBalanceUDPListenerPoolResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceUDPListenerPoolResponseStateEnumAlarm UpdateLoadBalanceUDPListenerPoolResponseStateEnum = "ALARM"
    UpdateLoadBalanceUDPListenerPoolResponseStateEnumForbidden UpdateLoadBalanceUDPListenerPoolResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceUDPListenerPoolResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceUDPListenerPoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceUDPListenerPoolResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerPoolResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerPoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetRequestId(v string) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetErrorMessage(v string) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetErrorCode(v string) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetState(v UpdateLoadBalanceUDPListenerPoolResponseStateEnum) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetBody(v string) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolResponse) SetErrorParams(v []string) *UpdateLoadBalanceUDPListenerPoolResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceUDPListenerPoolResponseBuilder struct {
	s *UpdateLoadBalanceUDPListenerPoolResponse
}

func NewUpdateLoadBalanceUDPListenerPoolResponseBuilder() *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	s := &UpdateLoadBalanceUDPListenerPoolResponse{}
	b := &UpdateLoadBalanceUDPListenerPoolResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) RequestId(v string) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) State(v UpdateLoadBalanceUDPListenerPoolResponseStateEnum) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) Body(v string) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceUDPListenerPoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolResponseBuilder) Build() *UpdateLoadBalanceUDPListenerPoolResponse {
	return b.s
}


