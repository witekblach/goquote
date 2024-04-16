package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func NewStorage(path string) (Storage, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	//seedDb(file)

	return Storage{path, file}, nil
}

func seedDb(file *os.File) {

	q1 := Quote{"AAA", 0}
	q2 := Quote{"BBB", 0}
	q3 := Quote{"CCC", 0}
	q4 := Quote{"DDD", 0}
	q5 := Quote{"EEE", 0}
	q6 := Quote{"FFF", 0}
	quotesToDb := Quotes{q1, q2, q3, q4, q5, q6}

	b, err := json.Marshal(quotesToDb)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(b)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) ReadAll() ([]byte, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from File: %w", err)
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

func (s *Storage) GetQuoteMatching(sToMatch string) ([]Quote, error) {
	var quotes []Quote
	var result []Quote

	fileContents, err := s.ReadAll()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContents, &quotes)
	if err != nil {
		panic(err)
	}

	for _, obj := range quotes {
		if strings.Contains(obj.Text, sToMatch) {
			result = append(result, obj)
		}
	}

	return result, nil
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
