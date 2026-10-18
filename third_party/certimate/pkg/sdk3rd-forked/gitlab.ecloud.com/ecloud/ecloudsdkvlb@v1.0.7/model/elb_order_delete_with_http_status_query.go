// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteWithHttpStatusQuery struct {
    position.Query
    // 弹性负载均衡的ID
	LoadBalanceId *string `json:"loadBalanceId,omitempty"`
}

func (s ElbOrderDeleteWithHttpStatusQuery) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteWithHttpStatusQuery) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteWithHttpStatusQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteWithHttpStatusQuery) SetLoadBalanceId(v string) *ElbOrderDeleteWithHttpStatusQuery {
	s.LoadBalanceId = &v
	return s
}


type ElbOrderDeleteWithHttpStatusQueryBuilder struct {
	s *ElbOrderDeleteWithHttpStatusQuery
}

func NewElbOrderDeleteWithHttpStatusQueryBuilder() *ElbOrderDeleteWithHttpStatusQueryBuilder {
	s := &ElbOrderDeleteWithHttpStatusQuery{}
	b := &ElbOrderDeleteWithHttpStatusQueryBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteWithHttpStatusQueryBuilder) LoadBalanceId(v string) *ElbOrderDeleteWithHttpStatusQueryBuilder {
	b.s.LoadBalanceId = &v
	return b
}

func (b *ElbOrderDeleteWithHttpStatusQueryBuilder) Build() *ElbOrderDeleteWithHttpStatusQuery {
	return b.s
}


