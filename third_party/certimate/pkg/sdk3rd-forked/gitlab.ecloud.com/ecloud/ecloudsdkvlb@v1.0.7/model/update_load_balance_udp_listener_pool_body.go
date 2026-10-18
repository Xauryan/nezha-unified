// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceUDPListenerPoolBody struct {
    position.Body
    // 监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadBalanceUDPListenerPoolBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceUDPListenerPoolBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceUDPListenerPoolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceUDPListenerPoolBody) SetPoolId(v string) *UpdateLoadBalanceUDPListenerPoolBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceUDPListenerPoolBody) SetId(v string) *UpdateLoadBalanceUDPListenerPoolBody {
	s.Id = &v
	return s
}


type UpdateLoadBalanceUDPListenerPoolBodyBuilder struct {
	s *UpdateLoadBalanceUDPListenerPoolBody
}

func NewUpdateLoadBalanceUDPListenerPoolBodyBuilder() *UpdateLoadBalanceUDPListenerPoolBodyBuilder {
	s := &UpdateLoadBalanceUDPListenerPoolBody{}
	b := &UpdateLoadBalanceUDPListenerPoolBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolBodyBuilder) PoolId(v string) *UpdateLoadBalanceUDPListenerPoolBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolBodyBuilder) Id(v string) *UpdateLoadBalanceUDPListenerPoolBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceUDPListenerPoolBodyBuilder) Build() *UpdateLoadBalanceUDPListenerPoolBody {
	return b.s
}


