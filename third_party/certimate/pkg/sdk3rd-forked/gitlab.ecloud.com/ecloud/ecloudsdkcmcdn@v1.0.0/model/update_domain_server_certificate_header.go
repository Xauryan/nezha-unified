// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDomainServerCertificateHeader struct {
    position.Header
    // 移动云官网用户user_id
	UserId *string `json:"user_Id,omitempty"`
}

func (s UpdateDomainServerCertificateHeader) String() string {
	return utils.Beautify(s)
}

func (s UpdateDomainServerCertificateHeader) GoString() string {
	return s.String()
}

func (s UpdateDomainServerCertificateHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDomainServerCertificateHeader) SetUserId(v string) *UpdateDomainServerCertificateHeader {
	s.UserId = &v
	return s
}


type UpdateDomainServerCertificateHeaderBuilder struct {
	s *UpdateDomainServerCertificateHeader
}

func NewUpdateDomainServerCertificateHeaderBuilder() *UpdateDomainServerCertificateHeaderBuilder {
	s := &UpdateDomainServerCertificateHeader{}
	b := &UpdateDomainServerCertificateHeaderBuilder{s: s}
	return b
}

func (b *UpdateDomainServerCertificateHeaderBuilder) UserId(v string) *UpdateDomainServerCertificateHeaderBuilder {
	b.s.UserId = &v
	return b
}

func (b *UpdateDomainServerCertificateHeaderBuilder) Build() *UpdateDomainServerCertificateHeader {
	return b.s
}


