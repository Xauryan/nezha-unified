// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceUDPListenerBodyGroupTypeEnum string

// List of GroupType
const (
    CreateLoadBalanceUDPListenerBodyGroupTypeEnumBlacklist CreateLoadBalanceUDPListenerBodyGroupTypeEnum = "blacklist"
    CreateLoadBalanceUDPListenerBodyGroupTypeEnumWhitelist CreateLoadBalanceUDPListenerBodyGroupTypeEnum = "whitelist"
)
type CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    CreateLoadBalanceUDPListenerBodyLbAlgorithmEnumRoundRobin CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum = "ROUND_ROBIN"
    CreateLoadBalanceUDPListenerBodyLbAlgorithmEnumLeastConnections CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    CreateLoadBalanceUDPListenerBodyLbAlgorithmEnumSourceIp CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum = "SOURCE_IP"
)
type CreateLoadBalanceUDPListenerBodyHealthTypeEnum string

// List of HealthType
const (
    CreateLoadBalanceUDPListenerBodyHealthTypeEnumPing CreateLoadBalanceUDPListenerBodyHealthTypeEnum = "PING"
    CreateLoadBalanceUDPListenerBodyHealthTypeEnumUdp CreateLoadBalanceUDPListenerBodyHealthTypeEnum = "UDP"
)
type CreateLoadBalanceUDPListenerBodyProtocolEnum string

// List of Protocol
const (
    CreateLoadBalanceUDPListenerBodyProtocolEnumUdp CreateLoadBalanceUDPListenerBodyProtocolEnum = "UDP"
)
type CreateLoadBalanceUDPListenerBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    CreateLoadBalanceUDPListenerBodySessionPersistenceEnumNone CreateLoadBalanceUDPListenerBodySessionPersistenceEnum = "NONE"
    CreateLoadBalanceUDPListenerBodySessionPersistenceEnumSourceIp CreateLoadBalanceUDPListenerBodySessionPersistenceEnum = "SOURCE_IP"
    CreateLoadBalanceUDPListenerBodySessionPersistenceEnumSourceIpPort CreateLoadBalanceUDPListenerBodySessionPersistenceEnum = "SOURCE_IP_PORT"
)
type CreateLoadBalanceUDPListenerBodyProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalanceUDPListenerBodyProductTypeEnumNormal CreateLoadBalanceUDPListenerBodyProductTypeEnum = "NORMAL"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumDeCloud CreateLoadBalanceUDPListenerBodyProductTypeEnum = "DE_CLOUD"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumAutoscaling CreateLoadBalanceUDPListenerBodyProductTypeEnum = "AUTOSCALING"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumVo CreateLoadBalanceUDPListenerBodyProductTypeEnum = "VO"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumCdn CreateLoadBalanceUDPListenerBodyProductTypeEnum = "CDN"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumPaasMaster CreateLoadBalanceUDPListenerBodyProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalanceUDPListenerBodyProductTypeEnumPaasSlave CreateLoadBalanceUDPListenerBodyProductTypeEnum = "PAAS_SLAVE"
)

type CreateLoadBalanceUDPListenerBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *CreateLoadBalanceUDPListenerBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *CreateLoadBalanceUDPListenerBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 监听器前端协议的端口号，取值为1-65535
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 监听器前后端使用的协议
	Protocol *CreateLoadBalanceUDPListenerBodyProtocolEnum `json:"protocol,omitempty"`
    // 监听器绑定的健康检查检查失败后的最大尝试检查次数，取值为1-10
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID，不为Null，说明使用该后端服务器组ID，不新建后端服务器组
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持方式：NONE,APP_COOKIE,HTTP_COOKIE,SOURCE_IP（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *CreateLoadBalanceUDPListenerBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *CreateLoadBalanceUDPListenerBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s CreateLoadBalanceUDPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceUDPListenerBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceUDPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceUDPListenerBody) SetHealthDelay(v int32) *CreateLoadBalanceUDPListenerBody {
	s.HealthDelay = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetGroupType(v CreateLoadBalanceUDPListenerBodyGroupTypeEnum) *CreateLoadBalanceUDPListenerBody {
	s.GroupType = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetKeepAliveTimeout(v int32) *CreateLoadBalanceUDPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetLbAlgorithm(v CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceUDPListenerBody {
	s.LbAlgorithm = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetNotes(v string) *CreateLoadBalanceUDPListenerBody {
	s.Notes = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetHealthType(v CreateLoadBalanceUDPListenerBodyHealthTypeEnum) *CreateLoadBalanceUDPListenerBody {
	s.HealthType = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetLoadBalanceId(v string) *CreateLoadBalanceUDPListenerBody {
	s.LoadBalanceId = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetProtocolPort(v int32) *CreateLoadBalanceUDPListenerBody {
	s.ProtocolPort = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetProtocol(v CreateLoadBalanceUDPListenerBodyProtocolEnum) *CreateLoadBalanceUDPListenerBody {
	s.Protocol = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetHealthMaxRetries(v int32) *CreateLoadBalanceUDPListenerBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetName(v string) *CreateLoadBalanceUDPListenerBody {
	s.Name = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetPoolId(v string) *CreateLoadBalanceUDPListenerBody {
	s.PoolId = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetSessionPersistence(v CreateLoadBalanceUDPListenerBodySessionPersistenceEnum) *CreateLoadBalanceUDPListenerBody {
	s.SessionPersistence = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetGroupEnabled(v bool) *CreateLoadBalanceUDPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetControlGroupId(v string) *CreateLoadBalanceUDPListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetHealthTimeout(v int32) *CreateLoadBalanceUDPListenerBody {
	s.HealthTimeout = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetProductType(v CreateLoadBalanceUDPListenerBodyProductTypeEnum) *CreateLoadBalanceUDPListenerBody {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalanceUDPListenerBody) SetPoolName(v string) *CreateLoadBalanceUDPListenerBody {
	s.PoolName = &v
	return s
}


type CreateLoadBalanceUDPListenerBodyBuilder struct {
	s *CreateLoadBalanceUDPListenerBody
}

func NewCreateLoadBalanceUDPListenerBodyBuilder() *CreateLoadBalanceUDPListenerBodyBuilder {
	s := &CreateLoadBalanceUDPListenerBody{}
	b := &CreateLoadBalanceUDPListenerBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) HealthDelay(v int32) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) GroupType(v CreateLoadBalanceUDPListenerBodyGroupTypeEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) KeepAliveTimeout(v int32) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) LbAlgorithm(v CreateLoadBalanceUDPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) Notes(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) HealthType(v CreateLoadBalanceUDPListenerBodyHealthTypeEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) LoadBalanceId(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) ProtocolPort(v int32) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) Protocol(v CreateLoadBalanceUDPListenerBodyProtocolEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) HealthMaxRetries(v int32) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) Name(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) PoolId(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) SessionPersistence(v CreateLoadBalanceUDPListenerBodySessionPersistenceEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) GroupEnabled(v bool) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) ControlGroupId(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) HealthTimeout(v int32) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) ProductType(v CreateLoadBalanceUDPListenerBodyProductTypeEnum) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) PoolName(v string) *CreateLoadBalanceUDPListenerBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *CreateLoadBalanceUDPListenerBodyBuilder) Build() *CreateLoadBalanceUDPListenerBody {
	return b.s
}


