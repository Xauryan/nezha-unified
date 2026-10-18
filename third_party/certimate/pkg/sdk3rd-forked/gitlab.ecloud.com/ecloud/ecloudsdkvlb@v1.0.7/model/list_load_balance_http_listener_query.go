// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPListenerQuery struct {
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

func (s ListLoadBalanceHTTPListenerQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerQuery) SetProtocolPorts(v string) *ListLoadBalanceHTTPListenerQuery {
	s.ProtocolPorts = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerQuery) SetListenerName(v string) *ListLoadBalanceHTTPListenerQuery {
	s.ListenerName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerQuery) SetPageSize(v int32) *ListLoadBalanceHTTPListenerQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerQuery) SetPage(v int32) *ListLoadBalanceHTTPListenerQuery {
	s.Page = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerQuery) SetMultiAzUuid(v string) *ListLoadBalanceHTTPListenerQuery {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceHTTPListenerQueryBuilder struct {
	s *ListLoadBalanceHTTPListenerQuery
}

func NewListLoadBalanceHTTPListenerQueryBuilder() *ListLoadBalanceHTTPListenerQueryBuilder {
	s := &ListLoadBalanceHTTPListenerQuery{}
	b := &ListLoadBalanceHTTPListenerQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) ProtocolPorts(v string) *ListLoadBalanceHTTPListenerQueryBuilder {
	b.s.ProtocolPorts = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) ListenerName(v string) *ListLoadBalanceHTTPListenerQueryBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) PageSize(v int32) *ListLoadBalanceHTTPListenerQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) Page(v int32) *ListLoadBalanceHTTPListenerQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) MultiAzUuid(v string) *ListLoadBalanceHTTPListenerQueryBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerQueryBuilder) Build() *ListLoadBalanceHTTPListenerQuery {
	return b.s
}


