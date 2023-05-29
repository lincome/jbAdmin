// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalRoleRelToActionDao is internal type for wrapping internal DAO implements.
type internalRoleRelToActionDao = *internal.RoleRelToActionDao

// roleRelToActionDao is the data access object for table auth_role_rel_to_action.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelToActionDao struct {
	internalRoleRelToActionDao
}

var (
	// RoleRelToAction is globally public accessible object for table auth_role_rel_to_action operations.
	RoleRelToAction = roleRelToActionDao{
		internal.NewRoleRelToActionDao(),
	}
)

// 解析insert
func (dao *roleRelToActionDao) ParseInsert(insert ...map[string]interface{}) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		insertData := []map[string]interface{}{}
		for index, item := range insert {
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][dao.PrimaryKey()] = v
				default:
					/* //数据库不存在的字段过滤掉
					if !garray.NewStrArrayFrom(dao.Column()).Contains(k) {
						continue
					} */
					insertData[index][k] = v
				}
			}
		}
		if len(insertData) == 1 {
			m = m.Data(insertData[0])
		} else {
			m = m.Data(insertData)
		}
		return m
	}
}

// 解析update
func (dao *roleRelToActionDao) ParseUpdate(update map[string]interface{}) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[dao.Table()+"."+dao.PrimaryKey()] = v
			default:
				/* //数据库不存在的字段过滤掉
				if !garray.NewStrArrayFrom(dao.Column()).Contains(k) {
						continue
				} */
				updateData[dao.Table()+"."+k] = v
			}
		}
		m = m.Data(updateData)
		return m
	}
}

// 解析field
func (dao *roleRelToActionDao) ParseField(field []string, afterField *[]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
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

// 解析filter
func (dao *roleRelToActionDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id":
				m = m.Where(dao.Table()+"."+dao.PrimaryKey(), v)
			case "excId":
				//m = m.Where(dao.Table()+"."+dao.PrimaryKey()+" <> ?", v)
				m = m.WhereNot(dao.Table()+"."+dao.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(dao.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(dao.Table()+".createTime", v)
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

// 解析group
func (dao *roleRelToActionDao) ParseGroup(group []string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(dao.Table() + "." + dao.PrimaryKey())
			default:
				if garray.NewStrArrayFrom(dao.Column()).Contains(v) {
					m = m.Group(dao.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (dao *roleRelToActionDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(dao.Table()+"."+dao.PrimaryKey(), v[1])
			default:
				if garray.NewStrArrayFrom(dao.Column()).Contains(v[0]) {
					m = m.Order(dao.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (dao *roleRelToActionDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(*joinCodeArr).Contains(joinCode) {
			*joinCodeArr = append(*joinCodeArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin("xxxx", "xxxx."+dao.PrimaryKey()+" = "+dao.Table()+"."+dao.PrimaryKey()) */
			}
		}
		return m
	}
}

// Fill with you ideas below.
