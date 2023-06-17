// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Http is the golang structure of table log_http for DAO operations like Where/Data.
type Http struct {
	g.Meta    `orm:"table:log_http, do:true"`
	HttpId    interface{} // http记录ID
	Url       interface{} // 地址
	Header    interface{} // 请求头
	ReqData   interface{} // 请求数据
	ResData   interface{} // 响应数据
	RunTime   interface{} // 运行时间（单位：毫秒）
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
