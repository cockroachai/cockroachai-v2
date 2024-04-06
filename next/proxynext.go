package next

import (
	"cockroachai/config"
	"cockroachai/utils"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	proxy    = &httputil.ReverseProxy{}
	UpStream = "https://chat.openai.com"
	u, _     = url.Parse(UpStream)
)

func ProxyNext(r *ghttp.Request) {
	ctx := r.Context()
	path := r.Request.URL.Path
	g.Log().Info(ctx, "ProxyNext:", path)
	// 替换path中的版本号
	path = strings.Replace(path, config.CacheBuildId, config.BuildId, 1)
	r.Request.URL.Path = path
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(config.Ja3Proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	proxy.Rewrite = func(proxyRequest *httputil.ProxyRequest) {
		proxyRequest.SetURL(u)
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		// 移除cookie
		response.Header.Del("Set-Cookie")
		return nil
	}
	header := r.Request.Header
	header.Set("Origin", "https://chat.openai.com")
	header.Set("Referer", "https://chat.openai.com/")
	utils.HeaderModify(&r.Request.Header)

	proxy.ServeHTTP(r.Response.RawWriter(), r.Request)

}
