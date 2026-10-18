// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserCrtListResponseList struct {

    // 证书联系人手机号
	ContactMobile *string `json:"contactMobile,omitempty"`
    // 绑定的域名id
	DomainIds *string `json:"domainIds,omitempty"`
    // 证书私钥
	PrivateKey *string `json:"privateKey,omitempty"`
    // 证书联系人邮箱
	ContactEmail *string `json:"contactEmail,omitempty"`
    // 创建时间
	CreateTime *string `json:"createTime,omitempty"`
    // 证书联系人
	ContactName *string `json:"contactName,omitempty"`
    // 过期时间
	ExpirationTime *string `json:"expirationTime,omitempty"`
    // 证书公钥
	Certificate *string `json:"certificate,omitempty"`
    // 证书名称
	CrtName *string `json:"crtName,omitempty"`
    // 绑定的域名名称
	DomainNames *string `json:"domainNames,omitempty"`
    // 证书唯一id
	UniqueId *string `json:"uniqueId,omitempty"`
}

func (s DescribeUserCrtListResponseList) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserCrtListResponseList) GoString() string {
	return s.String()
}

func (s DescribeUserCrtListResponseList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserCrtListResponseList) SetContactMobile(v string) *DescribeUserCrtListResponseList {
	s.ContactMobile = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetDomainIds(v string) *DescribeUserCrtListResponseList {
	s.DomainIds = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetPrivateKey(v string) *DescribeUserCrtListResponseList {
	s.PrivateKey = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetContactEmail(v string) *DescribeUserCrtListResponseList {
	s.ContactEmail = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetCreateTime(v string) *DescribeUserCrtListResponseList {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetContactName(v string) *DescribeUserCrtListResponseList {
	s.ContactName = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetExpirationTime(v string) *DescribeUserCrtListResponseList {
	s.ExpirationTime = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetCertificate(v string) *DescribeUserCrtListResponseList {
	s.Certificate = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetCrtName(v string) *DescribeUserCrtListResponseList {
	s.CrtName = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetDomainNames(v string) *DescribeUserCrtListResponseList {
	s.DomainNames = &v
	return s
}

func (s *DescribeUserCrtListResponseList) SetUniqueId(v string) *DescribeUserCrtListResponseList {
	s.UniqueId = &v
	return s
}


type DescribeUserCrtListResponseListBuilder struct {
	s *DescribeUserCrtListResponseList
}

func NewDescribeUserCrtListResponseListBuilder() *DescribeUserCrtListResponseListBuilder {
	s := &DescribeUserCrtListResponseList{}
	b := &DescribeUserCrtListResponseListBuilder{s: s}
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) ContactMobile(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.ContactMobile = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) DomainIds(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.DomainIds = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) PrivateKey(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.PrivateKey = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) ContactEmail(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.ContactEmail = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) CreateTime(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.CreateTime = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) ContactName(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.ContactName = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) ExpirationTime(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.ExpirationTime = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) Certificate(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.Certificate = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) CrtName(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.CrtName = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) DomainNames(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.DomainNames = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) UniqueId(v string) *DescribeUserCrtListResponseListBuilder {
	b.s.UniqueId = &v
	return b
}

func (b *DescribeUserCrtListResponseListBuilder) Build() *DescribeUserCrtListResponseList {
	return b.s
}


