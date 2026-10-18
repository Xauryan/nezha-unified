// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncWithHttpStatusResponseBody struct {

    // 资源ID
	ResourceId *string `json:"resourceId,omitempty"`
    // API订购返回的订单ID
	OrderId *string `json:"orderId,omitempty"`
    // API订购返回订单包含的订单项列表(包含批量订购)
	OrderExtResps *[]ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps `json:"orderExtResps,omitempty"`
    // API预付费流程返回的支付链接
	PayUrl *string `json:"payUrl,omitempty"`
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseBody) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncWithHttpStatusResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseBody) SetResourceId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBody {
	s.ResourceId = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseBody) SetOrderId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBody {
	s.OrderId = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseBody) SetOrderExtResps(v []ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) *ElbOrderCreateAsyncWithHttpStatusResponseBody {
	s.OrderExtResps = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusResponseBody) SetPayUrl(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBody {
	s.PayUrl = &v
	return s
}


type ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder struct {
	s *ElbOrderCreateAsyncWithHttpStatusResponseBody
}

func NewElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder() *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder {
	s := &ElbOrderCreateAsyncWithHttpStatusResponseBody{}
	b := &ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder) ResourceId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder) OrderId(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder) OrderExtResps(v []ElbOrderCreateAsyncWithHttpStatusResponseOrderExtResps) *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder {
	b.s.OrderExtResps = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder) PayUrl(v string) *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder {
	b.s.PayUrl = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusResponseBodyBuilder) Build() *ElbOrderCreateAsyncWithHttpStatusResponseBody {
	return b.s
}


