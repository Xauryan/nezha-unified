// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceResponseStateEnum string

// List of State
const (
    ListLoadbalanceResponseStateEnumOk ListLoadbalanceResponseStateEnum = "OK"
    ListLoadbalanceResponseStateEnumError ListLoadbalanceResponseStateEnum = "ERROR"
    ListLoadbalanceResponseStateEnumException ListLoadbalanceResponseStateEnum = "EXCEPTION"
    ListLoadbalanceResponseStateEnumAlarm ListLoadbalanceResponseStateEnum = "ALARM"
    ListLoadbalanceResponseStateEnumForbidden ListLoadbalanceResponseStateEnum = "FORBIDDEN"
)

type ListLoadbalanceResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadbalanceResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadbalanceResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadbalanceResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceResponse) GoString() string {
	return s.String()
}

func (s ListLoadbalanceResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceResponse) SetAsyncKey(v string) *ListLoadbalanceResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListLoadbalanceResponse) SetRequestId(v string) *ListLoadbalanceResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadbalanceResponse) SetErrorMessage(v string) *ListLoadbalanceResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadbalanceResponse) SetErrorCode(v string) *ListLoadbalanceResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadbalanceResponse) SetState(v ListLoadbalanceResponseStateEnum) *ListLoadbalanceResponse {
	s.State = &v
	return s
}

func (s *ListLoadbalanceResponse) SetBody(v *ListLoadbalanceResponseBody) *ListLoadbalanceResponse {
	s.Body = v
	return s
}

func (s *ListLoadbalanceResponse) SetErrorParams(v []string) *ListLoadbalanceResponse {
	s.ErrorParams = v
	return s
}


type ListLoadbalanceResponseBuilder struct {
	s *ListLoadbalanceResponse
}

func NewListLoadbalanceResponseBuilder() *ListLoadbalanceResponseBuilder {
	s := &ListLoadbalanceResponse{}
	b := &ListLoadbalanceResponseBuilder{s: s}
	return b
}

func (b *ListLoadbalanceResponseBuilder) AsyncKey(v string) *ListLoadbalanceResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListLoadbalanceResponseBuilder) RequestId(v string) *ListLoadbalanceResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadbalanceResponseBuilder) ErrorMessage(v string) *ListLoadbalanceResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadbalanceResponseBuilder) ErrorCode(v string) *ListLoadbalanceResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadbalanceResponseBuilder) State(v ListLoadbalanceResponseStateEnum) *ListLoadbalanceResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadbalanceResponseBuilder) Body(v *ListLoadbalanceResponseBody) *ListLoadbalanceResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadbalanceResponseBuilder) ErrorParams(v []string) *ListLoadbalanceResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadbalanceResponseBuilder) Build() *ListLoadbalanceResponse {
	return b.s
}


