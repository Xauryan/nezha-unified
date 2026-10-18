// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListPoolRespResponseStateEnum string

// List of State
const (
    ListPoolRespResponseStateEnumOk ListPoolRespResponseStateEnum = "OK"
    ListPoolRespResponseStateEnumError ListPoolRespResponseStateEnum = "ERROR"
    ListPoolRespResponseStateEnumException ListPoolRespResponseStateEnum = "EXCEPTION"
    ListPoolRespResponseStateEnumAlarm ListPoolRespResponseStateEnum = "ALARM"
    ListPoolRespResponseStateEnumForbidden ListPoolRespResponseStateEnum = "FORBIDDEN"
)

type ListPoolRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *ListPoolRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListPoolRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListPoolRespResponse) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespResponse) GoString() string {
	return s.String()
}

func (s ListPoolRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespResponse) SetAsyncKey(v string) *ListPoolRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListPoolRespResponse) SetRequestId(v string) *ListPoolRespResponse {
	s.RequestId = &v
	return s
}

func (s *ListPoolRespResponse) SetErrorMessage(v string) *ListPoolRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListPoolRespResponse) SetErrorCode(v string) *ListPoolRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListPoolRespResponse) SetState(v ListPoolRespResponseStateEnum) *ListPoolRespResponse {
	s.State = &v
	return s
}

func (s *ListPoolRespResponse) SetBody(v *ListPoolRespResponseBody) *ListPoolRespResponse {
	s.Body = v
	return s
}

func (s *ListPoolRespResponse) SetErrorParams(v []string) *ListPoolRespResponse {
	s.ErrorParams = v
	return s
}


type ListPoolRespResponseBuilder struct {
	s *ListPoolRespResponse
}

func NewListPoolRespResponseBuilder() *ListPoolRespResponseBuilder {
	s := &ListPoolRespResponse{}
	b := &ListPoolRespResponseBuilder{s: s}
	return b
}

func (b *ListPoolRespResponseBuilder) AsyncKey(v string) *ListPoolRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListPoolRespResponseBuilder) RequestId(v string) *ListPoolRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListPoolRespResponseBuilder) ErrorMessage(v string) *ListPoolRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListPoolRespResponseBuilder) ErrorCode(v string) *ListPoolRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListPoolRespResponseBuilder) State(v ListPoolRespResponseStateEnum) *ListPoolRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListPoolRespResponseBuilder) Body(v *ListPoolRespResponseBody) *ListPoolRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListPoolRespResponseBuilder) ErrorParams(v []string) *ListPoolRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListPoolRespResponseBuilder) Build() *ListPoolRespResponse {
	return b.s
}


