// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListNewServerPortsForMemberResponseStateEnum string

// List of State
const (
    ListNewServerPortsForMemberResponseStateEnumOk ListNewServerPortsForMemberResponseStateEnum = "OK"
    ListNewServerPortsForMemberResponseStateEnumError ListNewServerPortsForMemberResponseStateEnum = "ERROR"
    ListNewServerPortsForMemberResponseStateEnumException ListNewServerPortsForMemberResponseStateEnum = "EXCEPTION"
    ListNewServerPortsForMemberResponseStateEnumAlarm ListNewServerPortsForMemberResponseStateEnum = "ALARM"
    ListNewServerPortsForMemberResponseStateEnumForbidden ListNewServerPortsForMemberResponseStateEnum = "FORBIDDEN"
)

type ListNewServerPortsForMemberResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *ListNewServerPortsForMemberResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListNewServerPortsForMemberResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListNewServerPortsForMemberResponse) String() string {
	return utils.Beautify(s)
}

func (s ListNewServerPortsForMemberResponse) GoString() string {
	return s.String()
}

func (s ListNewServerPortsForMemberResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListNewServerPortsForMemberResponse) SetAsyncKey(v string) *ListNewServerPortsForMemberResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetRequestId(v string) *ListNewServerPortsForMemberResponse {
	s.RequestId = &v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetErrorMessage(v string) *ListNewServerPortsForMemberResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetErrorCode(v string) *ListNewServerPortsForMemberResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetState(v ListNewServerPortsForMemberResponseStateEnum) *ListNewServerPortsForMemberResponse {
	s.State = &v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetBody(v *ListNewServerPortsForMemberResponseBody) *ListNewServerPortsForMemberResponse {
	s.Body = v
	return s
}

func (s *ListNewServerPortsForMemberResponse) SetErrorParams(v []string) *ListNewServerPortsForMemberResponse {
	s.ErrorParams = v
	return s
}


type ListNewServerPortsForMemberResponseBuilder struct {
	s *ListNewServerPortsForMemberResponse
}

func NewListNewServerPortsForMemberResponseBuilder() *ListNewServerPortsForMemberResponseBuilder {
	s := &ListNewServerPortsForMemberResponse{}
	b := &ListNewServerPortsForMemberResponseBuilder{s: s}
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) AsyncKey(v string) *ListNewServerPortsForMemberResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) RequestId(v string) *ListNewServerPortsForMemberResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) ErrorMessage(v string) *ListNewServerPortsForMemberResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) ErrorCode(v string) *ListNewServerPortsForMemberResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) State(v ListNewServerPortsForMemberResponseStateEnum) *ListNewServerPortsForMemberResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) Body(v *ListNewServerPortsForMemberResponseBody) *ListNewServerPortsForMemberResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) ErrorParams(v []string) *ListNewServerPortsForMemberResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListNewServerPortsForMemberResponseBuilder) Build() *ListNewServerPortsForMemberResponse {
	return b.s
}


