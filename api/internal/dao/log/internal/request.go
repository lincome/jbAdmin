// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RequestDao is the data access object for table log_request.
type RequestDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns RequestColumns // columns contains all the column names of Table for convenient usage.
}

// RequestColumns defines and stores column names for table log_request.
type RequestColumns struct {
	LogId         string // 请求日志ID
	RequestUrl    string // 请求地址
	RequestHeader string // 请求头
	RequestData   string // 请求数据
	ResponseBody  string // 响应体
	RunTime       string // 运行时间（单位：毫秒）
	UpdateTime    string // 更新时间
	CreateTime    string // 创建时间
}

// requestColumns holds the columns for table log_request.
var requestColumns = RequestColumns{
	LogId:         "logId",
	RequestUrl:    "requestUrl",
	RequestHeader: "requestHeader",
	RequestData:   "requestData",
	ResponseBody:  "responseBody",
	RunTime:       "runTime",
	UpdateTime:    "updateTime",
	CreateTime:    "createTime",
}

// NewRequestDao creates and returns a new DAO object for table data access.
func NewRequestDao() *RequestDao {
	return &RequestDao{
		group:   "default",
		table:   "log_request",
		columns: requestColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RequestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RequestDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RequestDao) Columns() RequestColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RequestDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RequestDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RequestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *RequestDao) PrimaryKey() string {
	return reflect.ValueOf(dao.columns).Field(0).String()
}

// 所有字段的数组
func (dao *RequestDao) ColumnArr() []string {
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	column := make([]string, count)
	for i := 0; i < count; i++ {
		column[i] = v.Field(i).String()
	}
	return column
}

// 所有字段的数组（返回的格式更方便使用）
func (dao *RequestDao) ColumnArrG() *garray.StrArray {
	return garray.NewStrArrayFrom(dao.ColumnArr())
}
