// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncUpdateL7PolicyResponseStateEnum string

// List of State
const (
    AsyncUpdateL7PolicyResponseStateEnumOk AsyncUpdateL7PolicyResponseStateEnum = "OK"
    AsyncUpdateL7PolicyResponseStateEnumError AsyncUpdateL7PolicyResponseStateEnum = "ERROR"
    AsyncUpdateL7PolicyResponseStateEnumException AsyncUpdateL7PolicyResponseStateEnum = "EXCEPTION"
    AsyncUpdateL7PolicyResponseStateEnumAlarm AsyncUpdateL7PolicyResponseStateEnum = "ALARM"
    AsyncUpdateL7PolicyResponseStateEnumForbidden AsyncUpdateL7PolicyResponseStateEnum = "FORBIDDEN"
)

type AsyncUpdateL7PolicyResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *AsyncUpdateL7PolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncUpdateL7PolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateL7PolicyResponse) GoString() string {
	return s.String()
}

func (s AsyncUpdateL7PolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateL7PolicyResponse) SetAsyncKey(v string) *AsyncUpdateL7PolicyResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetRequestId(v string) *AsyncUpdateL7PolicyResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetErrorMessage(v string) *AsyncUpdateL7PolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetErrorCode(v string) *AsyncUpdateL7PolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetState(v AsyncUpdateL7PolicyResponseStateEnum) *AsyncUpdateL7PolicyResponse {
	s.State = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetBody(v string) *AsyncUpdateL7PolicyResponse {
	s.Body = &v
	return s
}

func (s *AsyncUpdateL7PolicyResponse) SetErrorParams(v []string) *AsyncUpdateL7PolicyResponse {
	s.ErrorParams = v
	return s
}


type AsyncUpdateL7PolicyResponseBuilder struct {
	s *AsyncUpdateL7PolicyResponse
}

func NewAsyncUpdateL7PolicyResponseBuilder() *AsyncUpdateL7PolicyResponseBuilder {
	s := &AsyncUpdateL7PolicyResponse{}
	b := &AsyncUpdateL7PolicyResponseBuilder{s: s}
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) AsyncKey(v string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) RequestId(v string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) ErrorMessage(v string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) ErrorCode(v string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) State(v AsyncUpdateL7PolicyResponseStateEnum) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) Body(v string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) ErrorParams(v []string) *AsyncUpdateL7PolicyResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncUpdateL7PolicyResponseBuilder) Build() *AsyncUpdateL7PolicyResponse {
	return b.s
}


