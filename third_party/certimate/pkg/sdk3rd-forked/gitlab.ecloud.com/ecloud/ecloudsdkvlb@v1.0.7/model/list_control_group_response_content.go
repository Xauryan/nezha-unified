// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupResponseContent struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 绑定监听器数量
	BindedListeners *int32 `json:"bindedListeners,omitempty"`
    // 是否被绑定
	IsBind *bool `json:"isBind,omitempty"`
    // 是否删除
	IsDelete *bool `json:"isDelete,omitempty"`
    // 用户ID
	Proposer *string `json:"proposer,omitempty"`
    // 是否能被监听器绑定
	EnableToBindListener *bool `json:"enableToBindListener,omitempty"`
    // 访问控制组描述
	Description *string `json:"description,omitempty"`
    // 包含的IP数量
	BindedIps *int32 `json:"bindedIps,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
    // 绑定监听器信息
	ListenerList *[]ListControlGroupResponseListenerList `json:"listenerList,omitempty"`
    // 客户ID
	CustomerId *string `json:"customerId,omitempty"`
    // 访问控制组名称
	Name *string `json:"name,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 修改人
	ModifiedBy *string `json:"modifiedBy,omitempty"`
    // 访问控制组ID
	Id *string `json:"id,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 包含的IP信息
	IpList []string `json:"ipList,omitempty"`
}

func (s ListControlGroupResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupResponseContent) GoString() string {
	return s.String()
}

func (s ListControlGroupResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupResponseContent) SetModifiedTime(v string) *ListControlGroupResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListControlGroupResponseContent) SetBindedListeners(v int32) *ListControlGroupResponseContent {
	s.BindedListeners = &v
	return s
}

func (s *ListControlGroupResponseContent) SetIsBind(v bool) *ListControlGroupResponseContent {
	s.IsBind = &v
	return s
}

func (s *ListControlGroupResponseContent) SetIsDelete(v bool) *ListControlGroupResponseContent {
	s.IsDelete = &v
	return s
}

func (s *ListControlGroupResponseContent) SetProposer(v string) *ListControlGroupResponseContent {
	s.Proposer = &v
	return s
}

func (s *ListControlGroupResponseContent) SetEnableToBindListener(v bool) *ListControlGroupResponseContent {
	s.EnableToBindListener = &v
	return s
}

func (s *ListControlGroupResponseContent) SetDescription(v string) *ListControlGroupResponseContent {
	s.Description = &v
	return s
}

func (s *ListControlGroupResponseContent) SetBindedIps(v int32) *ListControlGroupResponseContent {
	s.BindedIps = &v
	return s
}

func (s *ListControlGroupResponseContent) SetVaz(v string) *ListControlGroupResponseContent {
	s.Vaz = &v
	return s
}

func (s *ListControlGroupResponseContent) SetListenerList(v []ListControlGroupResponseListenerList) *ListControlGroupResponseContent {
	s.ListenerList = &v
	return s
}

func (s *ListControlGroupResponseContent) SetCustomerId(v string) *ListControlGroupResponseContent {
	s.CustomerId = &v
	return s
}

func (s *ListControlGroupResponseContent) SetName(v string) *ListControlGroupResponseContent {
	s.Name = &v
	return s
}

func (s *ListControlGroupResponseContent) SetCreatedTime(v string) *ListControlGroupResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListControlGroupResponseContent) SetModifiedBy(v string) *ListControlGroupResponseContent {
	s.ModifiedBy = &v
	return s
}

func (s *ListControlGroupResponseContent) SetId(v string) *ListControlGroupResponseContent {
	s.Id = &v
	return s
}

func (s *ListControlGroupResponseContent) SetRegion(v string) *ListControlGroupResponseContent {
	s.Region = &v
	return s
}

func (s *ListControlGroupResponseContent) SetIpList(v []string) *ListControlGroupResponseContent {
	s.IpList = v
	return s
}


type ListControlGroupResponseContentBuilder struct {
	s *ListControlGroupResponseContent
}

func NewListControlGroupResponseContentBuilder() *ListControlGroupResponseContentBuilder {
	s := &ListControlGroupResponseContent{}
	b := &ListControlGroupResponseContentBuilder{s: s}
	return b
}

func (b *ListControlGroupResponseContentBuilder) ModifiedTime(v string) *ListControlGroupResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) BindedListeners(v int32) *ListControlGroupResponseContentBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) IsBind(v bool) *ListControlGroupResponseContentBuilder {
	b.s.IsBind = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) IsDelete(v bool) *ListControlGroupResponseContentBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Proposer(v string) *ListControlGroupResponseContentBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) EnableToBindListener(v bool) *ListControlGroupResponseContentBuilder {
	b.s.EnableToBindListener = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Description(v string) *ListControlGroupResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) BindedIps(v int32) *ListControlGroupResponseContentBuilder {
	b.s.BindedIps = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Vaz(v string) *ListControlGroupResponseContentBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) ListenerList(v []ListControlGroupResponseListenerList) *ListControlGroupResponseContentBuilder {
	b.s.ListenerList = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) CustomerId(v string) *ListControlGroupResponseContentBuilder {
	b.s.CustomerId = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Name(v string) *ListControlGroupResponseContentBuilder {
	b.s.Name = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) CreatedTime(v string) *ListControlGroupResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) ModifiedBy(v string) *ListControlGroupResponseContentBuilder {
	b.s.ModifiedBy = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Id(v string) *ListControlGroupResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Region(v string) *ListControlGroupResponseContentBuilder {
	b.s.Region = &v
	return b
}

func (b *ListControlGroupResponseContentBuilder) IpList(v []string) *ListControlGroupResponseContentBuilder {
	b.s.IpList = v
	return b
}

func (b *ListControlGroupResponseContentBuilder) Build() *ListControlGroupResponseContent {
	return b.s
}


