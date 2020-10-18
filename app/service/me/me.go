package me

import (
	"errors"
	"time"

	"gf-app/app/model/me"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/grand"
)

const USER_INFO_P = "TOKEN:"
const KEY = "token"

type UserCache struct {
	Id     int64
	Name   string
	Avatar string
	Motto  string
}

/**
检查缓存是否有
*/
func CheckLogin(token string) (bool, error) {
	token = CovToken(token)
	conn := g.Redis().Conn()
	defer conn.Close()
	reply, err := conn.DoVar("EXISTS", (token))
	if err != nil {
		return false, err
	}
	return reply.Bool(), nil
}

/**

 */
func Login(name string, password string) (string, *me.Entity, error) {
	// 查库
	one, e := me.Model.FindOne("name=? and password=?", name, password)
	if e != nil {
		return "", nil, e
	}
	if one == nil {
		return "", nil, errors.New("查无此人")
	}
	// 生成token
	token, e := genToken()
	if e != nil {
		return "", nil, e
	}
	// 缓存
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.DoVarWithTimeout(time.Hour, "SET", CovToken(token), one)
	return token, one, nil
}

func Detail(token string) (*me.Entity, error) {
	token = CovToken(token)
	conn := g.Redis().Conn()
	defer conn.Close()
	reply, err := conn.DoVar("EXISTS", (token))
	if err != nil {
		glog.Debug(err.Error())
	}
	b := reply.Bool()
	if b == false {
		return nil, errors.New("用户不存在")
	}

	res, err := conn.DoVar("GET", token)
	if err != nil {
		return nil, err
	}
	var me *me.Entity
	err = res.Struct(&me)
	if err != nil {
		return nil, err
	}
	return me, nil
}

func Load(token string) (*UserCache, error) {
	token = CovToken(token)
	conn := g.Redis().Conn()
	defer conn.Close()
	res, err := conn.DoVar("GET", token)
	if err != nil {
		return nil, err
	}
	var me *UserCache
	err = res.Struct(&me)
	if err != nil {
		return nil, err
	}
	return me, nil
}

/**
获取token
*/
func genToken() (string, error) {
	encryptString, e := gmd5.EncryptString(grand.S(16))
	return encryptString, e
}

func CovToken(token string) string {
	return USER_INFO_P + token
}
