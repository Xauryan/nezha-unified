// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderCreateAsyncResponseStateEnum string

// List of State
const (
    ElbOrderCreateAsyncResponseStateEnumOk ElbOrderCreateAsyncResponseStateEnum = "OK"
    ElbOrderCreateAsyncResponseStateEnumError ElbOrderCreateAsyncResponseStateEnum = "ERROR"
    ElbOrderCreateAsyncResponseStateEnumException ElbOrderCreateAsyncResponseStateEnum = "EXCEPTION"
    ElbOrderCreateAsyncResponseStateEnumAlarm ElbOrderCreateAsyncResponseStateEnum = "ALARM"
    ElbOrderCreateAsyncResponseStateEnumForbidden ElbOrderCreateAsyncResponseStateEnum = "FORBIDDEN"
)

type ElbOrderCreateAsyncResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ElbOrderCreateAsyncResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ElbOrderCreateAsyncResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ElbOrderCreateAsyncResponse) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncResponse) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncResponse) SetAsyncKey(v string) *ElbOrderCreateAsyncResponse {
	s.AsyncKey = &v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetRequestId(v string) *ElbOrderCreateAsyncResponse {
	s.RequestId = &v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetErrorMessage(v string) *ElbOrderCreateAsyncResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetErrorCode(v string) *ElbOrderCreateAsyncResponse {
	s.ErrorCode = &v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetState(v ElbOrderCreateAsyncResponseStateEnum) *ElbOrderCreateAsyncResponse {
	s.State = &v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetBody(v *ElbOrderCreateAsyncResponseBody) *ElbOrderCreateAsyncResponse {
	s.Body = v
	return s
}

func (s *ElbOrderCreateAsyncResponse) SetErrorParams(v []string) *ElbOrderCreateAsyncResponse {
	s.ErrorParams = v
	return s
}


type ElbOrderCreateAsyncResponseBuilder struct {
	s *ElbOrderCreateAsyncResponse
}

func NewElbOrderCreateAsyncResponseBuilder() *ElbOrderCreateAsyncResponseBuilder {
	s := &ElbOrderCreateAsyncResponse{}
	b := &ElbOrderCreateAsyncResponseBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) AsyncKey(v string) *ElbOrderCreateAsyncResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) RequestId(v string) *ElbOrderCreateAsyncResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) ErrorMessage(v string) *ElbOrderCreateAsyncResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) ErrorCode(v string) *ElbOrderCreateAsyncResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) State(v ElbOrderCreateAsyncResponseStateEnum) *ElbOrderCreateAsyncResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) Body(v *ElbOrderCreateAsyncResponseBody) *ElbOrderCreateAsyncResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) ErrorParams(v []string) *ElbOrderCreateAsyncResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ElbOrderCreateAsyncResponseBuilder) Build() *ElbOrderCreateAsyncResponse {
	return b.s
}


