// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CertificationBindTagsRequest struct {


	CertificationBindTagsBody *CertificationBindTagsBody `json:"certificationBindTagsBody,omitempty"`
}

func (s CertificationBindTagsRequest) String() string {
	return utils.Beautify(s)
}

func (s CertificationBindTagsRequest) GoString() string {
	return s.String()
}

func (s CertificationBindTagsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationBindTagsRequest) SetCertificationBindTagsBody(v *CertificationBindTagsBody) *CertificationBindTagsRequest {
	s.CertificationBindTagsBody = v
	return s
}


type CertificationBindTagsRequestBuilder struct {
	s *CertificationBindTagsRequest
}

func NewCertificationBindTagsRequestBuilder() *CertificationBindTagsRequestBuilder {
	s := &CertificationBindTagsRequest{}
	b := &CertificationBindTagsRequestBuilder{s: s}
	return b
}

func (b *CertificationBindTagsRequestBuilder) CertificationBindTagsBody(v *CertificationBindTagsBody) *CertificationBindTagsRequestBuilder {
	b.s.CertificationBindTagsBody = v
	return b
}

func (b *CertificationBindTagsRequestBuilder) Build() *CertificationBindTagsRequest {
	return b.s
}


