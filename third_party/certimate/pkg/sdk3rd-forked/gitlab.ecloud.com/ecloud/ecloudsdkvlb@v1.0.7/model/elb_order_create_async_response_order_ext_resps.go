// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum string

// List of RelProductType
const (
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumRouter ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "router"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumNatgateway ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "natgateway"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumIpv6bandwidth ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "ipv6bandwidth"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumIp ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "ip"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumBandwidth ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "bandwidth"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumCbwp ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "cbwp"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumSharedfp ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "sharedfp"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumIpsecvpn ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "ipsecvpn"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumSslvpn ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "sslvpn"
    ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnumElb ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum = "elb"
)
type ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum string

// List of ProductType
const (
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumRouter ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "router"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumNatgateway ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "natgateway"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumIpv6bandwidth ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "ipv6bandwidth"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumIp ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "ip"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumBandwidth ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "bandwidth"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumCbwp ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "cbwp"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumSharedfp ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "sharedfp"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumIpsecvpn ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "ipsecvpn"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumSslvpn ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "sslvpn"
    ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnumElb ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum = "elb"
)

type ElbOrderCreateAsyncResponseOrderExtResps struct {

    // API订购的关联产品订单项ID, 如:公网ID产品关联的带宽产品订单项ID，弹性负载均衡暂不支持组合订购暂不支持该参数
	RelOrderExtId *string `json:"relOrderExtId,omitempty"`
    // API订购的产品订单项ID
	OrderExtId *string `json:"orderExtId,omitempty"`
    // API订购的关联产品类型, 如:公网IP产品关联的带宽产品
	RelProductType *ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum `json:"relProductType,omitempty"`
    // API订购的关联产品订单项, 如:公网IP关联的带宽订单项状态
	RelOrderExtStatus *int32 `json:"relOrderExtStatus,omitempty"`
    // API订购的产品订单项状态
	OrderExtStatus *int32 `json:"orderExtStatus,omitempty"`
    // API订购的产品类型
	ProductType *ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum `json:"productType,omitempty"`
}

func (s ElbOrderCreateAsyncResponseOrderExtResps) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncResponseOrderExtResps) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncResponseOrderExtResps) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetRelOrderExtId(v string) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.RelOrderExtId = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetOrderExtId(v string) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.OrderExtId = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetRelProductType(v ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.RelProductType = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetRelOrderExtStatus(v int32) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.RelOrderExtStatus = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetOrderExtStatus(v int32) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.OrderExtStatus = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseOrderExtResps) SetProductType(v ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum) *ElbOrderCreateAsyncResponseOrderExtResps {
	s.ProductType = &v
	return s
}


type ElbOrderCreateAsyncResponseOrderExtRespsBuilder struct {
	s *ElbOrderCreateAsyncResponseOrderExtResps
}

func NewElbOrderCreateAsyncResponseOrderExtRespsBuilder() *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	s := &ElbOrderCreateAsyncResponseOrderExtResps{}
	b := &ElbOrderCreateAsyncResponseOrderExtRespsBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) RelOrderExtId(v string) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.RelOrderExtId = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) OrderExtId(v string) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.OrderExtId = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) RelProductType(v ElbOrderCreateAsyncResponseOrderExtRespsRelProductTypeEnum) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.RelProductType = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) RelOrderExtStatus(v int32) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.RelOrderExtStatus = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) OrderExtStatus(v int32) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.OrderExtStatus = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) ProductType(v ElbOrderCreateAsyncResponseOrderExtRespsProductTypeEnum) *ElbOrderCreateAsyncResponseOrderExtRespsBuilder {
	b.s.ProductType = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseOrderExtRespsBuilder) Build() *ElbOrderCreateAsyncResponseOrderExtResps {
	return b.s
}


