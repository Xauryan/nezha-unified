// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DescribeCdnDomainDetailResponseSourcesBackProtocolEnum string

// List of BackProtocol
const (
    DescribeCdnDomainDetailResponseSourcesBackProtocolEnumFlow DescribeCdnDomainDetailResponseSourcesBackProtocolEnum = "FLOW"
    DescribeCdnDomainDetailResponseSourcesBackProtocolEnumHttp DescribeCdnDomainDetailResponseSourcesBackProtocolEnum = "HTTP"
    DescribeCdnDomainDetailResponseSourcesBackProtocolEnumHttps DescribeCdnDomainDetailResponseSourcesBackProtocolEnum = "HTTPS"
)
type DescribeCdnDomainDetailResponseSourcesBackTypeEnum string

// List of BackType
const (
    DescribeCdnDomainDetailResponseSourcesBackTypeEnumIp DescribeCdnDomainDetailResponseSourcesBackTypeEnum = "ip"
    DescribeCdnDomainDetailResponseSourcesBackTypeEnumDomain DescribeCdnDomainDetailResponseSourcesBackTypeEnum = "domain"
    DescribeCdnDomainDetailResponseSourcesBackTypeEnumEos DescribeCdnDomainDetailResponseSourcesBackTypeEnum = "eos"
)

type DescribeCdnDomainDetailResponseSources struct {

    // 回源协议
	BackProtocol *DescribeCdnDomainDetailResponseSourcesBackProtocolEnum `json:"backProtocol,omitempty"`
    // 回源配置id
	SourceId *int32 `json:"sourceId,omitempty"`
    // 回源地址，ip或域名
	BackAddr *string `json:"backAddr,omitempty"`
    // 回源地址类型
	BackType *DescribeCdnDomainDetailResponseSourcesBackTypeEnum `json:"backType,omitempty"`
    // 是否为主站
	Primary *bool `json:"primary,omitempty"`
}

func (s DescribeCdnDomainDetailResponseSources) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseSources) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseSources) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseSources) SetBackProtocol(v DescribeCdnDomainDetailResponseSourcesBackProtocolEnum) *DescribeCdnDomainDetailResponseSources {
	s.BackProtocol = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseSources) SetSourceId(v int32) *DescribeCdnDomainDetailResponseSources {
	s.SourceId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseSources) SetBackAddr(v string) *DescribeCdnDomainDetailResponseSources {
	s.BackAddr = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseSources) SetBackType(v DescribeCdnDomainDetailResponseSourcesBackTypeEnum) *DescribeCdnDomainDetailResponseSources {
	s.BackType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseSources) SetPrimary(v bool) *DescribeCdnDomainDetailResponseSources {
	s.Primary = &v
	return s
}


type DescribeCdnDomainDetailResponseSourcesBuilder struct {
	s *DescribeCdnDomainDetailResponseSources
}

func NewDescribeCdnDomainDetailResponseSourcesBuilder() *DescribeCdnDomainDetailResponseSourcesBuilder {
	s := &DescribeCdnDomainDetailResponseSources{}
	b := &DescribeCdnDomainDetailResponseSourcesBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) BackProtocol(v DescribeCdnDomainDetailResponseSourcesBackProtocolEnum) *DescribeCdnDomainDetailResponseSourcesBuilder {
	b.s.BackProtocol = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) SourceId(v int32) *DescribeCdnDomainDetailResponseSourcesBuilder {
	b.s.SourceId = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) BackAddr(v string) *DescribeCdnDomainDetailResponseSourcesBuilder {
	b.s.BackAddr = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) BackType(v DescribeCdnDomainDetailResponseSourcesBackTypeEnum) *DescribeCdnDomainDetailResponseSourcesBuilder {
	b.s.BackType = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) Primary(v bool) *DescribeCdnDomainDetailResponseSourcesBuilder {
	b.s.Primary = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseSourcesBuilder) Build() *DescribeCdnDomainDetailResponseSources {
	return b.s
}


