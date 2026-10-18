// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListL7PolicyRespRequest struct {


	ListL7PolicyRespQuery *ListL7PolicyRespQuery `json:"listL7PolicyRespQuery,omitempty"`
}

func (s ListL7PolicyRespRequest) String() string {
	return utils.Beautify(s)
}

func (s ListL7PolicyRespRequest) GoString() string {
	return s.String()
}

func (s ListL7PolicyRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListL7PolicyRespRequest) SetListL7PolicyRespQuery(v *ListL7PolicyRespQuery) *ListL7PolicyRespRequest {
	s.ListL7PolicyRespQuery = v
	return s
}


type ListL7PolicyRespRequestBuilder struct {
	s *ListL7PolicyRespRequest
}

func NewListL7PolicyRespRequestBuilder() *ListL7PolicyRespRequestBuilder {
	s := &ListL7PolicyRespRequest{}
	b := &ListL7PolicyRespRequestBuilder{s: s}
	return b
}

func (b *ListL7PolicyRespRequestBuilder) ListL7PolicyRespQuery(v *ListL7PolicyRespQuery) *ListL7PolicyRespRequestBuilder {
	b.s.ListL7PolicyRespQuery = v
	return b
}

func (b *ListL7PolicyRespRequestBuilder) Build() *ListL7PolicyRespRequest {
	return b.s
}


