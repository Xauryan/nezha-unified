// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceHTTPSListenerQuery struct {
    position.Query
    // 实例使用的监听端口
	ProtocolPorts *string `json:"protocolPorts,omitempty"`
    // 证书（容器）Uuid
	ContainerUuid *string `json:"containerUuid,omitempty"`
    // 监听器名称
	ListenerName *string `json:"listenerName,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 双AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
}

func (s ListLoadBalanceHTTPSListenerQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPSListenerQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPSListenerQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetProtocolPorts(v string) *ListLoadBalanceHTTPSListenerQuery {
	s.ProtocolPorts = &v
	return s
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetContainerUuid(v string) *ListLoadBalanceHTTPSListenerQuery {
	s.ContainerUuid = &v
	return s
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetListenerName(v string) *ListLoadBalanceHTTPSListenerQuery {
	s.ListenerName = &v
	return s
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetPageSize(v int32) *ListLoadBalanceHTTPSListenerQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetPage(v int32) *ListLoadBalanceHTTPSListenerQuery {
	s.Page = &v
	return s
}

func (s *ListLoadBalanceHTTPSListenerQuery) SetMultiAzUuid(v string) *ListLoadBalanceHTTPSListenerQuery {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceHTTPSListenerQueryBuilder struct {
	s *ListLoadBalanceHTTPSListenerQuery
}

func NewListLoadBalanceHTTPSListenerQueryBuilder() *ListLoadBalanceHTTPSListenerQueryBuilder {
	s := &ListLoadBalanceHTTPSListenerQuery{}
	b := &ListLoadBalanceHTTPSListenerQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) ProtocolPorts(v string) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.ProtocolPorts = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) ContainerUuid(v string) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) ListenerName(v string) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) PageSize(v int32) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) Page(v int32) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) MultiAzUuid(v string) *ListLoadBalanceHTTPSListenerQueryBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceHTTPSListenerQueryBuilder) Build() *ListLoadBalanceHTTPSListenerQuery {
	return b.s
}


