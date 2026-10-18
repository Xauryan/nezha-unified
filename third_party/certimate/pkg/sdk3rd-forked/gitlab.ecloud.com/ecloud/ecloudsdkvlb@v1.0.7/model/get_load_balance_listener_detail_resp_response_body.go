// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum string

// List of Protocol
const (
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumHttp GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "HTTP"
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumHttps GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "HTTPS"
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumTcp GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "TCP"
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumUdp GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "UDP"
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumTerminatedHttps GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "TERMINATED_HTTPS"
    GetLoadBalanceListenerDetailRespResponseBodyProtocolEnumSip GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum = "SIP"
)
type GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum string

// List of HealthType
const (
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumHttp GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "HTTP"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumHttps GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "HTTPS"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumTcp GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "TCP"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumPing GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "PING"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumUdp GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "UDP"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumHttp10 GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "HTTP10"
    GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnumHttp11 GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum = "HTTP11"
)
type GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumNone GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "NONE"
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumAppCookie GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "APP_COOKIE"
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumHttpCookie GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "HTTP_COOKIE"
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumSourceIp GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "SOURCE_IP"
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumSourceIpPort GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "SOURCE_IP_PORT"
    GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnumCallId GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum = "CALL_ID"
)
type GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum string

// List of OpStatus
const (
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumInactive GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "INACTIVE"
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumActive GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "ACTIVE"
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumCreating GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "CREATING"
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumDeleting GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "DELETING"
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumUpdating GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "UPDATING"
    GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnumError GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum = "ERROR"
)
type GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumActive GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "ACTIVE"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumBuild GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "BUILD"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumDown GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "DOWN"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumError GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "ERROR"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumPendingCreate GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "PENDING_CREATE"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumPendingUpdate GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "PENDING_UPDATE"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumPendingDelete GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "PENDING_DELETE"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumMigrating GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "MIGRATING"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumDeleted GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "DELETED"
    GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnumUnrecognized GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum = "UNRECOGNIZED"
)
type GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnumRoundRobin GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum = "ROUND_ROBIN"
    GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnumLeastConnections GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnumSourceIp GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum = "SOURCE_IP"
)

type GetLoadBalanceListenerDetailRespResponseBody struct {

    // 健康监控时间间隔(单位：秒)
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds *string `json:"sniContainerIds,omitempty"`
    // 是否是双AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 重定到的监听器名称
	RedirectToListenerName *string `json:"redirectToListenerName,omitempty"`
    // 监听器协议
	Protocol *GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum `json:"protocol,omitempty"`
    // 证书容器中的SNI证书ID列表
	SniContainerIdList []string `json:"sniContainerIdList,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 监听器的ID
	Id *string `json:"id,omitempty"`
    // 开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // 后端服务器组的名称
	PoolName *string `json:"poolName,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // TCP长连接超时时间
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 健康监控类型
	HealthType *GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡规格
	LoadBalanceFlavor *int32 `json:"loadBalanceFlavor,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 协议端口号
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 期望HTTP的状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 访问控制组名称
	GroupName *string `json:"groupName,omitempty"`
    // TLS安全策略名称
	TlsCustomizeProtocolName *string `json:"tlsCustomizeProtocolName,omitempty"`
    // 监听器名称
	Name *string `json:"name,omitempty"`
    // 后端服务器组的ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持
	SessionPersistence *GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 用于监控的url，必须以“/"开头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 监听器状态
	OpStatus *GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum `json:"opStatus,omitempty"`
    // 最大等待时间
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 绑定的访问控制组类型: 黑名单：blacklist，白名单：whitelist
	GroupType *string `json:"groupType,omitempty"`
    // 重定到监听器ID,不为null,说明该监听器开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 是否开启获取真实IP
	Transparent *bool `json:"transparent,omitempty"`
    // 监听器底层状态： ACTIVE（资源已激活），BUILD（资源已创建），DOWN（资源已创建但未激活），ERROR（资源无法工作），PENDING_CREATE,PENDING_UPDATE（资源正在创建当中），PENDING_DELETE（资源在删除当中），MIGRATING（迁移中），DELETED（资源已删除），UNRECOGNIZED（未知状态）
	ProvisioningStatus *GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 是否开启HTTP2协议
	Http2 *bool `json:"http2,omitempty"`
    // 证书容器中的服务器证书的访问链接中的UUID
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // TLS安全策略ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法）,LEAST_CONNECTIONS（最小连接算法）,SOURCE_IP（源IP地址算法）
	LbAlgorithm *GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 健康检查的资源ID
	HealthId *string `json:"healthId,omitempty"`
    // 连接数限制
	ConnectionLimit *int32 `json:"connectionLimit,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // 证书容器中的CA证书的ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 访问控制组的ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s GetLoadBalanceListenerDetailRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceListenerDetailRespResponseBody) GoString() string {
	return s.String()
}

func (s GetLoadBalanceListenerDetailRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthDelay(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthDelay = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetModifiedTime(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetNotes(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Notes = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetSniContainerIds(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.SniContainerIds = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetIsMultiAz(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetRedirectToListenerName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.RedirectToListenerName = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetProtocol(v GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetSniContainerIdList(v []string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.SniContainerIdList = v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetCreatedTime(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Id = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetMutualAuthenticationUp(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetPoolName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.PoolName = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetSniUp(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.SniUp = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetKeepAliveTimeout(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.KeepAliveTimeout = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthHttpMethod(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthHttpMethod = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthType(v GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthType = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetLoadBalanceFlavor(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.LoadBalanceFlavor = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetLoadBalanceId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.LoadBalanceId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetProtocolPort(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ProtocolPort = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthExpectedCode(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthExpectedCode = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetGroupName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetTlsCustomizeProtocolName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.TlsCustomizeProtocolName = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Name = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetPoolId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.PoolId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetSessionPersistence(v GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.SessionPersistence = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthUrlPath(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthUrlPath = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthAdminStateUp(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthAdminStateUp = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetOpStatus(v GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.OpStatus = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthTimeout(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthTimeout = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetMultiAzUuid(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.MultiAzUuid = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetGroupType(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.GroupType = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetRedirectToListenerId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.RedirectToListenerId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetDescription(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetTransparent(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Transparent = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetProvisioningStatus(v GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ProvisioningStatus = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHttp2(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Http2 = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetDefaultTlsContainerId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetCookieName(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.CookieName = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetTlsCustomizeProtocolId(v int64) *GetLoadBalanceListenerDetailRespResponseBody {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetLbAlgorithm(v GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum) *GetLoadBalanceListenerDetailRespResponseBody {
	s.LbAlgorithm = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetTlsUp(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.TlsUp = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetProxyProtocol(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ProxyProtocol = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetConnectionLimit(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ConnectionLimit = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetDeleted(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.Deleted = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetHealthMaxRetries(v int32) *GetLoadBalanceListenerDetailRespResponseBody {
	s.HealthMaxRetries = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetGroupEnabled(v bool) *GetLoadBalanceListenerDetailRespResponseBody {
	s.GroupEnabled = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetCaContainerId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.CaContainerId = &v
	return s
}

func (s *GetLoadBalanceListenerDetailRespResponseBody) SetControlGroupId(v string) *GetLoadBalanceListenerDetailRespResponseBody {
	s.ControlGroupId = &v
	return s
}


type GetLoadBalanceListenerDetailRespResponseBodyBuilder struct {
	s *GetLoadBalanceListenerDetailRespResponseBody
}

func NewGetLoadBalanceListenerDetailRespResponseBodyBuilder() *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	s := &GetLoadBalanceListenerDetailRespResponseBody{}
	b := &GetLoadBalanceListenerDetailRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthDelay(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ModifiedTime(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Notes(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) SniContainerIds(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.SniContainerIds = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) IsMultiAz(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) RedirectToListenerName(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.RedirectToListenerName = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Protocol(v GetLoadBalanceListenerDetailRespResponseBodyProtocolEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) SniContainerIdList(v []string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.SniContainerIdList = v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) CreatedTime(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Id(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) MutualAuthenticationUp(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) PoolName(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) SniUp(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.SniUp = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) KeepAliveTimeout(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthHttpMethod(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthType(v GetLoadBalanceListenerDetailRespResponseBodyHealthTypeEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthType = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) LoadBalanceFlavor(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.LoadBalanceFlavor = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) LoadBalanceId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ProtocolPort(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthExpectedCode(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) GroupName(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) TlsCustomizeProtocolName(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.TlsCustomizeProtocolName = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Name(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) PoolId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) SessionPersistence(v GetLoadBalanceListenerDetailRespResponseBodySessionPersistenceEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthUrlPath(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthAdminStateUp(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) OpStatus(v GetLoadBalanceListenerDetailRespResponseBodyOpStatusEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthTimeout(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) MultiAzUuid(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) GroupType(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) RedirectToListenerId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Description(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Transparent(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Transparent = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ProvisioningStatus(v GetLoadBalanceListenerDetailRespResponseBodyProvisioningStatusEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Http2(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Http2 = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) DefaultTlsContainerId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) CookieName(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.CookieName = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) TlsCustomizeProtocolId(v int64) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) LbAlgorithm(v GetLoadBalanceListenerDetailRespResponseBodyLbAlgorithmEnum) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) TlsUp(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ProxyProtocol(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ConnectionLimit(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Deleted(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.Deleted = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) HealthMaxRetries(v int32) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) GroupEnabled(v bool) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) CaContainerId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) ControlGroupId(v string) *GetLoadBalanceListenerDetailRespResponseBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *GetLoadBalanceListenerDetailRespResponseBodyBuilder) Build() *GetLoadBalanceListenerDetailRespResponseBody {
	return b.s
}


