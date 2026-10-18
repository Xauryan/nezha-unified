// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPListenerResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadBalanceHTTPListenerResponseContent `json:"content,omitempty"`
}

func (s ListLoadBalanceHTTPListenerResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerResponseBody) SetTotal(v int32) *ListLoadBalanceHTTPListenerResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseBody) SetContent(v []ListLoadBalanceHTTPListenerResponseContent) *ListLoadBalanceHTTPListenerResponseBody {
	s.Content = &v
	return s
}


type ListLoadBalanceHTTPListenerResponseBodyBuilder struct {
	s *ListLoadBalanceHTTPListenerResponseBody
}

func NewListLoadBalanceHTTPListenerResponseBodyBuilder() *ListLoadBalanceHTTPListenerResponseBodyBuilder {
	s := &ListLoadBalanceHTTPListenerResponseBody{}
	b := &ListLoadBalanceHTTPListenerResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBodyBuilder) Total(v int32) *ListLoadBalanceHTTPListenerResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBodyBuilder) Content(v []ListLoadBalanceHTTPListenerResponseContent) *ListLoadBalanceHTTPListenerResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseBodyBuilder) Build() *ListLoadBalanceHTTPListenerResponseBody {
	return b.s
}


