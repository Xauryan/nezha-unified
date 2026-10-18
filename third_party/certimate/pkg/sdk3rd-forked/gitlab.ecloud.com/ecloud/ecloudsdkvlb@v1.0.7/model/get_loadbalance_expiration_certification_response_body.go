// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum string

// List of ContainerType
const (
    GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnumGeneric GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum = "generic"
    GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnumRsa GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum = "rsa"
    GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnumCertificate GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum = "certificate"
)
type GetLoadbalanceExpirationCertificationResponseBodyTypeEnum string

// List of Type
const (
    GetLoadbalanceExpirationCertificationResponseBodyTypeEnumServer GetLoadbalanceExpirationCertificationResponseBodyTypeEnum = "SERVER"
    GetLoadbalanceExpirationCertificationResponseBodyTypeEnumCa GetLoadbalanceExpirationCertificationResponseBodyTypeEnum = "CA"
    GetLoadbalanceExpirationCertificationResponseBodyTypeEnumSni GetLoadbalanceExpirationCertificationResponseBodyTypeEnum = "SNI"
)
type GetLoadbalanceExpirationCertificationResponseBodyStatusEnum string

// List of Status
const (
    GetLoadbalanceExpirationCertificationResponseBodyStatusEnumPending GetLoadbalanceExpirationCertificationResponseBodyStatusEnum = "PENDING"
    GetLoadbalanceExpirationCertificationResponseBodyStatusEnumActive GetLoadbalanceExpirationCertificationResponseBodyStatusEnum = "ACTIVE"
    GetLoadbalanceExpirationCertificationResponseBodyStatusEnumError GetLoadbalanceExpirationCertificationResponseBodyStatusEnum = "ERROR"
)

type GetLoadbalanceExpirationCertificationResponseBody struct {

    // 证书所在容器的ID
	ContainerUuid *string `json:"containerUuid,omitempty"`
    // 证书绑定监听器数量
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 证书所在容器的类型，默认均为certificate
	ContainerType *GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum `json:"containerType,omitempty"`
    // 证书描述
	Description *string `json:"description,omitempty"`
    // 证书类型
	Type *GetLoadbalanceExpirationCertificationResponseBodyTypeEnum `json:"type,omitempty"`
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
    // 边缘云可用区
	Region *string `json:"region,omitempty"`
    // 容器状态
	Status *GetLoadbalanceExpirationCertificationResponseBodyStatusEnum `json:"status,omitempty"`
}

func (s GetLoadbalanceExpirationCertificationResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetLoadbalanceExpirationCertificationResponseBody) GoString() string {
	return s.String()
}

func (s GetLoadbalanceExpirationCertificationResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetContainerUuid(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.ContainerUuid = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetBindedListeners(v int32) *GetLoadbalanceExpirationCertificationResponseBody {
	s.BindedListeners = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetContainerType(v GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum) *GetLoadbalanceExpirationCertificationResponseBody {
	s.ContainerType = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetDescription(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Description = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetType(v GetLoadbalanceExpirationCertificationResponseBodyTypeEnum) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Type = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetDelete(v bool) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Delete = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetVaz(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Vaz = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetContainerRef(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.ContainerRef = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetExpirationTime(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetName(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Name = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetCreatedTime(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetId(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Id = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetRegion(v string) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Region = &v
	return s
}

func (s *GetLoadbalanceExpirationCertificationResponseBody) SetStatus(v GetLoadbalanceExpirationCertificationResponseBodyStatusEnum) *GetLoadbalanceExpirationCertificationResponseBody {
	s.Status = &v
	return s
}


type GetLoadbalanceExpirationCertificationResponseBodyBuilder struct {
	s *GetLoadbalanceExpirationCertificationResponseBody
}

func NewGetLoadbalanceExpirationCertificationResponseBodyBuilder() *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	s := &GetLoadbalanceExpirationCertificationResponseBody{}
	b := &GetLoadbalanceExpirationCertificationResponseBodyBuilder{s: s}
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) ContainerUuid(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.ContainerUuid = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) BindedListeners(v int32) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) ContainerType(v GetLoadbalanceExpirationCertificationResponseBodyContainerTypeEnum) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.ContainerType = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Description(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Type(v GetLoadbalanceExpirationCertificationResponseBodyTypeEnum) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Delete(v bool) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Delete = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Vaz(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) ContainerRef(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.ContainerRef = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) ExpirationTime(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.ExpirationTime = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Name(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) CreatedTime(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Id(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Region(v string) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Status(v GetLoadbalanceExpirationCertificationResponseBodyStatusEnum) *GetLoadbalanceExpirationCertificationResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetLoadbalanceExpirationCertificationResponseBodyBuilder) Build() *GetLoadbalanceExpirationCertificationResponseBody {
	return b.s
}


