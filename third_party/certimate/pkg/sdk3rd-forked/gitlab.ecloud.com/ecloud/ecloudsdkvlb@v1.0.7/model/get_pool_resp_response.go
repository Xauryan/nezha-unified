// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPoolRespResponseStateEnum string

// List of State
const (
    GetPoolRespResponseStateEnumOk GetPoolRespResponseStateEnum = "OK"
    GetPoolRespResponseStateEnumError GetPoolRespResponseStateEnum = "ERROR"
    GetPoolRespResponseStateEnumException GetPoolRespResponseStateEnum = "EXCEPTION"
    GetPoolRespResponseStateEnumAlarm GetPoolRespResponseStateEnum = "ALARM"
    GetPoolRespResponseStateEnumForbidden GetPoolRespResponseStateEnum = "FORBIDDEN"
)

type GetPoolRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *GetPoolRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPoolRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetPoolRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespResponse) GoString() string {
	return s.String()
}

func (s GetPoolRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespResponse) SetAsyncKey(v string) *GetPoolRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetPoolRespResponse) SetRequestId(v string) *GetPoolRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetPoolRespResponse) SetErrorMessage(v string) *GetPoolRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPoolRespResponse) SetErrorCode(v string) *GetPoolRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPoolRespResponse) SetState(v GetPoolRespResponseStateEnum) *GetPoolRespResponse {
	s.State = &v
	return s
}

func (s *GetPoolRespResponse) SetBody(v *GetPoolRespResponseBody) *GetPoolRespResponse {
	s.Body = v
	return s
}

func (s *GetPoolRespResponse) SetErrorParams(v []string) *GetPoolRespResponse {
	s.ErrorParams = v
	return s
}


type GetPoolRespResponseBuilder struct {
	s *GetPoolRespResponse
}

func NewGetPoolRespResponseBuilder() *GetPoolRespResponseBuilder {
	s := &GetPoolRespResponse{}
	b := &GetPoolRespResponseBuilder{s: s}
	return b
}

func (b *GetPoolRespResponseBuilder) AsyncKey(v string) *GetPoolRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetPoolRespResponseBuilder) RequestId(v string) *GetPoolRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPoolRespResponseBuilder) ErrorMessage(v string) *GetPoolRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPoolRespResponseBuilder) ErrorCode(v string) *GetPoolRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPoolRespResponseBuilder) State(v GetPoolRespResponseStateEnum) *GetPoolRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPoolRespResponseBuilder) Body(v *GetPoolRespResponseBody) *GetPoolRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPoolRespResponseBuilder) ErrorParams(v []string) *GetPoolRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetPoolRespResponseBuilder) Build() *GetPoolRespResponse {
	return b.s
}


