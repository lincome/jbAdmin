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

// 解析insert
func (dao *actionDao) ParseInsert(insert []map[string]interface{}, fill ...bool) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		insertData := []map[string]interface{}{}
		for index, item := range insert {
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][dao.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉
					if len(fill) > 0 && fill[0] && !dao.ColumnArrG().Contains(k) {
						continue
					}
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
func (dao *actionDao) ParseUpdate(update map[string]interface{}, fill ...bool) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[dao.Table()+"."+dao.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉
				if len(fill) > 0 && fill[0] && !dao.ColumnArrG().Contains(k) {
					continue
				}
				updateData[dao.Table()+"."+k] = v
			}
		}
		m = m.Data(updateData)
		return m
	}
}

// 解析field
func (dao *actionDao) ParseField(field []string, afterField *[]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			case "id":
				m = m.Fields(dao.Table() + "." + dao.PrimaryKey())
			default:
				if dao.ColumnArrG().Contains(v) {
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
func (dao *actionDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id":
				m = m.Where(dao.Table()+"."+dao.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(dao.Table()+"."+dao.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(dao.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(dao.Table()+".createTime", v)
			default:
				kArr := strings.Split(k, " ")
				if dao.ColumnArrG().Contains(kArr[0]) {
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
func (dao *actionDao) ParseGroup(group []string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(dao.Table() + "." + dao.PrimaryKey())
			default:
				if dao.ColumnArrG().Contains(v) {
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
func (dao *actionDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(dao.Table()+"."+dao.PrimaryKey(), v[1])
			default:
				if dao.ColumnArrG().Contains(v[0]) {
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
func (dao *actionDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
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
