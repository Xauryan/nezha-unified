// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceResponseContentProviderEnum string

// List of Provider
const (
    ListLoadbalanceResponseContentProviderEnumBrocade ListLoadbalanceResponseContentProviderEnum = "brocade"
    ListLoadbalanceResponseContentProviderEnumF5hardware ListLoadbalanceResponseContentProviderEnum = "f5hardware"
    ListLoadbalanceResponseContentProviderEnumNokia ListLoadbalanceResponseContentProviderEnum = "nokia"
    ListLoadbalanceResponseContentProviderEnumBcslb ListLoadbalanceResponseContentProviderEnum = "bcslb"
    ListLoadbalanceResponseContentProviderEnumF5networksS ListLoadbalanceResponseContentProviderEnum = "f5networks_s"
    ListLoadbalanceResponseContentProviderEnumNsb ListLoadbalanceResponseContentProviderEnum = "nsb"
    ListLoadbalanceResponseContentProviderEnumEslb ListLoadbalanceResponseContentProviderEnum = "eslb"
    ListLoadbalanceResponseContentProviderEnumRjhardf5 ListLoadbalanceResponseContentProviderEnum = "rjhardf5"
    ListLoadbalanceResponseContentProviderEnumVirtual ListLoadbalanceResponseContentProviderEnum = "virtual"
    ListLoadbalanceResponseContentProviderEnumH3clb ListLoadbalanceResponseContentProviderEnum = "h3clb"
    ListLoadbalanceResponseContentProviderEnumH3virtual ListLoadbalanceResponseContentProviderEnum = "h3virtual"
    ListLoadbalanceResponseContentProviderEnumAliyun ListLoadbalanceResponseContentProviderEnum = "aliyun"
    ListLoadbalanceResponseContentProviderEnumNfvslb ListLoadbalanceResponseContentProviderEnum = "nfvslb"
)
type ListLoadbalanceResponseContentProductTypeEnum string

// List of ProductType
const (
    ListLoadbalanceResponseContentProductTypeEnumNormal ListLoadbalanceResponseContentProductTypeEnum = "NORMAL"
    ListLoadbalanceResponseContentProductTypeEnumDeCloud ListLoadbalanceResponseContentProductTypeEnum = "DE_CLOUD"
    ListLoadbalanceResponseContentProductTypeEnumAutoscaling ListLoadbalanceResponseContentProductTypeEnum = "AUTOSCALING"
    ListLoadbalanceResponseContentProductTypeEnumVo ListLoadbalanceResponseContentProductTypeEnum = "VO"
    ListLoadbalanceResponseContentProductTypeEnumCdn ListLoadbalanceResponseContentProductTypeEnum = "CDN"
    ListLoadbalanceResponseContentProductTypeEnumPaasMaster ListLoadbalanceResponseContentProductTypeEnum = "PAAS_MASTER"
    ListLoadbalanceResponseContentProductTypeEnumPaasSlave ListLoadbalanceResponseContentProductTypeEnum = "PAAS_SLAVE"
    ListLoadbalanceResponseContentProductTypeEnumVcpe ListLoadbalanceResponseContentProductTypeEnum = "VCPE"
    ListLoadbalanceResponseContentProductTypeEnumEmr ListLoadbalanceResponseContentProductTypeEnum = "EMR"
    ListLoadbalanceResponseContentProductTypeEnumLogaudit ListLoadbalanceResponseContentProductTypeEnum = "LOGAUDIT"
    ListLoadbalanceResponseContentProductTypeEnumMse ListLoadbalanceResponseContentProductTypeEnum = "MSE"
    ListLoadbalanceResponseContentProductTypeEnumVirtual ListLoadbalanceResponseContentProductTypeEnum = "VIRTUAL"
    ListLoadbalanceResponseContentProductTypeEnumBastion ListLoadbalanceResponseContentProductTypeEnum = "BASTION"
    ListLoadbalanceResponseContentProductTypeEnumNextGenerationFirewall ListLoadbalanceResponseContentProductTypeEnum = "NEXT_GENERATION_FIREWALL"
)
type ListLoadbalanceResponseContentEcStatusEnum string

