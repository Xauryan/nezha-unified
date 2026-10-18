// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListNewServerPortsForMemberResponseContentOpStatusEnum string

// List of OpStatus
const (
    ListNewServerPortsForMemberResponseContentOpStatusEnumCreating ListNewServerPortsForMemberResponseContentOpStatusEnum = "CREATING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumOk ListNewServerPortsForMemberResponseContentOpStatusEnum = "OK"
    ListNewServerPortsForMemberResponseContentOpStatusEnumRebooting ListNewServerPortsForMemberResponseContentOpStatusEnum = "REBOOTING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumFrozen ListNewServerPortsForMemberResponseContentOpStatusEnum = "FROZEN"
    ListNewServerPortsForMemberResponseContentOpStatusEnumBackupRebuilding ListNewServerPortsForMemberResponseContentOpStatusEnum = "BACKUP_REBUILDING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumRebuilding ListNewServerPortsForMemberResponseContentOpStatusEnum = "REBUILDING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumResizing ListNewServerPortsForMemberResponseContentOpStatusEnum = "RESIZING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumShelving ListNewServerPortsForMemberResponseContentOpStatusEnum = "SHELVING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumUnshelving ListNewServerPortsForMemberResponseContentOpStatusEnum = "UNSHELVING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumStarting ListNewServerPortsForMemberResponseContentOpStatusEnum = "STARTING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumStoping ListNewServerPortsForMemberResponseContentOpStatusEnum = "STOPING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumPasswordUpdating ListNewServerPortsForMemberResponseContentOpStatusEnum = "PASSWORD_UPDATING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumBackuping ListNewServerPortsForMemberResponseContentOpStatusEnum = "BACKUPING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumImageCreating ListNewServerPortsForMemberResponseContentOpStatusEnum = "IMAGE_CREATING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumUnrecognized ListNewServerPortsForMemberResponseContentOpStatusEnum = "UNRECOGNIZED"
    ListNewServerPortsForMemberResponseContentOpStatusEnumBootVolumeResizing ListNewServerPortsForMemberResponseContentOpStatusEnum = "BOOT_VOLUME_RESIZING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumDeleting ListNewServerPortsForMemberResponseContentOpStatusEnum = "DELETING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumSecondDeleting ListNewServerPortsForMemberResponseContentOpStatusEnum = "SECOND_DELETING"
    ListNewServerPortsForMemberResponseContentOpStatusEnumMigrating ListNewServerPortsForMemberResponseContentOpStatusEnum = "MIGRATING"
)

type ListNewServerPortsForMemberResponseContent struct {

    // 子网ID
	SubnetId *string `json:"subnetId,omitempty"`
    // 绑定的资源ID
	ResourceId *string `json:"resourceId,omitempty"`
    // 内网IP
	PrivateIp *string `json:"privateIp,omitempty"`
    // 绑定的资源名称
	ResourceName *string `json:"resourceName,omitempty"`
    // 云主机状态，使用时注意实时查询云主机的状态，https://eop.vip.ha.core:18080/portal/product/vm
	OpStatus *ListNewServerPortsForMemberResponseContentOpStatusEnum `json:"opStatus,omitempty"`
}

func (s ListNewServerPortsForMemberResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListNewServerPortsForMemberResponseContent) GoString() string {
	return s.String()
}

func (s ListNewServerPortsForMemberResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListNewServerPortsForMemberResponseContent) SetSubnetId(v string) *ListNewServerPortsForMemberResponseContent {
	s.SubnetId = &v
	return s
}

func (s *ListNewServerPortsForMemberResponseContent) SetResourceId(v string) *ListNewServerPortsForMemberResponseContent {
	s.ResourceId = &v
	return s
}

func (s *ListNewServerPortsForMemberResponseContent) SetPrivateIp(v string) *ListNewServerPortsForMemberResponseContent {
	s.PrivateIp = &v
	return s
}

func (s *ListNewServerPortsForMemberResponseContent) SetResourceName(v string) *ListNewServerPortsForMemberResponseContent {
	s.ResourceName = &v
	return s
}

func (s *ListNewServerPortsForMemberResponseContent) SetOpStatus(v ListNewServerPortsForMemberResponseContentOpStatusEnum) *ListNewServerPortsForMemberResponseContent {
	s.OpStatus = &v
	return s
}


type ListNewServerPortsForMemberResponseContentBuilder struct {
	s *ListNewServerPortsForMemberResponseContent
}

func NewListNewServerPortsForMemberResponseContentBuilder() *ListNewServerPortsForMemberResponseContentBuilder {
	s := &ListNewServerPortsForMemberResponseContent{}
	b := &ListNewServerPortsForMemberResponseContentBuilder{s: s}
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) SubnetId(v string) *ListNewServerPortsForMemberResponseContentBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) ResourceId(v string) *ListNewServerPortsForMemberResponseContentBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) PrivateIp(v string) *ListNewServerPortsForMemberResponseContentBuilder {
	b.s.PrivateIp = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) ResourceName(v string) *ListNewServerPortsForMemberResponseContentBuilder {
	b.s.ResourceName = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) OpStatus(v ListNewServerPortsForMemberResponseContentOpStatusEnum) *ListNewServerPortsForMemberResponseContentBuilder {
	b.s.OpStatus = &v
	return b
}

func (b *ListNewServerPortsForMemberResponseContentBuilder) Build() *ListNewServerPortsForMemberResponseContent {
	return b.s
}


