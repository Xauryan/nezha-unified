// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalanceHTTPListenerBodyGroupTypeEnum string

// List of GroupType
const (
    CreateLoadBalanceHTTPListenerBodyGroupTypeEnumBlacklist CreateLoadBalanceHTTPListenerBodyGroupTypeEnum = "blacklist"
    CreateLoadBalanceHTTPListenerBodyGroupTypeEnumWhitelist CreateLoadBalanceHTTPListenerBodyGroupTypeEnum = "whitelist"
)
type CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnumRoundRobin CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum = "ROUND_ROBIN"
    CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnumLeastConnections CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnumSourceIp CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum = "SOURCE_IP"
)
type CreateLoadBalanceHTTPListenerBodyHealthTypeEnum string

// List of HealthType
const (
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumHttp CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "HTTP"
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumHttp10 CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "HTTP10"
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumHttp11 CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "HTTP11"
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumHttps CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "HTTPS"
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumTcp CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "TCP"
    CreateLoadBalanceHTTPListenerBodyHealthTypeEnumPing CreateLoadBalanceHTTPListenerBodyHealthTypeEnum = "PING"
)
type CreateLoadBalanceHTTPListenerBodyProtocolEnum string

// List of Protocol
const (
    CreateLoadBalanceHTTPListenerBodyProtocolEnumHttp CreateLoadBalanceHTTPListenerBodyProtocolEnum = "HTTP"
)
type CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    CreateLoadBalanceHTTPListenerBodySessionPersistenceEnumNone CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum = "NONE"
    CreateLoadBalanceHTTPListenerBodySessionPersistenceEnumAppCookie CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum = "APP_COOKIE"
    CreateLoadBalanceHTTPListenerBodySessionPersistenceEnumHttpCookie CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum = "HTTP_COOKIE"
    CreateLoadBalanceHTTPListenerBodySessionPersistenceEnumSourceIp CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum = "SOURCE_IP"
)
type CreateLoadBalanceHTTPListenerBodyProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumNormal CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "NORMAL"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumDeCloud CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "DE_CLOUD"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumAutoscaling CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "AUTOSCALING"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumVo CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "VO"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumCdn CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "CDN"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumPaasMaster CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalanceHTTPListenerBodyProductTypeEnumPaasSlave CreateLoadBalanceHTTPListenerBodyProductTypeEnum = "PAAS_SLAVE"
)

