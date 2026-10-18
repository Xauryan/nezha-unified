// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum string

// List of GroupType
const (
    CreateLoadBalanceHTTPSListenerBodyGroupTypeEnumBlacklist CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum = "blacklist"
    CreateLoadBalanceHTTPSListenerBodyGroupTypeEnumWhitelist CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum = "whitelist"
)
type CreateLoadBalanceHTTPSListenerBodyProtocolEnum string

// List of Protocol
const (
    CreateLoadBalanceHTTPSListenerBodyProtocolEnumHttps CreateLoadBalanceHTTPSListenerBodyProtocolEnum = "HTTPS"
    CreateLoadBalanceHTTPSListenerBodyProtocolEnumTerminatedHttps CreateLoadBalanceHTTPSListenerBodyProtocolEnum = "TERMINATED_HTTPS"
)
type CreateLoadBalanceHTTPSListenerBodyProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumNormal CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "NORMAL"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumDeCloud CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "DE_CLOUD"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumAutoscaling CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "AUTOSCALING"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumVo CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "VO"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumCdn CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "CDN"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumPaasMaster CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalanceHTTPSListenerBodyProductTypeEnumPaasSlave CreateLoadBalanceHTTPSListenerBodyProductTypeEnum = "PAAS_SLAVE"
)
type CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnumRoundRobin CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum = "ROUND_ROBIN"
    CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnumLeastConnections CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnumSourceIp CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum = "SOURCE_IP"
)
type CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum string

// List of HealthType
const (
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumHttp CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "HTTP"
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumHttp10 CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "HTTP10"
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumHttp11 CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "HTTP11"
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumHttps CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "HTTPS"
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumTcp CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "TCP"
    CreateLoadBalanceHTTPSListenerBodyHealthTypeEnumPing CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum = "PING"
)
type CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumNone CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "NONE"
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumAppCookie CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "APP_COOKIE"
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumHttpCookie CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "HTTP_COOKIE"
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumSourceIp CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "SOURCE_IP"
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumSourceIpPort CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "SOURCE_IP_PORT"
    CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnumCallId CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum = "CALL_ID"
)

type CreateLoadBalanceHTTPSListenerBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds []string `json:"sniContainerIds,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器前后端使用的协议：HTTPS（前端与后端协议均为 HTTPS），TERMINATED_HTTPS（前端为 HTTPS、后端为 HTTP）；NFV域监听器仅支持TERMINATED_HTTPS
	Protocol *CreateLoadBalanceHTTPSListenerBodyProtocolEnum `json:"protocol,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *CreateLoadBalanceHTTPSListenerBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum `json:"healthType,omitempty"`
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
    // 后端服务器组ID, 不为Null，说明使用该后端服务器组ID，不新建后端服务器组
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持方式：NONE,APP_COOKIE,HTTP_COOKIE,SOURCE_IP（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于监控的URL，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 证书容器中的CA证书的ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s CreateLoadBalanceHTTPSListenerBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPSListenerBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPSListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthDelay(v int32) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthDelay = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetGroupType(v CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.GroupType = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetNotes(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.Notes = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetSniContainerIds(v []string) *CreateLoadBalanceHTTPSListenerBody {
	s.SniContainerIds = v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetTransparent(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.Transparent = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetProtocol(v CreateLoadBalanceHTTPSListenerBodyProtocolEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.Protocol = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHttp2(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.Http2 = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetDefaultTlsContainerId(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetMutualAuthenticationUp(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetProductType(v CreateLoadBalanceHTTPSListenerBodyProductTypeEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetPoolName(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.PoolName = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetTlsCustomizeProtocolId(v int64) *CreateLoadBalanceHTTPSListenerBody {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetSniUp(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.SniUp = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetKeepAliveTimeout(v int32) *CreateLoadBalanceHTTPSListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetLbAlgorithm(v CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.LbAlgorithm = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetTlsUp(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.TlsUp = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthHttpMethod(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthType(v CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthType = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetLoadBalanceId(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.LoadBalanceId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetProtocolPort(v int32) *CreateLoadBalanceHTTPSListenerBody {
	s.ProtocolPort = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthExpectedCode(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthMaxRetries(v int32) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetName(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.Name = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetPoolId(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.PoolId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetSessionPersistence(v CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum) *CreateLoadBalanceHTTPSListenerBody {
	s.SessionPersistence = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetGroupEnabled(v bool) *CreateLoadBalanceHTTPSListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthUrlPath(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthUrlPath = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetCaContainerId(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.CaContainerId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetControlGroupId(v string) *CreateLoadBalanceHTTPSListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *CreateLoadBalanceHTTPSListenerBody) SetHealthTimeout(v int32) *CreateLoadBalanceHTTPSListenerBody {
	s.HealthTimeout = &v
	return s
}


type CreateLoadBalanceHTTPSListenerBodyBuilder struct {
	s *CreateLoadBalanceHTTPSListenerBody
}

func NewCreateLoadBalanceHTTPSListenerBodyBuilder() *CreateLoadBalanceHTTPSListenerBodyBuilder {
	s := &CreateLoadBalanceHTTPSListenerBody{}
	b := &CreateLoadBalanceHTTPSListenerBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthDelay(v int32) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) GroupType(v CreateLoadBalanceHTTPSListenerBodyGroupTypeEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Notes(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) SniContainerIds(v []string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.SniContainerIds = v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Transparent(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Protocol(v CreateLoadBalanceHTTPSListenerBodyProtocolEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Http2(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) DefaultTlsContainerId(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) MutualAuthenticationUp(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) ProductType(v CreateLoadBalanceHTTPSListenerBodyProductTypeEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) PoolName(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) TlsCustomizeProtocolId(v int64) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) SniUp(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) KeepAliveTimeout(v int32) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) LbAlgorithm(v CreateLoadBalanceHTTPSListenerBodyLbAlgorithmEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) TlsUp(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthHttpMethod(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthType(v CreateLoadBalanceHTTPSListenerBodyHealthTypeEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) LoadBalanceId(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) ProtocolPort(v int32) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthExpectedCode(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthMaxRetries(v int32) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Name(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) PoolId(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) SessionPersistence(v CreateLoadBalanceHTTPSListenerBodySessionPersistenceEnum) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) GroupEnabled(v bool) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthUrlPath(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) CaContainerId(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) ControlGroupId(v string) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) HealthTimeout(v int32) *CreateLoadBalanceHTTPSListenerBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *CreateLoadBalanceHTTPSListenerBodyBuilder) Build() *CreateLoadBalanceHTTPSListenerBody {
	return b.s
}


