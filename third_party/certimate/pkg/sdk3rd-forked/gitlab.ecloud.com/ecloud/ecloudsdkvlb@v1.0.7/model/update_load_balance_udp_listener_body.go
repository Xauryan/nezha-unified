// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceUDPListenerBody struct {
    position.Body
    // 绑定的访问控制组的名字
	GroupName *string `json:"groupName,omitempty"`
    // 绑定的访问控制组类型。黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
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
}

func (s UpdateLoadBalanceUDPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerBody) SetGroupName(v string) *UpdateLoadBalanceUDPListenerBody {
	s.GroupName = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetGroupType(v string) *UpdateLoadBalanceUDPListenerBody {
	s.GroupType = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetKeepAliveTimeout(v int32) *UpdateLoadBalanceUDPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetNotes(v string) *UpdateLoadBalanceUDPListenerBody {
	s.Notes = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetName(v string) *UpdateLoadBalanceUDPListenerBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetPoolId(v string) *UpdateLoadBalanceUDPListenerBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetGroupEnabled(v bool) *UpdateLoadBalanceUDPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetId(v string) *UpdateLoadBalanceUDPListenerBody {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerBody) SetControlGroupId(v string) *UpdateLoadBalanceUDPListenerBody {
	s.ControlGroupId = &v
	return s
}


type UpdateLoadBalanceUDPListenerBodyBuilder struct {
	s *UpdateLoadBalanceUDPListenerBody
}

func NewUpdateLoadBalanceUDPListenerBodyBuilder() *UpdateLoadBalanceUDPListenerBodyBuilder {
	s := &UpdateLoadBalanceUDPListenerBody{}
	b := &UpdateLoadBalanceUDPListenerBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) GroupName(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) GroupType(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) KeepAliveTimeout(v int32) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) Notes(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) Name(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) PoolId(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) GroupEnabled(v bool) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) Id(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) ControlGroupId(v string) *UpdateLoadBalanceUDPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerBodyBuilder) Build() *UpdateLoadBalanceUDPListenerBody {
	return b.s
}


