// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDomainServerCertificateResponse struct {

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

func (s DeleteDomainServerCertificateResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteDomainServerCertificateResponse) GoString() string {
	return s.String()
}

func (s DeleteDomainServerCertificateResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDomainServerCertificateResponse) SetRequestId(v string) *DeleteDomainServerCertificateResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainServerCertificateResponse) SetErrorMessage(v string) *DeleteDomainServerCertificateResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDomainServerCertificateResponse) SetErrorCode(v string) *DeleteDomainServerCertificateResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDomainServerCertificateResponse) SetState(v string) *DeleteDomainServerCertificateResponse {
	s.State = &v
	return s
}

func (s *DeleteDomainServerCertificateResponse) SetBody(v string) *DeleteDomainServerCertificateResponse {
	s.Body = &v
	return s
}


type DeleteDomainServerCertificateResponseBuilder struct {
	s *DeleteDomainServerCertificateResponse
}

func NewDeleteDomainServerCertificateResponseBuilder() *DeleteDomainServerCertificateResponseBuilder {
	s := &DeleteDomainServerCertificateResponse{}
	b := &DeleteDomainServerCertificateResponseBuilder{s: s}
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) RequestId(v string) *DeleteDomainServerCertificateResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) ErrorMessage(v string) *DeleteDomainServerCertificateResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) ErrorCode(v string) *DeleteDomainServerCertificateResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) State(v string) *DeleteDomainServerCertificateResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) Body(v string) *DeleteDomainServerCertificateResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteDomainServerCertificateResponseBuilder) Build() *DeleteDomainServerCertificateResponse {
	return b.s
}


