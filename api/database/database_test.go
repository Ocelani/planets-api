package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnection(t *testing.T) {
	var want *mongo.Database
	db := Connection()
	assert.IsType(t, want, db)
}
