package configs

import (
	"context"
	"dario/pokemon-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const collectionName = "pokemon"

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options.Database()
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client
}

var DB = Connect()

func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database(GetDatabaseName()).Collection(collectionName)
	return collection
}

func GetCollectionPaginated(ctx context.Context, client *mongo.Client, limit, page int) []models.Pokemon {
	collection := client.Database(GetDatabaseName()).Collection(collectionName)
	curr, err := collection.Find(ctx, bson.D{{}}, GetMongoPaginate(limit, page).GetPaginatedOpts())
	result := make([]models.Pokemon, 0)
	if err != nil {
		log.Fatal(err)
	}
	for curr.Next(ctx) {
		var pokemon models.Pokemon
		if err := curr.Decode(&pokemon); err != nil {
			log.Println(err)
		}

		result = append(result, pokemon)
	}
	return result
}
