// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum string

// List of ProductType
const (
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumNormal CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "NORMAL"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumDeCloud CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "DE_CLOUD"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumAutoscaling CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "AUTOSCALING"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumVo CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "VO"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumCdn CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "CDN"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumPaasMaster CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "PAAS_MASTER"
    CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnumPaasSlave CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum = "PAAS_SLAVE"
)

type CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 网卡类型，主机--vm，物理机--ironic
	PortType *string `json:"portType,omitempty"`
    // 业务主机用来被监听的端口
	Port *int32 `json:"port,omitempty"`
    // 业务主机的私有IP
	Ip *string `json:"ip,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 业务主机负载均衡的权重，取值范围：1-100。取值越大，在所有业务主机中占的权重越高，承担的负载均衡的比例越大
	Weight *int32 `json:"weight,omitempty"`
    // 主机ID，最多可添加20个
	VmHostId *string `json:"vmHostId,omitempty"`
    // 业务主机类型, 1：云主机，3：裸金属
	Type *int32 `json:"type,omitempty"`
    // 业务主机由何种类型产品创建，客户订购默认为NORMAL类型 取值如下： NORMAL：普通创建 DE_CLOUD：专属云创建 AUTOSCALING：弹性伸缩创建 VO：VO模板编配创建 CDN：CDN云主机创建 PAAS_MASTER：PAAS容器控制节点云主机创建 PAAS_SLAVE：PAAS容器业务节点云主机创建 VCPE：VCPE云主机创建 EMR：EMR产品云主机创建 LOGAUDIT：日志审计产品云主机创建 MSE：MSE产品云主机创建
	ProductType *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum `json:"productType,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
}

func (s CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) String() string {
	return utils.Beautify(s)
}

func (s CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) GoString() string {
	return s.String()
}

func (s CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetSubnetId(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.SubnetId = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetPortType(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.PortType = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetPort(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Port = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetIp(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Ip = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetDescription(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetWeight(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Weight = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetVmHostId(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.VmHostId = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetType(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Type = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetProductType(v CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.ProductType = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetVaz(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.Vaz = &v
	return s
}

func (s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs) SetIsMultiAz(v bool) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	s.IsMultiAz = &v
	return s
}


type CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder struct {
	s *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs
}

func NewCreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder() *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	s := &CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs{}
	b := &CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder{s: s}
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) SubnetId(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) PortType(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.PortType = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Port(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Port = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Ip(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Ip = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Description(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Weight(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Weight = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) VmHostId(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.VmHostId = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Type(v int32) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Type = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) ProductType(v CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsProductTypeEnum) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Vaz(v string) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.Vaz = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) IsMultiAz(v bool) *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqsBuilder) Build() *CreateLoadBalancePoolMembersRequestPoolMemberCreateReqs {
	return b.s
}


