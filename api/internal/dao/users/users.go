// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/users/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalUsersDao is internal type for wrapping internal DAO implements.
type internalUsersDao = *internal.UsersDao

// usersDao is the data access object for table users.
// You can define custom methods on it to extend its functionality as you wish.
type usersDao struct {
	internalUsersDao
}

var (
	// Users is globally public accessible object for table users operations.
	Users = usersDao{
		internal.NewUsersDao(),
	}
)

// 获取daoModel
func (daoThis *usersDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *usersDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *usersDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *usersDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().UserId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *usersDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return `COALESCE(NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Phone + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Email + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Account + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Nickname + `, ''))`
}

// 解析filter
func (daoThis *usersDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().UserId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().UserId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().UserId, v)
				}
			case `label`:
				m = m.Where(m.Builder().WhereLike(daoModel.DbTable+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Email, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Nickname, `%`+gconv.String(v)+`%`))
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case Privacy.Columns().IdCardNo, Privacy.Columns().IdCardGender, Privacy.Columns().IdCardBirthday:
				tablePrivacy := Privacy.ParseDbTable(m.GetCtx())
				m = m.Where(tablePrivacy+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tablePrivacy, daoModel))
			case Privacy.Columns().IdCardName:
				tablePrivacy := Privacy.ParseDbTable(m.GetCtx())
				m = m.WhereLike(tablePrivacy+`.`+k, `%`+gconv.String(v)+`%`)
				m = m.Handler(daoThis.ParseJoin(tablePrivacy, daoModel))
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
func (daoThis *usersDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
			case Privacy.Columns().Password, Privacy.Columns().Salt, Privacy.Columns().IdCardNo, Privacy.Columns().IdCardName, Privacy.Columns().IdCardGender, Privacy.Columns().IdCardBirthday, Privacy.Columns().IdCardAddress:
				tablePrivacy := Privacy.ParseDbTable(m.GetCtx())
				m = m.Fields(tablePrivacy + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tablePrivacy, daoModel))
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
		if daoModel.AfterField.Size() > 0 || len(daoModel.AfterFieldWithParam) > 0 {
			m = m.Hook(daoThis.HookSelect(daoModel))
		}
		return m
	}
}

// hook select
func (daoThis *usersDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *usersDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
			case daoThis.Columns().Phone:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().Email:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().Account:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().WxOpenid:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().WxUnionid:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case Privacy.Columns().Password, Privacy.Columns().Salt, Privacy.Columns().IdCardNo, Privacy.Columns().IdCardName, Privacy.Columns().IdCardGender, Privacy.Columns().IdCardBirthday, Privacy.Columns().IdCardAddress:
				if garray.NewStrArrayFrom([]string{``, `0`, `[]`, `{}`}).Contains(gconv.String(v)) { //gvar.New(v).IsEmpty()无法验证指针的值是空的数据
					continue
				}
				insertData, ok := daoModel.AfterInsert[`privacy`].(map[string]any)
				if !ok {
					insertData = map[string]any{}
				}
				insertData[k] = v
				daoModel.AfterInsert[`privacy`] = insertData
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
func (daoThis *usersDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `privacy`:
					insertData, _ := v.(map[string]any)
					insertData[Privacy.Columns().UserId] = id
					Privacy.CtxDaoModel(ctx).HookInsert(insertData).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *usersDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			case daoThis.Columns().Phone:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().Email:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().Account:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().WxOpenid:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().WxUnionid:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case Privacy.Columns().Password, Privacy.Columns().Salt, Privacy.Columns().IdCardNo, Privacy.Columns().IdCardName, Privacy.Columns().IdCardGender, Privacy.Columns().IdCardBirthday, Privacy.Columns().IdCardAddress:
				updateData, ok := daoModel.AfterUpdate[`privacy`].(map[string]any)
				if !ok {
					updateData = map[string]any{}
				}
				updateData[k] = v
				daoModel.AfterUpdate[`privacy`] = updateData
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[k] = v
				}
			}
		}
		m = m.Data(updateData)
		if len(daoModel.AfterUpdate) > 0 {
			m = m.Hook(daoThis.HookUpdate(daoModel))
			if len(updateData) == 0 {
				daoModel.IsOnlyAfterUpdate = true
			}
		}
		return m
	}
}

// hook update
func (daoThis *usersDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if daoModel.IsOnlyAfterUpdate {
				result = driver.RowsAffected(0)
			} else {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `privacy`:
					updateData, _ := v.(map[string]any)
					for _, id := range daoModel.IdArr {
						updateData[Privacy.Columns().UserId] = id
						if row, _ := Privacy.CtxDaoModel(ctx).Filter(Privacy.Columns().UserId, id).HookUpdate(updateData).UpdateAndGetAffected(); row == 0 { //更新失败，有可能记录不存在，这时做插入操作
							Privacy.CtxDaoModel(ctx).HookInsert(updateData).Insert()
						}
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
						daoModel.CloneNew().Filter(`id`, id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *usersDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			Privacy.CtxDaoModel(ctx).Filter(Privacy.Columns().UserId, daoModel.IdArr).Delete()
			return
		},
	}
}

// 解析group
func (daoThis *usersDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().UserId)
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
func (daoThis *usersDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().UserId, 1))
			case daoThis.Columns().Birthday:
				m = m.Order(daoModel.DbTable + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.Columns().UserId)
			case Privacy.Columns().IdCardBirthday:
				tablePrivacy := Privacy.ParseDbTable(m.GetCtx())
				m = m.Order(tablePrivacy + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.Columns().UserId)
				m = m.Handler(daoThis.ParseJoin(tablePrivacy, daoModel))
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
func (daoThis *usersDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case Privacy.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Privacy.Columns().UserId+` = `+daoModel.DbTable+`.`+daoThis.Columns().UserId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().UserId+` = `+daoModel.DbTable+`.`+daoThis.Columns().UserId)
		}
		return m
	}
}

// Fill with you ideas below.
