package api

import "github.com/gogf/gf/v2/net/ghttp"

// setup GET /setup
func Setup(r *ghttp.Request) {
	r.Response.WriteTpl("setup.html")
}
