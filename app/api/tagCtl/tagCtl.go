package tagCtl

import (
	"gf-app/app/api"
	"gf-app/app/service/tagService"
	"gf-app/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type Tag struct {
	*api.Di
}

func (this *Tag) Index(r *ghttp.Request) {
	var data *tagService.IndexRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	indexResponse, err := tagService.Index(data)
	if err != nil {
		response.JsonExit(r, -1, err.Error())
	}
	response.JsonOk(r, indexResponse)

}

func (this *Tag) Create(r *ghttp.Request) {
	var data *tagService.CreateRequest
	e := r.Parse(&data)
	if e != nil {
		response.JsonExit(r, 0, e.Error())
	}
	res, err := tagService.Create(data)
	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	response.JsonOk(r, res)
}

func (this *Tag) Update(r *ghttp.Request) {
	var data *tagService.UpdateRequest
	e := r.Parse(&data)
	if e != nil {
		response.JsonExit(r, 0, e.Error())
	}
	res, err := tagService.Update(data)
	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	response.JsonOk(r, res)
}

func (this *Tag) Delete(r *ghttp.Request) {
	var data *tagService.DeleteRequest
	e := r.Parse(&data)
	if e != nil {
		response.JsonExit(r, 0, e.Error())
	}
	res, err := tagService.Delete(data)
	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}
	response.JsonOk(r, res)
}
