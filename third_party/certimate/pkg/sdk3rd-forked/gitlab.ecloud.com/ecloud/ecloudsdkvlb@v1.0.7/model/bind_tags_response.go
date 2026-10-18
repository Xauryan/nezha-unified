// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindTagsResponseStateEnum string

// List of State
const (
    BindTagsResponseStateEnumOk BindTagsResponseStateEnum = "OK"
    BindTagsResponseStateEnumError BindTagsResponseStateEnum = "ERROR"
    BindTagsResponseStateEnumException BindTagsResponseStateEnum = "EXCEPTION"
    BindTagsResponseStateEnumAlarm BindTagsResponseStateEnum = "ALARM"
    BindTagsResponseStateEnumForbidden BindTagsResponseStateEnum = "FORBIDDEN"
)

type BindTagsResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *BindTagsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s BindTagsResponse) String() string {
	return utils.Beautify(s)
}

func (s BindTagsResponse) GoString() string {
	return s.String()
}

func (s BindTagsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindTagsResponse) SetAsyncKey(v string) *BindTagsResponse {
	s.AsyncKey = &v
	return s
}

func (s *BindTagsResponse) SetRequestId(v string) *BindTagsResponse {
	s.RequestId = &v
	return s
}

func (s *BindTagsResponse) SetErrorMessage(v string) *BindTagsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BindTagsResponse) SetErrorCode(v string) *BindTagsResponse {
	s.ErrorCode = &v
	return s
}

func (s *BindTagsResponse) SetState(v BindTagsResponseStateEnum) *BindTagsResponse {
	s.State = &v
	return s
}

func (s *BindTagsResponse) SetBody(v string) *BindTagsResponse {
	s.Body = &v
	return s
}

func (s *BindTagsResponse) SetErrorParams(v []string) *BindTagsResponse {
	s.ErrorParams = v
	return s
}


type BindTagsResponseBuilder struct {
	s *BindTagsResponse
}

func NewBindTagsResponseBuilder() *BindTagsResponseBuilder {
	s := &BindTagsResponse{}
	b := &BindTagsResponseBuilder{s: s}
	return b
}

func (b *BindTagsResponseBuilder) AsyncKey(v string) *BindTagsResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *BindTagsResponseBuilder) RequestId(v string) *BindTagsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BindTagsResponseBuilder) ErrorMessage(v string) *BindTagsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BindTagsResponseBuilder) ErrorCode(v string) *BindTagsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BindTagsResponseBuilder) State(v BindTagsResponseStateEnum) *BindTagsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BindTagsResponseBuilder) Body(v string) *BindTagsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *BindTagsResponseBuilder) ErrorParams(v []string) *BindTagsResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *BindTagsResponseBuilder) Build() *BindTagsResponse {
	return b.s
}


