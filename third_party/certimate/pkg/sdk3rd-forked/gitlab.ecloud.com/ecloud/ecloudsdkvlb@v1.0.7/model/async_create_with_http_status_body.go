// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreateWithHttpStatusBodyGroupTypeEnum string

// List of GroupType
const (
    AsyncCreateWithHttpStatusBodyGroupTypeEnumBlacklist AsyncCreateWithHttpStatusBodyGroupTypeEnum = "blacklist"
    AsyncCreateWithHttpStatusBodyGroupTypeEnumWhitelist AsyncCreateWithHttpStatusBodyGroupTypeEnum = "whitelist"
)
type AsyncCreateWithHttpStatusBodyProtocolEnum string

// List of Protocol
const (
    AsyncCreateWithHttpStatusBodyProtocolEnumHttp AsyncCreateWithHttpStatusBodyProtocolEnum = "HTTP"
    AsyncCreateWithHttpStatusBodyProtocolEnumHttps AsyncCreateWithHttpStatusBodyProtocolEnum = "HTTPS"
    AsyncCreateWithHttpStatusBodyProtocolEnumTcp AsyncCreateWithHttpStatusBodyProtocolEnum = "TCP"
    AsyncCreateWithHttpStatusBodyProtocolEnumTerminatedHttps AsyncCreateWithHttpStatusBodyProtocolEnum = "TERMINATED_HTTPS"
    AsyncCreateWithHttpStatusBodyProtocolEnumUdp AsyncCreateWithHttpStatusBodyProtocolEnum = "UDP"
)
type AsyncCreateWithHttpStatusBodyProductTypeEnum string

// List of ProductType
const (
    AsyncCreateWithHttpStatusBodyProductTypeEnumNormal AsyncCreateWithHttpStatusBodyProductTypeEnum = "NORMAL"
    AsyncCreateWithHttpStatusBodyProductTypeEnumDeCloud AsyncCreateWithHttpStatusBodyProductTypeEnum = "DE_CLOUD"
    AsyncCreateWithHttpStatusBodyProductTypeEnumAutoscaling AsyncCreateWithHttpStatusBodyProductTypeEnum = "AUTOSCALING"
    AsyncCreateWithHttpStatusBodyProductTypeEnumVo AsyncCreateWithHttpStatusBodyProductTypeEnum = "VO"
    AsyncCreateWithHttpStatusBodyProductTypeEnumCdn AsyncCreateWithHttpStatusBodyProductTypeEnum = "CDN"
    AsyncCreateWithHttpStatusBodyProductTypeEnumPaasMaster AsyncCreateWithHttpStatusBodyProductTypeEnum = "PAAS_MASTER"
    AsyncCreateWithHttpStatusBodyProductTypeEnumPaasSlave AsyncCreateWithHttpStatusBodyProductTypeEnum = "PAAS_SLAVE"
)
type AsyncCreateWithHttpStatusBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    AsyncCreateWithHttpStatusBodyLbAlgorithmEnumRoundRobin AsyncCreateWithHttpStatusBodyLbAlgorithmEnum = "ROUND_ROBIN"
    AsyncCreateWithHttpStatusBodyLbAlgorithmEnumLeastConnections AsyncCreateWithHttpStatusBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    AsyncCreateWithHttpStatusBodyLbAlgorithmEnumSourceIp AsyncCreateWithHttpStatusBodyLbAlgorithmEnum = "SOURCE_IP"
)
type AsyncCreateWithHttpStatusBodyHealthTypeEnum string

