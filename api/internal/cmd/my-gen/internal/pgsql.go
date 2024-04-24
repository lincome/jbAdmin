package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type pgsql struct {
	common
}

func (dbHandler pgsql) GetFieldList(ctx context.Context, group, table string) (fieldList []MyGenField) {
	/* fieldListTmp, _ := g.DB(group).GetAll(ctx, `SELECT *,col_description ('tes' :: REGCLASS, ordinal_position) AS column_comment FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`'`)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v[`column_name`].String(),
			FieldTypeRaw: v[`data_type`].String(),
			IsNull:       v[`is_nullable`].Bool(),
			Default:      v[`column_default`].Interface(),
			Comment:      v[`column_comment`].String(),
		}
		if !v[`column_default`].IsNil() {
			defaultStr := v[`column_default`].String()
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
		}
		fieldList[v[`ordinal_position`].Int()-1] = field
	} */
	fieldListTmp, _ := g.DB(group).TableFields(ctx, table)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v.Name,
			FieldTypeRaw: v.Type,
			IsNull:       v.Null,
			Default:      v.Default,
			Comment:      v.Comment,
		}
		if !gvar.New(v.Default).IsNil() {
			defaultStr := gconv.String(v.Default)
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
		}
		fieldList[v.Index] = field
	}
	return
}

func (dbHandler pgsql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	indrelid, _ := g.DB(group).GetValue(ctx, `SELECT oid FROM pg_class WHERE relname = '`+table+`'`)
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_index WHERE indrelid = '`+indrelid.String()+`'`)
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	for _, v := range keyListTmp {
		if !v[`indisvalid`].Bool() {
			continue
		}
		key := MyGenKey{
			IsPrimary: v[`indisprimary`].Bool(),
			IsUnique:  v[`indisunique`].Bool(),
		}
		// g.DB(group).GetValue(ctx, `SELECT indkey FROM pg_index WHERE indexrelid = `+v[`indexrelid`].String())
		fieldIndex := v[`indkey`].Int() //TODO indkey返回值有BUG。联合索引应该返回4 5，但这里却是0。
		if fieldIndex != 0 {
			key.FieldArr = append(key.FieldArr, fieldList[fieldIndex-1].FieldRaw)
			if fieldList[fieldIndex-1].IsAutoInc {
				key.IsAutoInc = true
			}
		}
		keyList = append(keyList, key)
	}
	return
}

func (dbHandler pgsql) GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	fieldLimitStr = fieldInfo[`character_maximum_length`].String()
	return
}

func (dbHandler pgsql) GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat [2]string) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	fieldLimitFloat = [2]string{fieldInfo[`numeric_precision_radix`].String(), fieldInfo[`numeric_scale`].String()}
	return
}

func (dbHandler pgsql) GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt int) {
	fieldLimitInt = 4
	if gstr.Pos(field.FieldTypeRaw, `smallint`) != -1 {
		fieldLimitInt = 2
	} else if gstr.Pos(field.FieldTypeRaw, `bigint`) != -1 {
		fieldLimitInt = 8
	}
	return
}

func (dbHandler pgsql) GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType) {
	if gstr.Pos(field.FieldTypeRaw, `int`) != -1 && gstr.Pos(field.FieldTypeRaw, `point`) == -1 { //int等类型
		fieldType = TypeInt
		/* if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 { //pgsql不分正负
			fieldType = TypeIntU
		} */
	} else if gstr.Pos(field.FieldTypeRaw, `numeric`) != -1 || gstr.Pos(field.FieldTypeRaw, `real`) != -1 || gstr.Pos(field.FieldTypeRaw, `double`) != -1 { //float类型
		fieldType = TypeFloat
		/* if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 { //pgsql不分正负
			fieldType = TypeFloatU
		} */
	} else if field.FieldTypeRaw == `character varying` { //varchar类型
		fieldType = TypeVarchar
	} else if field.FieldTypeRaw == `character` { //char类型
		fieldType = TypeChar
	} else if field.FieldTypeRaw == `text` { //text类型
		fieldType = TypeText
	} else if field.FieldTypeRaw == `json` { //json类型
		fieldType = TypeJson
	} else if gstr.Pos(field.FieldTypeRaw, `timestamp`) != -1 { //datetime类型（在pgsql中，timestamp类型就是datetime类型）
		fieldType = TypeDatetime
	} else if field.FieldTypeRaw == `date` { //date类型
		fieldType = TypeDate
	} else /* if gstr.Pos(field.FieldTypeRaw, `timestamp`) != -1 { //timestamp类型
		fieldType = TypeTimestamp //pgsql没有该类型
	} else  */if gstr.Pos(field.FieldTypeRaw, `time`) != -1 { //time类型
		fieldType = TypeTime
	}
	return
}