// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package org

import (
	"api/internal/cache"
	"api/internal/consts"
	daoIndex "api/internal/dao"
	"api/internal/dao/org/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalConfigDao is internal type for wrapping internal DAO implements.
type internalConfigDao = *internal.ConfigDao

// configDao is the data access object for table org_config.
// You can define custom methods on it to extend its functionality as you wish.
type configDao struct {
	internalConfigDao
}

var (
	// Config is globally public accessible object for table org_config operations.
	Config = configDao{
		internal.NewConfigDao(),
	}
)

// 获取daoModel
func (daoThis *configDao) CtxDaoModel(ctx context.Context, dbOpt ...any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *configDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *configDao) ParseDbTable(ctx context.Context, dbTableOpt ...any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *configDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return `CONCAT_WS('|', COALESCE(` + daoModel.DbTable + `.` + daoThis.Columns().OrgId + `, ''), COALESCE(` + daoModel.DbTable + `.` + daoThis.Columns().ConfigKey + `, ''))`
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *configDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().ConfigKey
}

// 解析filter
func (daoThis *configDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.Strings(v)
				}
				inStrArr := make([]string, len(idArr))
				for index, id := range idArr {
					inStrArr[index] = `('` + gstr.Replace(id, `|`, `', '`) + `')`
				}
				m = m.Where(`(` + daoModel.DbTable + `.` + daoThis.Columns().OrgId + `, ` + daoModel.DbTable + `.` + daoThis.Columns().ConfigKey + `) IN (` + gstr.Join(inStrArr, `, `) + `)`)
			case `exc_id`, `exc_id_arr`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.Strings(v)
				}
				inStrArr := make([]string, len(idArr))
				for index, id := range idArr {
					inStrArr[index] = `('` + gstr.Replace(id, `|`, `', '`) + `')`
				}
				m = m.Where(`(` + daoModel.DbTable + `.` + daoThis.Columns().OrgId + `, ` + daoModel.DbTable + `.` + daoThis.Columns().ConfigKey + `) NOT IN (` + gstr.Join(inStrArr, `, `) + `)`)
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().ConfigKey, `%`+gconv.String(v)+`%`)
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
func (daoThis *configDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
			case Org.Columns().OrgName:
				tableOrg := Org.ParseDbTable(m.GetCtx())
				m = m.Fields(tableOrg + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableOrg, daoModel))
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

// 处理afterField
func (daoThis *configDao) HandleAfterField(ctx context.Context, record gdb.Record, daoModel *daoIndex.DaoModel) {
	for _, v := range daoModel.AfterFieldSlice {
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

// hook select
func (daoThis *configDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil || len(result) == 0 {
				return
			}

			var wg sync.WaitGroup
			wg.Add(len(result))
			daoModel.AfterFieldSlice = daoModel.AfterField.Slice()
			for _, record := range result {
				go func(record gdb.Record) {
					defer wg.Done()
					daoThis.HandleAfterField(ctx, record, daoModel)
				}(record)
			}
			wg.Wait()
			return
		},
	}
}

// 解析insert
func (daoThis *configDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
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
func (daoThis *configDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
					daoModel.CloneNew().FilterPri(id).HookUpdate(g.Map{k: v}).Update()
				}
			} */
			return
		},
	}
}

// 解析update
func (daoThis *configDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[k] = v
				}
			}
		}
		m = m.Data(updateData)
		if len(daoModel.AfterUpdate) == 0 {
			return m
		}
		m = m.Hook(daoThis.HookUpdate(daoModel))
		if len(updateData) == 0 {
			daoModel.IsOnlyAfterUpdate = true
		}
		return m
	}
}

// hook update
func (daoThis *configDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case `xxxx`:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().FilterPri(id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *configDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *configDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().OrgId)
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().ConfigKey)
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
func (daoThis *configDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				suffix := gstr.TrimLeftStr(kArr[0], k, 1)
				m = m.Order(daoModel.DbTable + `.` + daoThis.Columns().OrgId + suffix)
				m = m.Order(daoModel.DbTable + `.` + daoThis.Columns().ConfigKey + suffix)
				remain := gstr.TrimLeftStr(gstr.TrimLeftStr(v, k+suffix, 1), `,`, 1)
				if remain != `` {
					m = m.Order(remain)
				}
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
func (daoThis *configDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case Org.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Org.Columns().OrgId+` = `+daoModel.DbTable+`.`+daoThis.Columns().OrgId)
		}
		return m
	}
}

// Fill with you ideas below.

// 获取单个配置
func (daoThis *configDao) GetOne(ctx context.Context, orgId string, configKey string) (configValue *gvar.Var) {
	// configValue, _ = daoThis.CtxDaoModel(ctx).FilterPri(orgId+`|`+configKey).Value(daoThis.Columns().ConfigValue)
	configValue, _ = cache.DbData.GetOrSet(ctx, daoThis.CtxDaoModel(ctx), orgId+`|`+configKey, consts.CACHE_TIME_DEFAULT, daoThis.Columns().ConfigValue)
	return
}

// 获取配置
func (daoThis *configDao) Get(ctx context.Context, orgId string, configKeyArr ...string) (config gdb.Record, err error) {
	idArr := make([]any, len(configKeyArr))
	for index, configKey := range configKeyArr {
		idArr[index] = orgId + `|` + configKey
	}
	// configTmp, err := daoThis.CtxDaoModel(ctx).FilterPri(idArr).PluckStr(daoThis.Columns().ConfigValue, daoThis.Columns().ConfigKey)
	configTmp, err := cache.DbData.GetOrSetPluck(ctx, daoThis.CtxDaoModel(ctx), idArr, consts.CACHE_TIME_DEFAULT, daoThis.Columns().ConfigValue)
	if err != nil {
		return
	}
	config = gdb.Record{}
	for k, v := range configTmp {
		config[gstr.Replace(k, orgId+`|`, ``, 1)] = v
	}
	return
}

// 保存配置
func (daoThis *configDao) Save(ctx context.Context, orgId string, config map[string]any) (err error) {
	idArr := make([]any, 0, len(config))
	daoModelThis := daoThis.CtxDaoModel(ctx)
	err = daoModelThis.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		for k, v := range config {
			_, err = tx.Model(daoModelThis.DbTable).Data(g.Map{
				daoThis.Columns().OrgId:       orgId,
				daoThis.Columns().ConfigKey:   k,
				daoThis.Columns().ConfigValue: v,
			}).OnConflict(daoThis.Columns().OrgId, daoThis.Columns().ConfigKey).Save()
			if err != nil {
				return
			}
			idArr = append(idArr, orgId+`|`+k)
		}
		return
	})
	if err != nil {
		return
	}
	cache.DbData.Del(ctx, daoModelThis, idArr...)
	return
}
