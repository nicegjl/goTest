package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now();
		fmt.Println("--- 中间件开始执行 ---")
		c.Set("request", "中间件")
		code := c.Writer.Status();
		fmt.Println("--- 中间件执行完毕 ---", code)
		t2 := time.Since(t1);
		fmt.Println("Use Time:", t2)
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())

	r.LoadHTMLFiles("template/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "Hello, World!"})
	})

	v1 := r.Group("v1/api")
	{
		v1.GET("/getData", getData)
		v1.GET("/user/:id/:name", getUserFromRestful)
		v1.GET("/user", getUserFromUrl)
		v1.POST("/register", handleRegister)
		v1.POST("/login", handleLogin)
	}


	fmt.Println("http://192.168.18.177:8080")

	r.Run(":8080")
}

type GetDataResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
func getData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})

	res := GetDataResponse{
		Code: 200,
		Msg: "ok",
		Data: "hello from getData",
	}
	c.JSON(http.StatusOK, res)
}

func getUserFromRestful(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
}

func getUserFromUrl(c *gin.Context) {
	name := c.DefaultQuery("name", "gongjiali")
	age := c.DefaultQuery("age", "18")
	c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
}


// 注册接口
var user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
func handleRegister(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, _ := c.Get("request");
	fmt.Println("Middle Data: ", req)

	username := user.Username
	email := user.Email
	password := user.Password

	fmt.Println(fmt.Sprintf("username: %s, email: %s, password: %s", username, email, password))

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"email": email,
		"password": password,
		"token": "HJUSHodjuoi334rnkjsKSHNw35BHJWQ32",
	})
}

// 登录接口
type LoginRequest struct {
	User string `from: "username" json: "user" binding: "required"`
	Password string `from: "password" json: "password" binding: "required"`
}
type LoginResponse struct {
	User string `json:"user"`
	Password string `json:"password"`
	Token string `json:"token"`
}
func handleLogin(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证登录信息
	if json.User != "admin" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 结构体响应
	var info LoginResponse
	info.User = json.User
	info.Password = json.Password
	info.Token = "HJUSHodjuoi334rnkjsKSHNw35BHJWQ32"

	c.JSON(http.StatusOK, info)
}