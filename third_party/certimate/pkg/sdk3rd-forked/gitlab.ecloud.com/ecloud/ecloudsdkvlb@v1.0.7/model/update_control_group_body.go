// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateControlGroupBody struct {
    position.Body
    // 访问控制组名称(description与name字段至少传一个)
	Name *string `json:"name,omitempty"`
    // 访问控制组描述(description与name字段至少传一个)
	Description *string `json:"description,omitempty"`
    // 访问控制组ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateControlGroupBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupBody) GoString() string {
	return s.String()
}

func (s UpdateControlGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupBody) SetName(v string) *UpdateControlGroupBody {
	s.Name = &v
	return s
}

func (s *UpdateControlGroupBody) SetDescription(v string) *UpdateControlGroupBody {
	s.Description = &v
	return s
}

func (s *UpdateControlGroupBody) SetId(v string) *UpdateControlGroupBody {
	s.Id = &v
	return s
}


type UpdateControlGroupBodyBuilder struct {
	s *UpdateControlGroupBody
}

func NewUpdateControlGroupBodyBuilder() *UpdateControlGroupBodyBuilder {
	s := &UpdateControlGroupBody{}
	b := &UpdateControlGroupBodyBuilder{s: s}
	return b
}

func (b *UpdateControlGroupBodyBuilder) Name(v string) *UpdateControlGroupBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateControlGroupBodyBuilder) Description(v string) *UpdateControlGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateControlGroupBodyBuilder) Id(v string) *UpdateControlGroupBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateControlGroupBodyBuilder) Build() *UpdateControlGroupBody {
	return b.s
}


