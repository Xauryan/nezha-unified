// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalancePoolMemberQuery struct {
    position.Query
    // 请输入后端服务器类型对应的数字，0：IP类型，1：云主机，3：裸金属
	Types []int32 `json:"types,omitempty"`
    // 后端服务器IP
	Ip *string `json:"ip,omitempty"`
    // 后端服务器名称
	Name *string `json:"name,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
}

func (s ListLoadBalancePoolMemberQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberQuery) SetTypes(v []int32) *ListLoadBalancePoolMemberQuery {
	s.Types = v
	return s
}

func (s *ListLoadBalancePoolMemberQuery) SetIp(v string) *ListLoadBalancePoolMemberQuery {
	s.Ip = &v
	return s
}

func (s *ListLoadBalancePoolMemberQuery) SetName(v string) *ListLoadBalancePoolMemberQuery {
	s.Name = &v
	return s
}

func (s *ListLoadBalancePoolMemberQuery) SetPageSize(v int32) *ListLoadBalancePoolMemberQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancePoolMemberQuery) SetPage(v int32) *ListLoadBalancePoolMemberQuery {
	s.Page = &v
	return s
}


type ListLoadBalancePoolMemberQueryBuilder struct {
	s *ListLoadBalancePoolMemberQuery
}

func NewListLoadBalancePoolMemberQueryBuilder() *ListLoadBalancePoolMemberQueryBuilder {
	s := &ListLoadBalancePoolMemberQuery{}
	b := &ListLoadBalancePoolMemberQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) Types(v []int32) *ListLoadBalancePoolMemberQueryBuilder {
	b.s.Types = v
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) Ip(v string) *ListLoadBalancePoolMemberQueryBuilder {
	b.s.Ip = &v
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) Name(v string) *ListLoadBalancePoolMemberQueryBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) PageSize(v int32) *ListLoadBalancePoolMemberQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) Page(v int32) *ListLoadBalancePoolMemberQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalancePoolMemberQueryBuilder) Build() *ListLoadBalancePoolMemberQuery {
	return b.s
}


