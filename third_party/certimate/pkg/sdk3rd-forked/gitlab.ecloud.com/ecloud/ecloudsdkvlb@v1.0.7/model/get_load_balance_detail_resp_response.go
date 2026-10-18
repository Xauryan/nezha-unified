// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadBalanceDetailRespResponseStateEnum string

// List of State
const (
    GetLoadBalanceDetailRespResponseStateEnumOk GetLoadBalanceDetailRespResponseStateEnum = "OK"
    GetLoadBalanceDetailRespResponseStateEnumError GetLoadBalanceDetailRespResponseStateEnum = "ERROR"
    GetLoadBalanceDetailRespResponseStateEnumException GetLoadBalanceDetailRespResponseStateEnum = "EXCEPTION"
    GetLoadBalanceDetailRespResponseStateEnumAlarm GetLoadBalanceDetailRespResponseStateEnum = "ALARM"
    GetLoadBalanceDetailRespResponseStateEnumForbidden GetLoadBalanceDetailRespResponseStateEnum = "FORBIDDEN"
)

type GetLoadBalanceDetailRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *GetLoadBalanceDetailRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetLoadBalanceDetailRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetLoadBalanceDetailRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespResponse) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespResponse) SetAsyncKey(v string) *GetLoadBalanceDetailRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetRequestId(v string) *GetLoadBalanceDetailRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetErrorMessage(v string) *GetLoadBalanceDetailRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetErrorCode(v string) *GetLoadBalanceDetailRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetState(v GetLoadBalanceDetailRespResponseStateEnum) *GetLoadBalanceDetailRespResponse {
	s.State = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetBody(v *GetLoadBalanceDetailRespResponseBody) *GetLoadBalanceDetailRespResponse {
	s.Body = v
	return s
}

func (s *GetLoadBalanceDetailRespResponse) SetErrorParams(v []string) *GetLoadBalanceDetailRespResponse {
	s.ErrorParams = v
	return s
}


type GetLoadBalanceDetailRespResponseBuilder struct {
	s *GetLoadBalanceDetailRespResponse
}

func NewGetLoadBalanceDetailRespResponseBuilder() *GetLoadBalanceDetailRespResponseBuilder {
	s := &GetLoadBalanceDetailRespResponse{}
	b := &GetLoadBalanceDetailRespResponseBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) AsyncKey(v string) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) RequestId(v string) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) ErrorMessage(v string) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) ErrorCode(v string) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) State(v GetLoadBalanceDetailRespResponseStateEnum) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) Body(v *GetLoadBalanceDetailRespResponseBody) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) ErrorParams(v []string) *GetLoadBalanceDetailRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBuilder) Build() *GetLoadBalanceDetailRespResponse {
	return b.s
}


