package storage

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Candy struct {
	ID       uuid.UUID `json:"id"`
	Created  time.Time `json:"created_at"`
	Modified time.Time `json:"modified_at"`
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
