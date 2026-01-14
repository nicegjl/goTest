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
	r.GET("/getData", getDataHandler)
	r.Run(":8080")
}

func getDataHandler(c *gin.Context) {
	resp := GetDataResponse{
		Code: 0,
		Msg:  "ok",
		Data: "hello from getData",
	}
	c.JSON(http.StatusOK, resp)
}
