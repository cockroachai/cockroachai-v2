package api

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Login(r *ghttp.Request) {
	// ctx := r.GetCtx()
	if r.Session.MustGet("userToken").String() != "" {
		r.Response.RedirectTo("/")
		return
	}
	// if config.LOGINCALLBACK != "" && config.LOGINCALLBACK != "/login" {
	// 	r.Response.RedirectTo(config.LOGINCALLBACK)
	// 	return
	// }
	// g.Log().Info(ctx, "login", config.PowerBy)

	r.Response.WriteTpl("login.html", g.Map{
		"powerBy": config.PowerBy,
	})
}
