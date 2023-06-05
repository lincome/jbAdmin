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

// internalMenuDao is internal type for wrapping internal DAO implements.
type internalMenuDao = *internal.MenuDao

// menuDao is the data access object for table auth_menu.
// You can define custom methods on it to extend its functionality as you wish.
type menuDao struct {
	internalMenuDao
}

var (
	// Menu is globally public accessible object for table auth_menu operations.
	Menu = menuDao{
		internal.NewMenuDao(),
	}
)

// 解析insert
func (daoMenu *menuDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoMenu.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoMenu.ColumnArrG().Contains(k) {
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
func (daoMenu *menuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoMenu.Table()+"."+daoMenu.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoMenu.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoMenu.Table()+"."+k] = v
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
func (daoMenu *menuDao) ParseField(field []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoMenu.ParseJoin("xxxx", joinCodeArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoMenu.Table() + "." + daoMenu.PrimaryKey() + " AS " + v)
			case "sceneName":
				m = m.Fields(Scene.Table() + "." + v)
				m = daoMenu.ParseJoin("scene", joinCodeArr)(m)
			case "pMenuName":
				m = m.Fields("p_" + daoMenu.Table() + ".menuName AS " + v)
				m = daoMenu.ParseJoin("pMenu", joinCodeArr)(m)
			default:
				if daoMenu.ColumnArrG().Contains(v) {
					m = m.Fields(daoMenu.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoMenu.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoMenu *menuDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoMenu.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoMenu.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoMenu.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoMenu.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoMenu.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoMenu.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoMenu.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoMenu.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoMenu *menuDao) ParseGroup(group []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoMenu.Table() + "." + daoMenu.PrimaryKey())
			default:
				if daoMenu.ColumnArrG().Contains(v) {
					m = m.Group(daoMenu.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoMenu *menuDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v[1])
			default:
				if daoMenu.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoMenu.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoMenu *menuDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if !garray.NewStrArrayFrom(*joinCodeArr).Contains(joinCode) {
			*joinCodeArr = append(*joinCodeArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin("xxxx", "xxxx."+dao.PrimaryKey()+" = "+dao.Table()+"."+dao.PrimaryKey()) */
			case "scene":
				m = m.LeftJoin(Scene.Table(), Scene.Table()+"."+Scene.PrimaryKey()+" = "+daoMenu.Table()+"."+Scene.PrimaryKey())
			case "pMenu":
				pMenuTable := "p_" + daoMenu.Table()
				m = m.LeftJoin(daoMenu.Table()+" AS "+pMenuTable, pMenuTable+"."+daoMenu.PrimaryKey()+" = "+daoMenu.Table()+".pid")
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoMenu *menuDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoMenu *menuDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoMenu *menuDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
