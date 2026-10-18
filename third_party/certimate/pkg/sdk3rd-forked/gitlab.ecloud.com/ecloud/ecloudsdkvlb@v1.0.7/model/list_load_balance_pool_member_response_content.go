// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadBalancePoolMemberResponseContent struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 后端服务器关联的业务主机的名称
	VmName *string `json:"vmName,omitempty"`
    // 是否删除
	IsDelete *int32 `json:"isDelete,omitempty"`
    // 创建人，用户ID，例如：CICD-U-xxxxxx
	Proposer *string `json:"proposer,omitempty"`
    // 后端服务器的私有IP
	Ip *string `json:"ip,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 业务主机负载均衡的权重，取值范围：1-100（取值越大，在所有业务主机中占的权重越高，承担的负载均衡的比例越大）
	Weight *int32 `json:"weight,omitempty"`
    // 业务主机类型，0：IP类型 1：云主机 3：裸金属
	Type *int32 `json:"type,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 主节点状态
	AdminUpState *bool `json:"adminUpState,omitempty"`
    // 主机的底层状态，取值如下： up：正常 down：异常
	HealthStatus *string `json:"healthStatus,omitempty"`
    // 后端服务器用来被监听的端口
	Port *int32 `json:"port,omitempty"`
    // 后端服务器绑定的后端服务器组的ID
	PoolId *string `json:"poolId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 后端服务器的ID
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

func (s ListLoadBalancePoolMemberResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadBalancePoolMemberResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadBalancePoolMemberResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadBalancePoolMemberResponseContent) SetSubnetId(v string) *ListLoadBalancePoolMemberResponseContent {
	s.SubnetId = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetVmName(v string) *ListLoadBalancePoolMemberResponseContent {
	s.VmName = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetIsDelete(v int32) *ListLoadBalancePoolMemberResponseContent {
	s.IsDelete = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetProposer(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Proposer = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetIp(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Ip = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetDescription(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetWeight(v int32) *ListLoadBalancePoolMemberResponseContent {
	s.Weight = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetType(v int32) *ListLoadBalancePoolMemberResponseContent {
	s.Type = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetIsMultiAz(v bool) *ListLoadBalancePoolMemberResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetAdminUpState(v bool) *ListLoadBalancePoolMemberResponseContent {
	s.AdminUpState = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetHealthStatus(v string) *ListLoadBalancePoolMemberResponseContent {
	s.HealthStatus = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetPort(v int32) *ListLoadBalancePoolMemberResponseContent {
	s.Port = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetPoolId(v string) *ListLoadBalancePoolMemberResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetCreatedTime(v string) *ListLoadBalancePoolMemberResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetId(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetVmHostId(v string) *ListLoadBalancePoolMemberResponseContent {
	s.VmHostId = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetRegion(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Region = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetMultiAzUuid(v string) *ListLoadBalancePoolMemberResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListLoadBalancePoolMemberResponseContent) SetStatus(v string) *ListLoadBalancePoolMemberResponseContent {
	s.Status = &v
	return s
}


type ListLoadBalancePoolMemberResponseContentBuilder struct {
	s *ListLoadBalancePoolMemberResponseContent
}

func NewListLoadBalancePoolMemberResponseContentBuilder() *ListLoadBalancePoolMemberResponseContentBuilder {
	s := &ListLoadBalancePoolMemberResponseContent{}
	b := &ListLoadBalancePoolMemberResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) SubnetId(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) VmName(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.VmName = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) IsDelete(v int32) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Proposer(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Ip(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Ip = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Description(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Weight(v int32) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Weight = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Type(v int32) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Type = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) IsMultiAz(v bool) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) AdminUpState(v bool) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.AdminUpState = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) HealthStatus(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.HealthStatus = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Port(v int32) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Port = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) PoolId(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) CreatedTime(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Id(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) VmHostId(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.VmHostId = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Region(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Region = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) MultiAzUuid(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Status(v string) *ListLoadBalancePoolMemberResponseContentBuilder {
	b.s.Status = &v
	return b
}

func (b *ListLoadBalancePoolMemberResponseContentBuilder) Build() *ListLoadBalancePoolMemberResponseContent {
	return b.s
}


