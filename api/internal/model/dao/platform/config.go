// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/platform/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalConfigDao is internal type for wrapping internal DAO implements.
type internalConfigDao = *internal.ConfigDao

// configDao is the data access object for table platform_config.
// You can define custom methods on it to extend its functionality as you wish.
type configDao struct {
	internalConfigDao
}

var (
	// Config is globally public accessible object for table platform_config operations.
	Config = configDao{
		internal.NewConfigDao(),
	}
)

// 解析insert
func (daoConfig *configDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoConfig.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoConfig.ColumnArrG().Contains(k) {
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
func (daoConfig *configDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoConfig.Table()+"."+daoConfig.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoConfig.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoConfig.Table()+"."+k] = v
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
func (daoConfig *configDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoConfig.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoConfig.Table() + "." + daoConfig.PrimaryKey() + " AS " + v)
			default:
				if daoConfig.ColumnArrG().Contains(v) {
					m = m.Fields(daoConfig.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoConfig.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoConfig *configDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoConfig.Table()+"."+daoConfig.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoConfig.Table()+"."+daoConfig.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoConfig.Table()+"."+daoConfig.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoConfig.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoConfig.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoConfig.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoConfig.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoConfig.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoConfig.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoConfig.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoConfig.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoConfig *configDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoConfig.Table() + "." + daoConfig.PrimaryKey())
			default:
				if daoConfig.ColumnArrG().Contains(v) {
					m = m.Group(daoConfig.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoConfig *configDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoConfig.Table()+"."+daoConfig.PrimaryKey(), v[1])
			default:
				if daoConfig.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoConfig.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoConfig *configDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(joinCode) {
			*joinTableArr = append(*joinTableArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin(xxxx.Table(), xxxx.Table()+"."+xxxx.PrimaryKey()+" = "+daoConfig.Table()+"."+xxxx.PrimaryKey()) */
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoConfig *configDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoConfig *configDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoConfig.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoConfig.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoConfig.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoConfig.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoConfig *configDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoConfig.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoConfig.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoConfig.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoConfig.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
