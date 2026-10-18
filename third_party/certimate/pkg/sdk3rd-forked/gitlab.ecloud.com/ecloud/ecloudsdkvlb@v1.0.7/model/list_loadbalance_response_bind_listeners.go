// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceResponseBindListenersProtocolEnum string

// List of Protocol
const (
    ListLoadbalanceResponseBindListenersProtocolEnumHttp ListLoadbalanceResponseBindListenersProtocolEnum = "HTTP"
    ListLoadbalanceResponseBindListenersProtocolEnumHttps ListLoadbalanceResponseBindListenersProtocolEnum = "HTTPS"
    ListLoadbalanceResponseBindListenersProtocolEnumTcp ListLoadbalanceResponseBindListenersProtocolEnum = "TCP"
    ListLoadbalanceResponseBindListenersProtocolEnumUdp ListLoadbalanceResponseBindListenersProtocolEnum = "UDP"
    ListLoadbalanceResponseBindListenersProtocolEnumTerminatedHttps ListLoadbalanceResponseBindListenersProtocolEnum = "TERMINATED_HTTPS"
    ListLoadbalanceResponseBindListenersProtocolEnumSip ListLoadbalanceResponseBindListenersProtocolEnum = "SIP"
)
type ListLoadbalanceResponseBindListenersOpStatusEnum string

// List of OpStatus
const (
    ListLoadbalanceResponseBindListenersOpStatusEnumInactive ListLoadbalanceResponseBindListenersOpStatusEnum = "INACTIVE"
    ListLoadbalanceResponseBindListenersOpStatusEnumActive ListLoadbalanceResponseBindListenersOpStatusEnum = "ACTIVE"
    ListLoadbalanceResponseBindListenersOpStatusEnumCreating ListLoadbalanceResponseBindListenersOpStatusEnum = "CREATING"
    ListLoadbalanceResponseBindListenersOpStatusEnumDeleting ListLoadbalanceResponseBindListenersOpStatusEnum = "DELETING"
    ListLoadbalanceResponseBindListenersOpStatusEnumUpdating ListLoadbalanceResponseBindListenersOpStatusEnum = "UPDATING"
    ListLoadbalanceResponseBindListenersOpStatusEnumError ListLoadbalanceResponseBindListenersOpStatusEnum = "ERROR"
)

type ListLoadbalanceResponseBindListeners struct {

    // 弹性负载均衡规格
	Flavor *string `json:"flavor,omitempty"`
    // 绑定的访问控制组类型: 黑名单，白名单
	GroupType *string `json:"groupType,omitempty"`
    // 监听器协议
	Protocol *ListLoadbalanceResponseBindListenersProtocolEnum `json:"protocol,omitempty"`
    // 弹性负载均衡厂商
	Provider *string `json:"provider,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 访问控制是否生效
	Active *bool `json:"active,omitempty"`
    // 监听器ID
	Id *string `json:"id,omitempty"`
    // 监听器状态
	OpStatus *ListLoadbalanceResponseBindListenersOpStatusEnum `json:"opStatus,omitempty"`
    // 弹性负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 协议端口号
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
}

func (s ListLoadbalanceResponseBindListeners) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceResponseBindListeners) GoString() string {
	return s.String()
}

func (s ListLoadbalanceResponseBindListeners) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceResponseBindListeners) SetFlavor(v string) *ListLoadbalanceResponseBindListeners {
	s.Flavor = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetGroupType(v string) *ListLoadbalanceResponseBindListeners {
	s.GroupType = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetProtocol(v ListLoadbalanceResponseBindListenersProtocolEnum) *ListLoadbalanceResponseBindListeners {
	s.Protocol = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetProvider(v string) *ListLoadbalanceResponseBindListeners {
	s.Provider = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetName(v string) *ListLoadbalanceResponseBindListeners {
	s.Name = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetActive(v bool) *ListLoadbalanceResponseBindListeners {
	s.Active = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetId(v string) *ListLoadbalanceResponseBindListeners {
	s.Id = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetOpStatus(v ListLoadbalanceResponseBindListenersOpStatusEnum) *ListLoadbalanceResponseBindListeners {
	s.OpStatus = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetLoadBalanceId(v string) *ListLoadbalanceResponseBindListeners {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadbalanceResponseBindListeners) SetProtocolPort(v int32) *ListLoadbalanceResponseBindListeners {
	s.ProtocolPort = &v
	return s
}


type ListLoadbalanceResponseBindListenersBuilder struct {
	s *ListLoadbalanceResponseBindListeners
}

func NewListLoadbalanceResponseBindListenersBuilder() *ListLoadbalanceResponseBindListenersBuilder {
	s := &ListLoadbalanceResponseBindListeners{}
	b := &ListLoadbalanceResponseBindListenersBuilder{s: s}
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Flavor(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Flavor = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) GroupType(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Protocol(v ListLoadbalanceResponseBindListenersProtocolEnum) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Provider(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Provider = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Name(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Active(v bool) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Active = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Id(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) OpStatus(v ListLoadbalanceResponseBindListenersOpStatusEnum) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) LoadBalanceId(v string) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) ProtocolPort(v int32) *ListLoadbalanceResponseBindListenersBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadbalanceResponseBindListenersBuilder) Build() *ListLoadbalanceResponseBindListeners {
	return b.s
}


