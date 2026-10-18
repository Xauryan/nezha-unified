// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceUDPListenerResponseContentProtocolEnum string

// List of Protocol
const (
    ListLoadBalanceUDPListenerResponseContentProtocolEnumHttp ListLoadBalanceUDPListenerResponseContentProtocolEnum = "HTTP"
    ListLoadBalanceUDPListenerResponseContentProtocolEnumHttps ListLoadBalanceUDPListenerResponseContentProtocolEnum = "HTTPS"
    ListLoadBalanceUDPListenerResponseContentProtocolEnumTcp ListLoadBalanceUDPListenerResponseContentProtocolEnum = "TCP"
    ListLoadBalanceUDPListenerResponseContentProtocolEnumUdp ListLoadBalanceUDPListenerResponseContentProtocolEnum = "UDP"
    ListLoadBalanceUDPListenerResponseContentProtocolEnumTerminatedHttps ListLoadBalanceUDPListenerResponseContentProtocolEnum = "TERMINATED_HTTPS"
    ListLoadBalanceUDPListenerResponseContentProtocolEnumSip ListLoadBalanceUDPListenerResponseContentProtocolEnum = "SIP"
)
type ListLoadBalanceUDPListenerResponseContentHealthTypeEnum string

// List of HealthType
const (
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumHttp ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "HTTP"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumHttps ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "HTTPS"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumTcp ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "TCP"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumPing ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "PING"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumUdp ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "UDP"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumHttp10 ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "HTTP10"
    ListLoadBalanceUDPListenerResponseContentHealthTypeEnumHttp11 ListLoadBalanceUDPListenerResponseContentHealthTypeEnum = "HTTP11"
)
type ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum string

// List of SessionPersistence
const (
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumNone ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "NONE"
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumAppCookie ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "APP_COOKIE"
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumHttpCookie ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "HTTP_COOKIE"
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumSourceIp ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP"
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumSourceIpPort ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP_PORT"
    ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnumCallId ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum = "CALL_ID"
)
type ListLoadBalanceUDPListenerResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumInactive ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "INACTIVE"
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumActive ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "ACTIVE"
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumCreating ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "CREATING"
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumDeleting ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "DELETING"
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumUpdating ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "UPDATING"
    ListLoadBalanceUDPListenerResponseContentOpStatusEnumError ListLoadBalanceUDPListenerResponseContentOpStatusEnum = "ERROR"
)
type ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumActive ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "ACTIVE"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumBuild ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "BUILD"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumDown ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "DOWN"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumError ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "ERROR"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumPendingCreate ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "PENDING_CREATE"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumPendingUpdate ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "PENDING_UPDATE"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumPendingDelete ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "PENDING_DELETE"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumMigrating ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "MIGRATING"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumDeleted ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "DELETED"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumUpgrading ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "UPGRADING"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumPartDeleted ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "PART_DELETED"
    ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnumUnrecognized ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum = "UNRECOGNIZED"
)
type ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum string

// List of LbAlgorithm
const (
    ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnumRoundRobin ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum = "ROUND_ROBIN"
    ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnumLeastConnections ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum = "LEAST_CONNECTIONS"
    ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnumSourceIp ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum = "SOURCE_IP"
)

