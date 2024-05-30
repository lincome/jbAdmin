package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthRole struct{}

func NewAuthRole() *sAuthRole {
	return &sAuthRole{}
}

func init() {
	service.RegisterAuthRole(NewAuthRole())
}

// 新增
func (logicThis *sAuthRole) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.CtxDaoModel(ctx)

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.SliceUint(data[`menu_id_arr`])) > 0 {
		menuIdArr := gconv.SliceUint(data[`menu_id_arr`])
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Menu.Columns().MenuId: menuIdArr, daoAuth.Menu.Columns().SceneId: data[`scene_id`]}).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	}
	if _, ok := data[`action_id_arr`]; ok && len(gconv.SliceUint(data[`action_id_arr`])) > 0 {
		actionIdArr := gconv.SliceUint(data[`action_id_arr`])
		if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: data[`scene_id`]}).Count(); count != len(actionIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.SliceUint(data[`menu_id_arr`])) > 0 {
		menuIdArr := gconv.SliceUint(data[`menu_id_arr`])
		filterTmp := g.Map{daoAuth.Menu.Columns().MenuId: menuIdArr}
		if _, ok := data[`scene_id`]; ok {
			filterTmp[daoAuth.Menu.Columns().SceneId] = data[`scene_id`]
		} else {
			sceneIdArr, _ := daoModelThis.CloneNew().Filter(`id`, daoModelThis.IdArr).ArrayUint(daoThis.Columns().SceneId)
			if garray.NewArrayFrom(gconv.SliceAny(sceneIdArr)).Unique().Len() != 1 {
				err = utils.NewErrorCode(ctx, 89999998, ``) //因菜单所属场景ID只能一个，故只能允许相同场景ID下的角色一起修改菜单
				return
			}
			filterTmp[daoAuth.Menu.Columns().SceneId] = sceneIdArr[0]
		}
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(filterTmp).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	}

	if _, ok := data[`action_id_arr`]; ok && len(gconv.SliceUint(data[`action_id_arr`])) > 0 {
		actionIdArr := gconv.SliceUint(data[`action_id_arr`])
		filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr}
		if _, ok := data[`scene_id`]; ok {
			filterTmp[daoAuth.ActionRelToScene.Columns().SceneId] = data[`scene_id`]
			if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(filterTmp).Count(); count != len(actionIdArr) {
				err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
				return
			}
		} else {
			for _, id := range daoModelThis.IdArr {
				oldInfo, _ := daoModelThis.CloneNew().Filter(`id`, id).One()
				filterTmp[daoAuth.ActionRelToScene.Columns().SceneId] = oldInfo[daoThis.Columns().SceneId]
				if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(filterTmp).Count(); count != len(actionIdArr) {
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthRole) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
