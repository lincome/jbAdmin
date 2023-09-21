package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type AdminListReq struct {
	g.Meta `path:"/admin/list" method:"post" tags:"平台后台/权限管理/管理员" sm:"列表"`
	Filter AdminListFilter `json:"filter" dc:"过滤条件"`
	Field  []string        `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string          `json:"sort" default:"id DESC" dc:"排序"`
	Page   int             `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int             `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type AdminListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	AdminId        *uint       `json:"adminId,omitempty" v:"integer|min:1" dc:"管理员ID"`
	Account        string      `json:"account,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Phone          string      `json:"phone,omitempty" v:"phone" dc:"手机号"`
	RoleId         *uint       `json:"roleId,omitempty" v:"integer|min:1" dc:"角色ID"`
	IsStop         *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type AdminListRes struct {
	Count int             `json:"count" dc:"总数"`
	List  []AdminListItem `json:"list" dc:"列表"`
}

type AdminListItem struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	AdminId   *uint       `json:"adminId,omitempty" dc:"管理员ID"`
	Account   *string     `json:"account,omitempty" dc:"账号"`
	Phone     *string     `json:"phone,omitempty" dc:"手机号"`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	Nickname  *string     `json:"nickname,omitempty" dc:"昵称"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type AdminInfoReq struct {
	g.Meta `path:"/admin/info" method:"post" tags:"平台后台/权限管理/管理员" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type AdminInfoRes struct {
	Info AdminInfo `json:"info" dc:"详情"`
}

type AdminInfo struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	AdminId   *uint       `json:"adminId,omitempty" dc:"管理员ID"`
	Account   *string     `json:"account,omitempty" dc:"账号"`
	Phone     *string     `json:"phone,omitempty" dc:"手机号"`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	Nickname  *string     `json:"nickname,omitempty" dc:"昵称"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	RoleIdArr []uint      `json:"roleIdArr,omitempty" dc:"角色ID列表"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type AdminCreateReq struct {
	g.Meta    `path:"/admin/create" method:"post" tags:"平台后台/权限管理/管理员" sm:"创建"`
	Account   *string `json:"account,omitempty" v:"required-without:Phone|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Phone     *string `json:"phone,omitempty" v:"required-without:Account|phone" dc:"手机号"`
	Password  *string `json:"password,omitempty" v:"required|size:32|regex:^[\\p{L}\\p{N}_-]+$" dc:"密码"`
	RoleIdArr *[]uint `json:"roleIdArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"角色ID列表"`
	Nickname  *string `json:"nickname,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"昵称"`
	Avatar    *string `json:"avatar,omitempty" v:"url|length:1,120" dc:"头像"`
	IsStop    *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type AdminUpdateReq struct {
	g.Meta    `path:"/admin/update" method:"post" tags:"平台后台/权限管理/管理员" sm:"更新"`
	IdArr     []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	Account   *string `json:"account,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Phone     *string `json:"phone,omitempty" v:"phone" dc:"手机号"`
	Password  *string `json:"password,omitempty" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$" dc:"密码"`
	RoleIdArr *[]uint `json:"roleIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"角色ID列表"`
	Nickname  *string `json:"nickname,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"昵称"`
	Avatar    *string `json:"avatar,omitempty" v:"url|length:1,120" dc:"头像"`
	IsStop    *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type AdminDeleteReq struct {
	g.Meta `path:"/admin/del" method:"post" tags:"平台后台/权限管理/管理员" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
