// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestDetail struct {

    // 选择域名配置，取值为以下之一 1: 启用白名单 2: 启用黑名单
	ListType *int32 `json:"list_type,omitempty"`
    // Referer配置，取值为以下之一 true：允许Referer为空 false：禁止Referer为空
	ReferrerNull *bool `json:"referrer_null,omitempty"`
    // 白/黑名单的域名列表;支持*通配符;允许最多添加100条
	DomainList *string `json:"domainList,omitempty"`
}

func (s SetCdnDomainConfigRequestDetail) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestDetail) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestDetail) SetListType(v int32) *SetCdnDomainConfigRequestDetail {
	s.ListType = &v
	return s
}

func (s *SetCdnDomainConfigRequestDetail) SetReferrerNull(v bool) *SetCdnDomainConfigRequestDetail {
	s.ReferrerNull = &v
	return s
}

func (s *SetCdnDomainConfigRequestDetail) SetDomainList(v string) *SetCdnDomainConfigRequestDetail {
	s.DomainList = &v
	return s
}


type SetCdnDomainConfigRequestDetailBuilder struct {
	s *SetCdnDomainConfigRequestDetail
}

func NewSetCdnDomainConfigRequestDetailBuilder() *SetCdnDomainConfigRequestDetailBuilder {
	s := &SetCdnDomainConfigRequestDetail{}
	b := &SetCdnDomainConfigRequestDetailBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestDetailBuilder) ListType(v int32) *SetCdnDomainConfigRequestDetailBuilder {
	b.s.ListType = &v
	return b
}

func (b *SetCdnDomainConfigRequestDetailBuilder) ReferrerNull(v bool) *SetCdnDomainConfigRequestDetailBuilder {
	b.s.ReferrerNull = &v
	return b
}

func (b *SetCdnDomainConfigRequestDetailBuilder) DomainList(v string) *SetCdnDomainConfigRequestDetailBuilder {
	b.s.DomainList = &v
	return b
}

func (b *SetCdnDomainConfigRequestDetailBuilder) Build() *SetCdnDomainConfigRequestDetail {
	return b.s
}


