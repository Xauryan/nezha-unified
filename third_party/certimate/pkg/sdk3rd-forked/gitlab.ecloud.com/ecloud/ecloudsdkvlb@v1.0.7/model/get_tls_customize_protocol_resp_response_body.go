// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum string

// List of TlsStatus
const (
    GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnumNormal GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum = "NORMAL"
    GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnumError GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum = "ERROR"
    GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnumRunning GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum = "RUNNING"
)

type GetTLSCustomizeProtocolRespResponseBody struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // TLS安全策略绑定监听器数量
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 是否删除
	IsDelete *bool `json:"isDelete,omitempty"`
    // 绑定监听器列表
	ListenerInfoList *[]GetTLSCustomizeProtocolRespResponseListenerInfoList `json:"listenerInfoList,omitempty"`
    // TLS安全策略状态
	TlsStatus *GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum `json:"tlsStatus,omitempty"`
    // 是否为边缘云
	Vaz *string `json:"vaz,omitempty"`
    // 绑定监听信息
	ListenerInfos *string `json:"listenerInfos,omitempty"`
    // TLS安全策略协议
	TlsProtocols *string `json:"tlsProtocols,omitempty"`
    // 是否是默认策略
	IsDefault *bool `json:"isDefault,omitempty"`
    // 创建人
	CreatedBy *string `json:"createdBy,omitempty"`
    // TLS安全策略加密套件
	CipherSuites *string `json:"cipherSuites,omitempty"`
    // 客户ID
	CustomerId *string `json:"customerId,omitempty"`
    // TLS安全策略名称
	Name *string `json:"name,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 修改人
	ModifiedBy *string `json:"modifiedBy,omitempty"`
    // TLS安全策略ID
	Id *int64 `json:"id,omitempty"`
}

func (s GetTLSCustomizeProtocolRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetTLSCustomizeProtocolRespResponseBody) GoString() string {
	return s.String()
}

func (s GetTLSCustomizeProtocolRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetModifiedTime(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetBindedListeners(v int32) *GetTLSCustomizeProtocolRespResponseBody {
	s.BindedListeners = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetIsDelete(v bool) *GetTLSCustomizeProtocolRespResponseBody {
	s.IsDelete = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetListenerInfoList(v []GetTLSCustomizeProtocolRespResponseListenerInfoList) *GetTLSCustomizeProtocolRespResponseBody {
	s.ListenerInfoList = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetTlsStatus(v GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum) *GetTLSCustomizeProtocolRespResponseBody {
	s.TlsStatus = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetVaz(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.Vaz = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetListenerInfos(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.ListenerInfos = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetTlsProtocols(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.TlsProtocols = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetIsDefault(v bool) *GetTLSCustomizeProtocolRespResponseBody {
	s.IsDefault = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetCreatedBy(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetCipherSuites(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.CipherSuites = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetCustomerId(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.CustomerId = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetName(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.Name = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetCreatedTime(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetModifiedBy(v string) *GetTLSCustomizeProtocolRespResponseBody {
	s.ModifiedBy = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseBody) SetId(v int64) *GetTLSCustomizeProtocolRespResponseBody {
	s.Id = &v
	return s
}


type GetTLSCustomizeProtocolRespResponseBodyBuilder struct {
	s *GetTLSCustomizeProtocolRespResponseBody
}

func NewGetTLSCustomizeProtocolRespResponseBodyBuilder() *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	s := &GetTLSCustomizeProtocolRespResponseBody{}
	b := &GetTLSCustomizeProtocolRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) ModifiedTime(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) BindedListeners(v int32) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) IsDelete(v bool) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) ListenerInfoList(v []GetTLSCustomizeProtocolRespResponseListenerInfoList) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.ListenerInfoList = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) TlsStatus(v GetTLSCustomizeProtocolRespResponseBodyTlsStatusEnum) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.TlsStatus = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) Vaz(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) ListenerInfos(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.ListenerInfos = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) TlsProtocols(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.TlsProtocols = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) IsDefault(v bool) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.IsDefault = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) CreatedBy(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) CipherSuites(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.CipherSuites = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) CustomerId(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.CustomerId = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) Name(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) CreatedTime(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) ModifiedBy(v string) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.ModifiedBy = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) Id(v int64) *GetTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseBodyBuilder) Build() *GetTLSCustomizeProtocolRespResponseBody {
	return b.s
}


