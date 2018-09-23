package routers

import (
	"ginweb/controllers"
	"ginweb/entity"
	"ginweb/middleware"
	"ginweb/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter setup router
func SetupRouter(config *entity.Config) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 判断是否开启跨域模式
	if service.GetConfigCors() {
		r.Use(middleware.Cors())
	}

	// 设置gin开发模式
	gin.SetMode(config.GinMode)

	// 静态资源
	r.Static("/assets", "./assets")
	//r.StaticFS("/assets/static", http.Dir("./assets/static"))
	//r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// Register Router
	new(controllers.RegisterController).Router(r)
	// Example Router
	new(controllers.ExampleController).Router(r)

	// 判断是否开启jwt验证
	if service.GetConfigJWT() {

		// 设置jwt secret
		middleware.SetSignKey(service.GetConfigJWTSecret())

		r.Use(middleware.JWTAuth())
	}

	return r
}
