// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteListenerAsyncResponseStateEnum string

// List of State
const (
    DeleteListenerAsyncResponseStateEnumOk DeleteListenerAsyncResponseStateEnum = "OK"
    DeleteListenerAsyncResponseStateEnumError DeleteListenerAsyncResponseStateEnum = "ERROR"
    DeleteListenerAsyncResponseStateEnumException DeleteListenerAsyncResponseStateEnum = "EXCEPTION"
    DeleteListenerAsyncResponseStateEnumAlarm DeleteListenerAsyncResponseStateEnum = "ALARM"
    DeleteListenerAsyncResponseStateEnumForbidden DeleteListenerAsyncResponseStateEnum = "FORBIDDEN"
)

type DeleteListenerAsyncResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *DeleteListenerAsyncResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteListenerAsyncResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerAsyncResponse) GoString() string {
	return s.String()
}

func (s DeleteListenerAsyncResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerAsyncResponse) SetAsyncKey(v string) *DeleteListenerAsyncResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetRequestId(v string) *DeleteListenerAsyncResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetErrorMessage(v string) *DeleteListenerAsyncResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetErrorCode(v string) *DeleteListenerAsyncResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetState(v DeleteListenerAsyncResponseStateEnum) *DeleteListenerAsyncResponse {
	s.State = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetBody(v string) *DeleteListenerAsyncResponse {
	s.Body = &v
	return s
}

func (s *DeleteListenerAsyncResponse) SetErrorParams(v []string) *DeleteListenerAsyncResponse {
	s.ErrorParams = v
	return s
}


type DeleteListenerAsyncResponseBuilder struct {
	s *DeleteListenerAsyncResponse
}

func NewDeleteListenerAsyncResponseBuilder() *DeleteListenerAsyncResponseBuilder {
	s := &DeleteListenerAsyncResponse{}
	b := &DeleteListenerAsyncResponseBuilder{s: s}
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) AsyncKey(v string) *DeleteListenerAsyncResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) RequestId(v string) *DeleteListenerAsyncResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) ErrorMessage(v string) *DeleteListenerAsyncResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) ErrorCode(v string) *DeleteListenerAsyncResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) State(v DeleteListenerAsyncResponseStateEnum) *DeleteListenerAsyncResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) Body(v string) *DeleteListenerAsyncResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) ErrorParams(v []string) *DeleteListenerAsyncResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteListenerAsyncResponseBuilder) Build() *DeleteListenerAsyncResponse {
	return b.s
}


