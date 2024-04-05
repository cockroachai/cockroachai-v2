package api

import (
	"cockroachai/config"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	// proxy    = &httputil.ReverseProxy{}
	UpStream = "https://chat.openai.com"
	u, _     = url.Parse(UpStream)
)

func ProxyApi(r *ghttp.Request) {
	ctx := r.Context()
	proxy := &httputil.ReverseProxy{}
	path := r.RequestURI
	g.Log().Info(ctx, "ProxyApi:", path)
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(config.Ja3Proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	proxy.Rewrite = func(proxyRequest *httputil.ProxyRequest) {
		proxyRequest.SetURL(u)
	}

	proxy.ServeHTTP(r.Response.RawWriter(), r.Request)

}
