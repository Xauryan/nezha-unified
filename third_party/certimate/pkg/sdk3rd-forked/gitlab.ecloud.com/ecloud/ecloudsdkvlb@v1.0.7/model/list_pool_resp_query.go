// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListPoolRespQueryProtocolEnum string

// List of Protocol
const (
    ListPoolRespQueryProtocolEnumHttp ListPoolRespQueryProtocolEnum = "HTTP"
    ListPoolRespQueryProtocolEnumHttps ListPoolRespQueryProtocolEnum = "HTTPS"
    ListPoolRespQueryProtocolEnumTcp ListPoolRespQueryProtocolEnum = "TCP"
    ListPoolRespQueryProtocolEnumUdp ListPoolRespQueryProtocolEnum = "UDP"
    ListPoolRespQueryProtocolEnumTerminatedHttps ListPoolRespQueryProtocolEnum = "TERMINATED_HTTPS"
    ListPoolRespQueryProtocolEnumSip ListPoolRespQueryProtocolEnum = "SIP"
)

type ListPoolRespQueryPoolProtocolsEnum string

// List of PoolProtocols
const (
    ListPoolRespQueryPoolProtocolsEnumHttp ListPoolRespQueryPoolProtocolsEnum = "HTTP"
    ListPoolRespQueryPoolProtocolsEnumHttps ListPoolRespQueryPoolProtocolsEnum = "HTTPS"
    ListPoolRespQueryPoolProtocolsEnumTcp ListPoolRespQueryPoolProtocolsEnum = "TCP"
    ListPoolRespQueryPoolProtocolsEnumUdp ListPoolRespQueryPoolProtocolsEnum = "UDP"
    ListPoolRespQueryPoolProtocolsEnumTerminatedHttps ListPoolRespQueryPoolProtocolsEnum = "TERMINATED_HTTPS"
    ListPoolRespQueryPoolProtocolsEnumSip ListPoolRespQueryPoolProtocolsEnum = "SIP"
)

type ListPoolRespQuery struct {
    position.Query
    // 监听器协议类型
	Protocol *ListPoolRespQueryProtocolEnum `json:"protocol,omitempty"`
    // 后端服务器组协议
	PoolProtocols []ListPoolRespQueryPoolProtocolsEnum `json:"poolProtocols,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 是否绑定弹性负载均衡监听器
	BindListener *bool `json:"bindListener,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ListPoolRespQuery) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespQuery) GoString() string {
	return s.String()
}

func (s ListPoolRespQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespQuery) SetProtocol(v ListPoolRespQueryProtocolEnum) *ListPoolRespQuery {
	s.Protocol = &v
	return s
}

func (s *ListPoolRespQuery) SetPoolProtocols(v []ListPoolRespQueryPoolProtocolsEnum) *ListPoolRespQuery {
	s.PoolProtocols = v
	return s
}

func (s *ListPoolRespQuery) SetPageSize(v int32) *ListPoolRespQuery {
	s.PageSize = &v
	return s
}

func (s *ListPoolRespQuery) SetPage(v int32) *ListPoolRespQuery {
	s.Page = &v
	return s
}

func (s *ListPoolRespQuery) SetBindListener(v bool) *ListPoolRespQuery {
	s.BindListener = &v
	return s
}

func (s *ListPoolRespQuery) SetPoolName(v string) *ListPoolRespQuery {
	s.PoolName = &v
	return s
}


type ListPoolRespQueryBuilder struct {
	s *ListPoolRespQuery
}

func NewListPoolRespQueryBuilder() *ListPoolRespQueryBuilder {
	s := &ListPoolRespQuery{}
	b := &ListPoolRespQueryBuilder{s: s}
	return b
}

func (b *ListPoolRespQueryBuilder) Protocol(v ListPoolRespQueryProtocolEnum) *ListPoolRespQueryBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListPoolRespQueryBuilder) PoolProtocols(v []ListPoolRespQueryPoolProtocolsEnum) *ListPoolRespQueryBuilder {
	b.s.PoolProtocols = v
	return b
}

func (b *ListPoolRespQueryBuilder) PageSize(v int32) *ListPoolRespQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListPoolRespQueryBuilder) Page(v int32) *ListPoolRespQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListPoolRespQueryBuilder) BindListener(v bool) *ListPoolRespQueryBuilder {
	b.s.BindListener = &v
	return b
}

func (b *ListPoolRespQueryBuilder) PoolName(v string) *ListPoolRespQueryBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListPoolRespQueryBuilder) Build() *ListPoolRespQuery {
	return b.s
}


