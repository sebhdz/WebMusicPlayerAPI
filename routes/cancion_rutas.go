package routes

import (
	"github.com/gin-gonic/gin"
	"pruebaAPI/controllers"
)

func RutasCanciones(router *gin.Engine) {
	albumes := router.Group("canciones")
	{
		albumes.GET("/", controllers.ObtenerCanciones)
	}
}
