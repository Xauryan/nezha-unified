// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncResponseBody struct {

    // 资源ID
	ResourceId *string `json:"resourceId,omitempty"`
    // API订购返回的订单ID
	OrderId *string `json:"orderId,omitempty"`
    // API订购返回订单包含的订单项列表(包含批量订购)
	OrderExtResps *[]ElbOrderCreateAsyncResponseOrderExtResps `json:"orderExtResps,omitempty"`
    // API预付费流程返回的支付链接
	PayUrl *string `json:"payUrl,omitempty"`
}

func (s ElbOrderCreateAsyncResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncResponseBody) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncResponseBody) SetResourceId(v string) *ElbOrderCreateAsyncResponseBody {
	s.ResourceId = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseBody) SetOrderId(v string) *ElbOrderCreateAsyncResponseBody {
	s.OrderId = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseBody) SetOrderExtResps(v []ElbOrderCreateAsyncResponseOrderExtResps) *ElbOrderCreateAsyncResponseBody {
	s.OrderExtResps = &v
	return s
}

func (s *ElbOrderCreateAsyncResponseBody) SetPayUrl(v string) *ElbOrderCreateAsyncResponseBody {
	s.PayUrl = &v
	return s
}


type ElbOrderCreateAsyncResponseBodyBuilder struct {
	s *ElbOrderCreateAsyncResponseBody
}

func NewElbOrderCreateAsyncResponseBodyBuilder() *ElbOrderCreateAsyncResponseBodyBuilder {
	s := &ElbOrderCreateAsyncResponseBody{}
	b := &ElbOrderCreateAsyncResponseBodyBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncResponseBodyBuilder) ResourceId(v string) *ElbOrderCreateAsyncResponseBodyBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBodyBuilder) OrderId(v string) *ElbOrderCreateAsyncResponseBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBodyBuilder) OrderExtResps(v []ElbOrderCreateAsyncResponseOrderExtResps) *ElbOrderCreateAsyncResponseBodyBuilder {
	b.s.OrderExtResps = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBodyBuilder) PayUrl(v string) *ElbOrderCreateAsyncResponseBodyBuilder {
	b.s.PayUrl = &v
	return b
}

func (b *ElbOrderCreateAsyncResponseBodyBuilder) Build() *ElbOrderCreateAsyncResponseBody {
	return b.s
}


