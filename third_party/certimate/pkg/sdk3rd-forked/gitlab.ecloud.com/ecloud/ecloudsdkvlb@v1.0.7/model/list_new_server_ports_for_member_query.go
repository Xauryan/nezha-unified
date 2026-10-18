// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListNewServerPortsForMemberQuery struct {
    position.Query
    // 网卡类型，主机--vm，物理机--ironic
	PortType *string `json:"portType,omitempty"`
    // 后端服务器组的ID
	PoolId *string `json:"poolId,omitempty"`
    // 列表分页每页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 可根据主机名称、内网地址进行模糊查询
	QueryWord *string `json:"queryWord,omitempty"`
    // 可根据主机名称、内网地址进行模糊查询，取值为NAME、IP
	QueryType *string `json:"queryType,omitempty"`
}

func (s ListNewServerPortsForMemberQuery) String() string {
	return utils.Beautify(s)
}

func (s ListNewServerPortsForMemberQuery) GoString() string {
	return s.String()
}

func (s ListNewServerPortsForMemberQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListNewServerPortsForMemberQuery) SetPortType(v string) *ListNewServerPortsForMemberQuery {
	s.PortType = &v
	return s
}

func (s *ListNewServerPortsForMemberQuery) SetPoolId(v string) *ListNewServerPortsForMemberQuery {
	s.PoolId = &v
	return s
}

func (s *ListNewServerPortsForMemberQuery) SetPageSize(v int32) *ListNewServerPortsForMemberQuery {
	s.PageSize = &v
	return s
}

func (s *ListNewServerPortsForMemberQuery) SetPage(v int32) *ListNewServerPortsForMemberQuery {
	s.Page = &v
	return s
}

func (s *ListNewServerPortsForMemberQuery) SetQueryWord(v string) *ListNewServerPortsForMemberQuery {
	s.QueryWord = &v
	return s
}

func (s *ListNewServerPortsForMemberQuery) SetQueryType(v string) *ListNewServerPortsForMemberQuery {
	s.QueryType = &v
	return s
}


type ListNewServerPortsForMemberQueryBuilder struct {
	s *ListNewServerPortsForMemberQuery
}

func NewListNewServerPortsForMemberQueryBuilder() *ListNewServerPortsForMemberQueryBuilder {
	s := &ListNewServerPortsForMemberQuery{}
	b := &ListNewServerPortsForMemberQueryBuilder{s: s}
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) PortType(v string) *ListNewServerPortsForMemberQueryBuilder {
	b.s.PortType = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) PoolId(v string) *ListNewServerPortsForMemberQueryBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) PageSize(v int32) *ListNewServerPortsForMemberQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) Page(v int32) *ListNewServerPortsForMemberQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) QueryWord(v string) *ListNewServerPortsForMemberQueryBuilder {
	b.s.QueryWord = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) QueryType(v string) *ListNewServerPortsForMemberQueryBuilder {
	b.s.QueryType = &v
	return b
}

func (b *ListNewServerPortsForMemberQueryBuilder) Build() *ListNewServerPortsForMemberQuery {
	return b.s
}


