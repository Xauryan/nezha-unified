// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalanceListenerRespQueryProtocolsEnum string

// List of Protocols
const (
    ListLoadBalanceListenerRespQueryProtocolsEnumHttp ListLoadBalanceListenerRespQueryProtocolsEnum = "HTTP"
    ListLoadBalanceListenerRespQueryProtocolsEnumHttps ListLoadBalanceListenerRespQueryProtocolsEnum = "HTTPS"
    ListLoadBalanceListenerRespQueryProtocolsEnumTcp ListLoadBalanceListenerRespQueryProtocolsEnum = "TCP"
    ListLoadBalanceListenerRespQueryProtocolsEnumUdp ListLoadBalanceListenerRespQueryProtocolsEnum = "UDP"
    ListLoadBalanceListenerRespQueryProtocolsEnumTerminatedHttps ListLoadBalanceListenerRespQueryProtocolsEnum = "TERMINATED_HTTPS"
    ListLoadBalanceListenerRespQueryProtocolsEnumSip ListLoadBalanceListenerRespQueryProtocolsEnum = "SIP"
)

type ListLoadBalanceListenerRespQuery struct {
    position.Query
    // 实例使用的监听端口
	ProtocolPorts *string `json:"protocolPorts,omitempty"`
    // 证书（容器）UUID
	ContainerUuid *string `json:"containerUuid,omitempty"`
    // 监听器名称
	ListenerName *string `json:"listenerName,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 协议
	Protocols []ListLoadBalanceListenerRespQueryProtocolsEnum `json:"protocols,omitempty"`
    // 双AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
}

func (s ListLoadBalanceListenerRespQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespQuery) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespQuery) SetProtocolPorts(v string) *ListLoadBalanceListenerRespQuery {
	s.ProtocolPorts = &v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetContainerUuid(v string) *ListLoadBalanceListenerRespQuery {
	s.ContainerUuid = &v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetListenerName(v string) *ListLoadBalanceListenerRespQuery {
	s.ListenerName = &v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetPageSize(v int32) *ListLoadBalanceListenerRespQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetPage(v int32) *ListLoadBalanceListenerRespQuery {
	s.Page = &v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetProtocols(v []ListLoadBalanceListenerRespQueryProtocolsEnum) *ListLoadBalanceListenerRespQuery {
	s.Protocols = v
	return s
}

func (s *ListLoadBalanceListenerRespQuery) SetMultiAzUuid(v string) *ListLoadBalanceListenerRespQuery {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceListenerRespQueryBuilder struct {
	s *ListLoadBalanceListenerRespQuery
}

func NewListLoadBalanceListenerRespQueryBuilder() *ListLoadBalanceListenerRespQueryBuilder {
	s := &ListLoadBalanceListenerRespQuery{}
	b := &ListLoadBalanceListenerRespQueryBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) ProtocolPorts(v string) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.ProtocolPorts = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) ContainerUuid(v string) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) ListenerName(v string) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) PageSize(v int32) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) Page(v int32) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) Protocols(v []ListLoadBalanceListenerRespQueryProtocolsEnum) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.Protocols = v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) MultiAzUuid(v string) *ListLoadBalanceListenerRespQueryBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceListenerRespQueryBuilder) Build() *ListLoadBalanceListenerRespQuery {
	return b.s
}


