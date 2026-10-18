// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateListenerBody struct {
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
    // 绑定的访问控制组名字
	GroupName *string `json:"groupName,omitempty"`
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
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
}

func (s UpdateListenerBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateListenerBody) GoString() string {
	return s.String()
}

func (s UpdateListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateListenerBody) SetSniUp(v bool) *UpdateListenerBody {
	s.SniUp = &v
	return s
}

func (s *UpdateListenerBody) SetGroupType(v string) *UpdateListenerBody {
	s.GroupType = &v
	return s
}

func (s *UpdateListenerBody) SetKeepAliveTimeout(v int32) *UpdateListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *UpdateListenerBody) SetTlsUp(v bool) *UpdateListenerBody {
	s.TlsUp = &v
	return s
}

func (s *UpdateListenerBody) SetProxyProtocol(v bool) *UpdateListenerBody {
	s.ProxyProtocol = &v
	return s
}

func (s *UpdateListenerBody) SetRedirectToListenerId(v string) *UpdateListenerBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *UpdateListenerBody) SetSniContainerIds(v []string) *UpdateListenerBody {
	s.SniContainerIds = v
	return s
}

func (s *UpdateListenerBody) SetTransparent(v bool) *UpdateListenerBody {
	s.Transparent = &v
	return s
}

func (s *UpdateListenerBody) SetGroupName(v string) *UpdateListenerBody {
	s.GroupName = &v
	return s
}

func (s *UpdateListenerBody) SetName(v string) *UpdateListenerBody {
	s.Name = &v
	return s
}

func (s *UpdateListenerBody) SetPoolId(v string) *UpdateListenerBody {
	s.PoolId = &v
	return s
}

func (s *UpdateListenerBody) SetGroupEnabled(v bool) *UpdateListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *UpdateListenerBody) SetHttp2(v bool) *UpdateListenerBody {
	s.Http2 = &v
	return s
}

func (s *UpdateListenerBody) SetCaContainerId(v string) *UpdateListenerBody {
	s.CaContainerId = &v
	return s
}

func (s *UpdateListenerBody) SetId(v string) *UpdateListenerBody {
	s.Id = &v
	return s
}

func (s *UpdateListenerBody) SetRedirectEnabled(v bool) *UpdateListenerBody {
	s.RedirectEnabled = &v
	return s
}

func (s *UpdateListenerBody) SetControlGroupId(v string) *UpdateListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *UpdateListenerBody) SetDefaultTlsContainerId(v string) *UpdateListenerBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *UpdateListenerBody) SetMutualAuthenticationUp(v bool) *UpdateListenerBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *UpdateListenerBody) SetTlsCustomizeProtocolId(v int64) *UpdateListenerBody {
	s.TlsCustomizeProtocolId = &v
	return s
}


type UpdateListenerBodyBuilder struct {
	s *UpdateListenerBody
}

func NewUpdateListenerBodyBuilder() *UpdateListenerBodyBuilder {
	s := &UpdateListenerBody{}
	b := &UpdateListenerBodyBuilder{s: s}
	return b
}

func (b *UpdateListenerBodyBuilder) SniUp(v bool) *UpdateListenerBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *UpdateListenerBodyBuilder) GroupType(v string) *UpdateListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateListenerBodyBuilder) KeepAliveTimeout(v int32) *UpdateListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *UpdateListenerBodyBuilder) TlsUp(v bool) *UpdateListenerBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *UpdateListenerBodyBuilder) ProxyProtocol(v bool) *UpdateListenerBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *UpdateListenerBodyBuilder) RedirectToListenerId(v string) *UpdateListenerBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) SniContainerIds(v []string) *UpdateListenerBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *UpdateListenerBodyBuilder) Transparent(v bool) *UpdateListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *UpdateListenerBodyBuilder) GroupName(v string) *UpdateListenerBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateListenerBodyBuilder) Name(v string) *UpdateListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateListenerBodyBuilder) PoolId(v string) *UpdateListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) GroupEnabled(v bool) *UpdateListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *UpdateListenerBodyBuilder) Http2(v bool) *UpdateListenerBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *UpdateListenerBodyBuilder) CaContainerId(v string) *UpdateListenerBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) Id(v string) *UpdateListenerBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateListenerBodyBuilder) RedirectEnabled(v bool) *UpdateListenerBodyBuilder {
	b.s.RedirectEnabled = &v
	return b
}

func (b *UpdateListenerBodyBuilder) ControlGroupId(v string) *UpdateListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) DefaultTlsContainerId(v string) *UpdateListenerBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) MutualAuthenticationUp(v bool) *UpdateListenerBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *UpdateListenerBodyBuilder) TlsCustomizeProtocolId(v int64) *UpdateListenerBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *UpdateListenerBodyBuilder) Build() *UpdateListenerBody {
	return b.s
}


