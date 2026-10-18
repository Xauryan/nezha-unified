// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CertificationUnbindTagRequest struct {


	CertificationUnbindTagPath *CertificationUnbindTagPath `json:"certificationUnbindTagPath,omitempty"`
}

func (s CertificationUnbindTagRequest) String() string {
	return utils.Beautify(s)
}

func (s CertificationUnbindTagRequest) GoString() string {
	return s.String()
}

func (s CertificationUnbindTagRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationUnbindTagRequest) SetCertificationUnbindTagPath(v *CertificationUnbindTagPath) *CertificationUnbindTagRequest {
	s.CertificationUnbindTagPath = v
	return s
}


type CertificationUnbindTagRequestBuilder struct {
	s *CertificationUnbindTagRequest
}

func NewCertificationUnbindTagRequestBuilder() *CertificationUnbindTagRequestBuilder {
	s := &CertificationUnbindTagRequest{}
	b := &CertificationUnbindTagRequestBuilder{s: s}
	return b
}

func (b *CertificationUnbindTagRequestBuilder) CertificationUnbindTagPath(v *CertificationUnbindTagPath) *CertificationUnbindTagRequestBuilder {
	b.s.CertificationUnbindTagPath = v
	return b
}

func (b *CertificationUnbindTagRequestBuilder) Build() *CertificationUnbindTagRequest {
	return b.s
}


