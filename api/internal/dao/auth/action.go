// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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

// 获取daoModel
func (daoThis *actionDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *actionDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *actionDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析filter
func (daoThis *actionDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().ActionName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().ActionName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case ActionRelToScene.Columns().SceneId:
				tableActionRelToScene := ActionRelToScene.ParseDbTable(m.GetCtx())
				m = m.Where(tableActionRelToScene+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableActionRelToScene, daoModel))
			case `self_action`: //获取当前登录身份可用的操作。参数：map[string]interface{}{`scene_code`: `场景标识`, `scene_id`=>场景id, `login_id`: 登录身份id}
				val := gconv.Map(v)
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().IsStop, 0)
				tableActionRelToScene := ActionRelToScene.ParseDbTable(m.GetCtx())
				m = m.Where(tableActionRelToScene+`.`+ActionRelToScene.Columns().SceneId, val[`scene_id`])
				m = m.Handler(daoThis.ParseJoin(tableActionRelToScene, daoModel))
				switch gconv.String(val[`scene_code`]) {
				case `platform`:
					if gconv.Uint(val[`login_id`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Uint() { //平台超级管理员，不再需要其它条件
						continue
					}
					tableRole := Role.ParseDbTable(m.GetCtx())
					m = m.Where(tableRole+`.`+Role.Columns().IsStop, 0)
					tableRoleRelToAction := RoleRelToAction.ParseDbTable(m.GetCtx())
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToAction, daoModel))
					m = m.Handler(daoThis.ParseJoin(tableRole, daoModel))

					tableRoleRelOfPlatformAdmin := RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelOfPlatformAdmin+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`login_id`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoModel))
				default:
					m = m.Where(`1 = 0`)
				}
			default:
				if daoThis.ColumnArr().Contains(k) {
					m = m.Where(daoModel.DbTable+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析field
func (daoThis *actionDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */
			case `id`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ActionName + ` AS ` + v)
			case `scene_id_arr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey())
				daoModel.AfterField.Add(v)
			default:
				if daoThis.ColumnArr().Contains(v) {
					m = m.Fields(daoModel.DbTable + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				daoModel.AfterFieldWithParam[k] = v
			}
		}
		return m
	}
}

// hook select
func (daoThis *actionDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
					case `scene_id_arr`:
						sceneIdArr, _ := ActionRelToScene.CtxDaoModel(ctx).Filter(ActionRelToScene.Columns().ActionId, record[daoThis.PrimaryKey()]).Array(ActionRelToScene.Columns().SceneId)
						record[v] = gvar.New(sceneIdArr)
					default:
						record[v] = gvar.New(nil)
					}
				}
				/* for k, v := range daoModel.AfterFieldWithParam {
					switch k {
					case `xxxx`:
						record[k] = gvar.New(v)
					}
				} */
			}
			return
		},
	}
}

// 解析insert
func (daoThis *actionDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `scene_id_arr`:
				daoModel.AfterInsert[k] = v
			default:
				if !daoThis.ColumnArr().Contains(k) {
					continue
				}
				insertData[k] = v
			}
		}
		m = m.Data(insertData)
		if len(daoModel.AfterInsert) > 0 {
			m = m.Hook(daoThis.HookInsert(daoModel))
		}
		return m
	}
}

// hook insert
func (daoThis *actionDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `scene_id_arr`:
					insertList := []map[string]interface{}{}
					for _, item := range gconv.SliceAny(v) {
						insertList = append(insertList, map[string]interface{}{
							ActionRelToScene.Columns().ActionId: id,
							ActionRelToScene.Columns().SceneId:  item,
						})
					}
					ActionRelToScene.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *actionDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `scene_id_arr`:
				daoModel.AfterUpdate[k] = v
			default:
				if !daoThis.ColumnArr().Contains(k) {
					continue
				}
				updateData[k] = v
			}
		}
		m = m.Data(updateData)
		return m
	}
}

// hook update
func (daoThis *actionDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if daoIndex.IsEmptyDataOfUpdate(ctx, daoModel.DbGroup, in.Data) {
				result = driver.RowsAffected(0)
			} else {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `scene_id_arr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &ActionRelToScene, ActionRelToScene.Columns().ActionId, ActionRelToScene.Columns().SceneId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.SliceStr(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &ActionRelToScene, ActionRelToScene.Columns().ActionId, ActionRelToScene.Columns().SceneId, id, valArr)
					}
				}
			}

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case `xxxx`:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *actionDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			ActionRelToScene.CtxDaoModel(ctx).Filter(ActionRelToScene.Columns().ActionId, daoModel.IdArr).Delete()
			RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().ActionId, daoModel.IdArr).Delete()
			return
		},
	}
}

// 解析group
func (daoThis *actionDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArr().Contains(v) {
					m = m.Group(daoModel.DbTable + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *actionDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			default:
				if daoThis.ColumnArr().Contains(k) {
					m = m.Order(daoModel.DbTable + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *actionDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		case ActionRelToScene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+ActionRelToScene.Columns().ActionId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		case Role.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Role.PrimaryKey()+` = `+RoleRelToAction.ParseDbTable(m.GetCtx())+`.`+RoleRelToAction.Columns().RoleId)
		case RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+RoleRelToAction.ParseDbTable(m.GetCtx())+`.`+RoleRelToAction.Columns().RoleId)
		case RoleRelToAction.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToAction.Columns().ActionId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.
