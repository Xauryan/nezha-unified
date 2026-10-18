// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteListenerResponseStateEnum string

// List of State
const (
    DeleteListenerResponseStateEnumOk DeleteListenerResponseStateEnum = "OK"
    DeleteListenerResponseStateEnumError DeleteListenerResponseStateEnum = "ERROR"
    DeleteListenerResponseStateEnumException DeleteListenerResponseStateEnum = "EXCEPTION"
    DeleteListenerResponseStateEnumAlarm DeleteListenerResponseStateEnum = "ALARM"
    DeleteListenerResponseStateEnumForbidden DeleteListenerResponseStateEnum = "FORBIDDEN"
)

type DeleteListenerResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *DeleteListenerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteListenerResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s DeleteListenerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerResponse) SetAsyncKey(v string) *DeleteListenerResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteListenerResponse) SetRequestId(v string) *DeleteListenerResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteListenerResponse) SetErrorMessage(v string) *DeleteListenerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteListenerResponse) SetErrorCode(v string) *DeleteListenerResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteListenerResponse) SetState(v DeleteListenerResponseStateEnum) *DeleteListenerResponse {
	s.State = &v
	return s
}

func (s *DeleteListenerResponse) SetBody(v string) *DeleteListenerResponse {
	s.Body = &v
	return s
}

func (s *DeleteListenerResponse) SetErrorParams(v []string) *DeleteListenerResponse {
	s.ErrorParams = v
	return s
}


type DeleteListenerResponseBuilder struct {
	s *DeleteListenerResponse
}

func NewDeleteListenerResponseBuilder() *DeleteListenerResponseBuilder {
	s := &DeleteListenerResponse{}
	b := &DeleteListenerResponseBuilder{s: s}
	return b
}

func (b *DeleteListenerResponseBuilder) AsyncKey(v string) *DeleteListenerResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteListenerResponseBuilder) RequestId(v string) *DeleteListenerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteListenerResponseBuilder) ErrorMessage(v string) *DeleteListenerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteListenerResponseBuilder) ErrorCode(v string) *DeleteListenerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteListenerResponseBuilder) State(v DeleteListenerResponseStateEnum) *DeleteListenerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteListenerResponseBuilder) Body(v string) *DeleteListenerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteListenerResponseBuilder) ErrorParams(v []string) *DeleteListenerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteListenerResponseBuilder) Build() *DeleteListenerResponse {
	return b.s
}


