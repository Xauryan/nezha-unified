// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum string

// List of LbAlgorithm
const (
    AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnumRoundRobin AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum = "ROUND_ROBIN"
    AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnumLeastConnections AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum = "LEAST_CONNECTIONS"
    AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnumSourceIp AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum = "SOURCE_IP"
)
type AsyncCreatePolicyRequestPoolCreateReqProtocolEnum string

// List of Protocol
const (
    AsyncCreatePolicyRequestPoolCreateReqProtocolEnumHttp AsyncCreatePolicyRequestPoolCreateReqProtocolEnum = "HTTP"
    AsyncCreatePolicyRequestPoolCreateReqProtocolEnumHttps AsyncCreatePolicyRequestPoolCreateReqProtocolEnum = "HTTPS"
    AsyncCreatePolicyRequestPoolCreateReqProtocolEnumTcp AsyncCreatePolicyRequestPoolCreateReqProtocolEnum = "TCP"
    AsyncCreatePolicyRequestPoolCreateReqProtocolEnumTerminatedHttps AsyncCreatePolicyRequestPoolCreateReqProtocolEnum = "TERMINATED_HTTPS"
)
type AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum string

// List of SessionPersistence
const (
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumNone AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "NONE"
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumAppCookie AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "APP_COOKIE"
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumHttpCookie AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "HTTP_COOKIE"
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumSourceIp AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "SOURCE_IP"
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumSourceIpPort AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "SOURCE_IP_PORT"
    AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnumCallId AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum = "CALL_ID"
)

type AsyncCreatePolicyRequestPoolCreateReq struct {

    // Cookie超时时间
	CookieTimeOut *int32 `json:"cookieTimeOut,omitempty"`
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 监听器的调度算法，ROUND_ROBIN（轮询算法），LEAST_CONNECTIONS（最小连接算法），SOURCE_IP（源IP地址算法） （lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 监听器协议类型
	Protocol *AsyncCreatePolicyRequestPoolCreateReqProtocolEnum `json:"protocol,omitempty"`
    // 健康检查相关创建参数
	HealthMonitorCreateReq *AsyncCreatePolicyRequestHealthMonitorCreateReq `json:"healthMonitorCreateReq,omitempty"`
    // 会话保持类型：NONE（无会话保持），APP_COOKIE（基于应用程序Cookie的会话保持），HTTP_COOKIE（基于负载均衡器植入Cookie的会话保持），SOURCE_IP（基于源IP地址的会话保持），SOURCE_IP_PORT（基于源IP地址和源端口的会话保持），CALL_ID（基于呼叫ID的会话保持-特定场景）（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // Cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 负载均衡ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 后端服务器组名称为5~64位的英文、数字、下划线、中划线等的组合，且必须以英文字母开头
	PoolName *string `json:"poolName,omitempty"`
}

func (s AsyncCreatePolicyRequestPoolCreateReq) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePolicyRequestPoolCreateReq) GoString() string {
	return s.String()
}

func (s AsyncCreatePolicyRequestPoolCreateReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetCookieTimeOut(v int32) *AsyncCreatePolicyRequestPoolCreateReq {
	s.CookieTimeOut = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetListenerId(v string) *AsyncCreatePolicyRequestPoolCreateReq {
	s.ListenerId = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetLbAlgorithm(v AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum) *AsyncCreatePolicyRequestPoolCreateReq {
	s.LbAlgorithm = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetProtocol(v AsyncCreatePolicyRequestPoolCreateReqProtocolEnum) *AsyncCreatePolicyRequestPoolCreateReq {
	s.Protocol = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetHealthMonitorCreateReq(v *AsyncCreatePolicyRequestHealthMonitorCreateReq) *AsyncCreatePolicyRequestPoolCreateReq {
	s.HealthMonitorCreateReq = v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetSessionPersistence(v AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum) *AsyncCreatePolicyRequestPoolCreateReq {
	s.SessionPersistence = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetDescription(v string) *AsyncCreatePolicyRequestPoolCreateReq {
	s.Description = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetCookieName(v string) *AsyncCreatePolicyRequestPoolCreateReq {
	s.CookieName = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetLoadBalanceId(v string) *AsyncCreatePolicyRequestPoolCreateReq {
	s.LoadBalanceId = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetIsMultiAz(v bool) *AsyncCreatePolicyRequestPoolCreateReq {
	s.IsMultiAz = &v
	return s
}

func (s *AsyncCreatePolicyRequestPoolCreateReq) SetPoolName(v string) *AsyncCreatePolicyRequestPoolCreateReq {
	s.PoolName = &v
	return s
}


type AsyncCreatePolicyRequestPoolCreateReqBuilder struct {
	s *AsyncCreatePolicyRequestPoolCreateReq
}

func NewAsyncCreatePolicyRequestPoolCreateReqBuilder() *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	s := &AsyncCreatePolicyRequestPoolCreateReq{}
	b := &AsyncCreatePolicyRequestPoolCreateReqBuilder{s: s}
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) CookieTimeOut(v int32) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.CookieTimeOut = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) ListenerId(v string) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) LbAlgorithm(v AsyncCreatePolicyRequestPoolCreateReqLbAlgorithmEnum) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) Protocol(v AsyncCreatePolicyRequestPoolCreateReqProtocolEnum) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.Protocol = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) HealthMonitorCreateReq(v *AsyncCreatePolicyRequestHealthMonitorCreateReq) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.HealthMonitorCreateReq = v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) SessionPersistence(v AsyncCreatePolicyRequestPoolCreateReqSessionPersistenceEnum) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) Description(v string) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.Description = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) CookieName(v string) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.CookieName = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) LoadBalanceId(v string) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) IsMultiAz(v bool) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) PoolName(v string) *AsyncCreatePolicyRequestPoolCreateReqBuilder {
	b.s.PoolName = &v
	return b
}

func (b *AsyncCreatePolicyRequestPoolCreateReqBuilder) Build() *AsyncCreatePolicyRequestPoolCreateReq {
	return b.s
}


