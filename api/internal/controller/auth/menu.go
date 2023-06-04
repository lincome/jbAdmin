package controller

import (
	apiAuth "api/api/auth"
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

func (c *Menu) List(r *ghttp.Request) {
	var param *apiAuth.MenuListReq
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	filter := gconv.Map(param.Filter) //条件过滤
	order := [2]string{"id", "DESC"}
	if param.Sort.Key != "" {
		order[0] = param.Sort.Key
	}
	if param.Sort.Order != "" {
		order[1] = param.Sort.Order
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 10
	}

	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		//isAuth := $this->checkAuth(__FUNCTION__, $sceneCode, false);
		isAuth := true
		/**--------参数处理 开始--------**/
		allowField := []string{"menuId", "menuName", "id"}
		if isAuth {
			allowField = daoAuth.Menu.ColumnArr()
			allowField = append(allowField, "id")
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------参数处理 结束--------**/
		count, err := service.Menu().Count(r.Context(), filter)
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		list, err := service.Menu().List(r.Context(), filter, field, order, int((param.Page-1)*param.Limit), int(param.Limit))
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"count": count, "list": list}, 0, "")
		/* r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
		r.Response.WriteJson(map[string]interface{}{
			"code": 0,
			"msg":  g.I18n().Tf(r.GetCtx(), "0"),
			"data": map[string]interface{}{
				"list": list,
			},
		}) */
	}
}