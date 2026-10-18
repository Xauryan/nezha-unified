// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequest struct {


	SetCdnDomainConfigBody *SetCdnDomainConfigBody `json:"setCdnDomainConfigBody,omitempty"`
}

func (s SetCdnDomainConfigRequest) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequest) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequest) SetSetCdnDomainConfigBody(v *SetCdnDomainConfigBody) *SetCdnDomainConfigRequest {
	s.SetCdnDomainConfigBody = v
	return s
}


type SetCdnDomainConfigRequestBuilder struct {
	s *SetCdnDomainConfigRequest
}

func NewSetCdnDomainConfigRequestBuilder() *SetCdnDomainConfigRequestBuilder {
	s := &SetCdnDomainConfigRequest{}
	b := &SetCdnDomainConfigRequestBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestBuilder) SetCdnDomainConfigBody(v *SetCdnDomainConfigBody) *SetCdnDomainConfigRequestBuilder {
	b.s.SetCdnDomainConfigBody = v
	return b
}

func (b *SetCdnDomainConfigRequestBuilder) Build() *SetCdnDomainConfigRequest {
	return b.s
}


