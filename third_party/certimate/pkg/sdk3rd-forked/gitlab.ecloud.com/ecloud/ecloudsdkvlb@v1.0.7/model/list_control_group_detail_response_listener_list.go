// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListControlGroupDetailResponseListenerListProtocolEnum string

// List of Protocol
const (
    ListControlGroupDetailResponseListenerListProtocolEnumHttp ListControlGroupDetailResponseListenerListProtocolEnum = "HTTP"
    ListControlGroupDetailResponseListenerListProtocolEnumHttps ListControlGroupDetailResponseListenerListProtocolEnum = "HTTPS"
    ListControlGroupDetailResponseListenerListProtocolEnumTcp ListControlGroupDetailResponseListenerListProtocolEnum = "TCP"
    ListControlGroupDetailResponseListenerListProtocolEnumUdp ListControlGroupDetailResponseListenerListProtocolEnum = "UDP"
    ListControlGroupDetailResponseListenerListProtocolEnumTerminatedHttps ListControlGroupDetailResponseListenerListProtocolEnum = "TERMINATED_HTTPS"
    ListControlGroupDetailResponseListenerListProtocolEnumSip ListControlGroupDetailResponseListenerListProtocolEnum = "SIP"
)
type ListControlGroupDetailResponseListenerListRedirectProtocolEnum string

// List of RedirectProtocol
const (
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumHttp ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "HTTP"
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumHttps ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "HTTPS"
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumTcp ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "TCP"
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumUdp ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "UDP"
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumTerminatedHttps ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "TERMINATED_HTTPS"
    ListControlGroupDetailResponseListenerListRedirectProtocolEnumSip ListControlGroupDetailResponseListenerListRedirectProtocolEnum = "SIP"
)

