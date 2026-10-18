// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListPoolRespResponseContentLbAlgorithmEnum string

// List of LbAlgorithm
const (
    ListPoolRespResponseContentLbAlgorithmEnumRoundRobin ListPoolRespResponseContentLbAlgorithmEnum = "ROUND_ROBIN"
    ListPoolRespResponseContentLbAlgorithmEnumLeastConnections ListPoolRespResponseContentLbAlgorithmEnum = "LEAST_CONNECTIONS"
    ListPoolRespResponseContentLbAlgorithmEnumSourceIp ListPoolRespResponseContentLbAlgorithmEnum = "SOURCE_IP"
)
type ListPoolRespResponseContentProtocolEnum string

// List of Protocol
const (
    ListPoolRespResponseContentProtocolEnumHttp ListPoolRespResponseContentProtocolEnum = "HTTP"
    ListPoolRespResponseContentProtocolEnumHttps ListPoolRespResponseContentProtocolEnum = "HTTPS"
    ListPoolRespResponseContentProtocolEnumTcp ListPoolRespResponseContentProtocolEnum = "TCP"
    ListPoolRespResponseContentProtocolEnumUdp ListPoolRespResponseContentProtocolEnum = "UDP"
    ListPoolRespResponseContentProtocolEnumTerminatedHttps ListPoolRespResponseContentProtocolEnum = "TERMINATED_HTTPS"
    ListPoolRespResponseContentProtocolEnumSip ListPoolRespResponseContentProtocolEnum = "SIP"
)
type ListPoolRespResponseContentProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    ListPoolRespResponseContentProvisioningStatusEnumActive ListPoolRespResponseContentProvisioningStatusEnum = "ACTIVE"
    ListPoolRespResponseContentProvisioningStatusEnumBuild ListPoolRespResponseContentProvisioningStatusEnum = "BUILD"
    ListPoolRespResponseContentProvisioningStatusEnumDown ListPoolRespResponseContentProvisioningStatusEnum = "DOWN"
    ListPoolRespResponseContentProvisioningStatusEnumError ListPoolRespResponseContentProvisioningStatusEnum = "ERROR"
    ListPoolRespResponseContentProvisioningStatusEnumPendingCreate ListPoolRespResponseContentProvisioningStatusEnum = "PENDING_CREATE"
    ListPoolRespResponseContentProvisioningStatusEnumPendingUpdate ListPoolRespResponseContentProvisioningStatusEnum = "PENDING_UPDATE"
    ListPoolRespResponseContentProvisioningStatusEnumPendingDelete ListPoolRespResponseContentProvisioningStatusEnum = "PENDING_DELETE"
    ListPoolRespResponseContentProvisioningStatusEnumMigrating ListPoolRespResponseContentProvisioningStatusEnum = "MIGRATING"
    ListPoolRespResponseContentProvisioningStatusEnumDeleted ListPoolRespResponseContentProvisioningStatusEnum = "DELETED"
    ListPoolRespResponseContentProvisioningStatusEnumUnrecognized ListPoolRespResponseContentProvisioningStatusEnum = "UNRECOGNIZED"
)
type ListPoolRespResponseContentSessionPersistenceEnum string

// List of SessionPersistence
const (
    ListPoolRespResponseContentSessionPersistenceEnumNone ListPoolRespResponseContentSessionPersistenceEnum = "NONE"
    ListPoolRespResponseContentSessionPersistenceEnumAppCookie ListPoolRespResponseContentSessionPersistenceEnum = "APP_COOKIE"
    ListPoolRespResponseContentSessionPersistenceEnumHttpCookie ListPoolRespResponseContentSessionPersistenceEnum = "HTTP_COOKIE"
    ListPoolRespResponseContentSessionPersistenceEnumSourceIp ListPoolRespResponseContentSessionPersistenceEnum = "SOURCE_IP"
    ListPoolRespResponseContentSessionPersistenceEnumSourceIpPort ListPoolRespResponseContentSessionPersistenceEnum = "SOURCE_IP_PORT"
    ListPoolRespResponseContentSessionPersistenceEnumCallId ListPoolRespResponseContentSessionPersistenceEnum = "CALL_ID"
)

