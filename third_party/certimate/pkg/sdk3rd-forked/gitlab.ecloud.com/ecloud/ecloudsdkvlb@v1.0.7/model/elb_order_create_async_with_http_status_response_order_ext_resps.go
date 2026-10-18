// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum string

// List of RelProductType
const (
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumRouter ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "router"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumNatgateway ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "natgateway"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumIpv6bandwidth ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "ipv6bandwidth"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumIp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "ip"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumBandwidth ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "bandwidth"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumCbwp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "cbwp"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumSharedfp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "sharedfp"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumIpsecvpn ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "ipsecvpn"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumSslvpn ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "sslvpn"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnumElb ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum = "elb"
)
type ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum string

// List of ProductType
const (
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumRouter ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "router"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumNatgateway ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "natgateway"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumIpv6bandwidth ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "ipv6bandwidth"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumIp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "ip"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumBandwidth ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "bandwidth"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumCbwp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "cbwp"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumSharedfp ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "sharedfp"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumIpsecvpn ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "ipsecvpn"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumSslvpn ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "sslvpn"
    ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnumElb ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum = "elb"
)

type ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps struct {

    // API订购的关联产品订单项ID, 如:公网ip产品关联的带宽产品订单项ID
	RelOrderExtId *string `json:"relOrderExtId,omitempty"`
    // API订购的产品订单项ID
	OrderExtId *string `json:"orderExtId,omitempty"`
    // API订购的关联产品类型, 如:公网IP产品关联的带宽产品
	RelProductType *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum `json:"relProductType,omitempty"`
    // API订购的关联产品订单项, 如:公网IP关联的带宽订单项状态
	RelOrderExtStatus *int32 `json:"relOrderExtStatus,omitempty"`
    // API订购的产品订单项状态
	OrderExtStatus *int32 `json:"orderExtStatus,omitempty"`
    // API订购的产品类型
	ProductType *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum `json:"productType,omitempty"`
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetRelOrderExtId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.RelOrderExtId = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetOrderExtId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.OrderExtId = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetRelProductType(v ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.RelProductType = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetRelOrderExtStatus(v int32) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.RelOrderExtStatus = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetOrderExtStatus(v int32) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.OrderExtStatus = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) SetProductType(v ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	s.ProductType = &v
	return s
}


type ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder struct {
	s *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps
}

func NewElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder() *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	s := &ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps{}
	b := &ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) RelOrderExtId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.RelOrderExtId = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) OrderExtId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.OrderExtId = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) RelProductType(v ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsRelProductTypeEnum) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.RelProductType = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) RelOrderExtStatus(v int32) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.RelOrderExtStatus = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) OrderExtStatus(v int32) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.OrderExtStatus = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) ProductType(v ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsProductTypeEnum) *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder {
	b.s.ProductType = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtRespsBuilder) Build() *ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps {
	return b.s
}


