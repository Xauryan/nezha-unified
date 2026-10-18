// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadBalanceHTTPSListenerPoolBody struct {
    position.Body
    // 监听器绑定的后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadBalanceHTTPSListenerPoolBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadBalanceHTTPSListenerPoolBody) GoString() string {
	return s.String()
}

func (s UpdateLoadBalanceHTTPSListenerPoolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadBalanceHTTPSListenerPoolBody) SetPoolId(v string) *UpdateLoadBalanceHTTPSListenerPoolBody {
	s.PoolId = &v
	return s
}

func (s *UpdateLoadBalanceHTTPSListenerPoolBody) SetId(v string) *UpdateLoadBalanceHTTPSListenerPoolBody {
	s.Id = &v
	return s
}


type UpdateLoadBalanceHTTPSListenerPoolBodyBuilder struct {
	s *UpdateLoadBalanceHTTPSListenerPoolBody
}

func NewUpdateLoadBalanceHTTPSListenerPoolBodyBuilder() *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder {
	s := &UpdateLoadBalanceHTTPSListenerPoolBody{}
	b := &UpdateLoadBalanceHTTPSListenerPoolBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder) PoolId(v string) *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder) Id(v string) *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadBalanceHTTPSListenerPoolBodyBuilder) Build() *UpdateLoadBalanceHTTPSListenerPoolBody {
	return b.s
}


