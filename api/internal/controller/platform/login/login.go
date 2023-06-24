package controller

import (
	"api/api"
	apiLogin "api/api/platform/login"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取加密盐
func (controllerThis *Login) EncryptStr(ctx context.Context, req *apiLogin.LoginEncryptStrReq) (res *api.CommonEncryptStrRes, err error) {
	encryptStr, err := service.Login().EncryptStr(ctx, `platform`, req.Account)
	if err != nil {
		return
	}
	res = &api.CommonEncryptStrRes{EncryptStr: encryptStr}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiLogin.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	token, err := service.Login().Login(ctx, `platform`, req.Account, req.Password)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}

// 用户详情
func (controllerThis *Login) Info(ctx context.Context, req *apiLogin.LoginInfoReq) (res *apiLogin.LoginInfoRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	res = &apiLogin.LoginInfoRes{}
	loginInfo.Struct(&res.Info)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), map[string]interface{}{`info`: loginInfo}, 0)
	return
}

// 修改个人信息
func (controllerThis *Login) Update(ctx context.Context, req *apiLogin.LoginUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter := map[string]interface{}{`id`: loginInfo[`adminId`]}
	/**--------参数处理 结束--------**/

	_, err = service.Admin().Update(ctx, filter, data)
	return
}

// 用户菜单树
func (controllerThis *Login) MenuTree(ctx context.Context, req *apiLogin.LoginMenuTreeReq) (res *apiLogin.LoginMenuTreeRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	filter := map[string]interface{}{}
	filter[`selfMenu`] = map[string]interface{}{
		`sceneCode`: sceneInfo[`sceneCode`].String(),
		`sceneId`:   sceneInfo[`sceneId`].Int(),
		`loginId`:   loginInfo[`adminId`].Int(),
	}
	field := []string{`menuTree`, `showMenu`}

	list, err := service.Menu().List(ctx, filter, field, []string{}, 0, 0)
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, `menuId`, `pid`)
	res = &apiLogin.LoginMenuTreeRes{}
	tree.Structs(&res.Tree)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), map[string]interface{}{`tree`: tree}, 0)
	return
}