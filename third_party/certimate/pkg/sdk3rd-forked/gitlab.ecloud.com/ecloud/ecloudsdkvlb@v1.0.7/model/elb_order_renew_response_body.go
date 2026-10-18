// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderRenewResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
    ElbOrderRenewResponseBodyProcedureCodeEnumPrepaid ElbOrderRenewResponseBodyProcedureCodeEnum = "PREPAID"
    ElbOrderRenewResponseBodyProcedureCodeEnumInternalapproval ElbOrderRenewResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
    ElbOrderRenewResponseBodyProcedureCodeEnumPostpaid ElbOrderRenewResponseBodyProcedureCodeEnum = "POSTPAID"
)

type ElbOrderRenewResponseBody struct {

    // 订单号
	OrderId *string `json:"orderId,omitempty"`
    // 付费类型
	ProcedureCode *ElbOrderRenewResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`
    // 续订下单返回的订单项集合
	Products *[]ElbOrderRenewResponseProducts `json:"products,omitempty"`
}

func (s ElbOrderRenewResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderRenewResponseBody) GoString() string {
	return s.String()
}

func (s ElbOrderRenewResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderRenewResponseBody) SetOrderId(v string) *ElbOrderRenewResponseBody {
	s.OrderId = &v
	return s
}

func (s *ElbOrderRenewResponseBody) SetProcedureCode(v ElbOrderRenewResponseBodyProcedureCodeEnum) *ElbOrderRenewResponseBody {
	s.ProcedureCode = &v
	return s
}

func (s *ElbOrderRenewResponseBody) SetProducts(v []ElbOrderRenewResponseProducts) *ElbOrderRenewResponseBody {
	s.Products = &v
	return s
}


type ElbOrderRenewResponseBodyBuilder struct {
	s *ElbOrderRenewResponseBody
}

func NewElbOrderRenewResponseBodyBuilder() *ElbOrderRenewResponseBodyBuilder {
	s := &ElbOrderRenewResponseBody{}
	b := &ElbOrderRenewResponseBodyBuilder{s: s}
	return b
}

func (b *ElbOrderRenewResponseBodyBuilder) OrderId(v string) *ElbOrderRenewResponseBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *ElbOrderRenewResponseBodyBuilder) ProcedureCode(v ElbOrderRenewResponseBodyProcedureCodeEnum) *ElbOrderRenewResponseBodyBuilder {
	b.s.ProcedureCode = &v
	return b
}

func (b *ElbOrderRenewResponseBodyBuilder) Products(v []ElbOrderRenewResponseProducts) *ElbOrderRenewResponseBodyBuilder {
	b.s.Products = &v
	return b
}

func (b *ElbOrderRenewResponseBodyBuilder) Build() *ElbOrderRenewResponseBody {
	return b.s
}


