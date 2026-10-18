// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeCdnDomainDetailResponseRasLinkDetail struct {

    // 选择域名配置，取值为以下之一 1: 启用白名单 2: 启用黑名单
	ListType *int32 `json:"list_type,omitempty"`
    // Referer配置，取值为以下之一 true：允许Referer为空 false：禁止Referer为空
	ReferrerNull *bool `json:"referrer_null,omitempty"`
    // 白/黑名单的域名列表;支持*通配符;允许最多添加100条，域名之间用分号分隔
	DomainList *string `json:"domainList,omitempty"`
}

func (s DescribeCdnDomainDetailResponseRasLinkDetail) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseRasLinkDetail) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseRasLinkDetail) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseRasLinkDetail) SetListType(v int32) *DescribeCdnDomainDetailResponseRasLinkDetail {
	s.ListType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseRasLinkDetail) SetReferrerNull(v bool) *DescribeCdnDomainDetailResponseRasLinkDetail {
	s.ReferrerNull = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseRasLinkDetail) SetDomainList(v string) *DescribeCdnDomainDetailResponseRasLinkDetail {
	s.DomainList = &v
	return s
}


type DescribeCdnDomainDetailResponseRasLinkDetailBuilder struct {
	s *DescribeCdnDomainDetailResponseRasLinkDetail
}

func NewDescribeCdnDomainDetailResponseRasLinkDetailBuilder() *DescribeCdnDomainDetailResponseRasLinkDetailBuilder {
	s := &DescribeCdnDomainDetailResponseRasLinkDetail{}
	b := &DescribeCdnDomainDetailResponseRasLinkDetailBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseRasLinkDetailBuilder) ListType(v int32) *DescribeCdnDomainDetailResponseRasLinkDetailBuilder {
	b.s.ListType = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseRasLinkDetailBuilder) ReferrerNull(v bool) *DescribeCdnDomainDetailResponseRasLinkDetailBuilder {
	b.s.ReferrerNull = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseRasLinkDetailBuilder) DomainList(v string) *DescribeCdnDomainDetailResponseRasLinkDetailBuilder {
	b.s.DomainList = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseRasLinkDetailBuilder) Build() *DescribeCdnDomainDetailResponseRasLinkDetail {
	return b.s
}


