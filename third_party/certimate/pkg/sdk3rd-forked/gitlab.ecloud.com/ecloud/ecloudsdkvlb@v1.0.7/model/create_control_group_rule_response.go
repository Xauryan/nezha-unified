// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateControlGroupRuleResponseStateEnum string

// List of State
const (
    CreateControlGroupRuleResponseStateEnumOk CreateControlGroupRuleResponseStateEnum = "OK"
    CreateControlGroupRuleResponseStateEnumError CreateControlGroupRuleResponseStateEnum = "ERROR"
    CreateControlGroupRuleResponseStateEnumException CreateControlGroupRuleResponseStateEnum = "EXCEPTION"
    CreateControlGroupRuleResponseStateEnumAlarm CreateControlGroupRuleResponseStateEnum = "ALARM"
    CreateControlGroupRuleResponseStateEnumForbidden CreateControlGroupRuleResponseStateEnum = "FORBIDDEN"
)

type CreateControlGroupRuleResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateControlGroupRuleResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateControlGroupRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupRuleResponse) GoString() string {
	return s.String()
}

func (s CreateControlGroupRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupRuleResponse) SetAsyncKey(v string) *CreateControlGroupRuleResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetRequestId(v string) *CreateControlGroupRuleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetErrorMessage(v string) *CreateControlGroupRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetErrorCode(v string) *CreateControlGroupRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetState(v CreateControlGroupRuleResponseStateEnum) *CreateControlGroupRuleResponse {
	s.State = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetBody(v string) *CreateControlGroupRuleResponse {
	s.Body = &v
	return s
}

func (s *CreateControlGroupRuleResponse) SetErrorParams(v []string) *CreateControlGroupRuleResponse {
	s.ErrorParams = v
	return s
}


type CreateControlGroupRuleResponseBuilder struct {
	s *CreateControlGroupRuleResponse
}

func NewCreateControlGroupRuleResponseBuilder() *CreateControlGroupRuleResponseBuilder {
	s := &CreateControlGroupRuleResponse{}
	b := &CreateControlGroupRuleResponseBuilder{s: s}
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) AsyncKey(v string) *CreateControlGroupRuleResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) RequestId(v string) *CreateControlGroupRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) ErrorMessage(v string) *CreateControlGroupRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) ErrorCode(v string) *CreateControlGroupRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) State(v CreateControlGroupRuleResponseStateEnum) *CreateControlGroupRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) Body(v string) *CreateControlGroupRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) ErrorParams(v []string) *CreateControlGroupRuleResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateControlGroupRuleResponseBuilder) Build() *CreateControlGroupRuleResponse {
	return b.s
}


