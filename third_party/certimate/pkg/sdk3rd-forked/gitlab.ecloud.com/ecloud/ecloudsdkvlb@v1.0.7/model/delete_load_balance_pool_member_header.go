// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteLoadBalancePoolMemberHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s DeleteLoadBalancePoolMemberHeader) String() string {
	return utils.Beautify(s)
}

func (s DeleteLoadBalancePoolMemberHeader) GoString() string {
	return s.String()
}

func (s DeleteLoadBalancePoolMemberHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteLoadBalancePoolMemberHeader) SetRegion(v string) *DeleteLoadBalancePoolMemberHeader {
	s.Region = &v
	return s
}


type DeleteLoadBalancePoolMemberHeaderBuilder struct {
	s *DeleteLoadBalancePoolMemberHeader
}

func NewDeleteLoadBalancePoolMemberHeaderBuilder() *DeleteLoadBalancePoolMemberHeaderBuilder {
	s := &DeleteLoadBalancePoolMemberHeader{}
	b := &DeleteLoadBalancePoolMemberHeaderBuilder{s: s}
	return b
}

func (b *DeleteLoadBalancePoolMemberHeaderBuilder) Region(v string) *DeleteLoadBalancePoolMemberHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *DeleteLoadBalancePoolMemberHeaderBuilder) Build() *DeleteLoadBalancePoolMemberHeader {
	return b.s
}


