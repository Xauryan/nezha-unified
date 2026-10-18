// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetL7PolicyRespResponseStateEnum string

// List of State
const (
    GetL7PolicyRespResponseStateEnumOk GetL7PolicyRespResponseStateEnum = "OK"
    GetL7PolicyRespResponseStateEnumError GetL7PolicyRespResponseStateEnum = "ERROR"
    GetL7PolicyRespResponseStateEnumException GetL7PolicyRespResponseStateEnum = "EXCEPTION"
    GetL7PolicyRespResponseStateEnumAlarm GetL7PolicyRespResponseStateEnum = "ALARM"
    GetL7PolicyRespResponseStateEnumForbidden GetL7PolicyRespResponseStateEnum = "FORBIDDEN"
)

type GetL7PolicyRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *GetL7PolicyRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetL7PolicyRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetL7PolicyRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetL7PolicyRespResponse) GoString() string {
	return s.String()
}

func (s GetL7PolicyRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetL7PolicyRespResponse) SetAsyncKey(v string) *GetL7PolicyRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetL7PolicyRespResponse) SetRequestId(v string) *GetL7PolicyRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetL7PolicyRespResponse) SetErrorMessage(v string) *GetL7PolicyRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetL7PolicyRespResponse) SetErrorCode(v string) *GetL7PolicyRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetL7PolicyRespResponse) SetState(v GetL7PolicyRespResponseStateEnum) *GetL7PolicyRespResponse {
	s.State = &v
	return s
}

func (s *GetL7PolicyRespResponse) SetBody(v *GetL7PolicyRespResponseBody) *GetL7PolicyRespResponse {
	s.Body = v
	return s
}

func (s *GetL7PolicyRespResponse) SetErrorParams(v []string) *GetL7PolicyRespResponse {
	s.ErrorParams = v
	return s
}


type GetL7PolicyRespResponseBuilder struct {
	s *GetL7PolicyRespResponse
}

func NewGetL7PolicyRespResponseBuilder() *GetL7PolicyRespResponseBuilder {
	s := &GetL7PolicyRespResponse{}
	b := &GetL7PolicyRespResponseBuilder{s: s}
	return b
}

func (b *GetL7PolicyRespResponseBuilder) AsyncKey(v string) *GetL7PolicyRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) RequestId(v string) *GetL7PolicyRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) ErrorMessage(v string) *GetL7PolicyRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) ErrorCode(v string) *GetL7PolicyRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) State(v GetL7PolicyRespResponseStateEnum) *GetL7PolicyRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) Body(v *GetL7PolicyRespResponseBody) *GetL7PolicyRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) ErrorParams(v []string) *GetL7PolicyRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetL7PolicyRespResponseBuilder) Build() *GetL7PolicyRespResponse {
	return b.s
}


