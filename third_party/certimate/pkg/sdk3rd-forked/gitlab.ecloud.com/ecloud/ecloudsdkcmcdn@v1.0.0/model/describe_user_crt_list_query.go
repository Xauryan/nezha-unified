// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserCrtListQuery struct {
    position.Query
    // 一页中个数，默认为10
	PageSize *string `json:"pageSize,omitempty"`
    // 页数，默认为1
	Page *string `json:"page,omitempty"`
    // 搜索关键字
	Keyword *string `json:"keyword,omitempty"`
}

func (s DescribeUserCrtListQuery) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserCrtListQuery) GoString() string {
	return s.String()
}

func (s DescribeUserCrtListQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserCrtListQuery) SetPageSize(v string) *DescribeUserCrtListQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeUserCrtListQuery) SetPage(v string) *DescribeUserCrtListQuery {
	s.Page = &v
	return s
}

func (s *DescribeUserCrtListQuery) SetKeyword(v string) *DescribeUserCrtListQuery {
	s.Keyword = &v
	return s
}


type DescribeUserCrtListQueryBuilder struct {
	s *DescribeUserCrtListQuery
}

func NewDescribeUserCrtListQueryBuilder() *DescribeUserCrtListQueryBuilder {
	s := &DescribeUserCrtListQuery{}
	b := &DescribeUserCrtListQueryBuilder{s: s}
	return b
}

func (b *DescribeUserCrtListQueryBuilder) PageSize(v string) *DescribeUserCrtListQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *DescribeUserCrtListQueryBuilder) Page(v string) *DescribeUserCrtListQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *DescribeUserCrtListQueryBuilder) Keyword(v string) *DescribeUserCrtListQueryBuilder {
	b.s.Keyword = &v
	return b
}

func (b *DescribeUserCrtListQueryBuilder) Build() *DescribeUserCrtListQuery {
	return b.s
}


