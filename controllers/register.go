package controllers

import (
	"ginweb/entity"
	"ginweb/middleware"
	"ginweb/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
)

// RegisterController register controller
type RegisterController struct {
}

// Router setup router
func (r *RegisterController) Router(c *gin.Engine) {
	c.POST("register", r.register)           // 用户注册获取token
	c.POST("wxGetUserInfo", r.wxGetUserInfo) // 用户登录获取token
}

// register 注册处理
func (r *RegisterController) register(c *gin.Context) {
	var arg = new(entity.RegisterArg)
	err := c.ShouldBindWith(arg, binding.FormPost)
	if err != nil {
		glog.Errorf("register arg err: %#v", err)
		showMsg(c, err.Error())
		return
	}

	glog.Infof("register arg %#v", arg)
	if arg.Code == "" {
		showMsg(c, "filed missing")
		return
	}
	openid, err := service.Register(arg)
	if err != nil {
		glog.Errorf("register err: %#v", err)
		showMsg(c, err.Error())
		return
	}
	token, err := middleware.GenToken(openid)
	if err != nil {
		glog.Errorf("register gen token err: %#v", err)
		showMsg(c, err.Error())
		return
	}
	data := gin.H{"token": token}
	jsonResult(c, data)
}

// wxGetUserInfo 获取微信用户信息
func (r *RegisterController) wxGetUserInfo(c *gin.Context) {
	// 解析token
	token := c.GetHeader("Authorization")
	if token == "" {
		glog.Error("wxGetUserInfo token empty")
		abortError(c, "token empty")
		return
	}
	openid, err := middleware.ParseToken(token)
	if err != nil {
		glog.Errorf("wxGetUserInfo ParseToken err:%v", err)
		abortError(c, err.Error())
		return
	}
	// 解析验证参数
	var arg = new(entity.WXUserInfoArg)
	err = c.ShouldBindWith(arg, binding.FormPost)
	if err != nil {
		glog.Errorf("wxGetUserInfo arg err: %#v", err)
		showMsg(c, err.Error())
		return
	}
	err = service.VerifyUserInfo(arg, openid)
	if err != nil {
		glog.Errorf("wxGetUserInfo verify err: %#v", err)
		showMsg(c, err.Error())
		return
	}
	// 签发token
	newToken, err := middleware.GenToken(openid)
	if err != nil {
		glog.Errorf("wxGetUserInfo genToken err: %#v", err)
		showMsg(c, err.Error())
		return
	}
	data := gin.H{"token": newToken}
	jsonResult(c, data)
}
