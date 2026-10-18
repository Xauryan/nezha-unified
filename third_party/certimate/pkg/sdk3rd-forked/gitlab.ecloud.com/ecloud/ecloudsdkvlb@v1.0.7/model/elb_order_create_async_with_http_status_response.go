// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderCreateAsyncWithHttpStatusResponseStateEnum string

// List of State
const (
    ElbOrderCreateAsyncWithHttpStatusResponseStateEnumOk ElbOrderCreateAsyncWithHttpStatusResponseStateEnum = "OK"
    ElbOrderCreateAsyncWithHttpStatusResponseStateEnumError ElbOrderCreateAsyncWithHttpStatusResponseStateEnum = "ERROR"
    ElbOrderCreateAsyncWithHttpStatusResponseStateEnumException ElbOrderCreateAsyncWithHttpStatusResponseStateEnum = "EXCEPTION"
    ElbOrderCreateAsyncWithHttpStatusResponseStateEnumAlarm ElbOrderCreateAsyncWithHttpStatusResponseStateEnum = "ALARM"
    ElbOrderCreateAsyncWithHttpStatusResponseStateEnumForbidden ElbOrderCreateAsyncWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type ElbOrderCreateAsyncWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ElbOrderCreateAsyncWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ElbOrderCreateAsyncWithHttpStatusResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ElbOrderCreateAsyncWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetAsyncKey(v string) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetRequestId(v string) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetErrorMessage(v string) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetErrorCode(v string) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetState(v ElbOrderCreateAsyncWithHttpStatusResponseStateEnum) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetBody(v *ElbOrderCreateAsyncWithHttpStatusResponseBody) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.Body = v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponse) SetErrorParams(v []string) *ElbOrderCreateAsyncWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type ElbOrderCreateAsyncWithHttpStatusResponseBuilder struct {
	s *ElbOrderCreateAsyncWithHttpStatusResponse
}

func NewElbOrderCreateAsyncWithHttpStatusResponseBuilder() *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	s := &ElbOrderCreateAsyncWithHttpStatusResponse{}
	b := &ElbOrderCreateAsyncWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) AsyncKey(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) RequestId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) ErrorMessage(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) ErrorCode(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) State(v ElbOrderCreateAsyncWithHttpStatusResponseStateEnum) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) Body(v *ElbOrderCreateAsyncWithHttpStatusResponseBody) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) ErrorParams(v []string) *ElbOrderCreateAsyncWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBuilder) Build() *ElbOrderCreateAsyncWithHttpStatusResponse {
	return b.s
}


