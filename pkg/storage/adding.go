package storage

import (
	"log"

	"github.com/google/uuid"
	"github.com/williepotgieter/candyshop/pkg/adding"
)

func (s *Storage) AddCandy(c adding.Candy) (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `
	INSERT INTO candies (category, name, price)
	VALUES ($1, $2, $3)
	RETURNING candy_id`
	err := s.db.QueryRow(sqlStatement, c.Category, c.Name, c.Price).Scan(&id)
	if err != nil {
		log.Println("error while trying to save to DB: ", err)
		return uuid.Nil, err
	}

	return id, nil
}
