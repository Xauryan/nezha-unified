// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePoolBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    AsyncCreatePoolBodyLbAlgorithmEnumRoundRobin AsyncCreatePoolBodyLbAlgorithmEnum = "ROUND_ROBIN"
    AsyncCreatePoolBodyLbAlgorithmEnumLeastConnections AsyncCreatePoolBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    AsyncCreatePoolBodyLbAlgorithmEnumSourceIp AsyncCreatePoolBodyLbAlgorithmEnum = "SOURCE_IP"
)
type AsyncCreatePoolBodyProtocolEnum string

// List of Protocol
const (
    AsyncCreatePoolBodyProtocolEnumHttp AsyncCreatePoolBodyProtocolEnum = "HTTP"
    AsyncCreatePoolBodyProtocolEnumHttps AsyncCreatePoolBodyProtocolEnum = "HTTPS"
    AsyncCreatePoolBodyProtocolEnumTcp AsyncCreatePoolBodyProtocolEnum = "TCP"
    AsyncCreatePoolBodyProtocolEnumUdp AsyncCreatePoolBodyProtocolEnum = "UDP"
    AsyncCreatePoolBodyProtocolEnumTerminatedHttps AsyncCreatePoolBodyProtocolEnum = "TERMINATED_HTTPS"
    AsyncCreatePoolBodyProtocolEnumSip AsyncCreatePoolBodyProtocolEnum = "SIP"
)
type AsyncCreatePoolBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    AsyncCreatePoolBodySessionPersistenceEnumNone AsyncCreatePoolBodySessionPersistenceEnum = "NONE"
    AsyncCreatePoolBodySessionPersistenceEnumAppCookie AsyncCreatePoolBodySessionPersistenceEnum = "APP_COOKIE"
    AsyncCreatePoolBodySessionPersistenceEnumHttpCookie AsyncCreatePoolBodySessionPersistenceEnum = "HTTP_COOKIE"
    AsyncCreatePoolBodySessionPersistenceEnumSourceIp AsyncCreatePoolBodySessionPersistenceEnum = "SOURCE_IP"
    AsyncCreatePoolBodySessionPersistenceEnumSourceIpPort AsyncCreatePoolBodySessionPersistenceEnum = "SOURCE_IP_PORT"
    AsyncCreatePoolBodySessionPersistenceEnumCallId AsyncCreatePoolBodySessionPersistenceEnum = "CALL_ID"
)

type AsyncCreatePoolBody struct {
    position.Body
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 监听器的调度算法，ROUND_ROBIN（轮询算法），LEAST_CONNECTIONS（最小连接算法），SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *AsyncCreatePoolBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 监听器协议类型
	Protocol *AsyncCreatePoolBodyProtocolEnum `json:"protocol,omitempty"`
    // 健康检查相关创建参数
	HealthMonitorCreateReq *AsyncCreatePoolRequestHealthMonitorCreateReq `json:"healthMonitorCreateReq,omitempty"`
    // 会话保持类型：NONE（无会话保持），APP_COOKIE（基于应用程序Cookie的会话保持），HTTP_COOKIE（基于负载均衡器植入Cookie的会话保持），SOURCE_IP（基于源IP地址的会话保持），SOURCE_IP_PORT（基于源IP地址和源端口的会话保持），CALL_ID（基于呼叫ID的会话保持-特定场景）（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *AsyncCreatePoolBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 后端服务器组名称为5~64位的英文、数字、下划线、中划线等的组合
	PoolName *string `json:"poolName,omitempty"`
}

func (s AsyncCreatePoolBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePoolBody) GoString() string {
	return s.String()
}

func (s AsyncCreatePoolBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePoolBody) SetListenerId(v string) *AsyncCreatePoolBody {
	s.ListenerId = &v
	return s
}

func (s *AsyncCreatePoolBody) SetLbAlgorithm(v AsyncCreatePoolBodyLbAlgorithmEnum) *AsyncCreatePoolBody {
	s.LbAlgorithm = &v
	return s
}

func (s *AsyncCreatePoolBody) SetProtocol(v AsyncCreatePoolBodyProtocolEnum) *AsyncCreatePoolBody {
	s.Protocol = &v
	return s
}

func (s *AsyncCreatePoolBody) SetHealthMonitorCreateReq(v *AsyncCreatePoolRequestHealthMonitorCreateReq) *AsyncCreatePoolBody {
	s.HealthMonitorCreateReq = v
	return s
}

func (s *AsyncCreatePoolBody) SetSessionPersistence(v AsyncCreatePoolBodySessionPersistenceEnum) *AsyncCreatePoolBody {
	s.SessionPersistence = &v
	return s
}

func (s *AsyncCreatePoolBody) SetLoadBalanceId(v string) *AsyncCreatePoolBody {
	s.LoadBalanceId = &v
	return s
}

func (s *AsyncCreatePoolBody) SetIsMultiAz(v bool) *AsyncCreatePoolBody {
	s.IsMultiAz = &v
	return s
}

func (s *AsyncCreatePoolBody) SetPoolName(v string) *AsyncCreatePoolBody {
	s.PoolName = &v
	return s
}


type AsyncCreatePoolBodyBuilder struct {
	s *AsyncCreatePoolBody
}

func NewAsyncCreatePoolBodyBuilder() *AsyncCreatePoolBodyBuilder {
	s := &AsyncCreatePoolBody{}
	b := &AsyncCreatePoolBodyBuilder{s: s}
	return b
}

func (b *AsyncCreatePoolBodyBuilder) ListenerId(v string) *AsyncCreatePoolBodyBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) LbAlgorithm(v AsyncCreatePoolBodyLbAlgorithmEnum) *AsyncCreatePoolBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) Protocol(v AsyncCreatePoolBodyProtocolEnum) *AsyncCreatePoolBodyBuilder {
	b.s.Protocol = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) HealthMonitorCreateReq(v *AsyncCreatePoolRequestHealthMonitorCreateReq) *AsyncCreatePoolBodyBuilder {
	b.s.HealthMonitorCreateReq = v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) SessionPersistence(v AsyncCreatePoolBodySessionPersistenceEnum) *AsyncCreatePoolBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) LoadBalanceId(v string) *AsyncCreatePoolBodyBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) IsMultiAz(v bool) *AsyncCreatePoolBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) PoolName(v string) *AsyncCreatePoolBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *AsyncCreatePoolBodyBuilder) Build() *AsyncCreatePoolBody {
	return b.s
}