type CreateLoadBalanceHTTPListenerBody struct {
    position.Body
    // 监听器绑定的健康检查的时间间隔，单位为秒，取值为1-100的整数
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *CreateLoadBalanceHTTPListenerBodyGroupTypeEnum `json:"groupType,omitempty"`
    // 空闲超时时间，取值为10-300
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 重定向到的监听器ID，不是Null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 监听器绑定的健康检查的检查方式
	HealthType *CreateLoadBalanceHTTPListenerBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 监听器前端协议的端口号，取值为1-65535
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器绑定的检查方式为HTTP的健康检查的期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器前后端使用的协议
	Protocol *CreateLoadBalanceHTTPListenerBodyProtocolEnum `json:"protocol,omitempty"`
    // 监听器绑定的健康检查检查失败后的最大尝试检查次数，取值为1-10
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器名称为5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID, 不为Null，说明使用该后端服务器组ID，不新建后端服务器组
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持方式：NONE,APP_COOKIE,HTTP_COOKIE,SOURCE_IP（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 监听器绑定的检查方式为HTTP或HTTPS的用于监控的URL，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 监听器绑定的健康检查每次检查的超时时间，以秒为单位，取值为1-100的整数
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *CreateLoadBalanceHTTPListenerBodyProductTypeEnum `json:"productType,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s CreateLoadBalanceHTTPListenerBody) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalanceHTTPListenerBody) GoString() string {
	return s.String()
}

func (s CreateLoadBalanceHTTPListenerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthDelay(v int32) *CreateLoadBalanceHTTPListenerBody {
	s.HealthDelay = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetGroupType(v CreateLoadBalanceHTTPListenerBodyGroupTypeEnum) *CreateLoadBalanceHTTPListenerBody {
	s.GroupType = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetKeepAliveTimeout(v int32) *CreateLoadBalanceHTTPListenerBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetLbAlgorithm(v CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceHTTPListenerBody {
	s.LbAlgorithm = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetNotes(v string) *CreateLoadBalanceHTTPListenerBody {
	s.Notes = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthHttpMethod(v string) *CreateLoadBalanceHTTPListenerBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetRedirectToListenerId(v string) *CreateLoadBalanceHTTPListenerBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthType(v CreateLoadBalanceHTTPListenerBodyHealthTypeEnum) *CreateLoadBalanceHTTPListenerBody {
	s.HealthType = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetLoadBalanceId(v string) *CreateLoadBalanceHTTPListenerBody {
	s.LoadBalanceId = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetProtocolPort(v int32) *CreateLoadBalanceHTTPListenerBody {
	s.ProtocolPort = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetTransparent(v bool) *CreateLoadBalanceHTTPListenerBody {
	s.Transparent = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthExpectedCode(v string) *CreateLoadBalanceHTTPListenerBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetProtocol(v CreateLoadBalanceHTTPListenerBodyProtocolEnum) *CreateLoadBalanceHTTPListenerBody {
	s.Protocol = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthMaxRetries(v int32) *CreateLoadBalanceHTTPListenerBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetName(v string) *CreateLoadBalanceHTTPListenerBody {
	s.Name = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetPoolId(v string) *CreateLoadBalanceHTTPListenerBody {
	s.PoolId = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetSessionPersistence(v CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum) *CreateLoadBalanceHTTPListenerBody {
	s.SessionPersistence = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetGroupEnabled(v bool) *CreateLoadBalanceHTTPListenerBody {
	s.GroupEnabled = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthUrlPath(v string) *CreateLoadBalanceHTTPListenerBody {
	s.HealthUrlPath = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetControlGroupId(v string) *CreateLoadBalanceHTTPListenerBody {
	s.ControlGroupId = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetHealthTimeout(v int32) *CreateLoadBalanceHTTPListenerBody {
	s.HealthTimeout = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetProductType(v CreateLoadBalanceHTTPListenerBodyProductTypeEnum) *CreateLoadBalanceHTTPListenerBody {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalanceHTTPListenerBody) SetPoolName(v string) *CreateLoadBalanceHTTPListenerBody {
	s.PoolName = &v
	return s
}


type CreateLoadBalanceHTTPListenerBodyBuilder struct {
	s *CreateLoadBalanceHTTPListenerBody
}

func NewCreateLoadBalanceHTTPListenerBodyBuilder() *CreateLoadBalanceHTTPListenerBodyBuilder {
	s := &CreateLoadBalanceHTTPListenerBody{}
	b := &CreateLoadBalanceHTTPListenerBodyBuilder{s: s}
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthDelay(v int32) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) GroupType(v CreateLoadBalanceHTTPListenerBodyGroupTypeEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) KeepAliveTimeout(v int32) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) LbAlgorithm(v CreateLoadBalanceHTTPListenerBodyLbAlgorithmEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) Notes(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthHttpMethod(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) RedirectToListenerId(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthType(v CreateLoadBalanceHTTPListenerBodyHealthTypeEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) LoadBalanceId(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) ProtocolPort(v int32) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) Transparent(v bool) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthExpectedCode(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) Protocol(v CreateLoadBalanceHTTPListenerBodyProtocolEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthMaxRetries(v int32) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) Name(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) PoolId(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) SessionPersistence(v CreateLoadBalanceHTTPListenerBodySessionPersistenceEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) GroupEnabled(v bool) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthUrlPath(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) ControlGroupId(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) HealthTimeout(v int32) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) ProductType(v CreateLoadBalanceHTTPListenerBodyProductTypeEnum) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) PoolName(v string) *CreateLoadBalanceHTTPListenerBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *CreateLoadBalanceHTTPListenerBodyBuilder) Build() *CreateLoadBalanceHTTPListenerBody {
	return b.s
}


