package service

import (
	"errors"
	"strings"

	"gf-app/app/model/articleModel"
	"gf-app/app/model/articleTagModel"
	"gf-app/app/model/tagModel"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type IndexRequest struct {
	Page     int `v:"page@required|min:1"`
	PageSize int `v:"page_size@required|min:1"`
	KeyWord  string
	TagId    int
}
type IndexResponse struct {
	Items interface{} `json:"items"`
	Count int         `json:"count"`
}

// 列表
func ArtIndex(meId int, data *IndexRequest) (*IndexResponse, error) {
	t1 := articleModel.Alias("t1")
	t2 := articleTagModel.Alias("t2")
	builder := g.DB().Model(t1).LeftJoin(t2, "t1.id = t2.article_id")
	builder = builder.Where("t1.me_id = ?", meId)
	if data.KeyWord != "" {
		builder = builder.Where("t1.title like ?", "%"+data.KeyWord+"%")
	}
	if data.TagId > 0 {
		builder = builder.Where("t2.tag_id = ?", data.TagId)
	}
	count, err := builder.Count()
	if err != nil {
		return nil, err
	}
	results, err := builder.Fields("t1.*").Page(data.Page, data.PageSize).FindAll()
	if err != nil {
		return nil, err
	}

	return &IndexResponse{
		Items: results.List(),
		Count: count,
	}, nil
}

type ArtDetailRequest struct {
	Id int `v:"id@min:1"`
}
type ArtDetailResponse struct {
	*articleModel.Entity
	Tags interface{} `json:"tags"`
}

func ArtDetail(meId int, data *ArtDetailRequest) (*ArtDetailResponse, error) {
	entity, err := articleModel.Model.FindOne("id = ? and me_id = ?", data.Id, meId)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("data not existed")
	}
	// 查询标签
	t1 := articleTagModel.Table + " t1"
	t2 := tagModel.Table + " t2"

	results, err := g.DB().Table(t1).InnerJoin(t2, "t1.tag_id = t2.id").Where("t1.article_id = ?", data.Id).Fields("t2.id,t2.name").FindAll()
	if err != nil {
		return nil, err
	}
	results.Array()
	return &ArtDetailResponse{entity, results}, nil
}

type ArticleService struct {
}

// 创建文章请求体
type CreateArtRequest struct {
	Title     string `v:"title@required|length:3,64"`
	Img       string `v:"img@required"`
	Content   string `v:"content@required|min-length:3"`
	IsPublish int    `v:"is_publish@required|boolean"`
	Tags      string `v:"tags@min-length:1"`
}

// 创建文章响应体
type ArtResponse struct {
	Id int `json:"id"`
}

// todo 还有优化空间
func CreatArt(meId int, art *CreateArtRequest) (*ArtResponse, error) {
	var sqlErr error
	tx, sqlErr := g.DB().Begin()
	defer func() {
		if sqlErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if sqlErr != nil {
		return nil, sqlErr
	}
	data := &articleModel.Entity{
		MeId:        meId,
		Title:       art.Title,
		Img:         art.Img,
		Content:     art.Content,
		IsPublish:   art.IsPublish,
		CreatedTime: gtime.Now(),
		UpdatedTime: gtime.Now(),
	}
	result, sqlErr := tx.Model(articleModel.Table).Insert(data)
	if sqlErr != nil {
		return nil, sqlErr
	}
	lastInsertId, sqlErr := result.LastInsertId()
	if sqlErr != nil {
		return nil, sqlErr
	}
	// 解析标签
	tagsList := buildTagRe(art.Tags, lastInsertId)
	_, sqlErr = tx.Model(articleTagModel.Table).Insert(tagsList)
	if sqlErr != nil {
		return nil, sqlErr
	}

	response := &ArtResponse{
		Id: gconv.Int(lastInsertId),
	}
	return response, nil
}

// 构建关系
func buildTagRe(tags string, articleId int64) []*g.Map {
	tagsList := make([]*g.Map, 0, 0)
	splitTags := strings.Split(tags, ",")
	for _, v := range splitTags {
		i := gconv.Int(v)
		if i == 0 {
			continue
		}
		tagsList = append(tagsList, &g.Map{"article_id": articleId, "tag_id": i})
	}
	return tagsList
}

type UpdateArtRequest struct {
	Id int `v:"id@min:1"`
	*CreateArtRequest
}

// 更新
func UpdateArt(meId int, art *UpdateArtRequest) (*ArtResponse, error) {
	var sqlErr error
	entity, err := articleModel.Model.First(art.Id)
	if err != nil {
		return nil, err
	}
	entity.Id = 0
	entity.Title = art.Title
	entity.Img = art.Img
	entity.Content = art.Content
	entity.IsPublish = art.IsPublish
	entity.UpdatedTime = gtime.Now()
	tx, sqlErr := g.DB().Begin()
	defer func() {
		if sqlErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 更新主表
	result, sqlErr := tx.Model(articleModel.Table).WherePri(art.Id).OmitEmpty().Update(entity)
	if sqlErr != nil {
		return nil, sqlErr
	}
	eff, sqlErr := result.RowsAffected()
	if sqlErr != nil {
		return nil, sqlErr
	}
	if eff != 1 {
		sqlErr := errors.New("更新失败")
		return nil, sqlErr
	}
	// 删除便签
	_, sqlErr = tx.Model(articleTagModel.Table).Delete("article_id = ?", art.Id)
	if sqlErr != nil {
		return nil, sqlErr
	}
	// 新增便签
	tagsList := buildTagRe(art.Tags, int64(art.Id))
	_, sqlErr = tx.Model(articleTagModel.Table).Insert(tagsList)
	if sqlErr != nil {
		return nil, sqlErr
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
	var sqlErr error
	entity, sqlErr := articleModel.Model.FindOne("id = ? and me_id = ?", data.Id, meId)
	if sqlErr != nil {
		return sqlErr
	}
	if entity == nil {
		return errors.New("data not existed")
	}
	tx, sqlErr := g.DB().Begin()
	defer func() {
		if sqlErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	result, sqlErr := tx.Model(articleModel.Table).WherePri(data.Id).Delete()
	if sqlErr != nil {
		return sqlErr
	}
	affected, sqlErr := result.RowsAffected()
	if sqlErr != nil {
		return sqlErr
	}
	if affected != 1 {
		sqlErr := errors.New("删除失败")
		return sqlErr
	}
	_, sqlErr = tx.Model(articleTagModel.Table).Delete("article_id = ?", data.Id)
	if sqlErr != nil {
		return sqlErr
	}
	return nil
}
