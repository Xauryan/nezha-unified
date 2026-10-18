// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceByTagResponseBindListenersProtocolEnum string

// List of Protocol
const (
    ListLoadbalanceByTagResponseBindListenersProtocolEnumHttp ListLoadbalanceByTagResponseBindListenersProtocolEnum = "HTTP"
    ListLoadbalanceByTagResponseBindListenersProtocolEnumHttps ListLoadbalanceByTagResponseBindListenersProtocolEnum = "HTTPS"
    ListLoadbalanceByTagResponseBindListenersProtocolEnumTcp ListLoadbalanceByTagResponseBindListenersProtocolEnum = "TCP"
    ListLoadbalanceByTagResponseBindListenersProtocolEnumUdp ListLoadbalanceByTagResponseBindListenersProtocolEnum = "UDP"
    ListLoadbalanceByTagResponseBindListenersProtocolEnumTerminatedHttps ListLoadbalanceByTagResponseBindListenersProtocolEnum = "TERMINATED_HTTPS"
    ListLoadbalanceByTagResponseBindListenersProtocolEnumSip ListLoadbalanceByTagResponseBindListenersProtocolEnum = "SIP"
)
type ListLoadbalanceByTagResponseBindListenersOpStatusEnum string

// List of OpStatus
const (
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumInactive ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "INACTIVE"
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumActive ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "ACTIVE"
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumCreating ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "CREATING"
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumDeleting ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "DELETING"
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumUpdating ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "UPDATING"
    ListLoadbalanceByTagResponseBindListenersOpStatusEnumError ListLoadbalanceByTagResponseBindListenersOpStatusEnum = "ERROR"
)

type ListLoadbalanceByTagResponseBindListeners struct {

    // 弹性负载均衡规格
	Flavor *string `json:"flavor,omitempty"`
    // 绑定的访问控制组类型: 黑名单:blacklist,白名单:whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 监听器协议
	Protocol *ListLoadbalanceByTagResponseBindListenersProtocolEnum `json:"protocol,omitempty"`
    // 弹性负载均衡厂商
	Provider *string `json:"provider,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 访问控制是否生效,true:生效;false:未生效
	Active *bool `json:"active,omitempty"`
    // 监听器ID
	Id *string `json:"id,omitempty"`
    // 监听器状态
	OpStatus *ListLoadbalanceByTagResponseBindListenersOpStatusEnum `json:"opStatus,omitempty"`
    // 弹性负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 协议端口号
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
}

func (s ListLoadbalanceByTagResponseBindListeners) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagResponseBindListeners) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagResponseBindListeners) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetFlavor(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.Flavor = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetGroupType(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.GroupType = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetProtocol(v ListLoadbalanceByTagResponseBindListenersProtocolEnum) *ListLoadbalanceByTagResponseBindListeners {
	s.Protocol = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetProvider(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.Provider = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetName(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.Name = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetActive(v bool) *ListLoadbalanceByTagResponseBindListeners {
	s.Active = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetId(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.Id = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetOpStatus(v ListLoadbalanceByTagResponseBindListenersOpStatusEnum) *ListLoadbalanceByTagResponseBindListeners {
	s.OpStatus = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetLoadBalanceId(v string) *ListLoadbalanceByTagResponseBindListeners {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseBindListeners) SetProtocolPort(v int32) *ListLoadbalanceByTagResponseBindListeners {
	s.ProtocolPort = &v
	return s
}


type ListLoadbalanceByTagResponseBindListenersBuilder struct {
	s *ListLoadbalanceByTagResponseBindListeners
}

func NewListLoadbalanceByTagResponseBindListenersBuilder() *ListLoadbalanceByTagResponseBindListenersBuilder {
	s := &ListLoadbalanceByTagResponseBindListeners{}
	b := &ListLoadbalanceByTagResponseBindListenersBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Flavor(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Flavor = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) GroupType(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Protocol(v ListLoadbalanceByTagResponseBindListenersProtocolEnum) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Provider(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Provider = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Name(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Active(v bool) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Active = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Id(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) OpStatus(v ListLoadbalanceByTagResponseBindListenersOpStatusEnum) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) LoadBalanceId(v string) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) ProtocolPort(v int32) *ListLoadbalanceByTagResponseBindListenersBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadbalanceByTagResponseBindListenersBuilder) Build() *ListLoadbalanceByTagResponseBindListeners {
	return b.s
}


