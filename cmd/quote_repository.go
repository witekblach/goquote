package cmd

import (
	"errors"
	"log/slog"
	"math/rand"
	"strings"
)

func (s *Storage) GetRandomQuote() (*Quote, error) {
	quotes, err := s.GetAllQuotes()
	if len(quotes) == 0 {
		return nil, err
	}
	randomIndex := rand.Intn(len(quotes))
	randomElement := &quotes[randomIndex]

	return randomElement, nil
}

func (s *Storage) GetAllQuotes() ([]Quote, error) {
	quotes, err := s.readQuotesFromFile()
	if err != nil {
		return []Quote{}, errors.New("reading file. Trouble creating file or file does not contain any quotes")
	}

	return quotes, nil
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
	allQuotes, _ := s.GetAllQuotes()

	quotes := append(allQuotes, Quote{quote, 0})

	s.writeQuotesToFile(quotes)

	return true
}

func (s *Storage) RemoveQuote(quote Quote) bool {
	allQuotes, err := s.GetAllQuotes()
	if err != nil {
		slog.Info("I KNOW WHATS WRONG WITH IT. IT AIN'T GOT NO GAS IN IT")
		return false
	}

	quotes := append(allQuotes, quote)

	s.writeQuotesToFile(quotes)

	return true
}
