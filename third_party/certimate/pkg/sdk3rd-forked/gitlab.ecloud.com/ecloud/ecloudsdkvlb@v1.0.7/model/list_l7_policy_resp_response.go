// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListL7PolicyRespResponseStateEnum string

// List of State
const (
    ListL7PolicyRespResponseStateEnumOk ListL7PolicyRespResponseStateEnum = "OK"
    ListL7PolicyRespResponseStateEnumError ListL7PolicyRespResponseStateEnum = "ERROR"
    ListL7PolicyRespResponseStateEnumException ListL7PolicyRespResponseStateEnum = "EXCEPTION"
    ListL7PolicyRespResponseStateEnumAlarm ListL7PolicyRespResponseStateEnum = "ALARM"
    ListL7PolicyRespResponseStateEnumForbidden ListL7PolicyRespResponseStateEnum = "FORBIDDEN"
)

type ListL7PolicyRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *ListL7PolicyRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListL7PolicyRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListL7PolicyRespResponse) String() string {
	return utils.Beautify(s)
}

func (s ListL7PolicyRespResponse) GoString() string {
	return s.String()
}

func (s ListL7PolicyRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListL7PolicyRespResponse) SetAsyncKey(v string) *ListL7PolicyRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListL7PolicyRespResponse) SetRequestId(v string) *ListL7PolicyRespResponse {
	s.RequestId = &v
	return s
}

func (s *ListL7PolicyRespResponse) SetErrorMessage(v string) *ListL7PolicyRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListL7PolicyRespResponse) SetErrorCode(v string) *ListL7PolicyRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListL7PolicyRespResponse) SetState(v ListL7PolicyRespResponseStateEnum) *ListL7PolicyRespResponse {
	s.State = &v
	return s
}

func (s *ListL7PolicyRespResponse) SetBody(v *ListL7PolicyRespResponseBody) *ListL7PolicyRespResponse {
	s.Body = v
	return s
}

func (s *ListL7PolicyRespResponse) SetErrorParams(v []string) *ListL7PolicyRespResponse {
	s.ErrorParams = v
	return s
}


type ListL7PolicyRespResponseBuilder struct {
	s *ListL7PolicyRespResponse
}

func NewListL7PolicyRespResponseBuilder() *ListL7PolicyRespResponseBuilder {
	s := &ListL7PolicyRespResponse{}
	b := &ListL7PolicyRespResponseBuilder{s: s}
	return b
}

func (b *ListL7PolicyRespResponseBuilder) AsyncKey(v string) *ListL7PolicyRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) RequestId(v string) *ListL7PolicyRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) ErrorMessage(v string) *ListL7PolicyRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) ErrorCode(v string) *ListL7PolicyRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) State(v ListL7PolicyRespResponseStateEnum) *ListL7PolicyRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) Body(v *ListL7PolicyRespResponseBody) *ListL7PolicyRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) ErrorParams(v []string) *ListL7PolicyRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListL7PolicyRespResponseBuilder) Build() *ListL7PolicyRespResponse {
	return b.s
}


