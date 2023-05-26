// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SceneDao is the data access object for table auth_scene.
type SceneDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns SceneColumns // columns contains all the column names of Table for convenient usage.
}

// SceneColumns defines and stores column names for table auth_scene.
type SceneColumns struct {
	SceneId     string // 权限场景ID
	SceneCode   string // 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
	SceneName   string // 名称
	SceneConfig string // 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
	IsStop      string // 是否停用：0否 1是
	UpdateTime  string // 更新时间
	CreateTime  string // 创建时间
}

// sceneColumns holds the columns for table auth_scene.
var sceneColumns = SceneColumns{
	SceneId:     "sceneId",
	SceneCode:   "sceneCode",
	SceneName:   "sceneName",
	SceneConfig: "sceneConfig",
	IsStop:      "isStop",
	UpdateTime:  "updateTime",
	CreateTime:  "createTime",
}

// NewSceneDao creates and returns a new DAO object for table data access.
func NewSceneDao() *SceneDao {
	return &SceneDao{
		group:   "default",
		table:   "auth_scene",
		columns: sceneColumns,
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
func (dao *SceneDao) Columns() SceneColumns {
	return dao.columns
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

// PrimaryKey returns the primary key name of current dao.
func (dao *SceneDao) PrimaryKey() string {
	return reflect.ValueOf(dao.columns).Field(0).String()
}

// Column returns all column names of current dao.
func (dao *SceneDao) Column() []string {
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	column := make([]string, count)
	for i := 0; i < count; i++ {
		column[i] = v.Field(i).String()
	}
	return column
}
