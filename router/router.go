package router

import (
	"gf-app/app/api/hello"
	"gf-app/app/api/me"
	"gf-app/app/model/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.Middleware(middleware.Inject)
		group.ALL("/", hello.Hello)
		group.Group("/me", func(group *ghttp.RouterGroup) {
			me := &me.Me{}
			group.POST("/login", me.Login)
			group.GET("/index", me.Index)
			group.GET("/update", me.Update)
		})
	})

}
