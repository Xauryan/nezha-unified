// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPoolRespResponseBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    GetPoolRespResponseBodyLbAlgorithmEnumRoundRobin GetPoolRespResponseBodyLbAlgorithmEnum = "ROUND_ROBIN"
    GetPoolRespResponseBodyLbAlgorithmEnumLeastConnections GetPoolRespResponseBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    GetPoolRespResponseBodyLbAlgorithmEnumSourceIp GetPoolRespResponseBodyLbAlgorithmEnum = "SOURCE_IP"
)
type GetPoolRespResponseBodyProtocolEnum string

// List of Protocol
const (
    GetPoolRespResponseBodyProtocolEnumHttp GetPoolRespResponseBodyProtocolEnum = "HTTP"
    GetPoolRespResponseBodyProtocolEnumHttps GetPoolRespResponseBodyProtocolEnum = "HTTPS"
    GetPoolRespResponseBodyProtocolEnumTcp GetPoolRespResponseBodyProtocolEnum = "TCP"
    GetPoolRespResponseBodyProtocolEnumUdp GetPoolRespResponseBodyProtocolEnum = "UDP"
    GetPoolRespResponseBodyProtocolEnumTerminatedHttps GetPoolRespResponseBodyProtocolEnum = "TERMINATED_HTTPS"
    GetPoolRespResponseBodyProtocolEnumSip GetPoolRespResponseBodyProtocolEnum = "SIP"
)
type GetPoolRespResponseBodyProvisioningStatusEnum string

// List of ProvisioningStatus
const (
    GetPoolRespResponseBodyProvisioningStatusEnumActive GetPoolRespResponseBodyProvisioningStatusEnum = "ACTIVE"
    GetPoolRespResponseBodyProvisioningStatusEnumBuild GetPoolRespResponseBodyProvisioningStatusEnum = "BUILD"
    GetPoolRespResponseBodyProvisioningStatusEnumDown GetPoolRespResponseBodyProvisioningStatusEnum = "DOWN"
    GetPoolRespResponseBodyProvisioningStatusEnumError GetPoolRespResponseBodyProvisioningStatusEnum = "ERROR"
    GetPoolRespResponseBodyProvisioningStatusEnumPendingCreate GetPoolRespResponseBodyProvisioningStatusEnum = "PENDING_CREATE"
    GetPoolRespResponseBodyProvisioningStatusEnumPendingUpdate GetPoolRespResponseBodyProvisioningStatusEnum = "PENDING_UPDATE"
    GetPoolRespResponseBodyProvisioningStatusEnumPendingDelete GetPoolRespResponseBodyProvisioningStatusEnum = "PENDING_DELETE"
    GetPoolRespResponseBodyProvisioningStatusEnumMigrating GetPoolRespResponseBodyProvisioningStatusEnum = "MIGRATING"
    GetPoolRespResponseBodyProvisioningStatusEnumDeleted GetPoolRespResponseBodyProvisioningStatusEnum = "DELETED"
    GetPoolRespResponseBodyProvisioningStatusEnumUnrecognized GetPoolRespResponseBodyProvisioningStatusEnum = "UNRECOGNIZED"
)
type GetPoolRespResponseBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    GetPoolRespResponseBodySessionPersistenceEnumNone GetPoolRespResponseBodySessionPersistenceEnum = "NONE"
    GetPoolRespResponseBodySessionPersistenceEnumAppCookie GetPoolRespResponseBodySessionPersistenceEnum = "APP_COOKIE"
    GetPoolRespResponseBodySessionPersistenceEnumHttpCookie GetPoolRespResponseBodySessionPersistenceEnum = "HTTP_COOKIE"
    GetPoolRespResponseBodySessionPersistenceEnumSourceIp GetPoolRespResponseBodySessionPersistenceEnum = "SOURCE_IP"
    GetPoolRespResponseBodySessionPersistenceEnumSourceIpPort GetPoolRespResponseBodySessionPersistenceEnum = "SOURCE_IP_PORT"
    GetPoolRespResponseBodySessionPersistenceEnumCallId GetPoolRespResponseBodySessionPersistenceEnum = "CALL_ID"
)

