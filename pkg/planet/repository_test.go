package planet

import (
	"planets-api/api/database"
	"planets-api/internal"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = database.Connection().Collection("planets")

func TestNewRepository(t *testing.T) {
	var want Repository
	assert.IsType(t, want, NewRepository(collection))
}

func Test_repository_Create(t *testing.T) {
	internal.GetPlanetsFromAPI()
	// var planets []Planet
	// json.Deco
	// tests := []struct {
	// 	name    string
	// 	fields  fields
	// 	args    args
	// 	want    *Planet
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		r := &repository{
	// 			Collection: tt.fields.Collection,
	// 		}
	// 		got, err := r.Create(tt.args.planet)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("repository.Create() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func Test_repository_ReadAll(t *testing.T) {
	type fields struct {
		Collection *mongo.Collection
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				Collection: tt.fields.Collection,
			}
			got, err := r.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_ReadOne(t *testing.T) {
	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				Collection: tt.fields.Collection,
			}
			got, err := r.ReadOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.ReadOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.ReadOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		planet *Planet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				Collection: tt.fields.Collection,
			}
			got, err := r.Update(tt.args.planet)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				Collection: tt.fields.Collection,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("repository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
