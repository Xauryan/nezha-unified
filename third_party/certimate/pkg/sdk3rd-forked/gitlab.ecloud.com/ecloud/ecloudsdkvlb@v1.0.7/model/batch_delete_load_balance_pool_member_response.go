// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BatchDeleteLoadBalancePoolMemberResponseStateEnum string

// List of State
const (
    BatchDeleteLoadBalancePoolMemberResponseStateEnumOk BatchDeleteLoadBalancePoolMemberResponseStateEnum = "OK"
    BatchDeleteLoadBalancePoolMemberResponseStateEnumError BatchDeleteLoadBalancePoolMemberResponseStateEnum = "ERROR"
    BatchDeleteLoadBalancePoolMemberResponseStateEnumException BatchDeleteLoadBalancePoolMemberResponseStateEnum = "EXCEPTION"
    BatchDeleteLoadBalancePoolMemberResponseStateEnumAlarm BatchDeleteLoadBalancePoolMemberResponseStateEnum = "ALARM"
    BatchDeleteLoadBalancePoolMemberResponseStateEnumForbidden BatchDeleteLoadBalancePoolMemberResponseStateEnum = "FORBIDDEN"
)

type BatchDeleteLoadBalancePoolMemberResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *BatchDeleteLoadBalancePoolMemberResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s BatchDeleteLoadBalancePoolMemberResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteLoadBalancePoolMemberResponse) GoString() string {
	return s.String()
}

func (s BatchDeleteLoadBalancePoolMemberResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetAsyncKey(v string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.AsyncKey = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetRequestId(v string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetErrorMessage(v string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetErrorCode(v string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetState(v BatchDeleteLoadBalancePoolMemberResponseStateEnum) *BatchDeleteLoadBalancePoolMemberResponse {
	s.State = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetBody(v string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.Body = &v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberResponse) SetErrorParams(v []string) *BatchDeleteLoadBalancePoolMemberResponse {
	s.ErrorParams = v
	return s
}


type BatchDeleteLoadBalancePoolMemberResponseBuilder struct {
	s *BatchDeleteLoadBalancePoolMemberResponse
}

func NewBatchDeleteLoadBalancePoolMemberResponseBuilder() *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	s := &BatchDeleteLoadBalancePoolMemberResponse{}
	b := &BatchDeleteLoadBalancePoolMemberResponseBuilder{s: s}
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) AsyncKey(v string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) RequestId(v string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) ErrorMessage(v string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) ErrorCode(v string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) State(v BatchDeleteLoadBalancePoolMemberResponseStateEnum) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) Body(v string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) ErrorParams(v []string) *BatchDeleteLoadBalancePoolMemberResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberResponseBuilder) Build() *BatchDeleteLoadBalancePoolMemberResponse {
	return b.s
}


