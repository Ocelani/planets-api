package planet

import (
	"os/exec"
	"planets-api/api/database"
	"strconv"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := exec.Command("docker-compose", "up", "-d").Run(); err != nil {
		panic(err)
	}
}

func TestNewRepository(t *testing.T) {
	var (
		repo = NewRepository()
		typ  *repository
		imp  *Repository
	)
	assert.IsType(t, typ, repo)
	assert.Implements(t, imp, repo)
}

func Test_repository_Create(t *testing.T) {
	r := &repository{
		Collection: database.Connection().Collection("planets"),
	}
	t.Parallel()
	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var tt Planet
			fuzz.New().Fuzz(&tt)
			got, err := r.Create(&tt)
			assert.NoError(t, err)
			assert.Equal(t, &tt, got)
		})
	}
}

func Test_repository_ReadAll(t *testing.T) {
	r := &repository{
		Collection: database.Connection().Collection("planets"),
	}
	got, err := r.ReadAll()
	assert.NoError(t, err)
	var want []*Planet
	assert.IsType(t, want, got)
}

func Test_repository_ReadOne(t *testing.T) {
	r := &repository{
		Collection: database.Connection().Collection("planets"),
	}
	t.Parallel()
	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var tt Planet
			fuzz.New().Fuzz(&tt)
			got, err := r.Create(&tt)
			assert.NoError(t, err)
			assert.Equal(t, &tt, got)
			id := tt.ID.Hex()
			got, err = r.ReadOne(id)
			assert.NoError(t, err)
			assert.Equal(t, &tt, got)
		})
	}
}

func Test_repository_Update(t *testing.T) {
	// TODO: Add test cases.
}

func Test_repository_Delete(t *testing.T) {
	r := &repository{
		Collection: database.Connection().Collection("planets"),
	}
	t.Parallel()
	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var tt Planet
			fuzz.New().Fuzz(&tt)
			got, err := r.Create(&tt)
			assert.NoError(t, err)
			assert.Equal(t, &tt, got)
			id := tt.ID.Hex()
			assert.NoError(t, r.Delete(id))
		})
	}
}
