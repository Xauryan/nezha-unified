// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"os"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddCdnDomainBodyCdnTypeEnum string

// List of CdnType
const (
	AddCdnDomainBodyCdnTypeEnumWeb        AddCdnDomainBodyCdnTypeEnum = "web"
	AddCdnDomainBodyCdnTypeEnumDownload   AddCdnDomainBodyCdnTypeEnum = "download"
	AddCdnDomainBodyCdnTypeEnumVideo      AddCdnDomainBodyCdnTypeEnum = "video"
	AddCdnDomainBodyCdnTypeEnumLivestream AddCdnDomainBodyCdnTypeEnum = "liveStream"
)

type AddCdnDomainBodyCdnProtocolEnum string

// List of CdnProtocol
const (
	AddCdnDomainBodyCdnProtocolEnumHttp      AddCdnDomainBodyCdnProtocolEnum = "HTTP"
	AddCdnDomainBodyCdnProtocolEnumHttps     AddCdnDomainBodyCdnProtocolEnum = "HTTPS"
	AddCdnDomainBodyCdnProtocolEnumHttpHttps AddCdnDomainBodyCdnProtocolEnum = "HTTP_HTTPS"
)

type AddCdnDomainBodyProductTypeEnum string

// List of ProductType
const (
	AddCdnDomainBodyProductTypeEnumEcdn AddCdnDomainBodyProductTypeEnum = "ECDN"
)

type AddCdnDomainBody struct {
	position.Body
	// 上传的文件
	File *os.File `json:"file,omitempty"`
	// 加速域名的业务类型：页面：web，下载：download，点播：video
	CdnType *AddCdnDomainBodyCdnTypeEnum `json:"cdnType,omitempty"`
	// 证书uniqueId
	CrtUniqueId *string `json:"crtUniqueId,omitempty"`
	// 域名
	DomainName *string `json:"domainName,omitempty"`
	// 协议类型：http，https
	CdnProtocol *AddCdnDomainBodyCdnProtocolEnum `json:"cdnProtocol,omitempty"`
	// 源站信息
	CdnSrcDetails *[]AddCdnDomainRequestCdnSrcDetails `json:"cdnSrcDetails,omitempty"`
	// 域名备案号
	IcpLicensing *string `json:"icpLicensing,omitempty"`
	// 产品类型
	ProductType *AddCdnDomainBodyProductTypeEnum `json:"productType,omitempty"`
}

func (s AddCdnDomainBody) String() string {
	return utils.Beautify(s)
}

func (s AddCdnDomainBody) GoString() string {
	return s.String()
}

func (s AddCdnDomainBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddCdnDomainBody) SetFile(v os.File) *AddCdnDomainBody {
	s.File = &v
	return s
}

func (s *AddCdnDomainBody) SetCdnType(v AddCdnDomainBodyCdnTypeEnum) *AddCdnDomainBody {
	s.CdnType = &v
	return s
}

func (s *AddCdnDomainBody) SetCrtUniqueId(v string) *AddCdnDomainBody {
	s.CrtUniqueId = &v
	return s
}

func (s *AddCdnDomainBody) SetDomainName(v string) *AddCdnDomainBody {
	s.DomainName = &v
	return s
}

func (s *AddCdnDomainBody) SetCdnProtocol(v AddCdnDomainBodyCdnProtocolEnum) *AddCdnDomainBody {
	s.CdnProtocol = &v
	return s
}

func (s *AddCdnDomainBody) SetCdnSrcDetails(v []AddCdnDomainRequestCdnSrcDetails) *AddCdnDomainBody {
	s.CdnSrcDetails = &v
	return s
}

func (s *AddCdnDomainBody) SetIcpLicensing(v string) *AddCdnDomainBody {
	s.IcpLicensing = &v
	return s
}

func (s *AddCdnDomainBody) SetProductType(v AddCdnDomainBodyProductTypeEnum) *AddCdnDomainBody {
	s.ProductType = &v
	return s
}

type AddCdnDomainBodyBuilder struct {
	s *AddCdnDomainBody
}

func NewAddCdnDomainBodyBuilder() *AddCdnDomainBodyBuilder {
	s := &AddCdnDomainBody{}
	b := &AddCdnDomainBodyBuilder{s: s}
	return b
}

func (b *AddCdnDomainBodyBuilder) File(v os.File) *AddCdnDomainBodyBuilder {
	b.s.File = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) CdnType(v AddCdnDomainBodyCdnTypeEnum) *AddCdnDomainBodyBuilder {
	b.s.CdnType = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) CrtUniqueId(v string) *AddCdnDomainBodyBuilder {
	b.s.CrtUniqueId = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) DomainName(v string) *AddCdnDomainBodyBuilder {
	b.s.DomainName = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) CdnProtocol(v AddCdnDomainBodyCdnProtocolEnum) *AddCdnDomainBodyBuilder {
	b.s.CdnProtocol = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) CdnSrcDetails(v []AddCdnDomainRequestCdnSrcDetails) *AddCdnDomainBodyBuilder {
	b.s.CdnSrcDetails = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) IcpLicensing(v string) *AddCdnDomainBodyBuilder {
	b.s.IcpLicensing = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) ProductType(v AddCdnDomainBodyProductTypeEnum) *AddCdnDomainBodyBuilder {
	b.s.ProductType = &v
	return b
}

func (b *AddCdnDomainBodyBuilder) Build() *AddCdnDomainBody {
	return b.s
}
