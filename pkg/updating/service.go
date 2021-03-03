package updating

import "github.com/google/uuid"

type Repository interface {
	UpdateCandyCategory(uuid.UUID, string) error
	UpdateCandyName(uuid.UUID, string) error
	UpdateCandyPrice(uuid.UUID, float32) error
}

type Service interface {
	UpdateCandyCategory(uuid.UUID, CandyCategory) error
	UpdateCandyName(uuid.UUID, CandyName) error
	UpdateCandyPrice(uuid.UUID, CandyPrice) error
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) UpdateCandyCategory(id uuid.UUID, c CandyCategory) error {
	err := s.r.UpdateCandyCategory(id, c.Category)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCandyName(id uuid.UUID, c CandyName) error {
	err := s.r.UpdateCandyName(id, c.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCandyPrice(id uuid.UUID, c CandyPrice) error {
	err := s.r.UpdateCandyPrice(id, c.Price)
	if err != nil {
		return err
	}

	return nil
}
