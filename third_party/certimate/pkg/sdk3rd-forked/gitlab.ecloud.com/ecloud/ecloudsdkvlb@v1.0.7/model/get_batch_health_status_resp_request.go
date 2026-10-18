// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBatchHealthStatusRespRequest struct {


	GetBatchHealthStatusRespPath *GetBatchHealthStatusRespPath `json:"getBatchHealthStatusRespPath,omitempty"`
}

func (s GetBatchHealthStatusRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetBatchHealthStatusRespRequest) GoString() string {
	return s.String()
}

func (s GetBatchHealthStatusRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBatchHealthStatusRespRequest) SetGetBatchHealthStatusRespPath(v *GetBatchHealthStatusRespPath) *GetBatchHealthStatusRespRequest {
	s.GetBatchHealthStatusRespPath = v
	return s
}


type GetBatchHealthStatusRespRequestBuilder struct {
	s *GetBatchHealthStatusRespRequest
}

func NewGetBatchHealthStatusRespRequestBuilder() *GetBatchHealthStatusRespRequestBuilder {
	s := &GetBatchHealthStatusRespRequest{}
	b := &GetBatchHealthStatusRespRequestBuilder{s: s}
	return b
}

func (b *GetBatchHealthStatusRespRequestBuilder) GetBatchHealthStatusRespPath(v *GetBatchHealthStatusRespPath) *GetBatchHealthStatusRespRequestBuilder {
	b.s.GetBatchHealthStatusRespPath = v
	return b
}

func (b *GetBatchHealthStatusRespRequestBuilder) Build() *GetBatchHealthStatusRespRequest {
	return b.s
}


