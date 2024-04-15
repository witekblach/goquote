package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func NewStorage(path string) (*Storage, error) {
	println(path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	q1 := Quote{"AAA", 0}
	q2 := Quote{"BBB", 0}
	println(q1.String(), q2.String())
	quotesToDb := Quotes{q1, q2}

	b, err := json.Marshal(quotesToDb)

	_, err = file.Write(b)

	return &Storage{path, file}, nil
}

func (s *Storage) ReadAll() ([]byte, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file: %w", err)
	}

	return data, nil
}

func (s *Storage) GetAllQuotes() ([]Quote, error) {
	var quotes []Quote
	fileContents, err := s.ReadAll()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		panic(err)
	}

	return quotes, nil
}

func (s *Storage) GetRandomQuote() (Quote, error) {
	quotes, err := s.GetAllQuotes()
	if err != nil {
		panic(err)
	}

	randomIndex := rand.Intn(len(quotes))
	randomElement := quotes[randomIndex]
	return randomElement, nil
}

func (s *Storage) AddQuote(quote string) (bool, error) {
	return true, nil
}

func (s *Storage) ExistsQuote(quote string) (bool, error) {
	return true, nil
}
