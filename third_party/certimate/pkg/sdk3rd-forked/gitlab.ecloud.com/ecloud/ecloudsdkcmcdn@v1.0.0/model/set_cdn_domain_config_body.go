// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigBody struct {
    position.Body
    // error page信息
	ErrorPageReq *SetCdnDomainConfigRequestErrorPageReq `json:"errorPageReq,omitempty"`
    // ip黑名单信息
	IpBlackList *SetCdnDomainConfigRequestIpBlackList `json:"ipBlackList,omitempty"`
    // 缓存类型与过期时间信息
	Cache *SetCdnDomainConfigRequestCache `json:"cache,omitempty"`
    // 防盗链信息
	ReferrerAntiStealingLink *SetCdnDomainConfigRequestReferrerAntiStealingLink `json:"referrerAntiStealingLink,omitempty"`
    // 回源host
	BackSourceHost *SetCdnDomainConfigRequestBackSourceHost `json:"backSourceHost,omitempty"`
    // http响应头配置信息
	RespHeaderReq *SetCdnDomainConfigRequestRespHeaderReq `json:"respHeaderReq,omitempty"`
    // 支持HTTP/2的协议配置
	Http2Status *SetCdnDomainConfigRequestHttp2Status `json:"http2Status,omitempty"`
    // 智能压缩功能
	GzipStatus *SetCdnDomainConfigRequestGzipStatus `json:"gzipStatus,omitempty"`
    // 域名Id
	DomainId *int32 `json:"domainId,omitempty"`
    // http请求头配置信息
	ReqHeaderReq *SetCdnDomainConfigRequestReqHeaderReq `json:"reqHeaderReq,omitempty"`
    // 过滤参数配置信息
	UrlParameterReq *SetCdnDomainConfigRequestUrlParameterReq `json:"urlParameterReq,omitempty"`
}

func (s SetCdnDomainConfigBody) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigBody) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigBody) SetErrorPageReq(v *SetCdnDomainConfigRequestErrorPageReq) *SetCdnDomainConfigBody {
	s.ErrorPageReq = v
	return s
}

func (s *SetCdnDomainConfigBody) SetIpBlackList(v *SetCdnDomainConfigRequestIpBlackList) *SetCdnDomainConfigBody {
	s.IpBlackList = v
	return s
}

func (s *SetCdnDomainConfigBody) SetCache(v *SetCdnDomainConfigRequestCache) *SetCdnDomainConfigBody {
	s.Cache = v
	return s
}

func (s *SetCdnDomainConfigBody) SetReferrerAntiStealingLink(v *SetCdnDomainConfigRequestReferrerAntiStealingLink) *SetCdnDomainConfigBody {
	s.ReferrerAntiStealingLink = v
	return s
}

func (s *SetCdnDomainConfigBody) SetBackSourceHost(v *SetCdnDomainConfigRequestBackSourceHost) *SetCdnDomainConfigBody {
	s.BackSourceHost = v
	return s
}

func (s *SetCdnDomainConfigBody) SetRespHeaderReq(v *SetCdnDomainConfigRequestRespHeaderReq) *SetCdnDomainConfigBody {
	s.RespHeaderReq = v
	return s
}

func (s *SetCdnDomainConfigBody) SetHttp2Status(v *SetCdnDomainConfigRequestHttp2Status) *SetCdnDomainConfigBody {
	s.Http2Status = v
	return s
}

func (s *SetCdnDomainConfigBody) SetGzipStatus(v *SetCdnDomainConfigRequestGzipStatus) *SetCdnDomainConfigBody {
	s.GzipStatus = v
	return s
}

func (s *SetCdnDomainConfigBody) SetDomainId(v int32) *SetCdnDomainConfigBody {
	s.DomainId = &v
	return s
}

func (s *SetCdnDomainConfigBody) SetReqHeaderReq(v *SetCdnDomainConfigRequestReqHeaderReq) *SetCdnDomainConfigBody {
	s.ReqHeaderReq = v
	return s
}

func (s *SetCdnDomainConfigBody) SetUrlParameterReq(v *SetCdnDomainConfigRequestUrlParameterReq) *SetCdnDomainConfigBody {
	s.UrlParameterReq = v
	return s
}


type SetCdnDomainConfigBodyBuilder struct {
	s *SetCdnDomainConfigBody
}

func NewSetCdnDomainConfigBodyBuilder() *SetCdnDomainConfigBodyBuilder {
	s := &SetCdnDomainConfigBody{}
	b := &SetCdnDomainConfigBodyBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) ErrorPageReq(v *SetCdnDomainConfigRequestErrorPageReq) *SetCdnDomainConfigBodyBuilder {
	b.s.ErrorPageReq = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) IpBlackList(v *SetCdnDomainConfigRequestIpBlackList) *SetCdnDomainConfigBodyBuilder {
	b.s.IpBlackList = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) Cache(v *SetCdnDomainConfigRequestCache) *SetCdnDomainConfigBodyBuilder {
	b.s.Cache = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) ReferrerAntiStealingLink(v *SetCdnDomainConfigRequestReferrerAntiStealingLink) *SetCdnDomainConfigBodyBuilder {
	b.s.ReferrerAntiStealingLink = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) BackSourceHost(v *SetCdnDomainConfigRequestBackSourceHost) *SetCdnDomainConfigBodyBuilder {
	b.s.BackSourceHost = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) RespHeaderReq(v *SetCdnDomainConfigRequestRespHeaderReq) *SetCdnDomainConfigBodyBuilder {
	b.s.RespHeaderReq = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) Http2Status(v *SetCdnDomainConfigRequestHttp2Status) *SetCdnDomainConfigBodyBuilder {
	b.s.Http2Status = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) GzipStatus(v *SetCdnDomainConfigRequestGzipStatus) *SetCdnDomainConfigBodyBuilder {
	b.s.GzipStatus = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) DomainId(v int32) *SetCdnDomainConfigBodyBuilder {
	b.s.DomainId = &v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) ReqHeaderReq(v *SetCdnDomainConfigRequestReqHeaderReq) *SetCdnDomainConfigBodyBuilder {
	b.s.ReqHeaderReq = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) UrlParameterReq(v *SetCdnDomainConfigRequestUrlParameterReq) *SetCdnDomainConfigBodyBuilder {
	b.s.UrlParameterReq = v
	return b
}

func (b *SetCdnDomainConfigBodyBuilder) Build() *SetCdnDomainConfigBody {
	return b.s
}


