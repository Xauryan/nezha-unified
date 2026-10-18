// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListControlGroupRuleResponseStateEnum string

// List of State
const (
    ListControlGroupRuleResponseStateEnumOk ListControlGroupRuleResponseStateEnum = "OK"
    ListControlGroupRuleResponseStateEnumError ListControlGroupRuleResponseStateEnum = "ERROR"
    ListControlGroupRuleResponseStateEnumException ListControlGroupRuleResponseStateEnum = "EXCEPTION"
    ListControlGroupRuleResponseStateEnumAlarm ListControlGroupRuleResponseStateEnum = "ALARM"
    ListControlGroupRuleResponseStateEnumForbidden ListControlGroupRuleResponseStateEnum = "FORBIDDEN"
)

type ListControlGroupRuleResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListControlGroupRuleResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListControlGroupRuleResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListControlGroupRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRuleResponse) GoString() string {
	return s.String()
}

func (s ListControlGroupRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRuleResponse) SetAsyncKey(v string) *ListControlGroupRuleResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListControlGroupRuleResponse) SetRequestId(v string) *ListControlGroupRuleResponse {
	s.RequestId = &v
	return s
}

func (s *ListControlGroupRuleResponse) SetErrorMessage(v string) *ListControlGroupRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListControlGroupRuleResponse) SetErrorCode(v string) *ListControlGroupRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListControlGroupRuleResponse) SetState(v ListControlGroupRuleResponseStateEnum) *ListControlGroupRuleResponse {
	s.State = &v
	return s
}

func (s *ListControlGroupRuleResponse) SetBody(v *ListControlGroupRuleResponseBody) *ListControlGroupRuleResponse {
	s.Body = v
	return s
}

func (s *ListControlGroupRuleResponse) SetErrorParams(v []string) *ListControlGroupRuleResponse {
	s.ErrorParams = v
	return s
}


type ListControlGroupRuleResponseBuilder struct {
	s *ListControlGroupRuleResponse
}

func NewListControlGroupRuleResponseBuilder() *ListControlGroupRuleResponseBuilder {
	s := &ListControlGroupRuleResponse{}
	b := &ListControlGroupRuleResponseBuilder{s: s}
	return b
}

func (b *ListControlGroupRuleResponseBuilder) AsyncKey(v string) *ListControlGroupRuleResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) RequestId(v string) *ListControlGroupRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) ErrorMessage(v string) *ListControlGroupRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) ErrorCode(v string) *ListControlGroupRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) State(v ListControlGroupRuleResponseStateEnum) *ListControlGroupRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) Body(v *ListControlGroupRuleResponseBody) *ListControlGroupRuleResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) ErrorParams(v []string) *ListControlGroupRuleResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListControlGroupRuleResponseBuilder) Build() *ListControlGroupRuleResponse {
	return b.s
}


