// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteControlGroupRuleResponseStateEnum string

// List of State
const (
    DeleteControlGroupRuleResponseStateEnumOk DeleteControlGroupRuleResponseStateEnum = "OK"
    DeleteControlGroupRuleResponseStateEnumError DeleteControlGroupRuleResponseStateEnum = "ERROR"
    DeleteControlGroupRuleResponseStateEnumException DeleteControlGroupRuleResponseStateEnum = "EXCEPTION"
    DeleteControlGroupRuleResponseStateEnumAlarm DeleteControlGroupRuleResponseStateEnum = "ALARM"
    DeleteControlGroupRuleResponseStateEnumForbidden DeleteControlGroupRuleResponseStateEnum = "FORBIDDEN"
)

type DeleteControlGroupRuleResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *DeleteControlGroupRuleResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteControlGroupRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupRuleResponse) GoString() string {
	return s.String()
}

func (s DeleteControlGroupRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupRuleResponse) SetAsyncKey(v string) *DeleteControlGroupRuleResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetRequestId(v string) *DeleteControlGroupRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetErrorMessage(v string) *DeleteControlGroupRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetErrorCode(v string) *DeleteControlGroupRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetState(v DeleteControlGroupRuleResponseStateEnum) *DeleteControlGroupRuleResponse {
	s.State = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetBody(v string) *DeleteControlGroupRuleResponse {
	s.Body = &v
	return s
}

func (s *DeleteControlGroupRuleResponse) SetErrorParams(v []string) *DeleteControlGroupRuleResponse {
	s.ErrorParams = v
	return s
}


type DeleteControlGroupRuleResponseBuilder struct {
	s *DeleteControlGroupRuleResponse
}

func NewDeleteControlGroupRuleResponseBuilder() *DeleteControlGroupRuleResponseBuilder {
	s := &DeleteControlGroupRuleResponse{}
	b := &DeleteControlGroupRuleResponseBuilder{s: s}
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) AsyncKey(v string) *DeleteControlGroupRuleResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) RequestId(v string) *DeleteControlGroupRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) ErrorMessage(v string) *DeleteControlGroupRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) ErrorCode(v string) *DeleteControlGroupRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) State(v DeleteControlGroupRuleResponseStateEnum) *DeleteControlGroupRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) Body(v string) *DeleteControlGroupRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) ErrorParams(v []string) *DeleteControlGroupRuleResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteControlGroupRuleResponseBuilder) Build() *DeleteControlGroupRuleResponse {
	return b.s
}


