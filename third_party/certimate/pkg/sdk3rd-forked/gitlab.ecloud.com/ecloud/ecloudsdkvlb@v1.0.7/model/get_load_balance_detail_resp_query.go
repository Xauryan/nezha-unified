// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceDetailRespQuery struct {
    position.Query
    // 用户是否可见
	Visible *bool `json:"visible,omitempty"`
}

func (s GetLoadBalanceDetailRespQuery) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespQuery) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespQuery) SetVisible(v bool) *GetLoadBalanceDetailRespQuery {
	s.Visible = &v
	return s
}


type GetLoadBalanceDetailRespQueryBuilder struct {
	s *GetLoadBalanceDetailRespQuery
}

func NewGetLoadBalanceDetailRespQueryBuilder() *GetLoadBalanceDetailRespQueryBuilder {
	s := &GetLoadBalanceDetailRespQuery{}
	b := &GetLoadBalanceDetailRespQueryBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespQueryBuilder) Visible(v bool) *GetLoadBalanceDetailRespQueryBuilder {
	b.s.Visible = &v
	return b
}

func (b *GetLoadBalanceDetailRespQueryBuilder) Build() *GetLoadBalanceDetailRespQuery {
	return b.s
}


