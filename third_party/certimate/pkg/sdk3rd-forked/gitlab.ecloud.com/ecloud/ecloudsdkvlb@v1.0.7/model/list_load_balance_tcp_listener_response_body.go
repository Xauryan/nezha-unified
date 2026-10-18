// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceTCPListenerResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadBalanceTCPListenerResponseContent `json:"content,omitempty"`
}

func (s ListLoadBalanceTCPListenerResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerResponseBody) SetTotal(v int32) *ListLoadBalanceTCPListenerResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseBody) SetContent(v []ListLoadBalanceTCPListenerResponseContent) *ListLoadBalanceTCPListenerResponseBody {
	s.Content = &v
	return s
}


type ListLoadBalanceTCPListenerResponseBodyBuilder struct {
	s *ListLoadBalanceTCPListenerResponseBody
}

func NewListLoadBalanceTCPListenerResponseBodyBuilder() *ListLoadBalanceTCPListenerResponseBodyBuilder {
	s := &ListLoadBalanceTCPListenerResponseBody{}
	b := &ListLoadBalanceTCPListenerResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBodyBuilder) Total(v int32) *ListLoadBalanceTCPListenerResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBodyBuilder) Content(v []ListLoadBalanceTCPListenerResponseContent) *ListLoadBalanceTCPListenerResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseBodyBuilder) Build() *ListLoadBalanceTCPListenerResponseBody {
	return b.s
}


