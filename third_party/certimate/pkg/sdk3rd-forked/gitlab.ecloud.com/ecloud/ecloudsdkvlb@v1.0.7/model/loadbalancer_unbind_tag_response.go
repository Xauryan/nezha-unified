// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LoadbalancerUnbindTagResponseStateEnum string

// List of State
const (
    LoadbalancerUnbindTagResponseStateEnumOk LoadbalancerUnbindTagResponseStateEnum = "OK"
    LoadbalancerUnbindTagResponseStateEnumError LoadbalancerUnbindTagResponseStateEnum = "ERROR"
    LoadbalancerUnbindTagResponseStateEnumException LoadbalancerUnbindTagResponseStateEnum = "EXCEPTION"
    LoadbalancerUnbindTagResponseStateEnumAlarm LoadbalancerUnbindTagResponseStateEnum = "ALARM"
    LoadbalancerUnbindTagResponseStateEnumForbidden LoadbalancerUnbindTagResponseStateEnum = "FORBIDDEN"
)

type LoadbalancerUnbindTagResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *LoadbalancerUnbindTagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s LoadbalancerUnbindTagResponse) String() string {
	return utils.Beautify(s)
}

func (s LoadbalancerUnbindTagResponse) GoString() string {
	return s.String()
}

func (s LoadbalancerUnbindTagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoadbalancerUnbindTagResponse) SetAsyncKey(v string) *LoadbalancerUnbindTagResponse {
	s.AsyncKey = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetRequestId(v string) *LoadbalancerUnbindTagResponse {
	s.RequestId = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetErrorMessage(v string) *LoadbalancerUnbindTagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetErrorCode(v string) *LoadbalancerUnbindTagResponse {
	s.ErrorCode = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetState(v LoadbalancerUnbindTagResponseStateEnum) *LoadbalancerUnbindTagResponse {
	s.State = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetBody(v string) *LoadbalancerUnbindTagResponse {
	s.Body = &v
	return s
}

func (s *LoadbalancerUnbindTagResponse) SetErrorParams(v []string) *LoadbalancerUnbindTagResponse {
	s.ErrorParams = v
	return s
}


type LoadbalancerUnbindTagResponseBuilder struct {
	s *LoadbalancerUnbindTagResponse
}

func NewLoadbalancerUnbindTagResponseBuilder() *LoadbalancerUnbindTagResponseBuilder {
	s := &LoadbalancerUnbindTagResponse{}
	b := &LoadbalancerUnbindTagResponseBuilder{s: s}
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) AsyncKey(v string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) RequestId(v string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) ErrorMessage(v string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) ErrorCode(v string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) State(v LoadbalancerUnbindTagResponseStateEnum) *LoadbalancerUnbindTagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) Body(v string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) ErrorParams(v []string) *LoadbalancerUnbindTagResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *LoadbalancerUnbindTagResponseBuilder) Build() *LoadbalancerUnbindTagResponse {
	return b.s
}


