// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDomainServerCertificateBody struct {
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
    // 证书唯一id，证书新增时无需填写，修改时为必填项
	UniqueId *string `json:"uniqueId,omitempty"`
}

func (s UpdateDomainServerCertificateBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateDomainServerCertificateBody) GoString() string {
	return s.String()
}

func (s UpdateDomainServerCertificateBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDomainServerCertificateBody) SetContactMobile(v string) *UpdateDomainServerCertificateBody {
	s.ContactMobile = &v
	return s
}

func (s *UpdateDomainServerCertificateBody) SetPrivateKey(v string) *UpdateDomainServerCertificateBody {
	s.PrivateKey = &v
	return s
}

func (s *UpdateDomainServerCertificateBody) SetContactEmail(v string) *UpdateDomainServerCertificateBody {
	s.ContactEmail = &v
	return s
}

func (s *UpdateDomainServerCertificateBody) SetContactName(v string) *UpdateDomainServerCertificateBody {
	s.ContactName = &v
	return s
}

func (s *UpdateDomainServerCertificateBody) SetCertificate(v string) *UpdateDomainServerCertificateBody {
	s.Certificate = &v
	return s
}

func (s *UpdateDomainServerCertificateBody) SetUniqueId(v string) *UpdateDomainServerCertificateBody {
	s.UniqueId = &v
	return s
}


type UpdateDomainServerCertificateBodyBuilder struct {
	s *UpdateDomainServerCertificateBody
}

func NewUpdateDomainServerCertificateBodyBuilder() *UpdateDomainServerCertificateBodyBuilder {
	s := &UpdateDomainServerCertificateBody{}
	b := &UpdateDomainServerCertificateBodyBuilder{s: s}
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) ContactMobile(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.ContactMobile = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) PrivateKey(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.PrivateKey = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) ContactEmail(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.ContactEmail = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) ContactName(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.ContactName = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) Certificate(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.Certificate = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) UniqueId(v string) *UpdateDomainServerCertificateBodyBuilder {
	b.s.UniqueId = &v
	return b
}

func (b *UpdateDomainServerCertificateBodyBuilder) Build() *UpdateDomainServerCertificateBody {
	return b.s
}


