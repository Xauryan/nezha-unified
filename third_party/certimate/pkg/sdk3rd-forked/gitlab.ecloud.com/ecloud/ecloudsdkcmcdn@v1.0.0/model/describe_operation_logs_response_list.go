// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeOperationLogsResponseList struct {

    // 操作类型 刷新,预热
	OperateType *string `json:"operateType,omitempty"`
    // 刷新任务url的索引
	Index *int32 `json:"index,omitempty"`
    // 最新更新时间
	UpdateTime *string `json:"updateTime,omitempty"`
    // 用户id
	UserId *string `json:"userId,omitempty"`
    // 内容更新URL或者path
	Content *string `json:"content,omitempty"`
    // 是否删除，0：否，1：是
	Deleted *bool `json:"deleted,omitempty"`
    // 创建时间
	CreateTime *string `json:"createTime,omitempty"`
    // 刷新类型 文件,目录
	RefreshType *string `json:"refreshType,omitempty"`
    // 任务开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 任务结束时间
	EndTime *string `json:"endTime,omitempty"`

	Id *int32 `json:"id,omitempty"`
    // 任务Id
	TaskId *string `json:"taskId,omitempty"`
    // 状态
	Status *string `json:"status,omitempty"`
}

func (s DescribeOperationLogsResponseList) String() string {
	return utils.Beautify(s)
}

func (s DescribeOperationLogsResponseList) GoString() string {
	return s.String()
}

func (s DescribeOperationLogsResponseList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeOperationLogsResponseList) SetOperateType(v string) *DescribeOperationLogsResponseList {
	s.OperateType = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetIndex(v int32) *DescribeOperationLogsResponseList {
	s.Index = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetUpdateTime(v string) *DescribeOperationLogsResponseList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetUserId(v string) *DescribeOperationLogsResponseList {
	s.UserId = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetContent(v string) *DescribeOperationLogsResponseList {
	s.Content = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetDeleted(v bool) *DescribeOperationLogsResponseList {
	s.Deleted = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetCreateTime(v string) *DescribeOperationLogsResponseList {
	s.CreateTime = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetRefreshType(v string) *DescribeOperationLogsResponseList {
	s.RefreshType = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetStartTime(v string) *DescribeOperationLogsResponseList {
	s.StartTime = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetEndTime(v string) *DescribeOperationLogsResponseList {
	s.EndTime = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetId(v int32) *DescribeOperationLogsResponseList {
	s.Id = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetTaskId(v string) *DescribeOperationLogsResponseList {
	s.TaskId = &v
	return s
}

func (s *DescribeOperationLogsResponseList) SetStatus(v string) *DescribeOperationLogsResponseList {
	s.Status = &v
	return s
}


type DescribeOperationLogsResponseListBuilder struct {
	s *DescribeOperationLogsResponseList
}

func NewDescribeOperationLogsResponseListBuilder() *DescribeOperationLogsResponseListBuilder {
	s := &DescribeOperationLogsResponseList{}
	b := &DescribeOperationLogsResponseListBuilder{s: s}
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) OperateType(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.OperateType = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Index(v int32) *DescribeOperationLogsResponseListBuilder {
	b.s.Index = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) UpdateTime(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.UpdateTime = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) UserId(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.UserId = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Content(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.Content = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Deleted(v bool) *DescribeOperationLogsResponseListBuilder {
	b.s.Deleted = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) CreateTime(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.CreateTime = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) RefreshType(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.RefreshType = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) StartTime(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.StartTime = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) EndTime(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.EndTime = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Id(v int32) *DescribeOperationLogsResponseListBuilder {
	b.s.Id = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) TaskId(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.TaskId = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Status(v string) *DescribeOperationLogsResponseListBuilder {
	b.s.Status = &v
	return b
}

func (b *DescribeOperationLogsResponseListBuilder) Build() *DescribeOperationLogsResponseList {
	return b.s
}


