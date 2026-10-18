// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeletePoolAsyncResponseStateEnum string

// List of State
const (
    DeletePoolAsyncResponseStateEnumOk DeletePoolAsyncResponseStateEnum = "OK"
    DeletePoolAsyncResponseStateEnumError DeletePoolAsyncResponseStateEnum = "ERROR"
    DeletePoolAsyncResponseStateEnumException DeletePoolAsyncResponseStateEnum = "EXCEPTION"
    DeletePoolAsyncResponseStateEnumAlarm DeletePoolAsyncResponseStateEnum = "ALARM"
    DeletePoolAsyncResponseStateEnumForbidden DeletePoolAsyncResponseStateEnum = "FORBIDDEN"
)

type DeletePoolAsyncResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *DeletePoolAsyncResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeletePoolAsyncResponse) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolAsyncResponse) GoString() string {
	return s.String()
}

func (s DeletePoolAsyncResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolAsyncResponse) SetAsyncKey(v string) *DeletePoolAsyncResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetRequestId(v string) *DeletePoolAsyncResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetErrorMessage(v string) *DeletePoolAsyncResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetErrorCode(v string) *DeletePoolAsyncResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetState(v DeletePoolAsyncResponseStateEnum) *DeletePoolAsyncResponse {
	s.State = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetBody(v string) *DeletePoolAsyncResponse {
	s.Body = &v
	return s
}

func (s *DeletePoolAsyncResponse) SetErrorParams(v []string) *DeletePoolAsyncResponse {
	s.ErrorParams = v
	return s
}


type DeletePoolAsyncResponseBuilder struct {
	s *DeletePoolAsyncResponse
}

func NewDeletePoolAsyncResponseBuilder() *DeletePoolAsyncResponseBuilder {
	s := &DeletePoolAsyncResponse{}
	b := &DeletePoolAsyncResponseBuilder{s: s}
	return b
}

func (b *DeletePoolAsyncResponseBuilder) AsyncKey(v string) *DeletePoolAsyncResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) RequestId(v string) *DeletePoolAsyncResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) ErrorMessage(v string) *DeletePoolAsyncResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) ErrorCode(v string) *DeletePoolAsyncResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) State(v DeletePoolAsyncResponseStateEnum) *DeletePoolAsyncResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) Body(v string) *DeletePoolAsyncResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) ErrorParams(v []string) *DeletePoolAsyncResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeletePoolAsyncResponseBuilder) Build() *DeletePoolAsyncResponse {
	return b.s
}


