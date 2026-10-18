// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalancePoolMemberResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadBalancePoolMemberResponseContent `json:"content,omitempty"`
}

func (s ListLoadBalancePoolMemberResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberResponseBody) SetTotal(v int32) *ListLoadBalancePoolMemberResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseBody) SetContent(v []ListLoadBalancePoolMemberResponseContent) *ListLoadBalancePoolMemberResponseBody {
	s.Content = &v
	return s
}


type ListLoadBalancePoolMemberResponseBodyBuilder struct {
	s *ListLoadBalancePoolMemberResponseBody
}

func NewListLoadBalancePoolMemberResponseBodyBuilder() *ListLoadBalancePoolMemberResponseBodyBuilder {
	s := &ListLoadBalancePoolMemberResponseBody{}
	b := &ListLoadBalancePoolMemberResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberResponseBodyBuilder) Total(v int32) *ListLoadBalancePoolMemberResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBodyBuilder) Content(v []ListLoadBalancePoolMemberResponseContent) *ListLoadBalancePoolMemberResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseBodyBuilder) Build() *ListLoadBalancePoolMemberResponseBody {
	return b.s
}


