// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceTCPListenerPoolResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceTCPListenerPoolResponseStateEnumOk UpdateLoadBalanceTCPListenerPoolResponseStateEnum = "OK"
    UpdateLoadBalanceTCPListenerPoolResponseStateEnumError UpdateLoadBalanceTCPListenerPoolResponseStateEnum = "ERROR"
    UpdateLoadBalanceTCPListenerPoolResponseStateEnumException UpdateLoadBalanceTCPListenerPoolResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceTCPListenerPoolResponseStateEnumAlarm UpdateLoadBalanceTCPListenerPoolResponseStateEnum = "ALARM"
    UpdateLoadBalanceTCPListenerPoolResponseStateEnumForbidden UpdateLoadBalanceTCPListenerPoolResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceTCPListenerPoolResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceTCPListenerPoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceTCPListenerPoolResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceTCPListenerPoolResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceTCPListenerPoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetRequestId(v string) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetErrorMessage(v string) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetErrorCode(v string) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetState(v UpdateLoadBalanceTCPListenerPoolResponseStateEnum) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetBody(v string) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolResponse) SetErrorParams(v []string) *UpdateLoadBalanceTCPListenerPoolResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceTCPListenerPoolResponseBuilder struct {
	s *UpdateLoadBalanceTCPListenerPoolResponse
}

func NewUpdateLoadBalanceTCPListenerPoolResponseBuilder() *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	s := &UpdateLoadBalanceTCPListenerPoolResponse{}
	b := &UpdateLoadBalanceTCPListenerPoolResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) RequestId(v string) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) State(v UpdateLoadBalanceTCPListenerPoolResponseStateEnum) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) Body(v string) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceTCPListenerPoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolResponseBuilder) Build() *UpdateLoadBalanceTCPListenerPoolResponse {
	return b.s
}


