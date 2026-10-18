// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceHTTPListenerPoolResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceHTTPListenerPoolResponseStateEnumOk UpdateLoadBalanceHTTPListenerPoolResponseStateEnum = "OK"
    UpdateLoadBalanceHTTPListenerPoolResponseStateEnumError UpdateLoadBalanceHTTPListenerPoolResponseStateEnum = "ERROR"
    UpdateLoadBalanceHTTPListenerPoolResponseStateEnumException UpdateLoadBalanceHTTPListenerPoolResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceHTTPListenerPoolResponseStateEnumAlarm UpdateLoadBalanceHTTPListenerPoolResponseStateEnum = "ALARM"
    UpdateLoadBalanceHTTPListenerPoolResponseStateEnumForbidden UpdateLoadBalanceHTTPListenerPoolResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceHTTPListenerPoolResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceHTTPListenerPoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerPoolResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerPoolResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerPoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetRequestId(v string) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetErrorMessage(v string) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetErrorCode(v string) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetState(v UpdateLoadBalanceHTTPListenerPoolResponseStateEnum) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetBody(v string) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolResponse) SetErrorParams(v []string) *UpdateLoadBalanceHTTPListenerPoolResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceHTTPListenerPoolResponseBuilder struct {
	s *UpdateLoadBalanceHTTPListenerPoolResponse
}

func NewUpdateLoadBalanceHTTPListenerPoolResponseBuilder() *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	s := &UpdateLoadBalanceHTTPListenerPoolResponse{}
	b := &UpdateLoadBalanceHTTPListenerPoolResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) RequestId(v string) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) State(v UpdateLoadBalanceHTTPListenerPoolResponseStateEnum) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) Body(v string) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceHTTPListenerPoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolResponseBuilder) Build() *UpdateLoadBalanceHTTPListenerPoolResponse {
	return b.s
}


