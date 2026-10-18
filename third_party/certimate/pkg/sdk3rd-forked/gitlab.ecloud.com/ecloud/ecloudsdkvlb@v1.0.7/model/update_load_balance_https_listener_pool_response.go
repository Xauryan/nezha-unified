// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceHTTPSListenerPoolResponseStateEnumOk UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum = "OK"
    UpdateLoadBalanceHTTPSListenerPoolResponseStateEnumError UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum = "ERROR"
    UpdateLoadBalanceHTTPSListenerPoolResponseStateEnumException UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceHTTPSListenerPoolResponseStateEnumAlarm UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum = "ALARM"
    UpdateLoadBalanceHTTPSListenerPoolResponseStateEnumForbidden UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceHTTPSListenerPoolResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerPoolResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerPoolResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerPoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetRequestId(v string) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetErrorMessage(v string) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetErrorCode(v string) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetState(v UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetBody(v string) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolResponse) SetErrorParams(v []string) *UpdateLoadBalanceHTTPSListenerPoolResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceHTTPSListenerPoolResponseBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerPoolResponse
}

func NewUpdateLoadBalanceHTTPSListenerPoolResponseBuilder() *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	s := &UpdateLoadBalanceHTTPSListenerPoolResponse{}
	b := &UpdateLoadBalanceHTTPSListenerPoolResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) RequestId(v string) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) State(v UpdateLoadBalanceHTTPSListenerPoolResponseStateEnum) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) Body(v string) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolResponseBuilder) Build() *UpdateLoadBalanceHTTPSListenerPoolResponse {
	return b.s
}


