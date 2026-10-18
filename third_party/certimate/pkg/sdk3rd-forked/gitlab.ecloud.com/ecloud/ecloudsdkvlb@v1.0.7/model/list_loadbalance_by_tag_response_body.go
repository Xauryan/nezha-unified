// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceByTagResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadbalanceByTagResponseContent `json:"content,omitempty"`
}

func (s ListLoadbalanceByTagResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagResponseBody) SetTotal(v int32) *ListLoadbalanceByTagResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBody) SetContent(v []ListLoadbalanceByTagResponseContent) *ListLoadbalanceByTagResponseBody {
	s.Content = &v
	return s
}


type ListLoadbalanceByTagResponseBodyBuilder struct {
	s *ListLoadbalanceByTagResponseBody
}

func NewListLoadbalanceByTagResponseBodyBuilder() *ListLoadbalanceByTagResponseBodyBuilder {
	s := &ListLoadbalanceByTagResponseBody{}
	b := &ListLoadbalanceByTagResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagResponseBodyBuilder) Total(v int32) *ListLoadbalanceByTagResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBodyBuilder) Content(v []ListLoadbalanceByTagResponseContent) *ListLoadbalanceByTagResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBodyBuilder) Build() *ListLoadbalanceByTagResponseBody {
	return b.s
}


