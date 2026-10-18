// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceHTTPListenerResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceHTTPListenerResponseStateEnumOk UpdateLoadBalanceHTTPListenerResponseStateEnum = "OK"
    UpdateLoadBalanceHTTPListenerResponseStateEnumError UpdateLoadBalanceHTTPListenerResponseStateEnum = "ERROR"
    UpdateLoadBalanceHTTPListenerResponseStateEnumException UpdateLoadBalanceHTTPListenerResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceHTTPListenerResponseStateEnumAlarm UpdateLoadBalanceHTTPListenerResponseStateEnum = "ALARM"
    UpdateLoadBalanceHTTPListenerResponseStateEnumForbidden UpdateLoadBalanceHTTPListenerResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceHTTPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceHTTPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetRequestId(v string) *UpdateLoadBalanceHTTPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetErrorMessage(v string) *UpdateLoadBalanceHTTPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetErrorCode(v string) *UpdateLoadBalanceHTTPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetState(v UpdateLoadBalanceHTTPListenerResponseStateEnum) *UpdateLoadBalanceHTTPListenerResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetBody(v string) *UpdateLoadBalanceHTTPListenerResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerResponse) SetErrorParams(v []string) *UpdateLoadBalanceHTTPListenerResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceHTTPListenerResponseBuilder struct {
	s *UpdateLoadBalanceHTTPListenerResponse
}

func NewUpdateLoadBalanceHTTPListenerResponseBuilder() *UpdateLoadBalanceHTTPListenerResponseBuilder {
	s := &UpdateLoadBalanceHTTPListenerResponse{}
	b := &UpdateLoadBalanceHTTPListenerResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) RequestId(v string) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) State(v UpdateLoadBalanceHTTPListenerResponseStateEnum) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) Body(v string) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerResponseBuilder) Build() *UpdateLoadBalanceHTTPListenerResponse {
	return b.s
}


