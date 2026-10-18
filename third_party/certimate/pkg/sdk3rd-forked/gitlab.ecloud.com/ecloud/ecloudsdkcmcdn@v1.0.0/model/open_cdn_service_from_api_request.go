// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OpenCdnServiceFromApiRequest struct {


	OpenCdnServiceFromApiBody *OpenCdnServiceFromApiBody `json:"openCdnServiceFromApiBody,omitempty"`
}

func (s OpenCdnServiceFromApiRequest) String() string {
	return utils.Beautify(s)
}

func (s OpenCdnServiceFromApiRequest) GoString() string {
	return s.String()
}

func (s OpenCdnServiceFromApiRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OpenCdnServiceFromApiRequest) SetOpenCdnServiceFromApiBody(v *OpenCdnServiceFromApiBody) *OpenCdnServiceFromApiRequest {
	s.OpenCdnServiceFromApiBody = v
	return s
}


type OpenCdnServiceFromApiRequestBuilder struct {
	s *OpenCdnServiceFromApiRequest
}

func NewOpenCdnServiceFromApiRequestBuilder() *OpenCdnServiceFromApiRequestBuilder {
	s := &OpenCdnServiceFromApiRequest{}
	b := &OpenCdnServiceFromApiRequestBuilder{s: s}
	return b
}

func (b *OpenCdnServiceFromApiRequestBuilder) OpenCdnServiceFromApiBody(v *OpenCdnServiceFromApiBody) *OpenCdnServiceFromApiRequestBuilder {
	b.s.OpenCdnServiceFromApiBody = v
	return b
}

func (b *OpenCdnServiceFromApiRequestBuilder) Build() *OpenCdnServiceFromApiRequest {
	return b.s
}


