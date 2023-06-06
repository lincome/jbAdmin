package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRole struct{}

func NewRole() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(NewRole())
}

// 总数
func (logicThis *sRole) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *sRole) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoThis.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *sRole) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicThis *sRole) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data["menuIdArr"]
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data["menuIdArr"])
		//menuIdArrCount, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(g.Map{"menuId": data["menuIdArr"], "sceneId": data["sceneId"]}).Count()
		menuIdArrCount, _ := daoAuth.Menu.ParseDbCtx(ctx).Handler(daoAuth.Menu.ParseFilter(g.Map{"menuId": data["menuIdArr"], "sceneId": data["sceneId"]}, &[]string{})).Count()
		if len(menuIdArr) != menuIdArrCount {
			err = utils.NewErrorCode(ctx, 89999998, "")
			return
		}
	}
	_, okActionIdArr := data["actionIdArr"]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data["actionIdArr"])
		//actionIdArrCount, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(g.Map{"actionId": data["actionIdArr"], "sceneId": data["sceneId"]}).Count()
		actionIdArrCount, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Handler(daoAuth.ActionRelToScene.ParseFilter(g.Map{"actionId": data["actionIdArr"], "sceneId": data["sceneId"]}, &[]string{})).Count()
		if len(actionIdArr) != actionIdArrCount {
			err = utils.NewErrorCode(ctx, 89999998, "")
			return
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		return
	}

	if okMenuIdArr {
		daoThis.SaveRelMenu(ctx, gconv.SliceInt(data["menuIdArr"]), int(id))
	}
	if okActionIdArr {
		daoThis.SaveRelAction(ctx, gconv.SliceInt(data["actionIdArr"]), int(id))
	}
	return
}

// 更新
func (logicThis *sRole) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &joinTableArr)).Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicThis *sRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &joinTableArr)).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}
