package main

import (
	"cockroachai/arkose"
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	// 加载arkose模块
	arkose.Init(ctx)

	// 启动HTTP服务
	s := g.Server()
	s.SetPort(config.PORT)
	s.Run()
}
