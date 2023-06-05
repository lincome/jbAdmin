// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
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
func (daoAction *actionDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoAction.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoAction.ColumnArrG().Contains(k) {
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
func (daoAction *actionDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoAction.Table()+"."+daoAction.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoAction.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoAction.Table()+"."+k] = v
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
func (daoAction *actionDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoAction.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoAction.Table() + "." + daoAction.PrimaryKey() + " AS " + v)
			case "sceneIdArr":
				//需要id字段
				m = m.Fields(daoAction.Table() + "." + daoAction.PrimaryKey())
				afterField = append(afterField, v)
			default:
				if daoAction.ColumnArrG().Contains(v) {
					m = m.Fields(daoAction.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoAction.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoAction *actionDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoAction.Table()+"."+daoAction.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoAction.Table()+"."+daoAction.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoAction.Table()+"."+daoAction.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoAction.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoAction.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoAction.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoAction.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoAction.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoAction.Table()+"."+keywordField, v)
				}
			case "sceneId":
				m = m.Where(ActionRelToScene.Table()+"."+k, v)

				m = daoAction.ParseJoin("actionRelToScene", joinTableArr)(m)
			case "selfAction": //获取当前登录身份可用的操作。参数：map[string]interface{}{"sceneCode": "场景标识", "loginId": 登录身份id}
				val := v.(map[string]interface{})
				ctx := m.GetCtx()
				sceneInfo := m.GetCtx().Value("sceneInfo").(gdb.Record)
				sceneId := 0
				if len(sceneInfo) == 0 {
					sceneIdTmp, _ := Scene.Ctx(ctx).Where("sceneCode", val["sceneCode"]).Value("sceneId")
					sceneId = sceneIdTmp.Int()
				} else {
					sceneId = sceneInfo["sceneId"].Int()
				}
				m = m.Where(daoAction.Table()+".isStop", 0)
				m = m.Where(ActionRelToScene.Table()+".sceneId", sceneId)
				m = daoAction.ParseJoin("actionRelToScene", joinTableArr)(m)

				switch val["sceneCode"].(string) {
				case "platformAdmin":
					//if val["loginId"] === getConfig('app.superPlatformAdminId') { //平台超级管理员，不再需要其他条件
					if val["loginId"] == 1 { //平台超级管理员，不再需要其他条件
						return m
					}
					m = m.Where(Role.Table()+".isStop", 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+".adminId", val["loginId"])

					m = daoAction.ParseJoin("roleRelToAction", joinTableArr)(m)
					m = daoAction.ParseJoin("role", joinTableArr)(m)
					m = daoAction.ParseJoin("roleRelOfPlatformAdmin", joinTableArr)(m)
				}
				m = daoAction.ParseGroup([]string{"id"}, joinTableArr)(m)
			default:
				kArr := strings.Split(k, " ")
				if daoAction.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoAction.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoAction *actionDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoAction.Table() + "." + daoAction.PrimaryKey())
			default:
				if daoAction.ColumnArrG().Contains(v) {
					m = m.Group(daoAction.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoAction *actionDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoAction.Table()+"."+daoAction.PrimaryKey(), v[1])
			default:
				if daoAction.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoAction.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoAction *actionDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoAction.PrimaryKey()+" = "+daoAction.Table()+"."+daoAction.PrimaryKey())
		} */
		case "actionRelToScene":
			actionRelToSceneTable := ActionRelToScene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(actionRelToSceneTable) {
				*joinTableArr = append(*joinTableArr, actionRelToSceneTable)
				m = m.LeftJoin(actionRelToSceneTable, actionRelToSceneTable+"."+daoAction.PrimaryKey()+" = "+daoAction.Table()+"."+daoAction.PrimaryKey())
			}
		case "roleRelToAction":
			roleRelToActionTable := RoleRelToAction.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelToActionTable) {
				*joinTableArr = append(*joinTableArr, roleRelToActionTable)
				m = m.LeftJoin(roleRelToActionTable, roleRelToActionTable+"."+daoAction.PrimaryKey()+" = "+daoAction.Table()+"."+daoAction.PrimaryKey())
			}
		case "role":
			roleTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleTable) {
				*joinTableArr = append(*joinTableArr, roleTable)
				roleRelToActionTable := RoleRelToAction.Table()
				m = m.LeftJoin(roleTable, roleTable+"."+Role.PrimaryKey()+" = "+roleRelToActionTable+"."+Role.PrimaryKey())
			}
		case "roleRelOfPlatformAdmin":
			roleRelOfPlatformAdminTable := RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelOfPlatformAdminTable) {
				*joinTableArr = append(*joinTableArr, roleRelOfPlatformAdminTable)
				roleRelToActionTable := RoleRelToAction.Table()
				m = m.LeftJoin(roleRelOfPlatformAdminTable, roleRelOfPlatformAdminTable+"."+Role.PrimaryKey()+" = "+roleRelToActionTable+"."+Role.PrimaryKey())
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoAction *actionDao) AfterField(afterField []string) gdb.HookHandler {
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
					case "sceneIdArr":
						sceneIdArr, _ := RoleRelToMenu.Ctx(ctx).Where("actionId", record[daoAction.PrimaryKey()]).Fields("sceneId").Array()
						record[v] = gvar.New(sceneIdArr)
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// 详情
func (daoAction *actionDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoAction *actionDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
