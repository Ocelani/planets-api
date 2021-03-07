package planet

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet entity.
type Planet struct {
	ID      primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name    string             `json:"name" binding:"required,min=2" bson:"name"` // Nome
	Climate string             `json:"climate" bson:"climate,omitempty"`          // Clima
	Terrain string             `json:"terrain" bson:"terrain,omitempty"`          // Terreno, Solo
	Films   string             `json:"films" bson:"films,omitempty"`              // Nº de aparições em filmes
}
