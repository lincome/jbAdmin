// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/log/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalRequestDao is internal type for wrapping internal DAO implements.
type internalRequestDao = *internal.RequestDao

// requestDao is the data access object for table log_request.
// You can define custom methods on it to extend its functionality as you wish.
type requestDao struct {
	internalRequestDao
}

var (
	// Request is globally public accessible object for table log_request operations.
	Request = requestDao{
		internal.NewRequestDao(),
	}
)

// 解析insert
func (daoRequest *requestDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoRequest.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoRequest.ColumnArrG().Contains(k) {
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
func (daoRequest *requestDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoRequest.Table()+"."+daoRequest.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoRequest.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoRequest.Table()+"."+k] = v
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
func (daoRequest *requestDao) ParseField(field []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoRequest.ParseJoin("xxxx", joinCodeArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoRequest.Table() + "." + daoRequest.PrimaryKey() + " AS " + v)
			default:
				if daoRequest.ColumnArrG().Contains(v) {
					m = m.Fields(daoRequest.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoRequest.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoRequest *requestDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoRequest.Table()+"."+daoRequest.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoRequest.Table()+"."+daoRequest.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoRequest.Table()+"."+daoRequest.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoRequest.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoRequest.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoRequest.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoRequest.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoRequest.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoRequest.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoRequest.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoRequest.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoRequest *requestDao) ParseGroup(group []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoRequest.Table() + "." + daoRequest.PrimaryKey())
			default:
				if daoRequest.ColumnArrG().Contains(v) {
					m = m.Group(daoRequest.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoRequest *requestDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoRequest.Table()+"."+daoRequest.PrimaryKey(), v[1])
			default:
				if daoRequest.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoRequest.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoRequest *requestDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if !garray.NewStrArrayFrom(*joinCodeArr).Contains(joinCode) {
			*joinCodeArr = append(*joinCodeArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin(xxxx.Table(), xxxx.Table()+"."+xxxx.PrimaryKey()+" = "+daoRequest.Table()+"."+xxxx.PrimaryKey()) */
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoRequest *requestDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoRequest *requestDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinCodeArr := []string{}
	model := daoRequest.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRequest.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRequest.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRequest.ParseOrder(order, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoRequest *requestDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinCodeArr := []string{}
	model := daoRequest.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRequest.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRequest.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRequest.ParseOrder(order, &joinCodeArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
