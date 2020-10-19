package middleware

import (
	"gf-app/app/entity"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 鉴权中间件
func Auth(r *ghttp.Request) {
	// 是否传递token
	token := r.GetHeader("token")
	if token == "" {
		response.JsonExit(r, -1, "token is missing")
	}
	exists, err := entity.ExistsMe(token)
	if err != nil {
		glog.Error(err)
		response.JsonExit(r, -1, "server error")
	}
	if exists == false {
		response.JsonExit(r, -1, "need login")
	}
	r.Middleware.Next()
}
