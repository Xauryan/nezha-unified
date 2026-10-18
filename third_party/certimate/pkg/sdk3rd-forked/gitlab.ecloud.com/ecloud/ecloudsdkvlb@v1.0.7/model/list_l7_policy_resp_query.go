// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListL7PolicyRespQuery struct {
    position.Query
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
}

func (s ListL7PolicyRespQuery) String() string {
	return utils.Beautify(s)
}

func (s ListL7PolicyRespQuery) GoString() string {
	return s.String()
}

func (s ListL7PolicyRespQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListL7PolicyRespQuery) SetListenerId(v string) *ListL7PolicyRespQuery {
	s.ListenerId = &v
	return s
}

func (s *ListL7PolicyRespQuery) SetPoolId(v string) *ListL7PolicyRespQuery {
	s.PoolId = &v
	return s
}

func (s *ListL7PolicyRespQuery) SetPageSize(v int32) *ListL7PolicyRespQuery {
	s.PageSize = &v
	return s
}

func (s *ListL7PolicyRespQuery) SetPage(v int32) *ListL7PolicyRespQuery {
	s.Page = &v
	return s
}


type ListL7PolicyRespQueryBuilder struct {
	s *ListL7PolicyRespQuery
}

func NewListL7PolicyRespQueryBuilder() *ListL7PolicyRespQueryBuilder {
	s := &ListL7PolicyRespQuery{}
	b := &ListL7PolicyRespQueryBuilder{s: s}
	return b
}

func (b *ListL7PolicyRespQueryBuilder) ListenerId(v string) *ListL7PolicyRespQueryBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *ListL7PolicyRespQueryBuilder) PoolId(v string) *ListL7PolicyRespQueryBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListL7PolicyRespQueryBuilder) PageSize(v int32) *ListL7PolicyRespQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListL7PolicyRespQueryBuilder) Page(v int32) *ListL7PolicyRespQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListL7PolicyRespQueryBuilder) Build() *ListL7PolicyRespQuery {
	return b.s
}


