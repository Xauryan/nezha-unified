// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateListenerResponseStateEnum string

// List of State
const (
    UpdateListenerResponseStateEnumOk UpdateListenerResponseStateEnum = "OK"
    UpdateListenerResponseStateEnumError UpdateListenerResponseStateEnum = "ERROR"
    UpdateListenerResponseStateEnumException UpdateListenerResponseStateEnum = "EXCEPTION"
    UpdateListenerResponseStateEnumAlarm UpdateListenerResponseStateEnum = "ALARM"
    UpdateListenerResponseStateEnumForbidden UpdateListenerResponseStateEnum = "FORBIDDEN"
)

type UpdateListenerResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateListenerResponse) GoString() string {
	return s.String()
}

func (s UpdateListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateListenerResponse) SetAsyncKey(v string) *UpdateListenerResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateListenerResponse) SetRequestId(v string) *UpdateListenerResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateListenerResponse) SetErrorMessage(v string) *UpdateListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateListenerResponse) SetErrorCode(v string) *UpdateListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateListenerResponse) SetState(v UpdateListenerResponseStateEnum) *UpdateListenerResponse {
	s.State = &v
	return s
}

func (s *UpdateListenerResponse) SetBody(v string) *UpdateListenerResponse {
	s.Body = &v
	return s
}

func (s *UpdateListenerResponse) SetErrorParams(v []string) *UpdateListenerResponse {
	s.ErrorParams = v
	return s
}


type UpdateListenerResponseBuilder struct {
	s *UpdateListenerResponse
}

func NewUpdateListenerResponseBuilder() *UpdateListenerResponseBuilder {
	s := &UpdateListenerResponse{}
	b := &UpdateListenerResponseBuilder{s: s}
	return b
}

func (b *UpdateListenerResponseBuilder) AsyncKey(v string) *UpdateListenerResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateListenerResponseBuilder) RequestId(v string) *UpdateListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateListenerResponseBuilder) ErrorMessage(v string) *UpdateListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateListenerResponseBuilder) ErrorCode(v string) *UpdateListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateListenerResponseBuilder) State(v UpdateListenerResponseStateEnum) *UpdateListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateListenerResponseBuilder) Body(v string) *UpdateListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateListenerResponseBuilder) ErrorParams(v []string) *UpdateListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateListenerResponseBuilder) Build() *UpdateListenerResponse {
	return b.s
}