type ListControlGroupDetailResponseListenerList struct {

    // 租户ID
	TenantId *string `json:"tenant_id,omitempty"`
    // 服务器证书信息
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
    // TLS安全策略协议
	TlsProtocols *string `json:"tls_protocols,omitempty"`
    // 负载均衡ID
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxy_protocol,omitempty"`
    // 监听器描述
	Description *string `json:"description,omitempty"`
    // 是否透传源ID
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器连接限制
	ConnectionLimit *int32 `json:"connection_limit,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutual_authentication_up,omitempty"`
    // 监听器前后端使用的协议
	Protocol *ListControlGroupDetailResponseListenerListProtocolEnum `json:"protocol,omitempty"`
    // 监听器是否可用
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
    // 是否开启http2
	Http2 *bool `json:"http2,omitempty"`
    // 重定向监听器前端协议的端口号
	RedirectPort *int32 `json:"redirect_port,omitempty"`
    // 监听器ID
	Id *string `json:"id,omitempty"`
    // 默认后端服务器组ID
	DefaultPoolId *string `json:"default_pool_id,omitempty"`
    // 服务器证书ID
	DefaultTlsContainerId *string `json:"default_tls_container_id,omitempty"`
    // 监听器前端协议的端口号
	ProtocolPort *int32 `json:"protocol_port,omitempty"`
    // 连接超时时长
	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`
    // SNI证书信息
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`
    // 加密套件
	CipherSuites *string `json:"cipher_suites,omitempty"`
    // 重定向监听器前后端使用的协议
	RedirectProtocol *ListControlGroupDetailResponseListenerListRedirectProtocolEnum `json:"redirect_protocol,omitempty"`
    // 是否开启重定向
	RedirectUp *bool `json:"redirect_up,omitempty"`
    // 证书容器中的CA证书ID
	CaContainerId *string `json:"ca_container_id,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 负载均衡信息
	Loadbalancers *[]ListControlGroupDetailResponseLoadbalancers `json:"loadbalancers,omitempty"`
}

func (s ListControlGroupDetailResponseListenerList) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailResponseListenerList) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailResponseListenerList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailResponseListenerList) SetTenantId(v string) *ListControlGroupDetailResponseListenerList {
	s.TenantId = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetDefaultTlsContainerRef(v string) *ListControlGroupDetailResponseListenerList {
	s.DefaultTlsContainerRef = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetTlsProtocols(v string) *ListControlGroupDetailResponseListenerList {
	s.TlsProtocols = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetLoadbalancerId(v string) *ListControlGroupDetailResponseListenerList {
	s.LoadbalancerId = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetProxyProtocol(v bool) *ListControlGroupDetailResponseListenerList {
	s.ProxyProtocol = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetDescription(v string) *ListControlGroupDetailResponseListenerList {
	s.Description = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetTransparent(v bool) *ListControlGroupDetailResponseListenerList {
	s.Transparent = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetConnectionLimit(v int32) *ListControlGroupDetailResponseListenerList {
	s.ConnectionLimit = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetMutualAuthenticationUp(v bool) *ListControlGroupDetailResponseListenerList {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetProtocol(v ListControlGroupDetailResponseListenerListProtocolEnum) *ListControlGroupDetailResponseListenerList {
	s.Protocol = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetAdminStateUp(v bool) *ListControlGroupDetailResponseListenerList {
	s.AdminStateUp = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetHttp2(v bool) *ListControlGroupDetailResponseListenerList {
	s.Http2 = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetRedirectPort(v int32) *ListControlGroupDetailResponseListenerList {
	s.RedirectPort = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetId(v string) *ListControlGroupDetailResponseListenerList {
	s.Id = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetDefaultPoolId(v string) *ListControlGroupDetailResponseListenerList {
	s.DefaultPoolId = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetDefaultTlsContainerId(v string) *ListControlGroupDetailResponseListenerList {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetProtocolPort(v int32) *ListControlGroupDetailResponseListenerList {
	s.ProtocolPort = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetKeepaliveTimeout(v int32) *ListControlGroupDetailResponseListenerList {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetSniContainerRefs(v []string) *ListControlGroupDetailResponseListenerList {
	s.SniContainerRefs = v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetCipherSuites(v string) *ListControlGroupDetailResponseListenerList {
	s.CipherSuites = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetRedirectProtocol(v ListControlGroupDetailResponseListenerListRedirectProtocolEnum) *ListControlGroupDetailResponseListenerList {
	s.RedirectProtocol = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetRedirectUp(v bool) *ListControlGroupDetailResponseListenerList {
	s.RedirectUp = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetCaContainerId(v string) *ListControlGroupDetailResponseListenerList {
	s.CaContainerId = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetName(v string) *ListControlGroupDetailResponseListenerList {
	s.Name = &v
	return s
}

func (s *ListControlGroupDetailResponseListenerList) SetLoadbalancers(v []ListControlGroupDetailResponseLoadbalancers) *ListControlGroupDetailResponseListenerList {
	s.Loadbalancers = &v
	return s
}


type ListControlGroupDetailResponseListenerListBuilder struct {
	s *ListControlGroupDetailResponseListenerList
}

func NewListControlGroupDetailResponseListenerListBuilder() *ListControlGroupDetailResponseListenerListBuilder {
	s := &ListControlGroupDetailResponseListenerList{}
	b := &ListControlGroupDetailResponseListenerListBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) TenantId(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.TenantId = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) DefaultTlsContainerRef(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.DefaultTlsContainerRef = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) TlsProtocols(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) LoadbalancerId(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.LoadbalancerId = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) ProxyProtocol(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Description(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Transparent(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) ConnectionLimit(v int32) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) MutualAuthenticationUp(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Protocol(v ListControlGroupDetailResponseListenerListProtocolEnum) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) AdminStateUp(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Http2(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) RedirectPort(v int32) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.RedirectPort = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Id(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) DefaultPoolId(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.DefaultPoolId = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) DefaultTlsContainerId(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) ProtocolPort(v int32) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) KeepaliveTimeout(v int32) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.KeepaliveTimeout = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) SniContainerRefs(v []string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.SniContainerRefs = v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) CipherSuites(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) RedirectProtocol(v ListControlGroupDetailResponseListenerListRedirectProtocolEnum) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.RedirectProtocol = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) RedirectUp(v bool) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.RedirectUp = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) CaContainerId(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Name(v string) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Name = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Loadbalancers(v []ListControlGroupDetailResponseLoadbalancers) *ListControlGroupDetailResponseListenerListBuilder {
	b.s.Loadbalancers = &v
	return b
}

func (b *ListControlGroupDetailResponseListenerListBuilder) Build() *ListControlGroupDetailResponseListenerList {
	return b.s
}


