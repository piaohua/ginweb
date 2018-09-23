package service

import (
	"errors"
	"time"

	"ginweb/entity"
	"ginweb/libs"
	"ginweb/models"

	"github.com/golang/glog"
)

// Register regist 授权注册
func Register(arg *entity.RegisterArg) (openid string, err error) {
	conf := GetConfig()
	if conf == nil {
		glog.Errorf("config missing register arg: %#v", arg)
		err = errors.New("config missing")
		return
	}
	// 请求微信接口
	result, err := libs.Jscode2Session(arg.Code, conf.WXAppid, conf.WXAppSecret)
	if err != nil {
		glog.Errorf("code2session fail arg: %#v, err %v", arg, err)
		return
	}
	if result.OpenId == "" {
		glog.Errorf("openid empty register arg: %#v, result: %#v", arg, result)
		err = errors.New("openid empty")
		return
	}

	engine := GetDefaultEngine()
	if engine == nil {
		glog.Errorf("orm colsed register arg: %#v", arg)
		err = errors.New("orm closed")
		return
	}
	// 更新数据库操作
	var data models.BaseData
	has, err := engine.Cols("id", "openid", "session_key", "is_new").Where("openid = ?", result.OpenId).Get(&data)
	if err != nil {
		return
	}
	if has {
		// 更新session key
		data.SessionKey = result.SessionKey
		affected, err2 := engine.Id(data.ID).Cols("session_key").Update(&data)
		glog.Infof("register update : %d", affected)
		if err2 != nil {
			glog.Errorf("register update err: %#v", err2)
			err = err2
		} else {
			openid = result.OpenId
		}
		return
	}
	// TODO 结构优化
	data.OpenID = result.OpenId
	data.SessionKey = result.SessionKey
	data.ShareType = arg.ShareType
	data.Scene = arg.Scene
	data.Cover = arg.Cover
	data.Gold = 50000  // TODO 配置优化
	data.Diamond = 200 // TODO 配置优化
	affected, err := engine.Insert(&data)
	glog.Infof("register insert : %d", affected)
	if err != nil {
		glog.Errorf("register insert err: %#v", err)
		return
	}
	openid = result.OpenId
	// TODO 添加日志
	return
}

// VerifyUserInfo 验证微信用户信息
func VerifyUserInfo(arg *entity.WXUserInfoArg, openid string) (err error) {
	sessionKey, err := GetSessionKey(openid)
	if err != nil {
		return
	}
	////验证sha1( rawData + sessionKey )
	sign := libs.Sha1Signature(arg.RawData, sessionKey)
	if sign != arg.Signature {
		err = errors.New("sign failed")
		return
	}
	// 验证敏感信息
	b, err := libs.DecryptWechatAppletUser(arg.EncryptedData, sessionKey, arg.Iv)
	if err != nil {
		return
	}
	wxUserInfo := new(entity.WxUserInfo)
	err = libs.ParseUserInfo(b, wxUserInfo)
	if err != nil {
		return
	}
	if wxUserInfo.Watermark.Appid != GetConfigWXAppid() {
		err = errors.New("appid error")
		return
	}
	// 10秒内请求有效
	if (time.Now().Unix() - wxUserInfo.Watermark.Timestamp) > 10 {
		err = errors.New("session expire")
		return
	}
	if openid != wxUserInfo.OpenId {
		err = errors.New("openid error")
	}
	// TODO 更新存储微信用户信息
	return
}

// GetSessionKey get sessionKey by openid
func GetSessionKey(openid string) (sessionKey string, err error) {
	engine := GetDefaultEngine()
	if engine == nil {
		glog.Errorf("orm colsed GetSessionKey openid: %s", openid)
		err = errors.New("db colsed")
		return
	}
	user := new(models.BaseData)
	has, err := engine.Cols("session_key").Where("openid = ?", openid).Get(user)
	if err != nil {
		return
	}
	if !has {
		err = errors.New("not exist")
		return
	}
	sessionKey = user.SessionKey
	return
}
