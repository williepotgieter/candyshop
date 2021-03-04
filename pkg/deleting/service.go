package deleting

import "github.com/google/uuid"

type Repository interface {
	DeleteCandy(id uuid.UUID) error
}

type Service interface {
	DeleteCandy(id uuid.UUID) error
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) DeleteCandy(id uuid.UUID) error {
	err := s.r.DeleteCandy(id)
	if err != nil {
		return err
	}

	return nil
}
