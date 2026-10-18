// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserDomainsResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *DescribeUserDomainsResponseBody `json:"body,omitempty"`
}

func (s DescribeUserDomainsResponse) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserDomainsResponse) GoString() string {
	return s.String()
}

func (s DescribeUserDomainsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserDomainsResponse) SetRequestId(v string) *DescribeUserDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetErrorMessage(v string) *DescribeUserDomainsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetErrorCode(v string) *DescribeUserDomainsResponse {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetState(v string) *DescribeUserDomainsResponse {
	s.State = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetBody(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponse {
	s.Body = v
	return s
}


type DescribeUserDomainsResponseBuilder struct {
	s *DescribeUserDomainsResponse
}

func NewDescribeUserDomainsResponseBuilder() *DescribeUserDomainsResponseBuilder {
	s := &DescribeUserDomainsResponse{}
	b := &DescribeUserDomainsResponseBuilder{s: s}
	return b
}

func (b *DescribeUserDomainsResponseBuilder) RequestId(v string) *DescribeUserDomainsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DescribeUserDomainsResponseBuilder) ErrorMessage(v string) *DescribeUserDomainsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DescribeUserDomainsResponseBuilder) ErrorCode(v string) *DescribeUserDomainsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DescribeUserDomainsResponseBuilder) State(v string) *DescribeUserDomainsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DescribeUserDomainsResponseBuilder) Body(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DescribeUserDomainsResponseBuilder) Build() *DescribeUserDomainsResponse {
	return b.s
}


