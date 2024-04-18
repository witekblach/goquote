package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
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

	//seedDb(file)

	return Storage{path, file}, nil
}

func (s *Storage) GetAllQuotes() ([]Quote, error) {
	var quotes []Quote

	fileContents, err := s.readAll()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		panic(err)
	}

	if len(quotes) == 0 {
		slog.Error("Gotta add a quote first boy!")
	}

	return quotes, nil
}

func (s *Storage) GetQuoteMatching(substringToMatch string) ([]Quote, error) {
	var quotes []Quote

	fileContents, err := s.readAll()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		panic(err)
	}

	var quotesMatching []Quote

	for _, quote := range quotes {
		if strings.Contains(quote.Text, substringToMatch) {
			quotesMatching = append(quotesMatching, quote)
		}
	}

	return quotesMatching, nil
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

func (s *Storage) AddQuote(quote Quote) (bool, error) {
	allQuotes, err := s.GetAllQuotes()
	if err != nil {
		panic(err)
	}

	quotes := append(allQuotes, quote)

	s.writeQuotesToFile(quotes)

	return true, nil
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

func seedDb(file *os.File) {

	q1 := Quote{"Better Software Faster", 0}
	q2 := Quote{"Sooner Safer Happier", 0}
	quotesToDb := Quotes{q1, q2}

	b, err := json.Marshal(quotesToDb)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(b)
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
