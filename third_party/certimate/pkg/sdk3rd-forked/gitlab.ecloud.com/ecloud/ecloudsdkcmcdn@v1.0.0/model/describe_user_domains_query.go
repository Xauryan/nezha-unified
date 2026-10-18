// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DescribeUserDomainsQueryCdnTypeEnum string

// List of CdnType
const (
    DescribeUserDomainsQueryCdnTypeEnumWeb DescribeUserDomainsQueryCdnTypeEnum = "web"
    DescribeUserDomainsQueryCdnTypeEnumDownload DescribeUserDomainsQueryCdnTypeEnum = "download"
    DescribeUserDomainsQueryCdnTypeEnumVideo DescribeUserDomainsQueryCdnTypeEnum = "video"
    DescribeUserDomainsQueryCdnTypeEnumLivestream DescribeUserDomainsQueryCdnTypeEnum = "liveStream"
)

type DescribeUserDomainsQuery struct {
    position.Query
    // 根据业务类型进行筛选
	CdnType *DescribeUserDomainsQueryCdnTypeEnum `json:"cdnType,omitempty"`
    // 一页中个数，默认为10
	PageSize *string `json:"pageSize,omitempty"`
    // 页数，默认为1
	Page *string `json:"page,omitempty"`
    // 模糊查询关键字
	QueryWord *string `json:"queryWord,omitempty"`
}

func (s DescribeUserDomainsQuery) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserDomainsQuery) GoString() string {
	return s.String()
}

func (s DescribeUserDomainsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserDomainsQuery) SetCdnType(v DescribeUserDomainsQueryCdnTypeEnum) *DescribeUserDomainsQuery {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsQuery) SetPageSize(v string) *DescribeUserDomainsQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsQuery) SetPage(v string) *DescribeUserDomainsQuery {
	s.Page = &v
	return s
}

func (s *DescribeUserDomainsQuery) SetQueryWord(v string) *DescribeUserDomainsQuery {
	s.QueryWord = &v
	return s
}


type DescribeUserDomainsQueryBuilder struct {
	s *DescribeUserDomainsQuery
}

func NewDescribeUserDomainsQueryBuilder() *DescribeUserDomainsQueryBuilder {
	s := &DescribeUserDomainsQuery{}
	b := &DescribeUserDomainsQueryBuilder{s: s}
	return b
}

func (b *DescribeUserDomainsQueryBuilder) CdnType(v DescribeUserDomainsQueryCdnTypeEnum) *DescribeUserDomainsQueryBuilder {
	b.s.CdnType = &v
	return b
}

func (b *DescribeUserDomainsQueryBuilder) PageSize(v string) *DescribeUserDomainsQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *DescribeUserDomainsQueryBuilder) Page(v string) *DescribeUserDomainsQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *DescribeUserDomainsQueryBuilder) QueryWord(v string) *DescribeUserDomainsQueryBuilder {
	b.s.QueryWord = &v
	return b
}

func (b *DescribeUserDomainsQueryBuilder) Build() *DescribeUserDomainsQuery {
	return b.s
}


