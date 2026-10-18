// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MonitorBandwidthFlowRequest struct {


	MonitorBandwidthFlowBody *MonitorBandwidthFlowBody `json:"monitorBandwidthFlowBody,omitempty"`
}

func (s MonitorBandwidthFlowRequest) String() string {
	return utils.Beautify(s)
}

func (s MonitorBandwidthFlowRequest) GoString() string {
	return s.String()
}

func (s MonitorBandwidthFlowRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MonitorBandwidthFlowRequest) SetMonitorBandwidthFlowBody(v *MonitorBandwidthFlowBody) *MonitorBandwidthFlowRequest {
	s.MonitorBandwidthFlowBody = v
	return s
}


type MonitorBandwidthFlowRequestBuilder struct {
	s *MonitorBandwidthFlowRequest
}

func NewMonitorBandwidthFlowRequestBuilder() *MonitorBandwidthFlowRequestBuilder {
	s := &MonitorBandwidthFlowRequest{}
	b := &MonitorBandwidthFlowRequestBuilder{s: s}
	return b
}

func (b *MonitorBandwidthFlowRequestBuilder) MonitorBandwidthFlowBody(v *MonitorBandwidthFlowBody) *MonitorBandwidthFlowRequestBuilder {
	b.s.MonitorBandwidthFlowBody = v
	return b
}

func (b *MonitorBandwidthFlowRequestBuilder) Build() *MonitorBandwidthFlowRequest {
	return b.s
}


