// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalancePoolMembersResponseStateEnum string

// List of State
const (
    CreateLoadBalancePoolMembersResponseStateEnumOk CreateLoadBalancePoolMembersResponseStateEnum = "OK"
    CreateLoadBalancePoolMembersResponseStateEnumError CreateLoadBalancePoolMembersResponseStateEnum = "ERROR"
    CreateLoadBalancePoolMembersResponseStateEnumException CreateLoadBalancePoolMembersResponseStateEnum = "EXCEPTION"
    CreateLoadBalancePoolMembersResponseStateEnumAlarm CreateLoadBalancePoolMembersResponseStateEnum = "ALARM"
    CreateLoadBalancePoolMembersResponseStateEnumForbidden CreateLoadBalancePoolMembersResponseStateEnum = "FORBIDDEN"
)

type CreateLoadBalancePoolMembersResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *CreateLoadBalancePoolMembersResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadBalancePoolMembersResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalancePoolMembersResponse) GoString() string {
	return s.String()
}

func (s CreateLoadBalancePoolMembersResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalancePoolMembersResponse) SetAsyncKey(v string) *CreateLoadBalancePoolMembersResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetRequestId(v string) *CreateLoadBalancePoolMembersResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetErrorMessage(v string) *CreateLoadBalancePoolMembersResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetErrorCode(v string) *CreateLoadBalancePoolMembersResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetState(v CreateLoadBalancePoolMembersResponseStateEnum) *CreateLoadBalancePoolMembersResponse {
	s.State = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetBody(v bool) *CreateLoadBalancePoolMembersResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadBalancePoolMembersResponse) SetErrorParams(v []string) *CreateLoadBalancePoolMembersResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadBalancePoolMembersResponseBuilder struct {
	s *CreateLoadBalancePoolMembersResponse
}

func NewCreateLoadBalancePoolMembersResponseBuilder() *CreateLoadBalancePoolMembersResponseBuilder {
	s := &CreateLoadBalancePoolMembersResponse{}
	b := &CreateLoadBalancePoolMembersResponseBuilder{s: s}
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) AsyncKey(v string) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) RequestId(v string) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) ErrorMessage(v string) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) ErrorCode(v string) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) State(v CreateLoadBalancePoolMembersResponseStateEnum) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) Body(v bool) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) ErrorParams(v []string) *CreateLoadBalancePoolMembersResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadBalancePoolMembersResponseBuilder) Build() *CreateLoadBalancePoolMembersResponse {
	return b.s
}


