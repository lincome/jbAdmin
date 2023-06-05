// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

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
func (daoRoleRelToAction *roleRelToActionDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoRoleRelToAction.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoRoleRelToAction.ColumnArrG().Contains(k) {
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
func (daoRoleRelToAction *roleRelToActionDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoRoleRelToAction.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoRoleRelToAction.Table()+"."+k] = v
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			fieldArr = append(fieldArr, k+" = ?")
			valueArr = append(valueArr, v)
		}
		data := []interface{}{strings.Join(fieldArr, ",")}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// 解析field
func (daoRoleRelToAction *roleRelToActionDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoRoleRelToAction.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoRoleRelToAction.Table() + "." + daoRoleRelToAction.PrimaryKey() + " AS " + v)
			default:
				if daoRoleRelToAction.ColumnArrG().Contains(v) {
					m = m.Fields(daoRoleRelToAction.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoRoleRelToAction.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoRoleRelToAction *roleRelToActionDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoRoleRelToAction.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoRoleRelToAction.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoRoleRelToAction.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoRoleRelToAction.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoRoleRelToAction.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoRoleRelToAction.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoRoleRelToAction.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoRoleRelToAction.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoRoleRelToAction *roleRelToActionDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoRoleRelToAction.Table() + "." + daoRoleRelToAction.PrimaryKey())
			default:
				if daoRoleRelToAction.ColumnArrG().Contains(v) {
					m = m.Group(daoRoleRelToAction.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoRoleRelToAction *roleRelToActionDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey(), v[1])
			default:
				if daoRoleRelToAction.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoRoleRelToAction.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoRoleRelToAction *roleRelToActionDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoRoleRelToAction.PrimaryKey()+" = "+daoRoleRelToAction.Table()+"."+daoRoleRelToAction.PrimaryKey())
		} */
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoRoleRelToAction *roleRelToActionDao) AfterField(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for i, record := range result {
				for _, v := range afterField {
					switch v {
					/* case "xxxx":
					record[v] = gvar.New("") */
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// 详情
func (daoRoleRelToAction *roleRelToActionDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoRoleRelToAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoRoleRelToAction *roleRelToActionDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoRoleRelToAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelToAction.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
