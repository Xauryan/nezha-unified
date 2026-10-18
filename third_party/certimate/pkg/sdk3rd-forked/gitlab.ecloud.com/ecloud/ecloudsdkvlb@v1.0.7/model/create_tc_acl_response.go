// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateTCAclResponseStateEnum string

// List of State
const (
    CreateTCAclResponseStateEnumOk CreateTCAclResponseStateEnum = "OK"
    CreateTCAclResponseStateEnumError CreateTCAclResponseStateEnum = "ERROR"
    CreateTCAclResponseStateEnumException CreateTCAclResponseStateEnum = "EXCEPTION"
    CreateTCAclResponseStateEnumAlarm CreateTCAclResponseStateEnum = "ALARM"
    CreateTCAclResponseStateEnumForbidden CreateTCAclResponseStateEnum = "FORBIDDEN"
)

type CreateTCAclResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateTCAclResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateTCAclResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateTCAclResponse) GoString() string {
	return s.String()
}

func (s CreateTCAclResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTCAclResponse) SetRequestId(v string) *CreateTCAclResponse {
	s.RequestId = &v
	return s
}

func (s *CreateTCAclResponse) SetErrorMessage(v string) *CreateTCAclResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTCAclResponse) SetErrorCode(v string) *CreateTCAclResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateTCAclResponse) SetState(v CreateTCAclResponseStateEnum) *CreateTCAclResponse {
	s.State = &v
	return s
}

func (s *CreateTCAclResponse) SetBody(v string) *CreateTCAclResponse {
	s.Body = &v
	return s
}

func (s *CreateTCAclResponse) SetErrorParams(v []string) *CreateTCAclResponse {
	s.ErrorParams = v
	return s
}


type CreateTCAclResponseBuilder struct {
	s *CreateTCAclResponse
}

func NewCreateTCAclResponseBuilder() *CreateTCAclResponseBuilder {
	s := &CreateTCAclResponse{}
	b := &CreateTCAclResponseBuilder{s: s}
	return b
}

func (b *CreateTCAclResponseBuilder) RequestId(v string) *CreateTCAclResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateTCAclResponseBuilder) ErrorMessage(v string) *CreateTCAclResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateTCAclResponseBuilder) ErrorCode(v string) *CreateTCAclResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateTCAclResponseBuilder) State(v CreateTCAclResponseStateEnum) *CreateTCAclResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateTCAclResponseBuilder) Body(v string) *CreateTCAclResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateTCAclResponseBuilder) ErrorParams(v []string) *CreateTCAclResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateTCAclResponseBuilder) Build() *CreateTCAclResponse {
	return b.s
}


