package api

import "github.com/gogf/gf/v2/frame/g"

func init() {
	s := g.Server()
	group := s.Group("/")
	group.GET("/", Index)
	group.GET("/api/auth/session", ProxyApi)
	group.ALL("/backend-anon/*any", ProxyApi)
	group.ALL("/public-api/*any", ProxyApi)
	group.GET("/api/auth/providers", AuthProviders)
	group.GET("/api/auth/csrf", AuthCsrf)
	group.POST("/api/auth/signin/login-web", AuthSigninLoginWeb)
	group.POST("/api/auth/signin/auth0", AuthSigninAuth0)

}

// Init initializes the api module.
func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "Api module initialized")
}
