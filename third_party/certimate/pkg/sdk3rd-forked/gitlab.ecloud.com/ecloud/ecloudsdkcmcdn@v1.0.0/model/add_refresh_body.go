// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddRefreshBodyRefreshTypeEnum string

// List of RefreshType
const (
    AddRefreshBodyRefreshTypeEnumFile AddRefreshBodyRefreshTypeEnum = "FILE"
    AddRefreshBodyRefreshTypeEnumCatalog AddRefreshBodyRefreshTypeEnum = "CATALOG"
)

type AddRefreshBody struct {
    position.Body
    // 要刷新的URL或者目录
	Contents []string `json:"contents,omitempty"`
    // 刷新类型 文件：FILE，目录：CATALOG
	RefreshType *AddRefreshBodyRefreshTypeEnum `json:"refreshType,omitempty"`
}

func (s AddRefreshBody) String() string {
	return utils.Beautify(s)
}

func (s AddRefreshBody) GoString() string {
	return s.String()
}

func (s AddRefreshBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddRefreshBody) SetContents(v []string) *AddRefreshBody {
	s.Contents = v
	return s
}

func (s *AddRefreshBody) SetRefreshType(v AddRefreshBodyRefreshTypeEnum) *AddRefreshBody {
	s.RefreshType = &v
	return s
}


type AddRefreshBodyBuilder struct {
	s *AddRefreshBody
}

func NewAddRefreshBodyBuilder() *AddRefreshBodyBuilder {
	s := &AddRefreshBody{}
	b := &AddRefreshBodyBuilder{s: s}
	return b
}

func (b *AddRefreshBodyBuilder) Contents(v []string) *AddRefreshBodyBuilder {
	b.s.Contents = v
	return b
}

func (b *AddRefreshBodyBuilder) RefreshType(v AddRefreshBodyRefreshTypeEnum) *AddRefreshBodyBuilder {
	b.s.RefreshType = &v
	return b
}

func (b *AddRefreshBodyBuilder) Build() *AddRefreshBody {
	return b.s
}


