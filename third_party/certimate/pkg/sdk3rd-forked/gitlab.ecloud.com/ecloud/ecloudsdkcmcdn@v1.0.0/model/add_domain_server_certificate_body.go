// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddDomainServerCertificateBody struct {
    position.Body
    // 证书联系人手机号
	ContactMobile *string `json:"contactMobile,omitempty"`
    // 证书私钥
	PrivateKey *string `json:"privateKey,omitempty"`
    // 证书联系人邮箱
	ContactEmail *string `json:"contactEmail,omitempty"`
    // 证书联系人
	ContactName *string `json:"contactName,omitempty"`
    // 证书公钥
	Certificate *string `json:"certificate,omitempty"`
    // 证书名称
	CrtName *string `json:"crtName,omitempty"`
    // 绑定的域名id
	DomainId *int32 `json:"domainId,omitempty"`
    // 证书唯一id，证书新增时无需填写，修改时为必填项
	UniqueId *string `json:"uniqueId,omitempty"`
}

func (s AddDomainServerCertificateBody) String() string {
	return utils.Beautify(s)
}

func (s AddDomainServerCertificateBody) GoString() string {
	return s.String()
}

func (s AddDomainServerCertificateBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddDomainServerCertificateBody) SetContactMobile(v string) *AddDomainServerCertificateBody {
	s.ContactMobile = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetPrivateKey(v string) *AddDomainServerCertificateBody {
	s.PrivateKey = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetContactEmail(v string) *AddDomainServerCertificateBody {
	s.ContactEmail = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetContactName(v string) *AddDomainServerCertificateBody {
	s.ContactName = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetCertificate(v string) *AddDomainServerCertificateBody {
	s.Certificate = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetCrtName(v string) *AddDomainServerCertificateBody {
	s.CrtName = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetDomainId(v int32) *AddDomainServerCertificateBody {
	s.DomainId = &v
	return s
}

func (s *AddDomainServerCertificateBody) SetUniqueId(v string) *AddDomainServerCertificateBody {
	s.UniqueId = &v
	return s
}


type AddDomainServerCertificateBodyBuilder struct {
	s *AddDomainServerCertificateBody
}

func NewAddDomainServerCertificateBodyBuilder() *AddDomainServerCertificateBodyBuilder {
	s := &AddDomainServerCertificateBody{}
	b := &AddDomainServerCertificateBodyBuilder{s: s}
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) ContactMobile(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.ContactMobile = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) PrivateKey(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.PrivateKey = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) ContactEmail(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.ContactEmail = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) ContactName(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.ContactName = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) Certificate(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.Certificate = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) CrtName(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.CrtName = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) DomainId(v int32) *AddDomainServerCertificateBodyBuilder {
	b.s.DomainId = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) UniqueId(v string) *AddDomainServerCertificateBodyBuilder {
	b.s.UniqueId = &v
	return b
}

func (b *AddDomainServerCertificateBodyBuilder) Build() *AddDomainServerCertificateBody {
	return b.s
}


