// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListLoadbalanceQueryJoinProductTypeEnum string

// List of JoinProductType
const (
    ListLoadbalanceQueryJoinProductTypeEnumFip ListLoadbalanceQueryJoinProductTypeEnum = "FIP"
)
type ListLoadbalanceQueryIpVersionEnum string

// List of IpVersion
const (
    ListLoadbalanceQueryIpVersionEnumV4 ListLoadbalanceQueryIpVersionEnum = "V4"
    ListLoadbalanceQueryIpVersionEnumV6 ListLoadbalanceQueryIpVersionEnum = "V6"
)

type ListLoadbalanceQuery struct {
    position.Query
    // 是否可见，默认值：【true】
	Visible *bool `json:"visible,omitempty"`
    // 标签值
	TagValue *string `json:"tagValue,omitempty"`
    // 是否显示主弹性负载均衡
	IsMain *bool `json:"isMain,omitempty"`
    // 对接产品类型（FIP：公网IP）
	JoinProductType *ListLoadbalanceQueryJoinProductTypeEnum `json:"joinProductType,omitempty"`
    // 标签ID列表
	TagIds []string `json:"tagIds,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 排序方式，取值为desc（默认）：根据LB的创建时间降序排列；asc：根据LB的创建时间升序排列
	OrderBy *string `json:"orderBy,omitempty"`
    // 弹性负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
    // 边缘云虚拟可用区
	Vaz *string `json:"vaz,omitempty"`
    // 查询类型，弹性负载均衡名称查询填LOAD_BALANCE、内网地址查询填PRIVATE_IP、公网IP查询填PUBLIC_IP、VPC名称查询填VPC_NAME、子网名称查询填SUBNET_NAME、订单ID查询填ORDER_ID（需与queryWord组合使用）
	QueryType *string `json:"queryType,omitempty"`
    // 弹性负载均衡IP的类型，V4或V6
	IpVersion *ListLoadbalanceQueryIpVersionEnum `json:"ipVersion,omitempty"`
    // 是否为闲置实例，默认值：【false】
	IsFree *bool `json:"isFree,omitempty"`
    // 是否展示虚拟订单，默认值：【false】
	ShowVirtualLoadBalancers *bool `json:"showVirtualLoadBalancers,omitempty"`
    // 路由器ID
	RouterId *string `json:"routerId,omitempty"`
    // 查询结果是否展示关联标签
	IsDescribeTags *bool `json:"isDescribeTags,omitempty"`
    // 是否已绑定弹性公网IP
	FipBind *bool `json:"fipBind,omitempty"`
    // 弹性负载均衡关联子网的网络ID
	NetworkId *string `json:"networkId,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 查询关键字，根据负载均衡名称、内网地址、公网IP、VPC名称、子网名称、订单ID（需与queryType组合使用）
	QueryWord *string `json:"queryWord,omitempty"`
    // 是否可用
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 标签键
	TagKey *string `json:"tagKey,omitempty"`
}

func (s ListLoadbalanceQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceQuery) GoString() string {
	return s.String()
}

func (s ListLoadbalanceQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceQuery) SetVisible(v bool) *ListLoadbalanceQuery {
	s.Visible = &v
	return s
}

func (s *ListLoadbalanceQuery) SetTagValue(v string) *ListLoadbalanceQuery {
	s.TagValue = &v
	return s
}

func (s *ListLoadbalanceQuery) SetIsMain(v bool) *ListLoadbalanceQuery {
	s.IsMain = &v
	return s
}

func (s *ListLoadbalanceQuery) SetJoinProductType(v ListLoadbalanceQueryJoinProductTypeEnum) *ListLoadbalanceQuery {
	s.JoinProductType = &v
	return s
}

func (s *ListLoadbalanceQuery) SetTagIds(v []string) *ListLoadbalanceQuery {
	s.TagIds = v
	return s
}

func (s *ListLoadbalanceQuery) SetPageSize(v int32) *ListLoadbalanceQuery {
	s.PageSize = &v
	return s
}

func (s *ListLoadbalanceQuery) SetOrderBy(v string) *ListLoadbalanceQuery {
	s.OrderBy = &v
	return s
}

func (s *ListLoadbalanceQuery) SetLoadBalanceId(v string) *ListLoadbalanceQuery {
	s.LoadBalanceId = &v
	return s
}

func (s *ListLoadbalanceQuery) SetVaz(v string) *ListLoadbalanceQuery {
	s.Vaz = &v
	return s
}