type ListPoolRespResponseContent struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法），LEAST_CONNECTIONS（最小连接算法），SOURCE_IP（源IP地址算法）
	LbAlgorithm *ListPoolRespResponseContentLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 健康检查相关信息
	HealthMonitorResp *ListPoolRespResponseHealthMonitorResp `json:"healthMonitorResp,omitempty"`
    // 监听器协议类型
	Protocol *ListPoolRespResponseContentProtocolEnum `json:"protocol,omitempty"`
    // 绑定的主机数量
	CountBindMember *int32 `json:"countBindMember,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 后端服务器组状态：ACTIVE（运行），BUILD（创建），DOWN（停止），ERROR（错误），PENDING_CREATE（创建挂起），PENDING_UPDATE（更新挂起），PENDING_DELETE（删除挂起），DELETED（已删除），MIGRATING（迁移中），UPGRADING（升级中），PART_DELETED（部分删除），UNRECOGNIZED（未识别状态）
	ProvisioningStatus *ListPoolRespResponseContentProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 绑定的转发策略
	L7PolicyResps *[]ListPoolRespResponseL7PolicyResps `json:"l7PolicyResps,omitempty"`
    // 负载均衡监听器名称
	ListenerName *string `json:"listenerName,omitempty"`
    // 服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持类型：NONE（无会话保持），APP_COOKIE（基于应用程序Cookie的会话保持），HTTP_COOKIE（基于负载均衡器植入Cookie的会话保持），SOURCE_IP（基于源IP地址的会话保持），SOURCE_IP_PORT（基于源IP地址和源端口的会话保持），CALL_ID（基于呼叫ID的会话保持-特定场景）
	SessionPersistence *ListPoolRespResponseContentSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ListPoolRespResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespResponseContent) GoString() string {
	return s.String()
}

func (s ListPoolRespResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespResponseContent) SetModifiedTime(v string) *ListPoolRespResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListPoolRespResponseContent) SetLbAlgorithm(v ListPoolRespResponseContentLbAlgorithmEnum) *ListPoolRespResponseContent {
	s.LbAlgorithm = &v
	return s
}

func (s *ListPoolRespResponseContent) SetDescription(v string) *ListPoolRespResponseContent {
	s.Description = &v
	return s
}

func (s *ListPoolRespResponseContent) SetLoadBalanceId(v string) *ListPoolRespResponseContent {
	s.LoadBalanceId = &v
	return s
}

func (s *ListPoolRespResponseContent) SetIsMultiAz(v bool) *ListPoolRespResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListPoolRespResponseContent) SetListenerId(v string) *ListPoolRespResponseContent {
	s.ListenerId = &v
	return s
}

func (s *ListPoolRespResponseContent) SetHealthMonitorResp(v *ListPoolRespResponseHealthMonitorResp) *ListPoolRespResponseContent {
	s.HealthMonitorResp = v
	return s
}

func (s *ListPoolRespResponseContent) SetProtocol(v ListPoolRespResponseContentProtocolEnum) *ListPoolRespResponseContent {
	s.Protocol = &v
	return s
}

func (s *ListPoolRespResponseContent) SetCountBindMember(v int32) *ListPoolRespResponseContent {
	s.CountBindMember = &v
	return s
}

func (s *ListPoolRespResponseContent) SetDeleted(v bool) *ListPoolRespResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListPoolRespResponseContent) SetProvisioningStatus(v ListPoolRespResponseContentProvisioningStatusEnum) *ListPoolRespResponseContent {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListPoolRespResponseContent) SetL7PolicyResps(v []ListPoolRespResponseL7PolicyResps) *ListPoolRespResponseContent {
	s.L7PolicyResps = &v
	return s
}

func (s *ListPoolRespResponseContent) SetListenerName(v string) *ListPoolRespResponseContent {
	s.ListenerName = &v
	return s
}

func (s *ListPoolRespResponseContent) SetPoolId(v string) *ListPoolRespResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListPoolRespResponseContent) SetSessionPersistence(v ListPoolRespResponseContentSessionPersistenceEnum) *ListPoolRespResponseContent {
	s.SessionPersistence = &v
	return s
}

func (s *ListPoolRespResponseContent) SetCreatedTime(v string) *ListPoolRespResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListPoolRespResponseContent) SetMultiAzUuid(v string) *ListPoolRespResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListPoolRespResponseContent) SetCookieName(v string) *ListPoolRespResponseContent {
	s.CookieName = &v
	return s
}

func (s *ListPoolRespResponseContent) SetPoolName(v string) *ListPoolRespResponseContent {
	s.PoolName = &v
	return s
}


type ListPoolRespResponseContentBuilder struct {
	s *ListPoolRespResponseContent
}

func NewListPoolRespResponseContentBuilder() *ListPoolRespResponseContentBuilder {
	s := &ListPoolRespResponseContent{}
	b := &ListPoolRespResponseContentBuilder{s: s}
	return b
}

func (b *ListPoolRespResponseContentBuilder) ModifiedTime(v string) *ListPoolRespResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) LbAlgorithm(v ListPoolRespResponseContentLbAlgorithmEnum) *ListPoolRespResponseContentBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) Description(v string) *ListPoolRespResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) LoadBalanceId(v string) *ListPoolRespResponseContentBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) IsMultiAz(v bool) *ListPoolRespResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) ListenerId(v string) *ListPoolRespResponseContentBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) HealthMonitorResp(v *ListPoolRespResponseHealthMonitorResp) *ListPoolRespResponseContentBuilder {
	b.s.HealthMonitorResp = v
	return b
}

func (b *ListPoolRespResponseContentBuilder) Protocol(v ListPoolRespResponseContentProtocolEnum) *ListPoolRespResponseContentBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) CountBindMember(v int32) *ListPoolRespResponseContentBuilder {
	b.s.CountBindMember = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) Deleted(v bool) *ListPoolRespResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) ProvisioningStatus(v ListPoolRespResponseContentProvisioningStatusEnum) *ListPoolRespResponseContentBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) L7PolicyResps(v []ListPoolRespResponseL7PolicyResps) *ListPoolRespResponseContentBuilder {
	b.s.L7PolicyResps = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) ListenerName(v string) *ListPoolRespResponseContentBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) PoolId(v string) *ListPoolRespResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) SessionPersistence(v ListPoolRespResponseContentSessionPersistenceEnum) *ListPoolRespResponseContentBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) CreatedTime(v string) *ListPoolRespResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) MultiAzUuid(v string) *ListPoolRespResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) CookieName(v string) *ListPoolRespResponseContentBuilder {
	b.s.CookieName = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) PoolName(v string) *ListPoolRespResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListPoolRespResponseContentBuilder) Build() *ListPoolRespResponseContent {
	return b.s
}


