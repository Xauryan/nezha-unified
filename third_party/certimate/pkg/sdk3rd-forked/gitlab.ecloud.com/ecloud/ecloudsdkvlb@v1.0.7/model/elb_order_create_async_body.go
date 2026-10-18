// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ElbOrderCreateAsyncBodyChargePeriodEnumEnum string

// List of ChargePeriodEnum
const (
    ElbOrderCreateAsyncBodyChargePeriodEnumEnumHour ElbOrderCreateAsyncBodyChargePeriodEnumEnum = "hour"
    ElbOrderCreateAsyncBodyChargePeriodEnumEnumMonth ElbOrderCreateAsyncBodyChargePeriodEnumEnum = "month"
    ElbOrderCreateAsyncBodyChargePeriodEnumEnumYear ElbOrderCreateAsyncBodyChargePeriodEnumEnum = "year"
)

type ElbOrderCreateAsyncBody struct {
    position.Body
    // 订购时长, 预付费订购且计费周期为month或者year时必填，后付费订购且计费周期为year时必填。 计费周期为month时，取值为1~12月，计费周期为year时，取值为[年数*12月]，最多5年
	Duration *int32 `json:"duration,omitempty"`
    // 弹性负载均衡的规格。1：普通型，2：优享型I，3：优享型II，4：高端型I，5：高端型II，6：旗舰型， 7：增强型， 10：简约型，21：性能保障型（选取时必填bandwidth参数）， 30：LCU规格，100：NFV托管平台
	Flavor *int32 `json:"flavor,omitempty"`
    // 弹性负载均衡名称，5~64位的英文、数字、下划线、中划线等的组合
	LoadBalanceName *string `json:"loadBalanceName,omitempty"`
    // 弹性负载均衡的子网资源ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 计费周期: 按时, 按月, 按年
	ChargePeriodEnum *ElbOrderCreateAsyncBodyChargePeriodEnumEnum `json:"chargePeriodEnum,omitempty"`
    // 带宽，仅订购性能保障型实例时可输入生效，单位：Mbps（flavro选择21性能保证型时，该字段为必填）
	Bandwidth *int32 `json:"bandwidth,omitempty"`
    // 指定内网IP地址创建弹性负载均衡实例
	IpAddress *string `json:"ipAddress,omitempty"`
    // 需要绑定的独享设备IP地址，多个地址以“,”分隔例如：1.2.3.4,1.2.3.5
	NodeIp *string `json:"nodeIp,omitempty"`
    // 规格为NFV托管平台订购时选择的云主机核数
	ServerCores *int32 `json:"serverCores,omitempty"`
    // 绑定标签
	Tags *[]ElbOrderCreateAsyncRequestTags `json:"tags,omitempty"`
}

func (s ElbOrderCreateAsyncBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncBody) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncBody) SetDuration(v int32) *ElbOrderCreateAsyncBody {
	s.Duration = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetFlavor(v int32) *ElbOrderCreateAsyncBody {
	s.Flavor = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetLoadBalanceName(v string) *ElbOrderCreateAsyncBody {
	s.LoadBalanceName = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetSubnetId(v string) *ElbOrderCreateAsyncBody {
	s.SubnetId = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetChargePeriodEnum(v ElbOrderCreateAsyncBodyChargePeriodEnumEnum) *ElbOrderCreateAsyncBody {
	s.ChargePeriodEnum = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetBandwidth(v int32) *ElbOrderCreateAsyncBody {
	s.Bandwidth = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetIpAddress(v string) *ElbOrderCreateAsyncBody {
	s.IpAddress = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetNodeIp(v string) *ElbOrderCreateAsyncBody {
	s.NodeIp = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetServerCores(v int32) *ElbOrderCreateAsyncBody {
	s.ServerCores = &v
	return s
}

func (s *ElbOrderCreateAsyncBody) SetTags(v []ElbOrderCreateAsyncRequestTags) *ElbOrderCreateAsyncBody {
	s.Tags = &v
	return s
}


type ElbOrderCreateAsyncBodyBuilder struct {
	s *ElbOrderCreateAsyncBody
}

func NewElbOrderCreateAsyncBodyBuilder() *ElbOrderCreateAsyncBodyBuilder {
	s := &ElbOrderCreateAsyncBody{}
	b := &ElbOrderCreateAsyncBodyBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) Duration(v int32) *ElbOrderCreateAsyncBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) Flavor(v int32) *ElbOrderCreateAsyncBodyBuilder {
	b.s.Flavor = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) LoadBalanceName(v string) *ElbOrderCreateAsyncBodyBuilder {
	b.s.LoadBalanceName = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) SubnetId(v string) *ElbOrderCreateAsyncBodyBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) ChargePeriodEnum(v ElbOrderCreateAsyncBodyChargePeriodEnumEnum) *ElbOrderCreateAsyncBodyBuilder {
	b.s.ChargePeriodEnum = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) Bandwidth(v int32) *ElbOrderCreateAsyncBodyBuilder {
	b.s.Bandwidth = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) IpAddress(v string) *ElbOrderCreateAsyncBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) NodeIp(v string) *ElbOrderCreateAsyncBodyBuilder {
	b.s.NodeIp = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) ServerCores(v int32) *ElbOrderCreateAsyncBodyBuilder {
	b.s.ServerCores = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) Tags(v []ElbOrderCreateAsyncRequestTags) *ElbOrderCreateAsyncBodyBuilder {
	b.s.Tags = &v
	return b
}

func (b *ElbOrderCreateAsyncBodyBuilder) Build() *ElbOrderCreateAsyncBody {
	return b.s
}


