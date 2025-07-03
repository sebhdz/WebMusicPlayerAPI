package main

import (
	"github.com/gin-gonic/gin"
	"pruebaAPI/database"
	"pruebaAPI/routes"
)

func main() {
	database.Conectar()
	router := gin.Default()
	routes.RutasAlbumes(router)
	router.Run(":3000")
}
