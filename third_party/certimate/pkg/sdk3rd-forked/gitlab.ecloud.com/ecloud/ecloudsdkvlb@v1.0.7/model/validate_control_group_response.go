// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ValidateControlGroupResponseStateEnum string

// List of State
const (
    ValidateControlGroupResponseStateEnumOk ValidateControlGroupResponseStateEnum = "OK"
    ValidateControlGroupResponseStateEnumError ValidateControlGroupResponseStateEnum = "ERROR"
    ValidateControlGroupResponseStateEnumException ValidateControlGroupResponseStateEnum = "EXCEPTION"
    ValidateControlGroupResponseStateEnumAlarm ValidateControlGroupResponseStateEnum = "ALARM"
    ValidateControlGroupResponseStateEnumForbidden ValidateControlGroupResponseStateEnum = "FORBIDDEN"
)

type ValidateControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ValidateControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ValidateControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s ValidateControlGroupResponse) GoString() string {
	return s.String()
}

func (s ValidateControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ValidateControlGroupResponse) SetAsyncKey(v string) *ValidateControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *ValidateControlGroupResponse) SetRequestId(v string) *ValidateControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ValidateControlGroupResponse) SetErrorMessage(v string) *ValidateControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ValidateControlGroupResponse) SetErrorCode(v string) *ValidateControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *ValidateControlGroupResponse) SetState(v ValidateControlGroupResponseStateEnum) *ValidateControlGroupResponse {
	s.State = &v
	return s
}

func (s *ValidateControlGroupResponse) SetBody(v string) *ValidateControlGroupResponse {
	s.Body = &v
	return s
}

func (s *ValidateControlGroupResponse) SetErrorParams(v []string) *ValidateControlGroupResponse {
	s.ErrorParams = v
	return s
}


type ValidateControlGroupResponseBuilder struct {
	s *ValidateControlGroupResponse
}

func NewValidateControlGroupResponseBuilder() *ValidateControlGroupResponseBuilder {
	s := &ValidateControlGroupResponse{}
	b := &ValidateControlGroupResponseBuilder{s: s}
	return b
}

func (b *ValidateControlGroupResponseBuilder) AsyncKey(v string) *ValidateControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) RequestId(v string) *ValidateControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) ErrorMessage(v string) *ValidateControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) ErrorCode(v string) *ValidateControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) State(v ValidateControlGroupResponseStateEnum) *ValidateControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) Body(v string) *ValidateControlGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ValidateControlGroupResponseBuilder) ErrorParams(v []string) *ValidateControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ValidateControlGroupResponseBuilder) Build() *ValidateControlGroupResponse {
	return b.s
}


