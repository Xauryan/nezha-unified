// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncUpdateWithHttpStatusResponseStateEnum string

// List of State
const (
    AsyncUpdateWithHttpStatusResponseStateEnumOk AsyncUpdateWithHttpStatusResponseStateEnum = "OK"
    AsyncUpdateWithHttpStatusResponseStateEnumError AsyncUpdateWithHttpStatusResponseStateEnum = "ERROR"
    AsyncUpdateWithHttpStatusResponseStateEnumException AsyncUpdateWithHttpStatusResponseStateEnum = "EXCEPTION"
    AsyncUpdateWithHttpStatusResponseStateEnumAlarm AsyncUpdateWithHttpStatusResponseStateEnum = "ALARM"
    AsyncUpdateWithHttpStatusResponseStateEnumForbidden AsyncUpdateWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type AsyncUpdateWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *AsyncUpdateWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncUpdateWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s AsyncUpdateWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateWithHttpStatusResponse) SetAsyncKey(v string) *AsyncUpdateWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetRequestId(v string) *AsyncUpdateWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetErrorMessage(v string) *AsyncUpdateWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetErrorCode(v string) *AsyncUpdateWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetState(v AsyncUpdateWithHttpStatusResponseStateEnum) *AsyncUpdateWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetBody(v string) *AsyncUpdateWithHttpStatusResponse {
	s.Body = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusResponse) SetErrorParams(v []string) *AsyncUpdateWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type AsyncUpdateWithHttpStatusResponseBuilder struct {
	s *AsyncUpdateWithHttpStatusResponse
}

func NewAsyncUpdateWithHttpStatusResponseBuilder() *AsyncUpdateWithHttpStatusResponseBuilder {
	s := &AsyncUpdateWithHttpStatusResponse{}
	b := &AsyncUpdateWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) AsyncKey(v string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) RequestId(v string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) ErrorMessage(v string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) ErrorCode(v string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) State(v AsyncUpdateWithHttpStatusResponseStateEnum) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) Body(v string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) ErrorParams(v []string) *AsyncUpdateWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncUpdateWithHttpStatusResponseBuilder) Build() *AsyncUpdateWithHttpStatusResponse {
	return b.s
}


