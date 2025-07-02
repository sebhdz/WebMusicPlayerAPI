package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	Nombre  string `json:"nombre"`
	Artista string `json:"artist"`
}

var albumes = []album{
	{Nombre: "Tu hogar", Artista: "Tino el pinguino"},
	{Nombre: "Lo legal", Artista: "El bebeto"},
}

func Pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumes)
}

func main() {
	router := gin.Default()
	router.GET("/ping", Pong)
	router.Run(":8080")
}
