// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadbalanceNameResponseStateEnum string

// List of State
const (
    UpdateLoadbalanceNameResponseStateEnumOk UpdateLoadbalanceNameResponseStateEnum = "OK"
    UpdateLoadbalanceNameResponseStateEnumError UpdateLoadbalanceNameResponseStateEnum = "ERROR"
    UpdateLoadbalanceNameResponseStateEnumException UpdateLoadbalanceNameResponseStateEnum = "EXCEPTION"
    UpdateLoadbalanceNameResponseStateEnumAlarm UpdateLoadbalanceNameResponseStateEnum = "ALARM"
    UpdateLoadbalanceNameResponseStateEnumForbidden UpdateLoadbalanceNameResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadbalanceNameResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadbalanceNameResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadbalanceNameResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceNameResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceNameResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceNameResponse) SetAsyncKey(v string) *UpdateLoadbalanceNameResponse {
	s.AsyncKey = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetRequestId(v string) *UpdateLoadbalanceNameResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetErrorMessage(v string) *UpdateLoadbalanceNameResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetErrorCode(v string) *UpdateLoadbalanceNameResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetState(v UpdateLoadbalanceNameResponseStateEnum) *UpdateLoadbalanceNameResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetBody(v string) *UpdateLoadbalanceNameResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadbalanceNameResponse) SetErrorParams(v []string) *UpdateLoadbalanceNameResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadbalanceNameResponseBuilder struct {
	s *UpdateLoadbalanceNameResponse
}

func NewUpdateLoadbalanceNameResponseBuilder() *UpdateLoadbalanceNameResponseBuilder {
	s := &UpdateLoadbalanceNameResponse{}
	b := &UpdateLoadbalanceNameResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) AsyncKey(v string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) RequestId(v string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) ErrorMessage(v string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) ErrorCode(v string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) State(v UpdateLoadbalanceNameResponseStateEnum) *UpdateLoadbalanceNameResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) Body(v string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) ErrorParams(v []string) *UpdateLoadbalanceNameResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadbalanceNameResponseBuilder) Build() *UpdateLoadbalanceNameResponse {
	return b.s
}