type GetPoolRespResponseBody struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 监听器的调度算法：ROUND_ROBIN（轮询算法），LEAST_CONNECTIONS（最小连接算法），SOURCE_IP（源IP地址算法）
	LbAlgorithm *GetPoolRespResponseBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 健康检查相关信息
	HealthMonitorResp *GetPoolRespResponseHealthMonitorResp `json:"healthMonitorResp,omitempty"`
    // 监听器协议
	Protocol *GetPoolRespResponseBodyProtocolEnum `json:"protocol,omitempty"`
    // 绑定的主机数量
	CountBindMember *int32 `json:"countBindMember,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 后端服务器组状态：ACTIVE（运行），BUILD（创建），DOWN（停止），ERROR（错误），PENDING_CREATE（创建挂起），PENDING_UPDATE（更新挂起），PENDING_DELETE（删除挂起），DELETED（已删除），MIGRATING（迁移中），UPGRADING（升级中），PART_DELETED（部分删除），UNRECOGNIZED（未识别状态）
	ProvisioningStatus *GetPoolRespResponseBodyProvisioningStatusEnum `json:"provisioningStatus,omitempty"`
    // 绑定的转发策略
	L7PolicyResps *[]GetPoolRespResponseL7PolicyResps `json:"l7PolicyResps,omitempty"`
    // 负载均衡监听器名称
	ListenerName *string `json:"listenerName,omitempty"`
    // 服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持类型：NONE（无会话保持），APP_COOKIE（基于应用程序Cookie的会话保持），HTTP_COOKIE（基于负载均衡器植入Cookie的会话保持），SOURCE_IP（基于源IP地址的会话保持），SOURCE_IP_PORT（基于源IP地址和源端口的会话保持），CALL_ID（基于呼叫ID的会话保持-特定场景）
	SessionPersistence *GetPoolRespResponseBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 服务器组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetPoolRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespResponseBody) GoString() string {
	return s.String()
}

func (s GetPoolRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespResponseBody) SetModifiedTime(v string) *GetPoolRespResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetPoolRespResponseBody) SetLbAlgorithm(v GetPoolRespResponseBodyLbAlgorithmEnum) *GetPoolRespResponseBody {
	s.LbAlgorithm = &v
	return s
}

func (s *GetPoolRespResponseBody) SetDescription(v string) *GetPoolRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetPoolRespResponseBody) SetLoadBalanceId(v string) *GetPoolRespResponseBody {
	s.LoadBalanceId = &v
	return s
}

func (s *GetPoolRespResponseBody) SetIsMultiAz(v bool) *GetPoolRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetPoolRespResponseBody) SetListenerId(v string) *GetPoolRespResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetPoolRespResponseBody) SetHealthMonitorResp(v *GetPoolRespResponseHealthMonitorResp) *GetPoolRespResponseBody {
	s.HealthMonitorResp = v
	return s
}

func (s *GetPoolRespResponseBody) SetProtocol(v GetPoolRespResponseBodyProtocolEnum) *GetPoolRespResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetPoolRespResponseBody) SetCountBindMember(v int32) *GetPoolRespResponseBody {
	s.CountBindMember = &v
	return s
}

func (s *GetPoolRespResponseBody) SetDeleted(v bool) *GetPoolRespResponseBody {
	s.Deleted = &v
	return s
}

func (s *GetPoolRespResponseBody) SetProvisioningStatus(v GetPoolRespResponseBodyProvisioningStatusEnum) *GetPoolRespResponseBody {
	s.ProvisioningStatus = &v
	return s
}

func (s *GetPoolRespResponseBody) SetL7PolicyResps(v []GetPoolRespResponseL7PolicyResps) *GetPoolRespResponseBody {
	s.L7PolicyResps = &v
	return s
}

func (s *GetPoolRespResponseBody) SetListenerName(v string) *GetPoolRespResponseBody {
	s.ListenerName = &v
	return s
}

func (s *GetPoolRespResponseBody) SetPoolId(v string) *GetPoolRespResponseBody {
	s.PoolId = &v
	return s
}

func (s *GetPoolRespResponseBody) SetSessionPersistence(v GetPoolRespResponseBodySessionPersistenceEnum) *GetPoolRespResponseBody {
	s.SessionPersistence = &v
	return s
}

func (s *GetPoolRespResponseBody) SetCreatedTime(v string) *GetPoolRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetPoolRespResponseBody) SetMultiAzUuid(v string) *GetPoolRespResponseBody {
	s.MultiAzUuid = &v
	return s
}

func (s *GetPoolRespResponseBody) SetCookieName(v string) *GetPoolRespResponseBody {
	s.CookieName = &v
	return s
}

func (s *GetPoolRespResponseBody) SetPoolName(v string) *GetPoolRespResponseBody {
	s.PoolName = &v
	return s
}


type GetPoolRespResponseBodyBuilder struct {
	s *GetPoolRespResponseBody
}

func NewGetPoolRespResponseBodyBuilder() *GetPoolRespResponseBodyBuilder {
	s := &GetPoolRespResponseBody{}
	b := &GetPoolRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetPoolRespResponseBodyBuilder) ModifiedTime(v string) *GetPoolRespResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) LbAlgorithm(v GetPoolRespResponseBodyLbAlgorithmEnum) *GetPoolRespResponseBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) Description(v string) *GetPoolRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) LoadBalanceId(v string) *GetPoolRespResponseBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) IsMultiAz(v bool) *GetPoolRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) ListenerId(v string) *GetPoolRespResponseBodyBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) HealthMonitorResp(v *GetPoolRespResponseHealthMonitorResp) *GetPoolRespResponseBodyBuilder {
	b.s.HealthMonitorResp = v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) Protocol(v GetPoolRespResponseBodyProtocolEnum) *GetPoolRespResponseBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) CountBindMember(v int32) *GetPoolRespResponseBodyBuilder {
	b.s.CountBindMember = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) Deleted(v bool) *GetPoolRespResponseBodyBuilder {
	b.s.Deleted = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) ProvisioningStatus(v GetPoolRespResponseBodyProvisioningStatusEnum) *GetPoolRespResponseBodyBuilder {
	b.s.ProvisioningStatus = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) L7PolicyResps(v []GetPoolRespResponseL7PolicyResps) *GetPoolRespResponseBodyBuilder {
	b.s.L7PolicyResps = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) ListenerName(v string) *GetPoolRespResponseBodyBuilder {
	b.s.ListenerName = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) PoolId(v string) *GetPoolRespResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) SessionPersistence(v GetPoolRespResponseBodySessionPersistenceEnum) *GetPoolRespResponseBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) CreatedTime(v string) *GetPoolRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) MultiAzUuid(v string) *GetPoolRespResponseBodyBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) CookieName(v string) *GetPoolRespResponseBodyBuilder {
	b.s.CookieName = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) PoolName(v string) *GetPoolRespResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPoolRespResponseBodyBuilder) Build() *GetPoolRespResponseBody {
	return b.s
}


