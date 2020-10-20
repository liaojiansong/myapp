package articleCtl

import (
	"gf-app/app/api"
	"gf-app/app/service"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Article struct {
	*api.Di
}

// 列表
func (this Article) Index(r *ghttp.Request) {
	var data *service.IndexRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	artResponse, err := service.ArtIndex(this.Me.Id, data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, artResponse)

}

// 详情
func (this Article) Detail(r *ghttp.Request) {
	var data *service.ArtDetailRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	artResponse, err := service.ArtDetail(this.Me.Id, data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, artResponse)
}

// 创建文章
func (this Article) Create(r *ghttp.Request) {
	var data *service.CreateArtRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	artResponse, err := service.CreatArt(this.Me.Id, data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, artResponse)
}

// 更新文章
func (this Article) Update(r *ghttp.Request) {
	var data *service.UpdateArtRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	artResponse, err := service.UpdateArt(this.Me.Id, data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, artResponse)
}

// 删除文章
func (this *Article) Delete(r *ghttp.Request) {
	var data *service.ArtDeleteRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	err := service.Delete(this.Me.Id, data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, nil)
}
