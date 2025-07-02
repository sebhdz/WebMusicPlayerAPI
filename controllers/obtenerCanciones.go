package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pruebaAPI/database"
)

func ObtenerCanciones(router *gin.Context) {
	router.IndentedJSON(http.StatusOK, database.Canciones)
}
