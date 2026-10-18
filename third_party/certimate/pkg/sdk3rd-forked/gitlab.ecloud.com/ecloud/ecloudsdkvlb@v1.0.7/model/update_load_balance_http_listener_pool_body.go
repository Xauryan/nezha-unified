// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPListenerPoolBody struct {
    position.Body
    // 监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadBalanceHTTPListenerPoolBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPListenerPoolBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPListenerPoolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPListenerPoolBody) SetPoolId(v string) *UpdateLoadBalanceHTTPListenerPoolBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPListenerPoolBody) SetId(v string) *UpdateLoadBalanceHTTPListenerPoolBody {
	s.Id = &v
	return s
}


type UpdateLoadBalanceHTTPListenerPoolBodyBuilder struct {
	s *UpdateLoadBalanceHTTPListenerPoolBody
}

func NewUpdateLoadBalanceHTTPListenerPoolBodyBuilder() *UpdateLoadBalanceHTTPListenerPoolBodyBuilder {
	s := &UpdateLoadBalanceHTTPListenerPoolBody{}
	b := &UpdateLoadBalanceHTTPListenerPoolBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolBodyBuilder) PoolId(v string) *UpdateLoadBalanceHTTPListenerPoolBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolBodyBuilder) Id(v string) *UpdateLoadBalanceHTTPListenerPoolBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceHTTPListenerPoolBodyBuilder) Build() *UpdateLoadBalanceHTTPListenerPoolBody {
	return b.s
}


