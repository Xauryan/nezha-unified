// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum string

// List of BackProtocol
const (
    AddCdnDomainRequestCdnSrcDetailsBackProtocolEnumFlow AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum = "FLOW"
    AddCdnDomainRequestCdnSrcDetailsBackProtocolEnumHttp AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum = "HTTP"
    AddCdnDomainRequestCdnSrcDetailsBackProtocolEnumHttps AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum = "HTTPS"
)
type AddCdnDomainRequestCdnSrcDetailsBackTypeEnum string

// List of BackType
const (
    AddCdnDomainRequestCdnSrcDetailsBackTypeEnumIp AddCdnDomainRequestCdnSrcDetailsBackTypeEnum = "ip"
    AddCdnDomainRequestCdnSrcDetailsBackTypeEnumDomain AddCdnDomainRequestCdnSrcDetailsBackTypeEnum = "domain"
    AddCdnDomainRequestCdnSrcDetailsBackTypeEnumEos AddCdnDomainRequestCdnSrcDetailsBackTypeEnum = "eos"
)

type AddCdnDomainRequestCdnSrcDetails struct {

    // 回源协议
	BackProtocol *AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum `json:"backProtocol,omitempty"`
    // 回源地址，ip或域名
	BackAddr *string `json:"backAddr,omitempty"`
    // 回源地址类型
	BackType *AddCdnDomainRequestCdnSrcDetailsBackTypeEnum `json:"backType,omitempty"`
    // 是否为主站
	Primary *bool `json:"primary,omitempty"`
}

func (s AddCdnDomainRequestCdnSrcDetails) String() string {
	return utils.Beautify(s)
}

func (s AddCdnDomainRequestCdnSrcDetails) GoString() string {
	return s.String()
}

func (s AddCdnDomainRequestCdnSrcDetails) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddCdnDomainRequestCdnSrcDetails) SetBackProtocol(v AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum) *AddCdnDomainRequestCdnSrcDetails {
	s.BackProtocol = &v
	return s
}

func (s *AddCdnDomainRequestCdnSrcDetails) SetBackAddr(v string) *AddCdnDomainRequestCdnSrcDetails {
	s.BackAddr = &v
	return s
}

func (s *AddCdnDomainRequestCdnSrcDetails) SetBackType(v AddCdnDomainRequestCdnSrcDetailsBackTypeEnum) *AddCdnDomainRequestCdnSrcDetails {
	s.BackType = &v
	return s
}

func (s *AddCdnDomainRequestCdnSrcDetails) SetPrimary(v bool) *AddCdnDomainRequestCdnSrcDetails {
	s.Primary = &v
	return s
}


type AddCdnDomainRequestCdnSrcDetailsBuilder struct {
	s *AddCdnDomainRequestCdnSrcDetails
}

func NewAddCdnDomainRequestCdnSrcDetailsBuilder() *AddCdnDomainRequestCdnSrcDetailsBuilder {
	s := &AddCdnDomainRequestCdnSrcDetails{}
	b := &AddCdnDomainRequestCdnSrcDetailsBuilder{s: s}
	return b
}

func (b *AddCdnDomainRequestCdnSrcDetailsBuilder) BackProtocol(v AddCdnDomainRequestCdnSrcDetailsBackProtocolEnum) *AddCdnDomainRequestCdnSrcDetailsBuilder {
	b.s.BackProtocol = &v
	return b
}

func (b *AddCdnDomainRequestCdnSrcDetailsBuilder) BackAddr(v string) *AddCdnDomainRequestCdnSrcDetailsBuilder {
	b.s.BackAddr = &v
	return b
}

func (b *AddCdnDomainRequestCdnSrcDetailsBuilder) BackType(v AddCdnDomainRequestCdnSrcDetailsBackTypeEnum) *AddCdnDomainRequestCdnSrcDetailsBuilder {
	b.s.BackType = &v
	return b
}

func (b *AddCdnDomainRequestCdnSrcDetailsBuilder) Primary(v bool) *AddCdnDomainRequestCdnSrcDetailsBuilder {
	b.s.Primary = &v
	return b
}

func (b *AddCdnDomainRequestCdnSrcDetailsBuilder) Build() *AddCdnDomainRequestCdnSrcDetails {
	return b.s
}


