// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

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
func (daoThis *roleDao) ParseDbGroup(dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	if len(dbGroupSeldata) > 0 { //分库逻辑
	}
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(dbTableSelData map[string]interface{}) string {
	table := daoThis.Table()
	if len(dbTableSelData) > 0 { //分表逻辑
	}
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *roleDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.Table()).Safe().Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.ParseDbTable(dbSelDataList[1])).Safe().Ctx(ctx)
	default:
		return daoThis.Ctx(ctx)
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoThis.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
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
func (daoThis *roleDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoThis.Table()+"."+daoThis.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+"."+k] = v
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
func (daoThis *roleDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoThis.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoThis.Table() + "." + daoThis.PrimaryKey() + " AS " + v)
			case "sceneName":
				m = m.Fields(Scene.Table() + "." + v)
				m = daoThis.ParseJoin("scene", joinTableArr)(m)
			case "menuIdArr", "actionIdArr":
				//需要id字段
				m = m.Fields(daoThis.Table() + "." + daoThis.PrimaryKey())

				afterField = append(afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				val := gconv.SliceInt(v)
				if len(val) == 1 {
					m = m.Where(daoThis.Table()+"."+daoThis.PrimaryKey(), val[0])
				} else {
					m = m.Where(daoThis.Table()+"."+daoThis.PrimaryKey(), val)
				}
			case "excId", "excIdArr":
				val := gconv.SliceInt(v)
				if len(val) == 1 {
					m = m.WhereNot(daoThis.Table()+"."+daoThis.PrimaryKey(), val[0])
				} else {
					m = m.WhereNotIn(daoThis.Table()+"."+daoThis.PrimaryKey(), val)
				}
			case "startTime":
				m = m.WhereGTE(daoThis.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoThis.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoThis.PrimaryKey(), "Id", "Name")
				m = m.WhereLike(daoThis.Table()+"."+keywordField, gconv.String(v))
			default:
				kArr := strings.Split(k, " ") //支持"id > ?"等k
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					if len(kArr) == 1 {
						if gstr.ToLower(gstr.SubStr(kArr[0], -2)) == "id" {
							val := gconv.SliceInt(v)
							if len(val) == 1 {
								m = m.Where(daoThis.Table()+"."+k, val[0])
							} else {
								m = m.Where(daoThis.Table()+"."+k, val)
							}
						} else if gstr.ToLower(gstr.SubStr(kArr[0], -4)) == "name" {
							m = m.WhereLike(daoThis.Table()+"."+k, gconv.String(v))
						} else {
							m = m.Where(daoThis.Table()+"."+k, v)
						}
					} else {
						m = m.Where(daoThis.Table()+"."+k, v)
					}
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
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoThis.Table() + "." + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoThis.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *roleDao) ParseOrder(order [][2]string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoThis.Table()+"."+daoThis.PrimaryKey(), v[1])
			default:
				if daoThis.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoThis.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
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
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoThis.PrimaryKey()+" = "+daoThis.Table()+"."+daoThis.PrimaryKey())
		} */
		case "scene":
			sceneTable := Scene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(sceneTable) {
				*joinTableArr = append(*joinTableArr, sceneTable)
				m = m.LeftJoin(sceneTable, sceneTable+"."+Scene.PrimaryKey()+" = "+daoThis.Table()+"."+Scene.PrimaryKey())
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoThis *roleDao) AfterField(afterField []string) gdb.HookHandler {
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
					case "menuIdArr":
						menuIdArr, _ := RoleRelToMenu.ParseDbCtx(ctx).Where("roleId", record[daoThis.PrimaryKey()]).Array("menuId")
						record[v] = gvar.New(menuIdArr)
					case "actionIdArr":
						actionIdArr, _ := RoleRelToAction.ParseDbCtx(ctx).Where("roleId", record[daoThis.PrimaryKey()]).Array("actionId")
						record[v] = gvar.New(actionIdArr)
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// Fill with you ideas below.
// 保存关联菜单
func (daoThis *roleDao) SaveRelMenu(ctx context.Context, menuIdArr []int, id int) {
	menuIdArrOfOldTmp, _ := RoleRelToMenu.ParseDbCtx(ctx).Where("roleId", id).Array("menuId")
	menuIdArrOfOld := gconv.SliceInt(menuIdArrOfOldTmp)
	/**----新增关联菜单 开始----**/
	insertMenuIdArr := gset.NewIntSetFrom(menuIdArr).Diff(gset.NewIntSetFrom(menuIdArrOfOld)).Slice()
	if len(insertMenuIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertMenuIdArr {
			insertList = append(insertList, map[string]interface{}{
				"roleId": id,
				"menuId": v,
			})
		}
		RoleRelToMenu.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联菜单 结束----**/

	/**----删除关联菜单 开始----**/
	deleteMenuIdArr := gset.NewIntSetFrom(menuIdArrOfOld).Diff(gset.NewIntSetFrom(menuIdArr)).Slice()
	if len(deleteMenuIdArr) > 0 {
		RoleRelToMenu.ParseDbCtx(ctx).Where(g.Map{"roleId": id, "menuId": deleteMenuIdArr}).Delete()
	}
	/**----删除关联菜单 结束----**/
}

// 保存关联操作
func (daoThis *roleDao) SaveRelAction(ctx context.Context, actionIdArr []int, id int) {
	actionIdArrOfOldTmp, _ := RoleRelToAction.ParseDbCtx(ctx).Where("roleId", id).Array("actionId")
	actionIdArrOfOld := gconv.SliceInt(actionIdArrOfOldTmp)

	/**----新增关联操作 开始----**/

	inserttActionIdArr := gset.NewIntSetFrom(actionIdArr).Diff(gset.NewIntSetFrom(actionIdArrOfOld)).Slice()
	if len(inserttActionIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range inserttActionIdArr {
			insertList = append(insertList, map[string]interface{}{
				"roleId":   id,
				"actionId": v,
			})
		}
		RoleRelToAction.ParseDbCtx(ctx).Data(insertList).Insert()
	}

	/**----删除关联操作 开始----**/
	deleteActionIdArr := gset.NewIntSetFrom(actionIdArrOfOld).Diff(gset.NewIntSetFrom(actionIdArr)).Slice()
	if len(deleteActionIdArr) > 0 {
		RoleRelToAction.ParseDbCtx(ctx).Where(g.Map{"roleId": id, "actionId": deleteActionIdArr}).Delete()
	}
	/**----删除关联操作 结束----**/
}
