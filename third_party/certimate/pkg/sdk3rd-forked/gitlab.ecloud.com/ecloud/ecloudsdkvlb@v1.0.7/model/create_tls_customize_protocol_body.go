// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateTLSCustomizeProtocolBody struct {
    position.Body
    // TLS安全策略协议【TLS1.2】【TLS1.2,TLS1.3】【TLS1.1,TLS1.2】【TLS1.0,TLS1.1,TLS1.2】【TLS1.1,TLS1.2,TLS1.3】【TLS1.0,TLS1.1,TLS1.2,TLS1.3】
	TlsProtocols *string `json:"tlsProtocols,omitempty"`
    // 加密算法套件 【TLS1.0,TLS1.1,TLS1.2支持：CAMELLIA128-SHA、CAMELLIA256-SHA、ECDHE-ECDSA-AES128-SHA、ECDHE-ECDSA-AES256-SHA、AES128-SHA、AES256-SHA、DES-CBC3-SHA】 【TLS1.2支持：ECDHE-ECDSA-AES128-GCM-SHA256、ECDHE-ECDSA-AES256-GCM-SHA384、ECDHE-ECDSA-AES128-SHA256、ECDHE-ECDSA-AES256-SHA384、ECDHE-RSA-AES128-GCM-SHA256、ECDHE-RSA-AES256-GCM-SHA384、ECDHE-RSA-AES128-SHA256、ECDHE-RSA-AES256-SHA384、AES128-GCM-SHA256、AES256-SHA256】 【TLS1.3支持：TLS13-AES128-GCM-SHA256、TLS13-AES256-GCM-SHA384、TLS13-CHACHA20-POLY1305-SHA256】
	CipherSuites *string `json:"cipherSuites,omitempty"`
    // TLS安全策略名称
	Name *string `json:"name,omitempty"`
}

func (s CreateTLSCustomizeProtocolBody) String() string {
	return utils.Beautify(s)
}

func (s CreateTLSCustomizeProtocolBody) GoString() string {
	return s.String()
}

func (s CreateTLSCustomizeProtocolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTLSCustomizeProtocolBody) SetTlsProtocols(v string) *CreateTLSCustomizeProtocolBody {
	s.TlsProtocols = &v
	return s
}

func (s *CreateTLSCustomizeProtocolBody) SetCipherSuites(v string) *CreateTLSCustomizeProtocolBody {
	s.CipherSuites = &v
	return s
}

func (s *CreateTLSCustomizeProtocolBody) SetName(v string) *CreateTLSCustomizeProtocolBody {
	s.Name = &v
	return s
}


type CreateTLSCustomizeProtocolBodyBuilder struct {
	s *CreateTLSCustomizeProtocolBody
}

func NewCreateTLSCustomizeProtocolBodyBuilder() *CreateTLSCustomizeProtocolBodyBuilder {
	s := &CreateTLSCustomizeProtocolBody{}
	b := &CreateTLSCustomizeProtocolBodyBuilder{s: s}
	return b
}

func (b *CreateTLSCustomizeProtocolBodyBuilder) TlsProtocols(v string) *CreateTLSCustomizeProtocolBodyBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *CreateTLSCustomizeProtocolBodyBuilder) CipherSuites(v string) *CreateTLSCustomizeProtocolBodyBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *CreateTLSCustomizeProtocolBodyBuilder) Name(v string) *CreateTLSCustomizeProtocolBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateTLSCustomizeProtocolBodyBuilder) Build() *CreateTLSCustomizeProtocolBody {
	return b.s
}


