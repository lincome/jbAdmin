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
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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

// 获取daoModel
func (daoThis *menuDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *menuDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *menuDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *menuDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().MenuId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *menuDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().MenuName
}

// 解析filter
func (daoThis *menuDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().MenuId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().MenuId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().MenuId, v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().MenuName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().MenuName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `p_id_path_of_old`: //父级ID路径（旧）
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().IdPath, gconv.String(v)+`-%`)
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `self_menu`: //获取当前登录身份可用的菜单。参数：map[string]any{`scene_code`: `场景标识`, `login_id`: 登录身份id, `scene_id`: 场景id（平台超级管理员用，不用再做一次查询）}
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().IsStop, 0)
				val := gconv.Map(v)
				switch gconv.String(val[`scene_code`]) {
				case `platform`:
					if gconv.Uint(val[`login_id`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Uint() { //平台超级管理员
						m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().SceneId, val[`scene_id`])
						continue
					}
					roleIdArr, _ := Role.CtxDaoModel(m.GetCtx()).Fields(Role.Columns().RoleId).Filter(`self_role`, val).Array()
					if len(roleIdArr) == 0 {
						m = m.Where(`1 = 0`)
						continue
					}
					/* // 方式一：联表查询（不推荐。原因：auth_role及其关联表，后期表数据只会越来越大，故不建议联表）
					tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelToMenu+`.`+RoleRelToMenu.Columns().RoleId, roleIdArr)
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, daoModel))
					m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().MenuId) */
					// 方式二：非联表查询
					menuIdArr, _ := RoleRelToMenu.CtxDaoModel(m.GetCtx()).Filter(RoleRelToMenu.Columns().RoleId, roleIdArr).Distinct().Array(RoleRelToMenu.Columns().MenuId)
					if len(menuIdArr) == 0 {
						m = m.Where(`1 = 0`)
						continue
					}
					m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().MenuId, menuIdArr)
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
func (daoThis *menuDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
			case Scene.Columns().SceneName:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
			case `p_menu_name`:
				tableP := `p_` + daoModel.DbTable
				m = m.Fields(tableP + `.` + daoThis.Columns().MenuName + ` AS ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoModel))
			case `tree`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuId)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().Pid)
				m = m.Handler(daoThis.ParseOrder([]string{`tree`}, daoModel))
			case `show_menu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuName)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuIcon)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuUrl)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ExtraData)
				// m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ExtraData + `->'$.i18n' AS i18n`)	//mysql5.6版本不支持
				// m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(` + daoThis.Columns().ExtraData + `, \`$.i18n\`)) AS i18n`))	//mysql不能直接转成对象返回
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
func (daoThis *menuDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
					case `show_menu`:
						extraDataJson := gjson.New(record[daoThis.Columns().ExtraData])
						record[`i18n`] = extraDataJson.Get(`i18n`)
						if record[`i18n`] == nil {
							record[`i18n`] = gvar.New(map[string]any{`title`: map[string]any{`zh-cn`: record[daoThis.Columns().MenuName]}})
						}
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
func (daoThis *menuDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if _, ok := insert[daoThis.Columns().Pid]; !ok {
			insert[daoThis.Columns().Pid] = 0
		}
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
			case daoThis.Columns().Pid:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoModel.CloneNew().Filter(daoThis.Columns().MenuId, v).One()
					daoModel.AfterInsert[`self_update`] = map[string]any{
						`p_id_path`: pInfo[daoThis.Columns().IdPath].String(),
						`p_level`:   pInfo[daoThis.Columns().Level].Uint(),
					}
				} else {
					daoModel.AfterInsert[`self_update`] = map[string]any{
						`p_id_path`: `0`,
						`p_level`:   0,
					}
				}
			case daoThis.Columns().ExtraData:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
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
func (daoThis *menuDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `self_update`: //更新自身的ID路径和层级。参数：map[string]any{`p_id_path`: `父级ID路径`, `p_level`: `父级层级`}
					val := v.(map[string]any)
					daoModel.CloneNew().Filter(daoThis.Columns().MenuId, id).HookUpdate(map[string]any{
						daoThis.Columns().IdPath: gconv.String(val[`p_id_path`]) + `-` + gconv.String(id),
						daoThis.Columns().Level:  gconv.Uint(val[`p_level`]) + 1,
					}).Update()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *menuDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			case daoThis.Columns().Pid:
				updateData[k] = v
				pIdPath := `0`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoModel.CloneNew().Filter(daoThis.Columns().MenuId, v).One()
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Uint()
				}
				updateData[daoThis.Columns().IdPath] = gdb.Raw(`CONCAT('` + pIdPath + `-', ` + daoThis.Columns().MenuId + `)`)
				updateData[daoThis.Columns().Level] = pLevel + 1
				//更新所有子孙级的ID路径和层级
				childUpdateList := []map[string]any{}
				oldList, _ := daoModel.CloneNew().Filter(`id`, daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().Pid].Uint() {
						childUpdateList = append(childUpdateList, map[string]any{
							`p_id_path_of_old`: oldInfo[daoThis.Columns().IdPath],
							`p_id_path_of_new`: pIdPath + `-` + oldInfo[daoThis.Columns().MenuId].String(),
							`p_level_of_old`:   oldInfo[daoThis.Columns().Level],
							`p_level_of_new`:   pLevel + 1,
						})
					}
				}
				if len(childUpdateList) > 0 {
					daoModel.AfterUpdate[`child_update_list`] = childUpdateList
				}
			case `child_id_path`: //更新所有子孙级的ID路径。参数：map[string]any{`p_id_path_of_old`: `父级ID路径（旧）`, `p_id_path_of_new`: `父级ID路径（新）`}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[`p_id_path_of_old`])
				pIdPathOfNew := gconv.String(val[`p_id_path_of_new`])
				updateData[daoThis.Columns().IdPath] = gdb.Raw(`REPLACE(` + daoThis.Columns().IdPath + `, '` + pIdPathOfOld + `', '` + pIdPathOfNew + `')`)
			case `child_level`: //更新所有子孙级的层级。参数：map[string]any{`p_level_of_old`: `父级层级（旧）`, `p_level_of_new`: `父级层级（新）`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`p_level_of_old`])
				pLevelOfNew := gconv.Uint(val[`p_level_of_new`])
				updateData[daoThis.Columns().Level] = gdb.Raw(daoModel.DbTable + `.` + daoThis.Columns().Level + ` + ` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoThis.Columns().Level] = gdb.Raw(daoModel.DbTable + `.` + daoThis.Columns().Level + ` - ` + gconv.String(pLevelOfOld-pLevelOfNew))
				}
			case daoThis.Columns().ExtraData:
				if gconv.String(v) == `` {
					updateData[k] = nil
					continue
				}
				updateData[k] = v
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
func (daoThis *menuDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `child_update_list`: //修改pid时，更新所有子孙级的ID路径和层级。参数：[]map[string]any{`p_id_path_of_old`: `父级ID路径（旧）`, `p_id_path_of_new`: `父级ID路径（新）`, `p_level_of_old`: `父级层级（旧）`, `p_level_of_new`: `父级层级（新）`}
					val := v.([]map[string]any)
					for _, v1 := range val {
						daoModel.CloneNew().Filter(`p_id_path_of_old`, v1[`p_id_path_of_old`]).HookUpdate(g.Map{
							`child_id_path`: g.Map{
								`p_id_path_of_old`: v1[`p_id_path_of_old`],
								`p_id_path_of_new`: v1[`p_id_path_of_new`],
							},
							`child_level`: g.Map{
								`p_level_of_old`: v1[`p_level_of_old`],
								`p_level_of_new`: v1[`p_level_of_new`],
							},
						}).Update()
					}
				}
			}
			return
		},
	}
}

