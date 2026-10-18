// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadBalanceListenerDetailRespResponseStateEnum string

// List of State
const (
    GetLoadBalanceListenerDetailRespResponseStateEnumOk GetLoadBalanceListenerDetailRespResponseStateEnum = "OK"
    GetLoadBalanceListenerDetailRespResponseStateEnumError GetLoadBalanceListenerDetailRespResponseStateEnum = "ERROR"
    GetLoadBalanceListenerDetailRespResponseStateEnumException GetLoadBalanceListenerDetailRespResponseStateEnum = "EXCEPTION"
    GetLoadBalanceListenerDetailRespResponseStateEnumAlarm GetLoadBalanceListenerDetailRespResponseStateEnum = "ALARM"
    GetLoadBalanceListenerDetailRespResponseStateEnumForbidden GetLoadBalanceListenerDetailRespResponseStateEnum = "FORBIDDEN"
)

type GetLoadBalanceListenerDetailRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *GetLoadBalanceListenerDetailRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetLoadBalanceListenerDetailRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetLoadBalanceListenerDetailRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceListenerDetailRespResponse) GoString() string {
	return s.String()
}

func (s GetLoadBalanceListenerDetailRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetAsyncKey(v string) *GetLoadBalanceListenerDetailRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetRequestId(v string) *GetLoadBalanceListenerDetailRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetErrorMessage(v string) *GetLoadBalanceListenerDetailRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetErrorCode(v string) *GetLoadBalanceListenerDetailRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetState(v GetLoadBalanceListenerDetailRespResponseStateEnum) *GetLoadBalanceListenerDetailRespResponse {
	s.State = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetBody(v *GetLoadBalanceListenerDetailRespResponseBody) *GetLoadBalanceListenerDetailRespResponse {
	s.Body = v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponse) SetErrorParams(v []string) *GetLoadBalanceListenerDetailRespResponse {
	s.ErrorParams = v
	return s
}


type GetLoadBalanceListenerDetailRespResponseBuilder struct {
	s *GetLoadBalanceListenerDetailRespResponse
}

func NewGetLoadBalanceListenerDetailRespResponseBuilder() *GetLoadBalanceListenerDetailRespResponseBuilder {
	s := &GetLoadBalanceListenerDetailRespResponse{}
	b := &GetLoadBalanceListenerDetailRespResponseBuilder{s: s}
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) AsyncKey(v string) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) RequestId(v string) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) ErrorMessage(v string) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) ErrorCode(v string) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) State(v GetLoadBalanceListenerDetailRespResponseStateEnum) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) Body(v *GetLoadBalanceListenerDetailRespResponseBody) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) ErrorParams(v []string) *GetLoadBalanceListenerDetailRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBuilder) Build() *GetLoadBalanceListenerDetailRespResponse {
	return b.s
}


