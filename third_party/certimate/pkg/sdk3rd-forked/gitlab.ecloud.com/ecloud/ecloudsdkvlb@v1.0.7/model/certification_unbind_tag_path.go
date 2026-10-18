// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CertificationUnbindTagPath struct {
    position.Path
    // 证书ID
	ContainerUuid *string `json:"containerUuid,omitempty"`
}

func (s CertificationUnbindTagPath) String() string {
	return utils.Beautify(s)
}

func (s CertificationUnbindTagPath) GoString() string {
	return s.String()
}

func (s CertificationUnbindTagPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationUnbindTagPath) SetContainerUuid(v string) *CertificationUnbindTagPath {
	s.ContainerUuid = &v
	return s
}


type CertificationUnbindTagPathBuilder struct {
	s *CertificationUnbindTagPath
}

func NewCertificationUnbindTagPathBuilder() *CertificationUnbindTagPathBuilder {
	s := &CertificationUnbindTagPath{}
	b := &CertificationUnbindTagPathBuilder{s: s}
	return b
}

func (b *CertificationUnbindTagPathBuilder) ContainerUuid(v string) *CertificationUnbindTagPathBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *CertificationUnbindTagPathBuilder) Build() *CertificationUnbindTagPath {
	return b.s
}


