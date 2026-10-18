// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateLoadbalanceCertificationResponseStateEnum string

// List of State
const (
    UpdateLoadbalanceCertificationResponseStateEnumOk UpdateLoadbalanceCertificationResponseStateEnum = "OK"
    UpdateLoadbalanceCertificationResponseStateEnumError UpdateLoadbalanceCertificationResponseStateEnum = "ERROR"
    UpdateLoadbalanceCertificationResponseStateEnumException UpdateLoadbalanceCertificationResponseStateEnum = "EXCEPTION"
    UpdateLoadbalanceCertificationResponseStateEnumAlarm UpdateLoadbalanceCertificationResponseStateEnum = "ALARM"
    UpdateLoadbalanceCertificationResponseStateEnumForbidden UpdateLoadbalanceCertificationResponseStateEnum = "FORBIDDEN"
)

type UpdateLoadbalanceCertificationResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *UpdateLoadbalanceCertificationResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s UpdateLoadbalanceCertificationResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceCertificationResponse) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceCertificationResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceCertificationResponse) SetRequestId(v string) *UpdateLoadbalanceCertificationResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadbalanceCertificationResponse) SetErrorMessage(v string) *UpdateLoadbalanceCertificationResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLoadbalanceCertificationResponse) SetErrorCode(v string) *UpdateLoadbalanceCertificationResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLoadbalanceCertificationResponse) SetState(v UpdateLoadbalanceCertificationResponseStateEnum) *UpdateLoadbalanceCertificationResponse {
	s.State = &v
	return s
}

func (s *UpdateLoadbalanceCertificationResponse) SetBody(v string) *UpdateLoadbalanceCertificationResponse {
	s.Body = &v
	return s
}

func (s *UpdateLoadbalanceCertificationResponse) SetErrorParams(v []string) *UpdateLoadbalanceCertificationResponse {
	s.ErrorParams = v
	return s
}


type UpdateLoadbalanceCertificationResponseBuilder struct {
	s *UpdateLoadbalanceCertificationResponse
}

func NewUpdateLoadbalanceCertificationResponseBuilder() *UpdateLoadbalanceCertificationResponseBuilder {
	s := &UpdateLoadbalanceCertificationResponse{}
	b := &UpdateLoadbalanceCertificationResponseBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) RequestId(v string) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) ErrorMessage(v string) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) ErrorCode(v string) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) State(v UpdateLoadbalanceCertificationResponseStateEnum) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) Body(v string) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) ErrorParams(v []string) *UpdateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *UpdateLoadbalanceCertificationResponseBuilder) Build() *UpdateLoadbalanceCertificationResponse {
	return b.s
}


