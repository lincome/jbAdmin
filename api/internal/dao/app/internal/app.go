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

// AppDao is the data access object for table app.
type AppDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   AppColumns       // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// AppColumns defines and stores column names for table app.
type AppColumns struct {
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	IsStop      string // 停用：0否 1是
	AppId       string // APPID
	AppType     string // 类型：0安卓 1苹果
	PackageName string // 包名
	PackageFile string // 安装包
	VerNo       string // 版本号
	VerName     string // 版本名称
	VerIntro    string // 版本介绍
	ExtraConfig string // 额外配置
	Remark      string // 备注
	IsForcePrev string // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}

// appColumns holds the columns for table app.
var appColumns = AppColumns{
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsStop:      "is_stop",
	AppId:       "app_id",
	AppType:     "app_type",
	PackageName: "package_name",
	PackageFile: "package_file",
	VerNo:       "ver_no",
	VerName:     "ver_name",
	VerIntro:    "ver_intro",
	ExtraConfig: "extra_config",
	Remark:      "remark",
	IsForcePrev: "is_force_prev",
}

// NewAppDao creates and returns a new DAO object for table data access.
func NewAppDao() *AppDao {
	return &AppDao{
		group:   `default`,
		table:   `app`,
		columns: appColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(appColumns)
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
func (dao *AppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *AppDao) Columns() *AppColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *AppDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
