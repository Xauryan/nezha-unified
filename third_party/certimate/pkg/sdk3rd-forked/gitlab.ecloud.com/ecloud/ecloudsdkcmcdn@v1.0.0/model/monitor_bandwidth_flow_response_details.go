// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MonitorBandwidthFlowResponseDetails struct {

    // 区域，取值为32个全国地区
	Area *string `json:"area,omitempty"`
    // 用量粒度时间
	DateTime *string `json:"dateTime,omitempty"`
    // 5分钟粒度时间内请命中数
	RequestHit *int64 `json:"requestHit,omitempty"`
    // 5分钟粒度时间内请求总数
	RequestNum *int64 `json:"requestNum,omitempty"`
    // 用量，流量KB
	FlowValue *float64 `json:"flowValue,omitempty"`
    // 用量，带宽Kbps
	BandwidthValue *float64 `json:"bandwidthValue,omitempty"`
    // 字节命中率%
	ByteHitrate *float64 `json:"byteHitrate,omitempty"`
    // 请求命中率，命中率=命中数/请求数，单位%
	RequestHitrate *float64 `json:"requestHitrate,omitempty"`
}

func (s MonitorBandwidthFlowResponseDetails) String() string {
	return utils.Beautify(s)
}

func (s MonitorBandwidthFlowResponseDetails) GoString() string {
	return s.String()
}

func (s MonitorBandwidthFlowResponseDetails) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MonitorBandwidthFlowResponseDetails) SetArea(v string) *MonitorBandwidthFlowResponseDetails {
	s.Area = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetDateTime(v string) *MonitorBandwidthFlowResponseDetails {
	s.DateTime = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetRequestHit(v int64) *MonitorBandwidthFlowResponseDetails {
	s.RequestHit = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetRequestNum(v int64) *MonitorBandwidthFlowResponseDetails {
	s.RequestNum = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetFlowValue(v float64) *MonitorBandwidthFlowResponseDetails {
	s.FlowValue = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetBandwidthValue(v float64) *MonitorBandwidthFlowResponseDetails {
	s.BandwidthValue = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetByteHitrate(v float64) *MonitorBandwidthFlowResponseDetails {
	s.ByteHitrate = &v
	return s
}

func (s *MonitorBandwidthFlowResponseDetails) SetRequestHitrate(v float64) *MonitorBandwidthFlowResponseDetails {
	s.RequestHitrate = &v
	return s
}


type MonitorBandwidthFlowResponseDetailsBuilder struct {
	s *MonitorBandwidthFlowResponseDetails
}

func NewMonitorBandwidthFlowResponseDetailsBuilder() *MonitorBandwidthFlowResponseDetailsBuilder {
	s := &MonitorBandwidthFlowResponseDetails{}
	b := &MonitorBandwidthFlowResponseDetailsBuilder{s: s}
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) Area(v string) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.Area = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) DateTime(v string) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.DateTime = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) RequestHit(v int64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.RequestHit = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) RequestNum(v int64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.RequestNum = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) FlowValue(v float64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.FlowValue = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) BandwidthValue(v float64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.BandwidthValue = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) ByteHitrate(v float64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.ByteHitrate = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) RequestHitrate(v float64) *MonitorBandwidthFlowResponseDetailsBuilder {
	b.s.RequestHitrate = &v
	return b
}

func (b *MonitorBandwidthFlowResponseDetailsBuilder) Build() *MonitorBandwidthFlowResponseDetails {
	return b.s
}


