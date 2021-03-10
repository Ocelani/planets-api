package planet

import "go.mongodb.org/mongo-driver/mongo"

type (
	// Service is an interface from which our api module can access our repository.
	// Service interface {
	// 	Insert(planet *Planet) (*Planet, error)
	// 	FindAll() (*[]Planet, error)
	// 	FindOneWithID(id string) (*Planet, error)
	// 	FindOneWithName(name string) (*Planet, error)
	// 	Update(planet *Planet) (*Planet, error)
	// 	Remove(id string) error
	// }
	Service struct {
		repository *Repository
	}
)

// NewService is used to create a single instance of the Service.
func NewService(collection *mongo.Collection) *Service {
	return &Service{repository: NewRepository(collection)}
}

// Insert Planet.
func (s *Service) Insert(planet *Planet) (*Planet, error) {
	return s.repository.Create(planet)
}

// FindAll Planets.
func (s *Service) FindAll() (*[]Planet, error) {
	return s.repository.ReadAll()
}

// FindOne Planet.
func (s *Service) FindOneWithID(id string) (*Planet, error) {
	return s.repository.ReadOneWithID(id)
}

// FindOne Planet.
func (s *Service) FindOneWithName(name string) (*Planet, error) {
	return s.repository.ReadOneWithName(name)
}

// Update Planet.
func (s *Service) Update(planet *Planet) (*Planet, error) {
	return s.repository.Update(planet)
}

// Remove Planet.
func (s *Service) Remove(id string) error {
	return s.repository.Delete(id)
}
