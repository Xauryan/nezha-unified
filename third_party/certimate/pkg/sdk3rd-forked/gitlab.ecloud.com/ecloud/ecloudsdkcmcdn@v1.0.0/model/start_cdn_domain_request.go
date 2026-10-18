// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartCdnDomainRequest struct {


	StartCdnDomainQuery *StartCdnDomainQuery `json:"startCdnDomainQuery,omitempty"`
}

func (s StartCdnDomainRequest) String() string {
	return utils.Beautify(s)
}

func (s StartCdnDomainRequest) GoString() string {
	return s.String()
}

func (s StartCdnDomainRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartCdnDomainRequest) SetStartCdnDomainQuery(v *StartCdnDomainQuery) *StartCdnDomainRequest {
	s.StartCdnDomainQuery = v
	return s
}


type StartCdnDomainRequestBuilder struct {
	s *StartCdnDomainRequest
}

func NewStartCdnDomainRequestBuilder() *StartCdnDomainRequestBuilder {
	s := &StartCdnDomainRequest{}
	b := &StartCdnDomainRequestBuilder{s: s}
	return b
}

func (b *StartCdnDomainRequestBuilder) StartCdnDomainQuery(v *StartCdnDomainQuery) *StartCdnDomainRequestBuilder {
	b.s.StartCdnDomainQuery = v
	return b
}

func (b *StartCdnDomainRequestBuilder) Build() *StartCdnDomainRequest {
	return b.s
}


