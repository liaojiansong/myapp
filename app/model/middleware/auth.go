package middleware

import (
	"gf-app/app/service/me"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var noCheck = map[string]bool{
	"/me/login":  true,
	"/me/logout": true,
}
var noInject = map[string]bool{
	"/me/login": true,
}

func needCheck(uri string) bool {
	_, ok := noCheck[uri]
	return ok
}

func needInject(uri string) bool {
	_, ok := noInject[uri]
	return !ok
}

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	if needCheck(r.Router.Uri) {
		token := r.GetHeader("token")
		isLogin, e := me.CheckLogin(token)
		if e != nil {
			response.JsonExit(r, 1, e.Error())
		}
		if isLogin == false {
			response.JsonExit(r, 1, "请重新登入")
		}
	}
	r.Middleware.Next()
}
