// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdatePoolNameResponseStateEnum string

// List of State
const (
    UpdatePoolNameResponseStateEnumOk UpdatePoolNameResponseStateEnum = "OK"
    UpdatePoolNameResponseStateEnumError UpdatePoolNameResponseStateEnum = "ERROR"
    UpdatePoolNameResponseStateEnumException UpdatePoolNameResponseStateEnum = "EXCEPTION"
    UpdatePoolNameResponseStateEnumAlarm UpdatePoolNameResponseStateEnum = "ALARM"
    UpdatePoolNameResponseStateEnumForbidden UpdatePoolNameResponseStateEnum = "FORBIDDEN"
)

type UpdatePoolNameResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *UpdatePoolNameResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdatePoolNameResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdatePoolNameResponse) GoString() string {
	return s.String()
}

func (s UpdatePoolNameResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePoolNameResponse) SetAsyncKey(v string) *UpdatePoolNameResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdatePoolNameResponse) SetRequestId(v string) *UpdatePoolNameResponse {
	s.RequestId = &v
	return s
}

func (s *UpdatePoolNameResponse) SetErrorMessage(v string) *UpdatePoolNameResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePoolNameResponse) SetErrorCode(v string) *UpdatePoolNameResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePoolNameResponse) SetState(v UpdatePoolNameResponseStateEnum) *UpdatePoolNameResponse {
	s.State = &v
	return s
}

func (s *UpdatePoolNameResponse) SetBody(v string) *UpdatePoolNameResponse {
	s.Body = &v
	return s
}

func (s *UpdatePoolNameResponse) SetErrorParams(v []string) *UpdatePoolNameResponse {
	s.ErrorParams = v
	return s
}


type UpdatePoolNameResponseBuilder struct {
	s *UpdatePoolNameResponse
}

func NewUpdatePoolNameResponseBuilder() *UpdatePoolNameResponseBuilder {
	s := &UpdatePoolNameResponse{}
	b := &UpdatePoolNameResponseBuilder{s: s}
	return b
}

func (b *UpdatePoolNameResponseBuilder) AsyncKey(v string) *UpdatePoolNameResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) RequestId(v string) *UpdatePoolNameResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) ErrorMessage(v string) *UpdatePoolNameResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) ErrorCode(v string) *UpdatePoolNameResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) State(v UpdatePoolNameResponseStateEnum) *UpdatePoolNameResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) Body(v string) *UpdatePoolNameResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdatePoolNameResponseBuilder) ErrorParams(v []string) *UpdatePoolNameResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdatePoolNameResponseBuilder) Build() *UpdatePoolNameResponse {
	return b.s
}


