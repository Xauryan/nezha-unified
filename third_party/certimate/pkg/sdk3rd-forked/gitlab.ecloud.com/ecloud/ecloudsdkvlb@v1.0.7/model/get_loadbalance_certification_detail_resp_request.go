// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadbalanceCertificationDetailRespRequest struct {


	GetLoadbalanceCertificationDetailRespPath *GetLoadbalanceCertificationDetailRespPath `json:"getLoadbalanceCertificationDetailRespPath,omitempty"`
}

func (s GetLoadbalanceCertificationDetailRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetLoadbalanceCertificationDetailRespRequest) GoString() string {
	return s.String()
}

func (s GetLoadbalanceCertificationDetailRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadbalanceCertificationDetailRespRequest) SetGetLoadbalanceCertificationDetailRespPath(v *GetLoadbalanceCertificationDetailRespPath) *GetLoadbalanceCertificationDetailRespRequest {
	s.GetLoadbalanceCertificationDetailRespPath = v
	return s
}


type GetLoadbalanceCertificationDetailRespRequestBuilder struct {
	s *GetLoadbalanceCertificationDetailRespRequest
}

func NewGetLoadbalanceCertificationDetailRespRequestBuilder() *GetLoadbalanceCertificationDetailRespRequestBuilder {
	s := &GetLoadbalanceCertificationDetailRespRequest{}
	b := &GetLoadbalanceCertificationDetailRespRequestBuilder{s: s}
	return b
}

func (b *GetLoadbalanceCertificationDetailRespRequestBuilder) GetLoadbalanceCertificationDetailRespPath(v *GetLoadbalanceCertificationDetailRespPath) *GetLoadbalanceCertificationDetailRespRequestBuilder {
	b.s.GetLoadbalanceCertificationDetailRespPath = v
	return b
}

func (b *GetLoadbalanceCertificationDetailRespRequestBuilder) Build() *GetLoadbalanceCertificationDetailRespRequest {
	return b.s
}


