package main

import (
	"common/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/img", api.SaveImg)

	r.GET("/img/:id", api.GetImg)

	r.Run(":8081")
}
