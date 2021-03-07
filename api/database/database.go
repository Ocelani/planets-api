package database

import (
	"context"
	"log"
	"planets-api/pkg/planet"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://username:password@localhost:27017/fiber"

func Connection() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	return client.Database("planets-api")
}

// Planets returns its collection repository service.
func Planets() planet.Service {
	db := Connection().Collection("planets")
	repo := planet.NewRepository(db)
	return planet.NewService(repo)
}
