// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPSListenerBody struct {
    position.Body
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // 绑定的访问控制组类型。黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds []string `json:"sniContainerIds,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 绑定的访问控制组的名字
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
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetSniUp(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.SniUp = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetGroupType(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.GroupType = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetKeepAliveTimeout(v int32) *UpdateLoadBalanceHTTPSListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetNotes(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.Notes = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetTlsUp(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.TlsUp = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetSniContainerIds(v []string) *UpdateLoadBalanceHTTPSListenerBody {
	s.SniContainerIds = v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetTransparent(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.Transparent = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetGroupName(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.GroupName = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetName(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetPoolId(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetGroupEnabled(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetHttp2(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.Http2 = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetCaContainerId(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.CaContainerId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetId(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetControlGroupId(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetDefaultTlsContainerId(v string) *UpdateLoadBalanceHTTPSListenerBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetMutualAuthenticationUp(v bool) *UpdateLoadBalanceHTTPSListenerBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerBody) SetTlsCustomizeProtocolId(v int64) *UpdateLoadBalanceHTTPSListenerBody {
	s.TlsCustomizeProtocolId = &v
	return s
}


type UpdateLoadBalanceHTTPSListenerBodyBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerBody
}

func NewUpdateLoadBalanceHTTPSListenerBodyBuilder() *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	s := &UpdateLoadBalanceHTTPSListenerBody{}
	b := &UpdateLoadBalanceHTTPSListenerBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) SniUp(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) GroupType(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) KeepAliveTimeout(v int32) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Notes(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) TlsUp(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) SniContainerIds(v []string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Transparent(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) GroupName(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Name(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) PoolId(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) GroupEnabled(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Http2(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) CaContainerId(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Id(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) ControlGroupId(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) DefaultTlsContainerId(v string) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) MutualAuthenticationUp(v bool) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) TlsCustomizeProtocolId(v int64) *UpdateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerBodyBuilder) Build() *UpdateLoadBalanceHTTPSListenerBody {
	return b.s
}


