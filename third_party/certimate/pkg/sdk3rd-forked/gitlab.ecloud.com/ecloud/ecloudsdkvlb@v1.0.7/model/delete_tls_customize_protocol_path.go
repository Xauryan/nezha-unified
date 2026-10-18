// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteTLSCustomizeProtocolPath struct {
    position.Path
    // TLS安全策略ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
}

func (s DeleteTLSCustomizeProtocolPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteTLSCustomizeProtocolPath) GoString() string {
	return s.String()
}

func (s DeleteTLSCustomizeProtocolPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteTLSCustomizeProtocolPath) SetTlsCustomizeProtocolId(v int64) *DeleteTLSCustomizeProtocolPath {
	s.TlsCustomizeProtocolId = &v
	return s
}


type DeleteTLSCustomizeProtocolPathBuilder struct {
	s *DeleteTLSCustomizeProtocolPath
}

func NewDeleteTLSCustomizeProtocolPathBuilder() *DeleteTLSCustomizeProtocolPathBuilder {
	s := &DeleteTLSCustomizeProtocolPath{}
	b := &DeleteTLSCustomizeProtocolPathBuilder{s: s}
	return b
}

func (b *DeleteTLSCustomizeProtocolPathBuilder) TlsCustomizeProtocolId(v int64) *DeleteTLSCustomizeProtocolPathBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *DeleteTLSCustomizeProtocolPathBuilder) Build() *DeleteTLSCustomizeProtocolPath {
	return b.s
}