// hook delete
func (daoThis *menuDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) { //有软删除字段时需改成Update事件
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			return
		},
	}
}

// 解析group
func (daoThis *menuDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().MenuId)
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
func (daoThis *menuDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().MenuId, 1))
			case `tree`:
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.Columns().Pid)
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.Columns().Sort)
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.Columns().MenuId)
			case daoThis.Columns().Level:
				m = m.Order(daoModel.DbTable + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.Columns().MenuId)
			case daoThis.Columns().Sort:
				m = m.Order(daoModel.DbTable + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.Columns().MenuId)
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
func (daoThis *menuDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case Scene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Scene.Columns().SceneId+` = `+daoModel.DbTable+`.`+daoThis.Columns().SceneId)
		case `p_` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+` AS `+joinTable, joinTable+`.`+daoThis.Columns().MenuId+` = `+daoModel.DbTable+`.`+daoThis.Columns().Pid)
		case RoleRelToMenu.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable+` AS `+joinTable, joinTable+`.`+RoleRelToMenu.Columns().MenuId+` = `+daoModel.DbTable+`.`+daoThis.Columns().MenuId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().MenuId+` = `+daoModel.DbTable+`.`+daoThis.Columns().MenuId)
		}
		return m
	}
}

// Fill with you ideas below.
