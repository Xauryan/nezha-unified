// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum string

// List of HostType
const (
    DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnumAccelerated DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum = "ACCELERATED"
    DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnumSource DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum = "SOURCE"
    DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnumCustomize DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum = "CUSTOMIZE"
)

type DescribeCdnDomainDetailResponseBackSourceHost struct {

    // 域名类型:0:加速域名，1:源站域名，2:自定义域名
	HostType *DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum `json:"hostType,omitempty"`
    // 回源host
	Host *string `json:"host,omitempty"`
    // 回源host配置id
	HostId *int32 `json:"hostId,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBackSourceHost) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseBackSourceHost) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseBackSourceHost) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseBackSourceHost) SetHostType(v DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum) *DescribeCdnDomainDetailResponseBackSourceHost {
	s.HostType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBackSourceHost) SetHost(v string) *DescribeCdnDomainDetailResponseBackSourceHost {
	s.Host = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBackSourceHost) SetHostId(v int32) *DescribeCdnDomainDetailResponseBackSourceHost {
	s.HostId = &v
	return s
}


type DescribeCdnDomainDetailResponseBackSourceHostBuilder struct {
	s *DescribeCdnDomainDetailResponseBackSourceHost
}

func NewDescribeCdnDomainDetailResponseBackSourceHostBuilder() *DescribeCdnDomainDetailResponseBackSourceHostBuilder {
	s := &DescribeCdnDomainDetailResponseBackSourceHost{}
	b := &DescribeCdnDomainDetailResponseBackSourceHostBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseBackSourceHostBuilder) HostType(v DescribeCdnDomainDetailResponseBackSourceHostHostTypeEnum) *DescribeCdnDomainDetailResponseBackSourceHostBuilder {
	b.s.HostType = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBackSourceHostBuilder) Host(v string) *DescribeCdnDomainDetailResponseBackSourceHostBuilder {
	b.s.Host = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBackSourceHostBuilder) HostId(v int32) *DescribeCdnDomainDetailResponseBackSourceHostBuilder {
	b.s.HostId = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBackSourceHostBuilder) Build() *DescribeCdnDomainDetailResponseBackSourceHost {
	return b.s
}


