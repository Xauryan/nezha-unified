// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum string

// List of TlsStatus
const (
    ListTLSCustomizeProtocolRespResponseContentTlsStatusEnumNormal ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum = "NORMAL"
    ListTLSCustomizeProtocolRespResponseContentTlsStatusEnumError ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum = "ERROR"
    ListTLSCustomizeProtocolRespResponseContentTlsStatusEnumRunning ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum = "RUNNING"
)

type ListTLSCustomizeProtocolRespResponseContent struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 绑定监听器个数
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 是否删除
	IsDelete *bool `json:"isDelete,omitempty"`
    // 绑定监听列表
	ListenerInfoList *[]ListTLSCustomizeProtocolRespResponseListenerInfoList `json:"listenerInfoList,omitempty"`
    // TLS安全策略状态
	TlsStatus *ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum `json:"tlsStatus,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
    // 绑定监听信息
	ListenerInfos *string `json:"listenerInfos,omitempty"`
    // TLS安全策略协议
	TlsProtocols *string `json:"tlsProtocols,omitempty"`
    // 是否是默认策略
	IsDefault *bool `json:"isDefault,omitempty"`
    // 创建人
	CreatedBy *string `json:"createdBy,omitempty"`
    // 加密套件
	CipherSuites *string `json:"cipherSuites,omitempty"`
    // 客户ID
	CustomerId *string `json:"customerId,omitempty"`
    // TLS安全策略名称
	Name *string `json:"name,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 修改人
	ModifiedBy *string `json:"modifiedBy,omitempty"`
    // TLS安全测率ID
	Id *int64 `json:"id,omitempty"`
}

func (s ListTLSCustomizeProtocolRespResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespResponseContent) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetModifiedTime(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetBindedListeners(v int32) *ListTLSCustomizeProtocolRespResponseContent {
	s.BindedListeners = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetIsDelete(v bool) *ListTLSCustomizeProtocolRespResponseContent {
	s.IsDelete = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetListenerInfoList(v []ListTLSCustomizeProtocolRespResponseListenerInfoList) *ListTLSCustomizeProtocolRespResponseContent {
	s.ListenerInfoList = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetTlsStatus(v ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum) *ListTLSCustomizeProtocolRespResponseContent {
	s.TlsStatus = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetVaz(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.Vaz = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetListenerInfos(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.ListenerInfos = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetTlsProtocols(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.TlsProtocols = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetIsDefault(v bool) *ListTLSCustomizeProtocolRespResponseContent {
	s.IsDefault = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetCreatedBy(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.CreatedBy = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetCipherSuites(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.CipherSuites = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetCustomerId(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.CustomerId = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetName(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.Name = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetCreatedTime(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetModifiedBy(v string) *ListTLSCustomizeProtocolRespResponseContent {
	s.ModifiedBy = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseContent) SetId(v int64) *ListTLSCustomizeProtocolRespResponseContent {
	s.Id = &v
	return s
}


type ListTLSCustomizeProtocolRespResponseContentBuilder struct {
	s *ListTLSCustomizeProtocolRespResponseContent
}

func NewListTLSCustomizeProtocolRespResponseContentBuilder() *ListTLSCustomizeProtocolRespResponseContentBuilder {
	s := &ListTLSCustomizeProtocolRespResponseContent{}
	b := &ListTLSCustomizeProtocolRespResponseContentBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) ModifiedTime(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) BindedListeners(v int32) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) IsDelete(v bool) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) ListenerInfoList(v []ListTLSCustomizeProtocolRespResponseListenerInfoList) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.ListenerInfoList = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) TlsStatus(v ListTLSCustomizeProtocolRespResponseContentTlsStatusEnum) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.TlsStatus = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) Vaz(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) ListenerInfos(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.ListenerInfos = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) TlsProtocols(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) IsDefault(v bool) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.IsDefault = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) CreatedBy(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) CipherSuites(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) CustomerId(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.CustomerId = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) Name(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) CreatedTime(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) ModifiedBy(v string) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.ModifiedBy = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) Id(v int64) *ListTLSCustomizeProtocolRespResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseContentBuilder) Build() *ListTLSCustomizeProtocolRespResponseContent {
	return b.s
}


