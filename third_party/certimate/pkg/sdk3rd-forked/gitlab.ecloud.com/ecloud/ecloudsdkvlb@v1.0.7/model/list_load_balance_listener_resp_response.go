// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceListenerRespResponseStateEnum string

// List of State
const (
    ListLoadBalanceListenerRespResponseStateEnumOk ListLoadBalanceListenerRespResponseStateEnum = "OK"
    ListLoadBalanceListenerRespResponseStateEnumError ListLoadBalanceListenerRespResponseStateEnum = "ERROR"
    ListLoadBalanceListenerRespResponseStateEnumException ListLoadBalanceListenerRespResponseStateEnum = "EXCEPTION"
    ListLoadBalanceListenerRespResponseStateEnumAlarm ListLoadBalanceListenerRespResponseStateEnum = "ALARM"
    ListLoadBalanceListenerRespResponseStateEnumForbidden ListLoadBalanceListenerRespResponseStateEnum = "FORBIDDEN"
)

type ListLoadBalanceListenerRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadBalanceListenerRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadBalanceListenerRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadBalanceListenerRespResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespResponse) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespResponse) SetAsyncKey(v string) *ListLoadBalanceListenerRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetRequestId(v string) *ListLoadBalanceListenerRespResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetErrorMessage(v string) *ListLoadBalanceListenerRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetErrorCode(v string) *ListLoadBalanceListenerRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetState(v ListLoadBalanceListenerRespResponseStateEnum) *ListLoadBalanceListenerRespResponse {
	s.State = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetBody(v *ListLoadBalanceListenerRespResponseBody) *ListLoadBalanceListenerRespResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalanceListenerRespResponse) SetErrorParams(v []string) *ListLoadBalanceListenerRespResponse {
	s.ErrorParams = v
	return s
}


type ListLoadBalanceListenerRespResponseBuilder struct {
	s *ListLoadBalanceListenerRespResponse
}

func NewListLoadBalanceListenerRespResponseBuilder() *ListLoadBalanceListenerRespResponseBuilder {
	s := &ListLoadBalanceListenerRespResponse{}
	b := &ListLoadBalanceListenerRespResponseBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) AsyncKey(v string) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) RequestId(v string) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) ErrorMessage(v string) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) ErrorCode(v string) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) State(v ListLoadBalanceListenerRespResponseStateEnum) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) Body(v *ListLoadBalanceListenerRespResponseBody) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) ErrorParams(v []string) *ListLoadBalanceListenerRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBuilder) Build() *ListLoadBalanceListenerRespResponse {
	return b.s
}


