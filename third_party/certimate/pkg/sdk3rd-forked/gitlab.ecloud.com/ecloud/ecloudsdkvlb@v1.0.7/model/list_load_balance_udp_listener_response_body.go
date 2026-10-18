// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceUDPListenerResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadBalanceUDPListenerResponseContent `json:"content,omitempty"`
}

func (s ListLoadBalanceUDPListenerResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerResponseBody) SetTotal(v int32) *ListLoadBalanceUDPListenerResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseBody) SetContent(v []ListLoadBalanceUDPListenerResponseContent) *ListLoadBalanceUDPListenerResponseBody {
	s.Content = &v
	return s
}


type ListLoadBalanceUDPListenerResponseBodyBuilder struct {
	s *ListLoadBalanceUDPListenerResponseBody
}

func NewListLoadBalanceUDPListenerResponseBodyBuilder() *ListLoadBalanceUDPListenerResponseBodyBuilder {
	s := &ListLoadBalanceUDPListenerResponseBody{}
	b := &ListLoadBalanceUDPListenerResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBodyBuilder) Total(v int32) *ListLoadBalanceUDPListenerResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBodyBuilder) Content(v []ListLoadBalanceUDPListenerResponseContent) *ListLoadBalanceUDPListenerResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseBodyBuilder) Build() *ListLoadBalanceUDPListenerResponseBody {
	return b.s
}


