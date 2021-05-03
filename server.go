package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yimialmonte/gin-rest-api/controller"
	"github.com/yimialmonte/gin-rest-api/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, videoController.Save(ctx))
	})

	server.Run(":8080")
}
