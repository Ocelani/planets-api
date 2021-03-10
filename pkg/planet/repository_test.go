package planet

import (
	"os/exec"
	"planets-api/api/database"
	"strconv"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

// genTests defines how many tests data
// will be generated per test function.
const genTests = 100

// r connects the Repository to mongoDB.
var r *Repository

func init() {
	if err := exec.Command("docker-compose", "up", "-d").Run(); err != nil {
		panic(err)
	}
	r = &Repository{
		Collection: database.Connection().Collection("planets"),
	}
}

func TestNewRepository(t *testing.T) {
	var (
		repo = NewRepository(r.Collection)
		typ  *Repository
	)
	assert.IsType(t, typ, repo)
}

func Test_repository_Create(t *testing.T) {
	t.Parallel()
	for i := 0; i < genTests; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var p Planet
			fuzz.New().Fuzz(&p)

			got, err := r.Create(&p)
			if p.Name == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, &p, got)
		})
	}
}

func Test_repository_ReadAll(t *testing.T) {
	var want *[]Planet
	got, err := r.ReadAll()
	assert.NoError(t, err)
	assert.IsType(t, want, got)
}

func Test_repository_ReadOneWithID(t *testing.T) {
	t.Parallel()
	for i := 0; i < genTests; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var p Planet
			fuzz.New().Fuzz(&p)

			// Create
			create, err := r.Create(&p)
			if p.Name == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, &p, create)

			// Read
			id := create.ID.Hex()
			read, err := r.ReadOneWithID(id)
			assert.NoError(t, err)
			assert.Equal(t, &p, read)
		})
	}
}

func Test_repository_ReadOneWithName(t *testing.T) {
	t.Parallel()
	for i := 0; i < genTests; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var p Planet
			fuzz.New().Fuzz(&p)

			// Create
			create, err := r.Create(&p)
			if p.Name == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, &p, create)

			// Read
			read, err := r.ReadOneWithName(p.Name)
			assert.NoError(t, err)
			assert.Equal(t, create, read)
		})
	}
}

func Test_repository_Update(t *testing.T) {
	t.Parallel()
	for i := 0; i < genTests; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// Create
			var p1 Planet
			fuzz.New().Fuzz(&p1)
			create, err := r.Create(&p1)
			if p1.Name == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, &p1, create)

			// Update
			var p2 Planet
			fuzz.New().Fuzz(&p2) // New data
			p2.ID = create.ID    // Same ID
			update, err := r.Update(&p2)

			assert.NoError(t, err)
			assert.Equal(t, &p2, update)
			assert.NotEqual(t, create, update)
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	t.Parallel()
	for i := 0; i < genTests; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var p Planet
			fuzz.New().Fuzz(&p)

			// Create
			got, err := r.Create(&p)
			if p.Name == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, &p, got)

			// Delete
			id := p.ID.Hex()
			del := r.Delete(id)
			assert.NoError(t, del)
		})
	}
}
