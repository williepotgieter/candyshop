package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Candy struct {
	Id       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
}

func SetupStorage() (*Storage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return &Storage{}, err
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) GetAllCandyNames() ([]Candy, error) {
	var candy Candy
	var candies []Candy

	rows, err := s.db.Query("SELECT * FROM candies;")
	if err != nil {
		return []Candy{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&candy.Id, &candy.Category, &candy.Name, &candy.Price)
		if err != nil {
			return []Candy{}, err
		}

		candies = append(candies, candy)
	}

	err = rows.Err()
	if err != nil {
		return []Candy{}, err
	}

	return candies, nil
}
