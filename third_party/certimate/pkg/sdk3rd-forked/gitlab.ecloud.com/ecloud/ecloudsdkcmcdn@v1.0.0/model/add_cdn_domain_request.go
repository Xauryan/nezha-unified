// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddCdnDomainRequest struct {


	AddCdnDomainBody *AddCdnDomainBody `json:"addCdnDomainBody,omitempty"`
}

func (s AddCdnDomainRequest) String() string {
	return utils.Beautify(s)
}

func (s AddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s AddCdnDomainRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddCdnDomainRequest) SetAddCdnDomainBody(v *AddCdnDomainBody) *AddCdnDomainRequest {
	s.AddCdnDomainBody = v
	return s
}


type AddCdnDomainRequestBuilder struct {
	s *AddCdnDomainRequest
}

func NewAddCdnDomainRequestBuilder() *AddCdnDomainRequestBuilder {
	s := &AddCdnDomainRequest{}
	b := &AddCdnDomainRequestBuilder{s: s}
	return b
}

func (b *AddCdnDomainRequestBuilder) AddCdnDomainBody(v *AddCdnDomainBody) *AddCdnDomainRequestBuilder {
	b.s.AddCdnDomainBody = v
	return b
}

func (b *AddCdnDomainRequestBuilder) Build() *AddCdnDomainRequest {
	return b.s
}


