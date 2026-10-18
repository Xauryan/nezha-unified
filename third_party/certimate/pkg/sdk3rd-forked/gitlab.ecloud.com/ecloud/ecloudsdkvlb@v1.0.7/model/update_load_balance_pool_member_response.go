// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalancePoolMemberResponseStateEnum string

// List of State
const (
    UpdateLoadBalancePoolMemberResponseStateEnumOk UpdateLoadBalancePoolMemberResponseStateEnum = "OK"
    UpdateLoadBalancePoolMemberResponseStateEnumError UpdateLoadBalancePoolMemberResponseStateEnum = "ERROR"
    UpdateLoadBalancePoolMemberResponseStateEnumException UpdateLoadBalancePoolMemberResponseStateEnum = "EXCEPTION"
    UpdateLoadBalancePoolMemberResponseStateEnumAlarm UpdateLoadBalancePoolMemberResponseStateEnum = "ALARM"
    UpdateLoadBalancePoolMemberResponseStateEnumForbidden UpdateLoadBalancePoolMemberResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalancePoolMemberResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *UpdateLoadBalancePoolMemberResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalancePoolMemberResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancePoolMemberResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancePoolMemberResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancePoolMemberResponse) SetAsyncKey(v string) *UpdateLoadBalancePoolMemberResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetRequestId(v string) *UpdateLoadBalancePoolMemberResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetErrorMessage(v string) *UpdateLoadBalancePoolMemberResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetErrorCode(v string) *UpdateLoadBalancePoolMemberResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetState(v UpdateLoadBalancePoolMemberResponseStateEnum) *UpdateLoadBalancePoolMemberResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetBody(v string) *UpdateLoadBalancePoolMemberResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalancePoolMemberResponse) SetErrorParams(v []string) *UpdateLoadBalancePoolMemberResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalancePoolMemberResponseBuilder struct {
	s *UpdateLoadBalancePoolMemberResponse
}

func NewUpdateLoadBalancePoolMemberResponseBuilder() *UpdateLoadBalancePoolMemberResponseBuilder {
	s := &UpdateLoadBalancePoolMemberResponse{}
	b := &UpdateLoadBalancePoolMemberResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) AsyncKey(v string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) RequestId(v string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) ErrorMessage(v string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) ErrorCode(v string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) State(v UpdateLoadBalancePoolMemberResponseStateEnum) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) Body(v string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) ErrorParams(v []string) *UpdateLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalancePoolMemberResponseBuilder) Build() *UpdateLoadBalancePoolMemberResponse {
	return b.s
}


