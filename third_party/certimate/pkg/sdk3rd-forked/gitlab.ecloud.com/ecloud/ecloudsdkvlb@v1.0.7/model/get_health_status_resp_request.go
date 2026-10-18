// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetHealthStatusRespRequest struct {


	GetHealthStatusRespPath *GetHealthStatusRespPath `json:"getHealthStatusRespPath,omitempty"`
}

func (s GetHealthStatusRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetHealthStatusRespRequest) GoString() string {
	return s.String()
}

func (s GetHealthStatusRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetHealthStatusRespRequest) SetGetHealthStatusRespPath(v *GetHealthStatusRespPath) *GetHealthStatusRespRequest {
	s.GetHealthStatusRespPath = v
	return s
}


type GetHealthStatusRespRequestBuilder struct {
	s *GetHealthStatusRespRequest
}

func NewGetHealthStatusRespRequestBuilder() *GetHealthStatusRespRequestBuilder {
	s := &GetHealthStatusRespRequest{}
	b := &GetHealthStatusRespRequestBuilder{s: s}
	return b
}

func (b *GetHealthStatusRespRequestBuilder) GetHealthStatusRespPath(v *GetHealthStatusRespPath) *GetHealthStatusRespRequestBuilder {
	b.s.GetHealthStatusRespPath = v
	return b
}

func (b *GetHealthStatusRespRequestBuilder) Build() *GetHealthStatusRespRequest {
	return b.s
}


