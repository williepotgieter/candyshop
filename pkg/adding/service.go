package adding

import "github.com/google/uuid"

type Repository interface {
	AddCandy(Candy) (uuid.UUID, error)
}

type Service interface {
	AddCandy(Candy) (uuid.UUID, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) AddCandy(c Candy) (uuid.UUID, error) {
	id, err := s.r.AddCandy(c)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