type ListLoadBalanceUDPListenerResponseContent struct {

    // 健康监控时间间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 监听器备注
	Notes *string `json:"notes,omitempty"`
    // 证书容器中的SNI证书ID
	SniContainerIds *string `json:"sniContainerIds,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 重定到的listener名称
	RedirectToListenerName *string `json:"redirectToListenerName,omitempty"`
    // 协议
	Protocol *ListLoadBalanceUDPListenerResponseContentProtocolEnum `json:"protocol,omitempty"`
    // 证书容器中的SNI证书IDList
	SniContainerIdList []string `json:"sniContainerIdList,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 监听器ID
	Id *string `json:"id,omitempty"`
    // 开启双向认证
	MutualAuthenticationUp *bool `json:"mutualAuthenticationUp,omitempty"`
    // 响应超时
	ResponseTimeout *int32 `json:"responseTimeout,omitempty"`
    // 后端服务器组名称
	PoolName *string `json:"poolName,omitempty"`
    // 是否开启SNI证书
	SniUp *bool `json:"sniUp,omitempty"`
    // TCP长连接超时时间
	KeepAliveTimeout *int32 `json:"keepAliveTimeout,omitempty"`
    // 用于监控的HTTP方法
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 健康监控类型
	HealthType *ListLoadBalanceUDPListenerResponseContentHealthTypeEnum `json:"healthType,omitempty"`
    // 负载均衡规格
	LoadBalanceFlavor *int32 `json:"loadBalanceFlavor,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 协议端口号
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
    // 期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 访问控制组名称
	GroupName *string `json:"groupName,omitempty"`
    // TLS安全策略Name
	TlsCustomizeProtocolName *string `json:"tlsCustomizeProtocolName,omitempty"`
    // 名称
	Name *string `json:"name,omitempty"`
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持
	SessionPersistence *ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 用于监控的URL，必须以/打头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 监听器状态
	OpStatus *ListLoadBalanceUDPListenerResponseContentOpStatusEnum `json:"opStatus,omitempty"`
    // 最大等待时间
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为Uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 控制Host内容是否保持不变
	HostHeadMode *string `json:"hostHeadMode,omitempty"`
    // 黑名单，白名单
	GroupType *string `json:"groupType,omitempty"`
    // 重定到listenerId，不为Null，说明该listener开启了重定向
	RedirectToListenerId *string `json:"redirectToListenerId,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 是否透传源ID
	Transparent *bool `json:"transparent,omitempty"`
    // 最大并发数
	MaxConcurrency *int32 `json:"maxConcurrency,omitempty"`
    // 新建连接数
	NewConnection *int32 `json:"newConnection,omitempty"`
    // 监听器底层状态
	ProvisioningStatus *ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 是否是HTTP2
	Http2 *bool `json:"http2,omitempty"`
    // 监听器启停状态
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 证书（容器）Uuid
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // Cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 请求超时时间
	RequestTimeout *int32 `json:"requestTimeout,omitempty"`
    // TLS安全策略ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // LB算法，轮询模式
	LbAlgorithm *ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 是否开启TLS安全策略
	TlsUp *bool `json:"tlsUp,omitempty"`
    // 监听器是否支持proxy_protocol
	ProxyProtocol *bool `json:"proxyProtocol,omitempty"`
    // 健康检查的资源ID
	HealthId *string `json:"healthId,omitempty"`
    // 是否支持数据压缩 true：支持 false：不支持
	DataCompression *bool `json:"dataCompression,omitempty"`
    // 最大连接数
	ConnectionLimit *int32 `json:"connectionLimit,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 是否开启访问控制
	GroupEnabled *bool `json:"groupEnabled,omitempty"`
    // Ca容器ID
	CaContainerId *string `json:"caContainerId,omitempty"`
    // 访问控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s ListLoadBalanceUDPListenerResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceUDPListenerResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadBalanceUDPListenerResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthDelay(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthDelay = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetModifiedTime(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetNotes(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.Notes = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetSniContainerIds(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.SniContainerIds = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetIsMultiAz(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetRedirectToListenerName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.RedirectToListenerName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetProtocol(v ListLoadBalanceUDPListenerResponseContentProtocolEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.Protocol = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetSniContainerIdList(v []string) *ListLoadBalanceUDPListenerResponseContent {
	s.SniContainerIdList = v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetCreatedTime(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetMutualAuthenticationUp(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetResponseTimeout(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.ResponseTimeout = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetPoolName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.PoolName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetSniUp(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.SniUp = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetKeepAliveTimeout(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.KeepAliveTimeout = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthHttpMethod(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthHttpMethod = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthType(v ListLoadBalanceUDPListenerResponseContentHealthTypeEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthType = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetLoadBalanceFlavor(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.LoadBalanceFlavor = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetLoadBalanceId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetProtocolPort(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.ProtocolPort = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthExpectedCode(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthExpectedCode = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetGroupName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.GroupName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetTlsCustomizeProtocolName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.TlsCustomizeProtocolName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetPoolId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetSessionPersistence(v ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.SessionPersistence = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthUrlPath(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthUrlPath = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthAdminStateUp(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthAdminStateUp = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetOpStatus(v ListLoadBalanceUDPListenerResponseContentOpStatusEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthTimeout(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthTimeout = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetMultiAzUuid(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHostHeadMode(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.HostHeadMode = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetGroupType(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.GroupType = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetRedirectToListenerId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.RedirectToListenerId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetDescription(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetTransparent(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.Transparent = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetMaxConcurrency(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.MaxConcurrency = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetNewConnection(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.NewConnection = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetProvisioningStatus(v ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHttp2(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.Http2 = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetAdminStateUp(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetDefaultTlsContainerId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetCookieName(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.CookieName = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetRequestTimeout(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.RequestTimeout = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetTlsCustomizeProtocolId(v int64) *ListLoadBalanceUDPListenerResponseContent {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetLbAlgorithm(v ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceUDPListenerResponseContent {
	s.LbAlgorithm = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetTlsUp(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.TlsUp = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetProxyProtocol(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.ProxyProtocol = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetDataCompression(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.DataCompression = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetConnectionLimit(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.ConnectionLimit = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetDeleted(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetHealthMaxRetries(v int32) *ListLoadBalanceUDPListenerResponseContent {
	s.HealthMaxRetries = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetGroupEnabled(v bool) *ListLoadBalanceUDPListenerResponseContent {
	s.GroupEnabled = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetCaContainerId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.CaContainerId = &v
	return s
}

func (s *ListLoadBalanceUDPListenerResponseContent) SetControlGroupId(v string) *ListLoadBalanceUDPListenerResponseContent {
	s.ControlGroupId = &v
	return s
}


type ListLoadBalanceUDPListenerResponseContentBuilder struct {
	s *ListLoadBalanceUDPListenerResponseContent
}

func NewListLoadBalanceUDPListenerResponseContentBuilder() *ListLoadBalanceUDPListenerResponseContentBuilder {
	s := &ListLoadBalanceUDPListenerResponseContent{}
	b := &ListLoadBalanceUDPListenerResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthDelay(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ModifiedTime(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Notes(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Notes = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) SniContainerIds(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.SniContainerIds = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) IsMultiAz(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) RedirectToListenerName(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.RedirectToListenerName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Protocol(v ListLoadBalanceUDPListenerResponseContentProtocolEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) SniContainerIdList(v []string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.SniContainerIdList = v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) CreatedTime(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Id(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) MutualAuthenticationUp(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ResponseTimeout(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ResponseTimeout = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) PoolName(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) SniUp(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.SniUp = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) KeepAliveTimeout(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthHttpMethod(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthType(v ListLoadBalanceUDPListenerResponseContentHealthTypeEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthType = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) LoadBalanceFlavor(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.LoadBalanceFlavor = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) LoadBalanceId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ProtocolPort(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthExpectedCode(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) GroupName(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) TlsCustomizeProtocolName(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Name(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) PoolId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) SessionPersistence(v ListLoadBalanceUDPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthUrlPath(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthAdminStateUp(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) OpStatus(v ListLoadBalanceUDPListenerResponseContentOpStatusEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthTimeout(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) MultiAzUuid(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HostHeadMode(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HostHeadMode = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) GroupType(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) RedirectToListenerId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Description(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Transparent(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) MaxConcurrency(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.MaxConcurrency = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) NewConnection(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.NewConnection = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ProvisioningStatus(v ListLoadBalanceUDPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Http2(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) AdminStateUp(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) DefaultTlsContainerId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) CookieName(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.CookieName = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) RequestTimeout(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.RequestTimeout = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) TlsCustomizeProtocolId(v int64) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) LbAlgorithm(v ListLoadBalanceUDPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) TlsUp(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ProxyProtocol(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) DataCompression(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.DataCompression = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ConnectionLimit(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Deleted(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) HealthMaxRetries(v int32) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) GroupEnabled(v bool) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) CaContainerId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) ControlGroupId(v string) *ListLoadBalanceUDPListenerResponseContentBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListLoadBalanceUDPListenerResponseContentBuilder) Build() *ListLoadBalanceUDPListenerResponseContent {
	return b.s
}


