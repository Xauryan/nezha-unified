// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadBalancePoolMemberRespResponseStateEnum string

// List of State
const (
    GetLoadBalancePoolMemberRespResponseStateEnumOk GetLoadBalancePoolMemberRespResponseStateEnum = "OK"
    GetLoadBalancePoolMemberRespResponseStateEnumError GetLoadBalancePoolMemberRespResponseStateEnum = "ERROR"
    GetLoadBalancePoolMemberRespResponseStateEnumException GetLoadBalancePoolMemberRespResponseStateEnum = "EXCEPTION"
    GetLoadBalancePoolMemberRespResponseStateEnumAlarm GetLoadBalancePoolMemberRespResponseStateEnum = "ALARM"
    GetLoadBalancePoolMemberRespResponseStateEnumForbidden GetLoadBalancePoolMemberRespResponseStateEnum = "FORBIDDEN"
)

type GetLoadBalancePoolMemberRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *GetLoadBalancePoolMemberRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetLoadBalancePoolMemberRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetLoadBalancePoolMemberRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalancePoolMemberRespResponse) GoString() string {
	return s.String()
}

func (s GetLoadBalancePoolMemberRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalancePoolMemberRespResponse) SetAsyncKey(v string) *GetLoadBalancePoolMemberRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetRequestId(v string) *GetLoadBalancePoolMemberRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetErrorMessage(v string) *GetLoadBalancePoolMemberRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetErrorCode(v string) *GetLoadBalancePoolMemberRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetState(v GetLoadBalancePoolMemberRespResponseStateEnum) *GetLoadBalancePoolMemberRespResponse {
	s.State = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetBody(v *GetLoadBalancePoolMemberRespResponseBody) *GetLoadBalancePoolMemberRespResponse {
	s.Body = v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponse) SetErrorParams(v []string) *GetLoadBalancePoolMemberRespResponse {
	s.ErrorParams = v
	return s
}


type GetLoadBalancePoolMemberRespResponseBuilder struct {
	s *GetLoadBalancePoolMemberRespResponse
}

func NewGetLoadBalancePoolMemberRespResponseBuilder() *GetLoadBalancePoolMemberRespResponseBuilder {
	s := &GetLoadBalancePoolMemberRespResponse{}
	b := &GetLoadBalancePoolMemberRespResponseBuilder{s: s}
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) AsyncKey(v string) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) RequestId(v string) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) ErrorMessage(v string) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) ErrorCode(v string) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) State(v GetLoadBalancePoolMemberRespResponseStateEnum) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) Body(v *GetLoadBalancePoolMemberRespResponseBody) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) ErrorParams(v []string) *GetLoadBalancePoolMemberRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBuilder) Build() *GetLoadBalancePoolMemberRespResponse {
	return b.s
}


