// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"os"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyCdnDomainBodyCdnProtocolEnum string

// List of CdnProtocol
const (
	ModifyCdnDomainBodyCdnProtocolEnumHttp      ModifyCdnDomainBodyCdnProtocolEnum = "HTTP"
	ModifyCdnDomainBodyCdnProtocolEnumHttps     ModifyCdnDomainBodyCdnProtocolEnum = "HTTPS"
	ModifyCdnDomainBodyCdnProtocolEnumHttpHttps ModifyCdnDomainBodyCdnProtocolEnum = "HTTP_HTTPS"
)

type ModifyCdnDomainBody struct {
	position.Body
	// 上传的文件列表
	File *os.File `json:"file,omitempty"`
	// 证书uniqueId
	CrtUniqueId *string `json:"crtUniqueId,omitempty"`
	// 回源地址列表
	Sources *[]ModifyCdnDomainRequestSources `json:"sources,omitempty"`
	// 客户访问服务节点的协议
	CdnProtocol *ModifyCdnDomainBodyCdnProtocolEnum `json:"cdnProtocol,omitempty"`
	// 域名Id
	DomainId *int32 `json:"domainId,omitempty"`
	// 已经存在的文件名
	ExistFile []string `json:"existFile,omitempty"`
}

func (s ModifyCdnDomainBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyCdnDomainBody) GoString() string {
	return s.String()
}

func (s ModifyCdnDomainBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyCdnDomainBody) SetFile(v os.File) *ModifyCdnDomainBody {
	s.File = &v
	return s
}

func (s *ModifyCdnDomainBody) SetCrtUniqueId(v string) *ModifyCdnDomainBody {
	s.CrtUniqueId = &v
	return s
}

func (s *ModifyCdnDomainBody) SetSources(v []ModifyCdnDomainRequestSources) *ModifyCdnDomainBody {
	s.Sources = &v
	return s
}

func (s *ModifyCdnDomainBody) SetCdnProtocol(v ModifyCdnDomainBodyCdnProtocolEnum) *ModifyCdnDomainBody {
	s.CdnProtocol = &v
	return s
}

func (s *ModifyCdnDomainBody) SetDomainId(v int32) *ModifyCdnDomainBody {
	s.DomainId = &v
	return s
}

func (s *ModifyCdnDomainBody) SetExistFile(v []string) *ModifyCdnDomainBody {
	s.ExistFile = v
	return s
}

type ModifyCdnDomainBodyBuilder struct {
	s *ModifyCdnDomainBody
}

func NewModifyCdnDomainBodyBuilder() *ModifyCdnDomainBodyBuilder {
	s := &ModifyCdnDomainBody{}
	b := &ModifyCdnDomainBodyBuilder{s: s}
	return b
}

func (b *ModifyCdnDomainBodyBuilder) File(v os.File) *ModifyCdnDomainBodyBuilder {
	b.s.File = &v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) CrtUniqueId(v string) *ModifyCdnDomainBodyBuilder {
	b.s.CrtUniqueId = &v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) Sources(v []ModifyCdnDomainRequestSources) *ModifyCdnDomainBodyBuilder {
	b.s.Sources = &v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) CdnProtocol(v ModifyCdnDomainBodyCdnProtocolEnum) *ModifyCdnDomainBodyBuilder {
	b.s.CdnProtocol = &v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) DomainId(v int32) *ModifyCdnDomainBodyBuilder {
	b.s.DomainId = &v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) ExistFile(v []string) *ModifyCdnDomainBodyBuilder {
	b.s.ExistFile = v
	return b
}

func (b *ModifyCdnDomainBodyBuilder) Build() *ModifyCdnDomainBody {
	return b.s
}
