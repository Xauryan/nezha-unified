// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadBalanceTCPListenerResponseContentProtocolEnum string

// List of Protocol
const (
    ListLoadBalanceTCPListenerResponseContentProtocolEnumHttp ListLoadBalanceTCPListenerResponseContentProtocolEnum = "HTTP"
    ListLoadBalanceTCPListenerResponseContentProtocolEnumHttps ListLoadBalanceTCPListenerResponseContentProtocolEnum = "HTTPS"
    ListLoadBalanceTCPListenerResponseContentProtocolEnumTcp ListLoadBalanceTCPListenerResponseContentProtocolEnum = "TCP"
    ListLoadBalanceTCPListenerResponseContentProtocolEnumUdp ListLoadBalanceTCPListenerResponseContentProtocolEnum = "UDP"
    ListLoadBalanceTCPListenerResponseContentProtocolEnumTerminatedHttps ListLoadBalanceTCPListenerResponseContentProtocolEnum = "TERMINATED_HTTPS"
    ListLoadBalanceTCPListenerResponseContentProtocolEnumSip ListLoadBalanceTCPListenerResponseContentProtocolEnum = "SIP"
)
type ListLoadBalanceTCPListenerResponseContentHealthTypeEnum string

// List of HealthType
const (
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumHttp ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "HTTP"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumHttps ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "HTTPS"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumTcp ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "TCP"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumPing ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "PING"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumUdp ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "UDP"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumHttp10 ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "HTTP10"
    ListLoadBalanceTCPListenerResponseContentHealthTypeEnumHttp11 ListLoadBalanceTCPListenerResponseContentHealthTypeEnum = "HTTP11"
)
type ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum string

// List of SessionPersistence
const (
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumNone ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "NONE"
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumAppCookie ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "APP_COOKIE"
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumHttpCookie ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "HTTP_COOKIE"
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumSourceIp ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP"
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumSourceIpPort ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "SOURCE_IP_PORT"
    ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnumCallId ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum = "CALL_ID"
)
type ListLoadBalanceTCPListenerResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumInactive ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "INACTIVE"
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumActive ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "ACTIVE"
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumCreating ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "CREATING"
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumDeleting ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "DELETING"
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumUpdating ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "UPDATING"
    ListLoadBalanceTCPListenerResponseContentOpStatusEnumError ListLoadBalanceTCPListenerResponseContentOpStatusEnum = "ERROR"
)
type ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumActive ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "ACTIVE"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumBuild ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "BUILD"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumDown ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "DOWN"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumError ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "ERROR"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumPendingCreate ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "PENDING_CREATE"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumPendingUpdate ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "PENDING_UPDATE"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumPendingDelete ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "PENDING_DELETE"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumMigrating ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "MIGRATING"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumDeleted ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "DELETED"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumUpgrading ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "UPGRADING"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumPartDeleted ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "PART_DELETED"
    ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnumUnrecognized ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum = "UNRECOGNIZED"
)
type ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum string

// List of LbAlgorithm
const (
    ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnumRoundRobin ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum = "ROUND_ROBIN"
    ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnumLeastConnections ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum = "LEAST_CONNECTIONS"
    ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnumSourceIp ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum = "SOURCE_IP"
)

