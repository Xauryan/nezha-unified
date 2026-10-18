// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceTCPListenerPoolBody struct {
    position.Body
    // 监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadBalanceTCPListenerPoolBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceTCPListenerPoolBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceTCPListenerPoolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceTCPListenerPoolBody) SetPoolId(v string) *UpdateLoadBalanceTCPListenerPoolBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceTCPListenerPoolBody) SetId(v string) *UpdateLoadBalanceTCPListenerPoolBody {
	s.Id = &v
	return s
}


type UpdateLoadBalanceTCPListenerPoolBodyBuilder struct {
	s *UpdateLoadBalanceTCPListenerPoolBody
}

func NewUpdateLoadBalanceTCPListenerPoolBodyBuilder() *UpdateLoadBalanceTCPListenerPoolBodyBuilder {
	s := &UpdateLoadBalanceTCPListenerPoolBody{}
	b := &UpdateLoadBalanceTCPListenerPoolBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolBodyBuilder) PoolId(v string) *UpdateLoadBalanceTCPListenerPoolBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolBodyBuilder) Id(v string) *UpdateLoadBalanceTCPListenerPoolBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceTCPListenerPoolBodyBuilder) Build() *UpdateLoadBalanceTCPListenerPoolBody {
	return b.s
}


