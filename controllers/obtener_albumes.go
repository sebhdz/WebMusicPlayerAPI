package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pruebaAPI/services"
)

func GetAlbums(router *gin.Context) {
	albumes, err := services.ObtenerAlbumes()
	if err != nil {
		router.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el albumes"})
		return
	}
	router.IndentedJSON(http.StatusOK, albumes)
}
