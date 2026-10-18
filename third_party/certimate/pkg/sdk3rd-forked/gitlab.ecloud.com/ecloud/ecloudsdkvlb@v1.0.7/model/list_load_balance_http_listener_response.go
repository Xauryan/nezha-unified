// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceHTTPListenerResponseStateEnum string

// List of State
const (
    ListLoadBalanceHTTPListenerResponseStateEnumOk ListLoadBalanceHTTPListenerResponseStateEnum = "OK"
    ListLoadBalanceHTTPListenerResponseStateEnumError ListLoadBalanceHTTPListenerResponseStateEnum = "ERROR"
    ListLoadBalanceHTTPListenerResponseStateEnumException ListLoadBalanceHTTPListenerResponseStateEnum = "EXCEPTION"
    ListLoadBalanceHTTPListenerResponseStateEnumAlarm ListLoadBalanceHTTPListenerResponseStateEnum = "ALARM"
    ListLoadBalanceHTTPListenerResponseStateEnumForbidden ListLoadBalanceHTTPListenerResponseStateEnum = "FORBIDDEN"
)

type ListLoadBalanceHTTPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadBalanceHTTPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadBalanceHTTPListenerResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadBalanceHTTPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerResponse) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerResponse) SetRequestId(v string) *ListLoadBalanceHTTPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponse) SetErrorMessage(v string) *ListLoadBalanceHTTPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponse) SetErrorCode(v string) *ListLoadBalanceHTTPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponse) SetState(v ListLoadBalanceHTTPListenerResponseStateEnum) *ListLoadBalanceHTTPListenerResponse {
	s.State = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponse) SetBody(v *ListLoadBalanceHTTPListenerResponseBody) *ListLoadBalanceHTTPListenerResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponse) SetErrorParams(v []string) *ListLoadBalanceHTTPListenerResponse {
	s.ErrorParams = v
	return s
}


type ListLoadBalanceHTTPListenerResponseBuilder struct {
	s *ListLoadBalanceHTTPListenerResponse
}

func NewListLoadBalanceHTTPListenerResponseBuilder() *ListLoadBalanceHTTPListenerResponseBuilder {
	s := &ListLoadBalanceHTTPListenerResponse{}
	b := &ListLoadBalanceHTTPListenerResponseBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) RequestId(v string) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) ErrorMessage(v string) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) ErrorCode(v string) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) State(v ListLoadBalanceHTTPListenerResponseStateEnum) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) Body(v *ListLoadBalanceHTTPListenerResponseBody) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) ErrorParams(v []string) *ListLoadBalanceHTTPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBuilder) Build() *ListLoadBalanceHTTPListenerResponse {
	return b.s
}


