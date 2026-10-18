// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceHTTPListenerResponseContentProtocolEnum string

// List of Protocol
const (
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumHttp ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "HTTP"
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumHttps ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "HTTPS"
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumTcp ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "TCP"
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumUdp ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "UDP"
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumTerminatedHttps ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "TERMINATED_HTTPS"
    ListLoadBalanceHTTPListenerResponseContentProtocolEnumSip ListLoadBalanceHTTPListenerResponseContentProtocolEnum = "SIP"
)
type ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum string

// List of HealthType
const (
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumHttp ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "HTTP"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumHttps ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "HTTPS"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumTcp ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "TCP"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumPing ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "PING"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumUdp ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "UDP"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumHttp10 ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "HTTP10"
    ListLoadBalanceHTTPListenerResponseContentHealthTypeEnumHttp11 ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum = "HTTP11"
)
type ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum string

// List of SessionPersistence
const (
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumNone ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "NONE"
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumAppCookie ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "APP_COOKIE"
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumHttpCookie ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "HTTP_COOKIE"
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumSourceIp ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP"
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumSourceIpPort ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP_PORT"
    ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnumCallId ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum = "CALL_ID"
)
type ListLoadBalanceHTTPListenerResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumInactive ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "INACTIVE"
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumActive ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "ACTIVE"
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumCreating ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "CREATING"
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumDeleting ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "DELETING"
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumUpdating ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "UPDATING"
    ListLoadBalanceHTTPListenerResponseContentOpStatusEnumError ListLoadBalanceHTTPListenerResponseContentOpStatusEnum = "ERROR"
)
type ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumActive ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "ACTIVE"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumBuild ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "BUILD"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumDown ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "DOWN"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumError ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "ERROR"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumPendingCreate ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "PENDING_CREATE"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumPendingUpdate ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "PENDING_UPDATE"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumPendingDelete ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "PENDING_DELETE"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumMigrating ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "MIGRATING"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumDeleted ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "DELETED"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumUpgrading ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "UPGRADING"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumPartDeleted ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "PART_DELETED"
    ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnumUnrecognized ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum = "UNRECOGNIZED"
)
type ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum string

// List of LbAlgorithm
const (
    ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnumRoundRobin ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum = "ROUND_ROBIN"
    ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnumLeastConnections ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum = "LEAST_CONNECTIONS"
    ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnumSourceIp ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum = "SOURCE_IP"
)

