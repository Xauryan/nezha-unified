// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateControlGroupResponseStateEnum string

// List of State
const (
    UpdateControlGroupResponseStateEnumOk UpdateControlGroupResponseStateEnum = "OK"
    UpdateControlGroupResponseStateEnumError UpdateControlGroupResponseStateEnum = "ERROR"
    UpdateControlGroupResponseStateEnumException UpdateControlGroupResponseStateEnum = "EXCEPTION"
    UpdateControlGroupResponseStateEnumAlarm UpdateControlGroupResponseStateEnum = "ALARM"
    UpdateControlGroupResponseStateEnumForbidden UpdateControlGroupResponseStateEnum = "FORBIDDEN"
)

type UpdateControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码:OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupResponse) GoString() string {
	return s.String()
}

func (s UpdateControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupResponse) SetAsyncKey(v string) *UpdateControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateControlGroupResponse) SetRequestId(v string) *UpdateControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateControlGroupResponse) SetErrorMessage(v string) *UpdateControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateControlGroupResponse) SetErrorCode(v string) *UpdateControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateControlGroupResponse) SetState(v UpdateControlGroupResponseStateEnum) *UpdateControlGroupResponse {
	s.State = &v
	return s
}

func (s *UpdateControlGroupResponse) SetBody(v string) *UpdateControlGroupResponse {
	s.Body = &v
	return s
}

func (s *UpdateControlGroupResponse) SetErrorParams(v []string) *UpdateControlGroupResponse {
	s.ErrorParams = v
	return s
}


type UpdateControlGroupResponseBuilder struct {
	s *UpdateControlGroupResponse
}

func NewUpdateControlGroupResponseBuilder() *UpdateControlGroupResponseBuilder {
	s := &UpdateControlGroupResponse{}
	b := &UpdateControlGroupResponseBuilder{s: s}
	return b
}

func (b *UpdateControlGroupResponseBuilder) AsyncKey(v string) *UpdateControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) RequestId(v string) *UpdateControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) ErrorMessage(v string) *UpdateControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) ErrorCode(v string) *UpdateControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) State(v UpdateControlGroupResponseStateEnum) *UpdateControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) Body(v string) *UpdateControlGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateControlGroupResponseBuilder) ErrorParams(v []string) *UpdateControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateControlGroupResponseBuilder) Build() *UpdateControlGroupResponse {
	return b.s
}


