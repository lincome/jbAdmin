package my_gen

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenDaoField struct {
	ImportDao []string

	InsertParseBefore myGenDataHandler
	InsertParse       myGenDataHandler
	InsertHook        myGenDataHandler

	UpdateParse      myGenDataHandler
	UpdateHookBefore myGenDataHandler
	UpdateHookAfter  myGenDataHandler

	FieldParse myGenDataHandler
	FieldHook  myGenDataHandler

	FilterParse myGenDataHandler

	OrderParse myGenDataHandler

	JoinParse myGenDataHandler
}

func (myGenDaoFieldThis *myGenDaoField) reset() {
	myGenDaoFieldThis = &myGenDaoField{}
}

func getDaoFieldList(tpl myGenTpl) (daoFieldList []myGenDaoField) {
	labelListLen := len(tpl.Handle.LabelList)
	if labelListLen > 0 {
		fieldParseStr := `case ` + "`label`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)`
		filterParseStr := `case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if labelListLen > 1 {
			fieldParseStrTmp := "` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + " + `"
			parseFilterStr := "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + ", `%`+gconv.String(v)+`%`)"
			for i := labelListLen - 2; i >= 0; i-- {
				fieldParseStrTmp = "IF(IFNULL(` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, '') != '', ` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, " + fieldParseStrTmp + ")"
				if i == 0 {
					parseFilterStr = "WhereLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				} else {
					parseFilterStr = "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				}
			}
			fieldParseStr = `case ` + "`label`" + `:
				m = m.Fields(` + "`" + fieldParseStrTmp + " AS ` + v)"
			filterParseStr = `case ` + "`label`" + `:
				m = m.Where(m.Builder().` + parseFilterStr + `)`
		}

		daoField := myGenDaoField{}
		daoField.FieldParse.Method = ReturnTypeName
		daoField.FieldParse.DataTypeName = append(daoField.FieldParse.DataTypeName, fieldParseStr)
		daoField.FilterParse.Method = ReturnTypeName
		daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, filterParseStr)

		daoFieldList = append(daoFieldList, daoField)
	}

	for _, v := range tpl.FieldList {
		daoField := myGenDaoField{}
		// 根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容）
		/*--------根据字段数据类型处理 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型`	// `char类型`
			if v.IndexRaw == `UNI` && v.IsNull {
				daoField.InsertParse.Method = ReturnType
				daoField.InsertParse.DataType = append(daoField.InsertParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				insertData[k] = v
				if gconv.String(v) == `+"``"+` {
					insertData[k] = nil
				}`)

				daoField.UpdateParse.Method = ReturnType
				daoField.UpdateParse.DataType = append(daoField.UpdateParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = v
				if gconv.String(v) == `+"``"+` {
					updateData[daoModel.DbTable+`+"`.`"+`+k] = nil
				}`)
			}
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			if v.IsNull {
				daoField.InsertParse.Method = ReturnType
				daoField.InsertParse.DataType = append(daoField.InsertParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				insertData[k] = v
				if gconv.String(v) == `+"``"+` {
					insertData[k] = nil
				}`)

				daoField.UpdateParse.Method = ReturnType
				daoField.UpdateParse.DataType = append(daoField.UpdateParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = gvar.New(v)
				if gconv.String(v) == `+"``"+` {
					updateData[daoModel.DbTable+`+"`.`"+`+k] = nil
				}`)
			}
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
			daoField.OrderParse.Method = ReturnType
			daoField.OrderParse.DataType = append(daoField.OrderParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		}
		/*--------根据字段数据类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
			daoField.FilterParse.Method = ReturnTypeName
			daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, `case `+"`timeRangeStart`"+`:
				m = m.WhereGTE(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`, v)
			case `+"`timeRangeEnd`"+`:
				m = m.WhereLTE(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`, v)`)
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
		case TypeNamePid: // pid；	类型：int等类型；
			daoField.FieldParse.Method = ReturnTypeName
			if len(tpl.Handle.LabelList) > 0 {
				daoField.FieldParse.DataTypeName = append(daoField.FieldParse.DataTypeName, `case `+"`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`"+`:
					tableP := `+"`p_`"+` + daoModel.DbTable
					m = m.Fields(tableP + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` + `+"` AS `"+` + v)
					m = m.Handler(daoThis.ParseJoin(tableP, daoModel))`)
			}
			daoField.FieldParse.DataTypeName = append(daoField.FieldParse.DataTypeName, `case `+"`tree`"+`:
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+`)
				m = m.Handler(daoThis.ParseOrder([]string{`+"`tree`"+`}, daoModel))`)

			orderParseStr := `case ` + "`tree`" + `:
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + v.FieldCaseCamel + `)`
			if tpl.Handle.Pid.Sort != `` {
				orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Sort) + `)`
			}
			orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
			daoField.OrderParse.Method = ReturnTypeName
			daoField.OrderParse.DataTypeName = append(daoField.OrderParse.DataTypeName, orderParseStr)

			daoField.JoinParse.Method = ReturnTypeName
			daoField.JoinParse.DataTypeName = append(daoField.JoinParse.DataTypeName, `case `+"`p_`"+` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+`+"` AS `"+`+joinTable, joinTable+`+"`.`"+`+daoThis.PrimaryKey()+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`)`)

			if tpl.Handle.Pid.IsCoexist {
				daoField.InsertParseBefore.Method = ReturnTypeName
				daoField.InsertParseBefore.DataTypeName = append(daoField.InsertParseBefore.DataTypeName, `if _, ok := insert[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; !ok {
			insert[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`] = 0
		}`)

				daoField.InsertParse.Method = ReturnTypeName
				daoField.InsertParse.DataTypeName = append(daoField.InsertParse.DataTypeName, `case daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					daoModel.AfterInsert[`+"`pIdPath`"+`] = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					daoModel.AfterInsert[`+"`pLevel`"+`] = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				} else {
					daoModel.AfterInsert[`+"`pIdPath`"+`] = `+"`0`"+`
					daoModel.AfterInsert[`+"`pLevel`"+`] = 0
				}`)

				daoField.InsertHook.Method = ReturnTypeName
				daoField.InsertHook.DataTypeName = append(daoField.InsertHook.DataTypeName, `updateSelfData := map[string]interface{}{}
			for k, v := range daoModel.AfterInsert {
				switch k {
				case `+"`pIdPath`"+`:
					updateSelfData[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gconv.String(v) + `+"`-`"+` + gconv.String(id)
				case `+"`pLevel`"+`:
					updateSelfData[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}`)

				daoField.UpdateParse.Method = ReturnTypeName
				daoField.UpdateParse.DataTypeName = append(daoField.UpdateParse.DataTypeName, `case daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = v
				pIdPath := `+"`0`"+`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					pLevel = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				}
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`CONCAT('`"+` + pIdPath + `+"`-', `"+` + daoThis.PrimaryKey() + `+"`)`"+`)
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							`+"`pIdPathOfOld`"+`: oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`],
							`+"`pIdPathOfNew`"+`: pIdPath + `+"`-`"+` + oldInfo[daoThis.PrimaryKey()].String(),
							`+"`pLevelOfOld`"+`:  oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`],
							`+"`pLevelOfNew`"+`:  pLevel + 1,
						})
					}
				}
				if len(updateChildIdPathAndLevelList) > 0 {
					daoModel.AfterUpdate[`+"`updateChildIdPathAndLevelList`"+`] = updateChildIdPathAndLevelList
				}
			case `+"`childIdPath`"+`: //更新所有子孙级的idPath。参数：map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[`+"`pIdPathOfOld`"+`])
				pIdPathOfNew := gconv.String(val[`+"`pIdPathOfNew`"+`])
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`REPLACE(`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+` + `+"`, '`"+` + pIdPathOfOld + `+"`', '`"+` + pIdPathOfNew + `+"`')`"+`)
			case `+"`childLevel`"+`: //更新所有子孙级的level。参数：map[string]interface{}{`+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`+"`pLevelOfOld`"+`])
				pLevelOfNew := gconv.Uint(val[`+"`pLevelOfNew`"+`])
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` + `"+` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` - `"+` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`)

				daoField.UpdateHookAfter.Method = ReturnTypeName
				daoField.UpdateHookAfter.DataTypeName = append(daoField.UpdateHookAfter.DataTypeName, `for k, v := range daoModel.AfterUpdate {
				switch k {
				case `+"`updateChildIdPathAndLevelList`"+`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`, `+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoModel.CloneNew().Filter(`+"`pIdPathOfOld`"+`, v1[`+"`pIdPathOfOld`"+`]).HookUpdate(g.Map{
							`+"`childIdPath`"+`: g.Map{
								`+"`pIdPathOfOld`"+`: v1[`+"`pIdPathOfOld`"+`],
								`+"`pIdPathOfNew`"+`: v1[`+"`pIdPathOfNew`"+`],
							},
							`+"`childLevel`"+`: g.Map{
								`+"`pLevelOfOld`"+`: v1[`+"`pLevelOfOld`"+`],
								`+"`pLevelOfNew`"+`: v1[`+"`pLevelOfNew`"+`],
							},
						}).Update()
					}
				}
			}`)

				daoField.FilterParse.Method = ReturnTypeName
				daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, `case `+"`pIdPathOfOld`"+`: //父级IdPath（旧）
				m = m.WhereLike(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`, gconv.String(v)+`+"`-%`"+`)`)
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			daoField.OrderParse.Method = ReturnTypeName
			daoField.OrderParse.DataTypeName = append(daoField.OrderParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			daoField = myGenDaoField{}
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			insertParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			updateParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			passwordMapKey := tpl.getHandlePasswordMapKey(v.FieldRaw)
			if tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				insertParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
				updateParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
			}
			insertParseStr += `
				insertData[k] = password`
			updateParseStr += `
				updateData[daoModel.DbTable+` + "`.`" + `+k] = password`

			daoField.InsertParse.Method = ReturnTypeName
			daoField.InsertParse.DataTypeName = append(daoField.InsertParse.DataTypeName, insertParseStr)

			daoField.UpdateParse.Method = ReturnTypeName
			daoField.UpdateParse.DataTypeName = append(daoField.UpdateParse.DataTypeName, updateParseStr)
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			daoField.FilterParse.Method = ReturnTypeName
			daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.WhereLike(daoModel.DbTable+`+"`.`"+`+k, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)`)
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				daoPath := relIdObj.tpl.TableCaseCamel
				if relIdObj.tpl.RemovePrefixAlone != tpl.RemovePrefixAlone {
					daoPath = `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
					daoField.ImportDao = append(daoField.ImportDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
				}

				if !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
					fieldParseStr := `case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + `:` + `
				table` + relIdObj.tpl.TableCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx())
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + `, daoModel))`
					if relIdObj.Suffix != `` {
						fieldParseStr = `case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`:" + `
				table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` := ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + `, daoModel))`
					}
					daoField.FieldParse.Method = ReturnTypeName
					daoField.FieldParse.DataTypeName = append(daoField.FieldParse.DataTypeName, fieldParseStr)
				}

				joinParseStr := `case ` + daoPath + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				if relIdObj.Suffix != `` {
					joinParseStr = `case ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `:
			m = m.LeftJoin(` + daoPath + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				}
				daoField.JoinParse.Method = ReturnTypeName
				daoField.JoinParse.DataTypeName = append(daoField.JoinParse.DataTypeName, joinParseStr)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			daoField.OrderParse.Method = ReturnTypeName
			daoField.OrderParse.DataTypeName = append(daoField.OrderParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			daoField.FilterParse.Method = ReturnTypeName
			daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, filterParseStr)
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			daoField.FilterParse.Method = ReturnTypeName
			daoField.FilterParse.DataTypeName = append(daoField.FilterParse.DataTypeName, filterParseStr)
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		daoFieldList = append(daoFieldList, daoField)
	}
	return
}
