// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTLSCustomizeProtocolRespResponseListenerInfoList struct {

    // 绑定监听器协议
	Protocol *string `json:"protocol,omitempty"`
    // 绑定监听器名称
	Name *string `json:"name,omitempty"`
    // 绑定监听器ID
	Id *string `json:"id,omitempty"`
    // 绑定监听器协议端口
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
}

func (s GetTLSCustomizeProtocolRespResponseListenerInfoList) String() string {
	return utils.Beautify(s)
}

func (s GetTLSCustomizeProtocolRespResponseListenerInfoList) GoString() string {
	return s.String()
}

func (s GetTLSCustomizeProtocolRespResponseListenerInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTLSCustomizeProtocolRespResponseListenerInfoList) SetProtocol(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Protocol = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseListenerInfoList) SetName(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Name = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseListenerInfoList) SetId(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Id = &v
	return s
}

func (s *GetTLSCustomizeProtocolRespResponseListenerInfoList) SetProtocolPort(v int32) *GetTLSCustomizeProtocolRespResponseListenerInfoList {
	s.ProtocolPort = &v
	return s
}


type GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder struct {
	s *GetTLSCustomizeProtocolRespResponseListenerInfoList
}

func NewGetTLSCustomizeProtocolRespResponseListenerInfoListBuilder() *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	s := &GetTLSCustomizeProtocolRespResponseListenerInfoList{}
	b := &GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder{s: s}
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Protocol(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Name(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Name = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Id(v string) *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Id = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder) ProtocolPort(v int32) *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *GetTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Build() *GetTLSCustomizeProtocolRespResponseListenerInfoList {
	return b.s
}


