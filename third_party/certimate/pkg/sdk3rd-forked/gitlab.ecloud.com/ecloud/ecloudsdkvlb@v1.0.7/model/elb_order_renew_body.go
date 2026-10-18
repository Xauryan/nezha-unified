// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderRenewBody struct {
    position.Body
    // 订购时长,增加的时长，单位为原来的周期单位。计费周期为month时，取值为1~60月，计费周期为year时，取值为[年数*12月]，最多5年
	Duration *int32 `json:"duration,omitempty"`
    // 弹性负载均衡的ID
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s ElbOrderRenewBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderRenewBody) GoString() string {
	return s.String()
}

func (s ElbOrderRenewBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderRenewBody) SetDuration(v int32) *ElbOrderRenewBody {
	s.Duration = &v
	return s
}

func (s *ElbOrderRenewBody) SetInstanceId(v string) *ElbOrderRenewBody {
	s.InstanceId = &v
	return s
}


type ElbOrderRenewBodyBuilder struct {
	s *ElbOrderRenewBody
}

func NewElbOrderRenewBodyBuilder() *ElbOrderRenewBodyBuilder {
	s := &ElbOrderRenewBody{}
	b := &ElbOrderRenewBodyBuilder{s: s}
	return b
}

func (b *ElbOrderRenewBodyBuilder) Duration(v int32) *ElbOrderRenewBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *ElbOrderRenewBodyBuilder) InstanceId(v string) *ElbOrderRenewBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *ElbOrderRenewBodyBuilder) Build() *ElbOrderRenewBody {
	return b.s
}


