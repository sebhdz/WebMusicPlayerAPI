package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pruebaAPI/services"
)

func GetCanciones(router *gin.Context) {
	idAlbum := router.Param("id")
	canciones, err := services.ObtenerCanciones(idAlbum)
	if err != nil {
		router.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener canciones"})
		return
	}
	router.IndentedJSON(http.StatusOK, canciones)
}
