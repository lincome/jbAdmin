// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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

// 解析分库
func (daoThis *roleDao) ParseDbGroup(ctx context.Context, dbGroupSelData ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSelData) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableSelData) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *roleDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx, dbSelDataList[1])).Ctx(ctx)
	default:
		return g.DB(daoThis.ParseDbGroup(ctx)).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case `menuIdArr`, `actionIdArr`:
				hookData[k] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					insertData[k] = v
				}
			}
		}
		m = m.Data(insertData)
		if len(hookData) > 0 {
			m = m.Hook(daoThis.HookInsert(hookData))
		}
		return m
	}
}

// hook insert
func (daoThis *roleDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range data {
				switch k {
				case `menuIdArr`:
					daoThis.SaveRelMenu(ctx, gconv.SliceUint(v), uint(id))
				case `actionIdArr`:
					daoThis.SaveRelAction(ctx, gconv.SliceUint(v), uint(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *roleDao) ParseUpdate(update map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[tableThis+`.`+daoThis.PrimaryKey()] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					updateData[tableThis+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
				}
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			_, ok := v.(gdb.Raw)
			if ok {
				fieldArr = append(fieldArr, k+` = `+gconv.String(v))
			} else {
				fieldArr = append(fieldArr, k+` = ?`)
				valueArr = append(valueArr, v)
			}
		}
		data := []interface{}{gstr.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// hook update
func (daoThis *roleDao) HookUpdate(data map[string]interface{}, idArr ...uint) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			/* //不能这样拿idArr，联表时会有bug
			var idArr []*gvar.Var
			if len(data) > 0 {
				idArr, _ = daoThis.ParseDbCtx(ctx).Where(in.Condition, in.Args[len(in.Args)-gstr.Count(in.Condition, `?`):]...).Array(daoThis.PrimaryKey())
			} */
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range data {
				switch k {
				case `menuIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range idArr {
						daoThis.SaveRelMenu(ctx, relIdArr, id)
					}
				case `actionIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range idArr {
						daoThis.SaveRelAction(ctx, relIdArr, id)
					}
				}
			}

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return
		},
	}
}

// hook delete
func (daoThis *roleDao) HookDelete(idArr ...uint) gdb.HookHandler {
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

			RoleRelToMenu.ParseDbCtx(ctx).Where(RoleRelToMenu.Columns().RoleId, idArr).Delete()
			RoleRelToAction.ParseDbCtx(ctx).Where(RoleRelToAction.Columns().RoleId, idArr).Delete()
			RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(RoleRelOfPlatformAdmin.Columns().RoleId, idArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *roleDao) ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.ParseDbTable(ctx), joinTableArr))
			*afterField = append(*afterField, v) */
			case `id`:
				m = m.Fields(tableThis + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(tableThis + `.` + daoThis.Columns().RoleName + ` AS ` + v)
			case Scene.Columns().SceneName: //因前端页面已用该字段名显示，故不存在时改成`sceneName`（控制器也要改）。同时下面Fields方法改成m = m.Fields(tableUser + `.` + Scene.Columns().Xxxx + ` AS ` + v)
				tableScene := Scene.ParseDbTable(ctx)
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, joinTableArr))
			case `menuIdArr`, `actionIdArr`:
				m = m.Fields(tableThis + `.` + daoThis.PrimaryKey())
				*afterField = append(*afterField, v)
			case `tableName`:
				m = m.Fields(tableThis + `.` + daoThis.Columns().TableId)
				tableScene := Scene.ParseDbTable(ctx)
				m = m.Fields(tableScene + `.` + Scene.Columns().SceneCode)
				m = m.Handler(daoThis.ParseJoin(tableScene, joinTableArr))
				*afterField = append(*afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(tableThis + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				afterFieldWithParam[k] = v
			}
		}
		return m
	}
}

