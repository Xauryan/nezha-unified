// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateTLSNameBody struct {
    position.Body
    // TLS安全策略名称
	Name *string `json:"name,omitempty"`
    // TLS安全策略ID
	Id *int64 `json:"id,omitempty"`
}

func (s UpdateTLSNameBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateTLSNameBody) GoString() string {
	return s.String()
}

func (s UpdateTLSNameBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateTLSNameBody) SetName(v string) *UpdateTLSNameBody {
	s.Name = &v
	return s
}

func (s *UpdateTLSNameBody) SetId(v int64) *UpdateTLSNameBody {
	s.Id = &v
	return s
}


type UpdateTLSNameBodyBuilder struct {
	s *UpdateTLSNameBody
}

func NewUpdateTLSNameBodyBuilder() *UpdateTLSNameBodyBuilder {
	s := &UpdateTLSNameBody{}
	b := &UpdateTLSNameBodyBuilder{s: s}
	return b
}

func (b *UpdateTLSNameBodyBuilder) Name(v string) *UpdateTLSNameBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateTLSNameBodyBuilder) Id(v int64) *UpdateTLSNameBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateTLSNameBodyBuilder) Build() *UpdateTLSNameBody {
	return b.s
}


