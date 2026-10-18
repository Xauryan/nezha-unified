// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupDetailResponseLoadbalancers struct {

    // 负载均衡ID信息
	Id *string `json:"id,omitempty"`
}

func (s ListControlGroupDetailResponseLoadbalancers) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailResponseLoadbalancers) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailResponseLoadbalancers) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailResponseLoadbalancers) SetId(v string) *ListControlGroupDetailResponseLoadbalancers {
	s.Id = &v
	return s
}


type ListControlGroupDetailResponseLoadbalancersBuilder struct {
	s *ListControlGroupDetailResponseLoadbalancers
}

func NewListControlGroupDetailResponseLoadbalancersBuilder() *ListControlGroupDetailResponseLoadbalancersBuilder {
	s := &ListControlGroupDetailResponseLoadbalancers{}
	b := &ListControlGroupDetailResponseLoadbalancersBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailResponseLoadbalancersBuilder) Id(v string) *ListControlGroupDetailResponseLoadbalancersBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupDetailResponseLoadbalancersBuilder) Build() *ListControlGroupDetailResponseLoadbalancers {
	return b.s
}


