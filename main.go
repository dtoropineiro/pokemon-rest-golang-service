package main

import (
	"dario/pokemon-service/configs"
	_ "dario/pokemon-service/controllers"
	"dario/pokemon-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Connect()
	router := gin.Default()
	routes.PokemonRoute(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
