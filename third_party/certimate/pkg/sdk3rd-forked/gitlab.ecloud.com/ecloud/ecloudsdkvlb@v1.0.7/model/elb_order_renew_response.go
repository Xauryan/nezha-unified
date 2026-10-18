// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderRenewResponseStateEnum string

// List of State
const (
    ElbOrderRenewResponseStateEnumOk ElbOrderRenewResponseStateEnum = "OK"
    ElbOrderRenewResponseStateEnumError ElbOrderRenewResponseStateEnum = "ERROR"
    ElbOrderRenewResponseStateEnumException ElbOrderRenewResponseStateEnum = "EXCEPTION"
    ElbOrderRenewResponseStateEnumAlarm ElbOrderRenewResponseStateEnum = "ALARM"
    ElbOrderRenewResponseStateEnumForbidden ElbOrderRenewResponseStateEnum = "FORBIDDEN"
)

type ElbOrderRenewResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ElbOrderRenewResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ElbOrderRenewResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ElbOrderRenewResponse) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderRenewResponse) GoString() string {
	return s.String()
}

func (s ElbOrderRenewResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderRenewResponse) SetRequestId(v string) *ElbOrderRenewResponse {
	s.RequestId = &v
	return s
}

func (s *ElbOrderRenewResponse) SetErrorMessage(v string) *ElbOrderRenewResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ElbOrderRenewResponse) SetErrorCode(v string) *ElbOrderRenewResponse {
	s.ErrorCode = &v
	return s
}

func (s *ElbOrderRenewResponse) SetState(v ElbOrderRenewResponseStateEnum) *ElbOrderRenewResponse {
	s.State = &v
	return s
}

func (s *ElbOrderRenewResponse) SetBody(v *ElbOrderRenewResponseBody) *ElbOrderRenewResponse {
	s.Body = v
	return s
}

func (s *ElbOrderRenewResponse) SetErrorParams(v []string) *ElbOrderRenewResponse {
	s.ErrorParams = v
	return s
}


type ElbOrderRenewResponseBuilder struct {
	s *ElbOrderRenewResponse
}

func NewElbOrderRenewResponseBuilder() *ElbOrderRenewResponseBuilder {
	s := &ElbOrderRenewResponse{}
	b := &ElbOrderRenewResponseBuilder{s: s}
	return b
}

func (b *ElbOrderRenewResponseBuilder) RequestId(v string) *ElbOrderRenewResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ElbOrderRenewResponseBuilder) ErrorMessage(v string) *ElbOrderRenewResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ElbOrderRenewResponseBuilder) ErrorCode(v string) *ElbOrderRenewResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ElbOrderRenewResponseBuilder) State(v ElbOrderRenewResponseStateEnum) *ElbOrderRenewResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ElbOrderRenewResponseBuilder) Body(v *ElbOrderRenewResponseBody) *ElbOrderRenewResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ElbOrderRenewResponseBuilder) ErrorParams(v []string) *ElbOrderRenewResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ElbOrderRenewResponseBuilder) Build() *ElbOrderRenewResponse {
	return b.s
}


