// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SceneDao is the data access object for table auth_scene.
type SceneDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    SceneColumns     // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  *garray.StrArray // 所有字段的数组
}

// SceneColumns defines and stores column names for table auth_scene.
type SceneColumns struct {
	SceneId     string // 场景ID
	SceneName   string // 名称
	SceneCode   string // 标识
	SceneConfig string // 配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}
	Remark      string // 备注
	IsStop      string // 停用：0否 1是
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
}

// sceneColumns holds the columns for table auth_scene.
var sceneColumns = SceneColumns{
	SceneId:     "sceneId",
	SceneName:   "sceneName",
	SceneCode:   "sceneCode",
	SceneConfig: "sceneConfig",
	Remark:      "remark",
	IsStop:      "isStop",
	UpdatedAt:   "updatedAt",
	CreatedAt:   "createdAt",
}

// NewSceneDao creates and returns a new DAO object for table data access.
func NewSceneDao() *SceneDao {
	return &SceneDao{
		group:   `default`,
		table:   `auth_scene`,
		columns: sceneColumns,
		primaryKey: func() string {
			return reflect.ValueOf(sceneColumns).Field(0).String()
		}(),
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(sceneColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return garray.NewStrArrayFrom(column)
		}(),
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SceneDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SceneDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *SceneDao) Columns() *SceneColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SceneDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SceneDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SceneDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *SceneDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *SceneDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
