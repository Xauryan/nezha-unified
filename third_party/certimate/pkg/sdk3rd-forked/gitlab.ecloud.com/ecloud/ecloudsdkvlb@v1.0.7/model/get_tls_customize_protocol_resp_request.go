// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTLSCustomizeProtocolRespRequest struct {


	GetTLSCustomizeProtocolRespPath *GetTLSCustomizeProtocolRespPath `json:"getTLSCustomizeProtocolRespPath,omitempty"`
}

func (s GetTLSCustomizeProtocolRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetTLSCustomizeProtocolRespRequest) GoString() string {
	return s.String()
}

func (s GetTLSCustomizeProtocolRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTLSCustomizeProtocolRespRequest) SetGetTLSCustomizeProtocolRespPath(v *GetTLSCustomizeProtocolRespPath) *GetTLSCustomizeProtocolRespRequest {
	s.GetTLSCustomizeProtocolRespPath = v
	return s
}


type GetTLSCustomizeProtocolRespRequestBuilder struct {
	s *GetTLSCustomizeProtocolRespRequest
}

func NewGetTLSCustomizeProtocolRespRequestBuilder() *GetTLSCustomizeProtocolRespRequestBuilder {
	s := &GetTLSCustomizeProtocolRespRequest{}
	b := &GetTLSCustomizeProtocolRespRequestBuilder{s: s}
	return b
}

func (b *GetTLSCustomizeProtocolRespRequestBuilder) GetTLSCustomizeProtocolRespPath(v *GetTLSCustomizeProtocolRespPath) *GetTLSCustomizeProtocolRespRequestBuilder {
	b.s.GetTLSCustomizeProtocolRespPath = v
	return b
}

func (b *GetTLSCustomizeProtocolRespRequestBuilder) Build() *GetTLSCustomizeProtocolRespRequest {
	return b.s
}


