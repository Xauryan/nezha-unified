// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalancePoolMemberRespResponseBody struct {

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
    // 业务主机类型，取值如下： 0：IP类型，1：云主机，3：裸金属
	Type *int32 `json:"type,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 后端服务器的底层状态，取值如下： up：正常 down：异常
	HealthStatus *string `json:"healthStatus,omitempty"`
    // 业务主机用来被监听的端口
	Port *int32 `json:"port,omitempty"`
    // 主机绑定的后端服务器组的ID
	PoolId *string `json:"poolId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 后端服务器ID
	Id *string `json:"id,omitempty"`
    // 后端服务器关联业务主机的ID
	VmHostId *string `json:"vmHostId,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 业务主机的底层状态
	Status *string `json:"status,omitempty"`
}

func (s GetLoadBalancePoolMemberRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalancePoolMemberRespResponseBody) GoString() string {
	return s.String()
}

func (s GetLoadBalancePoolMemberRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetSubnetId(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.SubnetId = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetVmName(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.VmName = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetIsDelete(v int32) *GetLoadBalancePoolMemberRespResponseBody {
	s.IsDelete = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetProposer(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.Proposer = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetIp(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.Ip = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetDescription(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetWeight(v int32) *GetLoadBalancePoolMemberRespResponseBody {
	s.Weight = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetType(v int32) *GetLoadBalancePoolMemberRespResponseBody {
	s.Type = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetIsMultiAz(v bool) *GetLoadBalancePoolMemberRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetHealthStatus(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetPort(v int32) *GetLoadBalancePoolMemberRespResponseBody {
	s.Port = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetPoolId(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.PoolId = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetCreatedTime(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetId(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.Id = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetVmHostId(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.VmHostId = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetMultiAzUuid(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.MultiAzUuid = &v
	return s
}

func (s *GetLoadBalancePoolMemberRespResponseBody) SetStatus(v string) *GetLoadBalancePoolMemberRespResponseBody {
	s.Status = &v
	return s
}


type GetLoadBalancePoolMemberRespResponseBodyBuilder struct {
	s *GetLoadBalancePoolMemberRespResponseBody
}

func NewGetLoadBalancePoolMemberRespResponseBodyBuilder() *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	s := &GetLoadBalancePoolMemberRespResponseBody{}
	b := &GetLoadBalancePoolMemberRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) SubnetId(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) VmName(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.VmName = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) IsDelete(v int32) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Proposer(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Proposer = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Ip(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Ip = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Description(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Weight(v int32) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Weight = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Type(v int32) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) IsMultiAz(v bool) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) HealthStatus(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.HealthStatus = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Port(v int32) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Port = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) PoolId(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) CreatedTime(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Id(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) VmHostId(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.VmHostId = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) MultiAzUuid(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Status(v string) *GetLoadBalancePoolMemberRespResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetLoadBalancePoolMemberRespResponseBodyBuilder) Build() *GetLoadBalancePoolMemberRespResponseBody {
	return b.s
}


