// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDomainServerCertificateRequest struct {


	DeleteDomainServerCertificatePath *DeleteDomainServerCertificatePath `json:"deleteDomainServerCertificatePath,omitempty"`
}

func (s DeleteDomainServerCertificateRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteDomainServerCertificateRequest) GoString() string {
	return s.String()
}

func (s DeleteDomainServerCertificateRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDomainServerCertificateRequest) SetDeleteDomainServerCertificatePath(v *DeleteDomainServerCertificatePath) *DeleteDomainServerCertificateRequest {
	s.DeleteDomainServerCertificatePath = v
	return s
}


type DeleteDomainServerCertificateRequestBuilder struct {
	s *DeleteDomainServerCertificateRequest
}

func NewDeleteDomainServerCertificateRequestBuilder() *DeleteDomainServerCertificateRequestBuilder {
	s := &DeleteDomainServerCertificateRequest{}
	b := &DeleteDomainServerCertificateRequestBuilder{s: s}
	return b
}

func (b *DeleteDomainServerCertificateRequestBuilder) DeleteDomainServerCertificatePath(v *DeleteDomainServerCertificatePath) *DeleteDomainServerCertificateRequestBuilder {
	b.s.DeleteDomainServerCertificatePath = v
	return b
}

func (b *DeleteDomainServerCertificateRequestBuilder) Build() *DeleteDomainServerCertificateRequest {
	return b.s
}


