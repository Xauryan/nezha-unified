// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDomainServerCertificateRequest struct {


	UpdateDomainServerCertificateHeader *UpdateDomainServerCertificateHeader `json:"updateDomainServerCertificateHeader,omitempty"`

	UpdateDomainServerCertificateBody *UpdateDomainServerCertificateBody `json:"updateDomainServerCertificateBody,omitempty"`
}

func (s UpdateDomainServerCertificateRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateDomainServerCertificateRequest) GoString() string {
	return s.String()
}

func (s UpdateDomainServerCertificateRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDomainServerCertificateRequest) SetUpdateDomainServerCertificateHeader(v *UpdateDomainServerCertificateHeader) *UpdateDomainServerCertificateRequest {
	s.UpdateDomainServerCertificateHeader = v
	return s
}

func (s *UpdateDomainServerCertificateRequest) SetUpdateDomainServerCertificateBody(v *UpdateDomainServerCertificateBody) *UpdateDomainServerCertificateRequest {
	s.UpdateDomainServerCertificateBody = v
	return s
}


type UpdateDomainServerCertificateRequestBuilder struct {
	s *UpdateDomainServerCertificateRequest
}

func NewUpdateDomainServerCertificateRequestBuilder() *UpdateDomainServerCertificateRequestBuilder {
	s := &UpdateDomainServerCertificateRequest{}
	b := &UpdateDomainServerCertificateRequestBuilder{s: s}
	return b
}

func (b *UpdateDomainServerCertificateRequestBuilder) UpdateDomainServerCertificateHeader(v *UpdateDomainServerCertificateHeader) *UpdateDomainServerCertificateRequestBuilder {
	b.s.UpdateDomainServerCertificateHeader = v
	return b
}

func (b *UpdateDomainServerCertificateRequestBuilder) UpdateDomainServerCertificateBody(v *UpdateDomainServerCertificateBody) *UpdateDomainServerCertificateRequestBuilder {
	b.s.UpdateDomainServerCertificateBody = v
	return b
}

func (b *UpdateDomainServerCertificateRequestBuilder) Build() *UpdateDomainServerCertificateRequest {
	return b.s
}


