package routes

import (
	"dario/pokemon-service/controllers"
	"github.com/gin-gonic/gin"
)

func PokemonRoute(router *gin.Engine) {
	router.GET("api/v1/pokemon/", controllers.GetAllPokemon())
	router.GET("api/v1/pokemon/:name", controllers.GetPokemonByName())
}
