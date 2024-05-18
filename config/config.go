package config

import (
	"net/url"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

var (
	PORT         = 8080                                // 端口
	AssetPrefix  = "https://oaistatic-cdn.closeai.biz" // 资源前缀
	CacheBuildId = "__VtdGuo2T55cu1fqCkoX"             // 缓存版本号
	BuildId      = "__VtdGuo2T55cu1fqCkoX"             // 线上版本号
	Gclient      = g.Client()                          // http客户端
	Ja3Proxy     *url.URL                              // ja3代理
	ArkoseUrl    = "/v2/"
	OPENAIURL, _ = url.Parse("https://chatgpt.com")

	envScriptTpl = `
	<script src="/list.js"></script>
	<script>
	window.__arkoseUrl="{{.ArkoseUrl}}";
	window.__assetPrefix="{{.AssetPrefix}}";
	</script>
	`
	PowerBy       = `<a href="https://github.com/cockroachai/" target="_blank">Powered By cockroachai</a>`
	ProxyClient   *gclient.Client
	AdminPassword = guid.S()
	SessionCache  = gcache.New()
)

func init() {
	ctx := gctx.GetInitCtx()
	// 读取端口
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	// 读取资源前缀
	assetPrefix := g.Cfg().MustGetWithEnv(ctx, "ASSET_PREFIX").String()
	if assetPrefix != "" {
		AssetPrefix = assetPrefix
	}
	g.Log().Info(ctx, "ASSET_PREFIX:", AssetPrefix)

	// 读取ja3代理
	ja3Proxy := g.Cfg().MustGetWithEnv(ctx, "JA3_PROXY").String()
	if ja3Proxy == "" {
		panic("JA3_PROXY is required")
	}
	u, err := url.Parse(ja3Proxy)
	if err != nil {
		panic(err)
	}
	Ja3Proxy = u
	g.Log().Info(ctx, "JA3_PROXY:", Ja3Proxy.String())

	// 读取powerBy
	powerBy := g.Cfg().MustGetWithEnv(ctx, "POWER_BY").String()
	if powerBy != "" {
		PowerBy = powerBy
	}
	g.Log().Info(ctx, "POWER_BY:", PowerBy)

	// 读取adminPassword
	adminPassword := g.Cfg().MustGetWithEnv(ctx, "ADMIN_PASSWORD").String()
	if adminPassword != "" {
		AdminPassword = adminPassword
	}
	g.Log().Info(ctx, "ADMIN_PASSWORD:", AdminPassword)

	// 检查版本号并同步资源
	cacheBuildId := CheckVersion(ctx, AssetPrefix)
	if cacheBuildId != "" {
		CacheBuildId = cacheBuildId
	}
	g.Log().Info(ctx, "CacheBuildId:", CacheBuildId)

	// 获取线上版本号
	buildId := GetBuildId(ctx)
	if buildId != "" {
		BuildId = buildId
	}
	g.Log().Info(ctx, "BuildId:", BuildId)
	// 每小时更新一次
	go func() {
		for {
			time.Sleep(time.Hour)
			build := GetBuildId(ctx)
			if build != "" {
				BuildId = build
			}
			g.Log().Info(ctx, "BuildId:", BuildId)
			cacheBuildId := CheckVersion(ctx, AssetPrefix)
			if cacheBuildId != "" {
				CacheBuildId = cacheBuildId
			}
			g.Log().Info(ctx, "CacheBuildId:", CacheBuildId)
		}
	}()
	ProxyClient = g.Client().Proxy(Ja3Proxy.String()).SetBrowserMode(true).SetHeaderMap(g.MapStrStr{
		"Origin":     "https://chatgpt.com",
		"Referer":    "https://chatgpt.com/",
		"Host":       "chatgpt.com",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
	})

	// 加载session
	_, err = LoadSession(ctx)
	if err != nil {
		g.Log().Error(ctx, "LoadSession Error: ", err)
	}
	g.Log().Info(ctx, "LoadSession Success")

}

// 检查版本号并同步资源
func CheckVersion(ctx g.Ctx, assetPrefix string) (CacheBuildId string) {
	// 读取 assetPrefix/version
	versionVar := Gclient.GetVar(ctx, assetPrefix+"/version.json")
	CacheBuildId = gjson.New(versionVar).Get("cacheBuildId").String()
	g.Log().Infof(ctx, "Get config From %s ,CacheBuildId: %s", AssetPrefix, CacheBuildId)
	if CacheBuildId == "" {
		return ""
	}
	// 读取buildDate目录索引
	indexUrl := assetPrefix + "/template/" + CacheBuildId + "/index.txt"
	g.Log().Info(ctx, "Get files From ", indexUrl)
	buildDateVar := Gclient.GetVar(ctx, indexUrl).String()
	if buildDateVar == "" {
		return ""
	}
	// 按回车分割
	buildDateList := gstr.Split(buildDateVar, "\n")
	g.Dump(buildDateList)
	// 遍历目录索引 如果没有就下载
	for _, v := range buildDateList {
		if v == "" {
			continue
		}
		// 检查文件是否存在
		if !gfile.Exists("./resource/template/dynamic_templates/" + CacheBuildId + "/" + v) {
			g.Log().Infof(ctx, "Download %s", v)
			// 下载文件
			res, err := Gclient.Get(ctx, assetPrefix+"/template/"+CacheBuildId+"/"+v)
			if err != nil {
				g.Log().Error(ctx, "Download  Error: ", v, err)
				return ""
			}
			defer res.Close()
			if res.StatusCode != 200 {
				g.Log().Error(ctx, "Download  Error: ", v, res.StatusCode)
				return ""
			}
			// 写入文件
			err = gfile.PutBytes("./resource/template/dynamic_templates/"+CacheBuildId+"/"+v, res.ReadAll())
			if err != nil {
				g.Log().Error(ctx, "Download  Error: ", v, err)
				return ""
			}
		}
	}

	return
}

func GetEnvScript(ctx g.Ctx) string {
	script, err := gview.ParseContent(ctx, envScriptTpl, g.Map{
		"ArkoseUrl":   ArkoseUrl,
		"AssetPrefix": AssetPrefix,
	})
	if err != nil {
		g.Log().Error(ctx, "GetEnvScript Error: ", err)
		return ""
	}
	return script
}

// 获取版本号
func GetBuildId(ctx g.Ctx) string {
	resVar := Gclient.GetVar(ctx, "https://tcr9i.xyhelper.cn/ping")
	// gjson.New(resVar).Dump()
	buildId := gjson.New(resVar).Get("buildId").String()
	return buildId

}

// 刷新账号信息
func RefreshSession(ctx g.Ctx, refreshCookie string) (session *gjson.Json, err error) {
	res, err := ProxyClient.SetCookie("__Secure-next-auth.session-token", refreshCookie).SetCookie("oai-dm-tgt-c-240329","2024-04-02").Get(ctx, "https://chatgpt.com/api/auth/session")
	if err != nil {
		g.Log().Error(ctx, "RefreshUserToken Error: ", err)
		return
	}
	defer res.Close()
	if res.StatusCode != 200 {
		err = gerror.Newf("RefreshUserToken Error: %d", res.StatusCode)
		g.Log().Error(ctx, err)
		return
	}
	rrefreshCookie := res.GetCookie("__Secure-next-auth.session-token")
	// g.Log().Info(ctx, "RefreshUserToken", refreshCookie)
	if rrefreshCookie != "" {
		refreshCookie = rrefreshCookie
	}
	session = gjson.New(res.ReadAll())
	session.Set("refreshCookie", refreshCookie)
	// 将session写入 config/session.json
	err = gfile.PutContents("./config/session.json", session.String())
	// 将session写入缓存
	SessionCache.Set(ctx, "session", session, 0)

	return
}

// 加载session
func LoadSession(ctx g.Ctx) (session *gjson.Json, err error) {
	sessionStr := gfile.GetContents("./config/session.json")
	if sessionStr == "" {
		err = gerror.New("session.json is empty")
		return
	}
	session = gjson.New(sessionStr)
	SessionCache.Set(ctx, "session", session, 0)
	return
}

// GetRefreshCookie 获取refreshCookie
func GetRefreshCookie(ctx g.Ctx) (refreshCookie string) {
	sessionVar := SessionCache.MustGet(ctx, "session")
	refreshCookie = gjson.New(sessionVar).Get("refreshCookie").String()
	return
}

// GetAccessToken 获取accessToken
func GetAccessToken(ctx g.Ctx) (accessToken string) {
	sessionVar := SessionCache.MustGet(ctx, "session")
	accessToken = gjson.New(sessionVar).Get("accessToken").String()
	return
}
