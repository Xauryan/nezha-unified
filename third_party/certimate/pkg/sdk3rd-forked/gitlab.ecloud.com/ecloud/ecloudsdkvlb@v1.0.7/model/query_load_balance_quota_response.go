// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type QueryLoadBalanceQuotaResponseStateEnum string

// List of State
const (
    QueryLoadBalanceQuotaResponseStateEnumOk QueryLoadBalanceQuotaResponseStateEnum = "OK"
    QueryLoadBalanceQuotaResponseStateEnumError QueryLoadBalanceQuotaResponseStateEnum = "ERROR"
    QueryLoadBalanceQuotaResponseStateEnumException QueryLoadBalanceQuotaResponseStateEnum = "EXCEPTION"
    QueryLoadBalanceQuotaResponseStateEnumAlarm QueryLoadBalanceQuotaResponseStateEnum = "ALARM"
    QueryLoadBalanceQuotaResponseStateEnumForbidden QueryLoadBalanceQuotaResponseStateEnum = "FORBIDDEN"
)

type QueryLoadBalanceQuotaResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码
	State *QueryLoadBalanceQuotaResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *QueryLoadBalanceQuotaResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s QueryLoadBalanceQuotaResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryLoadBalanceQuotaResponse) GoString() string {
	return s.String()
}

func (s QueryLoadBalanceQuotaResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryLoadBalanceQuotaResponse) SetAsyncKey(v string) *QueryLoadBalanceQuotaResponse {
	s.AsyncKey = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetRequestId(v string) *QueryLoadBalanceQuotaResponse {
	s.RequestId = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetErrorMessage(v string) *QueryLoadBalanceQuotaResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetErrorCode(v string) *QueryLoadBalanceQuotaResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetState(v QueryLoadBalanceQuotaResponseStateEnum) *QueryLoadBalanceQuotaResponse {
	s.State = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetBody(v *QueryLoadBalanceQuotaResponseBody) *QueryLoadBalanceQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryLoadBalanceQuotaResponse) SetErrorParams(v []string) *QueryLoadBalanceQuotaResponse {
	s.ErrorParams = v
	return s
}


type QueryLoadBalanceQuotaResponseBuilder struct {
	s *QueryLoadBalanceQuotaResponse
}

func NewQueryLoadBalanceQuotaResponseBuilder() *QueryLoadBalanceQuotaResponseBuilder {
	s := &QueryLoadBalanceQuotaResponse{}
	b := &QueryLoadBalanceQuotaResponseBuilder{s: s}
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) AsyncKey(v string) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) RequestId(v string) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) ErrorMessage(v string) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) ErrorCode(v string) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) State(v QueryLoadBalanceQuotaResponseStateEnum) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) Body(v *QueryLoadBalanceQuotaResponseBody) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) ErrorParams(v []string) *QueryLoadBalanceQuotaResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBuilder) Build() *QueryLoadBalanceQuotaResponse {
	return b.s
}


