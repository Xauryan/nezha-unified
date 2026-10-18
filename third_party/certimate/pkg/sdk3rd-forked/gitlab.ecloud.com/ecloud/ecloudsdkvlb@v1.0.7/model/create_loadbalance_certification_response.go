// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadbalanceCertificationResponseStateEnum string

// List of State
const (
    CreateLoadbalanceCertificationResponseStateEnumOk CreateLoadbalanceCertificationResponseStateEnum = "OK"
    CreateLoadbalanceCertificationResponseStateEnumError CreateLoadbalanceCertificationResponseStateEnum = "ERROR"
    CreateLoadbalanceCertificationResponseStateEnumException CreateLoadbalanceCertificationResponseStateEnum = "EXCEPTION"
    CreateLoadbalanceCertificationResponseStateEnumAlarm CreateLoadbalanceCertificationResponseStateEnum = "ALARM"
    CreateLoadbalanceCertificationResponseStateEnumForbidden CreateLoadbalanceCertificationResponseStateEnum = "FORBIDDEN"
)

type CreateLoadbalanceCertificationResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CreateLoadbalanceCertificationResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CreateLoadbalanceCertificationResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadbalanceCertificationResponse) GoString() string {
	return s.String()
}

func (s CreateLoadbalanceCertificationResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadbalanceCertificationResponse) SetRequestId(v string) *CreateLoadbalanceCertificationResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoadbalanceCertificationResponse) SetErrorMessage(v string) *CreateLoadbalanceCertificationResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLoadbalanceCertificationResponse) SetErrorCode(v string) *CreateLoadbalanceCertificationResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateLoadbalanceCertificationResponse) SetState(v CreateLoadbalanceCertificationResponseStateEnum) *CreateLoadbalanceCertificationResponse {
	s.State = &v
	return s
}

func (s *CreateLoadbalanceCertificationResponse) SetBody(v string) *CreateLoadbalanceCertificationResponse {
	s.Body = &v
	return s
}

func (s *CreateLoadbalanceCertificationResponse) SetErrorParams(v []string) *CreateLoadbalanceCertificationResponse {
	s.ErrorParams = v
	return s
}


type CreateLoadbalanceCertificationResponseBuilder struct {
	s *CreateLoadbalanceCertificationResponse
}

func NewCreateLoadbalanceCertificationResponseBuilder() *CreateLoadbalanceCertificationResponseBuilder {
	s := &CreateLoadbalanceCertificationResponse{}
	b := &CreateLoadbalanceCertificationResponseBuilder{s: s}
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) RequestId(v string) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) ErrorMessage(v string) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) ErrorCode(v string) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) State(v CreateLoadbalanceCertificationResponseStateEnum) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) Body(v string) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) ErrorParams(v []string) *CreateLoadbalanceCertificationResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CreateLoadbalanceCertificationResponseBuilder) Build() *CreateLoadbalanceCertificationResponse {
	return b.s
}


