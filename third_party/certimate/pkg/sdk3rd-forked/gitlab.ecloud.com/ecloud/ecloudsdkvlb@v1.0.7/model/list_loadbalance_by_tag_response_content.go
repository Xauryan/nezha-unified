// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceByTagResponseContentProviderEnum string

// List of Provider
const (
    ListLoadbalanceByTagResponseContentProviderEnumBrocade ListLoadbalanceByTagResponseContentProviderEnum = "brocade"
    ListLoadbalanceByTagResponseContentProviderEnumF5hardware ListLoadbalanceByTagResponseContentProviderEnum = "f5hardware"
    ListLoadbalanceByTagResponseContentProviderEnumNokia ListLoadbalanceByTagResponseContentProviderEnum = "nokia"
    ListLoadbalanceByTagResponseContentProviderEnumBcslb ListLoadbalanceByTagResponseContentProviderEnum = "bcslb"
    ListLoadbalanceByTagResponseContentProviderEnumF5networksS ListLoadbalanceByTagResponseContentProviderEnum = "f5networks_s"
    ListLoadbalanceByTagResponseContentProviderEnumNsb ListLoadbalanceByTagResponseContentProviderEnum = "nsb"
    ListLoadbalanceByTagResponseContentProviderEnumEslb ListLoadbalanceByTagResponseContentProviderEnum = "eslb"
    ListLoadbalanceByTagResponseContentProviderEnumRjhardf5 ListLoadbalanceByTagResponseContentProviderEnum = "rjhardf5"
    ListLoadbalanceByTagResponseContentProviderEnumVirtual ListLoadbalanceByTagResponseContentProviderEnum = "virtual"
    ListLoadbalanceByTagResponseContentProviderEnumH3clb ListLoadbalanceByTagResponseContentProviderEnum = "h3clb"
    ListLoadbalanceByTagResponseContentProviderEnumH3virtual ListLoadbalanceByTagResponseContentProviderEnum = "h3virtual"
    ListLoadbalanceByTagResponseContentProviderEnumAliyun ListLoadbalanceByTagResponseContentProviderEnum = "aliyun"
    ListLoadbalanceByTagResponseContentProviderEnumNfvslb ListLoadbalanceByTagResponseContentProviderEnum = "nfvslb"
)
type ListLoadbalanceByTagResponseContentProductTypeEnum string

// List of ProductType
const (
    ListLoadbalanceByTagResponseContentProductTypeEnumNormal ListLoadbalanceByTagResponseContentProductTypeEnum = "NORMAL"
    ListLoadbalanceByTagResponseContentProductTypeEnumDeCloud ListLoadbalanceByTagResponseContentProductTypeEnum = "DE_CLOUD"
    ListLoadbalanceByTagResponseContentProductTypeEnumAutoscaling ListLoadbalanceByTagResponseContentProductTypeEnum = "AUTOSCALING"
    ListLoadbalanceByTagResponseContentProductTypeEnumVo ListLoadbalanceByTagResponseContentProductTypeEnum = "VO"
    ListLoadbalanceByTagResponseContentProductTypeEnumCdn ListLoadbalanceByTagResponseContentProductTypeEnum = "CDN"
    ListLoadbalanceByTagResponseContentProductTypeEnumPaasMaster ListLoadbalanceByTagResponseContentProductTypeEnum = "PAAS_MASTER"
    ListLoadbalanceByTagResponseContentProductTypeEnumPaasSlave ListLoadbalanceByTagResponseContentProductTypeEnum = "PAAS_SLAVE"
    ListLoadbalanceByTagResponseContentProductTypeEnumVcpe ListLoadbalanceByTagResponseContentProductTypeEnum = "VCPE"
    ListLoadbalanceByTagResponseContentProductTypeEnumEmr ListLoadbalanceByTagResponseContentProductTypeEnum = "EMR"
    ListLoadbalanceByTagResponseContentProductTypeEnumLogaudit ListLoadbalanceByTagResponseContentProductTypeEnum = "LOGAUDIT"
    ListLoadbalanceByTagResponseContentProductTypeEnumMse ListLoadbalanceByTagResponseContentProductTypeEnum = "MSE"
    ListLoadbalanceByTagResponseContentProductTypeEnumVirtual ListLoadbalanceByTagResponseContentProductTypeEnum = "VIRTUAL"
    ListLoadbalanceByTagResponseContentProductTypeEnumBastion ListLoadbalanceByTagResponseContentProductTypeEnum = "BASTION"
    ListLoadbalanceByTagResponseContentProductTypeEnumNextGenerationFirewall ListLoadbalanceByTagResponseContentProductTypeEnum = "NEXT_GENERATION_FIREWALL"
)
type ListLoadbalanceByTagResponseContentEcStatusEnum string

