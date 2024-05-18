package api

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

func Index(r *ghttp.Request) {
	ctx := r.GetCtx()

	// if r.Session.MustGet("userToken").IsEmpty() {
	// 	r.Response.RedirectTo("/login")
	// 	// r.Response.Writer.Write([]byte("Hello XyHelper"))
	// 	return
	// }
	model := r.Get("model").String()

	propsJson := gjson.New(Props)
	propsJson.Set("query.model", model)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)
	r.Cookie.Set("oai-did", uuid.New().String())

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/chat.html", g.Map{
		"props":       propsJson,
		"arkoseUrl":   config.ArkoseUrl,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}
func C(r *ghttp.Request) {
	ctx := r.GetCtx()

	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	convId := r.GetRouter("convId").String()

	g.Log().Debug(r.GetCtx(), "convId", convId)

	propsJson := gjson.New(Props)
	propsJson.Set("query.default.1", convId)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/chat.html", g.Map{
		"props":       propsJson,
		"arkoseUrl":   config.ArkoseUrl,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// Discovery 发现
func Discovery(r *ghttp.Request) {
	ctx := r.GetCtx()
	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	props := `
  {
    "props": {
      "pageProps": {
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/gpts/discovery",
    "query": {},
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/discovery.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

func Gpts(r *ghttp.Request) {
	ctx := r.GetCtx()
	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	props := `
  {
    "props": {
      "pageProps": {
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/gpts",
    "query": {},
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/gpts.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// Editor 编辑器
func Editor(r *ghttp.Request) {
	ctx := r.GetCtx()

	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}

	props := `
  {
    "props": {
      "pageProps": {
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/gpts/editor",
    "query": {},
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	// if slug != "" {
	// 	propsJson.Set("page", "/gpts/editor/[slug]")
	// 	propsJson.Set("query.slug", slug)
	// }
	// propsJson.Dump()

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/editor.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// Slug 编辑器
func Slug(r *ghttp.Request) {
	ctx := r.GetCtx()

	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	slug := r.GetRouter("slug").String()

	props := `
  {
    "props": {
      "pageProps": {
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/gpts/editor/[slug]",
    "query": { "slug": "g-I2KQmH4yZ" },
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)

	propsJson.Set("query.slug", slug)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/slug.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// G 游戏
func G(r *ghttp.Request) {
	ctx := r.GetCtx()
	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	gizmoId := r.GetRouter("gizmoId").String()
	props := `
  {
    "props": {
      "pageProps": {
        "kind": "chat_page",
        "gizmo": null,
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/g/[gizmoId]",
    "query": { "gizmoId": "g-I2KQmH4yZ-unix2bj" },
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)
	propsJson.Set("query.gizmoId", gizmoId)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/g.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// GC 游戏会话
func GC(r *ghttp.Request) {
	ctx := r.GetCtx()

	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	gizmoId := r.GetRouter("gizmoId").String()
	convId := r.GetRouter("convId").String()
	g.Log().Debug(r.GetCtx(), "gizmoId", gizmoId)
	props := `
  {
    "props": {
      "pageProps": {
        "user": {
          "id": "user-xyhelper",
          "name": "admin@closeai.com",
          "email": "admin@closeai.com",
          "image": "/avatars.png",
          "picture": "/avatars.png",
          "idp": "auth0",
          "iat": 2699699364,
          "mfa": false,
          "groups": []
        },
        "serviceStatus": {},
        "userCountry": "US",
        "serviceAnnouncement": { "public": {}, "paid": {} },
        "serverPrimedAllowBrowserStorageValue": true,
        "canManageBrowserStorage": false,
        "ageVerificationDeadline": null,
        "showCookieConsentBanner": false
      },
      "__N_SSP": true
    },
    "page": "/g/[gizmoId]/c/[convId]",
    "query": {
      "gizmoId": "g-I2KQmH4yZ-unix2bj",
      "convId": "e5fa1ee7-f482-4892-86ad-12cf0e0f9dd7"
    },
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
  }
  `
	propsJson := gjson.New(props)
	propsJson.Set("query.gizmoId", gizmoId)
	propsJson.Set("query.convId", convId)
	propsJson.Set("buildId", config.BuildId)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/gc.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})
}

// Mine 我的
func Mine(r *ghttp.Request) {
	ctx := r.GetCtx()
	if r.Session.MustGet("userToken").IsEmpty() {
		r.Response.RedirectTo("/login")
		return
	}
	props := `
  {
    "props": {
        "pageProps": {
            "user": {
                "id": "user-xyhelper",
                "name": "admin@closeai.com",
                "email": "admin@closeai.com",
                "image": "/avatars.png",
                "picture": "/avatars.png",
                "idp": "auth0",
                "iat": 2699699364,
                "mfa": false,
                "groups": [],
                "intercom_hash": "30fd0a0ada1c07ce526be7c3d54c22904b80fa7e2713d978630e979e4315cf67"
            },
            "serviceStatus": {},
            "userCountry": "US",
            "serviceAnnouncement": {
                "paid": {},
                "public": {}
            },
            "serverPrimedAllowBrowserStorageValue": true,
            "canManageBrowserStorage": false,
            "ageVerificationDeadline": null,
            "showCookieConsentBanner": false
        },
        "__N_SSP": true
    },
    "page": "/gpts/mine",
    "query": {},
    "buildId": "wtXFegAXt6bfbujLr1e7S",
    "assetPrefix": "",
    "isFallback": false,
    "gssp": true,
    "scriptLoader": []
}`
	propsJson := gjson.New(props)
	propsJson.Set("buildId", config.BuildId)
	propsJson.Set("assetPrefix", config.AssetPrefix)

	r.Response.WriteTpl("dynamic_templates/"+config.CacheBuildId+"/mine.html", g.Map{
		"arkoseUrl":   config.ArkoseUrl,
		"props":       propsJson,
		"assetPrefix": config.AssetPrefix,
		"envScript":   config.GetEnvScript(ctx),
	})

}
