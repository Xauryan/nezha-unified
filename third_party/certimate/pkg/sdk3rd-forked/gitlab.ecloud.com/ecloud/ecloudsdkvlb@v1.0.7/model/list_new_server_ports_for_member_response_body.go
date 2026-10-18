// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListNewServerPortsForMemberResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListNewServerPortsForMemberResponseContent `json:"content,omitempty"`
}

func (s ListNewServerPortsForMemberResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListNewServerPortsForMemberResponseBody) GoString() string {
	return s.String()
}

func (s ListNewServerPortsForMemberResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListNewServerPortsForMemberResponseBody) SetTotal(v int32) *ListNewServerPortsForMemberResponseBody {
	s.Total = &v
	return s
}

func (s *ListNewServerPortsForMemberResponseBody) SetContent(v []ListNewServerPortsForMemberResponseContent) *ListNewServerPortsForMemberResponseBody {
	s.Content = &v
	return s
}


type ListNewServerPortsForMemberResponseBodyBuilder struct {
	s *ListNewServerPortsForMemberResponseBody
}

func NewListNewServerPortsForMemberResponseBodyBuilder() *ListNewServerPortsForMemberResponseBodyBuilder {
	s := &ListNewServerPortsForMemberResponseBody{}
	b := &ListNewServerPortsForMemberResponseBodyBuilder{s: s}
	return b
}

func (b *ListNewServerPortsForMemberResponseBodyBuilder) Total(v int32) *ListNewServerPortsForMemberResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBodyBuilder) Content(v []ListNewServerPortsForMemberResponseContent) *ListNewServerPortsForMemberResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseBodyBuilder) Build() *ListNewServerPortsForMemberResponseBody {
	return b.s
}


