// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetTLSCustomizeProtocolRespResponseStateEnum string

// List of State
const (
    GetTLSCustomizeProtocolRespResponseStateEnumOk GetTLSCustomizeProtocolRespResponseStateEnum = "OK"
    GetTLSCustomizeProtocolRespResponseStateEnumError GetTLSCustomizeProtocolRespResponseStateEnum = "ERROR"
    GetTLSCustomizeProtocolRespResponseStateEnumException GetTLSCustomizeProtocolRespResponseStateEnum = "EXCEPTION"
    GetTLSCustomizeProtocolRespResponseStateEnumAlarm GetTLSCustomizeProtocolRespResponseStateEnum = "ALARM"
    GetTLSCustomizeProtocolRespResponseStateEnumForbidden GetTLSCustomizeProtocolRespResponseStateEnum = "FORBIDDEN"
)

type GetTLSCustomizeProtocolRespResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *GetTLSCustomizeProtocolRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetTLSCustomizeProtocolRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetTLSCustomizeProtocolRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetTLSCustomizeProtocolRespResponse) GoString() string {
	return s.String()
}

func (s GetTLSCustomizeProtocolRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTLSCustomizeProtocolRespResponse) SetRequestId(v string) *GetTLSCustomizeProtocolRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponse) SetErrorMessage(v string) *GetTLSCustomizeProtocolRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponse) SetErrorCode(v string) *GetTLSCustomizeProtocolRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponse) SetState(v GetTLSCustomizeProtocolRespResponseStateEnum) *GetTLSCustomizeProtocolRespResponse {
	s.State = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponse) SetBody(v *GetTLSCustomizeProtocolRespResponseBody) *GetTLSCustomizeProtocolRespResponse {
	s.Body = v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponse) SetErrorParams(v []string) *GetTLSCustomizeProtocolRespResponse {
	s.ErrorParams = v
	return s
}


type GetTLSCustomizeProtocolRespResponseBuilder struct {
	s *GetTLSCustomizeProtocolRespResponse
}

func NewGetTLSCustomizeProtocolRespResponseBuilder() *GetTLSCustomizeProtocolRespResponseBuilder {
	s := &GetTLSCustomizeProtocolRespResponse{}
	b := &GetTLSCustomizeProtocolRespResponseBuilder{s: s}
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) RequestId(v string) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) ErrorMessage(v string) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) ErrorCode(v string) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) State(v GetTLSCustomizeProtocolRespResponseStateEnum) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) Body(v *GetTLSCustomizeProtocolRespResponseBody) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) ErrorParams(v []string) *GetTLSCustomizeProtocolRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBuilder) Build() *GetTLSCustomizeProtocolRespResponse {
	return b.s
}


