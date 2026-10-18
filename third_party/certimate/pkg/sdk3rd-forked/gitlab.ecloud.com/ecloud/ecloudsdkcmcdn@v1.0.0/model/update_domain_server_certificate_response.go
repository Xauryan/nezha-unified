// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDomainServerCertificateResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *string `json:"body,omitempty"`
}

func (s UpdateDomainServerCertificateResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateDomainServerCertificateResponse) GoString() string {
	return s.String()
}

func (s UpdateDomainServerCertificateResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDomainServerCertificateResponse) SetRequestId(v string) *UpdateDomainServerCertificateResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainServerCertificateResponse) SetErrorMessage(v string) *UpdateDomainServerCertificateResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDomainServerCertificateResponse) SetErrorCode(v string) *UpdateDomainServerCertificateResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDomainServerCertificateResponse) SetState(v string) *UpdateDomainServerCertificateResponse {
	s.State = &v
	return s
}

func (s *UpdateDomainServerCertificateResponse) SetBody(v string) *UpdateDomainServerCertificateResponse {
	s.Body = &v
	return s
}


type UpdateDomainServerCertificateResponseBuilder struct {
	s *UpdateDomainServerCertificateResponse
}

func NewUpdateDomainServerCertificateResponseBuilder() *UpdateDomainServerCertificateResponseBuilder {
	s := &UpdateDomainServerCertificateResponse{}
	b := &UpdateDomainServerCertificateResponseBuilder{s: s}
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) RequestId(v string) *UpdateDomainServerCertificateResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) ErrorMessage(v string) *UpdateDomainServerCertificateResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) ErrorCode(v string) *UpdateDomainServerCertificateResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) State(v string) *UpdateDomainServerCertificateResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) Body(v string) *UpdateDomainServerCertificateResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateDomainServerCertificateResponseBuilder) Build() *UpdateDomainServerCertificateResponse {
	return b.s
}


