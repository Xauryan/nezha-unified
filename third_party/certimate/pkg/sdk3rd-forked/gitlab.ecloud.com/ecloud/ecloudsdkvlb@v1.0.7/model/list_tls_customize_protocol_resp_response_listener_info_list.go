// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListTLSCustomizeProtocolRespResponseListenerInfoList struct {

    // TLS安全策略绑定监听器端口协议
	Protocol *string `json:"protocol,omitempty"`
    // TLS安全策略绑定监听器名称
	Name *string `json:"name,omitempty"`
    // TLS安全策略绑定监听器ID
	Id *string `json:"id,omitempty"`
    // TLS安全策略绑定监听器端口
	ProtocolPort *int32 `json:"protocolPort,omitempty"`
}

func (s ListTLSCustomizeProtocolRespResponseListenerInfoList) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespResponseListenerInfoList) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespResponseListenerInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespResponseListenerInfoList) SetProtocol(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Protocol = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseListenerInfoList) SetName(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Name = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseListenerInfoList) SetId(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoList {
	s.Id = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseListenerInfoList) SetProtocolPort(v int32) *ListTLSCustomizeProtocolRespResponseListenerInfoList {
	s.ProtocolPort = &v
	return s
}


type ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder struct {
	s *ListTLSCustomizeProtocolRespResponseListenerInfoList
}

func NewListTLSCustomizeProtocolRespResponseListenerInfoListBuilder() *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	s := &ListTLSCustomizeProtocolRespResponseListenerInfoList{}
	b := &ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Protocol(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Protocol = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Name(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Name = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Id(v string) *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.Id = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder) ProtocolPort(v int32) *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder {
	b.s.ProtocolPort = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseListenerInfoListBuilder) Build() *ListTLSCustomizeProtocolRespResponseListenerInfoList {
	return b.s
}


