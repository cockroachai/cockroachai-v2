package backendanon

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

var ()

func init() {

	s := g.Server()
	group := s.Group("/")
	group.ALL("/backend-anon/*any", ProxyAnno)
	group.GET("/backend-anon/prompt_library", PromptLibrary)

	// group.GET("/backend-anon/prompt_library", ProxyAnno)
}

func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "Backend Anon module initialized")
}

// /backend-anon/prompt_library/
func PromptLibrary(r *ghttp.Request) {
	ctx := r.Context()
	limit := r.Get("limit").String()
	offset := r.Get("offset").String()
	ProxyClient := gclient.New().Proxy(config.Ja3Proxy.String()).SetBrowserMode(true).SetHeaderMap(g.MapStrStr{
		"Origin":        "https://chatgpt.com",
		"Referer":       "https://chatgpt.com/",
		"Host":          "chatgpt.com",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		"OAI-Language":  r.Header.Get("OAI-Language"),
		"OAI-Device-Id": r.Header.Get("OAI-Device-Id"),
	})
	res, err := ProxyClient.Get(ctx, "https://chatgpt.com/backend-anon/prompt_library/", g.Map{"limit": limit, "offset": offset})
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
