package cmd

import (
	"fmt"
	"os"
)

type Storage struct {
	Path string
	File *os.File
}

type Quote struct {
	Text   string
	Viewed int
}

func (q Quote) String() string {
	return fmt.Sprintf("Text: %s, Viewed: %d", q.Text, q.Viewed)
}

type Quotes = []Quote
