// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetHealthStatusRespResponseStateEnum string

// List of State
const (
    GetHealthStatusRespResponseStateEnumOk GetHealthStatusRespResponseStateEnum = "OK"
    GetHealthStatusRespResponseStateEnumError GetHealthStatusRespResponseStateEnum = "ERROR"
    GetHealthStatusRespResponseStateEnumException GetHealthStatusRespResponseStateEnum = "EXCEPTION"
    GetHealthStatusRespResponseStateEnumAlarm GetHealthStatusRespResponseStateEnum = "ALARM"
    GetHealthStatusRespResponseStateEnumForbidden GetHealthStatusRespResponseStateEnum = "FORBIDDEN"
)

type GetHealthStatusRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *GetHealthStatusRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetHealthStatusRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetHealthStatusRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetHealthStatusRespResponse) GoString() string {
	return s.String()
}

func (s GetHealthStatusRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetHealthStatusRespResponse) SetAsyncKey(v string) *GetHealthStatusRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetHealthStatusRespResponse) SetRequestId(v string) *GetHealthStatusRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetHealthStatusRespResponse) SetErrorMessage(v string) *GetHealthStatusRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetHealthStatusRespResponse) SetErrorCode(v string) *GetHealthStatusRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetHealthStatusRespResponse) SetState(v GetHealthStatusRespResponseStateEnum) *GetHealthStatusRespResponse {
	s.State = &v
	return s
}

func (s *GetHealthStatusRespResponse) SetBody(v *GetHealthStatusRespResponseBody) *GetHealthStatusRespResponse {
	s.Body = v
	return s
}

func (s *GetHealthStatusRespResponse) SetErrorParams(v []string) *GetHealthStatusRespResponse {
	s.ErrorParams = v
	return s
}


type GetHealthStatusRespResponseBuilder struct {
	s *GetHealthStatusRespResponse
}

func NewGetHealthStatusRespResponseBuilder() *GetHealthStatusRespResponseBuilder {
	s := &GetHealthStatusRespResponse{}
	b := &GetHealthStatusRespResponseBuilder{s: s}
	return b
}

func (b *GetHealthStatusRespResponseBuilder) AsyncKey(v string) *GetHealthStatusRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) RequestId(v string) *GetHealthStatusRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) ErrorMessage(v string) *GetHealthStatusRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) ErrorCode(v string) *GetHealthStatusRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) State(v GetHealthStatusRespResponseStateEnum) *GetHealthStatusRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) Body(v *GetHealthStatusRespResponseBody) *GetHealthStatusRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) ErrorParams(v []string) *GetHealthStatusRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetHealthStatusRespResponseBuilder) Build() *GetHealthStatusRespResponse {
	return b.s
}


