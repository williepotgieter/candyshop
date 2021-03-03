package storage

import (
	"github.com/google/uuid"
)

func (s *Storage) UpdateCandyCategory(id uuid.UUID, category string) error {
	sqlStatement := `
	UPDATE candies
	SET category = $2, modified_at = NOW()
	WHERE id = $1;`
	_, err := s.db.Exec(sqlStatement, id, category)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateCandyName(id uuid.UUID, name string) error {
	sqlStatement := `
	UPDATE candies
	SET name = $2, modified_at = NOW()
	WHERE id = $1;`
	_, err := s.db.Exec(sqlStatement, id, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateCandyPrice(id uuid.UUID, price float32) error {
	sqlStatement := `
	UPDATE candies
	SET price = $2, modified_at = NOW()
	WHERE id = $1;`
	_, err := s.db.Exec(sqlStatement, id, price)
	if err != nil {
		return err
	}

	return nil
}
