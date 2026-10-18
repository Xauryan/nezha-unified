// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalancListenerNameBody struct {
    position.Body
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 监听器ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadBalancListenerNameBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalancListenerNameBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalancListenerNameBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalancListenerNameBody) SetName(v string) *UpdateLoadBalancListenerNameBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadBalancListenerNameBody) SetId(v string) *UpdateLoadBalancListenerNameBody {
	s.Id = &v
	return s
}


type UpdateLoadBalancListenerNameBodyBuilder struct {
	s *UpdateLoadBalancListenerNameBody
}

func NewUpdateLoadBalancListenerNameBodyBuilder() *UpdateLoadBalancListenerNameBodyBuilder {
	s := &UpdateLoadBalancListenerNameBody{}
	b := &UpdateLoadBalancListenerNameBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalancListenerNameBodyBuilder) Name(v string) *UpdateLoadBalancListenerNameBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadBalancListenerNameBodyBuilder) Id(v string) *UpdateLoadBalancListenerNameBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalancListenerNameBodyBuilder) Build() *UpdateLoadBalancListenerNameBody {
	return b.s
}


