// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadBalancListenerNameResponseStateEnum string

// List of State
const (
    UpdateLoadBalancListenerNameResponseStateEnumOk UpdateLoadBalancListenerNameResponseStateEnum = "OK"
    UpdateLoadBalancListenerNameResponseStateEnumError UpdateLoadBalancListenerNameResponseStateEnum = "ERROR"
    UpdateLoadBalancListenerNameResponseStateEnumException UpdateLoadBalancListenerNameResponseStateEnum = "EXCEPTION"
    UpdateLoadBalancListenerNameResponseStateEnumAlarm UpdateLoadBalancListenerNameResponseStateEnum = "ALARM"
    UpdateLoadBalancListenerNameResponseStateEnumForbidden UpdateLoadBalancListenerNameResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadBalancListenerNameResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadBalancListenerNameResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadBalancListenerNameResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancListenerNameResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancListenerNameResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancListenerNameResponse) SetAsyncKey(v string) *UpdateLoadBalancListenerNameResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetRequestId(v string) *UpdateLoadBalancListenerNameResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetErrorMessage(v string) *UpdateLoadBalancListenerNameResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetErrorCode(v string) *UpdateLoadBalancListenerNameResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetState(v UpdateLoadBalancListenerNameResponseStateEnum) *UpdateLoadBalancListenerNameResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetBody(v string) *UpdateLoadBalancListenerNameResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadBalancListenerNameResponse) SetErrorParams(v []string) *UpdateLoadBalancListenerNameResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadBalancListenerNameResponseBuilder struct {
	s *UpdateLoadBalancListenerNameResponse
}

func NewUpdateLoadBalancListenerNameResponseBuilder() *UpdateLoadBalancListenerNameResponseBuilder {
	s := &UpdateLoadBalancListenerNameResponse{}
	b := &UpdateLoadBalancListenerNameResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) AsyncKey(v string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) RequestId(v string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) ErrorMessage(v string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) ErrorCode(v string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) State(v UpdateLoadBalancListenerNameResponseStateEnum) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) Body(v string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) ErrorParams(v []string) *UpdateLoadBalancListenerNameResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadBalancListenerNameResponseBuilder) Build() *UpdateLoadBalancListenerNameResponse {
	return b.s
}


