package me

import (
	meModel "gf-app/app/model/me"
	meService "gf-app/app/service/me"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 数据验证放在控制器

type Me struct {
}

type LoginRequest struct {
	Name     string `v:"required#用户昵称必须"`
	Password string `v:"required#密码必须"`
}

type LoginResponse struct {
	*meModel.Entity
	Token string `json:"token"`
}

func (this *Me) Login(r *ghttp.Request) {
	var data *LoginRequest
	// 参数转换
	e := r.Parse(&data)
	if e != nil {
		glog.Error(e)
		response.JsonExit(r, 1, e.Error())
	}
	// 检验密码
	token, entity, e := meService.Login(data.Name, data.Password)
	if e != nil {
		response.JsonExit(r, 1, e.Error())
	}
	loginResponse := &LoginResponse{
		Entity: entity,
		Token:  token,
	}
	response.JsonOk(r, loginResponse)
}

func (this *Me) Index(r *ghttp.Request) {
	// 缓存拿
	token := r.GetHeader("token")
	entity, e := meService.Detail(token)
	if e != nil {
		response.JsonExit(r, 1, e.Error())
	}
	response.JsonOk(r, entity)

}

func (this *Me) Update(r *ghttp.Request) {

}
