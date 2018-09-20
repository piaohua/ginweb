package routers

import (
	"fmt"
	"net/http"

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

    // Register Router
	new(controllers.RegisterController).Router(r)

	// 判断是否开启jwt验证
	if service.GetConfigJWT() {

        // 设置jwt secret
        middleware.SetSignKey(service.GetConfigJWTSecret())

		r.Use(middleware.JWTAuth())
	}

	return r
}

var db = make(map[string]string)

// SetupRouter2 setup router
func SetupRouter2() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Printf("user %s\n", user)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "error"})
		}
	})

	return r
}
