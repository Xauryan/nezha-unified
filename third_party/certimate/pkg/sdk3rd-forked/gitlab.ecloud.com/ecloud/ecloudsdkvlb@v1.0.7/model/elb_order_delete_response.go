// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderDeleteResponseStateEnum string

// List of State
const (
    ElbOrderDeleteResponseStateEnumOk ElbOrderDeleteResponseStateEnum = "OK"
    ElbOrderDeleteResponseStateEnumError ElbOrderDeleteResponseStateEnum = "ERROR"
    ElbOrderDeleteResponseStateEnumException ElbOrderDeleteResponseStateEnum = "EXCEPTION"
    ElbOrderDeleteResponseStateEnumAlarm ElbOrderDeleteResponseStateEnum = "ALARM"
    ElbOrderDeleteResponseStateEnumForbidden ElbOrderDeleteResponseStateEnum = "FORBIDDEN"
)

type ElbOrderDeleteResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ElbOrderDeleteResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ElbOrderDeleteResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ElbOrderDeleteResponse) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteResponse) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteResponse) SetRequestId(v string) *ElbOrderDeleteResponse {
	s.RequestId = &v
	return s
}

func (s *ElbOrderDeleteResponse) SetErrorMessage(v string) *ElbOrderDeleteResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ElbOrderDeleteResponse) SetErrorCode(v string) *ElbOrderDeleteResponse {
	s.ErrorCode = &v
	return s
}

func (s *ElbOrderDeleteResponse) SetState(v ElbOrderDeleteResponseStateEnum) *ElbOrderDeleteResponse {
	s.State = &v
	return s
}

func (s *ElbOrderDeleteResponse) SetBody(v *ElbOrderDeleteResponseBody) *ElbOrderDeleteResponse {
	s.Body = v
	return s
}

func (s *ElbOrderDeleteResponse) SetErrorParams(v []string) *ElbOrderDeleteResponse {
	s.ErrorParams = v
	return s
}


type ElbOrderDeleteResponseBuilder struct {
	s *ElbOrderDeleteResponse
}

func NewElbOrderDeleteResponseBuilder() *ElbOrderDeleteResponseBuilder {
	s := &ElbOrderDeleteResponse{}
	b := &ElbOrderDeleteResponseBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteResponseBuilder) RequestId(v string) *ElbOrderDeleteResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) ErrorMessage(v string) *ElbOrderDeleteResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) ErrorCode(v string) *ElbOrderDeleteResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) State(v ElbOrderDeleteResponseStateEnum) *ElbOrderDeleteResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) Body(v *ElbOrderDeleteResponseBody) *ElbOrderDeleteResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) ErrorParams(v []string) *ElbOrderDeleteResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ElbOrderDeleteResponseBuilder) Build() *ElbOrderDeleteResponse {
	return b.s
}


