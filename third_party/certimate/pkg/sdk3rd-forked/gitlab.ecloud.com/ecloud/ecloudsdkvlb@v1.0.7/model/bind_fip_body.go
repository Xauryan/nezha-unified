// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindFipBodyTypeEnum string

// List of Type
const (
    BindFipBodyTypeEnumElb BindFipBodyTypeEnum = "elb"
)

type BindFipBody struct {
    position.Body
    // 要绑定的资源ID
	ResourceId *string `json:"resourceId,omitempty"`
    // 公网IP ID
	IpId *string `json:"ipId,omitempty"`
    // 网卡ID
	PortId *string `json:"portId,omitempty"`
    // 公网IP绑定类型,负载均衡:elb
	Type *BindFipBodyTypeEnum `json:"type,omitempty"`
}

func (s BindFipBody) String() string {
	return utils.Beautify(s)
}

func (s BindFipBody) GoString() string {
	return s.String()
}

func (s BindFipBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindFipBody) SetResourceId(v string) *BindFipBody {
	s.ResourceId = &v
	return s
}

func (s *BindFipBody) SetIpId(v string) *BindFipBody {
	s.IpId = &v
	return s
}

func (s *BindFipBody) SetPortId(v string) *BindFipBody {
	s.PortId = &v
	return s
}

func (s *BindFipBody) SetType(v BindFipBodyTypeEnum) *BindFipBody {
	s.Type = &v
	return s
}


type BindFipBodyBuilder struct {
	s *BindFipBody
}

func NewBindFipBodyBuilder() *BindFipBodyBuilder {
	s := &BindFipBody{}
	b := &BindFipBodyBuilder{s: s}
	return b
}

func (b *BindFipBodyBuilder) ResourceId(v string) *BindFipBodyBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *BindFipBodyBuilder) IpId(v string) *BindFipBodyBuilder {
	b.s.IpId = &v
	return b
}

func (b *BindFipBodyBuilder) PortId(v string) *BindFipBodyBuilder {
	b.s.PortId = &v
	return b
}

func (b *BindFipBodyBuilder) Type(v BindFipBodyTypeEnum) *BindFipBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *BindFipBodyBuilder) Build() *BindFipBody {
	return b.s
}


