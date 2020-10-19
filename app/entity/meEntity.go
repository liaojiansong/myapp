package entity

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
	"time"
)

const USER_INFO_P = "USER:"

type MeCache struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Motto  string `json:"motto"`
}

func NewMeCache(id int, name string, avatar string, motto string) *MeCache {
	return &MeCache{Id: id, Name: name, Avatar: avatar, Motto: motto}
}

// 是否存在
func ExistsMe(token string) (exists bool, err error) {
	token = USER_INFO_P + token
	conn := g.Redis().Conn()
	defer conn.Close()
	reply, err := conn.DoVar("EXISTS", token)
	return reply.Bool(), err
}

// 加载
func LoadMe(token string) (cache *MeCache, err error) {
	token = USER_INFO_P + token
	conn := g.Redis().Conn()
	defer conn.Close()

	res, err := conn.DoVar("GET", token)
	if err != nil {
		return nil, err
	}
	var me *MeCache
	err = res.Struct(&me)
	return me, nil
}

// 保存
func (this *MeCache) StoreMe() (token string, err error) {
	token, err = this.genToken()
	if err != nil {
		return "", err
	}
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.DoVarWithTimeout(time.Hour, "SET", USER_INFO_P+token, this)
	return
}

// 摧毁
func DestroyMe(token string) (ok bool, err error) {
	token = USER_INFO_P + token
	conn := g.Redis().Conn()
	defer conn.Close()
	res, err := conn.DoVar("DEL", token)
	if err != nil {
		return false, err
	}
	b := res.Bool()
	return b, nil
}

func (this *MeCache) genToken() (token string, err error) {
	token, err = gmd5.EncryptString(grand.S(16))
	return
}
