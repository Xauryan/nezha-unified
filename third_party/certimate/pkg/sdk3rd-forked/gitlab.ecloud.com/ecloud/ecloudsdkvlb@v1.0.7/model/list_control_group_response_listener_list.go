// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListControlGroupResponseListenerListProtocolEnum string

// List of Protocol
const (
    ListControlGroupResponseListenerListProtocolEnumHttp ListControlGroupResponseListenerListProtocolEnum = "HTTP"
    ListControlGroupResponseListenerListProtocolEnumHttps ListControlGroupResponseListenerListProtocolEnum = "HTTPS"
    ListControlGroupResponseListenerListProtocolEnumTcp ListControlGroupResponseListenerListProtocolEnum = "TCP"
    ListControlGroupResponseListenerListProtocolEnumUdp ListControlGroupResponseListenerListProtocolEnum = "UDP"
    ListControlGroupResponseListenerListProtocolEnumTerminatedHttps ListControlGroupResponseListenerListProtocolEnum = "TERMINATED_HTTPS"
    ListControlGroupResponseListenerListProtocolEnumSip ListControlGroupResponseListenerListProtocolEnum = "SIP"
)
type ListControlGroupResponseListenerListRedirectProtocolEnum string

// List of RedirectProtocol
const (
    ListControlGroupResponseListenerListRedirectProtocolEnumHttp ListControlGroupResponseListenerListRedirectProtocolEnum = "HTTP"
    ListControlGroupResponseListenerListRedirectProtocolEnumHttps ListControlGroupResponseListenerListRedirectProtocolEnum = "HTTPS"
    ListControlGroupResponseListenerListRedirectProtocolEnumTcp ListControlGroupResponseListenerListRedirectProtocolEnum = "TCP"
    ListControlGroupResponseListenerListRedirectProtocolEnumUdp ListControlGroupResponseListenerListRedirectProtocolEnum = "UDP"
    ListControlGroupResponseListenerListRedirectProtocolEnumTerminatedHttps ListControlGroupResponseListenerListRedirectProtocolEnum = "TERMINATED_HTTPS"
    ListControlGroupResponseListenerListRedirectProtocolEnumSip ListControlGroupResponseListenerListRedirectProtocolEnum = "SIP"
)

