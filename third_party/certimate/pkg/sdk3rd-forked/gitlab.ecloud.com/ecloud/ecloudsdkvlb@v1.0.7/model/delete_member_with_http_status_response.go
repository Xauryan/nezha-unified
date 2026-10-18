// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteMemberWithHttpStatusResponseStateEnum string

// List of State
const (
    DeleteMemberWithHttpStatusResponseStateEnumOk DeleteMemberWithHttpStatusResponseStateEnum = "OK"
    DeleteMemberWithHttpStatusResponseStateEnumError DeleteMemberWithHttpStatusResponseStateEnum = "ERROR"
    DeleteMemberWithHttpStatusResponseStateEnumException DeleteMemberWithHttpStatusResponseStateEnum = "EXCEPTION"
    DeleteMemberWithHttpStatusResponseStateEnumAlarm DeleteMemberWithHttpStatusResponseStateEnum = "ALARM"
    DeleteMemberWithHttpStatusResponseStateEnumForbidden DeleteMemberWithHttpStatusResponseStateEnum = "FORBIDDEN"
)

type DeleteMemberWithHttpStatusResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *DeleteMemberWithHttpStatusResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteMemberWithHttpStatusResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteMemberWithHttpStatusResponse) GoString() string {
	return s.String()
}

func (s DeleteMemberWithHttpStatusResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteMemberWithHttpStatusResponse) SetAsyncKey(v string) *DeleteMemberWithHttpStatusResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetRequestId(v string) *DeleteMemberWithHttpStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetErrorMessage(v string) *DeleteMemberWithHttpStatusResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetErrorCode(v string) *DeleteMemberWithHttpStatusResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetState(v DeleteMemberWithHttpStatusResponseStateEnum) *DeleteMemberWithHttpStatusResponse {
	s.State = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetBody(v string) *DeleteMemberWithHttpStatusResponse {
	s.Body = &v
	return s
}

func (s *DeleteMemberWithHttpStatusResponse) SetErrorParams(v []string) *DeleteMemberWithHttpStatusResponse {
	s.ErrorParams = v
	return s
}


type DeleteMemberWithHttpStatusResponseBuilder struct {
	s *DeleteMemberWithHttpStatusResponse
}

func NewDeleteMemberWithHttpStatusResponseBuilder() *DeleteMemberWithHttpStatusResponseBuilder {
	s := &DeleteMemberWithHttpStatusResponse{}
	b := &DeleteMemberWithHttpStatusResponseBuilder{s: s}
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) AsyncKey(v string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) RequestId(v string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) ErrorMessage(v string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) ErrorCode(v string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) State(v DeleteMemberWithHttpStatusResponseStateEnum) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) Body(v string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) ErrorParams(v []string) *DeleteMemberWithHttpStatusResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteMemberWithHttpStatusResponseBuilder) Build() *DeleteMemberWithHttpStatusResponse {
	return b.s
}


