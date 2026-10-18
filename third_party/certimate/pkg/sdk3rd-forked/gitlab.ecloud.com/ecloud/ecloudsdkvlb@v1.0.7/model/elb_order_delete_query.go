// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteQuery struct {
    position.Query
    // 弹性负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ElbOrderDeleteQuery) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteQuery) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteQuery) SetLoadBalanceId(v string) *ElbOrderDeleteQuery {
	s.LoadBalanceId = &v
	return s
}


type ElbOrderDeleteQueryBuilder struct {
	s *ElbOrderDeleteQuery
}

func NewElbOrderDeleteQueryBuilder() *ElbOrderDeleteQueryBuilder {
	s := &ElbOrderDeleteQuery{}
	b := &ElbOrderDeleteQueryBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteQueryBuilder) LoadBalanceId(v string) *ElbOrderDeleteQueryBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ElbOrderDeleteQueryBuilder) Build() *ElbOrderDeleteQuery {
	return b.s
}


