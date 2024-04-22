package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetStorage() (Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := homeDir + "/db.json" // straight up db.json in your home directory

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	return Storage{path, file}, nil
}

func (s *Storage) GetRandomQuote() (Quote, error) {
	quotes := s.GetAllQuotes()

	randomIndex := rand.Intn(len(quotes))
	randomElement := quotes[randomIndex]

	return randomElement, nil
}

func (s *Storage) GetAllQuotes() []Quote {
	var quotes []Quote

	s.readQuotesFromFile(&quotes)

	return quotes
}

func (s *Storage) GetQuoteMatching(substringToMatch string) ([]Quote, error) {
	var quotes []Quote

	s.readQuotesFromFile(&quotes)

	var quotesMatching []Quote

	for _, quote := range quotes {
		if strings.Contains(quote.Text, substringToMatch) {
			quotesMatching = append(quotesMatching, quote)
		}
	}

	return quotesMatching, nil
}

func (s *Storage) AddQuote(quote Quote) (bool, error) {
	allQuotes := s.GetAllQuotes()

	quotes := append(allQuotes, quote)

	s.writeQuotesToFile(quotes)

	return true, nil
}

func (s *Storage) readQuotesFromFile(quotes *[]Quote) {
	fileContents, err := s.readAll()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) writeQuotesToFile(quotes []Quote) {
	b, err := json.Marshal(quotes)
	if err != nil {
		panic(err)
	}

	_, err = s.File.Write(b)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) readAll() ([]byte, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from File: %w", err)
	}

	return data, nil
}
