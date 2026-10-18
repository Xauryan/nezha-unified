// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CertificationBindTagsBody struct {
    position.Body
    // 资源ID, 当资源所属类型不为空时 不可为空
	ResourceIds []string `json:"resourceIds,omitempty"`
    // 资源所属类型，证书类型为CIDC-RT-BARBICAN
	ResourceType *string `json:"resourceType,omitempty"`
    // 标签
	Tags *[]CertificationBindTagsRequestTags `json:"tags,omitempty"`
}

func (s CertificationBindTagsBody) String() string {
	return utils.Beautify(s)
}

func (s CertificationBindTagsBody) GoString() string {
	return s.String()
}

func (s CertificationBindTagsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationBindTagsBody) SetResourceIds(v []string) *CertificationBindTagsBody {
	s.ResourceIds = v
	return s
}

func (s *CertificationBindTagsBody) SetResourceType(v string) *CertificationBindTagsBody {
	s.ResourceType = &v
	return s
}

func (s *CertificationBindTagsBody) SetTags(v []CertificationBindTagsRequestTags) *CertificationBindTagsBody {
	s.Tags = &v
	return s
}


type CertificationBindTagsBodyBuilder struct {
	s *CertificationBindTagsBody
}

func NewCertificationBindTagsBodyBuilder() *CertificationBindTagsBodyBuilder {
	s := &CertificationBindTagsBody{}
	b := &CertificationBindTagsBodyBuilder{s: s}
	return b
}

func (b *CertificationBindTagsBodyBuilder) ResourceIds(v []string) *CertificationBindTagsBodyBuilder {
	b.s.ResourceIds = v
	return b
}

func (b *CertificationBindTagsBodyBuilder) ResourceType(v string) *CertificationBindTagsBodyBuilder {
	b.s.ResourceType = &v
	return b
}

func (b *CertificationBindTagsBodyBuilder) Tags(v []CertificationBindTagsRequestTags) *CertificationBindTagsBodyBuilder {
	b.s.Tags = &v
	return b
}

func (b *CertificationBindTagsBodyBuilder) Build() *CertificationBindTagsBody {
	return b.s
}


