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
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalSceneDao is internal type for wrapping internal DAO implements.
type internalSceneDao = *internal.SceneDao

// sceneDao is the data access object for table auth_scene.
// You can define custom methods on it to extend its functionality as you wish.
type sceneDao struct {
	internalSceneDao
}

var (
	// Scene is globally public accessible object for table auth_scene operations.
	Scene = sceneDao{
		internal.NewSceneDao(),
	}
)

// 获取daoModel
func (daoThis *sceneDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *sceneDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *sceneDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *sceneDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().SceneId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *sceneDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().SceneName
}

// 解析filter
func (daoThis *sceneDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().SceneId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().SceneId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().SceneId, v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().SceneName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().SceneName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
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
func (daoThis *sceneDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */
			case `id`:
				m = m.Fields(daoThis.ParseId(daoModel) + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.ParseLabel(daoModel) + ` AS ` + v)
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
func (daoThis *sceneDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil || len(result) == 0 {
				return
			}

			var wg sync.WaitGroup
			wg.Add(len(result))
			afterFieldHandleFunc := func(record gdb.Record) {
				defer wg.Done()
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
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
			for _, record := range result {
				go afterFieldHandleFunc(record)
			}
			wg.Wait()
			return
		},
	}
}

// 解析insert
func (daoThis *sceneDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			default:
				if daoThis.ColumnArr().Contains(k) {
					insertData[k] = v
				}
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
func (daoThis *sceneDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case `xxxx`:
					daoModel.CloneNew().Filter(`id`, id).HookUpdate(g.Map{k: v}).Update()
				}
			} */
			return
		},
	}
}

// 解析update
func (daoThis *sceneDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[k] = v
				}
			}
		}
		m = m.Data(updateData)
		return m
	}
}

// hook update
func (daoThis *sceneDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case `xxxx`:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().Filter(`id`, id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *sceneDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) { //有软删除字段时需改成Update事件
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			/* 高并发环境中并不能确保以下这些表不会产生脏数据（因权限模块基本不会让太多人使用，故不对并发做严格处理：情况2有几率出现）
			举例：
				请求A：删除场景->删除前验证关联菜单是否存在->删除
				请求B：新增菜单->新增前验证关联场景是否存在->新增
			两个请求的步骤2在并发时可能都验证成功，此时就存在以下两种情况：
				情况1：请求B先新增，请求A后删除，会导致菜单表插入一条脏数据（）
					解决方法：做删除后置处理，以下代码就是
				情况2：请求A先删除，请求B后新增，会导致菜单表插入一条脏数据
					解决方法：菜单表做触发器，插入前判断场景是否被删除（程序中的判断不能解决并发问题，但数据库层面可以解决）

				通用解决方法（对情况1和情况2都有效）：
					1、请求A和请求B都使用事务，且读取场景表时，设置排它锁
					2、菜单表做外键约束
			*/
			Menu.CtxDaoModel(ctx).Filter(Menu.Columns().SceneId, daoModel.IdArr).Delete()
			ActionRelToScene.CtxDaoModel(ctx).Filter(ActionRelToScene.Columns().SceneId, daoModel.IdArr).Delete()
			Role.CtxDaoModel(ctx).Filter(Role.Columns().SceneId, daoModel.IdArr).Delete()
			return
		},
	}
}

// 解析group
func (daoThis *sceneDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().SceneId)
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
func (daoThis *sceneDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().SceneId, 1))
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
func (daoThis *sceneDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().SceneId+` = `+daoModel.DbTable+`.`+daoThis.Columns().SceneId)
		}
		return m
	}
}

// Fill with you ideas below.
