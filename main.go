package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"pruebaAPI/database"
	"pruebaAPI/routes"
)

func main() {
	godotenv.Load()

	database.Conectar()

	router := gin.Default()
	routes.RutasAlbumes(router)

	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
