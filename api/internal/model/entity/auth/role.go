// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	RoleId     uint        `json:"roleId"     ` // 权限角色ID
	SceneId    uint        `json:"sceneId"    ` // 权限场景ID
	TableId    uint        `json:"tableId"    ` // 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
	RoleName   string      `json:"roleName"   ` // 名称
	IsStop     uint        `json:"isStop"     ` // 是否停用：0否 1是
	UpdateAt *gtime.Time `json:"updateAt" ` // 更新时间
	CreateAt *gtime.Time `json:"createAt" ` // 创建时间
}
