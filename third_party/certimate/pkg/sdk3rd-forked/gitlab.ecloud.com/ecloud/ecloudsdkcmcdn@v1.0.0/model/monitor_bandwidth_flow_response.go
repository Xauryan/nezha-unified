// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MonitorBandwidthFlowResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *MonitorBandwidthFlowResponseBody `json:"body,omitempty"`
}

func (s MonitorBandwidthFlowResponse) String() string {
	return utils.Beautify(s)
}

func (s MonitorBandwidthFlowResponse) GoString() string {
	return s.String()
}

func (s MonitorBandwidthFlowResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MonitorBandwidthFlowResponse) SetRequestId(v string) *MonitorBandwidthFlowResponse {
	s.RequestId = &v
	return s
}

func (s *MonitorBandwidthFlowResponse) SetErrorMessage(v string) *MonitorBandwidthFlowResponse {
	s.ErrorMessage = &v
	return s
}

func (s *MonitorBandwidthFlowResponse) SetErrorCode(v string) *MonitorBandwidthFlowResponse {
	s.ErrorCode = &v
	return s
}

func (s *MonitorBandwidthFlowResponse) SetState(v string) *MonitorBandwidthFlowResponse {
	s.State = &v
	return s
}

func (s *MonitorBandwidthFlowResponse) SetBody(v *MonitorBandwidthFlowResponseBody) *MonitorBandwidthFlowResponse {
	s.Body = v
	return s
}


type MonitorBandwidthFlowResponseBuilder struct {
	s *MonitorBandwidthFlowResponse
}

func NewMonitorBandwidthFlowResponseBuilder() *MonitorBandwidthFlowResponseBuilder {
	s := &MonitorBandwidthFlowResponse{}
	b := &MonitorBandwidthFlowResponseBuilder{s: s}
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) RequestId(v string) *MonitorBandwidthFlowResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) ErrorMessage(v string) *MonitorBandwidthFlowResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) ErrorCode(v string) *MonitorBandwidthFlowResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) State(v string) *MonitorBandwidthFlowResponseBuilder {
	b.s.State = &v
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) Body(v *MonitorBandwidthFlowResponseBody) *MonitorBandwidthFlowResponseBuilder {
	b.s.Body = v
	return b
}

func (b *MonitorBandwidthFlowResponseBuilder) Build() *MonitorBandwidthFlowResponse {
	return b.s
}


