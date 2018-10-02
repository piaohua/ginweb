package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应json数据
func jsonResult(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

// 响应msg错误信息
func showMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
	})
}

// 响应msg错误信息
func abortError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
	})
}
