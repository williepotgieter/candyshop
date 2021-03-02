package reading

import "github.com/williepotgieter/candyshop/pkg/storage"

type Repository interface {
	GetAllCandyNames() ([]storage.Candy, error)
}

type Service interface {
	GetAllCandyNames() ([]storage.Candy, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllCandyNames() ([]storage.Candy, error) {
	cs, err := s.r.GetAllCandyNames()
	if err != nil {
		return nil, err
	}

	return cs, nil
}
