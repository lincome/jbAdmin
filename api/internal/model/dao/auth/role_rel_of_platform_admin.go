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

// internalRoleRelOfPlatformAdminDao is internal type for wrapping internal DAO implements.
type internalRoleRelOfPlatformAdminDao = *internal.RoleRelOfPlatformAdminDao

// roleRelOfPlatformAdminDao is the data access object for table auth_role_rel_of_platform_admin.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelOfPlatformAdminDao struct {
	internalRoleRelOfPlatformAdminDao
}

var (
	// RoleRelOfPlatformAdmin is globally public accessible object for table auth_role_rel_of_platform_admin operations.
	RoleRelOfPlatformAdmin = roleRelOfPlatformAdminDao{
		internal.NewRoleRelOfPlatformAdminDao(),
	}
)

// 解析insert
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoRoleRelOfPlatformAdmin.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(k) {
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
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoRoleRelOfPlatformAdmin.Table()+"."+daoRoleRelOfPlatformAdmin.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoRoleRelOfPlatformAdmin.Table()+"."+k] = v
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
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoRoleRelOfPlatformAdmin.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoRoleRelOfPlatformAdmin.Table() + "." + daoRoleRelOfPlatformAdmin.PrimaryKey() + " AS " + v)
			default:
				if daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(v) {
					m = m.Fields(daoRoleRelOfPlatformAdmin.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoRoleRelOfPlatformAdmin.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoRoleRelOfPlatformAdmin.Table()+"."+daoRoleRelOfPlatformAdmin.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoRoleRelOfPlatformAdmin.Table()+"."+daoRoleRelOfPlatformAdmin.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoRoleRelOfPlatformAdmin.Table()+"."+daoRoleRelOfPlatformAdmin.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoRoleRelOfPlatformAdmin.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoRoleRelOfPlatformAdmin.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoRoleRelOfPlatformAdmin.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoRoleRelOfPlatformAdmin.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoRoleRelOfPlatformAdmin.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoRoleRelOfPlatformAdmin.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoRoleRelOfPlatformAdmin.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoRoleRelOfPlatformAdmin.Table() + "." + daoRoleRelOfPlatformAdmin.PrimaryKey())
			default:
				if daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(v) {
					m = m.Group(daoRoleRelOfPlatformAdmin.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoRoleRelOfPlatformAdmin.Table()+"."+daoRoleRelOfPlatformAdmin.PrimaryKey(), v[1])
			default:
				if daoRoleRelOfPlatformAdmin.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoRoleRelOfPlatformAdmin.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(joinCode) {
			*joinTableArr = append(*joinTableArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin(xxxx.Table(), xxxx.Table()+"."+xxxx.PrimaryKey()+" = "+daoRoleRelOfPlatformAdmin.Table()+"."+xxxx.PrimaryKey()) */
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoRoleRelOfPlatformAdmin.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoRoleRelOfPlatformAdmin *roleRelOfPlatformAdminDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoRoleRelOfPlatformAdmin.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelOfPlatformAdmin.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
