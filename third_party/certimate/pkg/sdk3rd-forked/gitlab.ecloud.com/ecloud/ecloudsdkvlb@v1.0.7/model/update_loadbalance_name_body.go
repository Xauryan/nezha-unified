// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadbalanceNameBody struct {
    position.Body
    // 负载均衡备注
	Notes *string `json:"notes,omitempty"`
    // 弹性负载均衡名称，5~64位的英文、数字、下划线、中划线等的组合
	Name *string `json:"name,omitempty"`
    // 弹性负载均衡ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadbalanceNameBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceNameBody) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceNameBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceNameBody) SetNotes(v string) *UpdateLoadbalanceNameBody {
	s.Notes = &v
	return s
}

func (s *UpdateLoadbalanceNameBody) SetName(v string) *UpdateLoadbalanceNameBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadbalanceNameBody) SetId(v string) *UpdateLoadbalanceNameBody {
	s.Id = &v
	return s
}


type UpdateLoadbalanceNameBodyBuilder struct {
	s *UpdateLoadbalanceNameBody
}

func NewUpdateLoadbalanceNameBodyBuilder() *UpdateLoadbalanceNameBodyBuilder {
	s := &UpdateLoadbalanceNameBody{}
	b := &UpdateLoadbalanceNameBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceNameBodyBuilder) Notes(v string) *UpdateLoadbalanceNameBodyBuilder {
	b.s.Notes = &v
	return b
}

func (b *UpdateLoadbalanceNameBodyBuilder) Name(v string) *UpdateLoadbalanceNameBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadbalanceNameBodyBuilder) Id(v string) *UpdateLoadbalanceNameBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadbalanceNameBodyBuilder) Build() *UpdateLoadbalanceNameBody {
	return b.s
}