// List of HealthType
const (
    AsyncCreateWithHttpStatusBodyHealthTypeEnumHttp10 AsyncCreateWithHttpStatusBodyHealthTypeEnum = "HTTP10"
    AsyncCreateWithHttpStatusBodyHealthTypeEnumHttp11 AsyncCreateWithHttpStatusBodyHealthTypeEnum = "HTTP11"
    AsyncCreateWithHttpStatusBodyHealthTypeEnumHttps AsyncCreateWithHttpStatusBodyHealthTypeEnum = "HTTPS"
    AsyncCreateWithHttpStatusBodyHealthTypeEnumTcp AsyncCreateWithHttpStatusBodyHealthTypeEnum = "TCP"
    AsyncCreateWithHttpStatusBodyHealthTypeEnumPing AsyncCreateWithHttpStatusBodyHealthTypeEnum = "PING"
)
type AsyncCreateWithHttpStatusBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    AsyncCreateWithHttpStatusBodySessionPersistenceEnumNone AsyncCreateWithHttpStatusBodySessionPersistenceEnum = "NONE"
    AsyncCreateWithHttpStatusBodySessionPersistenceEnumAppCookie AsyncCreateWithHttpStatusBodySessionPersistenceEnum = "APP_COOKIE"
    AsyncCreateWithHttpStatusBodySessionPersistenceEnumHttpCookie AsyncCreateWithHttpStatusBodySessionPersistenceEnum = "HTTP_COOKIE"
    AsyncCreateWithHttpStatusBodySessionPersistenceEnumSourceIp AsyncCreateWithHttpStatusBodySessionPersistenceEnum = "SOURCE_IP"
    AsyncCreateWithHttpStatusBodySessionPersistenceEnumCallId AsyncCreateWithHttpStatusBodySessionPersistenceEnum = "CALL_ID"
)

type AsyncCreateWithHttpStatusBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *AsyncCreateWithHttpStatusBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 重定向到的监听器ID，不是Null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds []string `json:"sniContainerIds,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 监听器前后端使用的协议
	Protocol *AsyncCreateWithHttpStatusBodyProtocolEnum `json:"protocol,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 监听器是否可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *AsyncCreateWithHttpStatusBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // TCP长连接超时时间，取值为10-3600
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法
	LbAlgorithm *AsyncCreateWithHttpStatusBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 监听器是否支持proxy_protocol（仅TCP协议监听器支持该参数，天池底层不支持该参数）
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *AsyncCreateWithHttpStatusBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡ID
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
	SessionPersistence *AsyncCreateWithHttpStatusBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于监控的URL，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 证书容器中的CA证书ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 健康检查开关（仅provider为F5时可用）
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 访问控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s AsyncCreateWithHttpStatusBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreateWithHttpStatusBody) GoString() string {
	return s.String()
}

func (s AsyncCreateWithHttpStatusBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthDelay(v int32) *AsyncCreateWithHttpStatusBody {
	s.HealthDelay = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetGroupType(v AsyncCreateWithHttpStatusBodyGroupTypeEnum) *AsyncCreateWithHttpStatusBody {
	s.GroupType = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetRedirectToListenerId(v string) *AsyncCreateWithHttpStatusBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetSniContainerIds(v []string) *AsyncCreateWithHttpStatusBody {
	s.SniContainerIds = v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetTransparent(v bool) *AsyncCreateWithHttpStatusBody {
	s.Transparent = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetIsMultiAz(v bool) *AsyncCreateWithHttpStatusBody {
	s.IsMultiAz = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetProtocol(v AsyncCreateWithHttpStatusBodyProtocolEnum) *AsyncCreateWithHttpStatusBody {
	s.Protocol = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHttp2(v bool) *AsyncCreateWithHttpStatusBody {
	s.Http2 = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetAdminStateUp(v bool) *AsyncCreateWithHttpStatusBody {
	s.AdminStateUp = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetDefaultTlsContainerId(v string) *AsyncCreateWithHttpStatusBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetMutualAuthenticationUp(v bool) *AsyncCreateWithHttpStatusBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetProductType(v AsyncCreateWithHttpStatusBodyProductTypeEnum) *AsyncCreateWithHttpStatusBody {
	s.ProductType = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetPoolName(v string) *AsyncCreateWithHttpStatusBody {
	s.PoolName = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetTlsCustomizeProtocolId(v int64) *AsyncCreateWithHttpStatusBody {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetSniUp(v bool) *AsyncCreateWithHttpStatusBody {
	s.SniUp = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetKeepAliveTimeout(v int32) *AsyncCreateWithHttpStatusBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetLbAlgorithm(v AsyncCreateWithHttpStatusBodyLbAlgorithmEnum) *AsyncCreateWithHttpStatusBody {
	s.LbAlgorithm = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetTlsUp(v bool) *AsyncCreateWithHttpStatusBody {
	s.TlsUp = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthHttpMethod(v string) *AsyncCreateWithHttpStatusBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetProxyProtocol(v bool) *AsyncCreateWithHttpStatusBody {
	s.ProxyProtocol = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthType(v AsyncCreateWithHttpStatusBodyHealthTypeEnum) *AsyncCreateWithHttpStatusBody {
	s.HealthType = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetLoadBalanceId(v string) *AsyncCreateWithHttpStatusBody {
	s.LoadBalanceId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetProtocolPort(v int32) *AsyncCreateWithHttpStatusBody {
	s.ProtocolPort = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthExpectedCode(v string) *AsyncCreateWithHttpStatusBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthMaxRetries(v int32) *AsyncCreateWithHttpStatusBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetName(v string) *AsyncCreateWithHttpStatusBody {
	s.Name = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetPoolId(v string) *AsyncCreateWithHttpStatusBody {
	s.PoolId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetSessionPersistence(v AsyncCreateWithHttpStatusBodySessionPersistenceEnum) *AsyncCreateWithHttpStatusBody {
	s.SessionPersistence = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetGroupEnabled(v bool) *AsyncCreateWithHttpStatusBody {
	s.GroupEnabled = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthUrlPath(v string) *AsyncCreateWithHttpStatusBody {
	s.HealthUrlPath = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetCaContainerId(v string) *AsyncCreateWithHttpStatusBody {
	s.CaContainerId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthAdminStateUp(v bool) *AsyncCreateWithHttpStatusBody {
	s.HealthAdminStateUp = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetControlGroupId(v string) *AsyncCreateWithHttpStatusBody {
	s.ControlGroupId = &v
	return s
}

func (s *AsyncCreateWithHttpStatusBody) SetHealthTimeout(v int32) *AsyncCreateWithHttpStatusBody {
	s.HealthTimeout = &v
	return s
}


type AsyncCreateWithHttpStatusBodyBuilder struct {
	s *AsyncCreateWithHttpStatusBody
}

func NewAsyncCreateWithHttpStatusBodyBuilder() *AsyncCreateWithHttpStatusBodyBuilder {
	s := &AsyncCreateWithHttpStatusBody{}
	b := &AsyncCreateWithHttpStatusBodyBuilder{s: s}
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthDelay(v int32) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) GroupType(v AsyncCreateWithHttpStatusBodyGroupTypeEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) RedirectToListenerId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) SniContainerIds(v []string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) Transparent(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) IsMultiAz(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) Protocol(v AsyncCreateWithHttpStatusBodyProtocolEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) Http2(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) AdminStateUp(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) DefaultTlsContainerId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) MutualAuthenticationUp(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) ProductType(v AsyncCreateWithHttpStatusBodyProductTypeEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) PoolName(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) TlsCustomizeProtocolId(v int64) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) SniUp(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) KeepAliveTimeout(v int32) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) LbAlgorithm(v AsyncCreateWithHttpStatusBodyLbAlgorithmEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) TlsUp(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthHttpMethod(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) ProxyProtocol(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthType(v AsyncCreateWithHttpStatusBodyHealthTypeEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) LoadBalanceId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) ProtocolPort(v int32) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthExpectedCode(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthMaxRetries(v int32) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) Name(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) PoolId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) SessionPersistence(v AsyncCreateWithHttpStatusBodySessionPersistenceEnum) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) GroupEnabled(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthUrlPath(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) CaContainerId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthAdminStateUp(v bool) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) ControlGroupId(v string) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) HealthTimeout(v int32) *AsyncCreateWithHttpStatusBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *AsyncCreateWithHttpStatusBodyBuilder) Build() *AsyncCreateWithHttpStatusBody {
	return b.s
}


