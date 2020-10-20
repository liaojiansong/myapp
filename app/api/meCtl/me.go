package meCtl

import (
	"gf-app/app/api"
	"gf-app/app/service/meService"
	"gf-app/library/response"
	"gf-app/library/weather"
	"github.com/gogf/gf/net/ghttp"
)

type Me struct {
	*api.Di
}

// 主页
type IndexResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Motto  string `json:"motto"`

	City      string `json:"city"`
	Area      string `json:"area"`
	Temp      string `json:"temp"`
	Condition string `json:"condition"`
}

// 主页
func (this *Me) Index(r *ghttp.Request) {
	index := &IndexResponse{}
	// 天气部分
	w := weather.GetWeather("887")
	index.City = w.Data.City.Secondaryname
	index.Area = w.Data.City.Pname
	index.Temp = w.Data.Condition.Temp
	index.Condition = w.Data.Condition.Condition

	index.Id = this.Me.Id
	index.Name = this.Me.Name
	index.Avatar = this.Me.Avatar
	index.Motto = this.Me.Motto

	response.JsonOk(r, index)
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
