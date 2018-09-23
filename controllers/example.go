package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	r.GET("/example/welcome", c.welcome)      //
	r.POST("/example/form_post", c.formPost)  //
	r.POST("/example/upload", c.upload)       //
	r.POST("/example/loginJSON", c.loginJSON) //
	r.POST("/example/loginForm", c.loginForm) //

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/example", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	authorized.POST("admin", c.admin)
	r.POST("/example/login", c.login)
	r.GET("/example/someJSON", c.someJSON)
	r.GET("/example/moreJSON", c.moreJSON)
	r.GET("/example/someXML", c.someXML)
	//
	r.LoadHTMLGlob("./views/*")
	r.GET("/example/index", c.index)
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

// admin
func (c *ExampleController) admin(r *gin.Context) {
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
}

// welcome
func (c *ExampleController) welcome(r *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

// formPost
func (c *ExampleController) formPost(r *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

// upload
func (c *ExampleController) upload(r *gin.Context) {

	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}

// Login Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Example for binding JSON ({"user": "manu", "password": "123"})
func (c *ExampleController) loginJSON(r *gin.Context) {
	var json Login
	if c.BindJSON(&json) == nil {
		if json.User == "manu" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

// Example for binding a HTML form (user=manu&password=123)
func (c *ExampleController) loginForm(r *gin.Context) {
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&form) == nil {
		if form.User == "manu" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

// LoginForm login form
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// curl -v --form user=user --form password=password http://localhost:8080/login
// login
func (c *ExampleController) login(r *gin.Context) {
	// you can bind multipart form with explicit binding declaration:
	// c.BindWith(&form, binding.Form)
	// or you can simply use autobinding with Bind method:
	var form LoginForm
	// in this case proper binding will be automatically selected
	if c.Bind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

// gin.H is a shortcut for map[string]interface{}
func (c *ExampleController) someJSON(r *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func (c *ExampleController) moreJSON(r *gin.Context) {
	// You also can use a struct
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// Note that msg.Name becomes "user" in the JSON
	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	c.JSON(http.StatusOK, msg)
}

func (c *ExampleController) someXML(r *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
func (c *ExampleController) index(r *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
