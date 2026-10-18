// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupResponseLoadbalancers struct {

    // 负载均衡ID
	Id *string `json:"id,omitempty"`
}

func (s ListControlGroupResponseLoadbalancers) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupResponseLoadbalancers) GoString() string {
	return s.String()
}

func (s ListControlGroupResponseLoadbalancers) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupResponseLoadbalancers) SetId(v string) *ListControlGroupResponseLoadbalancers {
	s.Id = &v
	return s
}


type ListControlGroupResponseLoadbalancersBuilder struct {
	s *ListControlGroupResponseLoadbalancers
}

func NewListControlGroupResponseLoadbalancersBuilder() *ListControlGroupResponseLoadbalancersBuilder {
	s := &ListControlGroupResponseLoadbalancers{}
	b := &ListControlGroupResponseLoadbalancersBuilder{s: s}
	return b
}

func (b *ListControlGroupResponseLoadbalancersBuilder) Id(v string) *ListControlGroupResponseLoadbalancersBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupResponseLoadbalancersBuilder) Build() *ListControlGroupResponseLoadbalancers {
	return b.s
}


