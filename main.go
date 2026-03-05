package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDataResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("template/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "Hello, World!"})
	})

	r.GET("/getData", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})

		res := GetDataResponse{
			Code: 200,
			Msg: "ok",
			Data: "hello from getData",
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/user/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
	})

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "gongjiali")
		age := c.DefaultQuery("age", "18")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})


	// 注册接口
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	r.POST("/register", func(c *gin.Context) {

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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
	})


	r.Run(":8080")
}
