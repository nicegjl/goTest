package main

import (
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


	r.Run(":8080")
}
