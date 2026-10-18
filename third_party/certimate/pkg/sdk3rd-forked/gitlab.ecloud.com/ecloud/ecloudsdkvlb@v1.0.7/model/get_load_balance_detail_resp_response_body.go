// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetLoadBalanceDetailRespResponseBodyEcStatusEnum string

// List of EcStatus
const (
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumActive GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "ACTIVE"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumBuild GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "BUILD"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumDown GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "DOWN"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumError GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "ERROR"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumPendingCreate GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "PENDING_CREATE"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumPendingUpdate GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "PENDING_UPDATE"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumPendingDelete GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "PENDING_DELETE"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumMigrating GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "MIGRATING"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumDeleted GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "DELETED"
    GetLoadBalanceDetailRespResponseBodyEcStatusEnumUnrecognized GetLoadBalanceDetailRespResponseBodyEcStatusEnum = "UNRECOGNIZED"
)

type GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum int32

// List of IpVersionTypes
const (
    GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum4 GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum = 4
    GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum6 GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum = 6
)
type GetLoadBalanceDetailRespResponseBodyProviderEnum string

// List of Provider
const (
    GetLoadBalanceDetailRespResponseBodyProviderEnumBrocade GetLoadBalanceDetailRespResponseBodyProviderEnum = "brocade"
    GetLoadBalanceDetailRespResponseBodyProviderEnumF5hardware GetLoadBalanceDetailRespResponseBodyProviderEnum = "f5hardware"
    GetLoadBalanceDetailRespResponseBodyProviderEnumNokia GetLoadBalanceDetailRespResponseBodyProviderEnum = "nokia"
    GetLoadBalanceDetailRespResponseBodyProviderEnumBcslb GetLoadBalanceDetailRespResponseBodyProviderEnum = "bcslb"
    GetLoadBalanceDetailRespResponseBodyProviderEnumF5networksS GetLoadBalanceDetailRespResponseBodyProviderEnum = "f5networks_s"
    GetLoadBalanceDetailRespResponseBodyProviderEnumNsb GetLoadBalanceDetailRespResponseBodyProviderEnum = "nsb"
    GetLoadBalanceDetailRespResponseBodyProviderEnumEslb GetLoadBalanceDetailRespResponseBodyProviderEnum = "eslb"
    GetLoadBalanceDetailRespResponseBodyProviderEnumRjhardf5 GetLoadBalanceDetailRespResponseBodyProviderEnum = "rjhardf5"
    GetLoadBalanceDetailRespResponseBodyProviderEnumVirtual GetLoadBalanceDetailRespResponseBodyProviderEnum = "virtual"
    GetLoadBalanceDetailRespResponseBodyProviderEnumH3clb GetLoadBalanceDetailRespResponseBodyProviderEnum = "h3clb"
    GetLoadBalanceDetailRespResponseBodyProviderEnumH3virtual GetLoadBalanceDetailRespResponseBodyProviderEnum = "h3virtual"
    GetLoadBalanceDetailRespResponseBodyProviderEnumAliyun GetLoadBalanceDetailRespResponseBodyProviderEnum = "aliyun"
    GetLoadBalanceDetailRespResponseBodyProviderEnumNfvslb GetLoadBalanceDetailRespResponseBodyProviderEnum = "nfvslb"
)
type GetLoadBalanceDetailRespResponseBodyProductTypeEnum string

// List of ProductType
const (
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumNormal GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "NORMAL"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumDeCloud GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "DE_CLOUD"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumAutoscaling GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "AUTOSCALING"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumVo GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "VO"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumCdn GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "CDN"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumPaasMaster GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "PAAS_MASTER"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumPaasSlave GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "PAAS_SLAVE"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumVcpe GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "VCPE"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumEmr GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "EMR"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumLogaudit GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "LOGAUDIT"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumMse GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "MSE"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumVirtual GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "VIRTUAL"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumBastion GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "BASTION"
    GetLoadBalanceDetailRespResponseBodyProductTypeEnumNextGenerationFirewall GetLoadBalanceDetailRespResponseBodyProductTypeEnum = "NEXT_GENERATION_FIREWALL"
)
type GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum string

