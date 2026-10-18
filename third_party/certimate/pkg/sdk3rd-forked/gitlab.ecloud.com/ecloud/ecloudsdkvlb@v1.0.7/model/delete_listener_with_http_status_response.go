// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteListenerWithHttpStatusResponseStateEnum string

// List of State
const (
    DeleteListenerWithHttpStatusResponseStateEnumOk DeleteListenerWithHttpStatusResponseStateEnum = "OK"
    DeleteListenerWithHttpStatusResponseStateEnumError DeleteListenerWithHttpStatusResponseStateEnum = "ERROR"
    DeleteListenerWithHttpStatusResponseStateEnumException DeleteListenerWithHttpStatusResponseStateEnum = "EXCEPTION"
    DeleteListenerWithHttpStatusResponseStateEnumAlarm DeleteListenerWithHttpStatusResponseStateEnum = "ALARM"
    DeleteListenerWithHttpStatusResponseStateEnumForbidden DeleteListenerWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type DeleteListenerWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *DeleteListenerWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteListenerWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s DeleteListenerWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerWithHttpStatusResponse) SetAsyncKey(v string) *DeleteListenerWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetRequestId(v string) *DeleteListenerWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetErrorMessage(v string) *DeleteListenerWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetErrorCode(v string) *DeleteListenerWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetState(v DeleteListenerWithHttpStatusResponseStateEnum) *DeleteListenerWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetBody(v string) *DeleteListenerWithHttpStatusResponse {
	s.Body = &v
	return s
}

func (s *DeleteListenerWithHttpStatusResponse) SetErrorParams(v []string) *DeleteListenerWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type DeleteListenerWithHttpStatusResponseBuilder struct {
	s *DeleteListenerWithHttpStatusResponse
}

func NewDeleteListenerWithHttpStatusResponseBuilder() *DeleteListenerWithHttpStatusResponseBuilder {
	s := &DeleteListenerWithHttpStatusResponse{}
	b := &DeleteListenerWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) AsyncKey(v string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) RequestId(v string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) ErrorMessage(v string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) ErrorCode(v string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) State(v DeleteListenerWithHttpStatusResponseStateEnum) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) Body(v string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) ErrorParams(v []string) *DeleteListenerWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteListenerWithHttpStatusResponseBuilder) Build() *DeleteListenerWithHttpStatusResponse {
	return b.s
}