func (s *ListLoadbalanceQuery) SetQueryType(v string) *ListLoadbalanceQuery {
	s.QueryType = &v
	return s
}

func (s *ListLoadbalanceQuery) SetIpVersion(v ListLoadbalanceQueryIpVersionEnum) *ListLoadbalanceQuery {
	s.IpVersion = &v
	return s
}

func (s *ListLoadbalanceQuery) SetIsFree(v bool) *ListLoadbalanceQuery {
	s.IsFree = &v
	return s
}

func (s *ListLoadbalanceQuery) SetShowVirtualLoadBalancers(v bool) *ListLoadbalanceQuery {
	s.ShowVirtualLoadBalancers = &v
	return s
}

func (s *ListLoadbalanceQuery) SetRouterId(v string) *ListLoadbalanceQuery {
	s.RouterId = &v
	return s
}

func (s *ListLoadbalanceQuery) SetIsDescribeTags(v bool) *ListLoadbalanceQuery {
	s.IsDescribeTags = &v
	return s
}

func (s *ListLoadbalanceQuery) SetFipBind(v bool) *ListLoadbalanceQuery {
	s.FipBind = &v
	return s
}

func (s *ListLoadbalanceQuery) SetNetworkId(v string) *ListLoadbalanceQuery {
	s.NetworkId = &v
	return s
}

func (s *ListLoadbalanceQuery) SetPage(v int32) *ListLoadbalanceQuery {
	s.Page = &v
	return s
}

func (s *ListLoadbalanceQuery) SetQueryWord(v string) *ListLoadbalanceQuery {
	s.QueryWord = &v
	return s
}

func (s *ListLoadbalanceQuery) SetAdminStateUp(v bool) *ListLoadbalanceQuery {
	s.AdminStateUp = &v
	return s
}

func (s *ListLoadbalanceQuery) SetTagKey(v string) *ListLoadbalanceQuery {
	s.TagKey = &v
	return s
}


type ListLoadbalanceQueryBuilder struct {
	s *ListLoadbalanceQuery
}

func NewListLoadbalanceQueryBuilder() *ListLoadbalanceQueryBuilder {
	s := &ListLoadbalanceQuery{}
	b := &ListLoadbalanceQueryBuilder{s: s}
	return b
}

func (b *ListLoadbalanceQueryBuilder) Visible(v bool) *ListLoadbalanceQueryBuilder {
	b.s.Visible = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) TagValue(v string) *ListLoadbalanceQueryBuilder {
	b.s.TagValue = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) IsMain(v bool) *ListLoadbalanceQueryBuilder {
	b.s.IsMain = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) JoinProductType(v ListLoadbalanceQueryJoinProductTypeEnum) *ListLoadbalanceQueryBuilder {
	b.s.JoinProductType = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) TagIds(v []string) *ListLoadbalanceQueryBuilder {
	b.s.TagIds = v
	return b
}

func (b *ListLoadbalanceQueryBuilder) PageSize(v int32) *ListLoadbalanceQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) OrderBy(v string) *ListLoadbalanceQueryBuilder {
	b.s.OrderBy = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) LoadBalanceId(v string) *ListLoadbalanceQueryBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) Vaz(v string) *ListLoadbalanceQueryBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) QueryType(v string) *ListLoadbalanceQueryBuilder {
	b.s.QueryType = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) IpVersion(v ListLoadbalanceQueryIpVersionEnum) *ListLoadbalanceQueryBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) IsFree(v bool) *ListLoadbalanceQueryBuilder {
	b.s.IsFree = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) ShowVirtualLoadBalancers(v bool) *ListLoadbalanceQueryBuilder {
	b.s.ShowVirtualLoadBalancers = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) RouterId(v string) *ListLoadbalanceQueryBuilder {
	b.s.RouterId = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) IsDescribeTags(v bool) *ListLoadbalanceQueryBuilder {
	b.s.IsDescribeTags = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) FipBind(v bool) *ListLoadbalanceQueryBuilder {
	b.s.FipBind = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) NetworkId(v string) *ListLoadbalanceQueryBuilder {
	b.s.NetworkId = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) Page(v int32) *ListLoadbalanceQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) QueryWord(v string) *ListLoadbalanceQueryBuilder {
	b.s.QueryWord = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) AdminStateUp(v bool) *ListLoadbalanceQueryBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) TagKey(v string) *ListLoadbalanceQueryBuilder {
	b.s.TagKey = &v
	return b
}

func (b *ListLoadbalanceQueryBuilder) Build() *ListLoadbalanceQuery {
	return b.s
}


