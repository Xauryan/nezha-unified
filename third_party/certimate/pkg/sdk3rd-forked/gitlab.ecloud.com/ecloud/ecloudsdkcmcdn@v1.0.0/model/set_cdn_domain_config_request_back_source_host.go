// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetCdnDomainConfigRequestBackSourceHostTypeEnum string

// List of Type
const (
    SetCdnDomainConfigRequestBackSourceHostTypeEnumAccelerated SetCdnDomainConfigRequestBackSourceHostTypeEnum = "ACCELERATED"
    SetCdnDomainConfigRequestBackSourceHostTypeEnumSource SetCdnDomainConfigRequestBackSourceHostTypeEnum = "SOURCE"
    SetCdnDomainConfigRequestBackSourceHostTypeEnumCustomize SetCdnDomainConfigRequestBackSourceHostTypeEnum = "CUSTOMIZE"
)

type SetCdnDomainConfigRequestBackSourceHost struct {

    // 回源host配置是否开启，不开启，默认host为加速域名
	Enable *bool `json:"enable,omitempty"`
    // 回源host
	Host *string `json:"host,omitempty"`
    // 域名类型:0:加速域名，1:源站域名，2:自定义域名
	Type *SetCdnDomainConfigRequestBackSourceHostTypeEnum `json:"type,omitempty"`
}

func (s SetCdnDomainConfigRequestBackSourceHost) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestBackSourceHost) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestBackSourceHost) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestBackSourceHost) SetEnable(v bool) *SetCdnDomainConfigRequestBackSourceHost {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestBackSourceHost) SetHost(v string) *SetCdnDomainConfigRequestBackSourceHost {
	s.Host = &v
	return s
}

func (s *SetCdnDomainConfigRequestBackSourceHost) SetType(v SetCdnDomainConfigRequestBackSourceHostTypeEnum) *SetCdnDomainConfigRequestBackSourceHost {
	s.Type = &v
	return s
}


type SetCdnDomainConfigRequestBackSourceHostBuilder struct {
	s *SetCdnDomainConfigRequestBackSourceHost
}

func NewSetCdnDomainConfigRequestBackSourceHostBuilder() *SetCdnDomainConfigRequestBackSourceHostBuilder {
	s := &SetCdnDomainConfigRequestBackSourceHost{}
	b := &SetCdnDomainConfigRequestBackSourceHostBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestBackSourceHostBuilder) Enable(v bool) *SetCdnDomainConfigRequestBackSourceHostBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestBackSourceHostBuilder) Host(v string) *SetCdnDomainConfigRequestBackSourceHostBuilder {
	b.s.Host = &v
	return b
}

func (b *SetCdnDomainConfigRequestBackSourceHostBuilder) Type(v SetCdnDomainConfigRequestBackSourceHostTypeEnum) *SetCdnDomainConfigRequestBackSourceHostBuilder {
	b.s.Type = &v
	return b
}

func (b *SetCdnDomainConfigRequestBackSourceHostBuilder) Build() *SetCdnDomainConfigRequestBackSourceHost {
	return b.s
}


