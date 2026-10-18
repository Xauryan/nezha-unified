// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupQuery struct {
    position.Query
    // 产商，取值：nfvslb，aliyun，其他传值不生效，默认返回常规域
	Provider *string `json:"provider,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，page小于1时，默认返回第一页
	Page *int32 `json:"page,omitempty"`
    // 根据名称或描述模糊查询
	QueryWord *string `json:"queryWord,omitempty"`
    // 虚拟可用区
	Vaz *string `json:"vaz,omitempty"`
}

func (s ListControlGroupQuery) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupQuery) GoString() string {
	return s.String()
}

func (s ListControlGroupQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupQuery) SetProvider(v string) *ListControlGroupQuery {
	s.Provider = &v
	return s
}

func (s *ListControlGroupQuery) SetPageSize(v int32) *ListControlGroupQuery {
	s.PageSize = &v
	return s
}

func (s *ListControlGroupQuery) SetPage(v int32) *ListControlGroupQuery {
	s.Page = &v
	return s
}

func (s *ListControlGroupQuery) SetQueryWord(v string) *ListControlGroupQuery {
	s.QueryWord = &v
	return s
}

func (s *ListControlGroupQuery) SetVaz(v string) *ListControlGroupQuery {
	s.Vaz = &v
	return s
}


type ListControlGroupQueryBuilder struct {
	s *ListControlGroupQuery
}

func NewListControlGroupQueryBuilder() *ListControlGroupQueryBuilder {
	s := &ListControlGroupQuery{}
	b := &ListControlGroupQueryBuilder{s: s}
	return b
}

func (b *ListControlGroupQueryBuilder) Provider(v string) *ListControlGroupQueryBuilder {
	b.s.Provider = &v
	return b
}

func (b *ListControlGroupQueryBuilder) PageSize(v int32) *ListControlGroupQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListControlGroupQueryBuilder) Page(v int32) *ListControlGroupQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListControlGroupQueryBuilder) QueryWord(v string) *ListControlGroupQueryBuilder {
	b.s.QueryWord = &v
	return b
}

func (b *ListControlGroupQueryBuilder) Vaz(v string) *ListControlGroupQueryBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListControlGroupQueryBuilder) Build() *ListControlGroupQuery {
	return b.s
}