// List of EcStatus
const (
    ListLoadbalanceResponseContentEcStatusEnumActive ListLoadbalanceResponseContentEcStatusEnum = "ACTIVE"
    ListLoadbalanceResponseContentEcStatusEnumBuild ListLoadbalanceResponseContentEcStatusEnum = "BUILD"
    ListLoadbalanceResponseContentEcStatusEnumDown ListLoadbalanceResponseContentEcStatusEnum = "DOWN"
    ListLoadbalanceResponseContentEcStatusEnumError ListLoadbalanceResponseContentEcStatusEnum = "ERROR"
    ListLoadbalanceResponseContentEcStatusEnumPendingCreate ListLoadbalanceResponseContentEcStatusEnum = "PENDING_CREATE"
    ListLoadbalanceResponseContentEcStatusEnumPendingUpdate ListLoadbalanceResponseContentEcStatusEnum = "PENDING_UPDATE"
    ListLoadbalanceResponseContentEcStatusEnumPendingDelete ListLoadbalanceResponseContentEcStatusEnum = "PENDING_DELETE"
    ListLoadbalanceResponseContentEcStatusEnumMigrating ListLoadbalanceResponseContentEcStatusEnum = "MIGRATING"
    ListLoadbalanceResponseContentEcStatusEnumDeleted ListLoadbalanceResponseContentEcStatusEnum = "DELETED"
    ListLoadbalanceResponseContentEcStatusEnumUnrecognized ListLoadbalanceResponseContentEcStatusEnum = "UNRECOGNIZED"
)
type ListLoadbalanceResponseContentIpVersionEnum int32

// List of IpVersion
const (
    ListLoadbalanceResponseContentIpVersionEnum4 ListLoadbalanceResponseContentIpVersionEnum = 4
    ListLoadbalanceResponseContentIpVersionEnum6 ListLoadbalanceResponseContentIpVersionEnum = 6
)
type ListLoadbalanceResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListLoadbalanceResponseContentOpStatusEnumNewBuilt ListLoadbalanceResponseContentOpStatusEnum = "NEW_BUILT"
    ListLoadbalanceResponseContentOpStatusEnumWaitToCreate ListLoadbalanceResponseContentOpStatusEnum = "WAIT_TO_CREATE"
    ListLoadbalanceResponseContentOpStatusEnumCreating ListLoadbalanceResponseContentOpStatusEnum = "CREATING"
    ListLoadbalanceResponseContentOpStatusEnumCreatedToSynchronize ListLoadbalanceResponseContentOpStatusEnum = "CREATED_TO_SYNCHRONIZE"
    ListLoadbalanceResponseContentOpStatusEnumCreated ListLoadbalanceResponseContentOpStatusEnum = "CREATED"
    ListLoadbalanceResponseContentOpStatusEnumCreateFailed ListLoadbalanceResponseContentOpStatusEnum = "CREATE_FAILED"
    ListLoadbalanceResponseContentOpStatusEnumFreezing ListLoadbalanceResponseContentOpStatusEnum = "FREEZING"
    ListLoadbalanceResponseContentOpStatusEnumFrozen ListLoadbalanceResponseContentOpStatusEnum = "FROZEN"
    ListLoadbalanceResponseContentOpStatusEnumFreezeFailed ListLoadbalanceResponseContentOpStatusEnum = "FREEZE_FAILED"
    ListLoadbalanceResponseContentOpStatusEnumClosing ListLoadbalanceResponseContentOpStatusEnum = "CLOSING"
    ListLoadbalanceResponseContentOpStatusEnumClosed ListLoadbalanceResponseContentOpStatusEnum = "CLOSED"
    ListLoadbalanceResponseContentOpStatusEnumCloseFailed ListLoadbalanceResponseContentOpStatusEnum = "CLOSE_FAILED"
    ListLoadbalanceResponseContentOpStatusEnumResuming ListLoadbalanceResponseContentOpStatusEnum = "RESUMING"
    ListLoadbalanceResponseContentOpStatusEnumResumeFailed ListLoadbalanceResponseContentOpStatusEnum = "RESUME_FAILED"
    ListLoadbalanceResponseContentOpStatusEnumCanceled ListLoadbalanceResponseContentOpStatusEnum = "CANCELED"
    ListLoadbalanceResponseContentOpStatusEnumWaitToChange ListLoadbalanceResponseContentOpStatusEnum = "WAIT_TO_CHANGE"
    ListLoadbalanceResponseContentOpStatusEnumChanging ListLoadbalanceResponseContentOpStatusEnum = "CHANGING"
    ListLoadbalanceResponseContentOpStatusEnumRenewed ListLoadbalanceResponseContentOpStatusEnum = "RENEWED"
    ListLoadbalanceResponseContentOpStatusEnumCloseToDelete ListLoadbalanceResponseContentOpStatusEnum = "CLOSE_TO_DELETE"
    ListLoadbalanceResponseContentOpStatusEnumFrozenToRenew ListLoadbalanceResponseContentOpStatusEnum = "FROZEN_TO_RENEW"
    ListLoadbalanceResponseContentOpStatusEnumCancelToSynchronize ListLoadbalanceResponseContentOpStatusEnum = "CANCEL_TO_SYNCHRONIZE"
    ListLoadbalanceResponseContentOpStatusEnumRenewing ListLoadbalanceResponseContentOpStatusEnum = "RENEWING"
    ListLoadbalanceResponseContentOpStatusEnumChangeClosing ListLoadbalanceResponseContentOpStatusEnum = "CHANGE_CLOSING"
    ListLoadbalanceResponseContentOpStatusEnumChanged ListLoadbalanceResponseContentOpStatusEnum = "CHANGED"
    ListLoadbalanceResponseContentOpStatusEnumChangeFailed ListLoadbalanceResponseContentOpStatusEnum = "CHANGE_FAILED"
    ListLoadbalanceResponseContentOpStatusEnumWaitToDelete ListLoadbalanceResponseContentOpStatusEnum = "WAIT_TO_DELETE"
)

