package api

import (
	"gf-app/app/entity"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Di struct {
	Me *entity.MeCache
}

/**
直接从缓存中拉取
*/
func (this *Di) Init(r *ghttp.Request) {
	token := r.GetHeader("token")
	meCache, err := entity.LoadMe(token)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	this.Me = meCache
}
