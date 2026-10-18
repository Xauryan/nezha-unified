// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateMembersResponseStateEnum string

// List of State
const (
    CreateMembersResponseStateEnumOk CreateMembersResponseStateEnum = "OK"
    CreateMembersResponseStateEnumError CreateMembersResponseStateEnum = "ERROR"
    CreateMembersResponseStateEnumException CreateMembersResponseStateEnum = "EXCEPTION"
    CreateMembersResponseStateEnumAlarm CreateMembersResponseStateEnum = "ALARM"
    CreateMembersResponseStateEnumForbidden CreateMembersResponseStateEnum = "FORBIDDEN"
)

type CreateMembersResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *CreateMembersResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateMembersResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersResponse) GoString() string {
	return s.String()
}

func (s CreateMembersResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersResponse) SetAsyncKey(v string) *CreateMembersResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateMembersResponse) SetRequestId(v string) *CreateMembersResponse {
	s.RequestId = &v
	return s
}

func (s *CreateMembersResponse) SetErrorMessage(v string) *CreateMembersResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMembersResponse) SetErrorCode(v string) *CreateMembersResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateMembersResponse) SetState(v CreateMembersResponseStateEnum) *CreateMembersResponse {
	s.State = &v
	return s
}

func (s *CreateMembersResponse) SetBody(v bool) *CreateMembersResponse {
	s.Body = &v
	return s
}

func (s *CreateMembersResponse) SetErrorParams(v []string) *CreateMembersResponse {
	s.ErrorParams = v
	return s
}


type CreateMembersResponseBuilder struct {
	s *CreateMembersResponse
}

func NewCreateMembersResponseBuilder() *CreateMembersResponseBuilder {
	s := &CreateMembersResponse{}
	b := &CreateMembersResponseBuilder{s: s}
	return b
}

func (b *CreateMembersResponseBuilder) AsyncKey(v string) *CreateMembersResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateMembersResponseBuilder) RequestId(v string) *CreateMembersResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateMembersResponseBuilder) ErrorMessage(v string) *CreateMembersResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateMembersResponseBuilder) ErrorCode(v string) *CreateMembersResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateMembersResponseBuilder) State(v CreateMembersResponseStateEnum) *CreateMembersResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateMembersResponseBuilder) Body(v bool) *CreateMembersResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateMembersResponseBuilder) ErrorParams(v []string) *CreateMembersResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateMembersResponseBuilder) Build() *CreateMembersResponse {
	return b.s
}


