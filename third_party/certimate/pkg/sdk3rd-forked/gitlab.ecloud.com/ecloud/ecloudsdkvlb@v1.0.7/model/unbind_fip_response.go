// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UnbindFipResponseStateEnum string

// List of State
const (
    UnbindFipResponseStateEnumOk UnbindFipResponseStateEnum = "OK"
    UnbindFipResponseStateEnumError UnbindFipResponseStateEnum = "ERROR"
    UnbindFipResponseStateEnumException UnbindFipResponseStateEnum = "EXCEPTION"
    UnbindFipResponseStateEnumAlarm UnbindFipResponseStateEnum = "ALARM"
    UnbindFipResponseStateEnumForbidden UnbindFipResponseStateEnum = "FORBIDDEN"
)

type UnbindFipResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码
	State *UnbindFipResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UnbindFipResponse) String() string {
	return utils.Beautify(s)
}

func (s UnbindFipResponse) GoString() string {
	return s.String()
}

func (s UnbindFipResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindFipResponse) SetAsyncKey(v string) *UnbindFipResponse {
	s.AsyncKey = &v
	return s
}

func (s *UnbindFipResponse) SetRequestId(v string) *UnbindFipResponse {
	s.RequestId = &v
	return s
}

func (s *UnbindFipResponse) SetErrorMessage(v string) *UnbindFipResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindFipResponse) SetErrorCode(v string) *UnbindFipResponse {
	s.ErrorCode = &v
	return s
}

func (s *UnbindFipResponse) SetState(v UnbindFipResponseStateEnum) *UnbindFipResponse {
	s.State = &v
	return s
}

func (s *UnbindFipResponse) SetBody(v bool) *UnbindFipResponse {
	s.Body = &v
	return s
}

func (s *UnbindFipResponse) SetErrorParams(v []string) *UnbindFipResponse {
	s.ErrorParams = v
	return s
}


type UnbindFipResponseBuilder struct {
	s *UnbindFipResponse
}

func NewUnbindFipResponseBuilder() *UnbindFipResponseBuilder {
	s := &UnbindFipResponse{}
	b := &UnbindFipResponseBuilder{s: s}
	return b
}

func (b *UnbindFipResponseBuilder) AsyncKey(v string) *UnbindFipResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UnbindFipResponseBuilder) RequestId(v string) *UnbindFipResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnbindFipResponseBuilder) ErrorMessage(v string) *UnbindFipResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UnbindFipResponseBuilder) ErrorCode(v string) *UnbindFipResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UnbindFipResponseBuilder) State(v UnbindFipResponseStateEnum) *UnbindFipResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UnbindFipResponseBuilder) Body(v bool) *UnbindFipResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UnbindFipResponseBuilder) ErrorParams(v []string) *UnbindFipResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UnbindFipResponseBuilder) Build() *UnbindFipResponse {
	return b.s
}


