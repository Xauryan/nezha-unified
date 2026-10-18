// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListTLSCustomizeProtocolRespRequest struct {


	ListTLSCustomizeProtocolRespQuery *ListTLSCustomizeProtocolRespQuery `json:"listTLSCustomizeProtocolRespQuery,omitempty"`
}

func (s ListTLSCustomizeProtocolRespRequest) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespRequest) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespRequest) SetListTLSCustomizeProtocolRespQuery(v *ListTLSCustomizeProtocolRespQuery) *ListTLSCustomizeProtocolRespRequest {
	s.ListTLSCustomizeProtocolRespQuery = v
	return s
}


type ListTLSCustomizeProtocolRespRequestBuilder struct {
	s *ListTLSCustomizeProtocolRespRequest
}

func NewListTLSCustomizeProtocolRespRequestBuilder() *ListTLSCustomizeProtocolRespRequestBuilder {
	s := &ListTLSCustomizeProtocolRespRequest{}
	b := &ListTLSCustomizeProtocolRespRequestBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespRequestBuilder) ListTLSCustomizeProtocolRespQuery(v *ListTLSCustomizeProtocolRespQuery) *ListTLSCustomizeProtocolRespRequestBuilder {
	b.s.ListTLSCustomizeProtocolRespQuery = v
	return b
}

func (b *ListTLSCustomizeProtocolRespRequestBuilder) Build() *ListTLSCustomizeProtocolRespRequest {
	return b.s
}


