// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalRoleDao is internal type for wrapping internal DAO implements.
type internalRoleDao = *internal.RoleDao

// roleDao is the data access object for table auth_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	internalRoleDao
}

var (
	// Role is globally public accessible object for table auth_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// 解析insert
func (daoRole *roleDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoRole.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoRole.ColumnArrG().Contains(k) {
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
func (daoRole *roleDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoRole.Table()+"."+daoRole.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoRole.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoRole.Table()+"."+k] = v
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
func (daoRole *roleDao) ParseField(field []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoRole.ParseJoin("xxxx", joinCodeArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoRole.Table() + "." + daoRole.PrimaryKey() + " AS " + v)
			default:
				if daoRole.ColumnArrG().Contains(v) {
					m = m.Fields(daoRole.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoRole.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoRole *roleDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoRole.Table()+"."+daoRole.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoRole.Table()+"."+daoRole.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoRole.Table()+"."+daoRole.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoRole.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoRole.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoRole.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoRole.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoRole.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoRole.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoRole.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoRole.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoRole *roleDao) ParseGroup(group []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoRole.Table() + "." + daoRole.PrimaryKey())
			default:
				if daoRole.ColumnArrG().Contains(v) {
					m = m.Group(daoRole.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoRole *roleDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoRole.Table()+"."+daoRole.PrimaryKey(), v[1])
			default:
				if daoRole.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoRole.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoRole *roleDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if !garray.NewStrArrayFrom(*joinCodeArr).Contains(joinCode) {
			*joinCodeArr = append(*joinCodeArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin(xxxx.Table(), xxxx.Table()+"."+xxxx.PrimaryKey()+" = "+daoRole.Table()+"."+xxxx.PrimaryKey()) */
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoRole *roleDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoRole *roleDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRole.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoRole *roleDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRole.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