type ListControlGroupResponseListenerList struct {

    // 租户ID
	TenantId *string `json:"tenant_id,omitempty"`
    // 服务器证书信息
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
    // TLS安全策略协议
	TlsProtocols *string `json:"tls_protocols,omitempty"`
    // 负载均衡ID
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
    // 监听器是否开启pp类型的透传IP
	ProxyProtocol *bool `json:"proxy_protocol,omitempty"`
    // 监听器描述
	Description *string `json:"description,omitempty"`
    // 是否透传源ID
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器连接最大限制
	ConnectionLimit *int32 `json:"connection_limit,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutual_authentication_up,omitempty"`
    // 监听器前后端使用的协议
	Protocol *ListControlGroupResponseListenerListProtocolEnum `json:"protocol,omitempty"`
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
	RedirectProtocol *ListControlGroupResponseListenerListRedirectProtocolEnum `json:"redirect_protocol,omitempty"`
    // 是否开启重定向
	RedirectUp *bool `json:"redirect_up,omitempty"`
    // 证书容器中的CA证书ID
	CaContainerId *string `json:"ca_container_id,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 负载均衡ID信息
	Loadbalancers *[]ListControlGroupResponseLoadbalancers `json:"loadbalancers,omitempty"`
}

func (s ListControlGroupResponseListenerList) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupResponseListenerList) GoString() string {
	return s.String()
}

func (s ListControlGroupResponseListenerList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupResponseListenerList) SetTenantId(v string) *ListControlGroupResponseListenerList {
	s.TenantId = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetDefaultTlsContainerRef(v string) *ListControlGroupResponseListenerList {
	s.DefaultTlsContainerRef = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetTlsProtocols(v string) *ListControlGroupResponseListenerList {
	s.TlsProtocols = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetLoadbalancerId(v string) *ListControlGroupResponseListenerList {
	s.LoadbalancerId = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetProxyProtocol(v bool) *ListControlGroupResponseListenerList {
	s.ProxyProtocol = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetDescription(v string) *ListControlGroupResponseListenerList {
	s.Description = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetTransparent(v bool) *ListControlGroupResponseListenerList {
	s.Transparent = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetConnectionLimit(v int32) *ListControlGroupResponseListenerList {
	s.ConnectionLimit = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetMutualAuthenticationUp(v bool) *ListControlGroupResponseListenerList {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetProtocol(v ListControlGroupResponseListenerListProtocolEnum) *ListControlGroupResponseListenerList {
	s.Protocol = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetAdminStateUp(v bool) *ListControlGroupResponseListenerList {
	s.AdminStateUp = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetHttp2(v bool) *ListControlGroupResponseListenerList {
	s.Http2 = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetRedirectPort(v int32) *ListControlGroupResponseListenerList {
	s.RedirectPort = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetId(v string) *ListControlGroupResponseListenerList {
	s.Id = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetDefaultPoolId(v string) *ListControlGroupResponseListenerList {
	s.DefaultPoolId = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetDefaultTlsContainerId(v string) *ListControlGroupResponseListenerList {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetProtocolPort(v int32) *ListControlGroupResponseListenerList {
	s.ProtocolPort = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetKeepaliveTimeout(v int32) *ListControlGroupResponseListenerList {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetSniContainerRefs(v []string) *ListControlGroupResponseListenerList {
	s.SniContainerRefs = v
	return s
}

func (s *ListControlGroupResponseListenerList) SetCipherSuites(v string) *ListControlGroupResponseListenerList {
	s.CipherSuites = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetRedirectProtocol(v ListControlGroupResponseListenerListRedirectProtocolEnum) *ListControlGroupResponseListenerList {
	s.RedirectProtocol = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetRedirectUp(v bool) *ListControlGroupResponseListenerList {
	s.RedirectUp = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetCaContainerId(v string) *ListControlGroupResponseListenerList {
	s.CaContainerId = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetName(v string) *ListControlGroupResponseListenerList {
	s.Name = &v
	return s
}

func (s *ListControlGroupResponseListenerList) SetLoadbalancers(v []ListControlGroupResponseLoadbalancers) *ListControlGroupResponseListenerList {
	s.Loadbalancers = &v
	return s
}


type ListControlGroupResponseListenerListBuilder struct {
	s *ListControlGroupResponseListenerList
}

func NewListControlGroupResponseListenerListBuilder() *ListControlGroupResponseListenerListBuilder {
	s := &ListControlGroupResponseListenerList{}
	b := &ListControlGroupResponseListenerListBuilder{s: s}
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) TenantId(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.TenantId = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) DefaultTlsContainerRef(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.DefaultTlsContainerRef = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) TlsProtocols(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) LoadbalancerId(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.LoadbalancerId = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) ProxyProtocol(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Description(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Transparent(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) ConnectionLimit(v int32) *ListControlGroupResponseListenerListBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) MutualAuthenticationUp(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Protocol(v ListControlGroupResponseListenerListProtocolEnum) *ListControlGroupResponseListenerListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) AdminStateUp(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Http2(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) RedirectPort(v int32) *ListControlGroupResponseListenerListBuilder {
	b.s.RedirectPort = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Id(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) DefaultPoolId(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.DefaultPoolId = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) DefaultTlsContainerId(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) ProtocolPort(v int32) *ListControlGroupResponseListenerListBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) KeepaliveTimeout(v int32) *ListControlGroupResponseListenerListBuilder {
	b.s.KeepaliveTimeout = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) SniContainerRefs(v []string) *ListControlGroupResponseListenerListBuilder {
	b.s.SniContainerRefs = v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) CipherSuites(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) RedirectProtocol(v ListControlGroupResponseListenerListRedirectProtocolEnum) *ListControlGroupResponseListenerListBuilder {
	b.s.RedirectProtocol = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) RedirectUp(v bool) *ListControlGroupResponseListenerListBuilder {
	b.s.RedirectUp = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) CaContainerId(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Name(v string) *ListControlGroupResponseListenerListBuilder {
	b.s.Name = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Loadbalancers(v []ListControlGroupResponseLoadbalancers) *ListControlGroupResponseListenerListBuilder {
	b.s.Loadbalancers = &v
	return b
}

func (b *ListControlGroupResponseListenerListBuilder) Build() *ListControlGroupResponseListenerList {
	return b.s
}


