// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceListenerRespResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadBalanceListenerRespResponseContent `json:"content,omitempty"`
}

func (s ListLoadBalanceListenerRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespResponseBody) SetTotal(v int32) *ListLoadBalanceListenerRespResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseBody) SetContent(v []ListLoadBalanceListenerRespResponseContent) *ListLoadBalanceListenerRespResponseBody {
	s.Content = &v
	return s
}


type ListLoadBalanceListenerRespResponseBodyBuilder struct {
	s *ListLoadBalanceListenerRespResponseBody
}

func NewListLoadBalanceListenerRespResponseBodyBuilder() *ListLoadBalanceListenerRespResponseBodyBuilder {
	s := &ListLoadBalanceListenerRespResponseBody{}
	b := &ListLoadBalanceListenerRespResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespResponseBodyBuilder) Total(v int32) *ListLoadBalanceListenerRespResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBodyBuilder) Content(v []ListLoadBalanceListenerRespResponseContent) *ListLoadBalanceListenerRespResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseBodyBuilder) Build() *ListLoadBalanceListenerRespResponseBody {
	return b.s
}


