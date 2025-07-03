package routes

import (
	"github.com/gin-gonic/gin"
	"pruebaAPI/controllers"
)

func RutasAlbumes(router *gin.Engine) {
	albumes := router.Group("albumes")
	{
		albumes.GET("/", controllers.GetAlbums)
	}
}
