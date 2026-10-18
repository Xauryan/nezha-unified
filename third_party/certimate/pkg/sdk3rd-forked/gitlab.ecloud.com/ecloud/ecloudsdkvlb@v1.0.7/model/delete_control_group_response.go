// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteControlGroupResponseStateEnum string

// List of State
const (
    DeleteControlGroupResponseStateEnumOk DeleteControlGroupResponseStateEnum = "OK"
    DeleteControlGroupResponseStateEnumError DeleteControlGroupResponseStateEnum = "ERROR"
    DeleteControlGroupResponseStateEnumException DeleteControlGroupResponseStateEnum = "EXCEPTION"
    DeleteControlGroupResponseStateEnumAlarm DeleteControlGroupResponseStateEnum = "ALARM"
    DeleteControlGroupResponseStateEnumForbidden DeleteControlGroupResponseStateEnum = "FORBIDDEN"
)

type DeleteControlGroupResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功）, ERROR（错误）,EXCEPTION（异常）,ALARM（警告）, FORBIDDEN（禁止访问）
	State *DeleteControlGroupResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeleteControlGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupResponse) GoString() string {
	return s.String()
}

func (s DeleteControlGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupResponse) SetAsyncKey(v string) *DeleteControlGroupResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeleteControlGroupResponse) SetRequestId(v string) *DeleteControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteControlGroupResponse) SetErrorMessage(v string) *DeleteControlGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteControlGroupResponse) SetErrorCode(v string) *DeleteControlGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteControlGroupResponse) SetState(v DeleteControlGroupResponseStateEnum) *DeleteControlGroupResponse {
	s.State = &v
	return s
}

func (s *DeleteControlGroupResponse) SetBody(v bool) *DeleteControlGroupResponse {
	s.Body = &v
	return s
}

func (s *DeleteControlGroupResponse) SetErrorParams(v []string) *DeleteControlGroupResponse {
	s.ErrorParams = v
	return s
}


type DeleteControlGroupResponseBuilder struct {
	s *DeleteControlGroupResponse
}

func NewDeleteControlGroupResponseBuilder() *DeleteControlGroupResponseBuilder {
	s := &DeleteControlGroupResponse{}
	b := &DeleteControlGroupResponseBuilder{s: s}
	return b
}

func (b *DeleteControlGroupResponseBuilder) AsyncKey(v string) *DeleteControlGroupResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) RequestId(v string) *DeleteControlGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) ErrorMessage(v string) *DeleteControlGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) ErrorCode(v string) *DeleteControlGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) State(v DeleteControlGroupResponseStateEnum) *DeleteControlGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) Body(v bool) *DeleteControlGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteControlGroupResponseBuilder) ErrorParams(v []string) *DeleteControlGroupResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeleteControlGroupResponseBuilder) Build() *DeleteControlGroupResponse {
	return b.s
}


