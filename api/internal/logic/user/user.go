package logic

import (
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(NewUser())
}

// 新增
func (logicThis *sUser) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoUser.User
	id, err = daoThis.HandlerCtx(ctx).HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sUser) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoUser.User
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoHandlerThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sUser) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoUser.User
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoHandlerThis.HookSelect().DeleteAndGetAffected()
	return
}
