package planet

type (
	// Service is an interface from which our api module can access our repository.
	Service interface {
		Insert(planet *Planet) (*Planet, error)
		FindAll() (*[]Planet, error)
		FindOne(id string) (*Planet, error)
		Update(planet *Planet) (*Planet, error)
		Remove(id string) error
	}
	service struct {
		repository Repository
	}
)

// NewService is used to create a single instance of the service.
func NewService(r Repository) Service {
	return &service{repository: r}
}

// Insert Planet.
func (s *service) Insert(planet *Planet) (*Planet, error) {
	return s.repository.Create(planet)
}

// FindAll Planets.
func (s *service) FindAll() (*[]Planet, error) {
	return s.repository.ReadAll()
}

// FindOne Planet.
func (s *service) FindOne(id string) (*Planet, error) {
	return s.repository.ReadOne(id)
}

// Update Planet.
func (s *service) Update(planet *Planet) (*Planet, error) {
	return s.repository.Update(planet)
}

// Remove Planet.
func (s *service) Remove(id string) error {
	return s.repository.Delete(id)
}
