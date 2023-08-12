// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"api/internal/utils"
	"context"
	"database/sql"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
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
func (daoThis *roleDao) ParseDbGroup(ctx context.Context, dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSeldata) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(ctx context.Context, dbTableSelData map[string]interface{}) string {
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
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx, g.Map{})).Safe().Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx, dbSelDataList[1])).Safe().Ctx(ctx)
	default:
		return g.DB(daoThis.ParseDbGroup(ctx, g.Map{})).Model(daoThis.ParseDbTable(ctx, g.Map{})).Safe().Ctx(ctx)
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert map[string]interface{}, fill ...bool) gdb.ModelHandler {
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
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				insertData[k] = v
			}
		}
		m = m.Data(insertData).Hook(daoThis.HookInsert(hookData))
		return m
	}
}

// hook insert
func (daoThis *roleDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range data {
				switch k {
				case `menuIdArr`:
					daoThis.SaveRelMenu(ctx, gconv.SliceInt(v), int(id))
				case `actionIdArr`:
					daoThis.SaveRelAction(ctx, gconv.SliceInt(v), int(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *roleDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
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
		data := []interface{}{strings.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// hook update
func (daoThis *roleDao) HookUpdate(data map[string]interface{}, idArr ...int) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			/* //不能这样拿idArr，联表时会有bug
			var idArr []*gvar.Var
			if len(data) > 0 {
				idArr, _ = daoThis.ParseDbCtx(ctx).Where(in.Condition, in.Args[len(in.Args)-gstr.Count(in.Condition, `?`):]...).Array(daoThis.PrimaryKey())
			} */
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}

			for k, v := range data {
				switch k {
				case `menuIdArr`:
					relIdArr := gconv.SliceInt(v)
					for _, id := range idArr {
						daoThis.SaveRelMenu(ctx, relIdArr, id)
					}
				case `actionIdArr`:
					relIdArr := gconv.SliceInt(v)
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
func (daoThis *roleDao) HookDelete(idArr ...int) gdb.HookHandler {
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

			RoleRelToMenu.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
			RoleRelToAction.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
			RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *roleDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(Xxxx.Table(), joinTableArr)(m)
			afterField = append(afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().RoleName + ` AS ` + v)
			case `sceneName`:
				m = m.Fields(Scene.Table() + `.` + v)
				m = daoThis.ParseJoin(Scene.Table(), joinTableArr)(m)
			case `menuIdArr`, `actionIdArr`:
				//需要id字段
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				afterField = append(afterField, v)
			case `tableName`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().TableId)
				m = m.Fields(Scene.Table() + `.` + Scene.Columns().SceneCode)
				m = daoThis.ParseJoin(Scene.Table(), joinTableArr)(m)
				afterField = append(afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.HookSelect(afterField))
		}
		return m
	}
}

// hook select
func (daoThis *roleDao) HookSelect(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for index, record := range result {
				for _, v := range afterField {
					switch v {
					/* case `xxxx`:
					record[v] = gvar.New(``) */
					case `menuIdArr`:
						idArr, _ := RoleRelToMenu.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToMenu.Columns().MenuId)
						record[v] = gvar.New(idArr)
					case `actionIdArr`:
						idArr, _ := RoleRelToAction.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToAction.Columns().ActionId)
						record[v] = gvar.New(idArr)
					case `tableName`:
						if record[daoThis.Columns().TableId].Int() == 0 {
							record[v] = gvar.New(`平台`)
							continue
						}
						switch record[Scene.Columns().SceneCode].String() {
						case `platform`:
						}
					}
				}
				result[index] = record
			}
			return
		},
	}
}

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `id`, `idArr`:
				m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
			case `excId`, `excIdArr`:
				val := gconv.SliceInt(v)
				switch len(val) {
				case 0: //gconv.SliceInt会把0转换成[]int{}，故不能用转换后的val。必须用原始数据v
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				case 1:
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), val[0])
				default:
					m = m.WhereNotIn(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				}
			case `timeRangeStart`:
				m = m.WhereGTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `label`:
				m = m.WhereLike(daoThis.Table()+`.`+daoThis.Columns().RoleName, `%`+gconv.String(v)+`%`)
			case `sceneCode`:
				m = m.Where(Scene.Table()+`.`+Scene.Columns().SceneCode, v)
				m = daoThis.ParseJoin(Scene.Table(), joinTableArr)(m)
			default:
				kArr := strings.Split(k, ` `) //支持`id > ?`等k
				if !daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(k, v)
					continue
				}
				if len(kArr) == 1 {
					if gstr.SubStr(gstr.CaseCamel(kArr[0]), -4) == `Name` {
						m = m.WhereLike(daoThis.Table()+`.`+k, `%`+gconv.String(v)+`%`)
						continue
					}
				}
				m = m.Where(daoThis.Table()+`.`+k, v)
			}
		}
		return m
	}
}

// 解析group
func (daoThis *roleDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoThis.Table() + `.` + v)
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
		for _, v := range order {
			kArr := strings.Split(v, ` `)
			if len(kArr) == 1 {
				kArr = append(kArr, `ASC`)
			}
			switch kArr[0] {
			case `id`:
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), kArr[1])
			default:
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Order(daoThis.Table()+`.`+kArr[0], kArr[1])
				} else {
					m = m.Order(kArr[0], kArr[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *roleDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case Xxxx.Table():
		relTable := Xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
			*joinTableArr = append(*joinTableArr, relTable)
			m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case Scene.Table():
			relTable := Scene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+Scene.PrimaryKey()+` = `+daoThis.Table()+`.`+Scene.PrimaryKey())
			}
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联菜单
func (daoThis *roleDao) SaveRelMenu(ctx context.Context, relIdArr []int, id int) {
	relDao := RoleRelToMenu
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().MenuId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceInt(relIdArrOfOldTmp)

	/**----新增关联菜单 开始----**/
	insertRelIdArr := gset.NewIntSetFrom(relIdArr).Diff(gset.NewIntSetFrom(relIdArrOfOld)).Slice()
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
	/**----新增关联菜单 结束----**/

	/**----删除关联菜单 开始----**/
	deleteRelIdArr := gset.NewIntSetFrom(relIdArrOfOld).Diff(gset.NewIntSetFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联菜单 结束----**/
}

// 保存关联操作
func (daoThis *roleDao) SaveRelAction(ctx context.Context, relIdArr []int, id int) {
	relDao := RoleRelToAction
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().ActionId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceInt(relIdArrOfOldTmp)

	/**----新增关联操作 开始----**/
	insertRelIdArr := gset.NewIntSetFrom(relIdArr).Diff(gset.NewIntSetFrom(relIdArrOfOld)).Slice()
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
	/**----新增关联操作 开始----**/

	/**----删除关联操作 结束----**/
	deleteRelIdArr := gset.NewIntSetFrom(relIdArrOfOld).Diff(gset.NewIntSetFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联操作 结束----**/
}
