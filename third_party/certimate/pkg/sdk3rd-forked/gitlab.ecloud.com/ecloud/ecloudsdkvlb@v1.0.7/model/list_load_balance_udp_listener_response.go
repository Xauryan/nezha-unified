// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceUDPListenerResponseStateEnum string

// List of State
const (
    ListLoadBalanceUDPListenerResponseStateEnumOk ListLoadBalanceUDPListenerResponseStateEnum = "OK"
    ListLoadBalanceUDPListenerResponseStateEnumError ListLoadBalanceUDPListenerResponseStateEnum = "ERROR"
    ListLoadBalanceUDPListenerResponseStateEnumException ListLoadBalanceUDPListenerResponseStateEnum = "EXCEPTION"
    ListLoadBalanceUDPListenerResponseStateEnumAlarm ListLoadBalanceUDPListenerResponseStateEnum = "ALARM"
    ListLoadBalanceUDPListenerResponseStateEnumForbidden ListLoadBalanceUDPListenerResponseStateEnum = "FORBIDDEN"
)

type ListLoadBalanceUDPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadBalanceUDPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadBalanceUDPListenerResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadBalanceUDPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerResponse) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerResponse) SetRequestId(v string) *ListLoadBalanceUDPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponse) SetErrorMessage(v string) *ListLoadBalanceUDPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponse) SetErrorCode(v string) *ListLoadBalanceUDPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponse) SetState(v ListLoadBalanceUDPListenerResponseStateEnum) *ListLoadBalanceUDPListenerResponse {
	s.State = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponse) SetBody(v *ListLoadBalanceUDPListenerResponseBody) *ListLoadBalanceUDPListenerResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalanceUDPListenerResponse) SetErrorParams(v []string) *ListLoadBalanceUDPListenerResponse {
	s.ErrorParams = v
	return s
}


type ListLoadBalanceUDPListenerResponseBuilder struct {
	s *ListLoadBalanceUDPListenerResponse
}

func NewListLoadBalanceUDPListenerResponseBuilder() *ListLoadBalanceUDPListenerResponseBuilder {
	s := &ListLoadBalanceUDPListenerResponse{}
	b := &ListLoadBalanceUDPListenerResponseBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) RequestId(v string) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) ErrorMessage(v string) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) ErrorCode(v string) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) State(v ListLoadBalanceUDPListenerResponseStateEnum) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) Body(v *ListLoadBalanceUDPListenerResponseBody) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) ErrorParams(v []string) *ListLoadBalanceUDPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBuilder) Build() *ListLoadBalanceUDPListenerResponse {
	return b.s
}


