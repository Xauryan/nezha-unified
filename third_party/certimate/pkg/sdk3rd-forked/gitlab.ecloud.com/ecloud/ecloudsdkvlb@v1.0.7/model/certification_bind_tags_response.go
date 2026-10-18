// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CertificationBindTagsResponseStateEnum string

// List of State
const (
    CertificationBindTagsResponseStateEnumOk CertificationBindTagsResponseStateEnum = "OK"
    CertificationBindTagsResponseStateEnumError CertificationBindTagsResponseStateEnum = "ERROR"
    CertificationBindTagsResponseStateEnumException CertificationBindTagsResponseStateEnum = "EXCEPTION"
    CertificationBindTagsResponseStateEnumAlarm CertificationBindTagsResponseStateEnum = "ALARM"
    CertificationBindTagsResponseStateEnumForbidden CertificationBindTagsResponseStateEnum = "FORBIDDEN"
)

type CertificationBindTagsResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *CertificationBindTagsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s CertificationBindTagsResponse) String() string {
	return utils.Beautify(s)
}

func (s CertificationBindTagsResponse) GoString() string {
	return s.String()
}

func (s CertificationBindTagsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationBindTagsResponse) SetAsyncKey(v string) *CertificationBindTagsResponse {
	s.AsyncKey = &v
	return s
}

func (s *CertificationBindTagsResponse) SetRequestId(v string) *CertificationBindTagsResponse {
	s.RequestId = &v
	return s
}

func (s *CertificationBindTagsResponse) SetErrorMessage(v string) *CertificationBindTagsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CertificationBindTagsResponse) SetErrorCode(v string) *CertificationBindTagsResponse {
	s.ErrorCode = &v
	return s
}

func (s *CertificationBindTagsResponse) SetState(v CertificationBindTagsResponseStateEnum) *CertificationBindTagsResponse {
	s.State = &v
	return s
}

func (s *CertificationBindTagsResponse) SetBody(v string) *CertificationBindTagsResponse {
	s.Body = &v
	return s
}

func (s *CertificationBindTagsResponse) SetErrorParams(v []string) *CertificationBindTagsResponse {
	s.ErrorParams = v
	return s
}


type CertificationBindTagsResponseBuilder struct {
	s *CertificationBindTagsResponse
}

func NewCertificationBindTagsResponseBuilder() *CertificationBindTagsResponseBuilder {
	s := &CertificationBindTagsResponse{}
	b := &CertificationBindTagsResponseBuilder{s: s}
	return b
}

func (b *CertificationBindTagsResponseBuilder) AsyncKey(v string) *CertificationBindTagsResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) RequestId(v string) *CertificationBindTagsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) ErrorMessage(v string) *CertificationBindTagsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) ErrorCode(v string) *CertificationBindTagsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) State(v CertificationBindTagsResponseStateEnum) *CertificationBindTagsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) Body(v string) *CertificationBindTagsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CertificationBindTagsResponseBuilder) ErrorParams(v []string) *CertificationBindTagsResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *CertificationBindTagsResponseBuilder) Build() *CertificationBindTagsResponse {
	return b.s
}


