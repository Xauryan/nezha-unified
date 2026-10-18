// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceListenerRespResponseContentProtocolEnum string

// List of Protocol
const (
    ListLoadBalanceListenerRespResponseContentProtocolEnumHttp ListLoadBalanceListenerRespResponseContentProtocolEnum = "HTTP"
    ListLoadBalanceListenerRespResponseContentProtocolEnumHttps ListLoadBalanceListenerRespResponseContentProtocolEnum = "HTTPS"
    ListLoadBalanceListenerRespResponseContentProtocolEnumTcp ListLoadBalanceListenerRespResponseContentProtocolEnum = "TCP"
    ListLoadBalanceListenerRespResponseContentProtocolEnumUdp ListLoadBalanceListenerRespResponseContentProtocolEnum = "UDP"
    ListLoadBalanceListenerRespResponseContentProtocolEnumTerminatedHttps ListLoadBalanceListenerRespResponseContentProtocolEnum = "TERMINATED_HTTPS"
    ListLoadBalanceListenerRespResponseContentProtocolEnumSip ListLoadBalanceListenerRespResponseContentProtocolEnum = "SIP"
)
type ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumActive ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "ACTIVE"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumBuild ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "BUILD"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumDown ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "DOWN"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumError ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "ERROR"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumPendingCreate ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "PENDING_CREATE"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumPendingUpdate ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "PENDING_UPDATE"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumPendingDelete ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "PENDING_DELETE"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumMigrating ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "MIGRATING"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumDeleted ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "DELETED"
    ListLoadBalanceListenerRespResponseContentProvisioningStatusEnumUnrecognized ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum = "UNRECOGNIZED"
)
type ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum string

// List of LbAlgorithm
const (
    ListLoadBalanceListenerRespResponseContentLbAlgorithmEnumRoundRobin ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum = "ROUND_ROBIN"
    ListLoadBalanceListenerRespResponseContentLbAlgorithmEnumLeastConnections ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum = "LEAST_CONNECTIONS"
    ListLoadBalanceListenerRespResponseContentLbAlgorithmEnumSourceIp ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum = "SOURCE_IP"
)
type ListLoadBalanceListenerRespResponseContentHealthTypeEnum string

// List of HealthType
const (
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumHttp ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "HTTP"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumHttps ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "HTTPS"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumTcp ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "TCP"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumPing ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "PING"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumUdp ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "UDP"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumHttp10 ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "HTTP10"
    ListLoadBalanceListenerRespResponseContentHealthTypeEnumHttp11 ListLoadBalanceListenerRespResponseContentHealthTypeEnum = "HTTP11"
)
type ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum string

// List of SessionPersistence
const (
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumNone ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "NONE"
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumAppCookie ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "APP_COOKIE"
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumHttpCookie ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "HTTP_COOKIE"
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumSourceIp ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "SOURCE_IP"
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumSourceIpPort ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "SOURCE_IP_PORT"
    ListLoadBalanceListenerRespResponseContentSessionPersistenceEnumCallId ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum = "CALL_ID"
)
type ListLoadBalanceListenerRespResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadBalanceListenerRespResponseContentOpStatusEnumInactive ListLoadBalanceListenerRespResponseContentOpStatusEnum = "INACTIVE"
    ListLoadBalanceListenerRespResponseContentOpStatusEnumActive ListLoadBalanceListenerRespResponseContentOpStatusEnum = "ACTIVE"
    ListLoadBalanceListenerRespResponseContentOpStatusEnumCreating ListLoadBalanceListenerRespResponseContentOpStatusEnum = "CREATING"
    ListLoadBalanceListenerRespResponseContentOpStatusEnumDeleting ListLoadBalanceListenerRespResponseContentOpStatusEnum = "DELETING"
    ListLoadBalanceListenerRespResponseContentOpStatusEnumUpdating ListLoadBalanceListenerRespResponseContentOpStatusEnum = "UPDATING"
    ListLoadBalanceListenerRespResponseContentOpStatusEnumError ListLoadBalanceListenerRespResponseContentOpStatusEnum = "ERROR"
)

