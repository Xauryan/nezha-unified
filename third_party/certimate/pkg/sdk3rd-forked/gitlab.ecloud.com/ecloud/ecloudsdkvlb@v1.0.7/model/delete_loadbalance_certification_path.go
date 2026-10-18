// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteLoadbalanceCertificationPath struct {
    position.Path
    // 证书ID
	ContainerUuid *string `json:"containerUuid,omitempty"`
}

func (s DeleteLoadbalanceCertificationPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteLoadbalanceCertificationPath) GoString() string {
	return s.String()
}

func (s DeleteLoadbalanceCertificationPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteLoadbalanceCertificationPath) SetContainerUuid(v string) *DeleteLoadbalanceCertificationPath {
	s.ContainerUuid = &v
	return s
}


type DeleteLoadbalanceCertificationPathBuilder struct {
	s *DeleteLoadbalanceCertificationPath
}

func NewDeleteLoadbalanceCertificationPathBuilder() *DeleteLoadbalanceCertificationPathBuilder {
	s := &DeleteLoadbalanceCertificationPath{}
	b := &DeleteLoadbalanceCertificationPathBuilder{s: s}
	return b
}

func (b *DeleteLoadbalanceCertificationPathBuilder) ContainerUuid(v string) *DeleteLoadbalanceCertificationPathBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *DeleteLoadbalanceCertificationPathBuilder) Build() *DeleteLoadbalanceCertificationPath {
	return b.s
}


