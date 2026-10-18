// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDomainServerCertificatePath struct {
    position.Path
    // 证书唯一编号
	UniqueId *string `json:"uniqueId,omitempty"`
}

func (s DeleteDomainServerCertificatePath) String() string {
	return utils.Beautify(s)
}

func (s DeleteDomainServerCertificatePath) GoString() string {
	return s.String()
}

func (s DeleteDomainServerCertificatePath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDomainServerCertificatePath) SetUniqueId(v string) *DeleteDomainServerCertificatePath {
	s.UniqueId = &v
	return s
}


type DeleteDomainServerCertificatePathBuilder struct {
	s *DeleteDomainServerCertificatePath
}

func NewDeleteDomainServerCertificatePathBuilder() *DeleteDomainServerCertificatePathBuilder {
	s := &DeleteDomainServerCertificatePath{}
	b := &DeleteDomainServerCertificatePathBuilder{s: s}
	return b
}

func (b *DeleteDomainServerCertificatePathBuilder) UniqueId(v string) *DeleteDomainServerCertificatePathBuilder {
	b.s.UniqueId = &v
	return b
}

func (b *DeleteDomainServerCertificatePathBuilder) Build() *DeleteDomainServerCertificatePath {
	return b.s
}


