package api

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// setup GET /setup
func Setup(r *ghttp.Request) {
	r.Response.WriteTpl("setup.html")
}

func SetupPost(r *ghttp.Request) {
	ctx := r.Context()
	result := ""
	refreshCookie := r.Get("refreshCookie").String()
	// 去掉首尾空格
	refreshCookie = gstr.Trim(refreshCookie)
	if refreshCookie == "" {
		result = "refreshCookie is required"
		r.Response.WriteTpl("setup.html", g.Map{
			"result": result,
		})
		return
	}
	adminPassword := r.Get("adminPassword").String()
	// 去掉首尾空格
	adminPassword = gstr.Trim(adminPassword)
	if adminPassword != config.AdminPassword {
		result = "adminPassword is invalid"
		r.Response.WriteTpl("setup.html", g.Map{
			"result": result,
		})
		return
	}
	// g.Log().Info(ctx, "SetupPost", refreshCookie)
	session, err := config.RefreshSession(ctx, refreshCookie)
	if err != nil {
		g.Log().Error(ctx, "RefreshUserToken", err)
		result = err.Error()
		r.Response.WriteTpl("setup.html", g.Map{
			"result": result,
		})
		return

	}

	r.Response.WriteTpl("setup.html", g.Map{
		"result": session,
	})
}
