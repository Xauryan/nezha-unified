// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindTagsBody struct {
    position.Body
    // 资源ID列表
	ResourceIds []string `json:"resourceIds,omitempty"`
    // 资源所属类型，弹性负载均衡实例类型为CIDC-RT-SLB
	ResourceType *string `json:"resourceType,omitempty"`
    // 标签
	Tags *[]BindTagsRequestTags `json:"tags,omitempty"`
}

func (s BindTagsBody) String() string {
	return utils.Beautify(s)
}

func (s BindTagsBody) GoString() string {
	return s.String()
}

func (s BindTagsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindTagsBody) SetResourceIds(v []string) *BindTagsBody {
	s.ResourceIds = v
	return s
}

func (s *BindTagsBody) SetResourceType(v string) *BindTagsBody {
	s.ResourceType = &v
	return s
}

func (s *BindTagsBody) SetTags(v []BindTagsRequestTags) *BindTagsBody {
	s.Tags = &v
	return s
}


type BindTagsBodyBuilder struct {
	s *BindTagsBody
}

func NewBindTagsBodyBuilder() *BindTagsBodyBuilder {
	s := &BindTagsBody{}
	b := &BindTagsBodyBuilder{s: s}
	return b
}

func (b *BindTagsBodyBuilder) ResourceIds(v []string) *BindTagsBodyBuilder {
	b.s.ResourceIds = v
	return b
}

func (b *BindTagsBodyBuilder) ResourceType(v string) *BindTagsBodyBuilder {
	b.s.ResourceType = &v
	return b
}

func (b *BindTagsBodyBuilder) Tags(v []BindTagsRequestTags) *BindTagsBodyBuilder {
	b.s.Tags = &v
	return b
}

func (b *BindTagsBodyBuilder) Build() *BindTagsBody {
	return b.s
}


