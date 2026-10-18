// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadbalanceCertificationBody struct {
    position.Body
    // 证书名称，5~64个字符，以英文大小写字母开头，可包含数字、下划线（_）和连字符（-）
	Name *string `json:"name,omitempty"`
    // 证书描述
	Description *string `json:"description,omitempty"`
    // 证书ID
	Id *string `json:"id,omitempty"`
}

func (s UpdateLoadbalanceCertificationBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceCertificationBody) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceCertificationBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceCertificationBody) SetName(v string) *UpdateLoadbalanceCertificationBody {
	s.Name = &v
	return s
}

func (s *UpdateLoadbalanceCertificationBody) SetDescription(v string) *UpdateLoadbalanceCertificationBody {
	s.Description = &v
	return s
}

func (s *UpdateLoadbalanceCertificationBody) SetId(v string) *UpdateLoadbalanceCertificationBody {
	s.Id = &v
	return s
}


type UpdateLoadbalanceCertificationBodyBuilder struct {
	s *UpdateLoadbalanceCertificationBody
}

func NewUpdateLoadbalanceCertificationBodyBuilder() *UpdateLoadbalanceCertificationBodyBuilder {
	s := &UpdateLoadbalanceCertificationBody{}
	b := &UpdateLoadbalanceCertificationBodyBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceCertificationBodyBuilder) Name(v string) *UpdateLoadbalanceCertificationBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateLoadbalanceCertificationBodyBuilder) Description(v string) *UpdateLoadbalanceCertificationBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateLoadbalanceCertificationBodyBuilder) Id(v string) *UpdateLoadbalanceCertificationBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateLoadbalanceCertificationBodyBuilder) Build() *UpdateLoadbalanceCertificationBody {
	return b.s
}


