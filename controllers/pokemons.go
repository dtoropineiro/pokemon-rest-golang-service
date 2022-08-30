package controllers

import (
	"context"
	"dario/pokemon-service/configs"
	"dario/pokemon-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var pokemonCollection = configs.GetCollection(configs.DB)

func GetAllPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var pokemons []models.Pokemon
		defer cancel()
		/* TODO: refactor to call paginated method
		res := configs.GetCollectionPaginated(ctx, configs.DB, 5, 1)
		*/

		results, err := pokemonCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {

			}
		}(results, ctx)
		for results.Next(ctx) {
			var singlePokemon models.Pokemon
			if err = results.Decode(&singlePokemon); err != nil {
				c.JSON(http.StatusNotFound, err.Error())
			}

			pokemons = append(pokemons, singlePokemon)
		}

		c.JSON(http.StatusOK, pokemons)
	}
}

func GetPokemonByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		name := c.Param("name")
		var pokemon models.Pokemon
		defer cancel()

		err := pokemonCollection.FindOne(ctx, bson.M{"name": name}).Decode(&pokemon)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

func GetPokemonsByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		pokemonType := c.Query("pokemonType")
		var pokemons []models.Pokemon
		defer cancel()
		results, err := pokemonCollection.Find(ctx, bson.M{"type": pokemonType})

		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		for results.Next(ctx) {
			var singlePokemon models.Pokemon
			if err = results.Decode(&singlePokemon); err != nil {
				c.JSON(http.StatusNotFound, err.Error())
			}
			pokemons = append(pokemons, singlePokemon)
		}
		c.JSON(http.StatusOK, pokemons)
	}
}
