// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserCrtListResponseBody struct {


	PageSize *int32 `json:"pageSize,omitempty"`

	Page *int32 `json:"page,omitempty"`

	List *[]DescribeUserCrtListResponseList `json:"list,omitempty"`
}

func (s DescribeUserCrtListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserCrtListResponseBody) GoString() string {
	return s.String()
}

func (s DescribeUserCrtListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserCrtListResponseBody) SetPageSize(v int32) *DescribeUserCrtListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserCrtListResponseBody) SetPage(v int32) *DescribeUserCrtListResponseBody {
	s.Page = &v
	return s
}

func (s *DescribeUserCrtListResponseBody) SetList(v []DescribeUserCrtListResponseList) *DescribeUserCrtListResponseBody {
	s.List = &v
	return s
}


type DescribeUserCrtListResponseBodyBuilder struct {
	s *DescribeUserCrtListResponseBody
}

func NewDescribeUserCrtListResponseBodyBuilder() *DescribeUserCrtListResponseBodyBuilder {
	s := &DescribeUserCrtListResponseBody{}
	b := &DescribeUserCrtListResponseBodyBuilder{s: s}
	return b
}

func (b *DescribeUserCrtListResponseBodyBuilder) PageSize(v int32) *DescribeUserCrtListResponseBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *DescribeUserCrtListResponseBodyBuilder) Page(v int32) *DescribeUserCrtListResponseBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *DescribeUserCrtListResponseBodyBuilder) List(v []DescribeUserCrtListResponseList) *DescribeUserCrtListResponseBodyBuilder {
	b.s.List = &v
	return b
}

func (b *DescribeUserCrtListResponseBodyBuilder) Build() *DescribeUserCrtListResponseBody {
	return b.s
}


