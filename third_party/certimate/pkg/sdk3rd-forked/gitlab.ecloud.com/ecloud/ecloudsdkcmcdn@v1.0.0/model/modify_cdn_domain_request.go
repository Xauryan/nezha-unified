// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyCdnDomainRequest struct {


	ModifyCdnDomainBody *ModifyCdnDomainBody `json:"modifyCdnDomainBody,omitempty"`
}

func (s ModifyCdnDomainRequest) String() string {
	return utils.Beautify(s)
}

func (s ModifyCdnDomainRequest) GoString() string {
	return s.String()
}

func (s ModifyCdnDomainRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyCdnDomainRequest) SetModifyCdnDomainBody(v *ModifyCdnDomainBody) *ModifyCdnDomainRequest {
	s.ModifyCdnDomainBody = v
	return s
}


type ModifyCdnDomainRequestBuilder struct {
	s *ModifyCdnDomainRequest
}

func NewModifyCdnDomainRequestBuilder() *ModifyCdnDomainRequestBuilder {
	s := &ModifyCdnDomainRequest{}
	b := &ModifyCdnDomainRequestBuilder{s: s}
	return b
}

func (b *ModifyCdnDomainRequestBuilder) ModifyCdnDomainBody(v *ModifyCdnDomainBody) *ModifyCdnDomainRequestBuilder {
	b.s.ModifyCdnDomainBody = v
	return b
}

func (b *ModifyCdnDomainRequestBuilder) Build() *ModifyCdnDomainRequest {
	return b.s
}


