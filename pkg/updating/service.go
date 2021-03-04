package updating

import "github.com/google/uuid"

type Repository interface {
	UpdateCandyCategory(uuid.UUID, string) error
	UpdateCandyName(uuid.UUID, string) error
	UpdateCandyPrice(uuid.UUID, float32) error
}

type Service interface {
	UpdateCandyCategory(uuid.UUID, string) error
	UpdateCandyName(uuid.UUID, string) error
	UpdateCandyPrice(uuid.UUID, float32) error
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) UpdateCandyCategory(id uuid.UUID, c string) error {
	err := s.r.UpdateCandyCategory(id, c)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCandyName(id uuid.UUID, n string) error {
	err := s.r.UpdateCandyName(id, n)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCandyPrice(id uuid.UUID, p float32) error {
	err := s.r.UpdateCandyPrice(id, p)
	if err != nil {
		return err
	}

	return nil
}
