// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActionRelToScene is the golang structure for table action_rel_to_scene.
type ActionRelToScene struct {
	ActionId  uint        `json:"actionId"  orm:"actionId"  ` // 操作ID
	SceneId   uint        `json:"sceneId"   orm:"sceneId"   ` // 场景ID
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
