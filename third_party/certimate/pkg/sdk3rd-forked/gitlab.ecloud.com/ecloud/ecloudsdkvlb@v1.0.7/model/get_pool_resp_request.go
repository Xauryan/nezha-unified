// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPoolRespRequest struct {


	GetPoolRespPath *GetPoolRespPath `json:"getPoolRespPath,omitempty"`
}

func (s GetPoolRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespRequest) GoString() string {
	return s.String()
}

func (s GetPoolRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespRequest) SetGetPoolRespPath(v *GetPoolRespPath) *GetPoolRespRequest {
	s.GetPoolRespPath = v
	return s
}


type GetPoolRespRequestBuilder struct {
	s *GetPoolRespRequest
}

func NewGetPoolRespRequestBuilder() *GetPoolRespRequestBuilder {
	s := &GetPoolRespRequest{}
	b := &GetPoolRespRequestBuilder{s: s}
	return b
}

func (b *GetPoolRespRequestBuilder) GetPoolRespPath(v *GetPoolRespPath) *GetPoolRespRequestBuilder {
	b.s.GetPoolRespPath = v
	return b
}

func (b *GetPoolRespRequestBuilder) Build() *GetPoolRespRequest {
	return b.s
}


