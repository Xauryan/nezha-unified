// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteLoadbalanceCertificationRequest struct {


	DeleteLoadbalanceCertificationPath *DeleteLoadbalanceCertificationPath `json:"deleteLoadbalanceCertificationPath,omitempty"`
}

func (s DeleteLoadbalanceCertificationRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteLoadbalanceCertificationRequest) GoString() string {
	return s.String()
}

func (s DeleteLoadbalanceCertificationRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteLoadbalanceCertificationRequest) SetDeleteLoadbalanceCertificationPath(v *DeleteLoadbalanceCertificationPath) *DeleteLoadbalanceCertificationRequest {
	s.DeleteLoadbalanceCertificationPath = v
	return s
}


type DeleteLoadbalanceCertificationRequestBuilder struct {
	s *DeleteLoadbalanceCertificationRequest
}

func NewDeleteLoadbalanceCertificationRequestBuilder() *DeleteLoadbalanceCertificationRequestBuilder {
	s := &DeleteLoadbalanceCertificationRequest{}
	b := &DeleteLoadbalanceCertificationRequestBuilder{s: s}
	return b
}

func (b *DeleteLoadbalanceCertificationRequestBuilder) DeleteLoadbalanceCertificationPath(v *DeleteLoadbalanceCertificationPath) *DeleteLoadbalanceCertificationRequestBuilder {
	b.s.DeleteLoadbalanceCertificationPath = v
	return b
}

func (b *DeleteLoadbalanceCertificationRequestBuilder) Build() *DeleteLoadbalanceCertificationRequest {
	return b.s
}


