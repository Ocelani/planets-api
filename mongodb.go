package main

import "go.mongodb.org/mongo-driver/mongo"

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Database settings (insert your own database name and connection URI)
const dbName = "fiber_test"
const mongoURI = "mongodb://user:password@localhost:27017/" + dbName
