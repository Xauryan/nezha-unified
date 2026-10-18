// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeCdnDomainDetailResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *DescribeCdnDomainDetailResponseBody `json:"body,omitempty"`
}

func (s DescribeCdnDomainDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponse) SetRequestId(v string) *DescribeCdnDomainDetailResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetErrorMessage(v string) *DescribeCdnDomainDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetErrorCode(v string) *DescribeCdnDomainDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetState(v string) *DescribeCdnDomainDetailResponse {
	s.State = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetBody(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponse {
	s.Body = v
	return s
}


type DescribeCdnDomainDetailResponseBuilder struct {
	s *DescribeCdnDomainDetailResponse
}

func NewDescribeCdnDomainDetailResponseBuilder() *DescribeCdnDomainDetailResponseBuilder {
	s := &DescribeCdnDomainDetailResponse{}
	b := &DescribeCdnDomainDetailResponseBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) RequestId(v string) *DescribeCdnDomainDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) ErrorMessage(v string) *DescribeCdnDomainDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) ErrorCode(v string) *DescribeCdnDomainDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) State(v string) *DescribeCdnDomainDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) Body(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DescribeCdnDomainDetailResponseBuilder) Build() *DescribeCdnDomainDetailResponse {
	return b.s
}