// List of RecycleStatus
const (
    GetLoadBalanceDetailRespResponseBodyRecycleStatusEnumUnrecycle GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum = "UNRECYCLE"
    GetLoadBalanceDetailRespResponseBodyRecycleStatusEnumRecycleInProgress GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum = "RECYCLE_IN_PROGRESS"
    GetLoadBalanceDetailRespResponseBodyRecycleStatusEnumRecycleToBeDeleted GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum = "RECYCLE_TO_BE_DELETED"
    GetLoadBalanceDetailRespResponseBodyRecycleStatusEnumRecycleRecoverInProgress GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum = "RECYCLE_RECOVER_IN_PROGRESS"
    GetLoadBalanceDetailRespResponseBodyRecycleStatusEnumRecycleRecoverFailed GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum = "RECYCLE_RECOVER_FAILED"
)
type GetLoadBalanceDetailRespResponseBodyIpVersionEnum int32

// List of IpVersion
const (
    GetLoadBalanceDetailRespResponseBodyIpVersionEnum4 GetLoadBalanceDetailRespResponseBodyIpVersionEnum = 4
    GetLoadBalanceDetailRespResponseBodyIpVersionEnum6 GetLoadBalanceDetailRespResponseBodyIpVersionEnum = 6
)

type GetLoadBalanceDetailRespResponseBody struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 订单ID
	OrderId *string `json:"orderId,omitempty"`
    // 弹性负载均衡占用VIP地址
	PrivateIp *string `json:"privateIp,omitempty"`
    // 弹性负载均衡占用VIP的网卡ID
	VipPortId *string `json:"vipPortId,omitempty"`
    // 规格为NFV托管平台订购时选择的云主机核数
	ServerCores *int32 `json:"serverCores,omitempty"`
    // 是否双AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // DNAT模式启用状态
	Dnat *bool `json:"dnat,omitempty"`
    // 公网IP的ID
	IpId *string `json:"ipId,omitempty"`
    // 客户是否停用
	Enable *bool `json:"enable,omitempty"`
    // 路由器ID
	RouterId *string `json:"routerId,omitempty"`
    // 是否放入回收站
	Recycle *bool `json:"recycle,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 弹性负载均衡关联子网的网络ID
	NetworkId *string `json:"networkId,omitempty"`
    // 弹性负载均衡的ID
	Id *string `json:"id,omitempty"`
    // 弹性负载均衡状态： ACTIVE（资源已激活），BUILD（资源已创建），DOWN（资源已创建但未激活），ERROR（资源无法工作），（PENDING_CREATE,PENDING_UPDATE）资源正在创建当中，PENDING_DELETE（资源在删除当中），MIGRATING（资源在迁移中），DELETED（资源已删除），UNRECOGNIZED（未知状态）
	EcStatus *GetLoadBalanceDetailRespResponseBodyEcStatusEnum `json:"ecStatus,omitempty"`
    // 是否已可见
	Visible *bool `json:"visible,omitempty"`
    // 是否是主弹性负载均衡
	IsMain *bool `json:"isMain,omitempty"`
    // 创建人，用户ID，例如：CICD-U-xxxxxx
	Proposer *string `json:"proposer,omitempty"`
    // 弹性负载均衡搭载主机支持的子网IP类型，V4/V6
	IpVersionTypes []GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum `json:"ipVersionTypes,omitempty"`
    // 弹性负载均衡的版本，默认为0
	Version *int32 `json:"version,omitempty"`
    // 七层访问日志开关，部分灰度用户使用
	AccessLog *bool `json:"accessLog,omitempty"`
    // 边缘云vaz
	Vaz *string `json:"vaz,omitempty"`
    // 标签
	Tags *[]GetLoadBalanceDetailRespResponseTags `json:"tags,omitempty"`
    // 规格
	Flavor *int32 `json:"flavor,omitempty"`
    // 弹性负载均衡占用的MAC地址，仅DNAT模式实例支持
	MacAddress map[string]*string `json:"macAddress,omitempty"`
    // 加入回收站次数
	RecycleCount *int32 `json:"recycleCount,omitempty"`
    // 弹性负载均衡名称
	Name *string `json:"name,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 放入回收站时间
	RecycleTime *string `json:"recycleTime,omitempty"`
    // VPC名称
	VpcName *string `json:"vpcName,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 需要绑定的独享设备IP地址，仅在isExclusive为true时生效，多个地址以“,”分隔例如：1.2.3.4,1.2.3.5
	NodeIp *string `json:"nodeIp,omitempty"`
    // 查看该负载均衡器的订单状态，也可以标记资源正在创建中
	OrderStatus *string `json:"orderStatus,omitempty"`
    // 子网名称
	SubnetName *string `json:"subnetName,omitempty"`
    // 最大并发连接数
	MaxConcurrency *int32 `json:"maxConcurrency,omitempty"`
    // 是否独享设备，true是，false不是
	IsExclusive *bool `json:"isExclusive,omitempty"`
    // 弹性负载均衡厂商
	Provider *GetLoadBalanceDetailRespResponseBodyProviderEnum `json:"provider,omitempty"`
    // 是否是新版本弹性负载均衡
	LbIsNew *bool `json:"lbIsNew,omitempty"`
    // 新建连接数
	NewConnection *int32 `json:"newConnection,omitempty"`
    // 负载均衡器是否可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 负载均衡信息
	AzMappings []interface{} `json:"azMappings,omitempty"`
    // SNAT POOL地址的子网id
	SnatSubnetId *string `json:"snatSubnetId,omitempty"`
    // 产品类型，客户订购默认为NORMAL类型
	ProductType *GetLoadBalanceDetailRespResponseBodyProductTypeEnum `json:"productType,omitempty"`
    // 带宽，单位：Mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`
    // 弹性负载均衡状态
	LoadbalancerStatus *string `json:"loadbalancerStatus,omitempty"`
    // 公网IP地址
	PublicIp *string `json:"publicIp,omitempty"`
    // 双AZ情况下主弹性负载均衡的ID
	ParentId *string `json:"parentId,omitempty"`
    // 回收状态
	RecycleStatus *GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum `json:"recycleStatus,omitempty"`
    // 是否已删除
	Deleted *bool `json:"deleted,omitempty"`
    // 子网类型
	IpVersion *GetLoadBalanceDetailRespResponseBodyIpVersionEnum `json:"ipVersion,omitempty"`
}

func (s GetLoadBalanceDetailRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespResponseBody) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespResponseBody) SetSubnetId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.SubnetId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetOrderId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetPrivateIp(v string) *GetLoadBalanceDetailRespResponseBody {
	s.PrivateIp = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetVipPortId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.VipPortId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetServerCores(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.ServerCores = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIsMultiAz(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetDnat(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.Dnat = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIpId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.IpId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetEnable(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.Enable = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRouterId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.RouterId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRecycle(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.Recycle = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetCreatedTime(v string) *GetLoadBalanceDetailRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetNetworkId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.NetworkId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Id = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetEcStatus(v GetLoadBalanceDetailRespResponseBodyEcStatusEnum) *GetLoadBalanceDetailRespResponseBody {
	s.EcStatus = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetVisible(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.Visible = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIsMain(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.IsMain = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetProposer(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Proposer = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIpVersionTypes(v []GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum) *GetLoadBalanceDetailRespResponseBody {
	s.IpVersionTypes = v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetVersion(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.Version = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetAccessLog(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.AccessLog = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetVaz(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Vaz = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetTags(v []GetLoadBalanceDetailRespResponseTags) *GetLoadBalanceDetailRespResponseBody {
	s.Tags = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetFlavor(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.Flavor = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetMacAddress(v map[string]*string) *GetLoadBalanceDetailRespResponseBody {
	s.MacAddress = v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRecycleCount(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.RecycleCount = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetName(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Name = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRegion(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Region = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRecycleTime(v string) *GetLoadBalanceDetailRespResponseBody {
	s.RecycleTime = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetVpcName(v string) *GetLoadBalanceDetailRespResponseBody {
	s.VpcName = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetDescription(v string) *GetLoadBalanceDetailRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetNodeIp(v string) *GetLoadBalanceDetailRespResponseBody {
	s.NodeIp = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetOrderStatus(v string) *GetLoadBalanceDetailRespResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetSubnetName(v string) *GetLoadBalanceDetailRespResponseBody {
	s.SubnetName = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetMaxConcurrency(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.MaxConcurrency = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIsExclusive(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.IsExclusive = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetProvider(v GetLoadBalanceDetailRespResponseBodyProviderEnum) *GetLoadBalanceDetailRespResponseBody {
	s.Provider = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetLbIsNew(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.LbIsNew = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetNewConnection(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.NewConnection = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetAdminStateUp(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.AdminStateUp = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetAzMappings(v []interface{}) *GetLoadBalanceDetailRespResponseBody {
	s.AzMappings = v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetSnatSubnetId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.SnatSubnetId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetProductType(v GetLoadBalanceDetailRespResponseBodyProductTypeEnum) *GetLoadBalanceDetailRespResponseBody {
	s.ProductType = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetBandwidth(v int32) *GetLoadBalanceDetailRespResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetLoadbalancerStatus(v string) *GetLoadBalanceDetailRespResponseBody {
	s.LoadbalancerStatus = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetPublicIp(v string) *GetLoadBalanceDetailRespResponseBody {
	s.PublicIp = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetParentId(v string) *GetLoadBalanceDetailRespResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetRecycleStatus(v GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum) *GetLoadBalanceDetailRespResponseBody {
	s.RecycleStatus = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetDeleted(v bool) *GetLoadBalanceDetailRespResponseBody {
	s.Deleted = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseBody) SetIpVersion(v GetLoadBalanceDetailRespResponseBodyIpVersionEnum) *GetLoadBalanceDetailRespResponseBody {
	s.IpVersion = &v
	return s
}


type GetLoadBalanceDetailRespResponseBodyBuilder struct {
	s *GetLoadBalanceDetailRespResponseBody
}

func NewGetLoadBalanceDetailRespResponseBodyBuilder() *GetLoadBalanceDetailRespResponseBodyBuilder {
	s := &GetLoadBalanceDetailRespResponseBody{}
	b := &GetLoadBalanceDetailRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) SubnetId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) OrderId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) PrivateIp(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.PrivateIp = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) VipPortId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.VipPortId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) ServerCores(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.ServerCores = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IsMultiAz(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Dnat(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Dnat = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IpId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IpId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Enable(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Enable = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) RouterId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.RouterId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Recycle(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Recycle = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) CreatedTime(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) NetworkId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.NetworkId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Id(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) EcStatus(v GetLoadBalanceDetailRespResponseBodyEcStatusEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.EcStatus = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Visible(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Visible = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IsMain(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IsMain = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Proposer(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Proposer = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IpVersionTypes(v []GetLoadBalanceDetailRespResponseBodyIpVersionTypesEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IpVersionTypes = v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Version(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Version = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) AccessLog(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.AccessLog = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Vaz(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Tags(v []GetLoadBalanceDetailRespResponseTags) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Tags = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Flavor(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Flavor = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) MacAddress(v map[string]*string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.MacAddress = v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) RecycleCount(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.RecycleCount = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Name(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Region(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) RecycleTime(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.RecycleTime = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) VpcName(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Description(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) NodeIp(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.NodeIp = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) OrderStatus(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.OrderStatus = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) SubnetName(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.SubnetName = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) MaxConcurrency(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.MaxConcurrency = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IsExclusive(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IsExclusive = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Provider(v GetLoadBalanceDetailRespResponseBodyProviderEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Provider = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) LbIsNew(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.LbIsNew = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) NewConnection(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.NewConnection = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) AdminStateUp(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) AzMappings(v []interface{}) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.AzMappings = v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) SnatSubnetId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.SnatSubnetId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) ProductType(v GetLoadBalanceDetailRespResponseBodyProductTypeEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Bandwidth(v int32) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Bandwidth = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) LoadbalancerStatus(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.LoadbalancerStatus = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) PublicIp(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.PublicIp = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) ParentId(v string) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.ParentId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) RecycleStatus(v GetLoadBalanceDetailRespResponseBodyRecycleStatusEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.RecycleStatus = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Deleted(v bool) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.Deleted = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) IpVersion(v GetLoadBalanceDetailRespResponseBodyIpVersionEnum) *GetLoadBalanceDetailRespResponseBodyBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseBodyBuilder) Build() *GetLoadBalanceDetailRespResponseBody {
	return b.s
}


