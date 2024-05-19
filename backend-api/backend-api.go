package backendapi

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	group := s.Group("/")
	group.ALL("/backend-api/*any", ProxyBackendApi)
	group.GET("/backend-api/prompt_library/", PromptLibrary)
	group.POST("/backend-api/lat/r", LatR)
}

func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "BackendApi module initialized")
}

// /backend-api/prompt_library/
func PromptLibrary(r *ghttp.Request) {
	ctx := r.Context()
	userToken := r.Session.MustGet("userToken").String()
	if userToken == "" {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{"detail": "Unauthorized"})
	}
	limit := r.Get("limit").String()
	offset := r.Get("offset").String()
	headerMap := g.MapStrStr{
		"Origin":       "https://chatgpt.com",
		"Referer":      "https://chatgpt.com/",
		"Host":         "chatgpt.com",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		"OAI-Language": r.Header.Get("OAI-Language"),
	}
	accessToken := config.GetAccessToken(ctx)
	if accessToken != "" {
		headerMap["Authorization"] = "Bearer " + accessToken
	}
	ProxyClient := gclient.New().Proxy(config.Ja3Proxy.String()).SetBrowserMode(true).SetHeaderMap(headerMap)
	res, err := ProxyClient.Get(ctx, "https://chatgpt.com/backend-api/prompt_library/", g.Map{"limit": limit, "offset": offset})
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteStatus(500, "")
		return
	}
	defer res.Close()
	// res.RawDump()
	r.Response.Status = res.StatusCode
	r.Response.Write(res.ReadAll())
}

// /backend-api/lat/r
func LatR(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"status": "success",
	})
}
