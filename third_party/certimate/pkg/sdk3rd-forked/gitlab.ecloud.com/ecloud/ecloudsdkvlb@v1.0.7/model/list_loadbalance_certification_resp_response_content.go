// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceCertificationRespResponseContentContainerTypeEnum string

// List of ContainerType
const (
    ListLoadbalanceCertificationRespResponseContentContainerTypeEnumGeneric ListLoadbalanceCertificationRespResponseContentContainerTypeEnum = "generic"
    ListLoadbalanceCertificationRespResponseContentContainerTypeEnumRsa ListLoadbalanceCertificationRespResponseContentContainerTypeEnum = "rsa"
    ListLoadbalanceCertificationRespResponseContentContainerTypeEnumCertificate ListLoadbalanceCertificationRespResponseContentContainerTypeEnum = "certificate"
)
type ListLoadbalanceCertificationRespResponseContentTypeEnum string

// List of Type
const (
    ListLoadbalanceCertificationRespResponseContentTypeEnumServer ListLoadbalanceCertificationRespResponseContentTypeEnum = "SERVER"
    ListLoadbalanceCertificationRespResponseContentTypeEnumCa ListLoadbalanceCertificationRespResponseContentTypeEnum = "CA"
    ListLoadbalanceCertificationRespResponseContentTypeEnumSni ListLoadbalanceCertificationRespResponseContentTypeEnum = "SNI"
)
type ListLoadbalanceCertificationRespResponseContentStatusEnum string

// List of Status
const (
    ListLoadbalanceCertificationRespResponseContentStatusEnumPending ListLoadbalanceCertificationRespResponseContentStatusEnum = "PENDING"
    ListLoadbalanceCertificationRespResponseContentStatusEnumActive ListLoadbalanceCertificationRespResponseContentStatusEnum = "ACTIVE"
    ListLoadbalanceCertificationRespResponseContentStatusEnumError ListLoadbalanceCertificationRespResponseContentStatusEnum = "ERROR"
)

type ListLoadbalanceCertificationRespResponseContent struct {

    // 证书ID
	ContainerUuid *string `json:"containerUuid,omitempty"`
    // 证书绑定监听器个数
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 证书所在容器的类型，默认均为certificate
	ContainerType *ListLoadbalanceCertificationRespResponseContentContainerTypeEnum `json:"containerType,omitempty"`
    // 证书描述
	Description *string `json:"description,omitempty"`
    // 证书类型
	Type *ListLoadbalanceCertificationRespResponseContentTypeEnum `json:"type,omitempty"`
    // 证书是否被删除，true：已删除，false：未删除
	Delete *bool `json:"delete,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
    // 证书所在容器的访问地址
	ContainerRef *string `json:"containerRef,omitempty"`
    // 证书过期时间
	ExpirationTime *string `json:"expirationTime,omitempty"`
    // 证书名称
	Name *string `json:"name,omitempty"`
    // 证书创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 证书ID
	Id *string `json:"id,omitempty"`
    // 可用区名称
	Region *string `json:"region,omitempty"`
    // 证书所在容器的状态，暂未使用
	Status *ListLoadbalanceCertificationRespResponseContentStatusEnum `json:"status,omitempty"`
}

func (s ListLoadbalanceCertificationRespResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceCertificationRespResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadbalanceCertificationRespResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetContainerUuid(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.ContainerUuid = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetBindedListeners(v int32) *ListLoadbalanceCertificationRespResponseContent {
	s.BindedListeners = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetContainerType(v ListLoadbalanceCertificationRespResponseContentContainerTypeEnum) *ListLoadbalanceCertificationRespResponseContent {
	s.ContainerType = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetDescription(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetType(v ListLoadbalanceCertificationRespResponseContentTypeEnum) *ListLoadbalanceCertificationRespResponseContent {
	s.Type = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetDelete(v bool) *ListLoadbalanceCertificationRespResponseContent {
	s.Delete = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetVaz(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.Vaz = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetContainerRef(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.ContainerRef = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetExpirationTime(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.ExpirationTime = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetName(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetCreatedTime(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetId(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetRegion(v string) *ListLoadbalanceCertificationRespResponseContent {
	s.Region = &v
	return s
}

func (s *ListLoadbalanceCertificationRespResponseContent) SetStatus(v ListLoadbalanceCertificationRespResponseContentStatusEnum) *ListLoadbalanceCertificationRespResponseContent {
	s.Status = &v
	return s
}


type ListLoadbalanceCertificationRespResponseContentBuilder struct {
	s *ListLoadbalanceCertificationRespResponseContent
}

func NewListLoadbalanceCertificationRespResponseContentBuilder() *ListLoadbalanceCertificationRespResponseContentBuilder {
	s := &ListLoadbalanceCertificationRespResponseContent{}
	b := &ListLoadbalanceCertificationRespResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) ContainerUuid(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) BindedListeners(v int32) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) ContainerType(v ListLoadbalanceCertificationRespResponseContentContainerTypeEnum) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.ContainerType = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Description(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Type(v ListLoadbalanceCertificationRespResponseContentTypeEnum) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Type = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Delete(v bool) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Delete = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Vaz(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) ContainerRef(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.ContainerRef = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) ExpirationTime(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.ExpirationTime = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Name(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) CreatedTime(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Id(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Region(v string) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Region = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Status(v ListLoadbalanceCertificationRespResponseContentStatusEnum) *ListLoadbalanceCertificationRespResponseContentBuilder {
	b.s.Status = &v
	return b
}

func (b *ListLoadbalanceCertificationRespResponseContentBuilder) Build() *ListLoadbalanceCertificationRespResponseContent {
	return b.s
}


