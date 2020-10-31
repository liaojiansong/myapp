package tagService

import (
	"errors"

	"gf-app/app/model/articleTagModel"
	"gf-app/app/model/tagModel"
	"github.com/gogf/gf/os/glog"
)

type IndexRequest struct {
	Page     int `v:"page@min:1"`
	PageSize int `v:"page_size@min:1"`
	KeyWord  string
}

type IndexResponse struct {
	Items []*tagModel.Entity `json:"items"`
	Count int                `json:"count"`
}

// 主页
func Index(index *IndexRequest) (res *IndexResponse, err error) {
	builder := tagModel.Model
	if index.KeyWord != "" {
		builder.Where("name like ?", "%"+index.KeyWord+"%")
	}
	entities, err := builder.Page(index.Page, index.PageSize).All()
	if err != nil {
		return nil, err
	}
	count, err := builder.Count()
	if err != nil {
		return nil, err
	}
	return &IndexResponse{
		Items: entities,
		Count: count,
	}, nil

}

type OptionRequest struct {
	Keyword string
}

// 选项
func Options(option *OptionRequest) (res []*tagModel.Entity, err error) {
	model := tagModel.Model
	if option.Keyword != "" {
		model = model.Where("name like ?", "%"+option.Keyword+"%")
	}
	entities, err := model.FindAll()
	if err != nil {
		return nil, err
	}
	return entities, nil
}

type CreateRequest struct {
	Name string `v:"name@required"`
}

type CreateResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 创建
func Create(data *CreateRequest) (res *CreateResponse, err error) {
	isExist, err := checkName(data.Name, 0)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("name is existed")
	}
	if isExist {
		return nil, errors.New("name is existed")
	}

	result, err := tagModel.Model.Insert(data)
	if err != nil {
		return nil, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	res = &CreateResponse{
		Id:   int(lastInsertId),
		Name: data.Name,
	}
	return res, nil
}

type UpdateRequest struct {
	Id   int    `v:"id@required|min:1"`
	Name string `v:"name@required"`
}
type UpdateResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 更新
func Update(data *UpdateRequest) (res *UpdateResponse, err error) {
	// 获取,查重名
	entity, err := tagModel.Model.FindOne(data.Id)
	if err != nil {
		return nil, err
	}
	isExist, err := checkName(data.Name, data.Id)
	if err != nil {
		glog.Error(err)
		return nil, errors.New("name is existed")
	}
	if isExist {
		return nil, errors.New("name is existed")
	}

	entity.Name = data.Name
	result, err := entity.Update()
	if err != nil {
		glog.Error(err)
		return nil, errors.New("update failed")
	}
	affected, err := result.RowsAffected()
	if affected != 1 {
		glog.Error(err)
		return nil, errors.New("update failed")
	}

	res = &UpdateResponse{
		Id:   data.Id,
		Name: data.Name,
	}
	return res, nil
}

type DeleteRequest struct {
	Id int `v:"id@min:1"`
}

// 删除
func Delete(data *DeleteRequest) (ok bool, err error) {
	// 先查找
	entity, err := tagModel.Model.FindOne(data.Id)
	if err != nil {
		return false, err
	}
	if entity == nil {
		return false, errors.New("data is not exist")
	}
	// 是否关联
	has, err := hasArt(data.Id)

	if err != nil {
		return false, err
	}
	if has == true {
		return false, errors.New("已经关联文章,不能删除")
	}
	// 删除
	result, err := entity.Delete()
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected != 1 {
		return false, errors.New("delete failed")
	}
	return true, nil
}

// 是否关联文章
func hasArt(tagId int) (ok bool, err error) {
	one, err := articleTagModel.Model.FindOne("tag_id = ?", tagId)
	if err != nil {
		return false, err
	}
	if one != nil {
		return true, nil
	}
	return false, nil
}

// 检查名字
func checkName(name string, excludeId int) (isExist bool, err error) {
	model := tagModel.Model.Where("name = ?", name)
	if excludeId != 0 {
		model.Where("id <> ?", excludeId)
	}
	entity, err := model.FindOne()
	if err != nil {
		return false, err
	}
	if entity != nil {
		return true, nil
	}
	return false, nil
}
