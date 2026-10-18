// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestHttp2Status struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
}

func (s SetCdnDomainConfigRequestHttp2Status) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestHttp2Status) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestHttp2Status) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestHttp2Status) SetEnable(v bool) *SetCdnDomainConfigRequestHttp2Status {
	s.Enable = &v
	return s
}


type SetCdnDomainConfigRequestHttp2StatusBuilder struct {
	s *SetCdnDomainConfigRequestHttp2Status
}

func NewSetCdnDomainConfigRequestHttp2StatusBuilder() *SetCdnDomainConfigRequestHttp2StatusBuilder {
	s := &SetCdnDomainConfigRequestHttp2Status{}
	b := &SetCdnDomainConfigRequestHttp2StatusBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestHttp2StatusBuilder) Enable(v bool) *SetCdnDomainConfigRequestHttp2StatusBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestHttp2StatusBuilder) Build() *SetCdnDomainConfigRequestHttp2Status {
	return b.s
}


