package repository

type PackRepository interface {
	GetPackSizes() ([]int, error)
	UpdatePackSizes(newSizes []int) error
}

// In-memory implementation for simplicity
type InMemoryPackRepository struct {
	packSizes []int
}

func NewInMemoryPackRepository(defaultSizes []int) *InMemoryPackRepository {
	return &InMemoryPackRepository{packSizes: defaultSizes}
}

func (r *InMemoryPackRepository) GetPackSizes() ([]int, error) {
	return r.packSizes, nil
}

func (r *InMemoryPackRepository) UpdatePackSizes(newSizes []int) error {
	r.packSizes = newSizes
	return nil
}
