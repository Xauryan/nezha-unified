// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListControlGroupResponseStateEnum string

// List of State
const (
    ListControlGroupResponseStateEnumOk ListControlGroupResponseStateEnum = "OK"
    ListControlGroupResponseStateEnumError ListControlGroupResponseStateEnum = "ERROR"
    ListControlGroupResponseStateEnumException ListControlGroupResponseStateEnum = "EXCEPTION"
    ListControlGroupResponseStateEnumAlarm ListControlGroupResponseStateEnum = "ALARM"
    ListControlGroupResponseStateEnumForbidden ListControlGroupResponseStateEnum = "FORBIDDEN"
)

type ListControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码:OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListControlGroupResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupResponse) GoString() string {
	return s.String()
}

func (s ListControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupResponse) SetAsyncKey(v string) *ListControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListControlGroupResponse) SetRequestId(v string) *ListControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ListControlGroupResponse) SetErrorMessage(v string) *ListControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListControlGroupResponse) SetErrorCode(v string) *ListControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListControlGroupResponse) SetState(v ListControlGroupResponseStateEnum) *ListControlGroupResponse {
	s.State = &v
	return s
}

func (s *ListControlGroupResponse) SetBody(v *ListControlGroupResponseBody) *ListControlGroupResponse {
	s.Body = v
	return s
}

func (s *ListControlGroupResponse) SetErrorParams(v []string) *ListControlGroupResponse {
	s.ErrorParams = v
	return s
}


type ListControlGroupResponseBuilder struct {
	s *ListControlGroupResponse
}

func NewListControlGroupResponseBuilder() *ListControlGroupResponseBuilder {
	s := &ListControlGroupResponse{}
	b := &ListControlGroupResponseBuilder{s: s}
	return b
}

func (b *ListControlGroupResponseBuilder) AsyncKey(v string) *ListControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListControlGroupResponseBuilder) RequestId(v string) *ListControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListControlGroupResponseBuilder) ErrorMessage(v string) *ListControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListControlGroupResponseBuilder) ErrorCode(v string) *ListControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListControlGroupResponseBuilder) State(v ListControlGroupResponseStateEnum) *ListControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListControlGroupResponseBuilder) Body(v *ListControlGroupResponseBody) *ListControlGroupResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListControlGroupResponseBuilder) ErrorParams(v []string) *ListControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListControlGroupResponseBuilder) Build() *ListControlGroupResponse {
	return b.s
}


