package middleware

import (
	"gf-app/app/service/me"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

/**

 */
func Inject(r *ghttp.Request) {
	if needInject(r.Router.Uri) == true {
		token := r.GetHeader(me.KEY)
		userCache, e := me.Load(token)
		if e != nil {
			response.JsonExit(r, 1, e.Error())
		}
		r.SetCtxVar("user", userCache)
	}
	r.Middleware.Next()
}
