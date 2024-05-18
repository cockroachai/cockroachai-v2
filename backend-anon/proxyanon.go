package backendanon

import (
	"cockroachai/config"
	"cockroachai/utils"
	"crypto/tls"
	"net/http"
	"net/http/httputil"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	proxy = &httputil.ReverseProxy{}
)

func ProxyAnno(r *ghttp.Request) {
	ctx := r.Context()
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(config.Ja3Proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	proxy.Rewrite = func(proxyRequest *httputil.ProxyRequest) {
		proxyRequest.SetURL(config.OPENAIURL)
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		// 移除cookie
		response.Header.Del("Set-Cookie")
		if response.StatusCode >= 400 {
			g.Log().Warning(ctx, "ProxyAnno:", response.StatusCode, response.Request.Method, response.Request.URL.String())
		} else {
			g.Log().Info(ctx, "ProxyAnno:", response.StatusCode, response.Request.Method, response.Request.URL.String())
		}
		return nil
	}
	header := r.Request.Header
	header.Set("Origin", "https://chatgpt.com")
	header.Set("Referer", "https://chatgpt.com/")
	utils.HeaderModify(&r.Request.Header)

	proxy.ServeHTTP(r.Response.RawWriter(), r.Request)

}