// List of EcStatus
const (
    ListLoadbalanceByTagResponseContentEcStatusEnumActive ListLoadbalanceByTagResponseContentEcStatusEnum = "ACTIVE"
    ListLoadbalanceByTagResponseContentEcStatusEnumBuild ListLoadbalanceByTagResponseContentEcStatusEnum = "BUILD"
    ListLoadbalanceByTagResponseContentEcStatusEnumDown ListLoadbalanceByTagResponseContentEcStatusEnum = "DOWN"
    ListLoadbalanceByTagResponseContentEcStatusEnumError ListLoadbalanceByTagResponseContentEcStatusEnum = "ERROR"
    ListLoadbalanceByTagResponseContentEcStatusEnumPendingCreate ListLoadbalanceByTagResponseContentEcStatusEnum = "PENDING_CREATE"
    ListLoadbalanceByTagResponseContentEcStatusEnumPendingUpdate ListLoadbalanceByTagResponseContentEcStatusEnum = "PENDING_UPDATE"
    ListLoadbalanceByTagResponseContentEcStatusEnumPendingDelete ListLoadbalanceByTagResponseContentEcStatusEnum = "PENDING_DELETE"
    ListLoadbalanceByTagResponseContentEcStatusEnumMigrating ListLoadbalanceByTagResponseContentEcStatusEnum = "MIGRATING"
    ListLoadbalanceByTagResponseContentEcStatusEnumDeleted ListLoadbalanceByTagResponseContentEcStatusEnum = "DELETED"
    ListLoadbalanceByTagResponseContentEcStatusEnumUnrecognized ListLoadbalanceByTagResponseContentEcStatusEnum = "UNRECOGNIZED"
)
type ListLoadbalanceByTagResponseContentIpVersionEnum int32

// List of IpVersion
const (
    ListLoadbalanceByTagResponseContentIpVersionEnum4 ListLoadbalanceByTagResponseContentIpVersionEnum = 4
    ListLoadbalanceByTagResponseContentIpVersionEnum6 ListLoadbalanceByTagResponseContentIpVersionEnum = 6
)
type ListLoadbalanceByTagResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadbalanceByTagResponseContentOpStatusEnumNewBuilt ListLoadbalanceByTagResponseContentOpStatusEnum = "NEW_BUILT"
    ListLoadbalanceByTagResponseContentOpStatusEnumWaitToCreate ListLoadbalanceByTagResponseContentOpStatusEnum = "WAIT_TO_CREATE"
    ListLoadbalanceByTagResponseContentOpStatusEnumCreating ListLoadbalanceByTagResponseContentOpStatusEnum = "CREATING"
    ListLoadbalanceByTagResponseContentOpStatusEnumCreatedToSynchronize ListLoadbalanceByTagResponseContentOpStatusEnum = "CREATED_TO_SYNCHRONIZE"
    ListLoadbalanceByTagResponseContentOpStatusEnumCreated ListLoadbalanceByTagResponseContentOpStatusEnum = "CREATED"
    ListLoadbalanceByTagResponseContentOpStatusEnumCreateFailed ListLoadbalanceByTagResponseContentOpStatusEnum = "CREATE_FAILED"
    ListLoadbalanceByTagResponseContentOpStatusEnumFreezing ListLoadbalanceByTagResponseContentOpStatusEnum = "FREEZING"
    ListLoadbalanceByTagResponseContentOpStatusEnumFrozen ListLoadbalanceByTagResponseContentOpStatusEnum = "FROZEN"
    ListLoadbalanceByTagResponseContentOpStatusEnumFreezeFailed ListLoadbalanceByTagResponseContentOpStatusEnum = "FREEZE_FAILED"
    ListLoadbalanceByTagResponseContentOpStatusEnumClosing ListLoadbalanceByTagResponseContentOpStatusEnum = "CLOSING"
    ListLoadbalanceByTagResponseContentOpStatusEnumClosed ListLoadbalanceByTagResponseContentOpStatusEnum = "CLOSED"
    ListLoadbalanceByTagResponseContentOpStatusEnumCloseFailed ListLoadbalanceByTagResponseContentOpStatusEnum = "CLOSE_FAILED"
    ListLoadbalanceByTagResponseContentOpStatusEnumResuming ListLoadbalanceByTagResponseContentOpStatusEnum = "RESUMING"
    ListLoadbalanceByTagResponseContentOpStatusEnumResumeFailed ListLoadbalanceByTagResponseContentOpStatusEnum = "RESUME_FAILED"
    ListLoadbalanceByTagResponseContentOpStatusEnumCanceled ListLoadbalanceByTagResponseContentOpStatusEnum = "CANCELED"
    ListLoadbalanceByTagResponseContentOpStatusEnumWaitToChange ListLoadbalanceByTagResponseContentOpStatusEnum = "WAIT_TO_CHANGE"
    ListLoadbalanceByTagResponseContentOpStatusEnumChanging ListLoadbalanceByTagResponseContentOpStatusEnum = "CHANGING"
    ListLoadbalanceByTagResponseContentOpStatusEnumRenewed ListLoadbalanceByTagResponseContentOpStatusEnum = "RENEWED"
    ListLoadbalanceByTagResponseContentOpStatusEnumCloseToDelete ListLoadbalanceByTagResponseContentOpStatusEnum = "CLOSE_TO_DELETE"
    ListLoadbalanceByTagResponseContentOpStatusEnumFrozenToRenew ListLoadbalanceByTagResponseContentOpStatusEnum = "FROZEN_TO_RENEW"
    ListLoadbalanceByTagResponseContentOpStatusEnumCancelToSynchronize ListLoadbalanceByTagResponseContentOpStatusEnum = "CANCEL_TO_SYNCHRONIZE"
    ListLoadbalanceByTagResponseContentOpStatusEnumRenewing ListLoadbalanceByTagResponseContentOpStatusEnum = "RENEWING"
    ListLoadbalanceByTagResponseContentOpStatusEnumChangeClosing ListLoadbalanceByTagResponseContentOpStatusEnum = "CHANGE_CLOSING"
    ListLoadbalanceByTagResponseContentOpStatusEnumChanged ListLoadbalanceByTagResponseContentOpStatusEnum = "CHANGED"
    ListLoadbalanceByTagResponseContentOpStatusEnumChangeFailed ListLoadbalanceByTagResponseContentOpStatusEnum = "CHANGE_FAILED"
    ListLoadbalanceByTagResponseContentOpStatusEnumWaitToDelete ListLoadbalanceByTagResponseContentOpStatusEnum = "WAIT_TO_DELETE"
)

