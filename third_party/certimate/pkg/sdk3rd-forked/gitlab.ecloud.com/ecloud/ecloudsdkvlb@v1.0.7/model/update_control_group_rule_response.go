// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateControlGroupRuleResponseStateEnum string

// List of State
const (
    UpdateControlGroupRuleResponseStateEnumOk UpdateControlGroupRuleResponseStateEnum = "OK"
    UpdateControlGroupRuleResponseStateEnumError UpdateControlGroupRuleResponseStateEnum = "ERROR"
    UpdateControlGroupRuleResponseStateEnumException UpdateControlGroupRuleResponseStateEnum = "EXCEPTION"
    UpdateControlGroupRuleResponseStateEnumAlarm UpdateControlGroupRuleResponseStateEnum = "ALARM"
    UpdateControlGroupRuleResponseStateEnumForbidden UpdateControlGroupRuleResponseStateEnum = "FORBIDDEN"
)

type UpdateControlGroupRuleResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateControlGroupRuleResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateControlGroupRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupRuleResponse) GoString() string {
	return s.String()
}

func (s UpdateControlGroupRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupRuleResponse) SetAsyncKey(v string) *UpdateControlGroupRuleResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetRequestId(v string) *UpdateControlGroupRuleResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetErrorMessage(v string) *UpdateControlGroupRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetErrorCode(v string) *UpdateControlGroupRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetState(v UpdateControlGroupRuleResponseStateEnum) *UpdateControlGroupRuleResponse {
	s.State = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetBody(v string) *UpdateControlGroupRuleResponse {
	s.Body = &v
	return s
}

func (s *UpdateControlGroupRuleResponse) SetErrorParams(v []string) *UpdateControlGroupRuleResponse {
	s.ErrorParams = v
	return s
}


type UpdateControlGroupRuleResponseBuilder struct {
	s *UpdateControlGroupRuleResponse
}

func NewUpdateControlGroupRuleResponseBuilder() *UpdateControlGroupRuleResponseBuilder {
	s := &UpdateControlGroupRuleResponse{}
	b := &UpdateControlGroupRuleResponseBuilder{s: s}
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) AsyncKey(v string) *UpdateControlGroupRuleResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) RequestId(v string) *UpdateControlGroupRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) ErrorMessage(v string) *UpdateControlGroupRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) ErrorCode(v string) *UpdateControlGroupRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) State(v UpdateControlGroupRuleResponseStateEnum) *UpdateControlGroupRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) Body(v string) *UpdateControlGroupRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) ErrorParams(v []string) *UpdateControlGroupRuleResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateControlGroupRuleResponseBuilder) Build() *UpdateControlGroupRuleResponse {
	return b.s
}


