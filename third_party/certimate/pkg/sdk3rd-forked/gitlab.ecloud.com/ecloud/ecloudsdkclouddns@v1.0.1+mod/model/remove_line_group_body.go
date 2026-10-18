// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcloudcore/position"
)

type RemoveLineGroupBody struct {
	position.Body
	// 待删除的线路分组ID列表
	GroupIds []string `json:"groupIds"`
}
