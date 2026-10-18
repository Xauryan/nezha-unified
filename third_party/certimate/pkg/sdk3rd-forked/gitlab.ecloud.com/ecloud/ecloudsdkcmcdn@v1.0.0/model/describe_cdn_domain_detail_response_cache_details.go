// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DescribeCdnDomainDetailResponseCacheDetailsUnitEnum string

// List of Unit
const (
    DescribeCdnDomainDetailResponseCacheDetailsUnitEnumMonth DescribeCdnDomainDetailResponseCacheDetailsUnitEnum = "month"
    DescribeCdnDomainDetailResponseCacheDetailsUnitEnumDay DescribeCdnDomainDetailResponseCacheDetailsUnitEnum = "day"
    DescribeCdnDomainDetailResponseCacheDetailsUnitEnumHour DescribeCdnDomainDetailResponseCacheDetailsUnitEnum = "hour"
    DescribeCdnDomainDetailResponseCacheDetailsUnitEnumMinute DescribeCdnDomainDetailResponseCacheDetailsUnitEnum = "minute"
    DescribeCdnDomainDetailResponseCacheDetailsUnitEnumSecond DescribeCdnDomainDetailResponseCacheDetailsUnitEnum = "second"
)
type DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum string

// List of CacheType
const (
    DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnumCatalog DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum = "CATALOG"
    DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnumFileExtension DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum = "FILE_EXTENSION"
    DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnumGlobal DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum = "GLOBAL"
)

type DescribeCdnDomainDetailResponseCacheDetails struct {

    // 过期时间
	Unit *DescribeCdnDomainDetailResponseCacheDetailsUnitEnum `json:"unit,omitempty"`
    // 同一个域名多个缓存规则的优先级，权重越高优先级越高。
	Weight *int32 `json:"weight,omitempty"`
    // 缓存类型
	CacheType *DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum `json:"cacheType,omitempty"`
    // 过期时间。缓存策略设置为0，即不缓存。
	Ttl *int32 `json:"ttl,omitempty"`
    // 具体设置过期时间的缓存内容
	Content *string `json:"content,omitempty"`
    // 缓存id
	CacheId *int32 `json:"cacheId,omitempty"`
}

func (s DescribeCdnDomainDetailResponseCacheDetails) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseCacheDetails) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseCacheDetails) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetUnit(v DescribeCdnDomainDetailResponseCacheDetailsUnitEnum) *DescribeCdnDomainDetailResponseCacheDetails {
	s.Unit = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetWeight(v int32) *DescribeCdnDomainDetailResponseCacheDetails {
	s.Weight = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetCacheType(v DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum) *DescribeCdnDomainDetailResponseCacheDetails {
	s.CacheType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetTtl(v int32) *DescribeCdnDomainDetailResponseCacheDetails {
	s.Ttl = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetContent(v string) *DescribeCdnDomainDetailResponseCacheDetails {
	s.Content = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseCacheDetails) SetCacheId(v int32) *DescribeCdnDomainDetailResponseCacheDetails {
	s.CacheId = &v
	return s
}


type DescribeCdnDomainDetailResponseCacheDetailsBuilder struct {
	s *DescribeCdnDomainDetailResponseCacheDetails
}

func NewDescribeCdnDomainDetailResponseCacheDetailsBuilder() *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	s := &DescribeCdnDomainDetailResponseCacheDetails{}
	b := &DescribeCdnDomainDetailResponseCacheDetailsBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) Unit(v DescribeCdnDomainDetailResponseCacheDetailsUnitEnum) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.Unit = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) Weight(v int32) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.Weight = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) CacheType(v DescribeCdnDomainDetailResponseCacheDetailsCacheTypeEnum) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.CacheType = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) Ttl(v int32) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.Ttl = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) Content(v string) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.Content = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) CacheId(v int32) *DescribeCdnDomainDetailResponseCacheDetailsBuilder {
	b.s.CacheId = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseCacheDetailsBuilder) Build() *DescribeCdnDomainDetailResponseCacheDetails {
	return b.s
}


