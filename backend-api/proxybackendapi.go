package backendapi

import (
	"bytes"
	"cockroachai/config"
	"cockroachai/utils"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ProxyBackendApi(r *ghttp.Request) {
	ctx := r.Context()
	userToken := r.Session.MustGet("userToken").String()
	if userToken == "" {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{"detail": "Unauthorized"})
	}

	// 屏蔽一些接口
	path := r.RequestURI
	if strings.Contains(path, "invite") ||
		strings.HasPrefix(path, "/backend-api/share/create") ||
		strings.HasPrefix(path, "/backend-api/shared_conversations") ||
		(strings.HasPrefix(path, "/backend-api/accounts") && r.Request.Method == "DELETE") ||
		strings.HasPrefix(path, "/backend-api/aip/p/") ||
		strings.HasPrefix(path, "/backend-api/gizmo_creator_profile") ||
		strings.HasPrefix(path, "/backend-api/payments/checkout") ||
		strings.HasPrefix(path, "/backend-api/payments/customer_portal") ||
		strings.HasPrefix(path, "/backend-api/user_system_messages") ||
		strings.HasPrefix(path, "/backend-api/accounts/deactivate") {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{"detail": "你无权进行此操作。"})
		return
	}

	// g.Log().Info(ctx, "ProxyBackendApi:", path)
	proxy := &httputil.ReverseProxy{}

	proxy.Transport = &http.Transport{
		Proxy: http.ProxyURL(config.Ja3Proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ForceAttemptHTTP2: true}

	proxy.Rewrite = func(proxyRequest *httputil.ProxyRequest) {
		proxyRequest.SetURL(config.OPENAIURL)
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		// 移除cookie
		response.Header.Del("Set-Cookie")
		if response.StatusCode >= 400 {
			g.Log().Warning(ctx, "ProxyBackendApi:", response.StatusCode, response.Request.Method, response.Request.URL.String())
		} else {
			g.Log().Info(ctx, "ProxyBackendApi:", response.StatusCode, response.Request.Method, response.Request.URL.String())
		}
		if path == "/backend-api/me" && response.StatusCode == 200 {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
			resJson := gjson.New(body)
			resJson.Set("email", "share@closeai.com")
			resJson.Set("name", "share")
			resJson.Set("picture", "/avatars.png")
			resJson.Set("phone_number", "+1911")
			resJson.Set("id", "user-"+userToken)
			resJson.Set("orgs.data.0.name", "user-"+userToken)
			resJson.Set("orgs.data.0.description", "OpenAI")
			body = resJson.MustToJson()
			response.Body = io.NopCloser(bytes.NewReader(body))
			response.ContentLength = int64(len(body))
			response.Header.Set("Content-Length", fmt.Sprint(len(body)))
			response.Header.Del("Content-Encoding")

		}

		return nil
	}
	header := r.Request.Header
	header.Set("Origin", "https://chat.openai.com")
	header.Set("Referer", "https://chat.openai.com/")
	// header.Del("Cookie")
	header.Del("Accept-Encoding")
	accessToken := config.GetAccessToken(ctx)
	if accessToken != "" {
		header.Set("Authorization", "Bearer "+accessToken)
	}
	utils.HeaderModify(&r.Request.Header)

	proxy.ServeHTTP(r.Response.RawWriter(), r.Request)

}