type ListLoadbalanceByTagResponseContent struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // VPC名称
	VpcName *string `json:"vpcName,omitempty"`
    // 绑定后端服务组数量
	BindedPools *int32 `json:"bindedPools,omitempty"`
    // 订单ID
	OrderId *string `json:"orderId,omitempty"`
    // 弹性负载均衡VIP内网地址
	PrivateIp *string `json:"privateIp,omitempty"`
    // 公网IP的带宽,单位:Mb
	BandwidthSize *int32 `json:"bandwidthSize,omitempty"`
    // 描述，存放负载均衡实地址
	Description *string `json:"description,omitempty"`
    // 需要绑定的独享设备IP地址，仅在isExclusive为true时生效，多个地址以“,”分隔例如：1.2.3.4,1.2.3.5
	NodeIp *string `json:"nodeIp,omitempty"`
    // 所绑定的监听器详情
	BindListeners *[]ListLoadbalanceByTagResponseBindListeners `json:"bindListeners,omitempty"`
    // 弹性负载均衡VIP网卡ID
	VipPortId *string `json:"vipPortId,omitempty"`
    // 规格为NFV托管平台订购时选择的云主机核数
	ServerCores *int32 `json:"serverCores,omitempty"`
    // 子网名称
	SubnetName *string `json:"subnetName,omitempty"`
    // 是否多AZ场景，true：是，false：不是
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 是否独享设备，true：是，false：不是
	IsExclusive *bool `json:"isExclusive,omitempty"`
    // 公网IP的ID
	IpId *string `json:"ipId,omitempty"`
    // 弹性负载均衡厂商
	Provider *ListLoadbalanceByTagResponseContentProviderEnum `json:"provider,omitempty"`
    // 路由器ID
	RouterId *string `json:"routerId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 弹性负载均衡关联子网的网络ID
	NetworkId *string `json:"networkId,omitempty"`
    // 弹性负载均衡ID
	Id *string `json:"id,omitempty"`
    // 弹性负载均衡是否可用,true:可用;false:不可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 双AZ子弹性负载均衡详情
	AzMappings []interface{} `json:"azMappings,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *ListLoadbalanceByTagResponseContentProductTypeEnum `json:"productType,omitempty"`
    // 计费方式
	MeasureType *string `json:"measureType,omitempty"`
    // 底层状态
	EcStatus *ListLoadbalanceByTagResponseContentEcStatusEnum `json:"ecStatus,omitempty"`
    // 绑定监听器数量
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 是否已可见,true:可见;false:不可见
	Visible *bool `json:"visible,omitempty"`
    // 是否开启了IPv6带宽，true：开通，false：未开通
	Ipv6Status *bool `json:"ipv6Status,omitempty"`
    // 是否是双az情况下的展示条目，true：是，false：不是
	IsMain *bool `json:"isMain,omitempty"`
    // 弹性负载均衡状态
	LoadbalancerStatus *string `json:"loadbalancerStatus,omitempty"`
    // 创建人，用户ID，例如：CICD-U-xxx
	Proposer *string `json:"proposer,omitempty"`
    // 公网IP地址
	PublicIp *string `json:"publicIp,omitempty"`
    // 订购用户
	UserName *string `json:"userName,omitempty"`
    // 数据版本
	Version *int32 `json:"version,omitempty"`
    // 双AZ弹性负载均衡的主实例ID
	ParentId *string `json:"parentId,omitempty"`
    // 边缘云标识
	Vaz *string `json:"vaz,omitempty"`
    // 标签信息
	Tags *[]ListLoadbalanceByTagResponseTags `json:"tags,omitempty"`
    // 弹性负载均衡规格
	Flavor *int32 `json:"flavor,omitempty"`
    // 设备配置
	EquipmentConfig *string `json:"equipmentConfig,omitempty"`
    // 是否已删除
	Deleted *bool `json:"deleted,omitempty"`
    // 子网类型
	IpVersion *ListLoadbalanceByTagResponseContentIpVersionEnum `json:"ipVersion,omitempty"`
    // 弹性负载均衡名称
	Name *string `json:"name,omitempty"`
    // 订单状态
	OpStatus *ListLoadbalanceByTagResponseContentOpStatusEnum `json:"opStatus,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
}

func (s ListLoadbalanceByTagResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagResponseContent) SetSubnetId(v string) *ListLoadbalanceByTagResponseContent {
	s.SubnetId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetVpcName(v string) *ListLoadbalanceByTagResponseContent {
	s.VpcName = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetBindedPools(v int32) *ListLoadbalanceByTagResponseContent {
	s.BindedPools = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetOrderId(v string) *ListLoadbalanceByTagResponseContent {
	s.OrderId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetPrivateIp(v string) *ListLoadbalanceByTagResponseContent {
	s.PrivateIp = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetBandwidthSize(v int32) *ListLoadbalanceByTagResponseContent {
	s.BandwidthSize = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetDescription(v string) *ListLoadbalanceByTagResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetNodeIp(v string) *ListLoadbalanceByTagResponseContent {
	s.NodeIp = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetBindListeners(v []ListLoadbalanceByTagResponseBindListeners) *ListLoadbalanceByTagResponseContent {
	s.BindListeners = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetVipPortId(v string) *ListLoadbalanceByTagResponseContent {
	s.VipPortId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetServerCores(v int32) *ListLoadbalanceByTagResponseContent {
	s.ServerCores = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetSubnetName(v string) *ListLoadbalanceByTagResponseContent {
	s.SubnetName = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIsMultiAz(v bool) *ListLoadbalanceByTagResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIsExclusive(v bool) *ListLoadbalanceByTagResponseContent {
	s.IsExclusive = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIpId(v string) *ListLoadbalanceByTagResponseContent {
	s.IpId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetProvider(v ListLoadbalanceByTagResponseContentProviderEnum) *ListLoadbalanceByTagResponseContent {
	s.Provider = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetRouterId(v string) *ListLoadbalanceByTagResponseContent {
	s.RouterId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetCreatedTime(v string) *ListLoadbalanceByTagResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetNetworkId(v string) *ListLoadbalanceByTagResponseContent {
	s.NetworkId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetId(v string) *ListLoadbalanceByTagResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetAdminStateUp(v bool) *ListLoadbalanceByTagResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetAzMappings(v []interface{}) *ListLoadbalanceByTagResponseContent {
	s.AzMappings = v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetProductType(v ListLoadbalanceByTagResponseContentProductTypeEnum) *ListLoadbalanceByTagResponseContent {
	s.ProductType = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetMeasureType(v string) *ListLoadbalanceByTagResponseContent {
	s.MeasureType = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetEcStatus(v ListLoadbalanceByTagResponseContentEcStatusEnum) *ListLoadbalanceByTagResponseContent {
	s.EcStatus = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetBindedListeners(v int32) *ListLoadbalanceByTagResponseContent {
	s.BindedListeners = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetVisible(v bool) *ListLoadbalanceByTagResponseContent {
	s.Visible = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIpv6Status(v bool) *ListLoadbalanceByTagResponseContent {
	s.Ipv6Status = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIsMain(v bool) *ListLoadbalanceByTagResponseContent {
	s.IsMain = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetLoadbalancerStatus(v string) *ListLoadbalanceByTagResponseContent {
	s.LoadbalancerStatus = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetProposer(v string) *ListLoadbalanceByTagResponseContent {
	s.Proposer = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetPublicIp(v string) *ListLoadbalanceByTagResponseContent {
	s.PublicIp = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetUserName(v string) *ListLoadbalanceByTagResponseContent {
	s.UserName = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetVersion(v int32) *ListLoadbalanceByTagResponseContent {
	s.Version = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetParentId(v string) *ListLoadbalanceByTagResponseContent {
	s.ParentId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetVaz(v string) *ListLoadbalanceByTagResponseContent {
	s.Vaz = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetTags(v []ListLoadbalanceByTagResponseTags) *ListLoadbalanceByTagResponseContent {
	s.Tags = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetFlavor(v int32) *ListLoadbalanceByTagResponseContent {
	s.Flavor = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetEquipmentConfig(v string) *ListLoadbalanceByTagResponseContent {
	s.EquipmentConfig = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetDeleted(v bool) *ListLoadbalanceByTagResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetIpVersion(v ListLoadbalanceByTagResponseContentIpVersionEnum) *ListLoadbalanceByTagResponseContent {
	s.IpVersion = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetName(v string) *ListLoadbalanceByTagResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetOpStatus(v ListLoadbalanceByTagResponseContentOpStatusEnum) *ListLoadbalanceByTagResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadbalanceByTagResponseContent) SetRegion(v string) *ListLoadbalanceByTagResponseContent {
	s.Region = &v
	return s
}


type ListLoadbalanceByTagResponseContentBuilder struct {
	s *ListLoadbalanceByTagResponseContent
}

func NewListLoadbalanceByTagResponseContentBuilder() *ListLoadbalanceByTagResponseContentBuilder {
	s := &ListLoadbalanceByTagResponseContent{}
	b := &ListLoadbalanceByTagResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) SubnetId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) VpcName(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.VpcName = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) BindedPools(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.BindedPools = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) OrderId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.OrderId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) PrivateIp(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.PrivateIp = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) BandwidthSize(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.BandwidthSize = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Description(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) NodeIp(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.NodeIp = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) BindListeners(v []ListLoadbalanceByTagResponseBindListeners) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.BindListeners = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) VipPortId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.VipPortId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) ServerCores(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.ServerCores = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) SubnetName(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.SubnetName = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) IsMultiAz(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) IsExclusive(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.IsExclusive = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) IpId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.IpId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Provider(v ListLoadbalanceByTagResponseContentProviderEnum) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Provider = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) RouterId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.RouterId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) CreatedTime(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) NetworkId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.NetworkId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Id(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) AdminStateUp(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) AzMappings(v []interface{}) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.AzMappings = v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) ProductType(v ListLoadbalanceByTagResponseContentProductTypeEnum) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.ProductType = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) MeasureType(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.MeasureType = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) EcStatus(v ListLoadbalanceByTagResponseContentEcStatusEnum) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.EcStatus = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) BindedListeners(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Visible(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Visible = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Ipv6Status(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Ipv6Status = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) IsMain(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.IsMain = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) LoadbalancerStatus(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.LoadbalancerStatus = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Proposer(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) PublicIp(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.PublicIp = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) UserName(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.UserName = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Version(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Version = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) ParentId(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.ParentId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Vaz(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Tags(v []ListLoadbalanceByTagResponseTags) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Tags = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Flavor(v int32) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Flavor = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) EquipmentConfig(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.EquipmentConfig = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Deleted(v bool) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) IpVersion(v ListLoadbalanceByTagResponseContentIpVersionEnum) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Name(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) OpStatus(v ListLoadbalanceByTagResponseContentOpStatusEnum) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Region(v string) *ListLoadbalanceByTagResponseContentBuilder {
	b.s.Region = &v
	return b
}

func (b *ListLoadbalanceByTagResponseContentBuilder) Build() *ListLoadbalanceByTagResponseContent {
	return b.s
}


