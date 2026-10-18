// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceTCPListenerBody struct {
    position.Body
    // 绑定的访问控制组的名字
	GroupName *string `json:"groupName,omitempty"`
    // 绑定的访问控制组类型。黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 修改监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
}

func (s UpdateLoadBalanceTCPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceTCPListenerBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceTCPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceTCPListenerBody) SetGroupName(v string) *UpdateLoadBalanceTCPListenerBody {
	s.GroupName = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetGroupType(v string) *UpdateLoadBalanceTCPListenerBody {
	s.GroupType = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetKeepAliveTimeout(v int32) *UpdateLoadBalanceTCPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetNotes(v string) *UpdateLoadBalanceTCPListenerBody {
	s.Notes = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetProxyProtocol(v bool) *UpdateLoadBalanceTCPListenerBody {
	s.ProxyProtocol = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetName(v string) *UpdateLoadBalanceTCPListenerBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetPoolId(v string) *UpdateLoadBalanceTCPListenerBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetGroupEnabled(v bool) *UpdateLoadBalanceTCPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetId(v string) *UpdateLoadBalanceTCPListenerBody {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetControlGroupId(v string) *UpdateLoadBalanceTCPListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerBody) SetTransparent(v bool) *UpdateLoadBalanceTCPListenerBody {
	s.Transparent = &v
	return s
}


type UpdateLoadBalanceTCPListenerBodyBuilder struct {
	s *UpdateLoadBalanceTCPListenerBody
}

func NewUpdateLoadBalanceTCPListenerBodyBuilder() *UpdateLoadBalanceTCPListenerBodyBuilder {
	s := &UpdateLoadBalanceTCPListenerBody{}
	b := &UpdateLoadBalanceTCPListenerBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) GroupName(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) GroupType(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) KeepAliveTimeout(v int32) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) Notes(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) ProxyProtocol(v bool) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) Name(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) PoolId(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) GroupEnabled(v bool) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) Id(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) ControlGroupId(v string) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) Transparent(v bool) *UpdateLoadBalanceTCPListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerBodyBuilder) Build() *UpdateLoadBalanceTCPListenerBody {
	return b.s
}


