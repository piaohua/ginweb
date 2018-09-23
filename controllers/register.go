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
	BaseController
}

// Router setup router
func (c *RegisterController) Router(r *gin.Engine) {
	r.POST("register", c.register)           // 用户注册获取token
	r.POST("wxGetUserInfo", c.wxGetUserInfo) // 用户登录获取token
}

// register 注册处理
func (c *RegisterController) register(r *gin.Context) {
	c.Context = r
	var arg = new(entity.RegisterArg)
	err := c.ShouldBindWith(arg, binding.FormPost)
	if err != nil {
		glog.Errorf("register arg err: %#v", err)
		c.showMsg(err.Error())
		return
	}

	glog.Infof("register arg %#v", arg)
	if arg.Code == "" {
		c.showMsg("filed missing")
		return
	}
	openid, err := service.Register(arg)
	if err != nil {
		glog.Errorf("register err: %#v", err)
		c.showMsg(err.Error())
		return
	}
	token, err := middleware.GenToken(openid)
	if err != nil {
		glog.Errorf("register gen token err: %#v", err)
		c.showMsg(err.Error())
		return
	}
	data := gin.H{"token": token}
	c.jsonResult(data)
}

// wxGetUserInfo 获取微信用户信息
func (c *RegisterController) wxGetUserInfo(r *gin.Context) {
	c.Context = r
	// 解析token
	token := c.GetHeader("Authorization")
	if token == "" {
		glog.Error("wxGetUserInfo token empty")
		c.abortError("token empty")
		return
	}
	openid, err := middleware.ParseToken(token)
	if err != nil {
		glog.Errorf("wxGetUserInfo ParseToken err:%v", err)
		c.abortError(err.Error())
		return
	}
	// 解析验证参数
	var arg = new(entity.WXUserInfoArg)
	err = c.ShouldBindWith(arg, binding.FormPost)
	if err != nil {
		glog.Errorf("wxGetUserInfo arg err: %#v", err)
		c.showMsg(err.Error())
		return
	}
	err = service.VerifyUserInfo(arg, openid)
	if err != nil {
		glog.Errorf("wxGetUserInfo verify err: %#v", err)
		c.showMsg(err.Error())
		return
	}
	// 签发token
	newToken, err := middleware.GenToken(openid)
	if err != nil {
		glog.Errorf("wxGetUserInfo genToken err: %#v", err)
		c.showMsg(err.Error())
		return
	}
	data := gin.H{"token": newToken}
	c.jsonResult(data)
}
