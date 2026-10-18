// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListTLSCustomizeProtocolRespResponseStateEnum string

// List of State
const (
    ListTLSCustomizeProtocolRespResponseStateEnumOk ListTLSCustomizeProtocolRespResponseStateEnum = "OK"
    ListTLSCustomizeProtocolRespResponseStateEnumError ListTLSCustomizeProtocolRespResponseStateEnum = "ERROR"
    ListTLSCustomizeProtocolRespResponseStateEnumException ListTLSCustomizeProtocolRespResponseStateEnum = "EXCEPTION"
    ListTLSCustomizeProtocolRespResponseStateEnumAlarm ListTLSCustomizeProtocolRespResponseStateEnum = "ALARM"
    ListTLSCustomizeProtocolRespResponseStateEnumForbidden ListTLSCustomizeProtocolRespResponseStateEnum = "FORBIDDEN"
)

type ListTLSCustomizeProtocolRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *ListTLSCustomizeProtocolRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListTLSCustomizeProtocolRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListTLSCustomizeProtocolRespResponse) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespResponse) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespResponse) SetAsyncKey(v string) *ListTLSCustomizeProtocolRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetRequestId(v string) *ListTLSCustomizeProtocolRespResponse {
	s.RequestId = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetErrorMessage(v string) *ListTLSCustomizeProtocolRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetErrorCode(v string) *ListTLSCustomizeProtocolRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetState(v ListTLSCustomizeProtocolRespResponseStateEnum) *ListTLSCustomizeProtocolRespResponse {
	s.State = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetBody(v *ListTLSCustomizeProtocolRespResponseBody) *ListTLSCustomizeProtocolRespResponse {
	s.Body = v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponse) SetErrorParams(v []string) *ListTLSCustomizeProtocolRespResponse {
	s.ErrorParams = v
	return s
}


type ListTLSCustomizeProtocolRespResponseBuilder struct {
	s *ListTLSCustomizeProtocolRespResponse
}

func NewListTLSCustomizeProtocolRespResponseBuilder() *ListTLSCustomizeProtocolRespResponseBuilder {
	s := &ListTLSCustomizeProtocolRespResponse{}
	b := &ListTLSCustomizeProtocolRespResponseBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) AsyncKey(v string) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) RequestId(v string) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) ErrorMessage(v string) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) ErrorCode(v string) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) State(v ListTLSCustomizeProtocolRespResponseStateEnum) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) Body(v *ListTLSCustomizeProtocolRespResponseBody) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) ErrorParams(v []string) *ListTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBuilder) Build() *ListTLSCustomizeProtocolRespResponse {
	return b.s
}


