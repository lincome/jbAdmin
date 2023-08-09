package initialize

// 多语言设置
import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.New()

	g.I18n().SetPath(g.Cfg().MustGet(ctx, `i18n.path`).String())         //设置资源目录
	g.I18n().SetLanguage(g.Cfg().MustGet(ctx, `i18n.language`).String()) //设置默认为中文（原默认为英文en）
}
