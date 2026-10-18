// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryLoadBalanceQuotaResponseBody struct {

    // 总配额
	TotalQuota *int32 `json:"totalQuota,omitempty"`
    // 已使用配额
	UsedQuota *int32 `json:"usedQuota,omitempty"`
}

func (s QueryLoadBalanceQuotaResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryLoadBalanceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s QueryLoadBalanceQuotaResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryLoadBalanceQuotaResponseBody) SetTotalQuota(v int32) *QueryLoadBalanceQuotaResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *QueryLoadBalanceQuotaResponseBody) SetUsedQuota(v int32) *QueryLoadBalanceQuotaResponseBody {
	s.UsedQuota = &v
	return s
}


type QueryLoadBalanceQuotaResponseBodyBuilder struct {
	s *QueryLoadBalanceQuotaResponseBody
}

func NewQueryLoadBalanceQuotaResponseBodyBuilder() *QueryLoadBalanceQuotaResponseBodyBuilder {
	s := &QueryLoadBalanceQuotaResponseBody{}
	b := &QueryLoadBalanceQuotaResponseBodyBuilder{s: s}
	return b
}

func (b *QueryLoadBalanceQuotaResponseBodyBuilder) TotalQuota(v int32) *QueryLoadBalanceQuotaResponseBodyBuilder {
	b.s.TotalQuota = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBodyBuilder) UsedQuota(v int32) *QueryLoadBalanceQuotaResponseBodyBuilder {
	b.s.UsedQuota = &v
	return b
}

func (b *QueryLoadBalanceQuotaResponseBodyBuilder) Build() *QueryLoadBalanceQuotaResponseBody {
	return b.s
}


