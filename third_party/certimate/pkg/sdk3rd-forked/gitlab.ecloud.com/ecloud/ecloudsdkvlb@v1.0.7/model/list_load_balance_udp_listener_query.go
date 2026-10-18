// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceUDPListenerQuery struct {
    position.Query
    // 实例使用的监听端口
	ProtocolPorts *string `json:"protocolPorts,omitempty"`
    // 监听器名称
	ListenerName *string `json:"listenerName,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 双AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
}

func (s ListLoadBalanceUDPListenerQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerQuery) SetProtocolPorts(v string) *ListLoadBalanceUDPListenerQuery {
	s.ProtocolPorts = &v
	return s
}

func (s *ListLoadBalanceUDPListenerQuery) SetListenerName(v string) *ListLoadBalanceUDPListenerQuery {
	s.ListenerName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerQuery) SetPageSize(v int32) *ListLoadBalanceUDPListenerQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalanceUDPListenerQuery) SetPage(v int32) *ListLoadBalanceUDPListenerQuery {
	s.Page = &v
	return s
}

func (s *ListLoadBalanceUDPListenerQuery) SetMultiAzUuid(v string) *ListLoadBalanceUDPListenerQuery {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceUDPListenerQueryBuilder struct {
	s *ListLoadBalanceUDPListenerQuery
}

func NewListLoadBalanceUDPListenerQueryBuilder() *ListLoadBalanceUDPListenerQueryBuilder {
	s := &ListLoadBalanceUDPListenerQuery{}
	b := &ListLoadBalanceUDPListenerQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) ProtocolPorts(v string) *ListLoadBalanceUDPListenerQueryBuilder {
	b.s.ProtocolPorts = &v
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) ListenerName(v string) *ListLoadBalanceUDPListenerQueryBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) PageSize(v int32) *ListLoadBalanceUDPListenerQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) Page(v int32) *ListLoadBalanceUDPListenerQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) MultiAzUuid(v string) *ListLoadBalanceUDPListenerQueryBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceUDPListenerQueryBuilder) Build() *ListLoadBalanceUDPListenerQuery {
	return b.s
}