type ListLoadBalanceListenerRespResponseContent struct {

    // 健康监控时间间隔(单位：秒)
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 绑定的访问控制组类型: blacklist:黑名单，whitelist:白名单
	GroupType *string `json:"groupType,omitempty"`
    // 重定向到的监听器ID，不是null，说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds *string `json:"sniContainerIds,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 是否是双AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 重定向到的监听器名称
	RedirectToListenerName *string `json:"redirectToListenerName,omitempty"`
    // 监听器协议
	Protocol *ListLoadBalanceListenerRespResponseContentProtocolEnum `json:"protocol,omitempty"`
    // 证书容器中的SNI证书ID列表
	SniContainerIdList []string `json:"sniContainerIdList,omitempty"`
    // 监听器底层状态： ACTIVE（资源已激活），BUILD（资源已创建），DOWN（资源已创建但未激活），ERROR（资源无法工作），PENDING_CREATE,PENDING_UPDATE（资源正在创建当中），PENDING_DELETE（资源在删除当中），MIGRATING（迁移中），DELETED（资源已删除），UNRECOGNIZED（未知状态）
	ProvisioningStatus *ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // 是否开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
    // TLS安全策略的ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // TCP长连接超时时间
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法）
	LbAlgorithm *ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 健康检查的资源ID
	HealthId *string `json:"healthId,omitempty"`
    // 健康监控类型
	HealthType *ListLoadBalanceListenerRespResponseContentHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡规格
	LoadBalanceFlavor *int32 `json:"loadBalanceFlavor,omitempty"`
    // 负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 协议端口号
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 访问控制组名称
	GroupName *string `json:"groupName,omitempty"`
    // 连接数限制
	ConnectionLimit *int32 `json:"connectionLimit,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // TLS安全策略的名称
	TlsCustomizeProtocolName *string `json:"tlsCustomizeProtocolName,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持类型
	SessionPersistence *ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 用于监控的url，必须以“/”开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 证书容器中的CA证书的ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 健康检查开关
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 监听器状态
	OpStatus *ListLoadBalanceListenerRespResponseContentOpStatusEnum `json:"opStatus,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 最大等待时间
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 双AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
}

func (s ListLoadBalanceListenerRespResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceListenerRespResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadBalanceListenerRespResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthDelay(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.HealthDelay = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetModifiedTime(v string) *ListLoadBalanceListenerRespResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetGroupType(v string) *ListLoadBalanceListenerRespResponseContent {
	s.GroupType = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetRedirectToListenerId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.RedirectToListenerId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetSniContainerIds(v string) *ListLoadBalanceListenerRespResponseContent {
	s.SniContainerIds = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetDescription(v string) *ListLoadBalanceListenerRespResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetTransparent(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.Transparent = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetIsMultiAz(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetRedirectToListenerName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.RedirectToListenerName = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetProtocol(v ListLoadBalanceListenerRespResponseContentProtocolEnum) *ListLoadBalanceListenerRespResponseContent {
	s.Protocol = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetSniContainerIdList(v []string) *ListLoadBalanceListenerRespResponseContent {
	s.SniContainerIdList = v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetProvisioningStatus(v ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum) *ListLoadBalanceListenerRespResponseContent {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetCreatedTime(v string) *ListLoadBalanceListenerRespResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHttp2(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.Http2 = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetDefaultTlsContainerId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetMutualAuthenticationUp(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetCookieName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.CookieName = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetPoolName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.PoolName = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetTlsCustomizeProtocolId(v int64) *ListLoadBalanceListenerRespResponseContent {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetSniUp(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.SniUp = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetKeepAliveTimeout(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.KeepAliveTimeout = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetLbAlgorithm(v ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum) *ListLoadBalanceListenerRespResponseContent {
	s.LbAlgorithm = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetTlsUp(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.TlsUp = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthHttpMethod(v string) *ListLoadBalanceListenerRespResponseContent {
	s.HealthHttpMethod = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetProxyProtocol(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.ProxyProtocol = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.HealthId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthType(v ListLoadBalanceListenerRespResponseContentHealthTypeEnum) *ListLoadBalanceListenerRespResponseContent {
	s.HealthType = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetLoadBalanceFlavor(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.LoadBalanceFlavor = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetLoadBalanceId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetProtocolPort(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.ProtocolPort = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthExpectedCode(v string) *ListLoadBalanceListenerRespResponseContent {
	s.HealthExpectedCode = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetGroupName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.GroupName = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetConnectionLimit(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.ConnectionLimit = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetDeleted(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthMaxRetries(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.HealthMaxRetries = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetTlsCustomizeProtocolName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.TlsCustomizeProtocolName = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetName(v string) *ListLoadBalanceListenerRespResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetPoolId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetSessionPersistence(v ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum) *ListLoadBalanceListenerRespResponseContent {
	s.SessionPersistence = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetGroupEnabled(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.GroupEnabled = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthUrlPath(v string) *ListLoadBalanceListenerRespResponseContent {
	s.HealthUrlPath = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetCaContainerId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.CaContainerId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthAdminStateUp(v bool) *ListLoadBalanceListenerRespResponseContent {
	s.HealthAdminStateUp = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetOpStatus(v ListLoadBalanceListenerRespResponseContentOpStatusEnum) *ListLoadBalanceListenerRespResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetControlGroupId(v string) *ListLoadBalanceListenerRespResponseContent {
	s.ControlGroupId = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetHealthTimeout(v int32) *ListLoadBalanceListenerRespResponseContent {
	s.HealthTimeout = &v
	return s
}

func (s *ListLoadBalanceListenerRespResponseContent) SetMultiAzUuid(v string) *ListLoadBalanceListenerRespResponseContent {
	s.MultiAzUuid = &v
	return s
}


type ListLoadBalanceListenerRespResponseContentBuilder struct {
	s *ListLoadBalanceListenerRespResponseContent
}

func NewListLoadBalanceListenerRespResponseContentBuilder() *ListLoadBalanceListenerRespResponseContentBuilder {
	s := &ListLoadBalanceListenerRespResponseContent{}
	b := &ListLoadBalanceListenerRespResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthDelay(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ModifiedTime(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) GroupType(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) RedirectToListenerId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) SniContainerIds(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.SniContainerIds = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Description(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Transparent(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) IsMultiAz(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) RedirectToListenerName(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.RedirectToListenerName = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Protocol(v ListLoadBalanceListenerRespResponseContentProtocolEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) SniContainerIdList(v []string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.SniContainerIdList = v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ProvisioningStatus(v ListLoadBalanceListenerRespResponseContentProvisioningStatusEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) CreatedTime(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Http2(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Id(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) DefaultTlsContainerId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) MutualAuthenticationUp(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) CookieName(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.CookieName = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) PoolName(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) TlsCustomizeProtocolId(v int64) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) SniUp(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.SniUp = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) KeepAliveTimeout(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) LbAlgorithm(v ListLoadBalanceListenerRespResponseContentLbAlgorithmEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) TlsUp(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthHttpMethod(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ProxyProtocol(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthType(v ListLoadBalanceListenerRespResponseContentHealthTypeEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthType = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) LoadBalanceFlavor(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.LoadBalanceFlavor = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) LoadBalanceId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ProtocolPort(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthExpectedCode(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) GroupName(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ConnectionLimit(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Deleted(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthMaxRetries(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) TlsCustomizeProtocolName(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.TlsCustomizeProtocolName = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Name(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) PoolId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) SessionPersistence(v ListLoadBalanceListenerRespResponseContentSessionPersistenceEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) GroupEnabled(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthUrlPath(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) CaContainerId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthAdminStateUp(v bool) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) OpStatus(v ListLoadBalanceListenerRespResponseContentOpStatusEnum) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) ControlGroupId(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) HealthTimeout(v int32) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) MultiAzUuid(v string) *ListLoadBalanceListenerRespResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceListenerRespResponseContentBuilder) Build() *ListLoadBalanceListenerRespResponseContent {
	return b.s
}


