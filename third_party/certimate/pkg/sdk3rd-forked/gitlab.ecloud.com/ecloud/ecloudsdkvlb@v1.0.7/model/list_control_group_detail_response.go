// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListControlGroupDetailResponseStateEnum string

// List of State
const (
    ListControlGroupDetailResponseStateEnumOk ListControlGroupDetailResponseStateEnum = "OK"
    ListControlGroupDetailResponseStateEnumError ListControlGroupDetailResponseStateEnum = "ERROR"
    ListControlGroupDetailResponseStateEnumException ListControlGroupDetailResponseStateEnum = "EXCEPTION"
    ListControlGroupDetailResponseStateEnumAlarm ListControlGroupDetailResponseStateEnum = "ALARM"
    ListControlGroupDetailResponseStateEnumForbidden ListControlGroupDetailResponseStateEnum = "FORBIDDEN"
)

type ListControlGroupDetailResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *ListControlGroupDetailResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListControlGroupDetailResponseBody `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s ListControlGroupDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailResponse) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailResponse) SetAsyncKey(v string) *ListControlGroupDetailResponse {
	s.AsyncKey = &v
	return s
}

func (s *ListControlGroupDetailResponse) SetRequestId(v string) *ListControlGroupDetailResponse {
	s.RequestId = &v
	return s
}

func (s *ListControlGroupDetailResponse) SetErrorMessage(v string) *ListControlGroupDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListControlGroupDetailResponse) SetErrorCode(v string) *ListControlGroupDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListControlGroupDetailResponse) SetState(v ListControlGroupDetailResponseStateEnum) *ListControlGroupDetailResponse {
	s.State = &v
	return s
}

func (s *ListControlGroupDetailResponse) SetBody(v *ListControlGroupDetailResponseBody) *ListControlGroupDetailResponse {
	s.Body = v
	return s
}

func (s *ListControlGroupDetailResponse) SetErrorParams(v []string) *ListControlGroupDetailResponse {
	s.ErrorParams = v
	return s
}


type ListControlGroupDetailResponseBuilder struct {
	s *ListControlGroupDetailResponse
}

func NewListControlGroupDetailResponseBuilder() *ListControlGroupDetailResponseBuilder {
	s := &ListControlGroupDetailResponse{}
	b := &ListControlGroupDetailResponseBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailResponseBuilder) AsyncKey(v string) *ListControlGroupDetailResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) RequestId(v string) *ListControlGroupDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) ErrorMessage(v string) *ListControlGroupDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) ErrorCode(v string) *ListControlGroupDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) State(v ListControlGroupDetailResponseStateEnum) *ListControlGroupDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) Body(v *ListControlGroupDetailResponseBody) *ListControlGroupDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) ErrorParams(v []string) *ListControlGroupDetailResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *ListControlGroupDetailResponseBuilder) Build() *ListControlGroupDetailResponse {
	return b.s
}


