package config

import (
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	PORT         = 8080                                // 端口
	AssetPrefix  = "https://oaistatic-cdn.closeai.biz" // 资源前缀
	CacheBuildId = "__VtdGuo2T55cu1fqCkoX"             // 缓存版本号
	BuildId      = "__VtdGuo2T55cu1fqCkoX"             // 线上版本号
	gclient      = g.Client()                          // http客户端
	Ja3Proxy     *url.URL                              // ja3代理
	ArkoseUrl    = "/v2/"

	envScriptTpl = `
	<script src="/list.js"></script>
	<script>
	window.__arkoseUrl="{{.ArkoseUrl}}";
	window.__assetPrefix="{{.AssetPrefix}}";
	</script>
	`
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

}

// 检查版本号并同步资源
func CheckVersion(ctx g.Ctx, assetPrefix string) (CacheBuildId string) {
	// 读取 assetPrefix/version
	versionVar := gclient.GetVar(ctx, assetPrefix+"/version.json")
	CacheBuildId = gjson.New(versionVar).Get("cacheBuildId").String()
	g.Log().Infof(ctx, "Get config From %s ,CacheBuildId: %s", AssetPrefix, CacheBuildId)
	if CacheBuildId == "" {
		return ""
	}
	// 读取buildDate目录索引
	indexUrl := assetPrefix + "/template/" + CacheBuildId + "/index.txt"
	g.Log().Info(ctx, "Get files From ", indexUrl)
	buildDateVar := gclient.GetVar(ctx, indexUrl).String()
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
			res, err := gclient.Get(ctx, assetPrefix+"/template/"+CacheBuildId+"/"+v)
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
	resVar := gclient.GetVar(ctx, "https://tcr9i.xyhelper.cn/ping")
	// gjson.New(resVar).Dump()
	buildId := gjson.New(resVar).Get("buildId").String()
	return buildId

}
