// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ModifyCdnDomainRequestSourcesBackProtocolEnum string

// List of BackProtocol
const (
    ModifyCdnDomainRequestSourcesBackProtocolEnumFlow ModifyCdnDomainRequestSourcesBackProtocolEnum = "FLOW"
    ModifyCdnDomainRequestSourcesBackProtocolEnumHttp ModifyCdnDomainRequestSourcesBackProtocolEnum = "HTTP"
    ModifyCdnDomainRequestSourcesBackProtocolEnumHttps ModifyCdnDomainRequestSourcesBackProtocolEnum = "HTTPS"
)
type ModifyCdnDomainRequestSourcesBackTypeEnum string

// List of BackType
const (
    ModifyCdnDomainRequestSourcesBackTypeEnumIp ModifyCdnDomainRequestSourcesBackTypeEnum = "ip"
    ModifyCdnDomainRequestSourcesBackTypeEnumDomain ModifyCdnDomainRequestSourcesBackTypeEnum = "domain"
    ModifyCdnDomainRequestSourcesBackTypeEnumEos ModifyCdnDomainRequestSourcesBackTypeEnum = "eos"
)

type ModifyCdnDomainRequestSources struct {

    // 回源协议
	BackProtocol *ModifyCdnDomainRequestSourcesBackProtocolEnum `json:"backProtocol,omitempty"`
    // 回源地址，ip或域名
	BackAddr *string `json:"backAddr,omitempty"`
    // 回源地址类型
	BackType *ModifyCdnDomainRequestSourcesBackTypeEnum `json:"backType,omitempty"`
    // 是否为主站
	Primary *bool `json:"primary,omitempty"`
}

func (s ModifyCdnDomainRequestSources) String() string {
	return utils.Beautify(s)
}

func (s ModifyCdnDomainRequestSources) GoString() string {
	return s.String()
}

func (s ModifyCdnDomainRequestSources) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyCdnDomainRequestSources) SetBackProtocol(v ModifyCdnDomainRequestSourcesBackProtocolEnum) *ModifyCdnDomainRequestSources {
	s.BackProtocol = &v
	return s
}

func (s *ModifyCdnDomainRequestSources) SetBackAddr(v string) *ModifyCdnDomainRequestSources {
	s.BackAddr = &v
	return s
}

func (s *ModifyCdnDomainRequestSources) SetBackType(v ModifyCdnDomainRequestSourcesBackTypeEnum) *ModifyCdnDomainRequestSources {
	s.BackType = &v
	return s
}

func (s *ModifyCdnDomainRequestSources) SetPrimary(v bool) *ModifyCdnDomainRequestSources {
	s.Primary = &v
	return s
}


type ModifyCdnDomainRequestSourcesBuilder struct {
	s *ModifyCdnDomainRequestSources
}

func NewModifyCdnDomainRequestSourcesBuilder() *ModifyCdnDomainRequestSourcesBuilder {
	s := &ModifyCdnDomainRequestSources{}
	b := &ModifyCdnDomainRequestSourcesBuilder{s: s}
	return b
}

func (b *ModifyCdnDomainRequestSourcesBuilder) BackProtocol(v ModifyCdnDomainRequestSourcesBackProtocolEnum) *ModifyCdnDomainRequestSourcesBuilder {
	b.s.BackProtocol = &v
	return b
}

func (b *ModifyCdnDomainRequestSourcesBuilder) BackAddr(v string) *ModifyCdnDomainRequestSourcesBuilder {
	b.s.BackAddr = &v
	return b
}

func (b *ModifyCdnDomainRequestSourcesBuilder) BackType(v ModifyCdnDomainRequestSourcesBackTypeEnum) *ModifyCdnDomainRequestSourcesBuilder {
	b.s.BackType = &v
	return b
}

func (b *ModifyCdnDomainRequestSourcesBuilder) Primary(v bool) *ModifyCdnDomainRequestSourcesBuilder {
	b.s.Primary = &v
	return b
}

func (b *ModifyCdnDomainRequestSourcesBuilder) Build() *ModifyCdnDomainRequestSources {
	return b.s
}


