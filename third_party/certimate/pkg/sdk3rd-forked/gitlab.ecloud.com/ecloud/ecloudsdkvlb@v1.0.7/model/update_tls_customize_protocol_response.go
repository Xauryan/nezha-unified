// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateTLSCustomizeProtocolResponseStateEnum string

// List of State
const (
    UpdateTLSCustomizeProtocolResponseStateEnumOk UpdateTLSCustomizeProtocolResponseStateEnum = "OK"
    UpdateTLSCustomizeProtocolResponseStateEnumError UpdateTLSCustomizeProtocolResponseStateEnum = "ERROR"
    UpdateTLSCustomizeProtocolResponseStateEnumException UpdateTLSCustomizeProtocolResponseStateEnum = "EXCEPTION"
    UpdateTLSCustomizeProtocolResponseStateEnumAlarm UpdateTLSCustomizeProtocolResponseStateEnum = "ALARM"
    UpdateTLSCustomizeProtocolResponseStateEnumForbidden UpdateTLSCustomizeProtocolResponseStateEnum = "FORBIDDEN"
)

type UpdateTLSCustomizeProtocolResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *UpdateTLSCustomizeProtocolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateTLSCustomizeProtocolResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSCustomizeProtocolResponse) GoString() string {
	return s.String()
}

func (s UpdateTLSCustomizeProtocolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSCustomizeProtocolResponse) SetAsyncKey(v string) *UpdateTLSCustomizeProtocolResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetRequestId(v string) *UpdateTLSCustomizeProtocolResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetErrorMessage(v string) *UpdateTLSCustomizeProtocolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetErrorCode(v string) *UpdateTLSCustomizeProtocolResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetState(v UpdateTLSCustomizeProtocolResponseStateEnum) *UpdateTLSCustomizeProtocolResponse {
	s.State = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetBody(v bool) *UpdateTLSCustomizeProtocolResponse {
	s.Body = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolResponse) SetErrorParams(v []string) *UpdateTLSCustomizeProtocolResponse {
	s.ErrorParams = v
	return s
}


type UpdateTLSCustomizeProtocolResponseBuilder struct {
	s *UpdateTLSCustomizeProtocolResponse
}

func NewUpdateTLSCustomizeProtocolResponseBuilder() *UpdateTLSCustomizeProtocolResponseBuilder {
	s := &UpdateTLSCustomizeProtocolResponse{}
	b := &UpdateTLSCustomizeProtocolResponseBuilder{s: s}
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) AsyncKey(v string) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) RequestId(v string) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) ErrorMessage(v string) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) ErrorCode(v string) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) State(v UpdateTLSCustomizeProtocolResponseStateEnum) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) Body(v bool) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) ErrorParams(v []string) *UpdateTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateTLSCustomizeProtocolResponseBuilder) Build() *UpdateTLSCustomizeProtocolResponse {
	return b.s
}


