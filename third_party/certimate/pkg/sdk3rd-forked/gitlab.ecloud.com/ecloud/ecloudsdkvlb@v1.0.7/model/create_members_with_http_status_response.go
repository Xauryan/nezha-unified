// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateMembersWithHttpStatusResponseStateEnum string

// List of State
const (
    CreateMembersWithHttpStatusResponseStateEnumOk CreateMembersWithHttpStatusResponseStateEnum = "OK"
    CreateMembersWithHttpStatusResponseStateEnumError CreateMembersWithHttpStatusResponseStateEnum = "ERROR"
    CreateMembersWithHttpStatusResponseStateEnumException CreateMembersWithHttpStatusResponseStateEnum = "EXCEPTION"
    CreateMembersWithHttpStatusResponseStateEnumAlarm CreateMembersWithHttpStatusResponseStateEnum = "ALARM"
    CreateMembersWithHttpStatusResponseStateEnumForbidden CreateMembersWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type CreateMembersWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *CreateMembersWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateMembersWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s CreateMembersWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersWithHttpStatusResponse) SetAsyncKey(v string) *CreateMembersWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetRequestId(v string) *CreateMembersWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetErrorMessage(v string) *CreateMembersWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetErrorCode(v string) *CreateMembersWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetState(v CreateMembersWithHttpStatusResponseStateEnum) *CreateMembersWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetBody(v bool) *CreateMembersWithHttpStatusResponse {
	s.Body = &v
	return s
}

func (s *CreateMembersWithHttpStatusResponse) SetErrorParams(v []string) *CreateMembersWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type CreateMembersWithHttpStatusResponseBuilder struct {
	s *CreateMembersWithHttpStatusResponse
}

func NewCreateMembersWithHttpStatusResponseBuilder() *CreateMembersWithHttpStatusResponseBuilder {
	s := &CreateMembersWithHttpStatusResponse{}
	b := &CreateMembersWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) AsyncKey(v string) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) RequestId(v string) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) ErrorMessage(v string) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) ErrorCode(v string) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) State(v CreateMembersWithHttpStatusResponseStateEnum) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) Body(v bool) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) ErrorParams(v []string) *CreateMembersWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateMembersWithHttpStatusResponseBuilder) Build() *CreateMembersWithHttpStatusResponse {
	return b.s
}


