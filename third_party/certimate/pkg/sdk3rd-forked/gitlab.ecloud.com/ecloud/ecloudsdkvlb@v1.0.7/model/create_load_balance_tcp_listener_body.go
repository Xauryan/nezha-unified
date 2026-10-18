// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceTCPListenerBodyGroupTypeEnum string

// List of GroupType
const (
    CreateLoadBalanceTCPListenerBodyGroupTypeEnumBlacklist CreateLoadBalanceTCPListenerBodyGroupTypeEnum = "blacklist"
    CreateLoadBalanceTCPListenerBodyGroupTypeEnumWhitelist CreateLoadBalanceTCPListenerBodyGroupTypeEnum = "whitelist"
)
type CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    CreateLoadBalanceTCPListenerBodyLbAlgorithmEnumRoundRobin CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum = "ROUND_ROBIN"
    CreateLoadBalanceTCPListenerBodyLbAlgorithmEnumLeastConnections CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    CreateLoadBalanceTCPListenerBodyLbAlgorithmEnumSourceIp CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum = "SOURCE_IP"
)
type CreateLoadBalanceTCPListenerBodyHealthTypeEnum string

// List of HealthType
const (
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumHttp CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "HTTP"
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumHttp10 CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "HTTP10"
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumHttp11 CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "HTTP11"
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumHttps CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "HTTPS"
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumTcp CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "TCP"
    CreateLoadBalanceTCPListenerBodyHealthTypeEnumPing CreateLoadBalanceTCPListenerBodyHealthTypeEnum = "PING"
)
type CreateLoadBalanceTCPListenerBodyProtocolEnum string

// List of Protocol
const (
    CreateLoadBalanceTCPListenerBodyProtocolEnumTcp CreateLoadBalanceTCPListenerBodyProtocolEnum = "TCP"
)
type CreateLoadBalanceTCPListenerBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    CreateLoadBalanceTCPListenerBodySessionPersistenceEnumNone CreateLoadBalanceTCPListenerBodySessionPersistenceEnum = "NONE"
    CreateLoadBalanceTCPListenerBodySessionPersistenceEnumSourceIp CreateLoadBalanceTCPListenerBodySessionPersistenceEnum = "SOURCE_IP"
    CreateLoadBalanceTCPListenerBodySessionPersistenceEnumSourceIpPort CreateLoadBalanceTCPListenerBodySessionPersistenceEnum = "SOURCE_IP_PORT"
)
type CreateLoadBalanceTCPListenerBodyProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalanceTCPListenerBodyProductTypeEnumNormal CreateLoadBalanceTCPListenerBodyProductTypeEnum = "NORMAL"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumDeCloud CreateLoadBalanceTCPListenerBodyProductTypeEnum = "DE_CLOUD"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumAutoscaling CreateLoadBalanceTCPListenerBodyProductTypeEnum = "AUTOSCALING"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumVo CreateLoadBalanceTCPListenerBodyProductTypeEnum = "VO"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumCdn CreateLoadBalanceTCPListenerBodyProductTypeEnum = "CDN"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumPaasMaster CreateLoadBalanceTCPListenerBodyProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalanceTCPListenerBodyProductTypeEnumPaasSlave CreateLoadBalanceTCPListenerBodyProductTypeEnum = "PAAS_SLAVE"
)

type CreateLoadBalanceTCPListenerBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *CreateLoadBalanceTCPListenerBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 监听器是否支持proxy_protocol（仅TCP协议监听器支持该参数，天池底层不支持该参数）
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *CreateLoadBalanceTCPListenerBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 监听器前端协议的端口号，取值为1-65535
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器绑定的检查方式为HTTP的健康检查的期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器前后端使用的协议
	Protocol *CreateLoadBalanceTCPListenerBodyProtocolEnum `json:"protocol,omitempty"`
    // 监听器绑定的健康检查检查失败后的最大尝试检查次数，取值为1-10
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID，不为Null，说明使用该后端服务器组ID，不新建后端服务器组
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持方式：NONE,APP_COOKIE,HTTP_COOKIE,SOURCE_IP（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *CreateLoadBalanceTCPListenerBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于监控的URL，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *CreateLoadBalanceTCPListenerBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s CreateLoadBalanceTCPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceTCPListenerBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceTCPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthDelay(v int32) *CreateLoadBalanceTCPListenerBody {
	s.HealthDelay = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetGroupType(v CreateLoadBalanceTCPListenerBodyGroupTypeEnum) *CreateLoadBalanceTCPListenerBody {
	s.GroupType = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetKeepAliveTimeout(v int32) *CreateLoadBalanceTCPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetLbAlgorithm(v CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceTCPListenerBody {
	s.LbAlgorithm = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetNotes(v string) *CreateLoadBalanceTCPListenerBody {
	s.Notes = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthHttpMethod(v string) *CreateLoadBalanceTCPListenerBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetProxyProtocol(v bool) *CreateLoadBalanceTCPListenerBody {
	s.ProxyProtocol = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthType(v CreateLoadBalanceTCPListenerBodyHealthTypeEnum) *CreateLoadBalanceTCPListenerBody {
	s.HealthType = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetLoadBalanceId(v string) *CreateLoadBalanceTCPListenerBody {
	s.LoadBalanceId = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetProtocolPort(v int32) *CreateLoadBalanceTCPListenerBody {
	s.ProtocolPort = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetTransparent(v bool) *CreateLoadBalanceTCPListenerBody {
	s.Transparent = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthExpectedCode(v string) *CreateLoadBalanceTCPListenerBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetProtocol(v CreateLoadBalanceTCPListenerBodyProtocolEnum) *CreateLoadBalanceTCPListenerBody {
	s.Protocol = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthMaxRetries(v int32) *CreateLoadBalanceTCPListenerBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetName(v string) *CreateLoadBalanceTCPListenerBody {
	s.Name = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetPoolId(v string) *CreateLoadBalanceTCPListenerBody {
	s.PoolId = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetSessionPersistence(v CreateLoadBalanceTCPListenerBodySessionPersistenceEnum) *CreateLoadBalanceTCPListenerBody {
	s.SessionPersistence = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetGroupEnabled(v bool) *CreateLoadBalanceTCPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthUrlPath(v string) *CreateLoadBalanceTCPListenerBody {
	s.HealthUrlPath = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetControlGroupId(v string) *CreateLoadBalanceTCPListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetHealthTimeout(v int32) *CreateLoadBalanceTCPListenerBody {
	s.HealthTimeout = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetProductType(v CreateLoadBalanceTCPListenerBodyProductTypeEnum) *CreateLoadBalanceTCPListenerBody {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalanceTCPListenerBody) SetPoolName(v string) *CreateLoadBalanceTCPListenerBody {
	s.PoolName = &v
	return s
}


type CreateLoadBalanceTCPListenerBodyBuilder struct {
	s *CreateLoadBalanceTCPListenerBody
}

func NewCreateLoadBalanceTCPListenerBodyBuilder() *CreateLoadBalanceTCPListenerBodyBuilder {
	s := &CreateLoadBalanceTCPListenerBody{}
	b := &CreateLoadBalanceTCPListenerBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthDelay(v int32) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) GroupType(v CreateLoadBalanceTCPListenerBodyGroupTypeEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) KeepAliveTimeout(v int32) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) LbAlgorithm(v CreateLoadBalanceTCPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) Notes(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthHttpMethod(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) ProxyProtocol(v bool) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthType(v CreateLoadBalanceTCPListenerBodyHealthTypeEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) LoadBalanceId(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) ProtocolPort(v int32) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) Transparent(v bool) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthExpectedCode(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) Protocol(v CreateLoadBalanceTCPListenerBodyProtocolEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthMaxRetries(v int32) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) Name(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) PoolId(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) SessionPersistence(v CreateLoadBalanceTCPListenerBodySessionPersistenceEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) GroupEnabled(v bool) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthUrlPath(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) ControlGroupId(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) HealthTimeout(v int32) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) ProductType(v CreateLoadBalanceTCPListenerBodyProductTypeEnum) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) PoolName(v string) *CreateLoadBalanceTCPListenerBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *CreateLoadBalanceTCPListenerBodyBuilder) Build() *CreateLoadBalanceTCPListenerBody {
	return b.s
}


