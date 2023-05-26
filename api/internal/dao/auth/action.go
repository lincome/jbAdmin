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

// internalActionDao is internal type for wrapping internal DAO implements.
type internalActionDao = *internal.ActionDao

// actionDao is the data access object for table auth_action.
// You can define custom methods on it to extend its functionality as you wish.
type actionDao struct {
	internalActionDao
}

var (
	// Action is globally public accessible object for table auth_action operations.
	Action = actionDao{
		internal.NewActionDao(),
	}
)

func (dao actionDao) ParseField(field []string, afterField *[]string, joinCode *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			case "id":
				m = m.Fields(dao.Table() + "." + dao.PrimaryKey())
			default:
				if garray.NewStrArrayFrom(dao.Column()).Contains(v) {
					m = m.Fields(dao.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		return m
	}
}

func (dao actionDao) ParseFilter(filter g.MapStrAny, joinCode *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id":
				m = m.Where(dao.Table()+"."+dao.PrimaryKey(), v)
			default:
				kArr := strings.Split(k, " ")
				if garray.NewStrArrayFrom(dao.Column()).Contains(kArr[0]) {
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
