package api

import "github.com/gogf/gf/v2/frame/g"

func init() {
	s := g.Server()
	group := s.Group("/")
	group.GET("/", Index)
	group.GET("/api/auth/session", ProxyApi)
	group.ALL("/backend-anon/*any", ProxyApi)
}

// Init initializes the api module.
func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "Api module initialized")
}
