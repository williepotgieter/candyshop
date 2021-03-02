package storage

func (s *Storage) GetAllCandyNames() ([]Candy, error) {
	var candy Candy
	var candies []Candy

	rows, err := s.db.Query("SELECT * FROM candies;")
	if err != nil {
		return []Candy{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&candy.ID, &candy.Created, &candy.Modified, &candy.Category, &candy.Name, &candy.Price)
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
