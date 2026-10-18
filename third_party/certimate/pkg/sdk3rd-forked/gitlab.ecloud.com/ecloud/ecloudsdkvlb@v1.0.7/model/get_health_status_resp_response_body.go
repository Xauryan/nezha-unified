// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetHealthStatusRespResponseBody struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 后端服务器关联的业务主机的名称
	VmName *string `json:"vmName,omitempty"`
    // 是否删除
	IsDelete *int32 `json:"isDelete,omitempty"`
    // 创建人，用户ID，例如：CICD-U-xxxxxx
	Proposer *string `json:"proposer,omitempty"`
    // 业务主机的私有IP
	Ip *string `json:"ip,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 业务主机负载均衡的权重，取值范围：1-100（取值越大，在所有业务主机中占的权重越高，承担的负载均衡的比例越大）
	Weight *int32 `json:"weight,omitempty"`
    // 业务主机类型
	Type *int32 `json:"type,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 主节点状态
	AdminUpState *bool `json:"adminUpState,omitempty"`
    // 后端业务主机的底层状态，取值如下： up：正常 down：异常
	HealthStatus *string `json:"healthStatus,omitempty"`
    // 业务主机用来被监听的端口
	Port *int32 `json:"port,omitempty"`
    // 后端服务器绑定的后端服务器组的ID
	PoolId *string `json:"poolId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 后端服务器ID
	Id *string `json:"id,omitempty"`
    // 后端服务器关联业务主机的ID
	VmHostId *string `json:"vmHostId,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 业务主机的底层状态
	Status *string `json:"status,omitempty"`
}

func (s GetHealthStatusRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetHealthStatusRespResponseBody) GoString() string {
	return s.String()
}

func (s GetHealthStatusRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetHealthStatusRespResponseBody) SetSubnetId(v string) *GetHealthStatusRespResponseBody {
	s.SubnetId = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetVmName(v string) *GetHealthStatusRespResponseBody {
	s.VmName = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetIsDelete(v int32) *GetHealthStatusRespResponseBody {
	s.IsDelete = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetProposer(v string) *GetHealthStatusRespResponseBody {
	s.Proposer = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetIp(v string) *GetHealthStatusRespResponseBody {
	s.Ip = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetDescription(v string) *GetHealthStatusRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetWeight(v int32) *GetHealthStatusRespResponseBody {
	s.Weight = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetType(v int32) *GetHealthStatusRespResponseBody {
	s.Type = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetIsMultiAz(v bool) *GetHealthStatusRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetAdminUpState(v bool) *GetHealthStatusRespResponseBody {
	s.AdminUpState = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetHealthStatus(v string) *GetHealthStatusRespResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetPort(v int32) *GetHealthStatusRespResponseBody {
	s.Port = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetPoolId(v string) *GetHealthStatusRespResponseBody {
	s.PoolId = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetCreatedTime(v string) *GetHealthStatusRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetId(v string) *GetHealthStatusRespResponseBody {
	s.Id = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetVmHostId(v string) *GetHealthStatusRespResponseBody {
	s.VmHostId = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetRegion(v string) *GetHealthStatusRespResponseBody {
	s.Region = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetMultiAzUuid(v string) *GetHealthStatusRespResponseBody {
	s.MultiAzUuid = &v
	return s
}

func (s *GetHealthStatusRespResponseBody) SetStatus(v string) *GetHealthStatusRespResponseBody {
	s.Status = &v
	return s
}


type GetHealthStatusRespResponseBodyBuilder struct {
	s *GetHealthStatusRespResponseBody
}

func NewGetHealthStatusRespResponseBodyBuilder() *GetHealthStatusRespResponseBodyBuilder {
	s := &GetHealthStatusRespResponseBody{}
	b := &GetHealthStatusRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) SubnetId(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) VmName(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.VmName = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) IsDelete(v int32) *GetHealthStatusRespResponseBodyBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Proposer(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Proposer = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Ip(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Ip = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Description(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Weight(v int32) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Weight = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Type(v int32) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) IsMultiAz(v bool) *GetHealthStatusRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) AdminUpState(v bool) *GetHealthStatusRespResponseBodyBuilder {
	b.s.AdminUpState = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) HealthStatus(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.HealthStatus = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Port(v int32) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Port = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) PoolId(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) CreatedTime(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Id(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) VmHostId(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.VmHostId = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Region(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) MultiAzUuid(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Status(v string) *GetHealthStatusRespResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetHealthStatusRespResponseBodyBuilder) Build() *GetHealthStatusRespResponseBody {
	return b.s
}


