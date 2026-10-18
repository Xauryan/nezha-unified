// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTLSCustomizeProtocolRespPath struct {
    position.Path
    // TLS安全策略ID
	TlsCustomizeProtocolId *string `json:"tlsCustomizeProtocolId,omitempty"`
}

func (s GetTLSCustomizeProtocolRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetTLSCustomizeProtocolRespPath) GoString() string {
	return s.String()
}

func (s GetTLSCustomizeProtocolRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTLSCustomizeProtocolRespPath) SetTlsCustomizeProtocolId(v string) *GetTLSCustomizeProtocolRespPath {
	s.TlsCustomizeProtocolId = &v
	return s
}


type GetTLSCustomizeProtocolRespPathBuilder struct {
	s *GetTLSCustomizeProtocolRespPath
}

func NewGetTLSCustomizeProtocolRespPathBuilder() *GetTLSCustomizeProtocolRespPathBuilder {
	s := &GetTLSCustomizeProtocolRespPath{}
	b := &GetTLSCustomizeProtocolRespPathBuilder{s: s}
	return b
}

func (b *GetTLSCustomizeProtocolRespPathBuilder) TlsCustomizeProtocolId(v string) *GetTLSCustomizeProtocolRespPathBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespPathBuilder) Build() *GetTLSCustomizeProtocolRespPath {
	return b.s
}


