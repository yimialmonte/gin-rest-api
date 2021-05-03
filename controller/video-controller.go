package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yimialmonte/gin-rest-api/entity"
	"github.com/yimialmonte/gin-rest-api/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(s service.VideoService) VideoController {
	return &controller{
		service: s,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
