package cmd

import (
	"encoding/json"
	"os"
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

func (s *Storage) readQuotesFromFile() ([]Quote, error) {
	fileContents, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
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

func (s *Storage) updateQuoteInFile(quote Quote) {
	quotes, err := s.readQuotesFromFile()
	if err != nil {
		return
	}

	for i, q := range quotes {
		if q.Text == quote.Text {
			quotes[i].Viewed = quote.Viewed
		}
	}

	s.writeQuotesToFile(quotes)
}
