package routes

import (
	"github.com/gin-gonic/gin"
	"pruebaAPI/controllers"
)

func RutasAlbumes(router *gin.Engine) {
	albumesGrupo := router.Group("albumes")
	{
		albumesGrupo.GET("", controllers.GetAlbums)
		albumesGrupo.GET("/:id", controllers.GetCanciones)
	}
}
