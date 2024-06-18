package controller

import (
	"api/api"
	apiOrg "api/api/org/org"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Admin struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewAdmin() *Admin {
	field := daoOrg.Admin.ColumnArr().Slice()
	field = gset.NewStrSetFrom(field).Diff(gset.NewStrSetFrom([]string{daoOrg.Admin.Columns().Password, daoOrg.Admin.Columns().Salt})).Slice() //移除敏感字段
	defaultFieldOfList := []string{`id`, `label`, daoOrg.Org.Columns().OrgName}
	defaultFieldOfInfo := []string{`id`, `label`, `role_id_arr`}
	return &Admin{
		defaultFieldOfList: append(field, defaultFieldOfList...),
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, gset.NewStrSetFrom(defaultFieldOfList).Merge(gset.NewStrSetFrom(defaultFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Admin) List(ctx context.Context, req *apiOrg.AdminListReq) (res *apiOrg.AdminListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]any{}
	}

	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfList
	}

	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter[daoOrg.Admin.Columns().OrgId] = loginInfo[daoOrg.Admin.Columns().OrgId].Int()
	filter[daoOrg.Admin.Columns().IsSuper] = 0 //机构超级管理员不显示，也不能修改和删除
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgAdminRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoOrg.Admin.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiOrg.AdminListRes{Count: count, List: []apiOrg.AdminInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Admin) Info(ctx context.Context, req *apiOrg.AdminInfoReq) (res *apiOrg.AdminInfoRes, err error) {
	/**--------参数处理 开始--------**/
	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfInfo
	}
	filter := map[string]any{`id`: req.Id}

	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter[daoOrg.Admin.Columns().OrgId] = loginInfo[daoOrg.Admin.Columns().OrgId].Int()
	filter[daoOrg.Admin.Columns().IsSuper] = 0 //机构超级管理员不显示，也不能修改和删除
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgAdminRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoOrg.Admin.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiOrg.AdminInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Admin) Create(ctx context.Context, req *apiOrg.AdminCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})

	loginInfo := utils.GetCtxLoginInfo(ctx)
	data[daoOrg.Admin.Columns().OrgId] = loginInfo[daoOrg.Admin.Columns().OrgId].Int()
	data[daoOrg.Admin.Columns().IsSuper] = 0 //只能是普通管理员
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgAdminCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.OrgAdmin().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Admin) Update(ctx context.Context, req *apiOrg.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, `id_arr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]any{`id`: req.IdArr}

	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter[daoOrg.Admin.Columns().OrgId] = loginInfo[daoOrg.Admin.Columns().OrgId].Int()
	filter[daoOrg.Admin.Columns().IsSuper] = 0 //机构超级管理员不显示，也不能修改和删除
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgAdminUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.OrgAdmin().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Admin) Delete(ctx context.Context, req *apiOrg.AdminDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]any{`id`: req.IdArr}

	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter[daoOrg.Admin.Columns().OrgId] = loginInfo[daoOrg.Admin.Columns().OrgId].Int()
	filter[daoOrg.Admin.Columns().IsSuper] = 0 //机构超级管理员不显示，也不能修改和删除
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgAdminDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.OrgAdmin().Delete(ctx, filter)
	return
}
