package api

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// /api/auth/providers
func AuthProviders(r *ghttp.Request) {
	jsonStr := `
	{
		"auth0": {
		  "id": "auth0",
		  "name": "Auth0",
		  "type": "oauth",
		  "signinUrl": "https://chat.openai.com/api/auth/signin/auth0",
		  "callbackUrl": "https://chat.openai.com/api/auth/callback/auth0"
		},
		"login-web": {
		  "id": "login-web",
		  "name": "Auth0",
		  "type": "oauth",
		  "signinUrl": "https://chat.openai.com/api/auth/signin/login-web",
		  "callbackUrl": "https://chat.openai.com/api/auth/callback/login-web"
		},
		"openai": {
		  "id": "openai",
		  "name": "openai",
		  "type": "oauth",
		  "signinUrl": "https://chat.openai.com/api/auth/signin/openai",
		  "callbackUrl": "https://chat.openai.com/api/auth/callback/openai"
		}
	  }`
	r.Response.WriteJson(gjson.New(jsonStr))
}

// /api/auth/csrf
func AuthCsrf(r *ghttp.Request) {
	jsonStr := `
	{
		"csrfToken": "084e72020630350580584426d7088865eeb0e7c0d8af0ee38931049bae4baa96"
	  }
	`
	r.Response.WriteJson(gjson.New(jsonStr))
}

// /api/auth/signin/login-web
func AuthSigninLoginWeb(r *ghttp.Request) {
	jsonStr := `
	{
		"url": "/login"
	  }
	`
	r.Response.WriteJson(gjson.New(jsonStr))
}

// /api/auth/signin/auth0
func AuthSigninAuth0(r *ghttp.Request) {
	jsonStr := `
	{
		"url": "/login"
	  }
	`
	r.Response.WriteJson(gjson.New(jsonStr))
}

// /api/auth/session
func AuthSession(r *ghttp.Request) {
	ctx := r.Context()
	userToken := r.Session.MustGet("userToken").String()
	if userToken == "" {
		r.Response.WriteJsonExit(g.Map{})
	} else {
		// 获取session缓存
		sessionVar := config.SessionCache.MustGet(ctx, "session")
		r.Response.WriteJsonExit(sessionVar)
	}
}

// /auth/logout
func AuthLogout(r *ghttp.Request) {
	r.Session.RemoveAll()
	r.Response.RedirectTo("/")
}
