package main

import (
	"github.com/gin-gonic/gin"
	"pruebaAPI/routes"
)

func main() {
	router := gin.Default()
	routes.RutasCanciones(router)
	router.Run(":3000")
}
