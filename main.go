package main

import (
	"github.com/gin-contrib/cors"
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

	router.Use(cors.Default())
	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	*/

	routes.RutasAlbumes(router)

	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
