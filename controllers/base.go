package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseController base controller
type BaseController struct {
    *gin.Context
    Openid string
}

// 响应json数据
func (c *BaseController) jsonResult(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "data": data,
    })
}

// 响应msg错误信息
func (c *BaseController) showMsg(msg string) {
	c.JSON(http.StatusOK, gin.H{
        "code": http.StatusBadRequest,
        "msg": msg,
    })
}

// 响应msg错误信息
func (c *BaseController) abortError(msg string) {
    c.AbortWithStatusJSON(http.StatusOK, gin.H{
        "code": http.StatusBadRequest,
        "msg": msg,
    })
}