// hook select
func (daoThis *roleDao) HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range *afterField {
					switch v {
					case `menuIdArr`:
						idArr, _ := RoleRelToMenu.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToMenu.Columns().MenuId)
						record[v] = gvar.New(idArr)
					case `actionIdArr`:
						idArr, _ := RoleRelToAction.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToAction.Columns().ActionId)
						record[v] = gvar.New(idArr)
					case `tableName`:
						if record[daoThis.Columns().TableId].Uint() == 0 {
							record[v] = gvar.New(`平台`)
							continue
						}
						switch record[Scene.Columns().SceneCode].String() {
						case `platform`:
						}
					default:
						record[v] = gvar.New(nil)
					}
				}
				/* for k, v := range afterFieldWithParam {
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

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for k, v := range filter {
			switch k {
			case `excId`, `excIdArr`:
				val := gconv.SliceUint(v)
				switch len(val) {
				case 0: //gconv.SliceUint会把0转换成[]uint{}，故不能用转换后的val。必须用原始数据v
					m = m.WhereNot(tableThis+`.`+daoThis.PrimaryKey(), v)
				case 1:
					m = m.WhereNot(tableThis+`.`+daoThis.PrimaryKey(), val[0])
				default:
					m = m.WhereNotIn(tableThis+`.`+daoThis.PrimaryKey(), v)
				}
			case `id`, `idArr`:
				m = m.Where(tableThis+`.`+daoThis.PrimaryKey(), v)
			case `label`:
				m = m.WhereLike(tableThis+`.`+daoThis.Columns().RoleName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().RoleName:
				m = m.WhereLike(tableThis+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(tableThis+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(tableThis+`.`+daoThis.Columns().CreatedAt, v)
			case Scene.Columns().SceneCode:
				tableScene := Scene.ParseDbTable(ctx)
				m = m.Where(tableScene+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableScene, joinTableArr))
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Where(tableThis+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *roleDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(tableThis + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(tableThis + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *roleDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(tableThis + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Order(tableThis + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *roleDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(*joinTableArr).Contains(joinCode) {
			return m
		}
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		*joinTableArr = append(*joinTableArr, joinCode)
		switch joinCode {
		/* case Xxxx.ParseDbTable(ctx):
		m = m.LeftJoin(joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+tableThis+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(ctx)+` AS `+joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+tableThis+`.`+daoThis.PrimaryKey()) */
		case Scene.ParseDbTable(ctx):
			m = m.LeftJoin(joinCode, joinCode+`.`+Scene.PrimaryKey()+` = `+tableThis+`.`+daoThis.Columns().SceneId)
		default:
			m = m.LeftJoin(joinCode, joinCode+`.`+daoThis.PrimaryKey()+` = `+tableThis+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联菜单
func (daoThis *roleDao) SaveRelMenu(ctx context.Context, relIdArr []uint, id uint) {
	relDao := RoleRelToMenu
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().MenuId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceUint(relIdArrOfOldTmp)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewFrom(relIdArr).Diff(gset.NewFrom(relIdArrOfOld)).Slice()
	if len(insertRelIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertRelIdArr {
			insertList = append(insertList, map[string]interface{}{
				priKey: id,
				relKey: v,
			})
		}
		relDao.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联 结束----**/

	/**----删除关联 开始----**/
	deleteRelIdArr := gset.NewFrom(relIdArrOfOld).Diff(gset.NewFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联 结束----**/
}

// 保存关联操作
func (daoThis *roleDao) SaveRelAction(ctx context.Context, relIdArr []uint, id uint) {
	relDao := RoleRelToAction
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().ActionId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceUint(relIdArrOfOldTmp)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewFrom(relIdArr).Diff(gset.NewFrom(relIdArrOfOld)).Slice()
	if len(insertRelIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertRelIdArr {
			insertList = append(insertList, map[string]interface{}{
				priKey: id,
				relKey: v,
			})
		}
		relDao.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联 开始----**/

	/**----删除关联 结束----**/
	deleteRelIdArr := gset.NewFrom(relIdArrOfOld).Diff(gset.NewFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联 结束----**/
}