type ListLoadbalanceResponseContent struct {

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
    // 描述，存放负载均衡器实地址
	Description *string `json:"description,omitempty"`
    // 需要绑定的独享设备IP地址，仅在isExclusive为true时生效，多个地址以“，”分隔例如：1.2.3.4,1.2.3.5
	NodeIp *string `json:"nodeIp,omitempty"`
    // 所绑定的监听器详情
	BindListeners *[]ListLoadbalanceResponseBindListeners `json:"bindListeners,omitempty"`
    // 弹性负载均衡VIP网卡ID
	VipPortId *string `json:"vipPortId,omitempty"`
    // 规格为NFV托管平台订购时选择的云主机核数
	ServerCores *int32 `json:"serverCores,omitempty"`
    // 子网名称
	SubnetName *string `json:"subnetName,omitempty"`
    // 是否双AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 是否独享设备，true是，false不是
	IsExclusive *bool `json:"isExclusive,omitempty"`
    // 公网IP的ID
	IpId *string `json:"ipId,omitempty"`
    // 弹性负载均衡厂商
	Provider *ListLoadbalanceResponseContentProviderEnum `json:"provider,omitempty"`
    // 路由器ID
	RouterId *string `json:"routerId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 弹性负载均衡关联子网的网络ID
	NetworkId *string `json:"networkId,omitempty"`
    // 弹性负载均衡的ID
	Id *string `json:"id,omitempty"`
    // 是否可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 双AZ子弹性负载均衡详情
	AzMappings []interface{} `json:"azMappings,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *ListLoadbalanceResponseContentProductTypeEnum `json:"productType,omitempty"`
    // 计费方式
	MeasureType *string `json:"measureType,omitempty"`
    // 弹性负载均衡状态： ACTIVE（资源已激活），BUILD（资源已创建），DOWN（资源已创建但未激活），ERROR（资源无法工作），（PENDING_CREATE,PENDING_UPDATE）资源正在创建当中，PENDING_DELETE（资源在删除当中），MIGRATING（资源在迁移中），DELETED（资源已删除），UNRECOGNIZED（未知状态）
	EcStatus *ListLoadbalanceResponseContentEcStatusEnum `json:"ecStatus,omitempty"`
    // 绑定监听器数量
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 是否已可见
	Visible *bool `json:"visible,omitempty"`
    // 是否开启了IPv6带宽，true开通，false未开通
	Ipv6Status *bool `json:"ipv6Status,omitempty"`
    // 双az情况下的展示条目
	IsMain *bool `json:"isMain,omitempty"`
    // 弹性负载均衡器状态
	LoadbalancerStatus *string `json:"loadbalancerStatus,omitempty"`
    // 创建人，用户ID，例如：CICD-U-xxxxxx
	Proposer *string `json:"proposer,omitempty"`
    // 公网IP地址
	PublicIp *string `json:"publicIp,omitempty"`
    // 订购用户
	UserName *string `json:"userName,omitempty"`
    // 数据版本
	Version *int32 `json:"version,omitempty"`
    // 双AZ弹性负载均衡的主实例ID
	ParentId *string `json:"parentId,omitempty"`
    // 边缘云vaz
	Vaz *string `json:"vaz,omitempty"`
    // 标签
	Tags *[]ListLoadbalanceResponseTags `json:"tags,omitempty"`
    // 弹性负载均衡规格
	Flavor *int32 `json:"flavor,omitempty"`
    // 设备配置
	EquipmentConfig *string `json:"equipmentConfig,omitempty"`
    // 是否已删除
	Deleted *bool `json:"deleted,omitempty"`
    // 子网类型
	IpVersion *ListLoadbalanceResponseContentIpVersionEnum `json:"ipVersion,omitempty"`
    // 名称
	Name *string `json:"name,omitempty"`
    // 订单状态： NEW_BUILT（新建）,WAIT_TO_CREATE（等待开通）,CREATING（正在开通）,CREATED_TO_SYNCHRONIZE（开通待同步）,CREATED（已开通）,CREATE_FAILED（开头失败）,FREEZING（正在冻结）,FROZEN（已冻结）,FREEZE_FAILED（冻结失败）,CLOSING（正在关闭）,CLOSED（已经闭）,CLOSE_FAILED（关闭失败）,RESUMING（正在恢复）,RESUME_FAILED（恢复失败）,CANCELED（已取消）,WAIT_TO_CHANGE（等待变更）,CHANGING（正在变更）,RENEWED（已续订）,CLOSE_TO_DELETE（关闭待删除）,FROZEN_TO_RENEW（冻结待续订）,CANCEL_TO_SYNCHRO
	OpStatus *ListLoadbalanceResponseContentOpStatusEnum `json:"opStatus,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
}

func (s ListLoadbalanceResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceResponseContent) GoString() string {
	return s.String()
}

func (s ListLoadbalanceResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceResponseContent) SetSubnetId(v string) *ListLoadbalanceResponseContent {
	s.SubnetId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetVpcName(v string) *ListLoadbalanceResponseContent {
	s.VpcName = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetBindedPools(v int32) *ListLoadbalanceResponseContent {
	s.BindedPools = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetOrderId(v string) *ListLoadbalanceResponseContent {
	s.OrderId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetPrivateIp(v string) *ListLoadbalanceResponseContent {
	s.PrivateIp = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetBandwidthSize(v int32) *ListLoadbalanceResponseContent {
	s.BandwidthSize = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetDescription(v string) *ListLoadbalanceResponseContent {
	s.Description = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetNodeIp(v string) *ListLoadbalanceResponseContent {
	s.NodeIp = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetBindListeners(v []ListLoadbalanceResponseBindListeners) *ListLoadbalanceResponseContent {
	s.BindListeners = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetVipPortId(v string) *ListLoadbalanceResponseContent {
	s.VipPortId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetServerCores(v int32) *ListLoadbalanceResponseContent {
	s.ServerCores = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetSubnetName(v string) *ListLoadbalanceResponseContent {
	s.SubnetName = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIsMultiAz(v bool) *ListLoadbalanceResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIsExclusive(v bool) *ListLoadbalanceResponseContent {
	s.IsExclusive = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIpId(v string) *ListLoadbalanceResponseContent {
	s.IpId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetProvider(v ListLoadbalanceResponseContentProviderEnum) *ListLoadbalanceResponseContent {
	s.Provider = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetRouterId(v string) *ListLoadbalanceResponseContent {
	s.RouterId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetCreatedTime(v string) *ListLoadbalanceResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetNetworkId(v string) *ListLoadbalanceResponseContent {
	s.NetworkId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetId(v string) *ListLoadbalanceResponseContent {
	s.Id = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetAdminStateUp(v bool) *ListLoadbalanceResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetAzMappings(v []interface{}) *ListLoadbalanceResponseContent {
	s.AzMappings = v
	return s
}

func (s *ListLoadbalanceResponseContent) SetProductType(v ListLoadbalanceResponseContentProductTypeEnum) *ListLoadbalanceResponseContent {
	s.ProductType = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetMeasureType(v string) *ListLoadbalanceResponseContent {
	s.MeasureType = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetEcStatus(v ListLoadbalanceResponseContentEcStatusEnum) *ListLoadbalanceResponseContent {
	s.EcStatus = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetBindedListeners(v int32) *ListLoadbalanceResponseContent {
	s.BindedListeners = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetVisible(v bool) *ListLoadbalanceResponseContent {
	s.Visible = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIpv6Status(v bool) *ListLoadbalanceResponseContent {
	s.Ipv6Status = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIsMain(v bool) *ListLoadbalanceResponseContent {
	s.IsMain = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetLoadbalancerStatus(v string) *ListLoadbalanceResponseContent {
	s.LoadbalancerStatus = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetProposer(v string) *ListLoadbalanceResponseContent {
	s.Proposer = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetPublicIp(v string) *ListLoadbalanceResponseContent {
	s.PublicIp = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetUserName(v string) *ListLoadbalanceResponseContent {
	s.UserName = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetVersion(v int32) *ListLoadbalanceResponseContent {
	s.Version = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetParentId(v string) *ListLoadbalanceResponseContent {
	s.ParentId = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetVaz(v string) *ListLoadbalanceResponseContent {
	s.Vaz = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetTags(v []ListLoadbalanceResponseTags) *ListLoadbalanceResponseContent {
	s.Tags = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetFlavor(v int32) *ListLoadbalanceResponseContent {
	s.Flavor = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetEquipmentConfig(v string) *ListLoadbalanceResponseContent {
	s.EquipmentConfig = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetDeleted(v bool) *ListLoadbalanceResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetIpVersion(v ListLoadbalanceResponseContentIpVersionEnum) *ListLoadbalanceResponseContent {
	s.IpVersion = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetName(v string) *ListLoadbalanceResponseContent {
	s.Name = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetOpStatus(v ListLoadbalanceResponseContentOpStatusEnum) *ListLoadbalanceResponseContent {
	s.OpStatus = &v
	return s
}

func (s *ListLoadbalanceResponseContent) SetRegion(v string) *ListLoadbalanceResponseContent {
	s.Region = &v
	return s
}


type ListLoadbalanceResponseContentBuilder struct {
	s *ListLoadbalanceResponseContent
}

func NewListLoadbalanceResponseContentBuilder() *ListLoadbalanceResponseContentBuilder {
	s := &ListLoadbalanceResponseContent{}
	b := &ListLoadbalanceResponseContentBuilder{s: s}
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) SubnetId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) VpcName(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.VpcName = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) BindedPools(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.BindedPools = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) OrderId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.OrderId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) PrivateIp(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.PrivateIp = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) BandwidthSize(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.BandwidthSize = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Description(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) NodeIp(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.NodeIp = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) BindListeners(v []ListLoadbalanceResponseBindListeners) *ListLoadbalanceResponseContentBuilder {
	b.s.BindListeners = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) VipPortId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.VipPortId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) ServerCores(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.ServerCores = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) SubnetName(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.SubnetName = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) IsMultiAz(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) IsExclusive(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.IsExclusive = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) IpId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.IpId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Provider(v ListLoadbalanceResponseContentProviderEnum) *ListLoadbalanceResponseContentBuilder {
	b.s.Provider = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) RouterId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.RouterId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) CreatedTime(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) NetworkId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.NetworkId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Id(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) AdminStateUp(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) AzMappings(v []interface{}) *ListLoadbalanceResponseContentBuilder {
	b.s.AzMappings = v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) ProductType(v ListLoadbalanceResponseContentProductTypeEnum) *ListLoadbalanceResponseContentBuilder {
	b.s.ProductType = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) MeasureType(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.MeasureType = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) EcStatus(v ListLoadbalanceResponseContentEcStatusEnum) *ListLoadbalanceResponseContentBuilder {
	b.s.EcStatus = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) BindedListeners(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Visible(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.Visible = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Ipv6Status(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.Ipv6Status = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) IsMain(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.IsMain = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) LoadbalancerStatus(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.LoadbalancerStatus = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Proposer(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) PublicIp(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.PublicIp = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) UserName(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.UserName = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Version(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.Version = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) ParentId(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.ParentId = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Vaz(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Tags(v []ListLoadbalanceResponseTags) *ListLoadbalanceResponseContentBuilder {
	b.s.Tags = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Flavor(v int32) *ListLoadbalanceResponseContentBuilder {
	b.s.Flavor = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) EquipmentConfig(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.EquipmentConfig = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Deleted(v bool) *ListLoadbalanceResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) IpVersion(v ListLoadbalanceResponseContentIpVersionEnum) *ListLoadbalanceResponseContentBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Name(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) OpStatus(v ListLoadbalanceResponseContentOpStatusEnum) *ListLoadbalanceResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Region(v string) *ListLoadbalanceResponseContentBuilder {
	b.s.Region = &v
	return b
}

func (b *ListLoadbalanceResponseContentBuilder) Build() *ListLoadbalanceResponseContent {
	return b.s
}


