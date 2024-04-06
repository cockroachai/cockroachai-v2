package next

import "github.com/gogf/gf/v2/frame/g"

func init() {
	s := g.Server()
	group := s.Group("/")
	group.GET("/_next/*any", ProxyNext)
}

func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "Next module initialized")
}
