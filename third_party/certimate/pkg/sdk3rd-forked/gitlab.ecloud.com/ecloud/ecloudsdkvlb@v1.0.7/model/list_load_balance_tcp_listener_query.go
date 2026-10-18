// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceTCPListenerQuery struct {
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

func (s ListLoadBalanceTCPListenerQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerQuery) SetProtocolPorts(v string) *ListLoadBalanceTCPListenerQuery {
	s.ProtocolPorts = &v
	return s
}

func (s *ListLoadBalanceTCPListenerQuery) SetListenerName(v string) *ListLoadBalanceTCPListenerQuery {
	s.ListenerName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerQuery) SetPageSize(v int32) *ListLoadBalanceTCPListenerQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalanceTCPListenerQuery) SetPage(v int32) *ListLoadBalanceTCPListenerQuery {
	s.Page = &v
	return s
}

func (s *ListLoadBalanceTCPListenerQuery) SetMultiAzUuid(v string) *ListLoadBalanceTCPListenerQuery {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceTCPListenerQueryBuilder struct {
	s *ListLoadBalanceTCPListenerQuery
}

func NewListLoadBalanceTCPListenerQueryBuilder() *ListLoadBalanceTCPListenerQueryBuilder {
	s := &ListLoadBalanceTCPListenerQuery{}
	b := &ListLoadBalanceTCPListenerQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) ProtocolPorts(v string) *ListLoadBalanceTCPListenerQueryBuilder {
	b.s.ProtocolPorts = &v
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) ListenerName(v string) *ListLoadBalanceTCPListenerQueryBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) PageSize(v int32) *ListLoadBalanceTCPListenerQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) Page(v int32) *ListLoadBalanceTCPListenerQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) MultiAzUuid(v string) *ListLoadBalanceTCPListenerQueryBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceTCPListenerQueryBuilder) Build() *ListLoadBalanceTCPListenerQuery {
	return b.s
}


