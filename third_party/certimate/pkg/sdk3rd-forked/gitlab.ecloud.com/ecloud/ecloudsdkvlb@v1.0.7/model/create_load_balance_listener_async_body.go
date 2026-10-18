// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceListenerAsyncBodyGroupTypeEnum string

// List of GroupType
const (
    CreateLoadBalanceListenerAsyncBodyGroupTypeEnumBlacklist CreateLoadBalanceListenerAsyncBodyGroupTypeEnum = "blacklist"
    CreateLoadBalanceListenerAsyncBodyGroupTypeEnumWhitelist CreateLoadBalanceListenerAsyncBodyGroupTypeEnum = "whitelist"
)
type CreateLoadBalanceListenerAsyncBodyProtocolEnum string

// List of Protocol
const (
    CreateLoadBalanceListenerAsyncBodyProtocolEnumHttp CreateLoadBalanceListenerAsyncBodyProtocolEnum = "HTTP"
    CreateLoadBalanceListenerAsyncBodyProtocolEnumHttps CreateLoadBalanceListenerAsyncBodyProtocolEnum = "HTTPS"
    CreateLoadBalanceListenerAsyncBodyProtocolEnumTcp CreateLoadBalanceListenerAsyncBodyProtocolEnum = "TCP"
    CreateLoadBalanceListenerAsyncBodyProtocolEnumTerminatedHttps CreateLoadBalanceListenerAsyncBodyProtocolEnum = "TERMINATED_HTTPS"
    CreateLoadBalanceListenerAsyncBodyProtocolEnumUdp CreateLoadBalanceListenerAsyncBodyProtocolEnum = "UDP"
    CreateLoadBalanceListenerAsyncBodyProtocolEnumSip CreateLoadBalanceListenerAsyncBodyProtocolEnum = "SIP"
)
type CreateLoadBalanceListenerAsyncBodyProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumNormal CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "NORMAL"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumDeCloud CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "DE_CLOUD"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumAutoscaling CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "AUTOSCALING"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumVo CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "VO"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumCdn CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "CDN"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumPaasMaster CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalanceListenerAsyncBodyProductTypeEnumPaasSlave CreateLoadBalanceListenerAsyncBodyProductTypeEnum = "PAAS_SLAVE"
)
type CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnumRoundRobin CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum = "ROUND_ROBIN"
    CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnumLeastConnections CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnumSourceIp CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum = "SOURCE_IP"
)
type CreateLoadBalanceListenerAsyncBodyHealthTypeEnum string

// List of HealthType
const (
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumHttp CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "HTTP"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumHttps CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "HTTPS"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumTcp CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "TCP"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumPing CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "PING"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumUdp CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "UDP"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumHttp10 CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "HTTP10"
    CreateLoadBalanceListenerAsyncBodyHealthTypeEnumHttp11 CreateLoadBalanceListenerAsyncBodyHealthTypeEnum = "HTTP11"
)
type CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumNone CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "NONE"
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumAppCookie CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "APP_COOKIE"
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumHttpCookie CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "HTTP_COOKIE"
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumSourceIp CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "SOURCE_IP"
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumSourceIpPort CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "SOURCE_IP_PORT"
    CreateLoadBalanceListenerAsyncBodySessionPersistenceEnumCallId CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum = "CALL_ID"
)

type CreateLoadBalanceListenerAsyncBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *CreateLoadBalanceListenerAsyncBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 重定向到的监听器ID，不是Null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds []string `json:"sniContainerIds,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 是否是双AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 监听器前后端使用的协议
	Protocol *CreateLoadBalanceListenerAsyncBodyProtocolEnum `json:"protocol,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 监听器是否可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *CreateLoadBalanceListenerAsyncBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // TCP长连接超时时间，取值为10-3600
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于检查的方法，取值范围：GET，POST
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 监听器是否支持proxy_protocol（仅TCP协议监听器支持该参数，天池底层不支持该参数）
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *CreateLoadBalanceListenerAsyncBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 监听器前端协议的端口号，取值为1-65535
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 监听器绑定的检查方式为HTTP的健康检查的期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器绑定的健康检查检查失败后的最大尝试检查次数，取值为1-10
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID, 不为Null，说明使用该poolId，不新建pool
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持方式：NONE,APP_COOKIE,HTTP_COOKIE,SOURCE_IP,SOURCE_IP_PORT,CALL_ID（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于监控的URL，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 证书容器中的CA证书的ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 健康检查开关（仅provider为F5时可用）
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s CreateLoadBalanceListenerAsyncBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceListenerAsyncBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceListenerAsyncBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthDelay(v int32) *CreateLoadBalanceListenerAsyncBody {
	s.HealthDelay = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetGroupType(v CreateLoadBalanceListenerAsyncBodyGroupTypeEnum) *CreateLoadBalanceListenerAsyncBody {
	s.GroupType = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetRedirectToListenerId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetSniContainerIds(v []string) *CreateLoadBalanceListenerAsyncBody {
	s.SniContainerIds = v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetTransparent(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.Transparent = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetIsMultiAz(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.IsMultiAz = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetProtocol(v CreateLoadBalanceListenerAsyncBodyProtocolEnum) *CreateLoadBalanceListenerAsyncBody {
	s.Protocol = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHttp2(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.Http2 = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetAdminStateUp(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.AdminStateUp = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetDefaultTlsContainerId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetMutualAuthenticationUp(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetProductType(v CreateLoadBalanceListenerAsyncBodyProductTypeEnum) *CreateLoadBalanceListenerAsyncBody {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetPoolName(v string) *CreateLoadBalanceListenerAsyncBody {
	s.PoolName = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetTlsCustomizeProtocolId(v int64) *CreateLoadBalanceListenerAsyncBody {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetSniUp(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.SniUp = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetKeepAliveTimeout(v int32) *CreateLoadBalanceListenerAsyncBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetLbAlgorithm(v CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum) *CreateLoadBalanceListenerAsyncBody {
	s.LbAlgorithm = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetTlsUp(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.TlsUp = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthHttpMethod(v string) *CreateLoadBalanceListenerAsyncBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetProxyProtocol(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.ProxyProtocol = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthType(v CreateLoadBalanceListenerAsyncBodyHealthTypeEnum) *CreateLoadBalanceListenerAsyncBody {
	s.HealthType = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetLoadBalanceId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.LoadBalanceId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetProtocolPort(v int32) *CreateLoadBalanceListenerAsyncBody {
	s.ProtocolPort = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthExpectedCode(v string) *CreateLoadBalanceListenerAsyncBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthMaxRetries(v int32) *CreateLoadBalanceListenerAsyncBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetName(v string) *CreateLoadBalanceListenerAsyncBody {
	s.Name = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetPoolId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.PoolId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetSessionPersistence(v CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum) *CreateLoadBalanceListenerAsyncBody {
	s.SessionPersistence = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetGroupEnabled(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.GroupEnabled = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthUrlPath(v string) *CreateLoadBalanceListenerAsyncBody {
	s.HealthUrlPath = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetCaContainerId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.CaContainerId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthAdminStateUp(v bool) *CreateLoadBalanceListenerAsyncBody {
	s.HealthAdminStateUp = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetControlGroupId(v string) *CreateLoadBalanceListenerAsyncBody {
	s.ControlGroupId = &v
	return s
}

func (s *CreateLoadBalanceListenerAsyncBody) SetHealthTimeout(v int32) *CreateLoadBalanceListenerAsyncBody {
	s.HealthTimeout = &v
	return s
}


type CreateLoadBalanceListenerAsyncBodyBuilder struct {
	s *CreateLoadBalanceListenerAsyncBody
}

func NewCreateLoadBalanceListenerAsyncBodyBuilder() *CreateLoadBalanceListenerAsyncBodyBuilder {
	s := &CreateLoadBalanceListenerAsyncBody{}
	b := &CreateLoadBalanceListenerAsyncBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthDelay(v int32) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) GroupType(v CreateLoadBalanceListenerAsyncBodyGroupTypeEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) RedirectToListenerId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) SniContainerIds(v []string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) Transparent(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) IsMultiAz(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) Protocol(v CreateLoadBalanceListenerAsyncBodyProtocolEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) Http2(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) AdminStateUp(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) DefaultTlsContainerId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) MutualAuthenticationUp(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) ProductType(v CreateLoadBalanceListenerAsyncBodyProductTypeEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) PoolName(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) TlsCustomizeProtocolId(v int64) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) SniUp(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) KeepAliveTimeout(v int32) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) LbAlgorithm(v CreateLoadBalanceListenerAsyncBodyLbAlgorithmEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) TlsUp(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthHttpMethod(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) ProxyProtocol(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthType(v CreateLoadBalanceListenerAsyncBodyHealthTypeEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) LoadBalanceId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) ProtocolPort(v int32) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthExpectedCode(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthMaxRetries(v int32) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) Name(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) PoolId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) SessionPersistence(v CreateLoadBalanceListenerAsyncBodySessionPersistenceEnum) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) GroupEnabled(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthUrlPath(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) CaContainerId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthAdminStateUp(v bool) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) ControlGroupId(v string) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) HealthTimeout(v int32) *CreateLoadBalanceListenerAsyncBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *CreateLoadBalanceListenerAsyncBodyBuilder) Build() *CreateLoadBalanceListenerAsyncBody {
	return b.s
}


