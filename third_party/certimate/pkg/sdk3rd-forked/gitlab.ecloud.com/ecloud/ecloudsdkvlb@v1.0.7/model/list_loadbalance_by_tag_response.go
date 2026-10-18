// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceByTagResponseStateEnum string

// List of State
const (
    ListLoadbalanceByTagResponseStateEnumOk ListLoadbalanceByTagResponseStateEnum = "OK"
    ListLoadbalanceByTagResponseStateEnumError ListLoadbalanceByTagResponseStateEnum = "ERROR"
    ListLoadbalanceByTagResponseStateEnumException ListLoadbalanceByTagResponseStateEnum = "EXCEPTION"
    ListLoadbalanceByTagResponseStateEnumAlarm ListLoadbalanceByTagResponseStateEnum = "ALARM"
    ListLoadbalanceByTagResponseStateEnumForbidden ListLoadbalanceByTagResponseStateEnum = "FORBIDDEN"
)

type ListLoadbalanceByTagResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListLoadbalanceByTagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListLoadbalanceByTagResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListLoadbalanceByTagResponse) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagResponse) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagResponse) SetAsyncKey(v string) *ListLoadbalanceByTagResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetRequestId(v string) *ListLoadbalanceByTagResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetErrorMessage(v string) *ListLoadbalanceByTagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetErrorCode(v string) *ListLoadbalanceByTagResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetState(v ListLoadbalanceByTagResponseStateEnum) *ListLoadbalanceByTagResponse {
	s.State = &v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetBody(v *ListLoadbalanceByTagResponseBody) *ListLoadbalanceByTagResponse {
	s.Body = v
	return s
}

func (s *ListLoadbalanceByTagResponse) SetErrorParams(v []string) *ListLoadbalanceByTagResponse {
	s.ErrorParams = v
	return s
}


type ListLoadbalanceByTagResponseBuilder struct {
	s *ListLoadbalanceByTagResponse
}

func NewListLoadbalanceByTagResponseBuilder() *ListLoadbalanceByTagResponseBuilder {
	s := &ListLoadbalanceByTagResponse{}
	b := &ListLoadbalanceByTagResponseBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) AsyncKey(v string) *ListLoadbalanceByTagResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) RequestId(v string) *ListLoadbalanceByTagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) ErrorMessage(v string) *ListLoadbalanceByTagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) ErrorCode(v string) *ListLoadbalanceByTagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) State(v ListLoadbalanceByTagResponseStateEnum) *ListLoadbalanceByTagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) Body(v *ListLoadbalanceByTagResponseBody) *ListLoadbalanceByTagResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) ErrorParams(v []string) *ListLoadbalanceByTagResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListLoadbalanceByTagResponseBuilder) Build() *ListLoadbalanceByTagResponse {
	return b.s
}