type ListLoadBalanceTCPListenerResponseContent struct {

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
	Protocol *ListLoadBalanceTCPListenerResponseContentProtocolEnum `json:"protocol,omitempty"`
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
	HealthType *ListLoadBalanceTCPListenerResponseContentHealthTypeEnum `json:"healthType,omitempty"`
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
	SessionPersistence *ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 用于监控的URL，必须以/打头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	HealthAdminStateUp *bool `json:"healthAdminStateUp,omitempty"`
    // 监听器状态
	OpStatus *ListLoadBalanceTCPListenerResponseContentOpStatusEnum `json:"opStatus,omitempty"`
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
	ProvisioningStatus *ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
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
	LbAlgorithm *ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
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
    // 访问控制组id
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s ListLoadBalanceTCPListenerResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalanceTCPListenerResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadBalanceTCPListenerResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthDelay(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthDelay = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetModifiedTime(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetNotes(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.Notes = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetSniContainerIds(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.SniContainerIds = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetIsMultiAz(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetRedirectToListenerName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.RedirectToListenerName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetProtocol(v ListLoadBalanceTCPListenerResponseContentProtocolEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.Protocol = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetSniContainerIdList(v []string) *ListLoadBalanceTCPListenerResponseContent {
	s.SniContainerIdList = v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetCreatedTime(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetMutualAuthenticationUp(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.MutualAuthenticationUp = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetResponseTimeout(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.ResponseTimeout = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetPoolName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.PoolName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetSniUp(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.SniUp = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetKeepAliveTimeout(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.KeepAliveTimeout = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthHttpMethod(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthHttpMethod = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthType(v ListLoadBalanceTCPListenerResponseContentHealthTypeEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthType = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetLoadBalanceFlavor(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.LoadBalanceFlavor = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetLoadBalanceId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetProtocolPort(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.ProtocolPort = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthExpectedCode(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthExpectedCode = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetGroupName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.GroupName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetTlsCustomizeProtocolName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.TlsCustomizeProtocolName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetPoolId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetSessionPersistence(v ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.SessionPersistence = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthUrlPath(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthUrlPath = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthAdminStateUp(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthAdminStateUp = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetOpStatus(v ListLoadBalanceTCPListenerResponseContentOpStatusEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthTimeout(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthTimeout = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetMultiAzUuid(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHostHeadMode(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.HostHeadMode = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetGroupType(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.GroupType = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetRedirectToListenerId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.RedirectToListenerId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetDescription(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetTransparent(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.Transparent = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetMaxConcurrency(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.MaxConcurrency = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetNewConnection(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.NewConnection = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetProvisioningStatus(v ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHttp2(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.Http2 = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetAdminStateUp(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetDefaultTlsContainerId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.DefaultTlsContainerId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetCookieName(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.CookieName = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetRequestTimeout(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.RequestTimeout = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetTlsCustomizeProtocolId(v int64) *ListLoadBalanceTCPListenerResponseContent {
	s.TlsCustomizeProtocolId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetLbAlgorithm(v ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceTCPListenerResponseContent {
	s.LbAlgorithm = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetTlsUp(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.TlsUp = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetProxyProtocol(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.ProxyProtocol = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetDataCompression(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.DataCompression = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetConnectionLimit(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.ConnectionLimit = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetDeleted(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetHealthMaxRetries(v int32) *ListLoadBalanceTCPListenerResponseContent {
	s.HealthMaxRetries = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetGroupEnabled(v bool) *ListLoadBalanceTCPListenerResponseContent {
	s.GroupEnabled = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetCaContainerId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.CaContainerId = &v
	return s
}

func (s *ListLoadBalanceTCPListenerResponseContent) SetControlGroupId(v string) *ListLoadBalanceTCPListenerResponseContent {
	s.ControlGroupId = &v
	return s
}


type ListLoadBalanceTCPListenerResponseContentBuilder struct {
	s *ListLoadBalanceTCPListenerResponseContent
}

func NewListLoadBalanceTCPListenerResponseContentBuilder() *ListLoadBalanceTCPListenerResponseContentBuilder {
	s := &ListLoadBalanceTCPListenerResponseContent{}
	b := &ListLoadBalanceTCPListenerResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthDelay(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ModifiedTime(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Notes(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Notes = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) SniContainerIds(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.SniContainerIds = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) IsMultiAz(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) RedirectToListenerName(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.RedirectToListenerName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Protocol(v ListLoadBalanceTCPListenerResponseContentProtocolEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) SniContainerIdList(v []string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.SniContainerIdList = v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) CreatedTime(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Id(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) MutualAuthenticationUp(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.MutualAuthenticationUp = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ResponseTimeout(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ResponseTimeout = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) PoolName(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) SniUp(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.SniUp = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) KeepAliveTimeout(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.KeepAliveTimeout = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthHttpMethod(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthType(v ListLoadBalanceTCPListenerResponseContentHealthTypeEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthType = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) LoadBalanceFlavor(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.LoadBalanceFlavor = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) LoadBalanceId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ProtocolPort(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthExpectedCode(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) GroupName(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) TlsCustomizeProtocolName(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Name(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) PoolId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) SessionPersistence(v ListLoadBalanceTCPListenerResponseContentSessionPersistenceEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthUrlPath(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthAdminStateUp(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthAdminStateUp = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) OpStatus(v ListLoadBalanceTCPListenerResponseContentOpStatusEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthTimeout(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) MultiAzUuid(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HostHeadMode(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HostHeadMode = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) GroupType(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) RedirectToListenerId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.RedirectToListenerId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Description(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Transparent(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Transparent = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) MaxConcurrency(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.MaxConcurrency = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) NewConnection(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.NewConnection = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ProvisioningStatus(v ListLoadBalanceTCPListenerResponseContentProvisioningStatusEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Http2(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Http2 = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) AdminStateUp(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) DefaultTlsContainerId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.DefaultTlsContainerId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) CookieName(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.CookieName = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) RequestTimeout(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.RequestTimeout = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) TlsCustomizeProtocolId(v int64) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.TlsCustomizeProtocolId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) LbAlgorithm(v ListLoadBalanceTCPListenerResponseContentLbAlgorithmEnum) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) TlsUp(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.TlsUp = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ProxyProtocol(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ProxyProtocol = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) DataCompression(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.DataCompression = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ConnectionLimit(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ConnectionLimit = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Deleted(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) HealthMaxRetries(v int32) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) GroupEnabled(v bool) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.GroupEnabled = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) CaContainerId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.CaContainerId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) ControlGroupId(v string) *ListLoadBalanceTCPListenerResponseContentBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListLoadBalanceTCPListenerResponseContentBuilder) Build() *ListLoadBalanceTCPListenerResponseContent {
	return b.s
}


