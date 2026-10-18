// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateWithHttpStatusBody struct {
    position.Body
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // 绑定的访问控制组类型。黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // TCP长连接超时时间，取值为10-3600
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 重定向到的监听器ID，不是Null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds []string `json:"sniContainerIds,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 修改监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 证书容器中的CA证书ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
    // 是否开启重定向
	RedirectEnabled *bool `json:"redirectEnabled,omitempty"`
    // 访问控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // TLS安全策略ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
}

func (s AsyncUpdateWithHttpStatusBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateWithHttpStatusBody) GoString() string {
	return s.String()
}

func (s AsyncUpdateWithHttpStatusBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateWithHttpStatusBody) SetSniUp(v bool) *AsyncUpdateWithHttpStatusBody {
	s.SniUp = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetGroupType(v string) *AsyncUpdateWithHttpStatusBody {
	s.GroupType = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetKeepAliveTimeout(v int32) *AsyncUpdateWithHttpStatusBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetTlsUp(v bool) *AsyncUpdateWithHttpStatusBody {
	s.TlsUp = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetProxyProtocol(v bool) *AsyncUpdateWithHttpStatusBody {
	s.ProxyProtocol = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetRedirectToListenerId(v string) *AsyncUpdateWithHttpStatusBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetSniContainerIds(v []string) *AsyncUpdateWithHttpStatusBody {
	s.SniContainerIds = v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetTransparent(v bool) *AsyncUpdateWithHttpStatusBody {
	s.Transparent = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetName(v string) *AsyncUpdateWithHttpStatusBody {
	s.Name = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetPoolId(v string) *AsyncUpdateWithHttpStatusBody {
	s.PoolId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetGroupEnabled(v bool) *AsyncUpdateWithHttpStatusBody {
	s.GroupEnabled = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetHttp2(v bool) *AsyncUpdateWithHttpStatusBody {
	s.Http2 = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetCaContainerId(v string) *AsyncUpdateWithHttpStatusBody {
	s.CaContainerId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetId(v string) *AsyncUpdateWithHttpStatusBody {
	s.Id = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetRedirectEnabled(v bool) *AsyncUpdateWithHttpStatusBody {
	s.RedirectEnabled = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetControlGroupId(v string) *AsyncUpdateWithHttpStatusBody {
	s.ControlGroupId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetDefaultTlsContainerId(v string) *AsyncUpdateWithHttpStatusBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetMutualAuthenticationUp(v bool) *AsyncUpdateWithHttpStatusBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *AsyncUpdateWithHttpStatusBody) SetTlsCustomizeProtocolId(v int64) *AsyncUpdateWithHttpStatusBody {
	s.TlsCustomizeProtocolId = &v
	return s
}


type AsyncUpdateWithHttpStatusBodyBuilder struct {
	s *AsyncUpdateWithHttpStatusBody
}

func NewAsyncUpdateWithHttpStatusBodyBuilder() *AsyncUpdateWithHttpStatusBodyBuilder {
	s := &AsyncUpdateWithHttpStatusBody{}
	b := &AsyncUpdateWithHttpStatusBodyBuilder{s: s}
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) SniUp(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) GroupType(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) KeepAliveTimeout(v int32) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) TlsUp(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) ProxyProtocol(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) RedirectToListenerId(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) SniContainerIds(v []string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) Transparent(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) Name(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) PoolId(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) GroupEnabled(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) Http2(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) CaContainerId(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) Id(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) RedirectEnabled(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.RedirectEnabled = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) ControlGroupId(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) DefaultTlsContainerId(v string) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) MutualAuthenticationUp(v bool) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) TlsCustomizeProtocolId(v int64) *AsyncUpdateWithHttpStatusBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *AsyncUpdateWithHttpStatusBodyBuilder) Build() *AsyncUpdateWithHttpStatusBody {
	return b.s
}


