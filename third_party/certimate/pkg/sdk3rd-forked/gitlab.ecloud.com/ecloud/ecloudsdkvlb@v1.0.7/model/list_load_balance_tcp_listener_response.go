// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceTCPListenerResponseStateEnum string

// List of State
const (
    ListLoadBalanceTCPListenerResponseStateEnumOk ListLoadBalanceTCPListenerResponseStateEnum = "OK"
    ListLoadBalanceTCPListenerResponseStateEnumError ListLoadBalanceTCPListenerResponseStateEnum = "ERROR"
    ListLoadBalanceTCPListenerResponseStateEnumException ListLoadBalanceTCPListenerResponseStateEnum = "EXCEPTION"
    ListLoadBalanceTCPListenerResponseStateEnumAlarm ListLoadBalanceTCPListenerResponseStateEnum = "ALARM"
    ListLoadBalanceTCPListenerResponseStateEnumForbidden ListLoadBalanceTCPListenerResponseStateEnum = "FORBIDDEN"
)

type ListLoadBalanceTCPListenerResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadBalanceTCPListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadBalanceTCPListenerResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadBalanceTCPListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerResponse) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerResponse) SetRequestId(v string) *ListLoadBalanceTCPListenerResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponse) SetErrorMessage(v string) *ListLoadBalanceTCPListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponse) SetErrorCode(v string) *ListLoadBalanceTCPListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponse) SetState(v ListLoadBalanceTCPListenerResponseStateEnum) *ListLoadBalanceTCPListenerResponse {
	s.State = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponse) SetBody(v *ListLoadBalanceTCPListenerResponseBody) *ListLoadBalanceTCPListenerResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalanceTCPListenerResponse) SetErrorParams(v []string) *ListLoadBalanceTCPListenerResponse {
	s.ErrorParams = v
	return s
}


type ListLoadBalanceTCPListenerResponseBuilder struct {
	s *ListLoadBalanceTCPListenerResponse
}

func NewListLoadBalanceTCPListenerResponseBuilder() *ListLoadBalanceTCPListenerResponseBuilder {
	s := &ListLoadBalanceTCPListenerResponse{}
	b := &ListLoadBalanceTCPListenerResponseBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) RequestId(v string) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) ErrorMessage(v string) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) ErrorCode(v string) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) State(v ListLoadBalanceTCPListenerResponseStateEnum) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) Body(v *ListLoadBalanceTCPListenerResponseBody) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) ErrorParams(v []string) *ListLoadBalanceTCPListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBuilder) Build() *ListLoadBalanceTCPListenerResponse {
	return b.s
}


