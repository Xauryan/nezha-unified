// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateTLSCustomizeProtocolBody struct {
    position.Body
    // TLS安全策略协议【TLS1.2】【TLS1.2,TLS1.3】【TLS1.1,TLS1.2】【TLS1.0,TLS1.1,TLS1.2】【TLS1.1,TLS1.2,TLS1.3】【TLS1.0,TLS1.1,TLS1.2,TLS1.3】
	TlsProtocols *string `json:"tlsProtocols,omitempty"`
    // 加密算法套件 【TLS1.0,TLS1.1,TLS1.2支持：CAMELLIA128-SHA、CAMELLIA256-SHA、ECDHE-ECDSA-AES128-SHA、ECDHE-ECDSA-AES256-SHA、AES128-SHA、AES256-SHA、DES-CBC3-SHA】 【TLS1.2支持：ECDHE-ECDSA-AES128-GCM-SHA256、ECDHE-ECDSA-AES256-GCM-SHA384、ECDHE-ECDSA-AES128-SHA256、ECDHE-ECDSA-AES256-SHA384、ECDHE-RSA-AES128-GCM-SHA256、ECDHE-RSA-AES256-GCM-SHA384、ECDHE-RSA-AES128-SHA256、ECDHE-RSA-AES256-SHA384、AES128-GCM-SHA256、AES256-SHA256】 【TLS1.3支持：TLS13-AES128-GCM-SHA256、TLS13-AES256-GCM-SHA384、TLS13-CHACHA20-POLY1305-SHA256】
	CipherSuites *string `json:"cipherSuites,omitempty"`
    // TLS安全策略名称
	Name *string `json:"name,omitempty"`
    // TLS安全策略ID
	Id *int64 `json:"id,omitempty"`
}

func (s UpdateTLSCustomizeProtocolBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSCustomizeProtocolBody) GoString() string {
	return s.String()
}

func (s UpdateTLSCustomizeProtocolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSCustomizeProtocolBody) SetTlsProtocols(v string) *UpdateTLSCustomizeProtocolBody {
	s.TlsProtocols = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolBody) SetCipherSuites(v string) *UpdateTLSCustomizeProtocolBody {
	s.CipherSuites = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolBody) SetName(v string) *UpdateTLSCustomizeProtocolBody {
	s.Name = &v
	return s
}

func (s *UpdateTLSCustomizeProtocolBody) SetId(v int64) *UpdateTLSCustomizeProtocolBody {
	s.Id = &v
	return s
}


type UpdateTLSCustomizeProtocolBodyBuilder struct {
	s *UpdateTLSCustomizeProtocolBody
}

func NewUpdateTLSCustomizeProtocolBodyBuilder() *UpdateTLSCustomizeProtocolBodyBuilder {
	s := &UpdateTLSCustomizeProtocolBody{}
	b := &UpdateTLSCustomizeProtocolBodyBuilder{s: s}
	return b
}

func (b *UpdateTLSCustomizeProtocolBodyBuilder) TlsProtocols(v string) *UpdateTLSCustomizeProtocolBodyBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolBodyBuilder) CipherSuites(v string) *UpdateTLSCustomizeProtocolBodyBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolBodyBuilder) Name(v string) *UpdateTLSCustomizeProtocolBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolBodyBuilder) Id(v int64) *UpdateTLSCustomizeProtocolBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateTLSCustomizeProtocolBodyBuilder) Build() *UpdateTLSCustomizeProtocolBody {
	return b.s
}


