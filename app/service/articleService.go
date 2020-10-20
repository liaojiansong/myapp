package service

import (
	"errors"
	"gf-app/app/model/articleModel"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type IndexRequest struct {
	Page     int `v:"page@required|min:1"`
	PageSize int `v:"page_size@required|min:1"`
	KeyWord  string
}
type IndexResponse struct {
	Items []*articleModel.Entity `json:"items"`
	Count int
}

// 列表
func ArtIndex(meId int, data *IndexRequest) (*IndexResponse, error) {
	builder := articleModel.Model.Where("me_id = ?", meId).Page(data.Page, data.PageSize)
	if data.KeyWord != "" {
		builder = builder.Where("title like ?", "%"+data.KeyWord+"%")
	}
	builder = builder.Where("me_id = ?", meId)
	count, err := builder.Count()
	if err != nil {
		return nil, err
	}
	builder = builder.Fields("id", "title", "img", "is_publish", "content", "created_time")
	entities, err := builder.Page(data.Page, data.PageSize).FindAll()
	return &IndexResponse{
		Items: entities,
		Count: count,
	}, nil
}

type ArtDetailRequest struct {
	Id int `v:"id@min:1"`
}
type ArtDetailResponse struct {
	*articleModel.Entity
}

func ArtDetail(meId int, data *ArtDetailRequest) (*ArtDetailResponse, error) {
	entity, err := articleModel.Model.FindOne("id = ? and me_id = ?", data.Id, meId)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("data not existed")
	}
	return &ArtDetailResponse{entity}, nil
}

type ArticleService struct {
}

// 创建文章请求体
type CreateArtRequest struct {
	Title     string `v:"title@required|length:3,64"`
	Img       string `v:"img@required"`
	Content   string `v:"content@required|min-length:3"`
	IsPublish int    `v:"is_publish@required|boolean"`
}

// 创建文章响应体
type ArtResponse struct {
	Id int `json:"id"`
}

func CreatArt(meId int, art *CreateArtRequest) (*ArtResponse, error) {
	model := &articleModel.Entity{
		MeId:        meId,
		Title:       art.Title,
		Img:         art.Img,
		Content:     art.Content,
		IsPublish:   art.IsPublish,
		CreatedTime: gtime.Now(),
		UpdatedTime: gtime.Now(),
	}
	result, err := model.Save()
	if err != nil {
		return nil, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	response := &ArtResponse{
		Id: gconv.Int(lastInsertId),
	}
	return response, nil
}

type UpdateArtRequest struct {
	Id int `v:"id@min:1"`
	*CreateArtRequest
}

// 更新
func UpdateArt(meId int, art *UpdateArtRequest) (*ArtResponse, error) {
	entity, err := articleModel.Model.First(art.Id)
	if err != nil {
		return nil, err
	}
	entity.Title = art.Title
	entity.Img = art.Img
	entity.Content = art.Content
	entity.IsPublish = art.IsPublish
	entity.UpdatedTime = gtime.Now()
	result, err := entity.Update()
	if err != nil {
		return nil, err
	}
	eff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if eff != 1 {
		return nil, errors.New("更新失败")
	}
	return &ArtResponse{
		Id: art.Id,
	}, nil
}

type ArtDeleteRequest struct {
	Id int `v:"id@min:1"`
}

// 删除
func Delete(meId int, data *ArtDeleteRequest) error {
	entity, err := articleModel.Model.FindOne("id = ? and me_id = ?", data.Id, meId)
	if err != nil {
		return err
	}
	if entity == nil {
		return errors.New("data not existed")
	}
	result, err := entity.Delete()
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("删除失败")
	}
	return nil
}
