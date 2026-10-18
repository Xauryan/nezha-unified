// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalanceHTTPSListenerResponseStateEnum string

// List of State
const (
    UpdateLoadBalanceHTTPSListenerResponseStateEnumOk UpdateLoadBalanceHTTPSListenerResponseStateEnum = "OK"
    UpdateLoadBalanceHTTPSListenerResponseStateEnumError UpdateLoadBalanceHTTPSListenerResponseStateEnum = "ERROR"
    UpdateLoadBalanceHTTPSListenerResponseStateEnumException UpdateLoadBalanceHTTPSListenerResponseStateEnum = "EXCEPTION"
    UpdateLoadBalanceHTTPSListenerResponseStateEnumAlarm UpdateLoadBalanceHTTPSListenerResponseStateEnum = "ALARM"
    UpdateLoadBalanceHTTPSListenerResponseStateEnumForbidden UpdateLoadBalanceHTTPSListenerResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalanceHTTPSListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalanceHTTPSListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetRequestId(v string) *UpdateLoadBalanceHTTPSListenerResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetErrorMessage(v string) *UpdateLoadBalanceHTTPSListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetErrorCode(v string) *UpdateLoadBalanceHTTPSListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetState(v UpdateLoadBalanceHTTPSListenerResponseStateEnum) *UpdateLoadBalanceHTTPSListenerResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetBody(v string) *UpdateLoadBalanceHTTPSListenerResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerResponse) SetErrorParams(v []string) *UpdateLoadBalanceHTTPSListenerResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalanceHTTPSListenerResponseBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerResponse
}

func NewUpdateLoadBalanceHTTPSListenerResponseBuilder() *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	s := &UpdateLoadBalanceHTTPSListenerResponse{}
	b := &UpdateLoadBalanceHTTPSListenerResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) RequestId(v string) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) ErrorMessage(v string) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) ErrorCode(v string) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) State(v UpdateLoadBalanceHTTPSListenerResponseStateEnum) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) Body(v string) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) ErrorParams(v []string) *UpdateLoadBalanceHTTPSListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerResponseBuilder) Build() *UpdateLoadBalanceHTTPSListenerResponse {
	return b.s
}


