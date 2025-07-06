package main

import (
	"clear-cart/config"
	"clear-cart/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis()
	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":3039"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
