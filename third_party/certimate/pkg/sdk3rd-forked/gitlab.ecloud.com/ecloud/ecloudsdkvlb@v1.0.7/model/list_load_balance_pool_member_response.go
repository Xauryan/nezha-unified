// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalancePoolMemberResponseStateEnum string

// List of State
const (
    ListLoadBalancePoolMemberResponseStateEnumOk ListLoadBalancePoolMemberResponseStateEnum = "OK"
    ListLoadBalancePoolMemberResponseStateEnumError ListLoadBalancePoolMemberResponseStateEnum = "ERROR"
    ListLoadBalancePoolMemberResponseStateEnumException ListLoadBalancePoolMemberResponseStateEnum = "EXCEPTION"
    ListLoadBalancePoolMemberResponseStateEnumAlarm ListLoadBalancePoolMemberResponseStateEnum = "ALARM"
    ListLoadBalancePoolMemberResponseStateEnumForbidden ListLoadBalancePoolMemberResponseStateEnum = "FORBIDDEN"
)

type ListLoadBalancePoolMemberResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *ListLoadBalancePoolMemberResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadBalancePoolMemberResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadBalancePoolMemberResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberResponse) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberResponse) SetAsyncKey(v string) *ListLoadBalancePoolMemberResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetRequestId(v string) *ListLoadBalancePoolMemberResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetErrorMessage(v string) *ListLoadBalancePoolMemberResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetErrorCode(v string) *ListLoadBalancePoolMemberResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetState(v ListLoadBalancePoolMemberResponseStateEnum) *ListLoadBalancePoolMemberResponse {
	s.State = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetBody(v *ListLoadBalancePoolMemberResponseBody) *ListLoadBalancePoolMemberResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalancePoolMemberResponse) SetErrorParams(v []string) *ListLoadBalancePoolMemberResponse {
	s.ErrorParams = v
	return s
}


type ListLoadBalancePoolMemberResponseBuilder struct {
	s *ListLoadBalancePoolMemberResponse
}

func NewListLoadBalancePoolMemberResponseBuilder() *ListLoadBalancePoolMemberResponseBuilder {
	s := &ListLoadBalancePoolMemberResponse{}
	b := &ListLoadBalancePoolMemberResponseBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) AsyncKey(v string) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) RequestId(v string) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) ErrorMessage(v string) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) ErrorCode(v string) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) State(v ListLoadBalancePoolMemberResponseStateEnum) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) Body(v *ListLoadBalancePoolMemberResponseBody) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) ErrorParams(v []string) *ListLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBuilder) Build() *ListLoadBalancePoolMemberResponse {
	return b.s
}


