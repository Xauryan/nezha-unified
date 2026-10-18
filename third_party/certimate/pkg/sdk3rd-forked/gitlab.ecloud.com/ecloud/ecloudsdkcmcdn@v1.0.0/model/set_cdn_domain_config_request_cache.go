// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestCache struct {

    // 缓存设置列表，仅当enable为true时有值，支持指定路径或者文件名后缀方式
	Detail *[]SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestCache) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestCache) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestCache) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestCache) SetDetail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestCache {
	s.Detail = &v
	return s
}


type SetCdnDomainConfigRequestCacheBuilder struct {
	s *SetCdnDomainConfigRequestCache
}

func NewSetCdnDomainConfigRequestCacheBuilder() *SetCdnDomainConfigRequestCacheBuilder {
	s := &SetCdnDomainConfigRequestCache{}
	b := &SetCdnDomainConfigRequestCacheBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestCacheBuilder) Detail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestCacheBuilder {
	b.s.Detail = &v
	return b
}

func (b *SetCdnDomainConfigRequestCacheBuilder) Build() *SetCdnDomainConfigRequestCache {
	return b.s
}


