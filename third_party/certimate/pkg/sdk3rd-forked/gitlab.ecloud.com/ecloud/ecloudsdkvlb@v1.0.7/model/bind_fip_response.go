// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindFipResponseStateEnum string

// List of State
const (
    BindFipResponseStateEnumOk BindFipResponseStateEnum = "OK"
    BindFipResponseStateEnumError BindFipResponseStateEnum = "ERROR"
    BindFipResponseStateEnumException BindFipResponseStateEnum = "EXCEPTION"
    BindFipResponseStateEnumAlarm BindFipResponseStateEnum = "ALARM"
    BindFipResponseStateEnumForbidden BindFipResponseStateEnum = "FORBIDDEN"
)

type BindFipResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码
	State *BindFipResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s BindFipResponse) String() string {
	return utils.Beautify(s)
}

func (s BindFipResponse) GoString() string {
	return s.String()
}

func (s BindFipResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindFipResponse) SetAsyncKey(v string) *BindFipResponse {
	s.AsyncKey = &v
	return s
}

func (s *BindFipResponse) SetRequestId(v string) *BindFipResponse {
	s.RequestId = &v
	return s
}

func (s *BindFipResponse) SetErrorMessage(v string) *BindFipResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BindFipResponse) SetErrorCode(v string) *BindFipResponse {
	s.ErrorCode = &v
	return s
}

func (s *BindFipResponse) SetState(v BindFipResponseStateEnum) *BindFipResponse {
	s.State = &v
	return s
}

func (s *BindFipResponse) SetBody(v string) *BindFipResponse {
	s.Body = &v
	return s
}

func (s *BindFipResponse) SetErrorParams(v []string) *BindFipResponse {
	s.ErrorParams = v
	return s
}


type BindFipResponseBuilder struct {
	s *BindFipResponse
}

func NewBindFipResponseBuilder() *BindFipResponseBuilder {
	s := &BindFipResponse{}
	b := &BindFipResponseBuilder{s: s}
	return b
}

func (b *BindFipResponseBuilder) AsyncKey(v string) *BindFipResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *BindFipResponseBuilder) RequestId(v string) *BindFipResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BindFipResponseBuilder) ErrorMessage(v string) *BindFipResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BindFipResponseBuilder) ErrorCode(v string) *BindFipResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BindFipResponseBuilder) State(v BindFipResponseStateEnum) *BindFipResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BindFipResponseBuilder) Body(v string) *BindFipResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BindFipResponseBuilder) ErrorParams(v []string) *BindFipResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *BindFipResponseBuilder) Build() *BindFipResponse {
	return b.s
}


