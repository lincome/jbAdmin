// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// internalActionRelToSceneDao is internal type for wrapping internal DAO implements.
type internalActionRelToSceneDao = *internal.ActionRelToSceneDao

// actionRelToSceneDao is the data access object for table auth_action_rel_to_scene.
// You can define custom methods on it to extend its functionality as you wish.
type actionRelToSceneDao struct {
	internalActionRelToSceneDao
}

var (
	// ActionRelToScene is globally public accessible object for table auth_action_rel_to_scene operations.
	ActionRelToScene = actionRelToSceneDao{
		internal.NewActionRelToSceneDao(),
	}
)

func (dao actionRelToSceneDao) Filter(filter g.MapStrAny, joinCode *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id":
				m = m.Where(dao.Table()+"."+dao.PrimaryKey(), v)
			default:
				kArr := strings.Split(k, " ")
				if garray.NewStrArrayFromCopy(dao.Column()).Contains(kArr[0]) {
					m = m.Where(dao.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// Fill with you ideas below.
