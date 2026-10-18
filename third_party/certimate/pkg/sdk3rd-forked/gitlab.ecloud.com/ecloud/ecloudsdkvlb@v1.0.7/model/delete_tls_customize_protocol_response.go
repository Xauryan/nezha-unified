// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteTLSCustomizeProtocolResponseStateEnum string

// List of State
const (
    DeleteTLSCustomizeProtocolResponseStateEnumOk DeleteTLSCustomizeProtocolResponseStateEnum = "OK"
    DeleteTLSCustomizeProtocolResponseStateEnumError DeleteTLSCustomizeProtocolResponseStateEnum = "ERROR"
    DeleteTLSCustomizeProtocolResponseStateEnumException DeleteTLSCustomizeProtocolResponseStateEnum = "EXCEPTION"
    DeleteTLSCustomizeProtocolResponseStateEnumAlarm DeleteTLSCustomizeProtocolResponseStateEnum = "ALARM"
    DeleteTLSCustomizeProtocolResponseStateEnumForbidden DeleteTLSCustomizeProtocolResponseStateEnum = "FORBIDDEN"
)

type DeleteTLSCustomizeProtocolResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *DeleteTLSCustomizeProtocolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteTLSCustomizeProtocolResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteTLSCustomizeProtocolResponse) GoString() string {
	return s.String()
}

func (s DeleteTLSCustomizeProtocolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteTLSCustomizeProtocolResponse) SetAsyncKey(v string) *DeleteTLSCustomizeProtocolResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetRequestId(v string) *DeleteTLSCustomizeProtocolResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetErrorMessage(v string) *DeleteTLSCustomizeProtocolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetErrorCode(v string) *DeleteTLSCustomizeProtocolResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetState(v DeleteTLSCustomizeProtocolResponseStateEnum) *DeleteTLSCustomizeProtocolResponse {
	s.State = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetBody(v bool) *DeleteTLSCustomizeProtocolResponse {
	s.Body = &v
	return s
}

func (s *DeleteTLSCustomizeProtocolResponse) SetErrorParams(v []string) *DeleteTLSCustomizeProtocolResponse {
	s.ErrorParams = v
	return s
}


type DeleteTLSCustomizeProtocolResponseBuilder struct {
	s *DeleteTLSCustomizeProtocolResponse
}

func NewDeleteTLSCustomizeProtocolResponseBuilder() *DeleteTLSCustomizeProtocolResponseBuilder {
	s := &DeleteTLSCustomizeProtocolResponse{}
	b := &DeleteTLSCustomizeProtocolResponseBuilder{s: s}
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) AsyncKey(v string) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) RequestId(v string) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) ErrorMessage(v string) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) ErrorCode(v string) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) State(v DeleteTLSCustomizeProtocolResponseStateEnum) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) Body(v bool) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) ErrorParams(v []string) *DeleteTLSCustomizeProtocolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteTLSCustomizeProtocolResponseBuilder) Build() *DeleteTLSCustomizeProtocolResponse {
	return b.s
}


