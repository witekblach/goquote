package cmd

import (
	"errors"
	"log/slog"
	"math/rand"
	"strings"
)

func (s *Storage) GetRandomQuote() (*Quote, error) {
	quotes := s.GetAllQuotes()
	if len(quotes) == 0 {
		return nil, errors.New("no quotes found")
	}
	randomIndex := rand.Intn(len(quotes))
	randomElement := &quotes[randomIndex]

	return randomElement, nil
}

func (s *Storage) GetAllQuotes() []Quote {
	quotes, err := s.readQuotesFromFile()
	if err != nil {
		slog.Error("had problems reading file :<")
	}

	return quotes
}

func (s *Storage) GetQuoteMatching(substringToMatch string) ([]Quote, error) {
	quotes, err := s.readQuotesFromFile()
	if err != nil {
		slog.Error("had problems reading file :<")
	}

	var quotesMatching []Quote
	for _, quote := range quotes {
		if strings.Contains(quote.Text, substringToMatch) {
			quotesMatching = append(quotesMatching, quote)
		}
	}

	return quotesMatching, nil
}

func (s *Storage) AddQuote(quote string) bool {
	allQuotes := s.GetAllQuotes()

	quotes := append(allQuotes, Quote{quote, 0})

	s.writeQuotesToFile(quotes)

	return true
}

func (s *Storage) RemoveQuote(quote Quote) bool {
	allQuotes := s.GetAllQuotes()

	quotes := append(allQuotes, quote)

	s.writeQuotesToFile(quotes)

	return true
}
