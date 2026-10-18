// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetL7PolicyRespRequest struct {


	GetL7PolicyRespPath *GetL7PolicyRespPath `json:"getL7PolicyRespPath,omitempty"`
}

func (s GetL7PolicyRespRequest) String() string {
	return utils.Beautify(s)
}

func (s GetL7PolicyRespRequest) GoString() string {
	return s.String()
}

func (s GetL7PolicyRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetL7PolicyRespRequest) SetGetL7PolicyRespPath(v *GetL7PolicyRespPath) *GetL7PolicyRespRequest {
	s.GetL7PolicyRespPath = v
	return s
}


type GetL7PolicyRespRequestBuilder struct {
	s *GetL7PolicyRespRequest
}

func NewGetL7PolicyRespRequestBuilder() *GetL7PolicyRespRequestBuilder {
	s := &GetL7PolicyRespRequest{}
	b := &GetL7PolicyRespRequestBuilder{s: s}
	return b
}

func (b *GetL7PolicyRespRequestBuilder) GetL7PolicyRespPath(v *GetL7PolicyRespPath) *GetL7PolicyRespRequestBuilder {
	b.s.GetL7PolicyRespPath = v
	return b
}

func (b *GetL7PolicyRespRequestBuilder) Build() *GetL7PolicyRespRequest {
	return b.s
}


