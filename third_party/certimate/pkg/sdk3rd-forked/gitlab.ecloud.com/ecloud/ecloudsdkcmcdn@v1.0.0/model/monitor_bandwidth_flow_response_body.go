// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MonitorBandwidthFlowResponseBody struct {


	Details *[]MonitorBandwidthFlowResponseDetails `json:"details,omitempty"`
    // 用量值
	Value *float64 `json:"value,omitempty"`
}

func (s MonitorBandwidthFlowResponseBody) String() string {
	return utils.Beautify(s)
}

func (s MonitorBandwidthFlowResponseBody) GoString() string {
	return s.String()
}

func (s MonitorBandwidthFlowResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MonitorBandwidthFlowResponseBody) SetDetails(v []MonitorBandwidthFlowResponseDetails) *MonitorBandwidthFlowResponseBody {
	s.Details = &v
	return s
}

func (s *MonitorBandwidthFlowResponseBody) SetValue(v float64) *MonitorBandwidthFlowResponseBody {
	s.Value = &v
	return s
}


type MonitorBandwidthFlowResponseBodyBuilder struct {
	s *MonitorBandwidthFlowResponseBody
}

func NewMonitorBandwidthFlowResponseBodyBuilder() *MonitorBandwidthFlowResponseBodyBuilder {
	s := &MonitorBandwidthFlowResponseBody{}
	b := &MonitorBandwidthFlowResponseBodyBuilder{s: s}
	return b
}

func (b *MonitorBandwidthFlowResponseBodyBuilder) Details(v []MonitorBandwidthFlowResponseDetails) *MonitorBandwidthFlowResponseBodyBuilder {
	b.s.Details = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBodyBuilder) Value(v float64) *MonitorBandwidthFlowResponseBodyBuilder {
	b.s.Value = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBodyBuilder) Build() *MonitorBandwidthFlowResponseBody {
	return b.s
}


