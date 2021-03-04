package storage

import "github.com/google/uuid"

func (s *Storage) DeleteCandy(id uuid.UUID) error {
	sqlStatement := `
	DELETE FROM candies
	WHERE id=$1`

	_, err := s.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
