// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchDeleteControlGroupResponseStateEnum string

// List of State
const (
    BatchDeleteControlGroupResponseStateEnumOk BatchDeleteControlGroupResponseStateEnum = "OK"
    BatchDeleteControlGroupResponseStateEnumError BatchDeleteControlGroupResponseStateEnum = "ERROR"
    BatchDeleteControlGroupResponseStateEnumException BatchDeleteControlGroupResponseStateEnum = "EXCEPTION"
    BatchDeleteControlGroupResponseStateEnumAlarm BatchDeleteControlGroupResponseStateEnum = "ALARM"
    BatchDeleteControlGroupResponseStateEnumForbidden BatchDeleteControlGroupResponseStateEnum = "FORBIDDEN"
)

type BatchDeleteControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *BatchDeleteControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s BatchDeleteControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteControlGroupResponse) GoString() string {
	return s.String()
}

func (s BatchDeleteControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteControlGroupResponse) SetAsyncKey(v string) *BatchDeleteControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetRequestId(v string) *BatchDeleteControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetErrorMessage(v string) *BatchDeleteControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetErrorCode(v string) *BatchDeleteControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetState(v BatchDeleteControlGroupResponseStateEnum) *BatchDeleteControlGroupResponse {
	s.State = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetBody(v string) *BatchDeleteControlGroupResponse {
	s.Body = &v
	return s
}

func (s *BatchDeleteControlGroupResponse) SetErrorParams(v []string) *BatchDeleteControlGroupResponse {
	s.ErrorParams = v
	return s
}


type BatchDeleteControlGroupResponseBuilder struct {
	s *BatchDeleteControlGroupResponse
}

func NewBatchDeleteControlGroupResponseBuilder() *BatchDeleteControlGroupResponseBuilder {
	s := &BatchDeleteControlGroupResponse{}
	b := &BatchDeleteControlGroupResponseBuilder{s: s}
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) AsyncKey(v string) *BatchDeleteControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) RequestId(v string) *BatchDeleteControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) ErrorMessage(v string) *BatchDeleteControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) ErrorCode(v string) *BatchDeleteControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) State(v BatchDeleteControlGroupResponseStateEnum) *BatchDeleteControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) Body(v string) *BatchDeleteControlGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) ErrorParams(v []string) *BatchDeleteControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *BatchDeleteControlGroupResponseBuilder) Build() *BatchDeleteControlGroupResponse {
	return b.s
}


