// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateTLSNameResponseStateEnum string

// List of State
const (
    UpdateTLSNameResponseStateEnumOk UpdateTLSNameResponseStateEnum = "OK"
    UpdateTLSNameResponseStateEnumError UpdateTLSNameResponseStateEnum = "ERROR"
    UpdateTLSNameResponseStateEnumException UpdateTLSNameResponseStateEnum = "EXCEPTION"
    UpdateTLSNameResponseStateEnumAlarm UpdateTLSNameResponseStateEnum = "ALARM"
    UpdateTLSNameResponseStateEnumForbidden UpdateTLSNameResponseStateEnum = "FORBIDDEN"
)

type UpdateTLSNameResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *UpdateTLSNameResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateTLSNameResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSNameResponse) GoString() string {
	return s.String()
}

func (s UpdateTLSNameResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSNameResponse) SetAsyncKey(v string) *UpdateTLSNameResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateTLSNameResponse) SetRequestId(v string) *UpdateTLSNameResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateTLSNameResponse) SetErrorMessage(v string) *UpdateTLSNameResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTLSNameResponse) SetErrorCode(v string) *UpdateTLSNameResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTLSNameResponse) SetState(v UpdateTLSNameResponseStateEnum) *UpdateTLSNameResponse {
	s.State = &v
	return s
}

func (s *UpdateTLSNameResponse) SetBody(v bool) *UpdateTLSNameResponse {
	s.Body = &v
	return s
}

func (s *UpdateTLSNameResponse) SetErrorParams(v []string) *UpdateTLSNameResponse {
	s.ErrorParams = v
	return s
}


type UpdateTLSNameResponseBuilder struct {
	s *UpdateTLSNameResponse
}

func NewUpdateTLSNameResponseBuilder() *UpdateTLSNameResponseBuilder {
	s := &UpdateTLSNameResponse{}
	b := &UpdateTLSNameResponseBuilder{s: s}
	return b
}

func (b *UpdateTLSNameResponseBuilder) AsyncKey(v string) *UpdateTLSNameResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) RequestId(v string) *UpdateTLSNameResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) ErrorMessage(v string) *UpdateTLSNameResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) ErrorCode(v string) *UpdateTLSNameResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) State(v UpdateTLSNameResponseStateEnum) *UpdateTLSNameResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) Body(v bool) *UpdateTLSNameResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateTLSNameResponseBuilder) ErrorParams(v []string) *UpdateTLSNameResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateTLSNameResponseBuilder) Build() *UpdateTLSNameResponse {
	return b.s
}


