// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPListenerBody struct {
    position.Body
    // 绑定的访问控制组的名字
	GroupName *string `json:"groupName,omitempty"`
    // 绑定的访问控制组类型。黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 重定向到的HTTP监听器ID，不是Null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 修改监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
    // 是否开启重定向
	RedirectEnabled *bool `json:"redirectEnabled,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetGroupName(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.GroupName = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetGroupType(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.GroupType = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetKeepAliveTimeout(v int32) *UpdateLoadBalanceHTTPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetNotes(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.Notes = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetRedirectToListenerId(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetName(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetPoolId(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetGroupEnabled(v bool) *UpdateLoadBalanceHTTPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetId(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.Id = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetRedirectEnabled(v bool) *UpdateLoadBalanceHTTPListenerBody {
	s.RedirectEnabled = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetControlGroupId(v string) *UpdateLoadBalanceHTTPListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerBody) SetTransparent(v bool) *UpdateLoadBalanceHTTPListenerBody {
	s.Transparent = &v
	return s
}


type UpdateLoadBalanceHTTPListenerBodyBuilder struct {
	s *UpdateLoadBalanceHTTPListenerBody
}

func NewUpdateLoadBalanceHTTPListenerBodyBuilder() *UpdateLoadBalanceHTTPListenerBodyBuilder {
	s := &UpdateLoadBalanceHTTPListenerBody{}
	b := &UpdateLoadBalanceHTTPListenerBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) GroupName(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) GroupType(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) KeepAliveTimeout(v int32) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) Notes(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) RedirectToListenerId(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) Name(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) PoolId(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) GroupEnabled(v bool) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) Id(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) RedirectEnabled(v bool) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.RedirectEnabled = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) ControlGroupId(v string) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) Transparent(v bool) *UpdateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerBodyBuilder) Build() *UpdateLoadBalanceHTTPListenerBody {
	return b.s
}


