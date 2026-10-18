// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderRenewResponseProducts struct {

    // 弹性负载均衡的ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单结束时间
	EndDate *string `json:"endDate,omitempty"`
    // 续订的产品订单项ID
	OrderExtId *string `json:"orderExtId,omitempty"`
}

func (s ElbOrderRenewResponseProducts) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderRenewResponseProducts) GoString() string {
	return s.String()
}

func (s ElbOrderRenewResponseProducts) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderRenewResponseProducts) SetInstanceId(v string) *ElbOrderRenewResponseProducts {
	s.InstanceId = &v
	return s
}

func (s *ElbOrderRenewResponseProducts) SetEndDate(v string) *ElbOrderRenewResponseProducts {
	s.EndDate = &v
	return s
}

func (s *ElbOrderRenewResponseProducts) SetOrderExtId(v string) *ElbOrderRenewResponseProducts {
	s.OrderExtId = &v
	return s
}


type ElbOrderRenewResponseProductsBuilder struct {
	s *ElbOrderRenewResponseProducts
}

func NewElbOrderRenewResponseProductsBuilder() *ElbOrderRenewResponseProductsBuilder {
	s := &ElbOrderRenewResponseProducts{}
	b := &ElbOrderRenewResponseProductsBuilder{s: s}
	return b
}

func (b *ElbOrderRenewResponseProductsBuilder) InstanceId(v string) *ElbOrderRenewResponseProductsBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *ElbOrderRenewResponseProductsBuilder) EndDate(v string) *ElbOrderRenewResponseProductsBuilder {
	b.s.EndDate = &v
	return b
}

func (b *ElbOrderRenewResponseProductsBuilder) OrderExtId(v string) *ElbOrderRenewResponseProductsBuilder {
	b.s.OrderExtId = &v
	return b
}

func (b *ElbOrderRenewResponseProductsBuilder) Build() *ElbOrderRenewResponseProducts {
	return b.s
}


