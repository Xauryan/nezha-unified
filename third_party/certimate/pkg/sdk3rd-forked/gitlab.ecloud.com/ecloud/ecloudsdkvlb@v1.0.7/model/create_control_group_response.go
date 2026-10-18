// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateControlGroupResponseStateEnum string

// List of State
const (
    CreateControlGroupResponseStateEnumOk CreateControlGroupResponseStateEnum = "OK"
    CreateControlGroupResponseStateEnumError CreateControlGroupResponseStateEnum = "ERROR"
    CreateControlGroupResponseStateEnumException CreateControlGroupResponseStateEnum = "EXCEPTION"
    CreateControlGroupResponseStateEnumAlarm CreateControlGroupResponseStateEnum = "ALARM"
    CreateControlGroupResponseStateEnumForbidden CreateControlGroupResponseStateEnum = "FORBIDDEN"
)

type CreateControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupResponse) GoString() string {
	return s.String()
}

func (s CreateControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupResponse) SetAsyncKey(v string) *CreateControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *CreateControlGroupResponse) SetRequestId(v string) *CreateControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateControlGroupResponse) SetErrorMessage(v string) *CreateControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateControlGroupResponse) SetErrorCode(v string) *CreateControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateControlGroupResponse) SetState(v CreateControlGroupResponseStateEnum) *CreateControlGroupResponse {
	s.State = &v
	return s
}

func (s *CreateControlGroupResponse) SetBody(v string) *CreateControlGroupResponse {
	s.Body = &v
	return s
}

func (s *CreateControlGroupResponse) SetErrorParams(v []string) *CreateControlGroupResponse {
	s.ErrorParams = v
	return s
}


type CreateControlGroupResponseBuilder struct {
	s *CreateControlGroupResponse
}

func NewCreateControlGroupResponseBuilder() *CreateControlGroupResponseBuilder {
	s := &CreateControlGroupResponse{}
	b := &CreateControlGroupResponseBuilder{s: s}
	return b
}

func (b *CreateControlGroupResponseBuilder) AsyncKey(v string) *CreateControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) RequestId(v string) *CreateControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) ErrorMessage(v string) *CreateControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) ErrorCode(v string) *CreateControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) State(v CreateControlGroupResponseStateEnum) *CreateControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) Body(v string) *CreateControlGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateControlGroupResponseBuilder) ErrorParams(v []string) *CreateControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateControlGroupResponseBuilder) Build() *CreateControlGroupResponse {
	return b.s
}


