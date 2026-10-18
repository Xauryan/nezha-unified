// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListPoolRespRequest struct {


	ListPoolRespPath *ListPoolRespPath `json:"listPoolRespPath,omitempty"`

	ListPoolRespQuery *ListPoolRespQuery `json:"listPoolRespQuery,omitempty"`
}

func (s ListPoolRespRequest) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespRequest) GoString() string {
	return s.String()
}

func (s ListPoolRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespRequest) SetListPoolRespPath(v *ListPoolRespPath) *ListPoolRespRequest {
	s.ListPoolRespPath = v
	return s
}

func (s *ListPoolRespRequest) SetListPoolRespQuery(v *ListPoolRespQuery) *ListPoolRespRequest {
	s.ListPoolRespQuery = v
	return s
}


type ListPoolRespRequestBuilder struct {
	s *ListPoolRespRequest
}

func NewListPoolRespRequestBuilder() *ListPoolRespRequestBuilder {
	s := &ListPoolRespRequest{}
	b := &ListPoolRespRequestBuilder{s: s}
	return b
}

func (b *ListPoolRespRequestBuilder) ListPoolRespPath(v *ListPoolRespPath) *ListPoolRespRequestBuilder {
	b.s.ListPoolRespPath = v
	return b
}

func (b *ListPoolRespRequestBuilder) ListPoolRespQuery(v *ListPoolRespQuery) *ListPoolRespRequestBuilder {
	b.s.ListPoolRespQuery = v
	return b
}

func (b *ListPoolRespRequestBuilder) Build() *ListPoolRespRequest {
	return b.s
}


