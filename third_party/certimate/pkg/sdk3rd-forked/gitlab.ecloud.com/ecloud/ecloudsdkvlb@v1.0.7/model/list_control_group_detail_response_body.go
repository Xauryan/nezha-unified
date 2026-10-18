// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupDetailResponseBody struct {

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
	ListenerList *[]ListControlGroupDetailResponseListenerList `json:"listenerList,omitempty"`
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

func (s ListControlGroupDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailResponseBody) SetModifiedTime(v string) *ListControlGroupDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetBindedListeners(v int32) *ListControlGroupDetailResponseBody {
	s.BindedListeners = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetIsBind(v bool) *ListControlGroupDetailResponseBody {
	s.IsBind = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetIsDelete(v bool) *ListControlGroupDetailResponseBody {
	s.IsDelete = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetProposer(v string) *ListControlGroupDetailResponseBody {
	s.Proposer = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetEnableToBindListener(v bool) *ListControlGroupDetailResponseBody {
	s.EnableToBindListener = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetDescription(v string) *ListControlGroupDetailResponseBody {
	s.Description = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetBindedIps(v int32) *ListControlGroupDetailResponseBody {
	s.BindedIps = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetVaz(v string) *ListControlGroupDetailResponseBody {
	s.Vaz = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetListenerList(v []ListControlGroupDetailResponseListenerList) *ListControlGroupDetailResponseBody {
	s.ListenerList = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetCustomerId(v string) *ListControlGroupDetailResponseBody {
	s.CustomerId = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetName(v string) *ListControlGroupDetailResponseBody {
	s.Name = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetCreatedTime(v string) *ListControlGroupDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetModifiedBy(v string) *ListControlGroupDetailResponseBody {
	s.ModifiedBy = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetId(v string) *ListControlGroupDetailResponseBody {
	s.Id = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetRegion(v string) *ListControlGroupDetailResponseBody {
	s.Region = &v
	return s
}

func (s *ListControlGroupDetailResponseBody) SetIpList(v []string) *ListControlGroupDetailResponseBody {
	s.IpList = v
	return s
}


type ListControlGroupDetailResponseBodyBuilder struct {
	s *ListControlGroupDetailResponseBody
}

func NewListControlGroupDetailResponseBodyBuilder() *ListControlGroupDetailResponseBodyBuilder {
	s := &ListControlGroupDetailResponseBody{}
	b := &ListControlGroupDetailResponseBodyBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) ModifiedTime(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) BindedListeners(v int32) *ListControlGroupDetailResponseBodyBuilder {
	b.s.BindedListeners = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) IsBind(v bool) *ListControlGroupDetailResponseBodyBuilder {
	b.s.IsBind = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) IsDelete(v bool) *ListControlGroupDetailResponseBodyBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Proposer(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) EnableToBindListener(v bool) *ListControlGroupDetailResponseBodyBuilder {
	b.s.EnableToBindListener = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Description(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) BindedIps(v int32) *ListControlGroupDetailResponseBodyBuilder {
	b.s.BindedIps = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Vaz(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) ListenerList(v []ListControlGroupDetailResponseListenerList) *ListControlGroupDetailResponseBodyBuilder {
	b.s.ListenerList = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) CustomerId(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.CustomerId = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Name(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) CreatedTime(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) ModifiedBy(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.ModifiedBy = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Id(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Region(v string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) IpList(v []string) *ListControlGroupDetailResponseBodyBuilder {
	b.s.IpList = v
	return b
}

func (b *ListControlGroupDetailResponseBodyBuilder) Build() *ListControlGroupDetailResponseBody {
	return b.s
}


