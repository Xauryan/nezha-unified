// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopCdnDomainRequest struct {


	StopCdnDomainQuery *StopCdnDomainQuery `json:"stopCdnDomainQuery,omitempty"`
}

func (s StopCdnDomainRequest) String() string {
	return utils.Beautify(s)
}

func (s StopCdnDomainRequest) GoString() string {
	return s.String()
}

func (s StopCdnDomainRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopCdnDomainRequest) SetStopCdnDomainQuery(v *StopCdnDomainQuery) *StopCdnDomainRequest {
	s.StopCdnDomainQuery = v
	return s
}


type StopCdnDomainRequestBuilder struct {
	s *StopCdnDomainRequest
}

func NewStopCdnDomainRequestBuilder() *StopCdnDomainRequestBuilder {
	s := &StopCdnDomainRequest{}
	b := &StopCdnDomainRequestBuilder{s: s}
	return b
}

func (b *StopCdnDomainRequestBuilder) StopCdnDomainQuery(v *StopCdnDomainQuery) *StopCdnDomainRequestBuilder {
	b.s.StopCdnDomainQuery = v
	return b
}

func (b *StopCdnDomainRequestBuilder) Build() *StopCdnDomainRequest {
	return b.s
}


