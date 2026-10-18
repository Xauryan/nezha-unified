// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetBatchHealthStatusRespResponseStateEnum string

// List of State
const (
    GetBatchHealthStatusRespResponseStateEnumOk GetBatchHealthStatusRespResponseStateEnum = "OK"
    GetBatchHealthStatusRespResponseStateEnumError GetBatchHealthStatusRespResponseStateEnum = "ERROR"
    GetBatchHealthStatusRespResponseStateEnumException GetBatchHealthStatusRespResponseStateEnum = "EXCEPTION"
    GetBatchHealthStatusRespResponseStateEnumAlarm GetBatchHealthStatusRespResponseStateEnum = "ALARM"
    GetBatchHealthStatusRespResponseStateEnumForbidden GetBatchHealthStatusRespResponseStateEnum = "FORBIDDEN"
)

type GetBatchHealthStatusRespResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *GetBatchHealthStatusRespResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]GetBatchHealthStatusRespResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetBatchHealthStatusRespResponse) String() string {
	return utils.Beautify(s)
}

func (s GetBatchHealthStatusRespResponse) GoString() string {
	return s.String()
}

func (s GetBatchHealthStatusRespResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBatchHealthStatusRespResponse) SetAsyncKey(v string) *GetBatchHealthStatusRespResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetRequestId(v string) *GetBatchHealthStatusRespResponse {
	s.RequestId = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetErrorMessage(v string) *GetBatchHealthStatusRespResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetErrorCode(v string) *GetBatchHealthStatusRespResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetState(v GetBatchHealthStatusRespResponseStateEnum) *GetBatchHealthStatusRespResponse {
	s.State = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetBody(v []GetBatchHealthStatusRespResponseBody) *GetBatchHealthStatusRespResponse {
	s.Body = &v
	return s
}

func (s *GetBatchHealthStatusRespResponse) SetErrorParams(v []string) *GetBatchHealthStatusRespResponse {
	s.ErrorParams = v
	return s
}


type GetBatchHealthStatusRespResponseBuilder struct {
	s *GetBatchHealthStatusRespResponse
}

func NewGetBatchHealthStatusRespResponseBuilder() *GetBatchHealthStatusRespResponseBuilder {
	s := &GetBatchHealthStatusRespResponse{}
	b := &GetBatchHealthStatusRespResponseBuilder{s: s}
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) AsyncKey(v string) *GetBatchHealthStatusRespResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) RequestId(v string) *GetBatchHealthStatusRespResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) ErrorMessage(v string) *GetBatchHealthStatusRespResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) ErrorCode(v string) *GetBatchHealthStatusRespResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) State(v GetBatchHealthStatusRespResponseStateEnum) *GetBatchHealthStatusRespResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) Body(v []GetBatchHealthStatusRespResponseBody) *GetBatchHealthStatusRespResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) ErrorParams(v []string) *GetBatchHealthStatusRespResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetBatchHealthStatusRespResponseBuilder) Build() *GetBatchHealthStatusRespResponse {
	return b.s
}


