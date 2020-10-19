package router

import (
	"gf-app/app/api"
	"gf-app/app/api/authController"
	"gf-app/app/api/meController"
	"gf-app/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		auth := &authController.AuthController{}
		group.ALL("/auth", auth)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		me := &meController.Me{&api.Di{}}
		group.ALL("/me", me)
		//group.Group("/me", func(group *ghttp.RouterGroup) {
		//	me := &meController.Me{&api.Di{}}
		//	//group.POST("/login", me.Login)
		//	//group.GET("/index", me.Index)
		//	//group.GET("/update", me.Update)
		//	//group.ALL()
		//})
	})

}
