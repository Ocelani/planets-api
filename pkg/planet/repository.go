package planet

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Repository interface allows us to access the CRUD Operations of mongoDB.
	Repository interface {
		Create(planet *Planet) (*Planet, error)
		ReadAll() (*[]Planet, error)
		ReadOne(id string) (*Planet, error)
		Update(planet *Planet) (*Planet, error)
		Delete(id string) error
	}
	repository struct {
		Collection *mongo.Collection
	}
)

// NewRepository constructor instantiates a new Repository.
func NewRepository(collection *mongo.Collection) Repository {
	return &repository{Collection: collection}
}

// Create just register a planet data in database.
func (r *repository) Create(planet *Planet) (*Planet, error) {
	planet.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(
		context.Background(),
		planet,
	)
	if err != nil {
		return nil, err
	}
	return planet, nil
}

// ReadAll returns the entire data found in planets mongoDB collection.
func (r *repository) ReadAll() (*[]Planet, error) {
	var planets []Planet
	cursor, err := r.Collection.Find(
		context.Background(),
		bson.D{},
	)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var planet Planet
		_ = cursor.Decode(&planet)
		planets = append(planets, planet)
	}
	return &planets, nil
}

// ReadOne finds and returns the data of a single planet.
func (r *repository) ReadOne(id string) (*Planet, error) {
	var planet Planet
	mg := r.Collection.FindOne(
		context.Background(),
		bson.M{"_id": id},
	)
	if err := mg.Decode(&planet); err != nil {
		return nil, err
	}
	return &planet, nil
}

// Update searches the planet parameter ID, then, updates its data in database.
func (r *repository) Update(planet *Planet) (*Planet, error) {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": planet.ID},
		bson.M{"$set": planet},
	)
	if err != nil {
		return nil, err
	}
	return planet, nil
}

// Delete the specific planet data in database with its id as a parameter.
func (r *repository) Delete(id string) error {
	planetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": planetID},
	)
	if err != nil {
		return err
	}
	return nil
}
