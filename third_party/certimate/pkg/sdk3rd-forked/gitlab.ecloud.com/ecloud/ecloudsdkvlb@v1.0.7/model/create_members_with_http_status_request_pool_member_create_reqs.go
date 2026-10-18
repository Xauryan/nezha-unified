// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum string

// List of ProductType
const (
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumNormal CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "NORMAL"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumDeCloud CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "DE_CLOUD"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumAutoscaling CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "AUTOSCALING"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumVo CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "VO"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumCdn CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "CDN"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumPaasMaster CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "PAAS_MASTER"
    CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnumPaasSlave CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum = "PAAS_SLAVE"
)

type CreateMembersWithHttpStatusRequestPoolMemberCreateReqs struct {

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
    // 业务主机类型，1：云主机，3：裸金属
	Type *int32 `json:"type,omitempty"`
    // 业务主机由何种类型产品创建，客户订购默认为NORMAL类型 取值如下： NORMAL：普通创建 DE_CLOUD：专属云创建 AUTOSCALING：弹性伸缩创建 VO：VO模板编配创建 CDN：CDN云主机创建 PAAS_MASTER：PAAS容器控制节点云主机创建 PAAS_SLAVE：PAAS容器业务节点云主机创建 VCPE：VCPE云主机创建 EMR：EMR产品云主机创建 LOGAUDIT：日志审计产品云主机创建 MSE：MSE产品云主机创建
	ProductType *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum `json:"productType,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
}

func (s CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) String() string {
	return utils.Beautify(s)
}

func (s CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) GoString() string {
	return s.String()
}

func (s CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetSubnetId(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.SubnetId = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetPortType(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.PortType = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetPort(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Port = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetIp(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Ip = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetDescription(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Description = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetWeight(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Weight = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetVmHostId(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.VmHostId = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetType(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Type = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetProductType(v CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.ProductType = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetVaz(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.Vaz = &v
	return s
}

func (s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs) SetIsMultiAz(v bool) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	s.IsMultiAz = &v
	return s
}


type CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder struct {
	s *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs
}

func NewCreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder() *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	s := &CreateMembersWithHttpStatusRequestPoolMemberCreateReqs{}
	b := &CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder{s: s}
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) SubnetId(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) PortType(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.PortType = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Port(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Port = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Ip(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Ip = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Description(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Weight(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Weight = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) VmHostId(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.VmHostId = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Type(v int32) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Type = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) ProductType(v CreateMembersWithHttpStatusRequestPoolMemberCreateReqsProductTypeEnum) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Vaz(v string) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.Vaz = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) IsMultiAz(v bool) *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *CreateMembersWithHttpStatusRequestPoolMemberCreateReqsBuilder) Build() *CreateMembersWithHttpStatusRequestPoolMemberCreateReqs {
	return b.s
}


