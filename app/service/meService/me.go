package meService

import (
	"errors"
	"gf-app/app/entity"
	"gf-app/app/model/meModel"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

// 登入数据
type LoginRequest struct {
	Name     string `v:"required"`
	Password string `v:"required|password"`
}

// 登入响应
type LoginResponse struct {
	*entity.MeCache
	Token string `json:"token"`
}

// 登入动作
func Login(rData *LoginRequest) (data *LoginResponse, err error) {
	// 查库
	one, e := meModel.Model.FindOne("name=? and password=?", rData.Name, rData.Password)
	if e != nil {
		return nil, e
	}
	if one == nil {
		return nil, errors.New("查无此人")
	}
	newMeCache := entity.NewMeCache(one.Id, one.Name, one.Avatar, one.Motto)
	store, err := newMeCache.StoreMe()
	return &LoginResponse{
		MeCache: newMeCache,
		Token:   store,
	}, nil
}

// 更新请求体
type UpdateRequest struct {
	Id        int    `v:"id@min:1"`
	Name      string `v:"required|min-length:3|max-length:24"`
	Password  string `v:"password@required|password"`
	Password2 string `v:"password_2@required|same:password"`
	Avatar    string `v:"avatar@required"`
	Motto     string `v:"motto@required"`
}

// 更新响应体
type UpdateResponse struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Avatar      string      `json:"avatar"`
	Motto       string      `json:"motto"`
	CreatedTime *gtime.Time `json:"created_time"`
}

// 更新
func Update(data *UpdateRequest, token string) (me *UpdateResponse, err error) {
	// 检验重名
	exists, err := meModel.Model.FindOne("id <> ? and name = ?", data.Id, data.Name)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("the same name is existed")
	}
	// 查找存在
	one, err := meModel.Model.First(data.Id)
	if err != nil {
		return nil, err
	}
	one.Name = data.Name
	one.Avatar = data.Avatar
	one.Motto = data.Motto
	one.Name = data.Name
	one.Password = data.Password
	one.UpdatedTime = gtime.Now()
	result, err := one.Update()

	if err != nil {
		glog.Error(err)
		return nil, errors.New("更新失败")
	}
	affected, err := result.RowsAffected()
	if affected < 1 {
		return nil, errors.New("更新失败")
	}
	_, err = entity.DestroyMe(token)
	if err != nil {
		glog.Error(err)
	}
	return &UpdateResponse{
		Id:          one.Id,
		Name:        one.Name,
		Avatar:      one.Avatar,
		Motto:       one.Motto,
		CreatedTime: one.CreatedTime,
	}, nil
}
