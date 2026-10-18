// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListTLSCustomizeProtocolRespQuery struct {
    position.Query
    // 是否是默认策略，默认为false
	IsDefault *bool `json:"isDefault,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	Size *int32 `json:"size,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 虚拟可用区
	Vaz *string `json:"vaz,omitempty"`
}

func (s ListTLSCustomizeProtocolRespQuery) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespQuery) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespQuery) SetIsDefault(v bool) *ListTLSCustomizeProtocolRespQuery {
	s.IsDefault = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespQuery) SetSize(v int32) *ListTLSCustomizeProtocolRespQuery {
	s.Size = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespQuery) SetPage(v int32) *ListTLSCustomizeProtocolRespQuery {
	s.Page = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespQuery) SetVaz(v string) *ListTLSCustomizeProtocolRespQuery {
	s.Vaz = &v
	return s
}


type ListTLSCustomizeProtocolRespQueryBuilder struct {
	s *ListTLSCustomizeProtocolRespQuery
}

func NewListTLSCustomizeProtocolRespQueryBuilder() *ListTLSCustomizeProtocolRespQueryBuilder {
	s := &ListTLSCustomizeProtocolRespQuery{}
	b := &ListTLSCustomizeProtocolRespQueryBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespQueryBuilder) IsDefault(v bool) *ListTLSCustomizeProtocolRespQueryBuilder {
	b.s.IsDefault = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespQueryBuilder) Size(v int32) *ListTLSCustomizeProtocolRespQueryBuilder {
	b.s.Size = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespQueryBuilder) Page(v int32) *ListTLSCustomizeProtocolRespQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespQueryBuilder) Vaz(v string) *ListTLSCustomizeProtocolRespQueryBuilder {
	b.s.Vaz = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespQueryBuilder) Build() *ListTLSCustomizeProtocolRespQuery {
	return b.s
}


