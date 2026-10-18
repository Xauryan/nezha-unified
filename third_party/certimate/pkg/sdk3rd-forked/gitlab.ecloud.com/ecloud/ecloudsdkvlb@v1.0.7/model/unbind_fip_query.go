// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnbindFipQuery struct {
    position.Query
    // 公网IP ID
	IpId *string `json:"ipId,omitempty"`
}

func (s UnbindFipQuery) String() string {
	return utils.Beautify(s)
}

func (s UnbindFipQuery) GoString() string {
	return s.String()
}

func (s UnbindFipQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindFipQuery) SetIpId(v string) *UnbindFipQuery {
	s.IpId = &v
	return s
}


type UnbindFipQueryBuilder struct {
	s *UnbindFipQuery
}

func NewUnbindFipQueryBuilder() *UnbindFipQueryBuilder {
	s := &UnbindFipQuery{}
	b := &UnbindFipQueryBuilder{s: s}
	return b
}

func (b *UnbindFipQueryBuilder) IpId(v string) *UnbindFipQueryBuilder {
	b.s.IpId = &v
	return b
}

func (b *UnbindFipQueryBuilder) Build() *UnbindFipQuery {
	return b.s
}


