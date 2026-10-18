// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateL7PolicyBody struct {
    position.Body
    // 七层转发策略名称，为5~64位的英文、数字、下划线中划线等的组合
	Name *string `json:"name,omitempty"`
    // 后端服务组ID
	PoolId *string `json:"poolId,omitempty"`
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
    // 七层转发策略优先级
	Position *int32 `json:"position,omitempty"`
}

func (s AsyncUpdateL7PolicyBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateL7PolicyBody) GoString() string {
	return s.String()
}

func (s AsyncUpdateL7PolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateL7PolicyBody) SetName(v string) *AsyncUpdateL7PolicyBody {
	s.Name = &v
	return s
}

func (s *AsyncUpdateL7PolicyBody) SetPoolId(v string) *AsyncUpdateL7PolicyBody {
	s.PoolId = &v
	return s
}

func (s *AsyncUpdateL7PolicyBody) SetL7PolicyId(v string) *AsyncUpdateL7PolicyBody {
	s.L7PolicyId = &v
	return s
}

func (s *AsyncUpdateL7PolicyBody) SetPosition(v int32) *AsyncUpdateL7PolicyBody {
	s.Position = &v
	return s
}


type AsyncUpdateL7PolicyBodyBuilder struct {
	s *AsyncUpdateL7PolicyBody
}

func NewAsyncUpdateL7PolicyBodyBuilder() *AsyncUpdateL7PolicyBodyBuilder {
	s := &AsyncUpdateL7PolicyBody{}
	b := &AsyncUpdateL7PolicyBodyBuilder{s: s}
	return b
}

func (b *AsyncUpdateL7PolicyBodyBuilder) Name(v string) *AsyncUpdateL7PolicyBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *AsyncUpdateL7PolicyBodyBuilder) PoolId(v string) *AsyncUpdateL7PolicyBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncUpdateL7PolicyBodyBuilder) L7PolicyId(v string) *AsyncUpdateL7PolicyBodyBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *AsyncUpdateL7PolicyBodyBuilder) Position(v int32) *AsyncUpdateL7PolicyBodyBuilder {
	b.s.Position = &v
	return b
}

func (b *AsyncUpdateL7PolicyBodyBuilder) Build() *AsyncUpdateL7PolicyBody {
	return b.s
}


