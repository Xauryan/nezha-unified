// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncUpdateBodyLbAlgorithmEnum string

// List of LbAlgorithm
const (
    AsyncUpdateBodyLbAlgorithmEnumRoundRobin AsyncUpdateBodyLbAlgorithmEnum = "ROUND_ROBIN"
    AsyncUpdateBodyLbAlgorithmEnumLeastConnections AsyncUpdateBodyLbAlgorithmEnum = "LEAST_CONNECTIONS"
    AsyncUpdateBodyLbAlgorithmEnumSourceIp AsyncUpdateBodyLbAlgorithmEnum = "SOURCE_IP"
)
type AsyncUpdateBodySessionPersistenceEnum string

// List of SessionPersistence
const (
    AsyncUpdateBodySessionPersistenceEnumNone AsyncUpdateBodySessionPersistenceEnum = "NONE"
    AsyncUpdateBodySessionPersistenceEnumAppCookie AsyncUpdateBodySessionPersistenceEnum = "APP_COOKIE"
    AsyncUpdateBodySessionPersistenceEnumHttpCookie AsyncUpdateBodySessionPersistenceEnum = "HTTP_COOKIE"
    AsyncUpdateBodySessionPersistenceEnumSourceIp AsyncUpdateBodySessionPersistenceEnum = "SOURCE_IP"
    AsyncUpdateBodySessionPersistenceEnumCallId AsyncUpdateBodySessionPersistenceEnum = "CALL_ID"
    AsyncUpdateBodySessionPersistenceEnumSourceIpPort AsyncUpdateBodySessionPersistenceEnum = "SOURCE_IP_PORT"
)

type AsyncUpdateBody struct {
    position.Body
    // 监听器的调度算法：ROUND_ROBIN（轮询算法），LEAST_CONNECTIONS（最小连接算法），SOURCE_IP（源IP地址算法）（lbAlgorithm字段设置的有默认值，如果填写其他非法值会默认选择ROUND_ROBIN）
	LbAlgorithm *AsyncUpdateBodyLbAlgorithmEnum `json:"lbAlgorithm,omitempty"`
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 会话保持类型：NONE（无会话保持），APP_COOKIE（基于应用程序Cookie的会话保持），HTTP_COOKIE（基于负载均衡器植入Cookie的会话保持），SOURCE_IP（基于源IP地址的会话保持），SOURCE_IP_PORT（基于源IP地址和源端口的会话保持），CALL_ID（基于呼叫ID的会话保持-特定场景）（sessionPersistence参数存在默认值，当取值错误或者不填是默认值为NONE）
	SessionPersistence *AsyncUpdateBodySessionPersistenceEnum `json:"sessionPersistence,omitempty"`
    // 健康检查相关参数
	HealthMonitorUpdateReq *AsyncUpdateRequestHealthMonitorUpdateReq `json:"healthMonitorUpdateReq,omitempty"`
    // Cookie名称
	CookieName *string `json:"cookieName,omitempty"`
    // 后端服务器组名称为5~64位的英文、数字、下划线、中划线等的组合
	PoolName *string `json:"poolName,omitempty"`
}

func (s AsyncUpdateBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateBody) GoString() string {
	return s.String()
}

func (s AsyncUpdateBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateBody) SetLbAlgorithm(v AsyncUpdateBodyLbAlgorithmEnum) *AsyncUpdateBody {
	s.LbAlgorithm = &v
	return s
}

func (s *AsyncUpdateBody) SetPoolId(v string) *AsyncUpdateBody {
	s.PoolId = &v
	return s
}

func (s *AsyncUpdateBody) SetSessionPersistence(v AsyncUpdateBodySessionPersistenceEnum) *AsyncUpdateBody {
	s.SessionPersistence = &v
	return s
}

func (s *AsyncUpdateBody) SetHealthMonitorUpdateReq(v *AsyncUpdateRequestHealthMonitorUpdateReq) *AsyncUpdateBody {
	s.HealthMonitorUpdateReq = v
	return s
}

func (s *AsyncUpdateBody) SetCookieName(v string) *AsyncUpdateBody {
	s.CookieName = &v
	return s
}

func (s *AsyncUpdateBody) SetPoolName(v string) *AsyncUpdateBody {
	s.PoolName = &v
	return s
}


type AsyncUpdateBodyBuilder struct {
	s *AsyncUpdateBody
}

func NewAsyncUpdateBodyBuilder() *AsyncUpdateBodyBuilder {
	s := &AsyncUpdateBody{}
	b := &AsyncUpdateBodyBuilder{s: s}
	return b
}

func (b *AsyncUpdateBodyBuilder) LbAlgorithm(v AsyncUpdateBodyLbAlgorithmEnum) *AsyncUpdateBodyBuilder {
	b.s.LbAlgorithm = &v
	return b
}

func (b *AsyncUpdateBodyBuilder) PoolId(v string) *AsyncUpdateBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncUpdateBodyBuilder) SessionPersistence(v AsyncUpdateBodySessionPersistenceEnum) *AsyncUpdateBodyBuilder {
	b.s.SessionPersistence = &v
	return b
}

func (b *AsyncUpdateBodyBuilder) HealthMonitorUpdateReq(v *AsyncUpdateRequestHealthMonitorUpdateReq) *AsyncUpdateBodyBuilder {
	b.s.HealthMonitorUpdateReq = v
	return b
}

func (b *AsyncUpdateBodyBuilder) CookieName(v string) *AsyncUpdateBodyBuilder {
	b.s.CookieName = &v
	return b
}

func (b *AsyncUpdateBodyBuilder) PoolName(v string) *AsyncUpdateBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *AsyncUpdateBodyBuilder) Build() *AsyncUpdateBody {
	return b.s
}


