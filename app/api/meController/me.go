package meController

import (
	"gf-app/app/api"
	"gf-app/app/service/meService"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Me struct {
	*api.Di
}

// 主页
func (this *Me) Index(r *ghttp.Request) {
	response.JsonOk(r, this.Me)
}

// 更新
func (this *Me) Update(r *ghttp.Request) {
	token := r.GetHeader("token")
	var data *meService.UpdateRequest
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	updateResponse, err := meService.Update(data, token)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, updateResponse)
}

// 上传头像
func (this Me) UploadAvatar(r *ghttp.Request) {
	file := r.GetUploadFile("avatar")
	if file == nil {
		response.JsonExit(r, -1, "please upload file")
	}
	filename, err := file.Save("./public/resource/image/avatars")
	if err != nil {
		response.JsonExit(r, -1, "save file failed")
	}
	response.JsonOk(r, filename)

}
