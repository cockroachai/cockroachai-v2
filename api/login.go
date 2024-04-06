package api

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/container/garray"
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

func LoginPost(r *ghttp.Request) {
	ctx := r.GetCtx()
	userToken := r.Get("password").String()
	if userToken == "" {
		r.Response.WriteJsonExit(g.Map{
			"error":   "userToken is required",
			"powerBy": config.PowerBy,
		})
		return
	}
	userTokens := garray.NewStrArrayFrom(g.Cfg().MustGet(ctx, "USERTOKENS").Strings())
	if !userTokens.Contains(userToken) {
		r.Response.WriteTpl("login.html", g.Map{
			"error":   "userToken is invalid",
			"powerBy": config.PowerBy,
		})
		return
	}
	r.Session.Set("userToken", userToken)
	r.Response.RedirectTo("/")
}
