// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreateWithHttpStatusResponseStateEnum string

// List of State
const (
    AsyncCreateWithHttpStatusResponseStateEnumOk AsyncCreateWithHttpStatusResponseStateEnum = "OK"
    AsyncCreateWithHttpStatusResponseStateEnumError AsyncCreateWithHttpStatusResponseStateEnum = "ERROR"
    AsyncCreateWithHttpStatusResponseStateEnumException AsyncCreateWithHttpStatusResponseStateEnum = "EXCEPTION"
    AsyncCreateWithHttpStatusResponseStateEnumAlarm AsyncCreateWithHttpStatusResponseStateEnum = "ALARM"
    AsyncCreateWithHttpStatusResponseStateEnumForbidden AsyncCreateWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type AsyncCreateWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *AsyncCreateWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncCreateWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreateWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s AsyncCreateWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreateWithHttpStatusResponse) SetAsyncKey(v string) *AsyncCreateWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetRequestId(v string) *AsyncCreateWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetErrorMessage(v string) *AsyncCreateWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetErrorCode(v string) *AsyncCreateWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetState(v AsyncCreateWithHttpStatusResponseStateEnum) *AsyncCreateWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetBody(v string) *AsyncCreateWithHttpStatusResponse {
	s.Body = &v
	return s
}

func (s *AsyncCreateWithHttpStatusResponse) SetErrorParams(v []string) *AsyncCreateWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type AsyncCreateWithHttpStatusResponseBuilder struct {
	s *AsyncCreateWithHttpStatusResponse
}

func NewAsyncCreateWithHttpStatusResponseBuilder() *AsyncCreateWithHttpStatusResponseBuilder {
	s := &AsyncCreateWithHttpStatusResponse{}
	b := &AsyncCreateWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) AsyncKey(v string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) RequestId(v string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) ErrorMessage(v string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) ErrorCode(v string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) State(v AsyncCreateWithHttpStatusResponseStateEnum) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) Body(v string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) ErrorParams(v []string) *AsyncCreateWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncCreateWithHttpStatusResponseBuilder) Build() *AsyncCreateWithHttpStatusResponse {
	return b.s
}


