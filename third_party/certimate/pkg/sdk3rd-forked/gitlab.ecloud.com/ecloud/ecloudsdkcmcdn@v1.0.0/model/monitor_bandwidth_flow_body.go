// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MonitorBandwidthFlowBody struct {
    position.Body
    // 域名id
	DomainIds []int32 `json:"domainIds,omitempty"`
    // 时间间隔，5:5分钟，60:小时，24:天
	Period *int32 `json:"period,omitempty"`
    // 日期，区分0今日，1昨日，2近7天，3近30天，7当月，8上月，4自定义
	DateType *int32 `json:"dateType,omitempty"`
    // 地区编码 全国:1001 北京:100 上海:210
	AreaIds []int32 `json:"areaIds,omitempty"`
    // 自定义，开始时间戳(毫秒)
	StartTime *int64 `json:"startTime,omitempty"`
    // 自定义，结束时间戳(毫秒)
	EndTime *int64 `json:"endTime,omitempty"`
    // 0带宽、1流量、2带宽和流量
	Type *int32 `json:"type,omitempty"`
}

func (s MonitorBandwidthFlowBody) String() string {
	return utils.Beautify(s)
}

func (s MonitorBandwidthFlowBody) GoString() string {
	return s.String()
}

func (s MonitorBandwidthFlowBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MonitorBandwidthFlowBody) SetDomainIds(v []int32) *MonitorBandwidthFlowBody {
	s.DomainIds = v
	return s
}

func (s *MonitorBandwidthFlowBody) SetPeriod(v int32) *MonitorBandwidthFlowBody {
	s.Period = &v
	return s
}

func (s *MonitorBandwidthFlowBody) SetDateType(v int32) *MonitorBandwidthFlowBody {
	s.DateType = &v
	return s
}

func (s *MonitorBandwidthFlowBody) SetAreaIds(v []int32) *MonitorBandwidthFlowBody {
	s.AreaIds = v
	return s
}

func (s *MonitorBandwidthFlowBody) SetStartTime(v int64) *MonitorBandwidthFlowBody {
	s.StartTime = &v
	return s
}

func (s *MonitorBandwidthFlowBody) SetEndTime(v int64) *MonitorBandwidthFlowBody {
	s.EndTime = &v
	return s
}

func (s *MonitorBandwidthFlowBody) SetType(v int32) *MonitorBandwidthFlowBody {
	s.Type = &v
	return s
}


type MonitorBandwidthFlowBodyBuilder struct {
	s *MonitorBandwidthFlowBody
}

func NewMonitorBandwidthFlowBodyBuilder() *MonitorBandwidthFlowBodyBuilder {
	s := &MonitorBandwidthFlowBody{}
	b := &MonitorBandwidthFlowBodyBuilder{s: s}
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) DomainIds(v []int32) *MonitorBandwidthFlowBodyBuilder {
	b.s.DomainIds = v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) Period(v int32) *MonitorBandwidthFlowBodyBuilder {
	b.s.Period = &v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) DateType(v int32) *MonitorBandwidthFlowBodyBuilder {
	b.s.DateType = &v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) AreaIds(v []int32) *MonitorBandwidthFlowBodyBuilder {
	b.s.AreaIds = v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) StartTime(v int64) *MonitorBandwidthFlowBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) EndTime(v int64) *MonitorBandwidthFlowBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) Type(v int32) *MonitorBandwidthFlowBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *MonitorBandwidthFlowBodyBuilder) Build() *MonitorBandwidthFlowBody {
	return b.s
}


