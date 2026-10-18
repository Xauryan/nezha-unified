// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadbalanceExpirationCertificationResponseStateEnum string

// List of State
const (
    GetLoadbalanceExpirationCertificationResponseStateEnumOk GetLoadbalanceExpirationCertificationResponseStateEnum = "OK"
    GetLoadbalanceExpirationCertificationResponseStateEnumError GetLoadbalanceExpirationCertificationResponseStateEnum = "ERROR"
    GetLoadbalanceExpirationCertificationResponseStateEnumException GetLoadbalanceExpirationCertificationResponseStateEnum = "EXCEPTION"
    GetLoadbalanceExpirationCertificationResponseStateEnumAlarm GetLoadbalanceExpirationCertificationResponseStateEnum = "ALARM"
    GetLoadbalanceExpirationCertificationResponseStateEnumForbidden GetLoadbalanceExpirationCertificationResponseStateEnum = "FORBIDDEN"
)

type GetLoadbalanceExpirationCertificationResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *GetLoadbalanceExpirationCertificationResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]GetLoadbalanceExpirationCertificationResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetLoadbalanceExpirationCertificationResponse) String() string {
	return utils.Beautify(s)
}

func (s GetLoadbalanceExpirationCertificationResponse) GoString() string {
	return s.String()
}

func (s GetLoadbalanceExpirationCertificationResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetRequestId(v string) *GetLoadbalanceExpirationCertificationResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetErrorMessage(v string) *GetLoadbalanceExpirationCertificationResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetErrorCode(v string) *GetLoadbalanceExpirationCertificationResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetState(v GetLoadbalanceExpirationCertificationResponseStateEnum) *GetLoadbalanceExpirationCertificationResponse {
	s.State = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetBody(v []GetLoadbalanceExpirationCertificationResponseBody) *GetLoadbalanceExpirationCertificationResponse {
	s.Body = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponse) SetErrorParams(v []string) *GetLoadbalanceExpirationCertificationResponse {
	s.ErrorParams = v
	return s
}


type GetLoadbalanceExpirationCertificationResponseBuilder struct {
	s *GetLoadbalanceExpirationCertificationResponse
}

func NewGetLoadbalanceExpirationCertificationResponseBuilder() *GetLoadbalanceExpirationCertificationResponseBuilder {
	s := &GetLoadbalanceExpirationCertificationResponse{}
	b := &GetLoadbalanceExpirationCertificationResponseBuilder{s: s}
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) RequestId(v string) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) ErrorMessage(v string) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) ErrorCode(v string) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) State(v GetLoadbalanceExpirationCertificationResponseStateEnum) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) Body(v []GetLoadbalanceExpirationCertificationResponseBody) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) ErrorParams(v []string) *GetLoadbalanceExpirationCertificationResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBuilder) Build() *GetLoadbalanceExpirationCertificationResponse {
	return b.s
}


