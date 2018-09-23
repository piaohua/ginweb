package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExampleController example controller
type ExampleController struct {
	BaseController
}

// Router setup router
func (c *ExampleController) Router(r *gin.Engine) {
	r.GET("/example/ping", c.ping)            //
	r.GET("/example/user/:name", c.userValue) //
}

var db = make(map[string]string)

// ping test
func (c *ExampleController) ping(r *gin.Context) {
	c.Context = r
	c.String(http.StatusOK, "pong")
}

// get user value
func (c *ExampleController) userValue(r *gin.Context) {
	c.Context = r
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

/*
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
*/
