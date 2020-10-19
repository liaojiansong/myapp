package authController

import (
	"gf-app/app/service/meService"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type AuthController struct {
}

// 登入
func (this *AuthController) Login(r *ghttp.Request) {
	var data *meService.LoginRequest
	// 参数转换
	e := r.Parse(&data)
	if e != nil {
		glog.Error(e)
		response.JsonExit(r, 1, e.Error())
	}
	// 检验密码
	loginResponse, e := meService.Login(data)
	if e != nil {
		response.JsonExit(r, 1, e.Error())
	}
	response.JsonOk(r, loginResponse)
}
