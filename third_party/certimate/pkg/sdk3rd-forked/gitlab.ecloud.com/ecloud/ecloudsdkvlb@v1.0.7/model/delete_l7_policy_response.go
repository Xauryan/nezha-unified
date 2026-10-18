// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteL7PolicyResponseStateEnum string

// List of State
const (
    DeleteL7PolicyResponseStateEnumOk DeleteL7PolicyResponseStateEnum = "OK"
    DeleteL7PolicyResponseStateEnumError DeleteL7PolicyResponseStateEnum = "ERROR"
    DeleteL7PolicyResponseStateEnumException DeleteL7PolicyResponseStateEnum = "EXCEPTION"
    DeleteL7PolicyResponseStateEnumAlarm DeleteL7PolicyResponseStateEnum = "ALARM"
    DeleteL7PolicyResponseStateEnumForbidden DeleteL7PolicyResponseStateEnum = "FORBIDDEN"
)

type DeleteL7PolicyResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *DeleteL7PolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteL7PolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteL7PolicyResponse) GoString() string {
	return s.String()
}

func (s DeleteL7PolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteL7PolicyResponse) SetAsyncKey(v string) *DeleteL7PolicyResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetRequestId(v string) *DeleteL7PolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetErrorMessage(v string) *DeleteL7PolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetErrorCode(v string) *DeleteL7PolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetState(v DeleteL7PolicyResponseStateEnum) *DeleteL7PolicyResponse {
	s.State = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetBody(v string) *DeleteL7PolicyResponse {
	s.Body = &v
	return s
}

func (s *DeleteL7PolicyResponse) SetErrorParams(v []string) *DeleteL7PolicyResponse {
	s.ErrorParams = v
	return s
}


type DeleteL7PolicyResponseBuilder struct {
	s *DeleteL7PolicyResponse
}

func NewDeleteL7PolicyResponseBuilder() *DeleteL7PolicyResponseBuilder {
	s := &DeleteL7PolicyResponse{}
	b := &DeleteL7PolicyResponseBuilder{s: s}
	return b
}

func (b *DeleteL7PolicyResponseBuilder) AsyncKey(v string) *DeleteL7PolicyResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) RequestId(v string) *DeleteL7PolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) ErrorMessage(v string) *DeleteL7PolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) ErrorCode(v string) *DeleteL7PolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) State(v DeleteL7PolicyResponseStateEnum) *DeleteL7PolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) Body(v string) *DeleteL7PolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) ErrorParams(v []string) *DeleteL7PolicyResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteL7PolicyResponseBuilder) Build() *DeleteL7PolicyResponse {
	return b.s
}


