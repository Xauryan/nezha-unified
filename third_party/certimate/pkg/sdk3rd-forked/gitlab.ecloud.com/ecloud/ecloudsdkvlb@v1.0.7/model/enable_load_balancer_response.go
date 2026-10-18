// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type EnableLoadBalancerResponseStateEnum string

// List of State
const (
    EnableLoadBalancerResponseStateEnumOk EnableLoadBalancerResponseStateEnum = "OK"
    EnableLoadBalancerResponseStateEnumError EnableLoadBalancerResponseStateEnum = "ERROR"
    EnableLoadBalancerResponseStateEnumException EnableLoadBalancerResponseStateEnum = "EXCEPTION"
    EnableLoadBalancerResponseStateEnumAlarm EnableLoadBalancerResponseStateEnum = "ALARM"
    EnableLoadBalancerResponseStateEnumForbidden EnableLoadBalancerResponseStateEnum = "FORBIDDEN"
)

type EnableLoadBalancerResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *EnableLoadBalancerResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s EnableLoadBalancerResponse) String() string {
	return utils.Beautify(s)
}

func (s EnableLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s EnableLoadBalancerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableLoadBalancerResponse) SetAsyncKey(v string) *EnableLoadBalancerResponse {
	s.AsyncKey = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetRequestId(v string) *EnableLoadBalancerResponse {
	s.RequestId = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetErrorMessage(v string) *EnableLoadBalancerResponse {
	s.ErrorMessage = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetErrorCode(v string) *EnableLoadBalancerResponse {
	s.ErrorCode = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetState(v EnableLoadBalancerResponseStateEnum) *EnableLoadBalancerResponse {
	s.State = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetBody(v string) *EnableLoadBalancerResponse {
	s.Body = &v
	return s
}

func (s *EnableLoadBalancerResponse) SetErrorParams(v []string) *EnableLoadBalancerResponse {
	s.ErrorParams = v
	return s
}


type EnableLoadBalancerResponseBuilder struct {
	s *EnableLoadBalancerResponse
}

func NewEnableLoadBalancerResponseBuilder() *EnableLoadBalancerResponseBuilder {
	s := &EnableLoadBalancerResponse{}
	b := &EnableLoadBalancerResponseBuilder{s: s}
	return b
}

func (b *EnableLoadBalancerResponseBuilder) AsyncKey(v string) *EnableLoadBalancerResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) RequestId(v string) *EnableLoadBalancerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) ErrorMessage(v string) *EnableLoadBalancerResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) ErrorCode(v string) *EnableLoadBalancerResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) State(v EnableLoadBalancerResponseStateEnum) *EnableLoadBalancerResponseBuilder {
	b.s.State = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) Body(v string) *EnableLoadBalancerResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) ErrorParams(v []string) *EnableLoadBalancerResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *EnableLoadBalancerResponseBuilder) Build() *EnableLoadBalancerResponse {
	return b.s
}


