package cmd

import (
	"fmt"
	"os"
)

type Storage struct {
	path string
	file *os.File
}

type Quote struct {
	text   string
	viewed int
}

func (q Quote) String() string {
	return fmt.Sprintf("text: %s, viewed: %d", q.text, q.viewed)
}

type Quotes = []Quote
