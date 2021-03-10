package planet

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Repository interface allows us to access the CRUD Operations of mongoDB.
	// Repository interface {
	// 	Create(planet *Planet) (*Planet, error)
	// 	ReadAll() (*[]Planet, error)
	// 	ReadOneWithID(id string) (*Planet, error)
	// 	ReadOneWithName(name string) (*Planet, error)
	// 	Update(planet *Planet) (*Planet, error)
	// 	Delete(id string) error
	// }
	Repository struct {
		Collection *mongo.Collection
	}
)

// NewRepository constructor instantiates a new Repository.
func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		Collection: collection,
	}
}

// Create just register a planet data in database.
func (r *Repository) Create(planet *Planet) (*Planet, error) {
	if planet.Name == "" {
		return nil, fmt.Errorf("Planet name was not provided")
	}
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
func (r *Repository) ReadAll() (*[]Planet, error) {
	var planets []Planet
	cursor, err := r.Collection.Find(
		context.Background(),
		bson.D{},
	)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var p Planet
		_ = cursor.Decode(&p)
		planets = append(planets, p)
	}

	return &planets, nil
}

// bson.D{{"$match", bson.D{{"podcast", id}}}}

// ReadOneWithID finds and returns the data of a single planet.
func (r *Repository) ReadOneWithID(id string) (*Planet, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	mg := r.Collection.FindOne(
		context.Background(),
		bson.M{"_id": oID},
	)
	var p Planet
	if err := mg.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

// ReadOneWithName finds and returns the data of a single planet.
func (r *Repository) ReadOneWithName(name string) (*Planet, error) {
	mg := r.Collection.FindOne(
		context.Background(),
		bson.M{"name": name},
	)
	var p Planet
	if err := mg.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

// Update searches the planet parameter ID, then, updates its data in database.
func (r *Repository) Update(planet *Planet) (*Planet, error) {
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
func (r *Repository) Delete(id string) error {
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