type ListLoadBalanceHTTPListenerResponseContent struct {

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
	Protocol *ListLoadBalanceHTTPListenerResponseContentProtocolEnum `json:"protocol,omitempty"`
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
	HealthType *ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum `json:"healthType,omitempty"`
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
	SessionPersistence *ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 用于监控的URL，必须以/打头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 监听器状态
	OpStatus *ListLoadBalanceHTTPListenerResponseContentOpStatusEnum `json:"opStatus,omitempty"`
    // 最大等待时间
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为Uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 控制Host内容是否保持不变
	HostHeadMode *string `json:"hostHeadMode,omitempty"`
    // 黑名单，白名单
	GroupType *string `json:"groupType,omitempty"`
    // 重定到listenerId,不为Null,说明该listener开启了重定向
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
	ProvisioningStatus *ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 是否是HTTP2
	Http2 *bool `json:"http2,omitempty"`
    // 监听器启停状态
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 证书（容器）Uuid
	DefaultTlsContainerId *string `json:"defaultTlsContainerId,omitempty"`
    // cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 请求超时时间
	RequestTimeout *int32 `json:"requestTimeout,omitempty"`
    // TLS安全策略ID
	TlsCustomizeProtocolId *int64 `json:"tlsCustomizeProtocolId,omitempty"`
    // LB算法，轮询模式
	LbAlgorithm *ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
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

func (s ListLoadBalanceHTTPListenerResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceHTTPListenerResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadBalanceHTTPListenerResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthDelay(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthDelay = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetModifiedTime(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetNotes(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.Notes = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetSniContainerIds(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.SniContainerIds = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetIsMultiAz(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetRedirectToListenerName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.RedirectToListenerName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetProtocol(v ListLoadBalanceHTTPListenerResponseContentProtocolEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.Protocol = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetSniContainerIdList(v []string) *ListLoadBalanceHTTPListenerResponseContent {
	s.SniContainerIdList = v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetCreatedTime(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetMutualAuthenticationUp(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetResponseTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.ResponseTimeout = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetPoolName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.PoolName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetSniUp(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.SniUp = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetKeepAliveTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.KeepAliveTimeout = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthHttpMethod(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthHttpMethod = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthType(v ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthType = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetLoadBalanceFlavor(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.LoadBalanceFlavor = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetLoadBalanceId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetProtocolPort(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.ProtocolPort = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthExpectedCode(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthExpectedCode = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetGroupName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.GroupName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetTlsCustomizeProtocolName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.TlsCustomizeProtocolName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetPoolId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetSessionPersistence(v ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.SessionPersistence = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthUrlPath(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthUrlPath = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthAdminStateUp(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthAdminStateUp = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetOpStatus(v ListLoadBalanceHTTPListenerResponseContentOpStatusEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthTimeout = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetMultiAzUuid(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHostHeadMode(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.HostHeadMode = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetGroupType(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.GroupType = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetRedirectToListenerId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.RedirectToListenerId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetDescription(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetTransparent(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.Transparent = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetMaxConcurrency(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.MaxConcurrency = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetNewConnection(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.NewConnection = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetProvisioningStatus(v ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHttp2(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.Http2 = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetAdminStateUp(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetDefaultTlsContainerId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetCookieName(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.CookieName = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetRequestTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.RequestTimeout = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetTlsCustomizeProtocolId(v int64) *ListLoadBalanceHTTPListenerResponseContent {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetLbAlgorithm(v ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceHTTPListenerResponseContent {
	s.LbAlgorithm = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetTlsUp(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.TlsUp = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetProxyProtocol(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.ProxyProtocol = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetDataCompression(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.DataCompression = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetConnectionLimit(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.ConnectionLimit = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetDeleted(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetHealthMaxRetries(v int32) *ListLoadBalanceHTTPListenerResponseContent {
	s.HealthMaxRetries = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetGroupEnabled(v bool) *ListLoadBalanceHTTPListenerResponseContent {
	s.GroupEnabled = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetCaContainerId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.CaContainerId = &v
	return s
}

func (s *ListLoadBalanceHTTPListenerResponseContent) SetControlGroupId(v string) *ListLoadBalanceHTTPListenerResponseContent {
	s.ControlGroupId = &v
	return s
}


type ListLoadBalanceHTTPListenerResponseContentBuilder struct {
	s *ListLoadBalanceHTTPListenerResponseContent
}

func NewListLoadBalanceHTTPListenerResponseContentBuilder() *ListLoadBalanceHTTPListenerResponseContentBuilder {
	s := &ListLoadBalanceHTTPListenerResponseContent{}
	b := &ListLoadBalanceHTTPListenerResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthDelay(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ModifiedTime(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Notes(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Notes = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) SniContainerIds(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.SniContainerIds = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) IsMultiAz(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) RedirectToListenerName(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.RedirectToListenerName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Protocol(v ListLoadBalanceHTTPListenerResponseContentProtocolEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) SniContainerIdList(v []string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.SniContainerIdList = v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) CreatedTime(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Id(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) MutualAuthenticationUp(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ResponseTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ResponseTimeout = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) PoolName(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) SniUp(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.SniUp = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) KeepAliveTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthHttpMethod(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthType(v ListLoadBalanceHTTPListenerResponseContentHealthTypeEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthType = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) LoadBalanceFlavor(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.LoadBalanceFlavor = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) LoadBalanceId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ProtocolPort(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthExpectedCode(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) GroupName(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) TlsCustomizeProtocolName(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Name(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) PoolId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) SessionPersistence(v ListLoadBalanceHTTPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthUrlPath(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthAdminStateUp(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) OpStatus(v ListLoadBalanceHTTPListenerResponseContentOpStatusEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) MultiAzUuid(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HostHeadMode(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HostHeadMode = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) GroupType(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) RedirectToListenerId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Description(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Transparent(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) MaxConcurrency(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.MaxConcurrency = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) NewConnection(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.NewConnection = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ProvisioningStatus(v ListLoadBalanceHTTPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Http2(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) AdminStateUp(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) DefaultTlsContainerId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) CookieName(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.CookieName = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) RequestTimeout(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.RequestTimeout = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) TlsCustomizeProtocolId(v int64) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) LbAlgorithm(v ListLoadBalanceHTTPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) TlsUp(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ProxyProtocol(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) DataCompression(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.DataCompression = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ConnectionLimit(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Deleted(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) HealthMaxRetries(v int32) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) GroupEnabled(v bool) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) CaContainerId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) ControlGroupId(v string) *ListLoadBalanceHTTPListenerResponseContentBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListLoadBalanceHTTPListenerResponseContentBuilder) Build() *ListLoadBalanceHTTPListenerResponseContent {
	return b.s
}


